package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/rs/zerolog/log"
)

type Context struct {
	Client     *api.GoFactoryClient `kong:"-"`
	ApiContext context.Context
	Trace      bool
}

type CheckHealthCommand struct{}

func (h *CheckHealthCommand) Run(ctx *Context) error {
	log.Trace().Msgf("health check command called with ctx %v and client %+v", ctx, ctx.Client)
	health, err := ctx.Client.GetServerHealth(ctx.ApiContext, "gofactory-cli-healthcheck-call")
	if err != nil {
		log.Fatal().Err(err).Msg("error during health check command")
	}
	if health == nil {
		log.Fatal().Msgf("health check command returned nil health response: %+v", ctx.Client)
	}
	if len(health.CustomData) > 0 {
		log.Info().Msgf("Server health: %+v\nCustom data reported: %s", health.Health, health.CustomData)
	}
	log.Info().Msgf("Server health: %+v", health.Health)
	return nil
}

type PasswordlessLoginCommand struct {
	Privilege string `arg:"" required:"" help:"The privilege to attach to the returned token"`
}

func (p *PasswordlessLoginCommand) Run(ctx *Context) error {
	// Make sure we have the correct casing for the privilege level.
	switch strings.ToLower(p.Privilege) {
	case "notauthenticated":
		p.Privilege = "NotAuthenticated"
	case "client":
		p.Privilege = "Client"
	case "administrator":
		p.Privilege = "Administrator"
	case "initialadmin:":
		p.Privilege = "InitialAdmin"
	case "apitoken":
		p.Privilege = "ApiToken"
	default:
		return fmt.Errorf("invalid privilege, use one of: %s | %s | %s | %s | %s",
			api.NOT_AUTHENTICATED_PRIVILEGE, api.CLIENT_PRIVILEGE, api.ADMINISTRATOR_PRIVILEGE,
			api.INITIAL_ADMIN_PRIVILEGE, api.API_TOKEN_PRIVILEGE)
	}
	err := ctx.Client.PasswordlessLogin(ctx.ApiContext, p.Privilege)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msgf("Successfully authenticated with privilege: %s", p.Privilege)
	log.Info().Msgf("Token returned: %s", ctx.Client.Token)
	log.Warn().Msgf("If you wish to use this token, make sure to replace your %s environment variable!", ENV_TOKEN)
	return nil
}

func (p *PasswordlessLoginCommand) Help() string {
	return fmt.Sprintf("Possible values for <privilege> are:"+
		"\n\t%s\t%s\n\t%s\t%s\n\t%s",
		api.NOT_AUTHENTICATED_PRIVILEGE, api.CLIENT_PRIVILEGE, api.ADMINISTRATOR_PRIVILEGE,
		api.INITIAL_ADMIN_PRIVILEGE, api.API_TOKEN_PRIVILEGE)
}

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
