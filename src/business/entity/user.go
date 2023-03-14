package entity

import "time"

type Users struct {
	ID               uint      `gorm:"type:bigint(20);primaryKey" json:"id"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	Email            string    `gorm:"type:varchar(50);unique" json:"email"`
	Username         string    `gorm:"type:varchar(255);unique" json:"username"`
	Fullname		 string	   `gorm:"type:varchar(255)" json:"fullname"`
	Address			 string	   `gorm:"type:varchar(255)" json:"address"`
	Phone			 string	   `gorm:"type:varchar(255)" json:"phone"`
	Password         string    `gorm:"type:varchar(255)" json:"password"`
	PremiumDue 		 time.Time `gorm:"default:null" json:"premiumDue"`
	RoleID			 uint	   `json:"roleId"`
	Role 			 Roles	   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VerificationCode string    `gorm:"type:varchar(255)" json:"verificationCode"`
	IsActive       	 bool      `gorm:"not null" json:"isActive"`
}

type UserResponse struct {
	ID		 uint `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Phone	 string	`json:"phone"`
	Role	 string `json:"role"`
	PremiumDue time.Time `json:"premiumDue"`
}

type UserSignup struct {
	Email    		 string `json:"email" binding:"required,email"`
	Username 		 string `json:"username" binding:"required"`
	Fullname		 string	`json:"fullname" binding:"required"`
	Address			 string	`json:"address"  binding:"required"`
	Phone			 string	`json:"phone"  binding:"required"`
	Password 		 string `json:"password" binding:"required"`
	ConfirmPass 	 string `json:"confirmPassword" binding:"required"`
}

type UserLogin struct {
	EmailUsername    string `json:"email_username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
