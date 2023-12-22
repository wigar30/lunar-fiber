package app

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/presenter/http/controller"
	roleCtrl "lunar-commerce-fiber/internal/presenter/http/controller/role"
	roleRepo "lunar-commerce-fiber/internal/repository/role"
	roleUC "lunar-commerce-fiber/internal/usecase/role"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var (
	AppSet = wire.NewSet(
		config.NewViper,
		config.NewLogger,
		config.NewConnMySql,
		config.NewListenApp,
		NewFiber,
	)

	ControllerSet = wire.NewSet(
		controller.NewController,
		roleCtrl.NewRoleController,
	)

	UseCaseSet = wire.NewSet(
		roleUC.NewRoleUseCase,
	)

	RepositorySet = wire.NewSet(
		roleRepo.NewRoleRepository,
	)

	AllSet = wire.NewSet(
		AppSet,
		ControllerSet,
		UseCaseSet,
		RepositorySet,
	)
)

func NewFiber(config *config.EnvConfigs) *fiber.App {
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
