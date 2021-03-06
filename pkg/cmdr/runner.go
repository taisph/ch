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
