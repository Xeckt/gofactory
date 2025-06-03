package api

import "context"

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
