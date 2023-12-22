package role

import (
	"lunar-commerce-fiber/internal/model"

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