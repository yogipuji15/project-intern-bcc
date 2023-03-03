package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindID(r string)(uint,error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db:db}
}

func (h *roleRepository) FindID(r string)(uint,error){
	var role entity.Roles
	err:=h.db.First(&role,"role=?",r).Error
	return role.ID , err
}