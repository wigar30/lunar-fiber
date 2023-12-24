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