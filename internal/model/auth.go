package model

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthUseCaseInterface interface {
	Login(c *fiber.Ctx, req *LoginRequest) (*LoginResponse, error)
}
