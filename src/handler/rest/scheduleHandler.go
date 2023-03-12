package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *rest) GetAllSchedule(c *gin.Context) {
	speakerId:= c.Param("id")
	month:= c.Query("month")
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
	
	schedules,statusCode,err:=h.uc.Schedule.GetAll(speakerId,month,user)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,schedules)
		return
	}

	h.SuccessResponse(c,statusCode,"Querying speaker's schedules data successfully",schedules)
}

func (h *rest) GetScheduleById(c *gin.Context) {
	Id:= c.Param("id")

	schedule,statusCode,err:=h.uc.Schedule.GetById(Id)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,"Failed to querying speaker's schedule data")
	}

	h.SuccessResponse(c,statusCode,"Querying speaker's schedule data successfully",schedule)
}