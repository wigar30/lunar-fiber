package productimage

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type ProductImageRepository struct {
	db *driver.Database
}

func NewProductImageRepository(db *driver.Database) model.ProductImageRepositoryInterface {
	return &ProductImageRepository{
		db: db,
	}
}