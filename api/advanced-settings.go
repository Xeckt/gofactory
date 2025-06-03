package api

import "context"

type AppliedAdvancedGameSettings struct {
	NoPower                         string `json:"FG.GameRules.NoPower,omitempty"`
	StartingTier                    string `json:"FG.GameRules.StartingTier,omitempty"`
	DisableArachnidCreatures        string `json:"FG.GameRules.DisableArachnidCreatures,omitempty"`
	NoUnlockCost                    string `json:"FG.GameRules.NoUnlockCost,omitempty"`
	SetGamePhase                    string `json:"FG.GameRules.SetGamePhase,omitempty"`
	UnlockAllResearchSchematics     string `json:"FG.GameRules.UnlockAllResearchSchematics,omitempty"`
	UnlockInstantAltRecipes         string `json:"FG.GameRules.UnlockInstantAltRecipes,omitempty"`
	UnlockAllResourceSinkSchematics string `json:"FG.GameRules.UnlockAllResourceSinkSchematics,omitempty"`
	NoBuildCost                     string `json:"FG.PlayerRules.NoBuildCost,omitempty"`
	GodMode                         string `json:"FG.PlayerRules.GodMode,omitempty"`
	FlightMode                      string `json:"FG.PlayerRules.FlightMode,omitempty"`
}

type AdvancedGameSettingsResponse struct {
	Data struct {
		Settings AppliedAdvancedGameSettings `json:"AppliedAdvancedGameSettings,omitempty"`
	} `json:"data,omitempty"`
}

func (c *GoFactoryClient) GetAdvancedGameSettings(ctx context.Context) (*AppliedAdvancedGameSettings, error) {
	appliedAdvanceSettingsResponse, err := CreateAndSendPostRequest[AdvancedGameSettingsResponse](ctx, c,
		GetAdvancedGameSettingsFunction,
		CreateGenericFunctionBody(GetAdvancedGameSettingsFunction))
	if err != nil {
		return nil, err
	}
	return &appliedAdvanceSettingsResponse.Data.Settings, nil
}
