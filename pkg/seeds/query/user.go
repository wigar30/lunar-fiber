package query

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/entity"

	"gorm.io/gorm"
)

func UserSeed(db *config.Database) error {
	db.Transaction(func(tx *gorm.DB) error {
		var count int64
		db.Model(&entity.User{}).Where(`"roleId" = ?`, "1").Count(&count)

		if count == 0 {
			tx.Create(&entity.User{
				Email: "admin@example.com",
				Name: "Admin Local",
				Password: "123456789",
				RoleID: "1",
				StatusID: "1",
			})
		}

		db.Model(&entity.User{}).Where(`"roleId" = ?`, "2").Count(&count)

		if count == 0 {
			tx.Create(&entity.User{
				Email: "user@example.com",
				Name: "User Local",
				Password: "123456789",
				RoleID: "2",
				StatusID: "1",
			})
		}

		db.Model(&entity.User{}).Where(`"statusId" = ?`, "2").Count(&count)

		if count == 0 {
			tx.Create(&entity.User{
				Email: "user.inactive@example.com",
				Name: "User Inactive Local",
				Password: "123456789",
				RoleID: "2",
				StatusID: "2",
			})
		}

		return nil
	})

	

	return nil
}