package main

import (
	"fmt"
	"os"
)

func main() {
	app, err := NewApp()

	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)

		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	app.Start()

	os.Exit(0)
}
