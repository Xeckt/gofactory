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
		SilenceUsage: true,
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

const (
	VERSION      = "0.0.1"
	ENV_GF_URL   = "GF_URL"
	ENV_GF_TOKEN = "GF_TOKEN"
)

func init() {
	Logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo)

	serverUrl = os.Getenv(ENV_GF_URL)
	serverToken = os.Getenv(ENV_GF_TOKEN)

	if len(serverToken) == 0 {
		Logger.Warn("check for empty environment variables", Logger.Args(serverUrl, serverToken))
	} else if len(serverUrl) == 0 {
		Logger.Fatal("GF_URL cannot be empty!")
	}

	client = api.NewGoFactoryClient(serverUrl, "", true)
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
		queryServer()
	case "health check":
		healthCheck()
	case "login":
		selected, err := loginMenu.Show()
		if err != nil {
			Logger.Fatal(err.Error())
		}

		privilege, err := privilegeMenu.Show()
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
			passwordLogin(password, privilege)
		case "passwordless":
			passwordlessLogin(privilege)
		}
	default:
		Logger.Error("Unknown option")
	}
}
