package config

import (
	"lunar-commerce-fiber/internal/presenter/http/controller"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var (
	AppSet = wire.NewSet(
		NewViper,
		NewLogger,
		NewFiber,
		NewConnMySql,
		NewListenApp,
	)

	ControllerSet = wire.NewSet(
		controller.NewController,
	)

	AllSet = wire.NewSet(
		AppSet,
		ControllerSet,
	)
)

func NewFiber(config *envConfigs) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: config.AppName,
		ErrorHandler: NewErrorHandler(),
		Prefork: config.AppEnv == "production",

		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

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
