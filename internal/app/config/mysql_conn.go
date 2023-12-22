package config

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnMySql(config *envConfigs, logrus *logrus.Logger) (db *gorm.DB, err error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbDatabase)

	w := logrus.Writer()
	defer w.Close()

	newLogger := logger.New(
		log.New(w, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	logrus.Info("Database connected")

	return
}
