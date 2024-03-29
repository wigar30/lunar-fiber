package tenant

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type TenantRepository struct {
	db *driver.Database
}

func NewTenantRepository(db *driver.Database) model.TenantRepositoryInterface {
	return &TenantRepository{
		db: db,
	}
}
