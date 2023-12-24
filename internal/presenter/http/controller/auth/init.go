package auth

import "lunar-commerce-fiber/internal/model"

type AuthController struct {
	authUseCase model.AuthUseCaseInterface
}

func NewAuthController(authUseCase model.AuthUseCaseInterface) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}