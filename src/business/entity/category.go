package entity

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Category string `gorm:"type:varchar(50);unique" json:"category"`
	Photo	 string `gorm:"type:varchar(100)" json:"photo"`
}
