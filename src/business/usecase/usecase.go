package usecase

import (
	"project-intern-bcc/src/business/repository"
	"project-intern-bcc/src/lib/auth"
	"project-intern-bcc/src/lib/storage"
)

type Usecase struct {
	User UserUsecase
	Category CategoryUsecase
	Company CompanyUsecase
	Order OrderUsecase
	Payment PaymentUsecase
	Proposal ProposalUsecase
	Review ReviewUsecase
	Role RoleUsecase
	Schedule ScheduleUsecase
	Speaker SpeakerUsecase
	CompanyCategory CompanyCategoryUsecase
}

func Init(storage storage.StorageInterface,auth auth.AuthInterface,repo *repository.Repository) *Usecase {
	return &Usecase{
		User: NewUserUsecase(repo.User, auth, storage),
		Category :NewCategoryUsecase(repo.Category),
		Company : NewCompanyUsecase(repo.Company),
		Order : NewOrderUsecase(repo.Order),
		Payment : NewPaymentUsecase(repo.Payment),
		Proposal : NewProposalUsecase(repo.Proposal),
		Review : NewReviewUsecase(repo.Review),
		Role : NewRoleUsecase(repo.Role),
		Schedule : NewScheduleUsecase(repo.Schedule),
		Speaker : NewSpeakerUsecase(repo.Speaker),
		CompanyCategory: NewCompanyCategoryUsecase(repo.CompanyCategory),
	}
}