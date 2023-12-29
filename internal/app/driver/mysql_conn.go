package driver

import (
	"fmt"
	"log"
	"lunar-commerce-fiber/internal/model"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewConnMySql(config *model.EnvConfigs, zaplog *model.Logger) *Database {
	// conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.database"))
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbDatabase)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		TranslateError: true,
	})

	if err != nil {
		zaplog.Error("failed to connect database: ", zapcore.Field{
			String: err.Error(),
		})
	}

	zaplog.Info("Database connected")

	return &Database{
		DB: db,
	}
}
