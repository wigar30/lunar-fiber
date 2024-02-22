package tenant

import (
	"errors"
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (tr *TenantRepository) GetAllByAuth(userId int64, p utils.Pagination) (*utils.Pagination, error) {
	var tenants []*entity.Tenant

	err := tr.db.
		Joins("LevelTenant").
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

func (tr *TenantRepository) GetByID(ID string) (*entity.Tenant, error) {
	var tenant *entity.Tenant

	err := tr.db.
		Joins("LevelTenant").
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
