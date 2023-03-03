package repository

import (
	"gorm.io/gorm"
	"project-intern-bcc/src/lib/storage"
)

type Repository struct {
	User      UserRepository
	Category  CategoryRepository
	Company	  CompanyRepository
	Order	  OrderRepository
	Payment   PaymentRepository
	Proposal  ProposalRepository
	Review    ReviewRepository
	Role      RoleRepository
	Schedule  ScheduleRepository
	Speaker   SpeakerRepository
	PremiumOrder PremiumOrderRepository
	CompanyCategory CompanyCategoryRepository
}

func Init(db *gorm.DB,storage storage.StorageInterface) *Repository {
	return &Repository{
		User:      NewUserRepository(db),
		Category: NewCategoryRepository(db),
		Company: NewCompanyRepository(db),
		Order: NewOrderRepository(db),
		Payment: NewPaymentRepository(db),
		Proposal: NewProposalRepository(db,storage),
		Review: NewReviewRepository(db),
		Role: NewRoleRepository(db),
		Schedule: NewScheduleRepository(db),
		Speaker: NewSpeakerRepository(db),
		PremiumOrder: NewPremiumOrderRepository(db),
		CompanyCategory: NewCompanyCategoryRepository(db),
	}
}