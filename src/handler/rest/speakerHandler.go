package rest

import (
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *rest) GetAllSpeakers(c *gin.Context) {
	// category:=c.Query("category")
	// keyword:=c.Query("keyword")
	var filter entity.FilterParam
	err:=c.ShouldBindWith(&filter,binding.Query)
	if err!=nil{
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read filter parameters")
		return
	}

	var pagination entity.Pagination
	err=c.ShouldBindWith(&pagination,binding.Query)
	if err!=nil{
		h.ErrorResponse(c, http.StatusBadRequest,err,"Failed to read pagination parameters")
		return
	}

	result, statusCode, err:= h.uc.Speaker.FindAll(filter,pagination)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}
	h.SuccessResponse(c, statusCode,"Querying Speakers data successfully", result)
}

func (h *rest) GetSpeakerById(c *gin.Context){
	speakerId:= c.Param("id")

	result, statusCode, err:= h.uc.Speaker.GetById(speakerId)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,result)
		return 
	}
	h.SuccessResponse(c, statusCode,"Querying Speakers data details successfully", result)
}
