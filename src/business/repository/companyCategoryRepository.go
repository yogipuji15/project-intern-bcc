package repository

import (

	"gorm.io/gorm"
)

type CompanyCategoryRepository interface {
}

type companyCategoryRepository struct {
	db *gorm.DB
}

func NewCompanyCategoryRepository(db *gorm.DB) CompanyCategoryRepository {
	return &companyCategoryRepository{db:db}
}

