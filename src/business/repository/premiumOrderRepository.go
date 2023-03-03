package repository

import "gorm.io/gorm"

type PremiumOrderRepository interface {
}

type premiumOrderRepository struct {
	db *gorm.DB
}

func NewPremiumOrderRepository(db *gorm.DB) PremiumOrderRepository {
	return &premiumOrderRepository{db:db}
}