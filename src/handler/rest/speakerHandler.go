package rest

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *rest) GetAllSpeakers(c *gin.Context) {
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

	var speakerId []uint
	if filter.Date!=""{
		id,exist:=c.Get("user")
		if exist==false{
			h.ErrorResponse(c,http.StatusUnauthorized,errors.New("Token is not found"),id)
			return
		}

		user,statusCode,err:=h.uc.User.GetById(id)
		if err!=nil{
			h.ErrorResponse(c,statusCode,err,"User is not found")
			return
		}

		if user.Role=="free-user"{
			h.ErrorResponse(c,http.StatusUnauthorized,errors.New("Free user can't access this page"),"Upgrade your account to premium to access this page")
			return
		}
		
		speakerId,statusCode,err:=h.uc.Schedule.GetSchedulesByDate(filter)
		if err!=nil{
			h.ErrorResponse(c,statusCode,err,speakerId)
			return 
		}
		result, statusCode, err:= h.uc.Speaker.FindAll(filter,pagination,speakerId)
		if err!=nil{
			h.ErrorResponse(c,statusCode,err,result)
			return 
		}

		h.SuccessResponse(c, statusCode,"Querying Speakers data successfully", result)
		return
	}

	result, statusCode, err:= h.uc.Speaker.FindAll(filter,pagination,speakerId)
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
		h.ErrorResponse(c,statusCode,err,"Failed to querying speaker's data")
		return 
	}
	h.SuccessResponse(c, statusCode,"Querying Speakers data details successfully", result)
}
