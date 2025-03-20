package console

import "github.com/spf13/cobra"

type Command interface {
	Run(_ *cobra.Command, _ []string)
	GetCmd() *cobra.Command
}
