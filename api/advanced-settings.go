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
	Data AppliedAdvancedGameSettings `json:"data,omitempty"`
}

type AppliedAdvancedGameSettings struct {
	Settings AdvancedGameSettings `json:"appliedAdvancedGameSettings,omitempty"`
}
type ApplyAdvancedGameSettingsRequest struct {
	Function string               `json:"function"`
	Data     AdvancedGameSettings `json:"appliedAdvancedGameSettings"`
}

func (c *GoFactoryClient) GetAdvancedGameSettings(ctx context.Context) (*AdvancedGameSettings, error) {
	appliedAdvanceSettingsResponse, err := CreateAndSendPostRequest[AppliedAdvancedGameSettings](ctx, c,
		GetAdvancedGameSettingsFunction,
		CreateGenericFunctionBody(GetAdvancedGameSettingsFunction))
	if err != nil {
		return nil, err
	}
	return &appliedAdvanceSettingsResponse.Settings, nil
}

func (c *GoFactoryClient) ApplyAdvancedGameSettings(ctx context.Context, settings AdvancedGameSettings) error {
	functionBody, err := json.Marshal(ApplyAdvancedGameSettingsRequest{
		Function: ApplyAdvancedGameSettingsFunction,
		Data:     settings,
	})
	if err != nil {
		return err
	}
	request, err := c.CreatePostRequest(ApplyAdvancedGameSettingsFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, request, &apiError)
	if err != nil {
		return err
	}

	if apiError != (APIError{}) {
		return &apiError
	}

	return nil
}
