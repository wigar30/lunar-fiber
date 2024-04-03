package model

import (
	"lunar-commerce-fiber/internal/entity"
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
	Status        *bool  `json:"status" gorm:"column:status;" validate:"required,boolean"`
	Avatar        string `json:"avatar" gorm:"column:avatar;" validate:""`
}

type ProductResponse struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	TotalStock    *int                   `json:"totalStock"`
	TotalSold     int                    `json:"totalSold"`
	Price         *int                   `json:"price"`
	Description   string                 `json:"description,omitempty"`
	Specification string                 `json:"specification,omitempty"`
	Status        *bool                  `json:"status"`
	ProductImages []*entity.ProductImage `json:"productImages"`
}

type ProductRepositoryInterface interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB) error
	RollbackTransaction(*gorm.DB) error

	GetAllByTenantID(string, utils.Pagination) (*utils.Pagination, error)
	GetByID(string, string) (*entity.Product, error)
	CreateProduct(*gorm.DB, CreateProduct) (string, error)
}

type ProductUseCaseInterface interface {
	GetAllByTenantID(string, PaginationRequest) (*utils.Pagination, error)
	GetByID(string, string, string) (*ProductResponse, error)
	CreateProduct(CreateProduct) (string, error)
}
