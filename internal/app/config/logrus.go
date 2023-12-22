package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(config *EnvConfigs) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.LogLever))

	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}