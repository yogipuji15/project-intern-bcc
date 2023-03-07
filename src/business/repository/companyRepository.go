package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	FindAll(filter entity.FilterParam, pagination entity.Pagination)([]entity.Companies,*entity.Pagination,error)
	GetById(id string) (entity.Companies,error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db:db}
}

func (h *companyRepository) FindAll(filter entity.FilterParam, pagination entity.Pagination)([]entity.Companies,*entity.Pagination,error){
	pg:= entity.FormatPaginationParam(pagination)
	
	

	var companies []entity.Companies
	err:= h.db.Joins("Category", h.db.Where(&entity.Categories{Category: filter.Category})).Where("company_name LIKE ? AND location LIKE ?", "%"+filter.Keyword+"%","%"+filter.Location+"%").Offset(int(pg.Offset)).Limit(int(pg.Limit)).Find(&companies).Error
	if err!=nil{
		return nil,nil,err
	}
	
	err = h.db.Joins("Category", h.db.Where(&entity.Categories{Category: filter.Category})).Model(&companies).Where("company_name LIKE ? AND location LIKE ?", "%"+filter.Keyword+"%","%"+filter.Location+"%").Count(&pg.TotalElement).Error
	if err!=nil{
		return nil,nil,err
	}
	pg.ProcessPagination(int64(len(companies)))

	return companies,&pg,err
}

func (h *companyRepository) GetById(id string) (entity.Companies,error){
	var company entity.Companies
	err:=h.db.Preload("Category").Where("id = ?", id).First(&company).Error
	
	return company,err
}