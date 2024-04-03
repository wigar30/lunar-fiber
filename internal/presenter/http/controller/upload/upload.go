package upload

import (
	"lunar-commerce-fiber/internal/model"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func (uc UploadController) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Validate size
	if file.Size > int64(uc.config.ImageSizeLimit) {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Request Entity Too Large",
		})
	}

	// Validate mime type
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/webp", "image/png"}
	if !slices.Contains(allowedTypes, file.Header.Get("Content-Type")) {
		return model.OnError(c, &model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Type",
		})
	}

	resp, err := uc.uploadUseCase.UploadImage(c, file, "avatar")
	if err, errC := err.(*model.ErrorResponse); errC {
		return model.OnError(c, &model.ErrorResponse{
			Code:    err.Code,
			Message: err.Error(),
		})
	}

	return model.OnSuccess(c, resp)
}
