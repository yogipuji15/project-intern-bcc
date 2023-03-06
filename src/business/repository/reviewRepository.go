package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetBySpeakerId(id string, pagination entity.Pagination)([]entity.Reviews,*entity.Pagination,error)
	Create(review entity.Reviews)(entity.Reviews,error)
	GetAll(id uint)([]entity.Reviews,error)
	GetAllBySpeakerAndUserId(userId uint, speakerId uint)(entity.Reviews,error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db:db}
}

func (h *reviewRepository) GetBySpeakerId(id string, pagination entity.Pagination)([]entity.Reviews,*entity.Pagination,error){
	pg:= entity.FormatPaginationParam(pagination)

	var reviews []entity.Reviews
	err:= h.db.Preload("User").Preload("Speaker").Where("speaker_id = ?", id).Offset(int(pg.Offset)).Limit(int(pg.Limit)).Find(&reviews).Error
	if err!=nil{
		return nil,nil,err
	}

	err = h.db.Order("rating desc").Model(&reviews).Preload("User").Preload("Speaker").Where("speaker_id = ?", id).Count(&pg.TotalElement).Error
	if err!=nil{
		return nil,nil,err
	}

	pg.ProcessPagination(int64(len(reviews)))

	return reviews,&pg,err
}

func (h *reviewRepository) GetAll(id uint)([]entity.Reviews,error){
	var reviews []entity.Reviews
	err:= h.db.Where("speaker_id = ?", id).Find(&reviews).Error
	if err!=nil{
		return nil,err
	}
	return reviews,nil
}

func (h *reviewRepository) Create(review entity.Reviews)(entity.Reviews,error){
	err:=h.db.Create(&review).Error
	return review,err
}

func (h *reviewRepository) GetAllBySpeakerAndUserId(userId uint, speakerId uint)(entity.Reviews,error){
	var reviews entity.Reviews
	err:= h.db.Where("speaker_id = ? AND user_id = ?", speakerId,userId).First(&reviews).Error
	if err!=nil{
		return reviews,err
	}
	return reviews,nil
}