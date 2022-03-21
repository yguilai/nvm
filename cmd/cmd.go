package cmd

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	installCmd,
	uninstallCmd,
	lsCmd,
	lsRemoteCmd,
	useCmd,
	cleanCmd,
}
