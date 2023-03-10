package rest

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *rest) CreateOrder(c *gin.Context) {
	var body entity.OrderInput
	if err:=h.BindBody(c,&body);err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	userId,exist:=c.Get("user")
	if exist==false{
		h.ErrorResponse(c,http.StatusUnauthorized,errors.New("User ID is not found in token"),"User ID doesn't exist")
		return
	}

	user,statusCode,err:=h.uc.User.GetById(userId)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,user)
		return 
	}

	speaker,statusCode,err:=h.uc.Speaker.GetById(strconv.Itoa(body.SpeakerID))
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,user)
		return 
	}

	rundown, err := c.FormFile("rundown")
	if err != nil {
		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to read rundown file")
		return
	}

	script, err := c.FormFile("script")
	if err != nil {
		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to read script file")
		return
	}

	result,statusCode,err:=h.uc.Order.CreateTransaction(speaker,user,body,rundown,script)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	h.SuccessResponse(c,statusCode,"Create order successfully",result)
}

func (h *rest) CheckOrderTransaction(c *gin.Context) {
	var body entity.CheckTransaction
	if err:=h.BindBody(c,&body);err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	result,_,statusCode,err:=h.uc.Order.UpdateOrderStatus(body)
	if err!=nil{
		resultPremiumOrder,premiumOrder,statusCodePremium,err:=h.uc.PremiumOrder.UpdatePremiumOrderStatus(body)
		if err!=nil{
			h.ErrorResponse(c,statusCodePremium,err,resultPremiumOrder)
			return
		}

		resultPremium,statusCodePremium,err :=h.uc.User.UpdateUserPremium(premiumOrder.UserID,premiumOrder.Quantity)
		if err!=nil{
			h.ErrorResponse(c,statusCodePremium,err,resultPremiumOrder)
			return
		}

		h.SuccessResponse(c,statusCodePremium,"Upgrade Premium User Successfully",resultPremium)
		return
	}

	h.SuccessResponse(c,statusCode,"Updating order status successfully",result)
}

func (h *rest) GetOrderHistory(c *gin.Context){
	userId,exist:=c.Get("user")
	if exist==false{
		h.ErrorResponse(c,http.StatusUnauthorized,errors.New("User ID is not found in token"),"User ID doesn't exist")
		return
	}

	var pagination entity.Pagination
	err:=c.ShouldBindWith(&pagination,binding.Query)
	if err!=nil{
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read pagination parameters")
		return
	}

	result,statusCode,err:=h.uc.Order.GetAllOrders(userId,pagination)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
	}

	h.SuccessResponse(c,statusCode,"Querying order's history successful",result)
}

func (h *rest) CreateTransactionByOrderCode(c *gin.Context) {
	var body entity.PayOrder
	if err:=h.BindBody(c,&body);err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	order,err:=h.uc.Order.GetOrderByOrderCode(body.OrderCode)
	if err!=nil{
		h.ErrorResponse(c,http.StatusNotFound,err,nil)
		return
	}

	speaker,statusCode,err:=h.uc.Speaker.GetById(strconv.FormatUint(uint64(order.SpeakerID), 10))
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,speaker)
		return
	}

	result,statusCode,err:=h.uc.Order.CreateTransactionByOrderCode(body.OrderCode,speaker)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	h.SuccessResponse(c,statusCode,"Create order transaction successfully",result)
}

