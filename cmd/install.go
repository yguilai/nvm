package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var installCmd = &cli.Command{
	Name:  "install",
	Usage: "install a nodejs version",
	Action: func(ctx *cli.Context) error {
		fmt.Println(ctx)
		return nil
	},
}
