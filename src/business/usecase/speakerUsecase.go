package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type SpeakerUsecase interface {
	FindAll(filter entity.FilterParam, pagination entity.Pagination)(interface{},int,error)
	GetById(speakerId string)(interface{},int,error)
}

type speakerUsecase struct {
	speakerRepository repository.SpeakerRepository
}

func NewSpeakerUsecase(r repository.SpeakerRepository) SpeakerUsecase {
	return &speakerUsecase{
		speakerRepository: r,
	}
}

func (h *speakerUsecase) FindAll(filter entity.FilterParam, pagination entity.Pagination)(interface{},int,error){
	speakers,pg,err:=h.speakerRepository.FindAll(filter, pagination)
	if err!=nil{
		return "Failed to Querying Speakers Data",http.StatusNotFound,err
	}

	result:=entity.SpeakersResponse{
		Speakers: speakers,
		Pagination: *pg,
	}

	return result,http.StatusOK,nil
}

func (h *speakerUsecase) GetById(speakerId string)(interface{},int,error){
	speaker,err:=h.speakerRepository.GetById(speakerId)
	if err!=nil{
		return "Failed to querying speaker's data",http.StatusNotFound,err
	}

	return speaker,http.StatusOK,err
}