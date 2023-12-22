package role

import (
	"lunar-commerce-fiber/internal/entity"
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (rr *RoleRepository) GetAll() ([]*entity.Role, error) {
	var roles []*entity.Role
	err := rr.db.Find(&roles).Error
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return roles, nil
}