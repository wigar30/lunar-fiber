package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type JwtClaims struct {
	ID   string `json:"id"`
	Role int    `json:"role"`
	jwt.StandardClaims
}

type AuthUseCaseInterface interface {
	Login(c *fiber.Ctx, req *LoginRequest) (*LoginResponse, error)
}
