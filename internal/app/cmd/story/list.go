package story

import (
	"fmt"
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/taisph/ch/internal/app/cmd/ch"
	"github.com/urfave/cli"
)

var (
	storyListCmd = cli.Command{Name: "list", Action: List, Usage: "List stories in specified project", Flags: []cli.Flag{
		cli.Int64Flag{Name: "project-id, p", Usage: "Id of project to create story in"},
	}}
)

func List(c *cli.Context) error {
	stories, err := ch.Client.Stories().List(c.Int64("project-id"))
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	w := tabwriter.NewWriter(c.App.Writer, 0, 0, 2, ' ', 0)
	for _, s := range stories {
		fmt.Fprintf(w, "%s\t%s\t%s\n", color.GreenString("ch"+strconv.FormatInt(s.Id, 10)), s.Name, s.AppUrl)
	}
	w.Flush()
	return nil
}
