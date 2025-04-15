package console

import "github.com/spf13/cobra"

type CommandAction func(cmd *cobra.Command, args []string)

func NewConsoleCommand(name string, description string, runCommand CommandAction) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: description,
		Run:   runCommand,
	}
}
