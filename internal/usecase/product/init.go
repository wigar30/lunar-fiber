package product

import "lunar-commerce-fiber/internal/model"

type ProductUseCase struct {
	productRepo    model.ProductRepositoryInterface
	membershipRepo model.MembershipRepositoryInterface
}

func NewProductUseCase(productRepo model.ProductRepositoryInterface, membershipRepo model.MembershipRepositoryInterface) model.ProductUseCaseInterface {
	return &ProductUseCase{
		productRepo:    productRepo,
		membershipRepo: membershipRepo,
	}
}
