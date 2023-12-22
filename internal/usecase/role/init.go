package role

import "lunar-commerce-fiber/internal/model"

type RoleUseCase struct {
	roleRepo model.RoleRepositoryInterface
}

func NewRoleUseCase(roleRepo model.RoleRepositoryInterface) model.RoleUseCaseInterface {
	return &RoleUseCase{
		roleRepo: roleRepo,
	}
}