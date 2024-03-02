package tenant

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/go-playground/validator/v10"
)

type TenantController struct {
	tenantUseCase model.TenantUseCaseInterface

	validator *validator.Validate
}

func NewTenantController(tenantUseCase model.TenantUseCaseInterface, validator *validator.Validate) *TenantController {
	return &TenantController{
		tenantUseCase: tenantUseCase,
		validator:     validator,
	}
}
