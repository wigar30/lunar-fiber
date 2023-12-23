package query

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

func TenantSeed(db *config.Database) error {
	db.Transaction(func(tx *gorm.DB) error {
		var count int64
		var tenant entity.Tenant
		db.Model(&tenant).Where("name = ?", "Tenant User").Count(&count)

		if count == 0 {
			tenant := entity.Tenant{
				Name: "Tenant User",
				TotalProduct: new(int),
			}
			tx.Create(&tenant)

			var user entity.User
			err := db.Model(&user).Where("name = ?", "User Local").First(&user).Error
			if err != nil {
				return err
			}
	
			tx.Create(&entity.Membership{
				UserRefer: user.ID,
				TenantRefer: tenant.ID,
				RoleID: user.RoleID,
				StatusID: user.StatusID,
			})
		}


		return nil
	})

	

	return nil
}