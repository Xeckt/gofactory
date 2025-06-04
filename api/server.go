package api

import (
	"context"
	"encoding/json"
)

// GetServerOptionsResponse represents the response from the Satisfactory dedicated server
// when querying the current and pending server options.
type GetServerOptionsResponse struct {
	// Data contains the current and pending server options.
	Data GetServerOptionsData `json:"data,omitempty"`
}

// GetServerOptionsData holds the current and pending server options.
type GetServerOptionsData struct {
	// ServerOptions represents the current server options.
	ServerOptions ServerOptions `json:"serverOptions,omitempty"`

	// PendingServerOptions represents server options that are pending to be applied.
	PendingServerOptions ServerOptions `json:"pendingServerOptions,omitempty"`
}

// ServerOptions defines various configuration options for the Satisfactory server.
type ServerOptions struct {
	// AutoPause controls whether the server will pause automatically when no-one is connected.
	AutoPause string `json:"FG.DSAutoPause,omitempty"`

	// AutoSaveOnDisconnect determines whether to automatically save the game when a player disconnects.
	AutoSaveOnDisconnect string `json:"FG.DSAutoSaveOnDisconnect,omitempty"`

	// DisableSeasonalEvents disables any seasonal events in the game.
	DisableSeasonalEvents string `json:"FG.DisableSeasonalEvents,omitempty"`

	// AutosaveInterval is the interval at which the game autosaves.
	AutosaveInterval string `json:"FG.AutosaveInterval,omitempty"`

	// ServerRestartTimeSlot defines the time slot for scheduled server restarts.
	ServerRestartTimeSlot string `json:"FG.ServerRestartTimeSlot,omitempty"`

	// SendGameplayData determines whether to send gameplay data to Ficsit.
	SendGameplayData string `json:"FG.SendGameplayData,omitempty"`

	// NetworkQuality sets the network quality mode for the server.
	NetworkQuality string `json:"FG.NetworkQuality,omitempty"`
}

// ApplyServerOptionsRequest represents a request to apply new server options.
type ApplyServerOptionsRequest struct {
	// Function specifies the API function to call.
	Function string `json:"function"`

	// Data contains the new server options to apply.
	Data ApplyServerOptionsRequestData `json:"data,omitempty"`
}

// ApplyServerOptionsRequestData holds the new server options to apply.
type ApplyServerOptionsRequestData struct {
	// ServerOptions represents the updated server options.
	ServerOptions ServerOptions `json:"updatedServerOptions,omitempty"`
}

// ApplyServerOptions applies new server configuration options to the Satisfactory dedicated server.
func (c *GoFactoryClient) ApplyServerOptions(ctx context.Context, options ServerOptions) error {
	functionBody, err := json.Marshal(ApplyServerOptionsRequest{
		Function: ApplyAdvancedGameSettingsFunction,
		Data:     ApplyServerOptionsRequestData{ServerOptions: options},
	})
	if err != nil {
		return err
	}

	request, err := c.CreatePostRequest(ApplyServerOptionsFunction, functionBody)
	if err != nil {
		return err
	}

	err = c.SendPostRequest(ctx, request, functionBody)
	if err != nil {
		return err
	}

	return nil
}

// GetServerOptions retrieves the current and pending server options
// from the Satisfactory dedicated server.
func (c *GoFactoryClient) GetServerOptions(ctx context.Context) (*GetServerOptionsData, error) {

	optionsResponse, err := CreateAndSendPostRequest[GetServerOptionsResponse](ctx, c,
		GetServerOptionsFunction,
		CreateGenericFunctionBody(GetServerOptionsFunction))
	if err != nil {
		return nil, err
	}
	return &optionsResponse.Data, nil
}

// QueryServerStateData represents the current state of the Satisfactory server.
type QueryServerStateData struct {
	// ActiveSessionName is the name of the currently loaded game session.
	ActiveSessionName string `json:"activeSessionName,omitempty"`

	// NumConnectedPlayers is the number of connected players.
	NumConnectedPlayers int `json:"numConnectedPlayers,omitempty"`

	// PlayerLimit is the maximum number of players that can be connected.
	PlayerLimit int `json:"playerLimit,omitempty"`

	// TechTier is the maximum tech tier of all Schematics currently unlocked
	TechTier int `json:"techTier,omitempty"`

	// ActiveSchematic is the schematic currently set as the active milestone.
	ActiveSchematic string `json:"activeSchematic,omitempty"`

	// GamePhase is the current game phase. None is no game is running.
	GamePhase string `json:"gamePhase,omitempty"`

	// IsGameRunning indicates whether a save is loaded, or if it's waiting for a session to be created.
	IsGameRunning bool `json:"isGameRunning,omitempty"`

	// TotalGameDuration is the total time the current save has been loaded in seconds.
	TotalGameDuration int `json:"totalGameDuration,omitempty"`

	// IsGamePaused indicates whether the game is currently paused.
	IsGamePaused bool `json:"isGamePaused,omitempty"`

	// AverageTickRate is the average server tick rate.
	AverageTickRate float64 `json:"averageTickRate,omitempty"`

	// AutoLoadSessionName is the name of the session set to auto-load.
	AutoLoadSessionName string `json:"autoLoadSessionName,omitempty"`
}

// QueryServerStateResponse represents the response from the server when querying its state.
type QueryServerStateResponse struct {
	// Data contains the current server game state.
	Data struct {
		State QueryServerStateData `json:"serverGameState,omitempty"`
	} `json:"data"`
}

// QueryServerState queries the current state of the Satisfactory server and returning all state data.
func (c *GoFactoryClient) QueryServerState(ctx context.Context) (*QueryServerStateData, error) {
	queryServerResponse, err := CreateAndSendPostRequest[QueryServerStateResponse](ctx, c,
		QueryServerStateFunction,
		CreateGenericFunctionBody(QueryServerStateFunction))
	if err != nil {
		return nil, err
	}
	return &queryServerResponse.Data.State, nil
}

// SetAutoLoadSessionRequest represents a request to set the automatically loaded session during
// server startup.
type SetAutoLoadSessionRequest struct {
	// Function specifies the API function to call.
	Function string `json:"function"`

	// Data contains the new auto-load session name.
	Data SetAutoLoadSessionRequestData `json:"data"`
}

// SetAutoLoadSessionRequestData holds the new auto-load session name.
type SetAutoLoadSessionRequestData struct {
	// SessionName is the name of the session to set for auto-load.
	SessionName string `json:"sessionName"`
}

// SetAutoLoadSessionName sets the auto-load session name on the Satisfactory server.
func (c *GoFactoryClient) SetAutoLoadSessionName(ctx context.Context, sessionName string) (bool, error) {
	functionBody, err := json.Marshal(SetAutoLoadSessionRequest{
		Function: SetAutoLoadSessionNameFunction,
		Data: SetAutoLoadSessionRequestData{
			SessionName: sessionName,
		},
	})
	if err != nil {
		return false, err
	}
	request, err := c.CreatePostRequest(SetAutoLoadSessionNameFunction, functionBody)
	if err != nil {
		return false, err
	}

	err = c.SendPostRequest(ctx, request, functionBody)
	if err != nil {
		return false, err
	}

	return true, nil
}

// RunCommandRequest represents a request to execute a console command on the server.
type RunCommandRequest struct {
	// Function specifies the API function to call for running a command.
	Function string `json:"function"`

	// Data contains the command to execute.
	Data RunCommandRequestData `json:"data,omitempty"`
}

// RunCommandRequestData holds the console command to run on the server.
type RunCommandRequestData struct {
	// Command is the console command to execute.
	Command string `json:"command"`
}

// RunServerCommand sends a console command to the Satisfactory dedicated server for execution.
func (c *GoFactoryClient) RunServerCommand(ctx context.Context, command string) error {
	functionBody, err := json.Marshal(RunCommandRequest{
		Function: RunCommandFunction,
		Data: RunCommandRequestData{
			Command: command,
		},
	})
	if err != nil {
		return err
	}
	request, err := c.CreatePostRequest(RunCommandFunction, functionBody)
	if err != nil {
		return err
	}

	err = c.SendPostRequest(ctx, request, functionBody)
	if err != nil {
		return err
	}

	return nil
}

// ShutdownServer sends a request to shut down the Satisfactory dedicated server.
func (c *GoFactoryClient) ShutdownServer(ctx context.Context) error {
	functionBody := CreateGenericFunctionBody(ShutdownFunction)

	request, err := c.CreatePostRequest(ShutdownFunction, functionBody)
	if err != nil {
		return err
	}

	err = c.SendPostRequest(ctx, request, functionBody)
	if err != nil {
		return err
	}

	return nil
}
