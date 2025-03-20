package main

import (
	"example.com/m/v2/internal/logger"
	"log"
)

type Application struct {
	logger logger.Service
}

func NewApplication(
	logger logger.Service,
) *Application {
	return &Application{
		logger: logger,
	}
}

func (a *Application) Start() {
	log.Println("asdasdasda")
}
