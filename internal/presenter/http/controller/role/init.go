package role

import (
	"lunar-commerce-fiber/internal/model"
)

type RoleController struct {
	roleUseCase model.RoleUseCaseInterface
}

func NewRoleController(roleUseCase model.RoleUseCaseInterface) *RoleController {
	return &RoleController{
		roleUseCase: roleUseCase,
	}
}