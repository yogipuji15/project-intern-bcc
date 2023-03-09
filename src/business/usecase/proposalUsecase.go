package usecase

import (
	"errors"
	"mime/multipart"
	"net/http"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/business/repository"
)

type ProposalUsecase interface {
	Create(user entity.UserResponse,proposalInput entity.InputProposal, file *multipart.FileHeader)(interface{},int,error)
}

type proposalUsecase struct {
	proposalRepository repository.ProposalRepository
}

func NewProposalUsecase(r repository.ProposalRepository) ProposalUsecase {
	return &proposalUsecase{
		proposalRepository: r,
	}
}

func (h *proposalUsecase) Create(user entity.UserResponse,proposalInput entity.InputProposal, file *multipart.FileHeader)(interface{},int,error){
	totalProposal,err:=h.proposalRepository.GetTotalProposalByUserId(user.ID)
	if err!=nil{
		return "Failed to querying proposal's data",http.StatusNotFound,err
	}
	
	if totalProposal>=3 && user.Role=="free-user"{
		return "Failed to apply proposal",http.StatusUnauthorized,errors.New("You have applied 3 proposals, upgrade your account to premium account")
	}

	_,err= h.proposalRepository.GetByUserAndCompanyId(user.ID,proposalInput.CompanyID)
	if err==nil{
		return "Failed to apply proposal",http.StatusBadRequest,errors.New("You have applied proposal for this company")
	}

	proposal := entity.Proposals{
		Name 		:proposalInput.Name,
		Status 		:"WAITING TO APPROVE",
		Message		:proposalInput.Message,
		Email 		:proposalInput.Email,
		Phone 		:proposalInput.Phone,
		UserID 		:user.ID,
		CompanyID 	:uint(proposalInput.CompanyID),
	}

	proposalResponse,err:=h.proposalRepository.Create(proposal,file)
	if err!=nil{
		return "Failed to apply proposal",http.StatusInternalServerError,err
	}

	return proposalResponse,http.StatusOK,nil
}

