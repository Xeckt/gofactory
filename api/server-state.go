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

func (c *GoFactoryClient) QueryServerState() (*QueryServerStateData, *APIError, error) {
	request, err := c.createPostRequest(QueryServerStateFunction, createGenericFunctionBody(QueryServerStateFunction))
	if err != nil {
		return nil, nil, err
	}

	var query queryServerStateResponse
	apiErr, err := c.sendPostRequest(request, &query)
	if err != nil {
		return nil, nil, err
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return &query.Data.State, nil, nil
}
