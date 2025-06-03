package api

import (
	"context"
	"encoding/json"
)

type AdvancedGameSettings struct {
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

type AdvancedGameSettingsData struct {
	Data AdvancedGameSettingsResponse `json:"data,omitempty"`
}

type AdvancedGameSettingsResponse struct {
	Settings AdvancedGameSettings `json:"appliedAdvancedGameSettings,omitempty"`
}
type ApplyAdvancedGameSettingsRequest struct {
	Function string               `json:"function"`
	Data     AdvancedGameSettings `json:"appliedAdvancedGameSettings"`
}

func (c *GoFactoryClient) GetAdvancedGameSettings(ctx context.Context) (*AdvancedGameSettings, error) {
	appliedAdvanceSettingsResponse, err := CreateAndSendPostRequest[AdvancedGameSettingsResponse](ctx, c,
		GetAdvancedGameSettingsFunction,
		CreateGenericFunctionBody(GetAdvancedGameSettingsFunction))
	if err != nil {
		return nil, err
	}
	return &appliedAdvanceSettingsResponse.Settings, nil
}

func (c *GoFactoryClient) ApplyAdvancedGameSettings(ctx context.Context, settings AdvancedGameSettings) (bool, error) {
	// Function doesn't reply with a body of info just status code, so handle this specifically.
	functionBody, err := json.Marshal(ApplyAdvancedGameSettingsRequest{
		Function: ApplyAdvancedGameSettingsFunction,
		Data:     settings,
	})
	if err != nil {
		return false, err
	}
	request, err := c.CreatePostRequest(ApplyAdvancedGameSettingsFunction, functionBody)
	if err != nil {
		return false, err
	}
	apiErr, err := c.SendPostRequest(ctx, request, functionBody)
	if err != nil {
		return false, err
	}
	if apiErr != nil {
		return false, apiErr
	}
	return true, nil
}
