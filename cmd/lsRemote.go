package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yguilai/nvm/version"
)

var lsRemoteCmd = &cli.Command{
	Name:  "ls-remote",
	Usage: "list all valid versions of nodejs",
	Action: func(ctx *cli.Context) error {
		_, err := version.FindAllValidVersions("https://registry.npmmirror.com/-/binary/node/", version.Taobao)
		if err != nil {
			panic(err)
		}
		//dump.P(versions)
		return nil
	},
}
