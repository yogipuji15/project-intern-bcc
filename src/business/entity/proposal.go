package entity

import "gorm.io/gorm"

type Proposals struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)" json:"name"`
	Proposal string `gorm:"type:varchar(255)" json:"proposal"`
	Status string `gorm:"type:varchar(255)" json:"status"`
	Message string `gorm:"type:longtext" json:"message"`
	Email string `gorm:"type:varchar(50)" json:"email"`
	Phone string `gorm:"type:varchar(50)" json:"phone"`
	UserID uint `json:"userId"`
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyID uint `json:"companyId"`
	Company Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type InputProposal struct{
	Name 		string `form:"name" binding:"required"`
	Email 		string `form:"email" binding:"required"`
	Phone 		string `form:"phone" binding:"required"`
	Message		string `form:"message" binding:"required"`
	CompanyID 	int `form:"companyId" binding:"required,number"`
}