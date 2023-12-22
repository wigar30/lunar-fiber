package main

import (
	"lunar-commerce-fiber/internal/app/config"
)

func main() {
	app := config.NewWire()
	app.ListenApp()
}
