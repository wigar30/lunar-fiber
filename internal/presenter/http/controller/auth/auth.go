package auth

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func (ac AuthController) Login(c *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := c.BodyParser(&req); err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := Validator.Struct(req)
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	} 

	resp, err := ac.authUseCase.Login(c, req)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
