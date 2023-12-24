package app

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/presenter/http/controller"
	authCtrl "lunar-commerce-fiber/internal/presenter/http/controller/auth"
	roleCtrl "lunar-commerce-fiber/internal/presenter/http/controller/role"
	roleRepo "lunar-commerce-fiber/internal/repository/role"
	userRepo "lunar-commerce-fiber/internal/repository/user"
	authUC "lunar-commerce-fiber/internal/usecase/auth"
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
		authCtrl.NewAuthController,
	)

	UseCaseSet = wire.NewSet(
		roleUC.NewRoleUseCase,
		authUC.NewAuthUseCase,
	)

	RepositorySet = wire.NewSet(
		roleRepo.NewRoleRepository,
		userRepo.NewUserRepository,
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
