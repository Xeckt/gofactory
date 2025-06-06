package cmd

import (
	"context"
	"fmt"
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
	Logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)

	serverUrl = os.Getenv(ENV_URL)
	serverToken = os.Getenv(ENV_TOKEN)

	if len(serverUrl) == 0 || len(serverToken) == 0 {
		Logger.Fatal("One of the required environment variables are not set", Logger.Args(
			fmt.Sprintf("%s", ENV_URL), serverUrl,
			fmt.Sprintf("%s", ENV_TOKEN), serverToken))
	}

	client = api.NewGoFactoryClient(serverUrl, serverToken, true)
	ctx = context.Background()

	Root.PersistentFlags().BoolVarP(&Trace, "trace", "t", false, "set the cli to trace mode")

	Root.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if Trace {
			Logger.Level = pterm.LogLevelTrace
			Logger = Logger.WithCaller()
		}
	}
}
