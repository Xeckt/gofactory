package api

type QueryServerStateData struct {
	ActiveSessionName   string  `json:"activeSessionName"`
	NumConnectedPlayers int     `json:"numConnectedPlayers"`
	PlayerLimit         int     `json:"playerLimit"`
	TechTier            int     `json:"techTier"`
	ActiveSchematic     string  `json:"activeSchematic"`
	GamePhase           string  `json:"gamePhase"`
	IsGameRunning       bool    `json:"isGameRunning"`
	TotalGameDuration   int     `json:"totalGameDuration"`
	IsGamePaused        bool    `json:"isGamePaused"`
	AverageTickRate     float64 `json:"averageTickRate"`
	AutoLoadSessionName string  `json:"autoLoadSessionName"`
}

type queryServerStateResponse struct {
	Data struct {
		State QueryServerStateData `json:"serverGameState"`
	} `json:"data"`
}

func (c *GoFactoryClient) QueryServerState() (*QueryServerStateData, error) {
	queryServerResponse, err := createAndSendPostRequest[queryServerStateResponse](c,
		QueryServerStateFunction,
		createGenericFunctionBody(QueryServerStateFunction))
	if err != nil {
		return nil, err
	}
	return &queryServerResponse.Data.State, nil
}
