package model

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type File struct {
	FilePath string `json:"file_path"`
}

type UploadUseCaseInterface interface {
	UploadImage(*fiber.Ctx, *multipart.FileHeader, string) (*File, error)
	UploadFile(*fiber.Ctx, *multipart.FileHeader, string) (*File, error)
}
