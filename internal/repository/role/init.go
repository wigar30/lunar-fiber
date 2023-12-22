package role

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/app/config"
)

type RoleRepository struct {
	db *config.Database
}

func NewRoleRepository(db *config.Database) model.RoleRepositoryInterface {
	return &RoleRepository{
		db: db,
	}
}