package cmd

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var queryServerCommand = &cobra.Command{
	Use:   "queryserver",
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

func init() {
	Root.AddCommand(queryServerCommand)
}
