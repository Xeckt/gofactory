package cmd

import (
	"github.com/spf13/cobra"
)

var healthCheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Run basic health check against the HTTPS api",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		healthCheck()
	},
}

func healthCheck() {
	Logger.Trace("health check command", Logger.Args(
		"Context", ctx,
		"Client pointer", &client),
	)
	server, err := client.GetServerHealth(ctx, "gofactory-cli-healthcheck-call")
	if err != nil {
		Logger.Fatal("healthcheck error", Logger.Args("error", err, "context", ctx, "object", server))
	}
	if server == nil {
		Logger.Fatal("healthcheck command returned nil response")
	}
	if len(server.CustomData) > 0 {
		Logger.Info("health check returned custom data", Logger.Args("Health", server,
			"Custom Data", server.CustomData))
	}
	Logger.Info("health check returned successfully", Logger.Args("Health", server.Health))
}

func init() {
	Root.AddCommand(healthCheckCmd)
}
