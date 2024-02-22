package query

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

func LevelTenantSeed(db *driver.Database) error {
	db.Transaction(func(tx *gorm.DB) error {
		var count int64
		db.Model(&entity.LevelTenant{}).Where("level = ?", "Tenant").Count(&count)

		if count == 0 {
			tx.Create(&entity.LevelTenant{
				Level: "Tenant",
			})
		}

		db.Model(&entity.LevelTenant{}).Where("level = ?", "Tenant Plus").Count(&count)

		if count == 0 {
			tx.Create(&entity.LevelTenant{
				Level: "Tenant Plus",
			})
		}

		db.Model(&entity.LevelTenant{}).Where("level = ?", "Mall").Count(&count)

		if count == 0 {
			tx.Create(&entity.LevelTenant{
				Level: "Mall",
			})
		}

		return nil
	})

	return nil
}
