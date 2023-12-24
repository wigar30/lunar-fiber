package query

import (
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

func RoleSeed(db *driver.Database) error {
	db.Transaction(func(tx *gorm.DB) error {
		var count int64
		db.Model(&entity.Role{}).Where("name = ?", "Admin").Count(&count)

		if count == 0 {
			tx.Create(&entity.Role{
				Name: "Admin",
			})
		}

		db.Model(&entity.Role{}).Where("name = ?", "User").Count(&count)

		if count == 0 {
			tx.Create(&entity.Role{
				Name: "User",
			})
		}

		db.Model(&entity.Role{}).Where("name = ?", "Tenant Owner").Count(&count)

		if count == 0 {
			tx.Create(&entity.Role{
				Name: "Tenant Owner",
			})
		}

		db.Model(&entity.Role{}).Where("name = ?", "Tenant Member").Count(&count)

		if count == 0 {
			tx.Create(&entity.Role{
				Name: "Tenant Member",
			})
		}

		return nil
	})

	

	return nil
}