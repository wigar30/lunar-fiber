package model

import (
	"lunar-commerce-fiber/internal/entity"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateProductImage struct {
	TenantID  string                `json:"tenantId" gorm:"not null;column:tenantId" validate:"required"`
	ProductID string                `json:"productId" gorm:"not null;column:productId" validate:"required"`
	StatusID  string                `json:"statusId" gorm:"not null;column:statusId" validate:"required"`
	Title     string                `json:"title" gorm:"not null" validate:"required"`
	Image     string                `json:"image" gorm:"column:image;"`
	ImageFile *multipart.FileHeader `json:"imageFile" validate:"required"`
}

type ProductImageResponse struct {
	ID            string                 `json:"id"`
	Description   string                 `json:"description,omitempty"`
	Specification string                 `json:"specification,omitempty"`
	StatusID      string                 `json:"statusId,omitempty"`
	Status        *entity.Status         `json:"status,omitempty"`
	ProductImages []*entity.ProductImage `json:"productImages"`
}

type ProductImageRepositoryInterface interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(*gorm.DB) error
	RollbackTransaction(*gorm.DB) error

	// GetAllByProductID(string, utils.Pagination) (*utils.Pagination, error)
	CreateBulkProductImage(*gorm.DB, []CreateProductImage) error
}

type ProductImageUseCaseInterface interface {
	CreateBulkProductImage(*fiber.Ctx, []CreateProductImage) (bool, error)
}
