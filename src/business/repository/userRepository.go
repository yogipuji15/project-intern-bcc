package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Test()(entity.Users,error)
	Create(user entity.Users)(entity.Users,error)
	FindByEmailUsername(emailUsername string)(entity.Users,error)
	Update(user entity.Users)(error)
	Delete(user entity.Users)(error)
	FindUserByToken(token string)(entity.Users,error)
	FindByEmail(email string)(entity.Users,error)
	FindById(id any)(entity.Users,error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (h *userRepository) Test()(entity.Users,error){
	var user entity.Users
	err:= h.db.Find(&user).Error
	return user,err
}

func (h *userRepository) Create(user entity.Users)(entity.Users,error){
	err:=h.db.Create(&user).Error
	return user,err
}

func (h *userRepository) Update(user entity.Users)(error){
	err:=h.db.Save(&user).Error
	return err
}

func (h *userRepository) FindByEmailUsername(emailUsername string)(entity.Users,error){
	var user entity.Users
	err:=h.db.Where("email = ?", emailUsername).Or("username = ?", emailUsername).First(&user).Error
	return user,err
}

func (h *userRepository) FindByEmail(email string)(entity.Users,error){
	var user entity.Users
	err:=h.db.Where("email = ?", email).First(&user).Error
	return user,err
}

func (h *userRepository) Delete(user entity.Users)(error){
	err:=h.db.Delete(&user).Error
	return err
}

func (h *userRepository) FindUserByToken(token string)(entity.Users,error){
	var user entity.Users
	err:=h.db.First(&user, "verification_code=?",token).Error
	return user,err
}

func (h *userRepository) FindById(id any)(entity.Users,error){
	var user entity.Users
	err:=h.db.First(&user, "id=?",id).Error
	return user,err
}