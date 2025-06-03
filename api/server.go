package api

import (
	"context"
	"encoding/json"
)

type GetServerOptionsResponse struct {
	Data GetServerOptionsData `json:"data,omitempty"`
}

type GetServerOptionsData struct {
	ServerOptions        ServerOptions `json:"serverOptions,omitempty"`
	PendingServerOptions ServerOptions `json:"pendingServerOptions,omitempty"`
}

type ServerOptions struct {
	AutoPause             string `json:"FG.DSAutoPause,omitempty"`
	AutoSaveOnDisconnect  string `json:"FG.DSAutoSaveOnDisconnect,omitempty"`
	DisableSeasonalEvents string `json:"FG.DisableSeasonalEvents,omitempty"`
	AutosaveInterval      string `json:"FG.AutosaveInterval,omitempty"`
	ServerRestartTimeSlot string `json:"FG.ServerRestartTimeSlo,omitempty"`
	SendGameplayData      string `json:"FG.SendGameplayData,omitempty"`
	NetworkQuality        string `json:"FG.NetworkQuality,omitempty"`
}

func (c *GoFactoryClient) GetServerOptions(ctx context.Context) (*GetServerOptionsData, error) {

	optionsResponse, err := CreateAndSendPostRequest[GetServerOptionsResponse](ctx, c,
		GetServerOptionsFunction,
		CreateGenericFunctionBody(GetServerOptionsFunction))
	if err != nil {
		return nil, err
	}
	return &optionsResponse.Data, nil
}

type QueryServerStateData struct {
	ActiveSessionName   string  `json:"activeSessionName,omitempty"`
	NumConnectedPlayers int     `json:"numConnectedPlayers,omitempty"`
	PlayerLimit         int     `json:"playerLimit,omitempty"`
	TechTier            int     `json:"techTier,omitempty"`
	ActiveSchematic     string  `json:"activeSchematic,omitempty"`
	GamePhase           string  `json:"gamePhase,omitempty"`
	IsGameRunning       bool    `json:"isGameRunning,omitempty"`
	TotalGameDuration   int     `json:"totalGameDuration,omitempty"`
	IsGamePaused        bool    `json:"isGamePaused,omitempty"`
	AverageTickRate     float64 `json:"averageTickRate,omitempty"`
	AutoLoadSessionName string  `json:"autoLoadSessionName,omitempty"`
}

type QueryServerStateResponse struct {
	Data struct {
		State QueryServerStateData `json:"serverGameState,omitempty"`
	} `json:"data"`
}

func (c *GoFactoryClient) QueryServerState(ctx context.Context) (*QueryServerStateData, error) {
	queryServerResponse, err := CreateAndSendPostRequest[QueryServerStateResponse](ctx, c,
		QueryServerStateFunction,
		CreateGenericFunctionBody(QueryServerStateFunction))
	if err != nil {
		return nil, err
	}
	return &queryServerResponse.Data.State, nil
}

type SetAutoLoadSessionRequest struct {
	Function string                        `json:"function"`
	Data     SetAutoLoadSessionRequestData `json:"data,omitempty"`
}

type SetAutoLoadSessionRequestData struct {
	SessionName string `json:"sessionName"`
}

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
