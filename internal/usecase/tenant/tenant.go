package tenant

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"
)

func (tu *TenantUseCase) GetAllByAuth(userId int64, p model.PaginationRequest) (*utils.Pagination, error) {
	tenants, err := tu.tenantRepo.GetAllByAuth(userId, utils.Pagination{
		Page: p.Page,
		Limit: p.Limit,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return tenants, nil
}

func (tu *TenantUseCase) GetByID(userId string, ID string) (*model.TenantResponse, error) {
	tenant, err := tu.tenantRepo.GetByID(userId, ID)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code:    errC.Code,
			Message: errC.Error(),
		}
	}

	return &model.TenantResponse{
		ID: tenant.ID,
		Name: tenant.Name,
		TotalProduct: int64(tenant.TotalProduct),
		LevelID: tenant.LevelID,
		Level: &tenant.LevelTenant,
		Memberships: tenant.Memberships,
		SummaryStat: tenant.SummaryStat,
	}, nil
}