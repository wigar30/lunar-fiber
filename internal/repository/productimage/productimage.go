package productimage

import (
	"lunar-commerce-fiber/internal/model"

	// "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (pr *ProductImageRepository) BeginTransaction() *gorm.DB {
	return pr.db.Begin()
}

func (pr *ProductImageRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (pr *ProductImageRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (pr *ProductImageRepository) CreateBulkProductImage(tx *gorm.DB, productImages []model.CreateProductImage) error {
	for _, p := range productImages {
		p.ImageFile = nil
	}

	println(productImages)
	// err := tx.Create(&productImages).Error
	// if err != nil {
	// 	return &model.ErrorResponse{
	// 		Code:    fiber.StatusInternalServerError,
	// 		Message: err.Error(),
	// 	}
	// }

	return nil
}
