package usecase

import (
	"errors"
	"net/http"
	"os"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
	"project-intern-bcc/src/lib/auth"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PremiumOrderUsecase interface {
	UpgradePremium(userId uint, input entity.InputPremiumOrder)(interface{},int,error)
	UpdatePremiumOrderStatus(body entity.CheckTransaction) (interface{},entity.PremiumOrders,int,error)
}

type premiumOrderUsecase struct {
	premiumOrderRepository repository.PremiumOrderRepository
	auth auth.AuthInterface
}

func NewPremiumOrderUsecase(r repository.PremiumOrderRepository, auth auth.AuthInterface) PremiumOrderUsecase {
	return &premiumOrderUsecase{
		premiumOrderRepository: r,
		auth :auth,
	}
}

func (h *premiumOrderUsecase) UpgradePremium(userId uint, input entity.InputPremiumOrder)(interface{},int,error){
	var price int
	if input.Month==12{
		price=35000
	}else if input.Month==1{
		price=32000
	}else{
		return "Premium package is not available",http.StatusBadRequest,errors.New("Premium package is not available")
	}

	var paymentId int
	if input.PaymentType=="bank-bca"{
		paymentId=1
	}else if input.PaymentType=="bank-bri"{
		paymentId=2
	}else if input.PaymentType=="gopay"{
		paymentId=3
	}else{
		return "Failed create order",http.StatusBadRequest,errors.New("payment type not found")
	}

	premiumOrder := entity.PremiumOrders{
		OrderCode 	: uuid.NewString(),
		Status 	  	: false,
		Quantity 	: input.Month,
		TotalPrice 	: price,
		UserID 		: userId,
		PaymentID 	: uint(paymentId),
	}
	
	premiumOrder,resp,err:=h.premiumOrderRepository.Create(premiumOrder)
	if err!=nil{
		return "Failed to create premium account order",http.StatusInternalServerError,err
	}

	response:=h.MidtransTransactionResponse(resp,premiumOrder)


	return response,http.StatusOK,nil
}

func (h *premiumOrderUsecase) MidtransTransactionResponse(resp *coreapi.ChargeResponse,order entity.PremiumOrders) entity.TransactionResponse{
	var payment interface{}
	if resp.PaymentType=="bank_transfer"{
		payment=resp.VaNumbers
	}else{
		payment=resp.Actions
	}
	return entity.TransactionResponse{
		TransactionID	: resp.TransactionID,
		TransactionTime : resp.TransactionTime,
		OrderID 	  	: resp.OrderID,
		PaymentResult 	: payment,
		TotalPrice  	: resp.GrossAmount,
		Status	 		: resp.TransactionStatus,
		Order			: order,
	}
}

func (h *premiumOrderUsecase) UpdatePremiumOrderStatus(body entity.CheckTransaction) (interface{},entity.PremiumOrders,int,error){
	var order entity.PremiumOrders
	
	mySignature:=body.OrderID+body.StatusCode+body.GrossAmount+os.Getenv("SERVER_KEY")

	if body.SignatureKey !=h.auth.Hash512(mySignature){
		return "Signature key is invalid",order,http.StatusUnauthorized,errors.New("Signature key is invalid")
	}

	order,err:=h.premiumOrderRepository.FindByOrderCode(body.OrderID)
	if err!=nil{
		return "Failed to querying premium order's data",order,http.StatusNotFound,err
	}

	if body.TransactionStatus=="settlement"{
		order.Status=true
	}else if body.TransactionStatus=="expired"{
		order.Status=false
	}else if body.TransactionStatus=="failure"{
		order.Status=false
	}

	err=h.premiumOrderRepository.Update(order)
	if err!=nil{
		return "Failed to update order's status data",order,http.StatusInternalServerError,err
	}

	return "Updating order status successfully",order,http.StatusOK,err
}