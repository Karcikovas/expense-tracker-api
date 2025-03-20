//go:build wireinject
// +build wireinject

package main

import (
	"expense-tracker-api/internal/http"
	"expense-tracker-api/internal/logger"
	"github.com/google/wire"
)

func NewApp() (*Application, error) {
	wire.Build(
		logger.DepSet,
		http.DepSet,
		NewApplication,
	)
	return &Application{}, nil
}
