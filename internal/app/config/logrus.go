package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(config *envConfigs) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.LogLever))

	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}