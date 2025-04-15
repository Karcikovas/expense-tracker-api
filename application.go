package main

import (
	"context"
	"expense-tracker-api/cmd"
	"expense-tracker-api/internal/http"
	"expense-tracker-api/internal/logger"
	"os"
	"os/signal"
)

type Application struct {
	logger logger.Service
	server *http.Server
	cmd    *cmd.Root
}

func NewApplication(
	logger logger.Service,
	server *http.Server,
	cmd *cmd.Root,
) *Application {
	return &Application{
		logger: logger,
		server: server,
		cmd:    cmd,
	}
}

func (a *Application) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	rootCmd := a.cmd.GetCmd()
	rootCmd.SetContext(ctx)

	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
