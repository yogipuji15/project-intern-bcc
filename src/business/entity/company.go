package entity

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	CompanyName 		string `gorm:"type:varchar(255);unique" json:"company"`
	Email       		string `gorm:"type:varchar(50);unique" json:"email"`
	Photo 				string `gorm:"type:varchar(255)" json:"photo"`
	Bio					string `gorm:"longtext" json:"bio"`
	Description			string `gorm:"longtext" json:"description"`
	Location 	  		string `gorm:"type:varchar(255)" json:"location"`
	ProposalsAccepted	int `gorm:"type:int" json:"proposalsAccepted"`
	Partner				string `gorm:"type:varchar(255)" json:"partner"`
	Website				string `gorm:"type:varchar(255)" json:"website"`
	Phone			 	string `gorm:"type:varchar(255)" json:"phone"`
	SponsorAssistance	string `gorm:"longtext" json:"sponsorAssistance"`
	CategoryID			uint `json:"categoryId"`
	Category 			Categories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CompaniesResponse struct{
	Companies []Companies
	Pagination Pagination
}