package app

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/internal/presenter/http/controller"
	authCtrl "lunar-commerce-fiber/internal/presenter/http/controller/auth"
	roleCtrl "lunar-commerce-fiber/internal/presenter/http/controller/role"
	userCtrl "lunar-commerce-fiber/internal/presenter/http/controller/user"
	"lunar-commerce-fiber/internal/presenter/http/middleware"
	roleRepo "lunar-commerce-fiber/internal/repository/role"
	userRepo "lunar-commerce-fiber/internal/repository/user"
	authUC "lunar-commerce-fiber/internal/usecase/auth"
	roleUC "lunar-commerce-fiber/internal/usecase/role"
	userUC "lunar-commerce-fiber/internal/usecase/user"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var (
	AppSet = wire.NewSet(
		config.NewViper,
		config.NewLogger,
		config.NewListenApp,
		driver.NewConnMySql,
		NewFiber,
	)

	ControllerSet = wire.NewSet(
		controller.NewController,
		roleCtrl.NewRoleController,
		authCtrl.NewAuthController,
		userCtrl.NewUserController,
	)

	UseCaseSet = wire.NewSet(
		roleUC.NewRoleUseCase,
		authUC.NewAuthUseCase,
		userUC.NewUserUseCase,
	)

	RepositorySet = wire.NewSet(
		roleRepo.NewRoleRepository,
		userRepo.NewUserRepository,
	)

	MiddlewareSet = wire.NewSet(
		middleware.NewMiddleware,
	)

	AllSet = wire.NewSet(
		AppSet,
		ControllerSet,
		UseCaseSet,
		RepositorySet,

		MiddlewareSet,
	)
)

func NewFiber(config *model.EnvConfigs) *fiber.App {
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
