package usecase

import "project-intern-bcc/src/business/repository"

type OrderUsecase interface {
}

type orderUsecase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(r repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepository: r,
	}
}