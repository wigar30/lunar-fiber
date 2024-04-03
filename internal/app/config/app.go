package config

import (
	"fmt"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller"
	"lunar-commerce-fiber/internal/presenter/http/middleware"
	"lunar-commerce-fiber/internal/presenter/http/router"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HTTPServiceInterface interface {
	ListenApp() error
}

type HTTPService struct {
	app    *fiber.App
	ctrl   *controller.Controller
	config *model.EnvConfigs
	log    *model.Logger
	mdlwr  *middleware.Middleware
}

func NewListenApp(app *fiber.App, ctrl *controller.Controller, config *model.EnvConfigs, log *model.Logger, mdlwr *middleware.Middleware) HTTPServiceInterface {
	return &HTTPService{
		app:    app,
		ctrl:   ctrl,
		config: config,
		log:    log,
		mdlwr:  mdlwr,
	}
}

func (f *HTTPService) ListenApp() error {
	if f.config.AppEnv == "production" {
		f.app.Use(
			limiter.New(limiter.Config{
				Max: 100,
				LimitReached: func(c *fiber.Ctx) error {
					return model.OnError(c, &model.ErrorResponse{
						Code:    fiber.StatusTooManyRequests,
						Message: "Too Many Requests",
					})
				},
			}),
		)
	}
	f.app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
		}),
		logger.New(),
		// cache.New(),
		helmet.New(),
		recover.New(),
	)

	router.Route(f.app, f.ctrl, f.mdlwr)

	f.app.Static("/storages", "./storages")

	f.app.Use(func(c *fiber.Ctx) error {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "Nothing found",
		})
	})

	f.log.Info("Application is running...")
	port := f.config.AppPort
	return f.app.Listen(fmt.Sprintf(":%s", port))
}
