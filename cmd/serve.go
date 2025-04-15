package cmd

import (
	"expense-tracker-api/internal/console"
	"expense-tracker-api/internal/http"
	"github.com/spf13/cobra"
)

type ServeCMD struct {
	server *http.Server
}

func NewServeCMD(
	server *http.Server,
) *ServeCMD {
	return &ServeCMD{
		server: server,
	}
}

func (r *ServeCMD) Run(_ *cobra.Command, _ []string) {
	r.server.Start()
}

func (r *ServeCMD) GetCmd() *cobra.Command {
	return console.NewConsoleCommand(
		"serve",
		"Serve http server",
		func(cmd *cobra.Command, args []string) {
			r.Run(cmd, args)
		},
	)
}
