package usecase

import (
	"project-intern-bcc/src/business/repository"
)

type CompanyCategoryUsecase interface {
}

type companyCategoryUsecase struct {
	companyCategoryRepository repository.CompanyCategoryRepository
}

func NewCompanyCategoryUsecase(r repository.CompanyCategoryRepository) CompanyCategoryUsecase {
	return &companyCategoryUsecase{
		companyCategoryRepository: r,
	}
}