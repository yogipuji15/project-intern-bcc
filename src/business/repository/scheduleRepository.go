package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule entity.Schedules) (entity.Schedules, error)
	GetAllBySpeakerId(speakerId string, month string) ([]entity.Schedules, error)
	GetById(Id string) (entity.Schedules, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (h *scheduleRepository) Create(schedule entity.Schedules) (entity.Schedules, error) {
	err := h.db.Create(&schedule).Error
	return schedule, err
}

func (h *scheduleRepository) GetAllBySpeakerId(speakerId string, month string) ([]entity.Schedules, error) {
	var schedules []entity.Schedules
	err := h.db.Where("time_start BETWEEN ? AND ?", month+"-01 00:00:00", month+"-28 23:59:59").Or("time_start BETWEEN ? AND ?", month+"-01 00:00:00", month+"-30 23:59:59").Or("time_start BETWEEN ? AND ?", month+"-01 00:00:00", month+"-31 23:59:59").Or("time_start BETWEEN ? AND ?", month+"-01 00:00:00", month+"-29 23:59:59").Find(&schedules, "speaker_id = ?", speakerId).Error

	return schedules, err
}

func (h *scheduleRepository) GetById(Id string) (entity.Schedules, error) {
	var schedule entity.Schedules
	err := h.db.Where("id = ?", Id).Find(&schedule).Error
	return schedule, err
}
