package middleware

import (
	"lunar-commerce-fiber/internal/model"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	userRepo model.UserRepositoryInterface

	config *model.EnvConfigs
}

func NewAuthMiddleware(userRepo model.UserRepositoryInterface, config *model.EnvConfigs) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo: userRepo,
		config:   config,
	}
}

func (am *AuthMiddleware) ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Unauthorized",
			})
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		jwtClaims, err := jwt.ParseWithClaims(token, &model.JwtClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &model.ErrorResponse{
					Code:    401,
					Message: "Invalid Token",
				}
			}
			return []byte(am.config.JwtSecret), nil
		})
		if err != nil {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
		}

		claims, ok := jwtClaims.Claims.(*model.JwtClaims)
		claimID, err := strconv.ParseInt(claims.ID, 10, 64)
		if err != nil {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
		}

		user, err := am.userRepo.GetUserByID(claimID, false)
		if err != nil {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
		}

		if user.StatusID != "1" {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "User not verified",
			})
		}

		if !ok || !jwtClaims.Valid {
			return model.OnError(c, &model.ErrorResponse{
				Code:    401,
				Message: "Invalid Token",
			})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
