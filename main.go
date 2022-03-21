package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yguilai/nvm/cmd"
	"os"
)

var app = &cli.App{
	Name:     "nvm",
	Usage:    "nodejs version manager",
	Version:  cmd.Version(),
	Commands: cmd.Commands,
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
