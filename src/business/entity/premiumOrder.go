package entity

import "gorm.io/gorm"

type PremiumOrders struct {
	gorm.Model
	OrderCode string `gorm:"type:varchar(255);unique" json:"orderCcode"`
	Status bool `gorm:"type:bool;default:null" json:"status"`
	Quantity int `gorm:"type:int" json:"quantity"`
	TotalPrice int `gorm:"type:int" json:"totalPrice"`
	UserID uint
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentID uint
	Payment Payments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}