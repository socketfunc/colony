package main

import (
	"os"

	"github.com/socketfunc/colony/cmd/colony-worker/app"
)

func main() {
	cmd := app.NewAppCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
