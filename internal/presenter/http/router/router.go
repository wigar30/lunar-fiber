package router

import (
	"lunar-commerce-fiber/internal/presenter/http/controller"
	"lunar-commerce-fiber/internal/presenter/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Route(f *fiber.App, ctrl *controller.Controller, m *middleware.Middleware) {
	api := f.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/ping", ctrl.Ping)

	auth := v1.Group("auth")
	auth.Post("/login", ctrl.Auth.Login)

	user := v1.Group("user", m.AuthMiddleware.ValidateToken())
	user.Get("/profile", ctrl.User.GetProfile)

	role := v1.Group("role", cache.New(), m.AuthMiddleware.ValidateToken())
	role.Get("/", m.RbacMiddleware.ValidateRoleUser([]string{"All"}), ctrl.Role.GetAll)
	role.Get("/:id", m.RbacMiddleware.ValidateRoleUser([]string{"All"}), ctrl.Role.GetByID)

	tenant := v1.Group("tenant", m.AuthMiddleware.ValidateToken())
	tenant.Post("/", ctrl.Tenant.CreateTenant)
	tenant.Get("/auth", ctrl.Tenant.GetAllByAuth)
	tenant.Get("/:id", ctrl.Tenant.GetByID)
	tenant.Get("/:id/product", ctrl.Product.GetAllByTenantID)
	tenant.Get("/:id/product/:productId", ctrl.Product.GetByID)

	product := v1.Group("product", m.AuthMiddleware.ValidateToken())
	product.Post("/", ctrl.Product.CreateProduct)

}
