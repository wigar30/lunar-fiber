package config

import (
	"log"

	"github.com/spf13/viper"
)

var Env *EnvConfigs

type EnvConfigs struct {
	AppName string
	AppEnv  string
	AppPort string

	DbConnection string
	DbHost       string
	DbPort       string
	DbDatabase   string
	DbUsername   string
	DbPassword   string

	LogLever int32 `mapstructure:"DB_PASSWORD"`
}

func NewViper() *EnvConfigs {
	config := viper.New()

	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath("./")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return &EnvConfigs{
		AppName: config.GetString("APP_NAME"),
		AppEnv:  config.GetString("APP_ENV"),
		AppPort: config.GetString("APP_PORT"),

		DbConnection: config.GetString("DB_CONNECTION"),
		DbHost:       config.GetString("DB_HOST"),
		DbPort:       config.GetString("DB_PORT"),
		DbDatabase:   config.GetString("DB_DATABASE"),
		DbUsername:   config.GetString("DB_USERNAME"),
		DbPassword:   config.GetString("DB_PASSWORD"),

		LogLever: config.GetInt32("LOG_LEVEL"),
	}
}
