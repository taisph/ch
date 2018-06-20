/*
 * Copyright 2018 Tais P. Hansen
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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
		cli.StringFlag{Name: "team, t", Usage: "Name of team to create story in (requires project, overridden by project-id)"},
		cli.StringFlag{Name: "project, p", Usage: "Name of project in specified team to create story in (requires team, overridden by project-id)"},
		cli.Int64Flag{Name: "project-id", Usage: "Id of project to create story in"},
		cli.StringFlag{Name: "type, y", Usage: "Story type (feature, bug or chore)"},
		cli.StringSliceFlag{Name: "label, l", Usage: "Label to attach to the story (case-sensitive, can be specified multiple times)"},
		cli.StringSliceFlag{Name: "owner, w", Usage: "Owner of the story (can be specified multiple times)"},
		cli.StringSliceFlag{Name: "follower, f", Usage: "Follower of the story (can be specified multiple times)"},
		cli.BoolFlag{Name: "open, o", Usage: "Open story in browser after creating"},
	}}
)

func Create(c *cli.Context) error {
	var cs *v2.CreateStory
	if id := c.Int64("project-id"); id != 0 {
		cs = &v2.CreateStory{Name: c.Args().First(), ProjectId: id}
	} else if tn, pn := c.String("team"), c.String("project"); tn != "" && pn != "" {
		team, err := ch.Client.Teams().GetByName(tn)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		if team == nil {
			return cli.NewExitError(fmt.Sprintf("error: team not found: %s", tn), 1)
		}

		project, err := ch.Client.Projects().GetByName(pn, team.Id)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		cs = &v2.CreateStory{Name: c.Args().First(), ProjectId: project.Id}
	} else {
		return cli.NewExitError(fmt.Sprintf("error: you must specify either team and project names, or project id"), 1)
	}

	if t := c.String("type"); t != "" {
		cs.StoryType = v2.StoryType(t)
	}

	if lbl := c.StringSlice("label"); len(lbl) > 0 {
		for _, n := range lbl {
			cs.Labels = append(cs.Labels, v2.CreateLabelParams{Name: n})
		}
	}

	if o := c.StringSlice("owner"); len(o) > 0 {
		for _, n := range o {
			member, err := ch.Client.Members().GetByName(n)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			if member == nil {
				fmt.Fprintf(c.App.Writer, "warning: owner not found: %s", n)
				continue
			}

			cs.OwnerIds = append(cs.OwnerIds, member.Id)
		}
	}

	if f := c.StringSlice("follower"); len(f) > 0 {
		for _, n := range f {
			member, err := ch.Client.Members().GetByName(n)
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			if member == nil {
				fmt.Fprintf(c.App.Writer, "warning: follower not found: %s", n)
				continue
			}

			cs.FollowerIds = append(cs.FollowerIds, member.Id)
		}
	}

	story, err := ch.Client.Stories().Create(cs)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(story.AppUrl)
	if c.Bool("open") {
		browser.Open(story.AppUrl)
	}
	return nil
}
