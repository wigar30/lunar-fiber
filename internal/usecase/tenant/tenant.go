package tenant

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (tu *TenantUseCase) GetAllByAuth(userId int64, p model.PaginationRequest) (*utils.Pagination, error) {
	tenants, err := tu.tenantRepo.GetAllByAuth(userId, utils.Pagination{
		Page:  p.Page,
		Limit: p.Limit,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code:    errC.Code,
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
		ID:           tenant.ID,
		Name:         tenant.Name,
		TotalProduct: int64(tenant.TotalProduct),
		LevelID:      tenant.LevelID,
		Level:        &tenant.LevelTenant,
		Memberships:  tenant.Memberships,
		SummaryStat:  tenant.SummaryStat,
	}, nil
}

func (tu *TenantUseCase) CreateTenant(userId int64, tenant model.CreateTenant) (string, error) {
	tenantCount, err := tu.tenantRepo.GetCountAuthTenant(userId)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return "", &model.ErrorResponse{
			Code:    errC.Code,
			Message: errC.Error(),
		}
	}

	if tenantCount == 3 {
		return "", &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "You have reached maximum of tenant",
		}
	}

	tx := tu.tenantRepo.BeginTransaction()
	ID, err := tu.tenantRepo.CreateTenant(tx, model.CreateTenant{
		Name: tenant.Name,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		tu.tenantRepo.RollbackTransaction(tx)
		return "", &model.ErrorResponse{
			Code:    errC.Code,
			Message: errC.Error(),
		}
	}

	_, err = tu.membershipRepo.CreateMembership(tx, model.CreateMembership{
		UserID: strconv.FormatInt(userId, 10),
		RoleID: "3",
		TenantID: ID,
		StatusID: "1",
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		tu.tenantRepo.RollbackTransaction(tx)
		return "", &model.ErrorResponse{
			Code:    errC.Code,
			Message: errC.Error(),
		}
	}

	return ID, tu.tenantRepo.CommitTransaction(tx)
}
