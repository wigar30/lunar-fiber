package controller

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller/role"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Role *role.RoleController
}

func NewController(
	RoleController *role.RoleController,
) *Controller {
	return &Controller{
		Role: RoleController,
	}
}

func (ctrl *Controller) Ping(c *fiber.Ctx) error {

	return model.OnSuccess(c, &model.PingResponse{
		Ping: "pong",
	})
}
