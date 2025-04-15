//go:build wireinject
// +build wireinject

package main

import (
	"expense-tracker-api/cmd"
	"expense-tracker-api/internal/config"
	"expense-tracker-api/internal/http"
	"expense-tracker-api/internal/logger"
	"github.com/google/wire"
)

func NewApp() (*Application, error) {
	wire.Build(
		cmd.DepSet,
		logger.DepSet,
		http.DepSet,
		config.DepSet,
		NewApplication,
	)
	return &Application{}, nil
}
