package main

import (
	"github.com/alchemicalkube/gofactory/cli/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		cmd.Logger.Trace(err.Error())
	}
}
