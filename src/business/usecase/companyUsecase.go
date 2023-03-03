package usecase

import (
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type CompanyUsecase interface {
	FindAll(category string, keyword string, pagination entity.Pagination)(interface{},int,error)
}

type companyUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyUsecase(r repository.CompanyRepository) CompanyUsecase {
	return &companyUsecase{
		companyRepository: r,
	}
}

func (h *companyUsecase) FindAll(category string, keyword string, pagination entity.Pagination)(interface{},int,error){
	companies,pg,err:=h.companyRepository.FindAll(category,keyword, pagination)
	if err!=nil{
		return "Failed to Querying sponsor companies Data",http.StatusNotFound,err
	}

	result:=entity.CompaniesResponse{
		Companies: companies,
		Pagination: *pg,
	}

	return result,http.StatusOK,nil
}