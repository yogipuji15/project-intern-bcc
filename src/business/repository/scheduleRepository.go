package repository

import "gorm.io/gorm"

type ScheduleRepository interface {
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db:db}
}