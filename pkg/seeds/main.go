package main

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/pkg/seeds/query"
)

func main() {
	envConfigs := config.NewViper()
	logger := config.NewLogger(envConfigs)
	db := driver.NewConnMySql(envConfigs, logger)

	var err error
	err = query.RoleSeed(db)
	if err != nil {
		logger.Fatalf("error: failed to seed role data")
	}

	err = query.StatusSeed(db)
	if err != nil {
		logger.Fatalf("error: failed to seed status data")
	}

	err = query.UserSeed(db)
	if err != nil {
		logger.Fatalf("error: failed to seed user data")
	}

	err = query.TenantSeed(db)
	if err != nil {
		logger.Fatalf("error: failed to seed user data")
	}
}