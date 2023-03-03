package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type ReviewUsecase interface {
	GetReviewsByUserId(id string, pagination entity.Pagination)(interface{},int,error)
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
	reviews,pg,err := h.reviewRepository.GetByUserId(id,pagination)
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