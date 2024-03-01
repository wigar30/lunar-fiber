package product

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/model"
)

type ProductRepository struct {
	db *driver.Database
}

func NewProductRepository(db *driver.Database) model.ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}