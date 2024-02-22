package controller

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller/auth"
	"lunar-commerce-fiber/internal/presenter/http/controller/role"
	"lunar-commerce-fiber/internal/presenter/http/controller/tenant"
	"lunar-commerce-fiber/internal/presenter/http/controller/user"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Role   *role.RoleController
	Auth   *auth.AuthController
	User   *user.UserController
	Tenant *tenant.TenantController
}

func NewController(
	RoleController *role.RoleController,
	AuthController *auth.AuthController,
	UserController *user.UserController,
	TenantController *tenant.TenantController,
) *Controller {
	return &Controller{
		Role:   RoleController,
		Auth:   AuthController,
		User:   UserController,
		Tenant: TenantController,
	}
}

func (ctrl *Controller) Ping(c *fiber.Ctx) error {

	return model.OnSuccess(c, &model.PingResponse{
		Ping: "pong fiber",
	})
}
