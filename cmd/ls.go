package cmd

import "github.com/urfave/cli/v2"

var lsCmd = &cli.Command{
	Name:  "ls",
	Usage: "listing has been installed nodejs versions",
	Action: func(ctx *cli.Context) error {
		panic("implement me")
	},
}
