package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	Review 		string `gorm:"type:longtext" json:"review"`
	Star 		int `gorm:"type:int" json:"star"`
	Date 		datatypes.Date `gorm:"type:date" json:"date"`
	UserID 		uint
	User Users 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpeakerID 	uint
	Speaker 	Speakers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ReviewsResponse struct{
	Reviews []ReviewResponse
	Pagination Pagination
}

type ReviewResponse struct{
	Review		string ` json:"review"`
	Star   		int ` json:"star"`
	Date 		datatypes.Date ` json:"date"`
	Username	string ` json:"username"`
}

type PostReview struct{
	Review 		string `json:"review" binding:"required"`
	Star 		int `json:"star" binding:"required"`
	SpeakerId   uint `json:"speakerId" binding:"required"`
}