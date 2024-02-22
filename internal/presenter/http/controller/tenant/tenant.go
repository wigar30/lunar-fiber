package tenant

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (tc TenantController) GetAllByAuth(c *fiber.Ctx) error {
	claims, _ := c.Locals("claims").(*model.JwtClaims)
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code: fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	userID, err := strconv.ParseInt(claims.ID, 10, 64)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	resp, err := tc.tenantUseCase.GetAllByAuth(userID, model.PaginationRequest{
		Page: page,
		Limit: limit,
	})
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code: err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
