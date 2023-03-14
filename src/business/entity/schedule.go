package entity

import (
	"time"
	"gorm.io/gorm"
)

type Schedules struct {
	gorm.Model
	TimeStart 	time.Time `json:"timeStart"`
	TimeEnd 	time.Time `json:"timeEnd"`
	Duration    int `json:"duration"`
	SpeakerID	uint `json:"speakerId"`
	Speaker		Speakers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ScheduleResponse struct{
	TimeStart 	time.Time `json:"timeStart"`
	TimeEnd 	time.Time `json:"timeEnd"`
	Duration    int `json:"duration"`
	SpeakerID	uint `json:"speakerId"`
}