package tenant

import "lunar-commerce-fiber/internal/model"

type TenantUseCase struct {
	tenantRepo     model.TenantRepositoryInterface
	membershipRepo model.MembershipRepositoryInterface
}

func NewTenantUseCase(tenantRepo model.TenantRepositoryInterface, membershipRepo model.MembershipRepositoryInterface) model.TenantUseCaseInterface {
	return &TenantUseCase{
		tenantRepo:     tenantRepo,
		membershipRepo: membershipRepo,
	}
}
