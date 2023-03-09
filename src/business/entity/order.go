package entity

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderCode string `gorm:"type:varchar(255);unique" json:"orderCode"`
	EventName string `gorm:"type:varchar(255)" json:"eventName"`
	Status string `gorm:"type:varchar(255)" json:"status"`
	BookTimeStart time.Time `json:"bookTimeStart"`
	BookTimeEnd time.Time `json:"bookTimeEnd"`
	Description string `gorm:"longtext" json:"description"`
	Duration int `json:"duration"`
	TotalPrice int `gorm:"type:int" json:"totalPrice"`
	Rundown string `gorm:"type:varchar(255)" json:"rundown"`
	Script string `gorm:"type:varchar(255)" json:"script"`
	UserID uint
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpeakerID uint
	Speaker Speakers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentID uint
	Payment Payments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderInput struct{
	EventName string `form:"eventName" binding:"required"`
	Description string `form:"description" binding:"required"`
	PaymentType string `form:"paymentType" binding:"required"`
	Duration int `form:"duration" binding:"required,number"`
	SpeakerID int `form:"speakerId" binding:"required,number"`
	BookTime string `form:"bookTime" binding:"required"`
}

type MidtransTransactionResponse struct{
	TransactionID	string `json:"transactionID"`
	TransactionTime string `json:"transactionTime"`
	OrderID 	  	string `json:"orderID"`
	PaymentResult 	interface{} `json:"payment"`
	TotalPrice  	string `json:"totalPrice"`
	Status	 		string	`json:"status"`
	Order			OrderResponse
}

type OrderResponse struct{
	OrderCode 		string `json:"orderCode"`
	EventName 		string `json:"eventName"`
	Status 			string `json:"status"`
	BookTimeStart 	time.Time `json:"bookTimeStart"`
	BookTimeEnd 	time.Time `json:"bookTimeEnd"`
	Description 	string `json:"description"`
	Duration 		int `json:"duration"`
	TotalPrice 		int `json:"totalPrice"`
	PaymentType		string `json:"paymentType"`
	Speaker 		Speakers
}

type CheckTransaction struct{
	TransactionStatus string `json:"transaction_status"`
	StatusCode		  string `json:"status_code"`
	SignatureKey	  string `json:"signature_key"`
	PaymentType		  string `json:"payment_type"`
	OrderID			  string `json:"order_id"`
	GrossAmount		  string `json:"gross_amount"`
}

type OrderHistoryResponse struct{
	Order 	   []Orders
	Pagination Pagination
}