package api

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
