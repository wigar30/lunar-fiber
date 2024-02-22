package tenant

import "lunar-commerce-fiber/internal/model"

type TenantUseCase struct {
	tenantRepo model.TenantRepositoryInterface
}

func NewTenantUseCase(tenantRepo model.TenantRepositoryInterface) model.TenantUseCaseInterface {
	return &TenantUseCase{
		tenantRepo: tenantRepo,
	}
}