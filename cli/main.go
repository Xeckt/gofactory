package main

import (
	"github.com/alchemicalkube/gofactory/cli/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		log.Fatal().Err(err)
	}
}
