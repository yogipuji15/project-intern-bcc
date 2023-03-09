package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type CompanyUsecase interface {
	FindAll(filter entity.FilterParam, pagination entity.Pagination)(interface{},int,error)
	GetById(companyId string)(entity.Companies,int,error)
}

type companyUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyUsecase(r repository.CompanyRepository) CompanyUsecase {
	return &companyUsecase{
		companyRepository: r,
	}
}

func (h *companyUsecase) FindAll(filter entity.FilterParam, pagination entity.Pagination)(interface{},int,error){
	companies,pg,err:=h.companyRepository.FindAll(filter, pagination)
	if err!=nil{
		return "Failed to Querying sponsor companies Data",http.StatusNotFound,err
	}

	result:=entity.CompaniesResponse{
		Companies: companies,
		Pagination: *pg,
	}

	return result,http.StatusOK,nil
}

func (h *companyUsecase) GetById(companyId string)(entity.Companies,int,error){
	company,err:=h.companyRepository.GetById(companyId)
	if err!=nil{
		return company,http.StatusNotFound,err
	}

	return company,http.StatusOK,err
}