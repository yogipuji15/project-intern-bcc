package usecase

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type ScheduleUsecase interface {
	Create(order entity.Orders)(interface{},int,error)
	GetAll(speakerId string,month string, user entity.UserResponse)(interface{},int,error)
	GetById(Id string)(entity.Schedules,int,error)
	GetSchedulesByDate(filter entity.FilterParam)([]uint,int,error)
}

type scheduleUsecase struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleUsecase(r repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepository: r,
	}
}

func (h *scheduleUsecase) Create(order entity.Orders)(interface{},int,error){
	schedule:= entity.Schedules{
		TimeStart: order.BookTimeStart,
		TimeEnd: order.BookTimeEnd,
		SpeakerID: order.SpeakerID,
	}

	schedule,err:=h.scheduleRepository.Create(schedule)
	if err!=nil{
		return "Failed to create speaker's schedule",http.StatusInternalServerError,err
	}

	return schedule,http.StatusOK,nil
}

func (h *scheduleUsecase) GetAll(speakerId string,month string, user entity.UserResponse)(interface{},int,error){
	if user.Role!="premium-user"{
		return "Upgrade your account to see speaker's schedule",http.StatusUnauthorized,errors.New("Upgrade your account to premium to access this page")
	}
	
	schedules,err:=h.scheduleRepository.GetAllBySpeakerId(speakerId,month)
	if err != nil{
		return "Failed to querying speaker's schedules data",http.StatusNotFound,err
	}
	
	var schedulesResponse []entity.ScheduleResponse
	for _,s:=range schedules{
		schedulesResponse=append(schedulesResponse, h.ConvertToScheduleResponse(s))
	}

	return schedulesResponse,http.StatusOK,nil
}

func (h *scheduleUsecase) GetById(Id string)(entity.Schedules,int,error){
	schedule,err:=h.scheduleRepository.GetById(Id)
	if err!=nil{
		return schedule,http.StatusNotFound,err
	}

	return schedule,http.StatusOK,nil
}

func (h *scheduleUsecase) ConvertToScheduleResponse(schedules entity.Schedules)(entity.ScheduleResponse){
	return entity.ScheduleResponse{
		TimeStart 	:schedules.TimeStart,
		TimeEnd 	:schedules.TimeEnd,
		Duration    :schedules.Duration,
		SpeakerID	:schedules.SpeakerID,
	}
}

func (h *scheduleUsecase) GetSchedulesByDate(filter entity.FilterParam)([]uint,int,error){
	schedule,err:=h.scheduleRepository.GetByDate(filter)
	if err!=nil{
		return nil,http.StatusNotFound,err
	}

	
	var speakerId []uint
	for _,s:=range schedule{
		speakerId=append(speakerId, s.SpeakerID)
	}
	
	return speakerId,http.StatusOK,nil
}