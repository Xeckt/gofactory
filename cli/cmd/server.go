package cmd

import (
	"fmt"
	"reflect"

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
	},
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
			for i := 0; i < s.NumField(); i++ {
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
	},
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

func init() {
	Root.AddCommand(serverCommand)
	serverCommand.AddCommand(queryServerCommand)
	serverCommand.AddCommand(serverOptionsCommand)
	serverOptionsCommand.AddCommand(getServerOptionsCommand)
	serverOptionsCommand.AddCommand(setServerOptionsCommand)
}
