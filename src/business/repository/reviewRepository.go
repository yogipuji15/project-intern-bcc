package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetByUserId(id string, pagination entity.Pagination)([]entity.Reviews,*entity.Pagination,error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db:db}
}

func (h *reviewRepository) GetByUserId(id string, pagination entity.Pagination)([]entity.Reviews,*entity.Pagination,error){
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