package main

import (
	"expense-tracker-api/internal/http"
	"expense-tracker-api/internal/logger"
)

type Application struct {
	logger logger.Service
	server *http.Server
}

func NewApplication(
	logger logger.Service,
	server *http.Server,
) *Application {
	return &Application{
		logger: logger,
		server: server,
	}
}

func (a *Application) Start() {
	a.server.Start()
}
