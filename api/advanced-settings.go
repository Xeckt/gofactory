package api

type AppliedAdvancedGameSettings struct {
	NoPower                         string `json:"FG.GameRules.NoPower"`
	StartingTier                    string `json:"FG.GameRules.StartingTier"`
	DisableArachnidCreatures        string `json:"FG.GameRules.DisableArachnidCreatures"`
	NoUnlockCost                    string `json:"FG.GameRules.NoUnlockCost"`
	SetGamePhase                    string `json:"FG.GameRules.SetGamePhase"`
	UnlockAllResearchSchematics     string `json:"FG.GameRules.UnlockAllResearchSchematics"`
	UnlockInstantAltRecipes         string `json:"FG.GameRules.UnlockInstantAltRecipes"`
	UnlockAllResourceSinkSchematics string `json:"FG.GameRules.UnlockAllResourceSinkSchematics"`
	NoBuildCost                     string `json:"FG.PlayerRules.NoBuildCost"`
	GodMode                         string `json:"FG.PlayerRules.GodMode"`
	FlightMode                      string `json:"FG.PlayerRules.FlightMode"`
}

type advancedGameSettingsResponse struct {
	Data struct {
		Settings AppliedAdvancedGameSettings `json:"AppliedAdvancedGameSettings"`
	} `json:"data"`
}

func (c *GoFactoryClient) GetAdvancedGameSettings() (*AppliedAdvancedGameSettings, error) {
	appliedAdvanceSettingsResponse, err := createAndSendPostRequest[advancedGameSettingsResponse](c,
		GetAdvancedGameSettingsFunction,
		createGenericFunctionBody(GetAdvancedGameSettingsFunction))
	if err != nil {
		return nil, err
	}
	return &appliedAdvanceSettingsResponse.Data.Settings, nil
}
