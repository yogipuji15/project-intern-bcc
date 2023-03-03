package repository

import (
	"project-intern-bcc/src/business/entity"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	FindAll(category string, keyword string, pagination entity.Pagination)([]entity.Companies,*entity.Pagination,error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db:db}
}

func (h *companyRepository) FindAll(category string, keyword string, pagination entity.Pagination)([]entity.Companies,*entity.Pagination,error){
	pg:= entity.FormatPaginationParam(pagination)
	
	

	var companies []entity.Companies
	err:= h.db.Where("company_name LIKE ?", "%"+keyword+"%").Offset(int(pg.Offset)).Limit(int(pg.Limit)).Find(&companies).Error
	if err!=nil{
		return nil,nil,err
	}

	err = h.db.Model(&companies).Where("company_name LIKE ?", "%"+keyword+"%").Count(&pg.TotalElement).Error
	if err!=nil{
		return nil,nil,err
	}
	pg.ProcessPagination(int64(len(companies)))

	return companies,&pg,err
}