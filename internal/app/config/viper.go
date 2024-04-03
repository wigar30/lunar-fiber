package config

import (
	"fmt"
	"log"
	"lunar-commerce-fiber/internal/model"

	"github.com/spf13/viper"
)

func NewViper() *model.EnvConfigs {
	config := viper.New()

	config.AddConfigPath(".")
	config.SetConfigType("env")
	config.SetConfigFile(".env")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	config.BindEnv("AppName", "APP_NAME")
	config.BindEnv("AppEnv", "APP_ENV")
	config.BindEnv("AppPort", "APP_PORT")

	config.BindEnv("DbConnection", "DB_CONNECTION")
	config.BindEnv("DbHost", "DB_HOST")
	config.BindEnv("DbPort", "DB_PORT")
	config.BindEnv("DbDatabase", "DB_DATABASE")
	config.BindEnv("DbUsername", "DB_USERNAME")
	config.BindEnv("DbPassword", "DB_PASSWORD")

	config.BindEnv("FileSizeLimit", "FILE_SIZE_LIMIT")
	config.BindEnv("ImageSizeLimit", "IMAGE_SIZE_LIMIT")

	config.BindEnv("LogLevel", "LOG_LEVEL")

	config.BindEnv("JwtExpiredIn", "AUTH_JWT_SECRET")
	config.BindEnv("JwtSecret", "AUTH_JWT_TOKEN_EXPIRES_IN")

	var env *model.EnvConfigs
	if err := config.Unmarshal(&env); err != nil {
		fmt.Printf("Error unmarshalling config: %s\n", err)
	}

	return env
}
