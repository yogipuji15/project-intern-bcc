package entity

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	PaymentType string `gorm:"type:varchar(50);unique" json:"paymentType"`
}