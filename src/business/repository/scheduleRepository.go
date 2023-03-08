package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule entity.Schedules)(entity.Schedules,error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db:db}
}

func (h *scheduleRepository) Create(schedule entity.Schedules)(entity.Schedules,error){
	err:=h.db.Create(&schedule).Error
	return schedule,err
}