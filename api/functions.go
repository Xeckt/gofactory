package api

import (
	"encoding/json"
	"log"
)

// genericFunctionBody represents a simple JSON body containing
// only the function name to call in the Satisfactory server API.
type genericFunctionBody struct {
	Function string `json:"function"`
}

// CreateGenericFunctionBody marshals the specified function name
// into a JSON-encoded body for API requests. Specifically used
// with POST requests that require no additional parameters.
func CreateGenericFunctionBody(function string) []byte {
	body, err := json.Marshal(genericFunctionBody{
		Function: function,
	})
	if err != nil {
		log.Fatal(err)
	}
	return body
}

// String constants representing API function names used by the Satisfactory dedicated server API.
const (
	// HealthCheckFunction is the API function name for the server health check.
	HealthCheckFunction = "HealthCheck"

	// VerifyAuthTokenFunction verifies the provided authentication token.
	VerifyAuthTokenFunction = "VerifyAuthenticationToken"

	// PasswordlessLoginFunction initiates passwordless login for the client.
	PasswordlessLoginFunction = "PasswordlessLogin"

	// PasswordLoginFunction initiates password-based login for the client.
	PasswordLoginFunction = "PasswordLogin"

	// QueryServerStateFunction queries the current state of the server.
	QueryServerStateFunction = "QueryServerState"

	// GetServerOptionsFunction retrieves the current server options.
	GetServerOptionsFunction = "GetServerOptions"

	// GetAdvancedGameSettingsFunction retrieves advanced game settings for the save file.
	GetAdvancedGameSettingsFunction = "GetAdvancedGameSettings"

	// ApplyAdvancedGameSettingsFunction applies advanced game settings to the save file.
	ApplyAdvancedGameSettingsFunction = "ApplyAdvancedGameSettings"

	// ClaimServerFunction claims the server for administration.
	ClaimServerFunction = "ClaimServer"

	// RenameServerFunction renames the server.
	RenameServerFunction = "RenameServer"

	// SetClientPasswordFunction sets the client password for joining the server.
	SetClientPasswordFunction = "SetClientPassword"

	// SetAdminPasswordFunction sets the administrator password for the server.
	SetAdminPasswordFunction = "SetAdminPassword"

	// SetAutoLoadSessionNameFunction sets the session name to auto-load on server startup.
	SetAutoLoadSessionNameFunction = "SetAutoLoadSessionName"

	// RunCommandFunction executes a console command on the server.
	RunCommandFunction = "RunCommand"

	// ShutdownFunction shuts down the Satisfactory server.
	ShutdownFunction = "Shutdown"

	// ApplyServerOptionsFunction applies new server options.
	ApplyServerOptionsFunction = "ApplyServerOptions"

	// CreateNewGameFunction creates a new game session.
	CreateNewGameFunction = "CreateNewGame"

	// SaveGameFunction saves the current game session.
	SaveGameFunction = "SaveGame"

	// DeleteSaveFileFunction deletes a save file from the server.
	DeleteSaveFileFunction = "DeleteSaveFile"

	// DeleteSaveSessionFunction deletes a game session from the server.
	DeleteSaveSessionFunction = "DeleteSaveSession"

	// EnumerateSessionsFunction enumerates all available game sessions.
	EnumerateSessionsFunction = "EnumerateSessions"

	// LoadGameFunction loads a game session.
	LoadGameFunction = "LoadGame"

	// UploadSaveGameFunction uploads a save game file to the server.
	UploadSaveGameFunction = "UploadSaveGame"

	// DownloadSaveGameFunction downloads a save game file from the server.
	DownloadSaveGameFunction = "DownloadSaveGame"
)
