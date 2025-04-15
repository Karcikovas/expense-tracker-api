package cmd

import (
	"expense-tracker-api/internal/console"
	"github.com/spf13/cobra"
)

type Root struct {
	commands []console.Command
}

func NewRoot(commands []console.Command) *Root {
	return &Root{
		commands: commands,
	}
}

func (r *Root) Run(_ *cobra.Command, _ []string) {}

func (r *Root) GetCmd() *cobra.Command {
	rootCmd := console.NewConsoleCommand(
		"root",
		"",
		r.Run,
	)

	for _, command := range r.commands {
		rootCmd.AddCommand(command.GetCmd())
	}

	return rootCmd
}
