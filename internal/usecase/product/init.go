package product

import "lunar-commerce-fiber/internal/model"

type ProductUseCase struct {
	productRepo model.ProductRepositoryInterface
}

func NewProductUseCase(productRepo model.ProductRepositoryInterface) model.ProductUseCaseInterface {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}