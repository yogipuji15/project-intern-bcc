package rest

import "github.com/gin-gonic/gin"

func (h *rest) GetAllSchedule(c *gin.Context) {
	speakerId:= c.Param("id")
	month:= c.Query("month")

	schedules,statusCode,err:=h.uc.Schedule.GetAll(speakerId,month)
	if err!=nil{
		h.ErrorResponse(c,statusCode,err,schedules)
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