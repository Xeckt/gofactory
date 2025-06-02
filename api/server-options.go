package api

type GetServerOptionsResponse struct {
	Data GetServerOptionsData `json:"data"`
}

type GetServerOptionsData struct {
	ServerOptions        ServerOptions `json:"serverOptions"`
	PendingServerOptions ServerOptions `json:"pendingServerOptions,omitempty"`
}

type ServerOptions struct {
	AutoPause             string `json:"FG.DSAutoPause"`
	AutoSaveOnDisconnect  string `json:"FG.DSAutoSaveOnDisconnect"`
	DisableSeasonalEvents string `json:"FG.DisableSeasonalEvents"`
	AutosaveInterval      string `json:"FG.AutosaveInterval"`
	ServerRestartTimeSlot string `json:"FG.ServerRestartTimeSlo"`
	SendGameplayData      string `json:"FG.SendGameplayData"`
	NetworkQuality        string `json:"FG.NetworkQuality"`
}

func (c *GoFactoryClient) GetServerOptions() (*GetServerOptionsData, error) {

	optionsResponse, err := CreateAndSendPostRequest[GetServerOptionsResponse](c,
		GetServerOptionsFunction,
		createGenericFunctionBody(GetServerOptionsFunction))
	if err != nil {
		return nil, err
	}
	return &optionsResponse.Data, nil
}
