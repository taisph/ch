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
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/taisph/ch/internal/app/cmd/ch"
	"github.com/urfave/cli"
)

var (
	teamListCmd = cli.Command{Name: "list", Action: List, Usage: "List teams", Flags: []cli.Flag{}}
)

func List(c *cli.Context) error {
	teams, err := ch.Client.Teams().List()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	w := tabwriter.NewWriter(c.App.Writer, 0, 0, 2, ' ', 0)
	for _, s := range teams {
		fmt.Fprintf(w, "%s\t%s\t%s\n", color.GreenString(strconv.FormatInt(s.Id, 10)), s.Name, *s.Description)
	}
	w.Flush()
	return nil
}
