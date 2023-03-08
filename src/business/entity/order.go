package entity

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderCode string `gorm:"type:varchar(255);unique" json:"orderCode"`
	EventName string `gorm:"type:varchar(255)" json:"eventName"`
	Status bool `gorm:"type:bool;default:null" json:"status"`
	BookTimeStart time.Time `gorm:"type:date" json:"bookTimeStart"`
	BookTimeEnd time.Time `gorm:"type:date" json:"bookTimeEnd"`
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
	EventName string `form:"eventName" `
	Description string `form:"description" `
	PaymentType string `form:"paymentType" `
	Duration int `form:"duration" `
	SpeakerID int `form:"speakerId" `
	BookTime string `form:"bookTime" `
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
	Status 			bool `json:"status"`
	BookTimeStart 	time.Time `json:"bookTimeStart"`
	BookTimeEnd 	time.Time `json:"bookTimeEnd"`
	Description 	string `json:"description"`
	Duration 		int `json:"duration"`
	TotalPrice 		int `json:"totalPrice"`
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