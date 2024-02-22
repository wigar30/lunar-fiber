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