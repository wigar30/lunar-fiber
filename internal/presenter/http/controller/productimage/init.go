package productimage

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/go-playground/validator/v10"
)

type ProductImageController struct {
	productImageUseCase model.ProductImageUseCaseInterface

	validator *validator.Validate
}

func NewProductImageController(productImageUseCase model.ProductImageUseCaseInterface, validator *validator.Validate) *ProductImageController {
	return &ProductImageController{
		productImageUseCase: productImageUseCase,
		validator:           validator,
	}
}
