package productimage

import "lunar-commerce-fiber/internal/model"

type ProductImageUseCase struct {
	productImageRepo model.ProductImageRepositoryInterface

	uploadUseCase model.UploadUseCaseInterface
}

func NewProductImageUseCase(productImageRepo model.ProductImageRepositoryInterface, uploadUseCase model.UploadUseCaseInterface) model.ProductImageUseCaseInterface {
	return &ProductImageUseCase{
		productImageRepo: productImageRepo,
		uploadUseCase:    uploadUseCase,
	}
}
