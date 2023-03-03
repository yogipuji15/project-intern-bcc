package usecase

import "project-intern-bcc/src/business/repository"

type RoleUsecase interface {
}

type roleUsecase struct {
	roleRepository repository.RoleRepository
}

func NewRoleUsecase(r repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepository: r,
	}
}

func (h *roleUsecase) FindID(role string) (uint,error){
	return h.roleRepository.FindID(role)
}