package ch

import (
	"os"

	"github.com/taisph/ch/pkg/clubhouse/v2"
	"github.com/taisph/ch/pkg/cmdr"
	"github.com/urfave/cli"
)

var (
	Client *v2.Client
)

func init() {
	app := cmdr.CmdRunner.App
	app.Name = "ch"
	app.Usage = "A Clubhouse CLI"
	app.HelpName = os.Args[0]
	app.Description = "Clubhouse CLI"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "token, t", Usage: "Clubhouse API token", EnvVar: "CLUBHOUSE_API_TOKEN"},
	}
	app.Before = func(c *cli.Context) error {
		Client, _ = v2.NewClient(c.GlobalString("token"))
		return nil
	}
}
