package role

import (
	"lunar-commerce-fiber/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (ru RoleUseCase) GetAll(c *fiber.Ctx) (*model.RolesResponse, error) {
	roles, err := ru.roleRepo.GetAll()
	if err != nil {
		return nil, &model.ErrorResponse{
			Code: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	var roleResponses []*model.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, &model.RoleResponse{
			// Copy values from role to RoleResponse
			ID:   role.ID,
			Name: role.Name,
			// Add other fields as needed
		})
	}

	return &model.RolesResponse{
		Items: roleResponses,
	}, nil
}