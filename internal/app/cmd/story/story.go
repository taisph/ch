package story

import (
	"github.com/taisph/ch/pkg/cmdr"
	"github.com/urfave/cli"
)

var (
	storyCmd = cli.Command{
		Name:  "story",
		Usage: "Create or list stories",
		Subcommands: cli.Commands{
			storyCreateCmd,
			storyListCmd,
		},
	}
)

func init() {
	cmdr.CmdRunner.Use(storyCmd)
}
