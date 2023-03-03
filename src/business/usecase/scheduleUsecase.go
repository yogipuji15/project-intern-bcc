package usecase

import "project-intern-bcc/src/business/repository"

type ScheduleUsecase interface {
}

type scheduleUsecase struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleUsecase(r repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepository: r,
	}
}