package entity

import "gorm.io/gorm"

type CompanyCategories struct {
	gorm.Model
	SponsorCategory string `gorm:"type:varchar(50);unique" json:"sponsorCategory"`
}