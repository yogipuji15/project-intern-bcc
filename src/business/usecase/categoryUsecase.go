package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/repository"
)

type CategoryUsecase interface {
	FindAll()(interface{},int,error)
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(r repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: r,
	}
}

func (h *categoryUsecase) FindAll()(interface{},int,error){
	categories,err:=h.categoryRepository.FindAll()
	if err!=nil{
		return "Categories data not found",http.StatusNotFound,err
	}
	return categories,http.StatusOK,nil
}