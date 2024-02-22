package tenant

import "lunar-commerce-fiber/internal/model"

type TenantController struct {
	tenantUseCase model.TenantUseCaseInterface
}

func NewTenantController(tenantUseCase model.TenantUseCaseInterface) *TenantController {
	return &TenantController{
		tenantUseCase: tenantUseCase,
	}
}
