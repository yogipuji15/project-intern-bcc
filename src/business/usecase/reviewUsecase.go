package usecase

import (
	"errors"
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
	"time"

	"gorm.io/datatypes"
)

type ReviewUsecase interface {
	GetReviewsByUserId(id string, pagination entity.Pagination)(interface{},int,error)
	CreateReview(userId uint, reviewInput entity.PostReview)(interface{},int,error,float32)
}

type reviewUsecase struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewUsecase(r repository.ReviewRepository) ReviewUsecase {
	return &reviewUsecase{
		reviewRepository: r,
	}
}

func (h *reviewUsecase) GetReviewsByUserId(id string, pagination entity.Pagination)(interface{},int,error){
	reviews,pg,err := h.reviewRepository.GetBySpeakerId(id,pagination)
	if err != nil{
		return "Reviews not found",http.StatusNotFound,err
	}

	var reviewResponse []entity.ReviewResponse
	for _,r := range reviews{
		reviewResponse=append(reviewResponse, h.ConvertToReviewResponse(r))
	}

	result:= entity.ReviewsResponse{
		Pagination: *pg,
		Reviews: reviewResponse,
	}

	return result,http.StatusOK,err
}

func (h *reviewUsecase) ConvertToReviewResponse(review entity.Reviews)(entity.ReviewResponse){
	return entity.ReviewResponse{
		Star: review.Star,
		Review: review.Review,
		Date: review.Date,
		Username: review.User.Username,
	}
}


func (h *reviewUsecase) CreateReview(userId uint, reviewInput entity.PostReview)(interface{},int,error,float32){
	var totalRating float32
	totalRating=0
	_,err:=h.reviewRepository.GetAllBySpeakerAndUserId(userId,reviewInput.SpeakerId)
	if err==nil{
		return "Failed to create review",http.StatusBadRequest,errors.New("User have written review for this speaker"),totalRating
	}
	

	review := entity.Reviews{
		Review 		: reviewInput.Review, 
		Star 		: reviewInput.Star,
		Date 		: datatypes.Date(time.Now()),
		UserID 		: userId,
		SpeakerID   : reviewInput.SpeakerId,
	}

	speakerReviews,err:=h.reviewRepository.GetAll(reviewInput.SpeakerId)
	if err!=nil{
		return "Failed querying speaker's reviews data",http.StatusNotFound,err,totalRating
	}

	totalRating=totalRating+float32(reviewInput.Star)
	for _,r:=range speakerReviews{
		totalRating=totalRating+float32(r.Star)
	}

	review,err=h.reviewRepository.Create(review)
	if err!=nil{
		return "Failed to create review",http.StatusInternalServerError,err,totalRating
	}

	return review,http.StatusOK,err,totalRating
}