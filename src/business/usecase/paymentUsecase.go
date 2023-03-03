package usecase

import "project-intern-bcc/src/business/repository"

type PaymentUsecase interface {
}

type paymentUsecase struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepository: r,
	}
}