package query

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

func StatusSeed(db *driver.Database) error {
	db.Transaction(func(tx *gorm.DB) error {
		var count int64
		db.Model(&entity.Status{}).Where("name = ?", "Active").Count(&count)

		if count == 0 {
			tx.Create(&entity.Status{
				Name: "Active",
			})
		}

		db.Model(&entity.Status{}).Where("name = ?", "Inactive").Count(&count)

		if count == 0 {
			tx.Create(&entity.Status{
				Name: "Inactive",
			})
		}

		return nil
	})

	

	return nil
}