package config

import (
	"fmt"
	"lunar-commerce-fiber/internal/presenter/http/controller"
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
	config *EnvConfigs
	log    *logrus.Logger
}

func NewListenApp(app *fiber.App, ctrl *controller.Controller, config *EnvConfigs, log *logrus.Logger) HTTPServiceInterface {
	return &HTTPService{
		app:    app,
		ctrl:   ctrl,
		config: config,
		log:    log,
	}
}

func (f *HTTPService) ListenApp() error {
	f.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.Route(f.app, f.ctrl)

	f.log.Info("Application is running...")
	port := f.config.AppPort
	return f.app.Listen(fmt.Sprintf(":%s", port))
}
