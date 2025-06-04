package api

import (
	"context"
	"encoding/json"
)

// AdvancedGameSettings represents advanced game rules and player settings
// for a Satisfactory save file.
type AdvancedGameSettings struct {
	// NoPower disables power consumption in the game world. Boolean represented as a string.
	NoPower string `json:"FG.GameRules.NoPower,omitempty"`

	// StartingTier sets the starting tier of the game for progression. Integer represented as a string.
	StartingTier string `json:"FG.GameRules.StartingTier,omitempty"`

	// DisableArachnidCreatures disables spider-like creatures in the game world. Boolean represented as a string.
	DisableArachnidCreatures string `json:"FG.GameRules.DisableArachnidCreatures,omitempty"`

	// NoUnlockCost removes the cost for unlocking milestones and upgrades. Boolean represented as a string.
	NoUnlockCost string `json:"FG.GameRules.NoUnlockCost,omitempty"`

	// SetGamePhase sets the current phase of the game (e.g., early, mid, late game). Integer represented as a string.
	SetGamePhase string `json:"FG.GameRules.SetGamePhase,omitempty"`

	// UnlockAllResearchSchematics instantly unlocks all research schematics. Boolean represented as a string.
	UnlockAllResearchSchematics string `json:"FG.GameRules.UnlockAllResearchSchematics,omitempty"`

	// UnlockInstantAltRecipes instantly unlocks all alternate recipes.Boolean represented as a string.
	UnlockInstantAltRecipes string `json:"FG.GameRules.UnlockInstantAltRecipes,omitempty"`

	// UnlockAllResourceSinkSchematics instantly unlocks all schematics
	// in the resource sink shop. Boolean represented as a string.
	UnlockAllResourceSinkSchematics string `json:"FG.GameRules.UnlockAllResourceSinkSchematics,omitempty"`

	// NoBuildCost removes the cost of building structures and machines. Boolean represented as a string.
	NoBuildCost string `json:"FG.PlayerRules.NoBuildCost,omitempty"`

	// GodMode enables invincibility for the player. Boolean represented as a string.
	GodMode string `json:"FG.PlayerRules.GodMode,omitempty"`

	// FlightMode enables the ability to fly around the game world. Boolean represented as a string.
	FlightMode string `json:"FG.PlayerRules.FlightMode,omitempty"`
}

// AdvancedGameSettingsData wraps the applied advanced game settings data
// for a Satisfactory save file.
type AdvancedGameSettingsData struct {
	// Data holds the actual applied advanced game settings.
	Data AppliedAdvancedGameSettings `json:"data,omitempty"`
}

// AppliedAdvancedGameSettings represents the specific advanced game settings
// that have been applied in a Satisfactory save file.
type AppliedAdvancedGameSettings struct {
	// Settings contains the advanced game rules and player settings.
	Settings AdvancedGameSettings `json:"appliedAdvancedGameSettings,omitempty"`
}

// ApplyAdvancedGameSettingsRequest represents a request to apply
// advanced game settings to a Satisfactory save file.
type ApplyAdvancedGameSettingsRequest struct {
	// Function specifies the function to call when applying the advanced settings.
	Function string `json:"function"`

	// Data contains the advanced game settings to apply.
	Data AdvancedGameSettings `json:"appliedAdvancedGameSettings"`
}

// GetAdvancedGameSettings retrieves the currently applied advanced game settings
// from the active Satisfactory save file.
func (c *GoFactoryClient) GetAdvancedGameSettings(ctx context.Context) (*AdvancedGameSettings, error) {
	appliedAdvanceSettingsResponse, err := CreateAndSendPostRequest[AppliedAdvancedGameSettings](ctx, c,
		GetAdvancedGameSettingsFunction,
		CreateGenericFunctionBody(GetAdvancedGameSettingsFunction))
	if err != nil {
		return nil, err
	}
	return &appliedAdvanceSettingsResponse.Settings, nil
}

// ApplyAdvancedGameSettings applies the provided AdvancedGameSettings
// to the current Satisfactory save file.
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
