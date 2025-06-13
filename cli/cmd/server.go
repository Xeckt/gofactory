package cmd

import (
	"fmt"
	"reflect"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "command to handle server specific commands",
	Args:  cobra.ExactArgs(1),
}

var queryServerCommand = &cobra.Command{
	Use:   "query",
	Short: "Receive information about the current server state.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		queryServer()
	},
}

func queryServer() {
	state, err := client.QueryServerState(ctx)
	Logger.Trace("query server command",
		Logger.Args(
			"context", ctx,
			"client pointer", &client,
			"client object", client,
			"state pointer", &state,
			"state object", state),
	)
	if err != nil {
		Logger.Fatal("query server error", Logger.Args("error", err))
	}
	if state == nil {
		Logger.Fatal("query server state returned nil response")
	}

	runningTime := fmt.Sprintf("%02d:%02d:%02d",
		state.TotalGameDuration/3600,
		(state.TotalGameDuration%3600)/60,
		state.TotalGameDuration%60)

	Logger.Info("server state", Logger.Args(
		"active session", state.ActiveSessionName,
		"connected players", state.NumConnectedPlayers,
		"player limit", state.PlayerLimit,
		"tech tier", state.TechTier,
		"phase", state.GamePhase,
		"server running", state.IsGameRunning,
		"running time", runningTime,
		"paused", state.IsGamePaused,
		"average tick rate", state.AverageTickRate,
		"auto load session name", state.AutoLoadSessionName))
}

var serverOptionsCommand = &cobra.Command{
	Use:   "options",
	Short: "command to handle server options",
	Args:  cobra.ExactArgs(1),
}

var getServerOptionsCommand = &cobra.Command{
	Use:   "get",
	Short: "get server options",
	Run: func(cmd *cobra.Command, args []string) {
		getServerOptions()
	},
}

func getServerOptions() {
	Logger.Trace("get server options", Logger.Args(
		"context", ctx,
		"client pointer", &client,
		"client object", client,
	))

	options, err := client.GetServerOptions(ctx)
	if err != nil {
		Logger.Fatal("get server options error", Logger.Args("error", err))
	}

	if options == nil {
		Logger.Fatal("get server options returned nil response")
	}

	Logger.Trace("server options response", Logger.Args(
		"options pointer", &options,
		"options object", options,
	))

	Logger.Info("applied server options", Logger.Args(
		"automatic pause", options.ServerOptions.AutoPause,
		"auto save on disconnect", options.ServerOptions.AutoSaveOnDisconnect,
		"disable seasonal events", options.ServerOptions.DisableSeasonalEvents,
		"autosave interval", options.ServerOptions.AutosaveInterval,
		"server restart time", options.ServerOptions.ServerRestartTimeSlot,
		"send gameplay data", options.ServerOptions.SendGameplayData,
		"network quality", options.ServerOptions.NetworkQuality))

	if !reflect.ValueOf(options.PendingServerOptions).IsZero() {
		// I f*cking love pterm.

		m := make(map[string]any)
		pendingOptionsStyle := make(map[string]pterm.Style)

		s := reflect.ValueOf(options.PendingServerOptions)
		for i := range s.NumField() {
			if !s.Field(i).IsZero() {
				pendingOptionsStyle[s.Type().Field(i).Name] = *pterm.NewStyle(pterm.FgYellow)
				m[s.Type().Field(i).Name] = s.Field(i).String()
			}
		}

		if len(m) > 0 {
			Logger.AppendKeyStyles(pendingOptionsStyle)
			Logger.Info("pending server options", Logger.ArgsFromMap(m))
		}
	}
}

var setServerOptionsCommand = &cobra.Command{
	Use:   "set",
	Short: "set server options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		/*
				Since there are a lot of options here, I'll need to think about how to actually apply them.
				By standard, we will allow multiple flags to specify the server options as sometimes you may only want
				to set one or two options. But if you want to do loads, might be worth specifying a file?

			TODO: Think carefully about this one.

		*/
	},
}

var serverNameFlag string

var renameServerCommand = &cobra.Command{
	Use:   "rename",
	Short: "rename server",
	Run: func(cmd *cobra.Command, args []string) {
		renameServer(serverNameFlag)
	},
}

func renameServer(name string) {
	if len(serverNameFlag) == 0 {
		Logger.Fatal("you must specify --name")
	}
	err := client.RenameServer(ctx, name)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("server renamed", Logger.Args("new name", name))
}

var claimServerCommand = &cobra.Command{
	Use:   "claim",
	Short: "claim server",
	Run: func(cmd *cobra.Command, args []string) {
		claimServer(serverNameFlag, passwordFlag)
	},
}

func claimServer(serverName string, password string) {
	if len(passwordFlag) == 0 || len(serverNameFlag) == 0 {
		Logger.Fatal("you must specify --password and --name")
	}

	if len(client.Token) != 0 {
		Logger.Fatal("your GF_TOKEN environment variable is not empty")
	}

	claimData := api.ClaimRequestData{
		ServerName:    serverName,
		AdminPassword: password,
	}

	err := client.ClaimServer(ctx, claimData)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("server claimed", Logger.Args("server name:", serverName, "password", password, "token", client.Token))
	Logger.Warn("make sure to update your GF_TOKEN environment variable with the new one!")
}

var setPasswordCommand = &cobra.Command{
	Use:   "set-password",
	Short: "allows you to set admin or client password",
	Args:  cobra.ExactArgs(1),
}

var setClientPasswordCommand = &cobra.Command{
	Use:   "client",
	Short: "command to set client password",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			Logger.Fatal("no password specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setClientPassword(args[0])
	},
}

func setClientPassword(password string) {
	err := client.SetClientPassword(ctx, password)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("client password set")
}

var setAdminPasswordCommand = &cobra.Command{
	Use:   "admin",
	Short: "command to set admin password",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			Logger.Fatal("no password specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setAdminPassword(args[0])
	},
}

func setAdminPassword(password string) {
	err := client.SetAdminPassword(ctx, password)
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("admin password set")
}

var runCommand = &cobra.Command{
	Use:   "run-command",
	Short: "command to either shutdown the server or run a command on the server",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			Logger.Fatal("please specify a command to run")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		commandRunner(args[0])
	},
}

func commandRunner(command string) {
	if command == "shutdown" {
		err := client.ShutdownServer(ctx)
		if err != nil {
			Logger.Fatal(err.Error())
		}
	} else {
		err := client.RunServerCommand(ctx, command)
		if err != nil {
			Logger.Fatal(err.Error())
		}
	}
	Logger.Info("successful", Logger.Args("command", command))
}

func init() {
	Root.AddCommand(serverCommand)

	serverCommand.AddCommand(claimServerCommand)
	serverCommand.AddCommand(queryServerCommand)
	serverCommand.AddCommand(serverOptionsCommand)
	serverCommand.AddCommand(renameServerCommand)
	serverCommand.AddCommand(setPasswordCommand)
	serverCommand.AddCommand(runCommand)

	serverCommand.PersistentFlags().StringVarP(&passwordFlag, "password", "p", "", "flag to supply a password to required commands")
	serverCommand.PersistentFlags().StringVarP(&serverNameFlag, "name", "n", "", "flag to supply a server name to required commands")

	serverOptionsCommand.AddCommand(getServerOptionsCommand)
	serverOptionsCommand.AddCommand(setServerOptionsCommand)

	setPasswordCommand.AddCommand(setAdminPasswordCommand)
	setPasswordCommand.AddCommand(setClientPasswordCommand)
}
