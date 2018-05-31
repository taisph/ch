package main

import (
	"fmt"

	_ "github.com/taisph/ch/internal/app/cmd/ch"
	_ "github.com/taisph/ch/internal/app/cmd/project"
	_ "github.com/taisph/ch/internal/app/cmd/story"
	"github.com/taisph/ch/pkg/cmdr"
)

func main() {
	err := cmdr.CmdRunner.Run()
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
