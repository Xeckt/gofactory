package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var healthCheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Run basic health check against the HTTPS api",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Trace().Msgf("health check command called with ctx %v and client %+v", ctx, client)
		health, err := client.GetServerHealth(ctx, "gofactory-cli-healthcheck-call")
		if err != nil {
			log.Fatal().Err(err).Msg("error during health check command")
		}
		if health == nil {
			log.Fatal().Msgf("health check command returned nil health response: %+v", ctx)
		}
		if len(health.CustomData) > 0 {
			log.Info().Msgf("Server health: %+v\nCustom data reported: %s", health.Health, health.CustomData)
		}
		log.Info().Msgf("Server health: %+v", health.Health)
	},
}

func init() {
	Root.AddCommand(healthCheckCmd)
}
