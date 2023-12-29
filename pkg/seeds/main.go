package main

import (
	"lunar-commerce-fiber/internal/app/config"
	"lunar-commerce-fiber/internal/app/driver"
	"lunar-commerce-fiber/pkg/seeds/query"
)

func main() {
	envConfigs := config.NewViper()
	zaplog := config.NewLogger()
	db := driver.NewConnMySql(envConfigs, zaplog)

	var err error
	err = query.RoleSeed(db)
	if err != nil {
		zaplog.Fatal("error: failed to seed role data")
	}

	err = query.StatusSeed(db)
	if err != nil {
		zaplog.Fatal("error: failed to seed status data")
	}

	err = query.UserSeed(db)
	if err != nil {
		zaplog.Fatal("error: failed to seed user data")
	}

	err = query.TenantSeed(db)
	if err != nil {
		zaplog.Fatal("error: failed to seed user data")
	}
}