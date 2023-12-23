package model

import (
	"lunar-commerce-fiber/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type RoleResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
}

type RolesResponse struct {
	Items []*RoleResponse `json:"items"`
}

type RoleRepositoryInterface interface {
	// Create new role
	//  @param role *Role, role object
	GetAll() ([]*entity.Role, error)
	GetByID(int64) (*entity.Role, error)
}

type RoleUseCaseInterface interface {
	GetAll(c *fiber.Ctx) (*RolesResponse, error)
	GetByID(int64) (*RoleResponse, error)
}
