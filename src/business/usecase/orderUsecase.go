package usecase

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderUsecase interface {
	CreateTransaction(speaker entity.Speakers, user entity.UserResponse,orderInput entity.OrderInput,rundown *multipart.FileHeader,script *multipart.FileHeader) (interface{},int,error)
	UpdateOrderStatus(body entity.CheckTransaction) (interface{},entity.Orders,int,error)
}

type orderUsecase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(r repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepository: r,
	}
}

func (h *orderUsecase) CreateTransaction(speaker entity.Speakers, user entity.UserResponse,orderInput entity.OrderInput,rundown *multipart.FileHeader,script *multipart.FileHeader) (interface{},int,error){
	orderId:=uuid.NewString()
	
	var paymentId int
	if orderInput.PaymentType=="bank-bca"{
		paymentId=1
	}else if orderInput.PaymentType=="bank-bri"{
		paymentId=2
	}else if orderInput.PaymentType=="gopay"{
		paymentId=3
	}else{
		return "Failed create order",http.StatusBadRequest,errors.New("payment type not found")
	}

	bookTimeStart,err:=time.Parse("2006-01-02 15:04:05 MST", orderInput.BookTime+" WIB")
	if err!=nil{
		return "Failed to parsing time",http.StatusBadRequest,err
	}
	
	var status string
	if user.Role=="premium-user"{
		status="Waiting For Payment"
	}else{
		status="Waiting To Approve"
	}
	order:=entity.Orders{
		OrderCode 		:orderId,
		EventName 		:orderInput.EventName,
		BookTimeStart 	:bookTimeStart,
		BookTimeEnd		:bookTimeStart.Add(time.Hour * time.Duration(orderInput.Duration)),
		Description 	:orderInput.Description,
		Duration 		:orderInput.Duration,
		TotalPrice  	:speaker.Price*orderInput.Duration,
		UserID 			:user.ID,
		Status			:status,
		SpeakerID 		:speaker.ID,
		PaymentID 		:uint(paymentId),
	}
	
	order,resp,err:=h.orderRepository.Create(speaker,user,order,rundown,script)
	if err!=nil{
		return "failed to create booking transaction",http.StatusInternalServerError,err
	}
	
	if resp!=nil{
		return h.MidtransTransactionResponse(resp,order,speaker),http.StatusOK,nil
	}
	return h.OrderResponse(order,speaker),http.StatusOK,nil
}


func (h *orderUsecase) MidtransTransactionResponse(resp *coreapi.ChargeResponse,order entity.Orders,speaker entity.Speakers) entity.MidtransTransactionResponse{
	var payment interface{}
	if resp.PaymentType=="bank_transfer"{
		payment=resp.VaNumbers
	}else{
		payment=resp.Actions
	}
	return entity.MidtransTransactionResponse{
		TransactionID	: resp.TransactionID,
		TransactionTime : resp.TransactionTime,
		OrderID 	  	: resp.OrderID,
		PaymentResult 	: payment,
		TotalPrice  	: resp.GrossAmount,
		Status	 		: resp.TransactionStatus,
		Order			: h.OrderResponse(order,speaker),
	}
}

func (h *orderUsecase) OrderResponse(order entity.Orders,speaker entity.Speakers) entity.OrderResponse{
	return entity.OrderResponse{
		OrderCode 		:order.OrderCode,
		EventName 		:order.EventName,
		Status 			:order.Status,
		BookTimeStart 	:order.BookTimeStart,
		BookTimeEnd 	:order.BookTimeEnd,
		Description 	:order.Description,
		Duration 		:order.Duration,
		TotalPrice 		:order.TotalPrice,
		Speaker 		:order.Speaker,
	}
}

func (h *orderUsecase) UpdateOrderStatus(body entity.CheckTransaction) (interface{},entity.Orders,int,error){
	var order entity.Orders
	
	mySignature:=body.OrderID+body.StatusCode+body.GrossAmount+os.Getenv("SERVER_KEY")

	if body.SignatureKey !=h.Hash512(mySignature){
		return "Signature key is invalid",order,http.StatusUnauthorized,errors.New("Signature key is invalid")
	}

	order,err:=h.orderRepository.FindByOrderCode(body.OrderID)
	if err!=nil{
		return "Failed to querying order data",order,http.StatusNotFound,err
	}

	if body.TransactionStatus=="settlement"{
		order.Status="Success Payment"
	}

	err=h.orderRepository.Update(order)
	if err!=nil{
		return "Failed to update order's status data",order,http.StatusInternalServerError,err
	}

	return "Updating order status successfully",order,http.StatusOK,err
}

func (h *orderUsecase) Hash512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	pass := hex.EncodeToString(hash.Sum(nil))
	return pass
}