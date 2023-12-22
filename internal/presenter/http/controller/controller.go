package controller

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Ping(c *fiber.Ctx) error {

	return model.OnSuccess(c, &model.PingResponse{
		Ping: "pong",
	})
}