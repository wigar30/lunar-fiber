package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	DbConnection string `mapstructure:"DB_CONNECTION"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbDatabase   string `mapstructure:"DB_DATABASE"`
	DbUsername   string `mapstructure:"DB_USERNAME"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`

	LogLever string `mapstructure:"DB_PASSWORD"`
}

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath("./")
	config.AutomaticEnv()


	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return config
}
