package entity

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	CompanyName string `gorm:"type:varchar(255);unique" json:"company"`
	Email       string `gorm:"type:varchar(50);unique" json:"email"`
	Photo 		string `gorm:"type:varchar(255)" json:"photo"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	CategoryID uint
	Category CompanyCategories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CompaniesResponse struct{
	Companies []Companies
	Pagination Pagination
}