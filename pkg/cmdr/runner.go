package cmdr

import (
	"os"

	"github.com/urfave/cli"
)

var (
	CmdRunner = NewRunner()
)

type Runner struct {
	App *cli.App
}

func NewRunner() *Runner {
	return &Runner{App: cli.NewApp()}
}

func (r *Runner) Run() error {
	return r.App.Run(os.Args)
}

func (r *Runner) Use(command cli.Command) {
	r.App.Commands = append(r.App.Commands, command)
}
