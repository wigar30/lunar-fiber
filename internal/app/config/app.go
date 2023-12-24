package config

import (
	"fmt"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller"
	"lunar-commerce-fiber/internal/presenter/http/middleware"
	"lunar-commerce-fiber/internal/presenter/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

type HTTPServiceInterface interface {
	ListenApp() error
}

type HTTPService struct {
	app    *fiber.App
	ctrl   *controller.Controller
	config *model.EnvConfigs
	log    *logrus.Logger
	mdlwr  *middleware.Middleware
}

func NewListenApp(app *fiber.App, ctrl *controller.Controller, config *model.EnvConfigs, log *logrus.Logger, mdlwr *middleware.Middleware) HTTPServiceInterface {
	return &HTTPService{
		app:    app,
		ctrl:   ctrl,
		config: config,
		log:    log,
		mdlwr:  mdlwr,
	}
}

func (f *HTTPService) ListenApp() error {
	f.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.Route(f.app, f.ctrl, f.mdlwr)

	f.log.Info("Application is running...")
	port := f.config.AppPort
	return f.app.Listen(fmt.Sprintf(":%s", port))
}
