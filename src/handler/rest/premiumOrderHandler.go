package rest

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
)

func (h *rest) OrderPremiumAccount(c *gin.Context) {
	var body entity.InputPremiumOrder
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

	result,statusCode,err:=h.uc.PremiumOrder.UpgradePremium(user.ID,body)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}

	h.SuccessResponse(c,statusCode,"Create premium account order successfully",result)
}

func (h *rest) CheckPremiumOrderTransaction(c *gin.Context) {
	var body entity.CheckTransaction
	if err:=h.BindBody(c,&body);err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	result,_,statusCode,err:=h.uc.PremiumOrder.UpdatePremiumOrderStatus(body)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	h.SuccessResponse(c,statusCode,"Updating premium order status successfully",result)
}