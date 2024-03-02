package tenant

import (
	"fmt"
	"lunar-commerce-fiber/internal/model"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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

func (tc TenantController) GetByID(c *fiber.Ctx) error {
	claims, _ := c.Locals("claims").(*model.JwtClaims)
	paramsId := c.Params("id")

	resp, err := tc.tenantUseCase.GetByID(claims.ID, paramsId)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}

func (tc TenantController) CreateTenant(c *fiber.Ctx) error {
	payload := new(model.CreateTenant)

	if err := c.BodyParser(payload); err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	errs := tc.validator.Struct(payload)
	if errs != nil {
		errMsgs := make([]string, 0)
		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: Needs to implement '%s'",
				err.Field(),
				err.Tag(),
			))
		}
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errMsgs, " | "),
		})
	}

	claims, _ := c.Locals("claims").(*model.JwtClaims)
	userID, err := strconv.ParseInt(claims.ID, 10, 64)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	resp, err := tc.tenantUseCase.CreateTenant(userID, *payload)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
