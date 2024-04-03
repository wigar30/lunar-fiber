package upload

import (
	"fmt"
	"lunar-commerce-fiber/internal/model"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (uu *UploadUseCase) UploadImage(c *fiber.Ctx, image *multipart.FileHeader, category string) (*model.File, error) {
	destinationFolder := fmt.Sprintf("storages/%s", category)
	destination := fmt.Sprintf("./%s/%s_%s", destinationFolder, strconv.FormatInt(time.Now().UnixMilli(), 10), image.Filename)

	if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
			return nil, &model.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if err := c.SaveFile(image, destination); err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &model.File{
		FilePath: destination,
	}, nil
}

func (uu *UploadUseCase) UploadFile(c *fiber.Ctx, image *multipart.FileHeader, category string) (*model.File, error) {
	now := time.Now()
	destinationFolder := fmt.Sprintf("storages/%s/%s/%s/%s", category, strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()))
	destination := fmt.Sprintf("./%s/%s_%s", destinationFolder, strconv.FormatInt(time.Now().UnixMilli(), 10), image.Filename)

	if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
			return nil, &model.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	if err := c.SaveFile(image, destination); err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &model.File{
		FilePath: destination,
	}, nil
}
