package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"
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
		log.Trace().Msgf("query server stated command called with ctx %v and client %+v", ctx, client)
		state, err := client.QueryServerState(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("error during query server state command")
		}
		if state == nil {
			log.Fatal().Msgf("query server state command returned nil response: %+v", client)
		}

		message := fmt.Sprintf("-------"+
			"\nActive Session: %s\n"+
			"Connected Players: %d\n"+
			"Player Limit: %d\n"+
			"Tech Tier: %d\n"+
			"Active Schematic: %s\n"+
			"Phase: %s\n"+
			"Server running: %t\n"+
			"Running time: %02d:%02d:%02d\n"+
			"Paused: %t\n"+
			"Average Tick Rate: %f\n"+
			"Auto Load Session Name: %s\n"+
			"-------",
			state.ActiveSessionName,
			state.NumConnectedPlayers,
			state.PlayerLimit,
			state.TechTier,
			state.ActiveSchematic,
			state.GamePhase,
			state.IsGameRunning,
			state.TotalGameDuration/3600,
			(state.TotalGameDuration%3600)/60,
			state.TotalGameDuration%60,
			state.IsGamePaused,
			state.AverageTickRate,
			state.AutoLoadSessionName)

		for _, line := range strings.Split(message, "\n") {
			log.Info().Msgf("%s", line)
		}
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
		log.Trace().Msgf("get server options called with ctx %v", ctx)
		options, err := client.GetServerOptions(ctx)
		if err != nil {
			log.Fatal().Err(err)
		}
		if options == nil {
			log.Fatal().Msgf("get server options returned empty: %+v", client)
		}
		log.Trace().Msgf("received options: %+v", &options)

		appliedMessage := fmt.Sprintf("\n-------APPLIED OPTIONS------\n"+
			"Automatic Pause: %s\n"+
			"Auto Save on Disconnect: %s\n"+
			"Disable Seasonal Events: %s\n"+
			"Autosave Interval: %s\n"+
			"Server Restart Time Slot: %s\n"+
			"Send Gameplay Data: %s\n"+
			"Network Quality: %s\n"+
			"-------APPLIED OPTIONS------\n",
			options.ServerOptions.AutoPause, options.ServerOptions.AutoSaveOnDisconnect,
			options.ServerOptions.DisableSeasonalEvents, options.ServerOptions.AutosaveInterval,
			options.ServerOptions.ServerRestartTimeSlot, options.ServerOptions.SendGameplayData,
			options.ServerOptions.NetworkQuality)

		if !reflect.ValueOf(options.PendingServerOptions).IsZero() {
			// Going to require some reflection magic to do this dynamically for pretty printing
			// Will only print values that aren't zero so we dont spam the console with nil info.
			s := reflect.ValueOf(options.PendingServerOptions)
			for i := 0; i < s.NumField(); i++ {
				if !s.Field(i).IsZero() {
					appliedMessage += fmt.Sprintf(
						"\n-------PENDING OPTIONS------\n"+
							"%s %s"+
							"\n-------PENDING OPTIONS------\n",
						s.Type().Field(i).Name, s.Field(i).String())
				}
			}

			/*appliedMessage += fmt.Sprintf("\n-------PENDING OPTIONS------\n"+
			"Automatic Pause: %s\n"+
			"Auto Save on Disconnect: %s\n"+
			"Disable Seasonal Events: %s\n"+
			"Autosave Interval: %s\n"+
			"Server Restart Time Slot: %s\n"+
			"Send Gameplay Data: %s\n"+
			"Network Quality: %s\n"+
			"-------PENDING OPTIONS------\n",
			options.PendingServerOptions.AutoPause, options.PendingServerOptions.AutoSaveOnDisconnect,
			options.PendingServerOptions.DisableSeasonalEvents, options.PendingServerOptions.AutosaveInterval,
			options.PendingServerOptions.ServerRestartTimeSlot, options.PendingServerOptions.SendGameplayData,
			options.PendingServerOptions.NetworkQuality)*/
		}

		for _, line := range strings.Split(appliedMessage, "\n") {
			log.Info().Msgf("%s", line)
		}
	},
}

var setServerOptionsCommand = &cobra.Command{
	Use:   "set",
	Short: "set server options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Trace().Msgf("set server options called with ctx %v", ctx)
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
