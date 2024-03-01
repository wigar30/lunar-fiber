package product

import (
	"fmt"
	"lunar-commerce-fiber/internal/model"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (pc *ProductController) GetAllByTenantID(c *fiber.Ctx) error {
	paramsId := c.Params("id")
	search := c.Query("search")
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	resp, err := pc.productUseCase.GetAllByTenantID(paramsId, model.PaginationRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	payload := new(model.CreateProduct)

	if err := c.BodyParser(payload); err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	errs := pc.validator.Struct(payload)
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

	resp, err := pc.productUseCase.CreateProduct(*payload)
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
