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

package team

import (
	"fmt"

	"github.com/taisph/ch/pkg/clubhouse/v2"

	"github.com/taisph/ch/internal/app/cmd/ch"
	"github.com/urfave/cli"
)

var (
	teamShowCmd = cli.Command{Name: "show", Action: Show, Usage: "Show team", Flags: []cli.Flag{
		cli.StringFlag{Name: "team", Usage: "Name of team to show"},
		cli.Int64Flag{Name: "team-id", Usage: "Id of team to show (overrides team)"},
	}}
)

func Show(c *cli.Context) error {
	var team *v2.Team
	var err error
	if id := c.Int64("team-id"); id != 0 {
		team, err = ch.Client.Teams().Get(id)
	} else if name := c.String("team"); name != "" {
		team, err = ch.Client.Teams().GetByName(name)
	} else {
		err = fmt.Errorf("error: you must specify either team or team-id")
	}
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Printf("%s\n%s\n%d projects, %d workflow states\n", team.Name, *team.Description, len(team.ProjectIds), len(team.Workflow.States))

	return nil
}
