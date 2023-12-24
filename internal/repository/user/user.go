package user

import (
	"errors"
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(id int64, relation bool) (*entity.User, error) {
	var user *entity.User
	var query *gorm.DB
	if relation {
		query = ur.db.Joins("Role").Joins("Status").First(&user, id)
	} else {
		query = ur.db.Select("statusId").First(&user, id)
	}
	err := query.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusNotFound,
			Message: err.Error(),
		}
	}
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return user, nil
}