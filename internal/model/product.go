package model

import (
	"lunar-commerce-fiber/pkg/utils"

	"gorm.io/gorm"
)

type CreateProduct struct {
	TenantID      string `json:"tenantId,omitempty" gorm:"not null;column:tenantId" validate:"required"`
	Name          string `json:"name" gorm:"not null" validate:"required,min=3"`
	TotalStock    *int   `json:"totalStock,omitempty" gorm:"not null;column:totalStock;default:0" validate:"required,numeric"`
	Price         *int   `json:"price,omitempty" gorm:"not null;column:price;default:0" validate:"required,numeric"`
	Description   string `json:"description,omitempty" gorm:"column:description;" validate:"required"`
	Specification string `json:"specification,omitempty" gorm:"column:specification;" validate:"required"`
}

type ProductRepositoryInterface interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB) error
	RollbackTransaction(*gorm.DB) error
	
	GetAllByTenantID(string, utils.Pagination) (*utils.Pagination, error)
	CreateProduct(*gorm.DB, CreateProduct) (string, error)
}

type ProductUseCaseInterface interface {
	GetAllByTenantID(string, PaginationRequest) (*utils.Pagination, error)
	CreateProduct(CreateProduct) (string, error)
}
