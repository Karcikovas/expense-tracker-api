package cmd

import "expense-tracker-api/internal/console"

func ProvideCMD(
	server *ServeCMD,
) []console.Command {
	return []console.Command{server}
}
