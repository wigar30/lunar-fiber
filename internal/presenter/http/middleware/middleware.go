package middleware

import (
	"lunar-commerce-fiber/internal/model"
)

type Middleware struct {
	AuthMiddleware *AuthMiddleware
	RbacMiddleware *RbacMiddleware
}

func NewMiddleware(userRepo model.UserRepositoryInterface, roleRepo model.RoleRepositoryInterface, config *model.EnvConfigs) *Middleware {
	return &Middleware{
		AuthMiddleware: NewAuthMiddleware(userRepo, config),
		RbacMiddleware: NewRbacMiddleware(roleRepo, config),
	}
}