package main

import (
	"context"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/rs/zerolog/log"
)

type Context struct {
	Client     *api.GoFactoryClient `kong:"-"`
	ApiContext context.Context
}

type CheckHealthCommand struct{}

func (h *CheckHealthCommand) Run(ctx *Context) error {
	log.Trace().Msgf("health check command called with ctx %v and client %+v", ctx, ctx.Client)
	health, err := ctx.Client.GetServerHealth(ctx.ApiContext, "")
	if err != nil {
		log.Fatal().Err(err).Msg("error during health check command")
	}
	if health == nil {
		log.Fatal().Msgf("health check command returned nil health response: %+v", ctx.Client)
	}
	if health.CustomData == "" || health.Health == "" {
		log.Error().Msgf("Health check looks empty, do you have a running session?")
		return nil
	}
	log.Info().Msgf("Health check returned: %+v", health.Health)
	return nil
}

type PasswordlessLoginCommand struct{}

type PasswordLoginCommand struct{}

type QueryServerStateCommand struct{}

func (q *QueryServerStateCommand) Run(ctx *Context) error {
	log.Trace().Msgf("query server stated command called with ctx %v and client %+v", ctx, ctx.Client)
	state, err := ctx.Client.QueryServerState(ctx.ApiContext)
	if err != nil {
		log.Fatal().Err(err).Msg("error during query server state command")
	}
	if state == nil {
		log.Fatal().Msgf("query server state command returned nil response: %+v", ctx.Client)
	}
	log.Info().Msgf("\nActive Session: %s\n"+
		"Connected Players: %d\n"+
		"Player Limit: %d\n"+
		"Tech Tier: %d\n"+
		"Active Schematic: %s\n"+
		"Phase: %s\n"+
		"Server running: %t\n"+
		"Running time: %02d:%02d:%02d\n"+
		"Paused: %t\n"+
		"Average Tick Rate: %f\n"+
		"Auto Load Session Name: %s\n",
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

	return nil
}

type GetServerOptionsCommand struct{}

type GetAdvancedGameSettingsCommand struct{}

type ApplyAdvancedGameSettingsCommand struct{}

type ClaimServerCommand struct{}

type RenameServerCommand struct{}

type SetClientPasswordCommand struct{}

type SetAdminPasswordCommand struct{}

type SetAutoLoadSessionNameCommand struct{}

type RunConsoleCommand struct{}

type ShutdownServerCommand struct{}

type ApplyServerOptionsCommand struct{}

type CreateNewGameCommand struct{}

type SaveGameCommand struct{}

type LoadGameCommand struct{}

type UploadSaveGameCommand struct{}

type DownloadSaveGameCommand struct{}

type DeleteSaveFileCommand struct{}

type DeleteSaveSessionCommand struct{}

type EnumerateSessionsCommand struct{}
