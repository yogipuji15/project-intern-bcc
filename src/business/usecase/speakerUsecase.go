package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
	"strconv"
)

type SpeakerUsecase interface {
	FindAll(filter entity.FilterParam, pagination entity.Pagination)(interface{},int,error)
	GetById(speakerId string)(entity.Speakers,int,error)
	UpdateRating(speakerId uint, totalRating float32)(interface{},int,error)
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

func (h *speakerUsecase) GetById(speakerId string)(entity.Speakers,int,error){
	speaker,err:=h.speakerRepository.GetById(speakerId)
	if err!=nil{
		return speaker,http.StatusNotFound,err
	}

	return speaker,http.StatusOK,err
}

func (h *speakerUsecase) UpdateRating(speakerId uint, totalRating float32)(interface{},int,error){
	speaker,err:=h.speakerRepository.GetById(strconv. FormatUint(uint64(speakerId),10))
	if err!=nil{
		return "Failed to qurying speaker's data",http.StatusNotFound,err
	}

	speaker.TotalReviews=speaker.TotalReviews+1
	var rating float32
	rating=totalRating/float32(speaker.TotalReviews)
	speaker.Rating=float32(rating)

	err=h.speakerRepository.UpdateRating(speaker)
	if err!=nil{
		return "Failed to update speaker's data",http.StatusInternalServerError,err
	}
	return "Update speaker's data successfully",http.StatusOK,nil
}