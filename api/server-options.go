package api

type getServerOptionsResponse struct {
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

// Pointer generics without T any interfaces would be cool right about now....
func (c *GoFactoryClient) GetServerOptions() (*GetServerOptionsData, *APIError, error) {
	request, err := c.createPostRequest(GetServerOptionsFunction, createGenericFunctionBody(GetServerOptionsFunction))
	if err != nil {
		return nil, nil, err
	}

	var options getServerOptionsResponse
	apiErr, err := c.sendPostRequest(request, &options)
	if err != nil {
		return nil, nil, err
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}

	return &options.Data, nil, nil
}
