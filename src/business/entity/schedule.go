package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Schedules struct {
	gorm.Model
	Date 		datatypes.Date `gorm:"type:date" json:"date"`
	Time 		datatypes.Time `gorm:"type:time" json:"time"`
	SpeakerID	uint
	Speaker		Speakers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}