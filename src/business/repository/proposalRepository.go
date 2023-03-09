package repository

import (
	"mime/multipart"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/lib/storage"

	"gorm.io/gorm"
)

type ProposalRepository interface {
	Create(proposal entity.Proposals,file *multipart.FileHeader)(entity.Proposals,error)
	GetTotalProposalByUserId(userId uint)(int64,error)
	GetByUserAndCompanyId(userId uint,companyId int)(entity.Proposals,error)
}

type proposalRepository struct {
	db *gorm.DB
	storage storage.StorageInterface
}

func NewProposalRepository(db *gorm.DB,storage storage.StorageInterface) ProposalRepository {
	return &proposalRepository{
		db:db,
		storage: storage,
	}
}

func (h *proposalRepository) Create(proposal entity.Proposals,file *multipart.FileHeader)(entity.Proposals,error){
	proposalLink, err :=h.storage.UploadFile(file)
	if err != nil{
		return proposal,err
	}
	
	proposal.Proposal=proposalLink

	err=h.db.Create(&proposal).Error

	return proposal,err
}

func (h *proposalRepository) GetTotalProposalByUserId(userId uint)(int64,error){
	var proposals []entity.Proposals
	var total int64

	err := h.db.Where("user_id = ?", userId).Find(&proposals).Error
	if err!=nil{
		return total,err
	}
	
	err = h.db.Model(&proposals).Where("user_id = ?", userId).Count(&total).Error
	
	return total,err
}

func (h *proposalRepository) GetByUserAndCompanyId(userId uint,companyId int)(entity.Proposals,error){
	var proposal entity.Proposals
	err:=h.db.Where("user_id = ? AND company_id = ?", userId,companyId).First(&proposal).Error
	return proposal,err
}