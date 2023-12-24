package user

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type UserRepository struct {
	db *driver.Database
}

func NewUserRepository(db *driver.Database) model.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}