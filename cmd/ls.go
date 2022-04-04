package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yguilai/nvm/env"
	"github.com/yguilai/nvm/version"
	"os"
	"path/filepath"
)

const versionsDir = "versions"

var lsCmd = &cli.Command{
	Name:  "ls",
	Usage: "listing has been installed nodejs versions",
	Action: func(ctx *cli.Context) error {
		verPath := filepath.Join(env.NvmHome(), versionsDir)
		dirs, err := os.ReadDir(verPath)
		if err != nil {
			return err
		}

		for _, dir := range dirs {
			ok, _ := version.IsVersionDir(dir.Name())
			if !ok {
				continue
			}
			// TODO: if a version used now, should flag it
			fmt.Println(dir.Name())
		}
		return nil
	},
}
