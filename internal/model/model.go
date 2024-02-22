package model

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Data   interface{}   `json:"data"`
	Status string        `json:"status,omitempty"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"details,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

func OnSuccess(c *fiber.Ctx, data interface{}) error {
	resp := ApiResponse{
		Status: "OK",
		Data:   data,
	}

	return c.JSON(resp)
}

func OnError(ctx *fiber.Ctx, err *ErrorResponse) error {

	return ctx.Status(err.Code).JSON(err)
}
