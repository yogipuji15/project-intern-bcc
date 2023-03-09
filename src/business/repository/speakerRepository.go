package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type SpeakerRepository interface {
	FindAll(filter entity.FilterParam, pagination entity.Pagination)([]entity.Speakers,*entity.Pagination,error)
	GetById(id string) (entity.Speakers,error)
	UpdateRating(speaker entity.Speakers) (error)
}

type speakerRepository struct {
	db *gorm.DB
}

func NewSpeakerRepository(db *gorm.DB) SpeakerRepository {
	return &speakerRepository{db:db}
}

func (h *speakerRepository) FindAll(filter entity.FilterParam, pagination entity.Pagination)([]entity.Speakers,*entity.Pagination,error){
	pg:= entity.FormatPaginationParam(pagination)
	maxPrice:=filter.MaxPrice
	if maxPrice==0{
		maxPrice=999999999999999999
	}

	var speakers []entity.Speakers
	err:= h.db.Order("rating desc").Joins("Category", h.db.Where(&entity.Categories{Category: filter.Category})).Where("name LIKE ? AND location LIKE ? AND price BETWEEN ? AND ?", "%"+filter.Keyword+"%", "%"+filter.Location+"%",filter.MinPrice,maxPrice).Offset(int(pg.Offset)).Limit(int(pg.Limit)).Find(&speakers).Error
	if err!=nil{
		return nil,nil,err
	}

	err = h.db.Order("rating desc").Model(&speakers).Joins("Category", h.db.Where(&entity.Categories{Category: filter.Category})).Where("name LIKE ? AND location LIKE ? AND price BETWEEN ? AND ?", "%"+filter.Keyword+"%", "%"+filter.Location+"%",filter.MinPrice,maxPrice).Count(&pg.TotalElement).Error
	if err!=nil{
		return nil,nil,err
	}

	pg.ProcessPagination(int64(len(speakers)))

	return speakers,&pg,err
}

func (h *speakerRepository) GetById(id string) (entity.Speakers,error){
	var speaker entity.Speakers
	err:=h.db.Preload("Category").Where("id = ?", id).First(&speaker).Error
	
	return speaker,err
}

func (h *speakerRepository) UpdateRating(speaker entity.Speakers) (error){
	err:=h.db.Save(&speaker).Error
	return err
}