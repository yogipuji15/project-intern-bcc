package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderCode string `gorm:"type:varchar(255);unique" json:"orderCcode"`
	Status bool `gorm:"type:bool;default:null" json:"status"`
	BookDate datatypes.Date `gorm:"type:date" json:"bookDate"`
	BookTime datatypes.Time `gorm:"type:time" json:"bookTime"`
	Quantity int `gorm:"type:int" json:"quantity"`
	TotalPrice int `gorm:"type:int" json:"totalPrice"`
	UserID uint
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpeakerID uint
	Speaker Speakers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentID uint
	Payment Payments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}