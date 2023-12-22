//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"lunar-commerce-fiber/internal/app/config"

)

func NewWire() config.HTTPServiceInterface {
	wire.Build(AllSet)

	return &config.HTTPService{}
}