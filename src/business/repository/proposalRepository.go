package repository

import (
	"gorm.io/gorm"
	"project-intern-bcc/src/lib/storage"
)

type ProposalRepository interface {
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