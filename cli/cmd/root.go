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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if Trace {
				Logger.Warn("TRACING WILL DISPLAY SENSITIVE INFORMATION!")
				Logger.Level = pterm.LogLevelTrace
				Logger = Logger.WithCaller()
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) == 1 {
				StartUi()
			}
		},
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
const ENV_GF_URL = "GF_URL"
const ENV_GF_TOKEN = "GF_TOKEN"

func init() {
	Logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)

	serverUrl = os.Getenv(ENV_GF_URL)
	serverToken = os.Getenv(ENV_GF_TOKEN)

	if len(serverUrl) == 0 || len(serverToken) == 0 {
		Logger.Fatal("check for empty environment variables", Logger.Args(
			fmt.Sprintf("%s", ENV_GF_URL), serverUrl,
			fmt.Sprintf("%s", ENV_GF_TOKEN), serverToken))
	}

	client = api.NewGoFactoryClient(serverUrl, serverToken, true)
	ctx = context.Background()

	Root.PersistentFlags().BoolVarP(&Trace, "trace", "t", false, "set the cli to trace mode")
}

func StartUi() {
	selected, err := selectMenu.Show()
	if err != nil {
		Logger.Fatal(err.Error())
	}

	switch selected {
	case "query server":
		queryServerCommand.Run(queryServerCommand, []string{})
	case "health check":
		healthCheckCmd.Run(healthCheckCmd, []string{})
	case "login":
		selected, err := loginMenu.Show()
		if err != nil {
			Logger.Fatal(err.Error())
		}
		switch selected {
		case "password":
			pInput := pterm.DefaultInteractiveTextInput.WithMask("*")
			password, err := pInput.Show("Enter password to authenticate")
			if err != nil {
				Logger.Fatal(err.Error())
			}
			privilege, err := privilegeMenu.Show()
			if err != nil {
				Logger.Fatal(err.Error())
			}
			fmt.Println(password, privilege)
		}
	default:
		Logger.Error("Unknown option")
	}
	StartUi()
}
