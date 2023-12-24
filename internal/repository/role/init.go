package role

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type RoleRepository struct {
	db *driver.Database
}

func NewRoleRepository(db *driver.Database) model.RoleRepositoryInterface {
	return &RoleRepository{
		db: db,
	}
}