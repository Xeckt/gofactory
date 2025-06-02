package api

type SendFunctionRequest struct {
	Function string `json:"function"`
}

type HealthCheckCustomData struct {
	CustomData string `json:"clientCustomData"`
}

type HealthCheckRequest struct {
	Function string                `json:"function"`
	Data     HealthCheckCustomData `json:"data"`
}

type HealthCheckResponse struct {
	Data struct {
		Health     string `json:"health"`
		CustomData string `json:"serverCustomData"`
	} `json:"data"`
}

type GetServerOptionsResponse struct {
	Data struct {
		ServerOptions        ServerOptions `json:"serverOptions"`
		PendingServerOptions ServerOptions `json:"pendingServerOptions"`
	} `json:"data"`
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
