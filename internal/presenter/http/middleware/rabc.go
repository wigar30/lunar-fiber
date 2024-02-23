package middleware

import (
	"lunar-commerce-fiber/internal/model"
	"slices"

	"github.com/gofiber/fiber/v2"
)

type RbacMiddleware struct {
	roleRepo model.RoleRepositoryInterface

	config *model.EnvConfigs
}

func NewRbacMiddleware(roleRepo model.RoleRepositoryInterface, config *model.EnvConfigs) *RbacMiddleware {
	return &RbacMiddleware{
		roleRepo: roleRepo,
		config:   config,
	}
}

func (rm *RbacMiddleware) ValidateRoleUser(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if slices.Contains(roles, "All") {
			return c.Next()
		}

		claims, _ := c.Locals("claims").(*model.JwtClaims)
		role, err := rm.roleRepo.GetByID(int64(claims.Role))
		if err != nil {
			return model.OnError(c, &model.ErrorResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid Resource",
			})
		}

		if !slices.Contains(roles, role.Name) {
			return model.OnError(c, &model.ErrorResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid Resource",
			})
		}

		return c.Next()
	}
}

func (rm *RbacMiddleware) ValidateRoleMembership(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if slices.Contains(roles, "All") {
			return c.Next()
		}
		
		return c.Next()
	}
}
