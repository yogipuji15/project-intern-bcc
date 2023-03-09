package rest

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
)

func (h *rest) UploadProposal(c *gin.Context) {
	var body entity.InputProposal
	err:=h.BindBody(c,&body)
	if err!=nil{
		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to read body")
		return
	}
	proposal, err := c.FormFile("proposal")
	if err != nil {
		h.ErrorResponse(c,http.StatusBadRequest,err,"Failed to read rundown file")
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

	result,statusCode,err:=h.uc.Proposal.Create(user,body,proposal)
	if err!= nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	h.SuccessResponse(c,statusCode,"Apply proposal successfully",result)
}