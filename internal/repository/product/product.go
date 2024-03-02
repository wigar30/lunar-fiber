package product

import (
	"errors"
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"
	"lunar-commerce-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (pr *ProductRepository) BeginTransaction() *gorm.DB {
	return pr.db.Begin()
}

func (pr *ProductRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (pr *ProductRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (pr *ProductRepository) GetAllByTenantID(tenantID string, p utils.Pagination) (*utils.Pagination, error) {
	var products []entity.Product

	db := pr.db.
		Where(&entity.Product{TenantID: tenantID})

	if p.Search != "" {
		search := "%" + p.Search + "%"
		db = db.Where(db.Where("name like ?", search).Or("description like ?", search))
	}

	err := db.
		Scopes(utils.Paginate(&p, pr.db.Where(&entity.Product{TenantID: tenantID}), &products)).
		Find(&products).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	p.Items = products

	return &p, nil
}

func (pr *ProductRepository) CreateProduct(tx *gorm.DB, product model.CreateProduct) (string, error) {
	productResult := &entity.Product{
		TenantID: product.TenantID,
		Name: product.Name,
		TotalStock: product.TotalStock,
		Description: product.Description,
		Specification: product.Specification,
		Price: product.Price,
	}
	err := tx.Create(&productResult).Error
	if err != nil {
		return "", &model.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return productResult.ID, nil
}
