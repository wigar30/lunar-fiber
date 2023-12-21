package main

import (
	"fmt"
	"lunar-commerce-fiber/internal/app/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)

	app, _ := config.NewWire()

	port := viperConfig.GetString("APP_PORT")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
	
}
