package router

import (
	"lunar-commerce-fiber/internal/presenter/http/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(f *fiber.App, ctrl *controller.Controller) {
	api := f.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/ping", ctrl.Ping)

	auth := v1.Group("auth")

	auth.Post("/login", ctrl.Auth.Login)

	role := v1.Group("role")

	role.Get("/", ctrl.Role.GetAll)
	role.Get("/:id", ctrl.Role.GetByID)
}