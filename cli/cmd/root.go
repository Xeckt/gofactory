package cmd

import (
	"context"
	"os"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	Root = &cobra.Command{
		Use:   "gofactory",
		Short: "cli tool for interacting with a satisfactory dedicated server",
	}

	Trace bool
)

var (
	serverUrl   string
	serverToken string
	client      *api.GoFactoryClient
	ctx         context.Context
	Logger      *pterm.Logger
)

const VERSION = "0.0.1"
const ENV_URL = "GF_URL"
const ENV_TOKEN = "GF_TOKEN"

func init() {
	serverUrl = os.Getenv(ENV_URL)
	serverToken = os.Getenv(ENV_TOKEN)

	if len(serverUrl) == 0 || len(serverToken) == 0 {
		log.Fatal().Msgf("One of the required environment variables are not set:\n\tURL: %s\n\tTOKEN: %s",
			serverUrl, serverToken)
	}

	client = api.NewGoFactoryClient(serverUrl, serverToken, true)
	ctx = context.Background()

	Root.PersistentFlags().BoolVarP(&Trace, "trace", "t", false, "set the cli to trace mode")

	Logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)

	Root.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if Trace {
			Logger.Level = pterm.LogLevelTrace
			Logger = Logger.WithCaller()
		}
	}
}
