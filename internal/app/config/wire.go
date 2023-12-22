//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/google/wire"
)

func NewWire() HTTPServiceInterface {
	wire.Build(AllSet)

	return &HTTPService{}
}