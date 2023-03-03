package rest

import (
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