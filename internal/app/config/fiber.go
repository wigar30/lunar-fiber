package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	AppSet = wire.NewSet(
		NewViper,
		NewLogger,
		NewFiber,
	)

	AllSet = wire.NewSet(
		AppSet,
	)
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: config.GetString("APP_NAME"),
		ErrorHandler: NewErrorHandler(),
		Prefork: config.GetString("APP_ENV") == "production",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	logrus.Info("Application is running...")
	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
