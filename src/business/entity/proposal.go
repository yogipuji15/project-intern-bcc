package entity

import "gorm.io/gorm"

type Proposals struct {
	gorm.Model
	Title string `gorm:"type:varchar(255)" json:"title"`
	Proposal string `gorm:"type:varchar(255)" json:"proposal"`
	Status bool `gorm:"type:bool;default:null" json:"status"`
	Description string `gorm:"type:longtext" json:"title"`
	UserID uint
	User Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CompanyID uint
	Company Companies `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}