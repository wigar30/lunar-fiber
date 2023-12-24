package controller

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller/auth"
	"lunar-commerce-fiber/internal/presenter/http/controller/role"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Role *role.RoleController
	Auth *auth.AuthController
}

func NewController(
	RoleController *role.RoleController,
	AuthController *auth.AuthController,
) *Controller {
	return &Controller{
		Role: RoleController,
		Auth: AuthController,
	}
}

func (ctrl *Controller) Ping(c *fiber.Ctx) error {

	return model.OnSuccess(c, &model.PingResponse{
		Ping: "pong",
	})
}
