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
	app.Version = "0.0.2"
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
