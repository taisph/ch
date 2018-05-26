package story

import (
	"fmt"

	"github.com/taisph/ch/internal/app/cmd/ch"
	"github.com/taisph/ch/internal/pkg/browser"
	"github.com/taisph/ch/pkg/clubhouse/v2"
	"github.com/urfave/cli"
)

var (
	storyCreateCmd = cli.Command{Name: "create", Action: Create, Usage: "Create a new story", ArgsUsage: "<name>", Flags: []cli.Flag{
		cli.Int64Flag{Name: "project-id, p", Usage: "Id of project to create story in"},
		cli.BoolFlag{Name: "open, o", Usage: "Open story in browser after creating"},
	}}
)

func Create(c *cli.Context) error {
	story, err := ch.Client.Stories().Create(&v2.CreateStory{Name: c.Args().First(), ProjectId: c.Int64("project-id")})
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(story.AppUrl)
	if c.Bool("open") {
		browser.Open(story.AppUrl)
	}
	return nil
}
