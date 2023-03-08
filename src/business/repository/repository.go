package repository

import (
	"project-intern-bcc/src/lib/midtrans"
	"project-intern-bcc/src/lib/storage"

	"gorm.io/gorm"
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

func Init(db *gorm.DB,storage storage.StorageInterface, midtrans midtrans.MidtransInterface) *Repository {
	return &Repository{
		User:      NewUserRepository(db),
		Category: NewCategoryRepository(db),
		Company: NewCompanyRepository(db),
		Order: NewOrderRepository(db,storage,midtrans),
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