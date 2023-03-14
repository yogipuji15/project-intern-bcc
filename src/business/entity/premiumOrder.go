package entity

import "gorm.io/gorm"

type PremiumOrders struct {
	gorm.Model
	OrderCode string `gorm:"type:varchar(255);unique" json:"orderCcode"`
	Status bool `gorm:"type:bool" json:"status"`
	Quantity int `gorm:"type:int" json:"quantity"`
	TotalPrice int `gorm:"type:int" json:"totalPrice"`
	UserID uint `json:"userId"`
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentID uint `json:"paymentId"`
	Payment Payments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type InputPremiumOrder struct{
	Month int `json:"month" binding:"required"`
	PaymentType string `form:"paymentType" binding:"required"`
}

type TransactionResponse struct{
	TransactionID	string `json:"transactionID"`
	TransactionTime string `json:"transactionTime"`
	OrderID 	  	string `json:"orderID"`
	PaymentResult 	interface{} `json:"payment"`
	TotalPrice  	string `json:"totalPrice"`
	Status	 		string	`json:"status"`
	Order			PremiumOrders
}