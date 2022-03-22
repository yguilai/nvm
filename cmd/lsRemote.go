package cmd

import (
	"github.com/urfave/cli/v2"
)

var lsRemoteCmd = &cli.Command{
	Name:  "ls-remote",
	Usage: "list all valid versions of nodejs",
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
