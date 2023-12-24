package auth

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/model"
)

type AuthUseCase struct {
	userRepo model.UserRepositoryInterface
	config   *config.EnvConfigs
}

func NewAuthUseCase(userRepo model.UserRepositoryInterface, config *config.EnvConfigs) model.AuthUseCaseInterface {
	return &AuthUseCase{
		userRepo: userRepo,
		config:   config,
	}
}
