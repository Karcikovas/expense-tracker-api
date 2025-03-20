//go:build wireinject
// +build wireinject

package main

import (
	"example.com/m/v2/internal/logger"
	"github.com/google/wire"
)

func NewApp() (*Application, error) {
	wire.Build(
		logger.DepSet,
		NewApplication,
	)
	return &Application{}, nil
}
