package tenant

import (
	"errors"
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (tr *TenantRepository) BeginTransaction() *gorm.DB {
	return tr.db.Begin()
}

func (tr *TenantRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (tr *TenantRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (tr *TenantRepository) GetAllByAuth(userId int64, p utils.Pagination) (*utils.Pagination, error) {
	var tenants []*entity.Tenant

	err := tr.db.
		Joins("LevelTenant").
		Joins("SummaryStat").
		Where("EXISTS(?)", tr.db.Table("memberships").Select("1").Where("tenants.id = memberships.tenantId AND userId = ?", userId)).
		Scopes(utils.Paginate(&p, tr.db.Where("EXISTS(?)", tr.db.Table("memberships").Select("1").Where("tenants.id = memberships.tenantId AND userId = ?", userId)), &tenants)).
		Find(&tenants).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	p.Items = tenants

	return &p, nil
}

func (tr *TenantRepository) GetCountAuthTenant(userId int64) (int64, error) {
	var tenantCount int64
	var tenant entity.Tenant
	err := tr.db.
		Model(&tenant).
		Where("EXISTS(?)", tr.db.Table("memberships").Select("1").Where("tenants.id = memberships.tenantId AND userId = ?", userId)).
		Count(&tenantCount).
		Error
	if err != nil {
		return 0, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return tenantCount, nil
}

func (tr *TenantRepository) GetByID(userId string, ID string) (*entity.Tenant, error) {
	var tenant *entity.Tenant
	var member *entity.Membership

	err := tr.db.Table("memberships").Select("id").Where(&entity.Membership{UserID: userId, TenantID: ID}).First(&member).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = tr.db.
		Joins("LevelTenant").
		Joins("SummaryStat").
		Preload("Memberships", func(db *gorm.DB) *gorm.DB {
			return db.
				Select("memberships.id, memberships.userId, memberships.roleId, memberships.tenantId").
				Joins("User", tr.db.Select("id", "name")).
				Joins("Role", tr.db.Select("id", "name"))
		}).
		First(&tenant, ID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return tenant, nil
}

func (tr *TenantRepository) CreateTenant(tx *gorm.DB, tenant model.CreateTenant) (string, error) {
	tenantResult := &entity.Tenant{
		Name: tenant.Name,
		LevelID: "1",
	}

	err := tx.Create(&tenantResult).Error
	if err != nil {
		return "", &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return tenantResult.ID, nil
}
