package project

import (
	"github.com/taisph/ch/pkg/cmdr"
	"github.com/urfave/cli"
)

var (
	projectCmd = cli.Command{
		Name:  "project",
		Usage: "List projects",
		Subcommands: cli.Commands{
			projectListCmd,
		},
	}
)

func init() {
	cmdr.CmdRunner.Use(projectCmd)
}
