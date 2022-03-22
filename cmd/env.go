package cmd

import "github.com/urfave/cli/v2"

var envCmd = &cli.Command{
	Name:  "env",
	Usage: "print environment variables about nvm and node",
}
