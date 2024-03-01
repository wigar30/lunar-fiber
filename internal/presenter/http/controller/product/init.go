package product

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	productUseCase model.ProductUseCaseInterface

	validator *validator.Validate
}

func NewProductController(productUseCase model.ProductUseCaseInterface, validator *validator.Validate) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
		validator:      validator,
	}
}
