package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yguilai/nvm/env"
	"github.com/yguilai/nvm/version"
	"sort"
)

var lsRemoteCmd = &cli.Command{
	Name:  "ls-remote",
	Usage: "listing all valid version of nodejs from remote server",
	Action: func(ctx *cli.Context) error {
		source, st := env.NvmSource()
		verArr, err := version.FindAllValidVersions(source, st)
		if err != nil {
			return err
		}
		sort.Slice(verArr, func(i, j int) bool {
			return verArr[i].Sort < verArr[j].Sort
		})

		for i, v := range verArr {
			fmt.Printf("%-15s", v.Name)
			if (i+1)%3 == 0 {
				fmt.Println()
			}
		}
		return nil
	},
}
