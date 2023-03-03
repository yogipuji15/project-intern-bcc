package usecase

import "project-intern-bcc/src/business/repository"

type PremiumOrderUsecase interface {
}

type premiumOrderUsecase struct {
	premiumOrderRepository repository.PremiumOrderRepository
}

func NewPremiumOrderUsecase(r repository.PremiumOrderRepository) PremiumOrderUsecase {
	return &premiumOrderUsecase{
		premiumOrderRepository: r,
	}
}