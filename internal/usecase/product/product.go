package product

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"
)

func (pu *ProductUseCase) GetAllByTenantID(tenantID string, p model.PaginationRequest) (*utils.Pagination, error) {
	products, err := pu.productRepo.GetAllByTenantID(tenantID, utils.Pagination{
		Page: p.Page,
		Limit: p.Limit,
		Search: p.Search,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return products, nil
}

func(pu *ProductUseCase) CreateProduct(product model.CreateProduct) (string, error) {
	tx := pu.productRepo.BeginTransaction()

	ID, err := pu.productRepo.CreateProduct(tx, model.CreateProduct{
		TenantID: product.TenantID,
		Name: product.Name,
		TotalStock: product.TotalStock,
		Description: product.Description,
		Specification: product.Description,
		Price: product.Price,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		return "", &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return ID, nil
}
