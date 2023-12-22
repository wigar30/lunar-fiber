package main

import (
	"lunar-commerce-fiber/internal/app"
)

func main() {
	app := app.NewWire()
	app.ListenApp()
}
