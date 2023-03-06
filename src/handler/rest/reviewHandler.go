package rest

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *rest) GetReviewsByUserId(c *gin.Context) {
	id:=c.Param("id")
	var pagination entity.Pagination
	err:=c.ShouldBindWith(&pagination,binding.Query)
	if err!=nil{
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read parameters")
		return
	}

	result, statusCode, err:= h.uc.Review.GetReviewsByUserId(id,pagination)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}

	h.SuccessResponse(c,statusCode,"Get speaker's reviews successfully",result)
}

func (h *rest) PostReview(c *gin.Context) {
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

	var body entity.PostReview
	if err := c.BindJSON(&body); err != nil {
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read body")
		return
	}

	result,statusCode,err,totalRating:=h.uc.Review.CreateReview(user.ID,body)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return
	}

	speaker,statusCode,err:=h.uc.Speaker.UpdateRating(body.SpeakerId,totalRating)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,speaker)
		return
	}
	
	h.SuccessResponse(c,statusCode,"Create speaker's review successfully",result)
}