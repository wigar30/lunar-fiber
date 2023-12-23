package role

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (rc RoleController) GetAll(c *fiber.Ctx) error {
	resp, err := rc.roleUseCase.GetAll(c)
	if err != nil {
		return model.OnError(c, fiber.StatusInternalServerError, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}

func (rc RoleController) GetByID(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		return model.OnError(c, fiber.StatusInternalServerError, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	resp, err := rc.roleUseCase.GetByID(id)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, err.Code, &model.ErrorResponse{
			Code: err.Code,
			Message: err.Error(),
		})
	} else if !errC {
			return model.OnError(c, fiber.StatusInternalServerError, &model.ErrorResponse{
				Code: fiber.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return model.OnSuccess(c, resp)
}