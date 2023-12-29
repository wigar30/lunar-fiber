package config

import (
	"log"
	"lunar-commerce-fiber/internal/model"

	"go.uber.org/zap"
)

func NewLogger() *model.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to init logger: %s", err)
	}
	return &model.Logger{
		Logger: logger,
	}
}