package usecase

import "project-intern-bcc/src/business/repository"

type ProposalUsecase interface {
}

type proposalUsecase struct {
	proposalRepository repository.ProposalRepository
}

func NewProposalUsecase(r repository.ProposalRepository) ProposalUsecase {
	return &proposalUsecase{
		proposalRepository: r,
	}
}