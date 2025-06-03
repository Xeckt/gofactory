package api

import "context"

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
