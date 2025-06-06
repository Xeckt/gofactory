package cmd

import (
	"context"
	"os"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	consoleWriter.TimeFormat = "15:04:05"

	multi := zerolog.MultiLevelWriter(consoleWriter)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger().Level(zerolog.InfoLevel)

	Root.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if Trace {
			log.Logger = log.Logger.Level(zerolog.TraceLevel)
			log.Trace().Msg("tracing enabled")
		}
	}
}
