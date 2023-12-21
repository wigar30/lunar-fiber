//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func NewWire() (*fiber.App, error) {
	wire.Build(wire.NewSet(AllSet))

	return nil, nil
}