package user

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/model"
)

type UserRepository struct {
	db *config.Database
}

func NewUserRepository(db *config.Database) model.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}