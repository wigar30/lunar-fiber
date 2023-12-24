package auth

import (
	"lunar-commerce-fiber/internal/app/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(config *config.EnvConfigs, id string, role int) (string, error) {
	token_lifespan, err := strconv.Atoi(config.JwtExpiredIn)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	})

	return token.SignedString([]byte(config.JwtSecret))
}