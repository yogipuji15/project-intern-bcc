package entity

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Role string `gorm:"type:varchar(50);unique" json:"role"`
}