package role

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (rc RoleController) GetAll(c *fiber.Ctx) error {
	// claimsCtx, ok := c.Locals("claims").(*model.JwtClaims)
	resp, err := rc.roleUseCase.GetAll(c)
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}

func (rc RoleController) GetByID(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	id, err := strconv.ParseInt(paramsId, 10, 64)
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	resp, err := rc.roleUseCase.GetByID(id)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
