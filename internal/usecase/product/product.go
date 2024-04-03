package product

import (
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
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

func (pu *ProductUseCase) GetByID(userID string, tenantID string, productID string) (*model.ProductResponse, error) {
	isTenantAuth, err := pu.membershipRepo.CheckIsAuthTenant(userID, tenantID)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}
	if (!isTenantAuth) {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusNotFound,
			Message: "record not found",
		}
	}

	product, err := pu.productRepo.GetByID(tenantID, productID)
	if errC, ok := err.(*model.ErrorResponse); ok {
		return nil, &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return &model.ProductResponse{
		ID: product.ID,
		Name: product.Name,
		TotalStock: product.TotalStock,
		TotalSold: *product.TotalSold,
		Price: product.Price,
		Description: product.Description,
		Specification: product.Specification,
		Status: product.Status,
		ProductImages: product.ProductImages,
	}, nil
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
		Status: product.Status,
	})
	if errC, ok := err.(*model.ErrorResponse); ok {
		pu.productRepo.RollbackTransaction(tx)
		return "", &model.ErrorResponse{
			Code: errC.Code,
			Message: errC.Error(),
		}
	}

	return ID, pu.productRepo.CommitTransaction(tx)
}
