package router

import (
	"lunar-commerce-fiber/internal/presenter/http/controller"
	"lunar-commerce-fiber/internal/presenter/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(f *fiber.App, ctrl *controller.Controller, m *middleware.Middleware) {
	api := f.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/ping", ctrl.Ping)

	auth := v1.Group("auth")
	auth.Post("/login", ctrl.Auth.Login)

	user := v1.Group("user", m.AuthMiddleware.ValidateToken())
	user.Get("/profile", ctrl.User.GetProfile)

	role := v1.Group("role", m.AuthMiddleware.ValidateToken())
	role.Get("/", ctrl.Role.GetAll)
	role.Get("/:id", ctrl.Role.GetByID)
}