package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll()([]entity.Categories,error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db:db}
}

func (h *categoryRepository) FindAll()([]entity.Categories,error){
	var category []entity.Categories
	err:= h.db.Find(&category).Error
	return category,err
}