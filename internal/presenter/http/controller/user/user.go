package user

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (uc UserController) GetProfile(c *fiber.Ctx) error {
	claims, _ := c.Locals("claims").(*model.JwtClaims)

	userID, err := strconv.ParseInt(claims.ID, 10, 64)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	resp, err := uc.userUseCase.GetUserByID(userID)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code: err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
