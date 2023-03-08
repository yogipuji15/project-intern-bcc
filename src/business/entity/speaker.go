package entity

import "gorm.io/gorm"

type Speakers struct {
	gorm.Model
	Name 			string    `gorm:"type:varchar(255)" json:"title"`
	Description 	string 	  `gorm:"type:longtext" json:"description"`
	Price 			int       `gorm:"type:int" json:"price"`
	Rating 			float32   `gorm:"type:float" json:"rating"`
	TotalReviews 	int       `gorm:"type:int" json:"totalReviews"`
	Photo 			string 	  `gorm:"type:varchar(255)" json:"photo"`
	Location 		string 	  `gorm:"type:varchar(255)" json:"location"`
	Email			string    `gorm:"type:varchar(50);unique" json:"email"`
	Portfolio		string 	  `gorm:"type:varchar(255)" json:"portfolio"`
	CategoryID		uint
	Category 		Categories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SpeakersResponse struct{
	Speakers []Speakers
	Pagination Pagination
}