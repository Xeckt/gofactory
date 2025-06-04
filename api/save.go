package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
)

// CreateNewGameRequest represents the request payload for creating a new game session.
type CreateNewGameRequest struct {
	// Function specifies the API function to call for creating a new game session.
	Function string `json:"function"`

	// Data contains the details needed to initialize the new game session.
	Data CreateNewGameData `json:"data"`
}

// CreateNewGameData wraps the actual game data for creating a new game session.
type CreateNewGameData struct {
	// GameData holds the parameters and configuration for the new game session.
	GameData CreateNewGameRequestData `json:"newGameData"`
}

// CreateNewGameRequestData contains the parameters required to create
// a new game session in the Satisfactory dedicated server.
type CreateNewGameRequestData struct {
	// SessionName is the name of the new game session.
	SessionName string `json:"sessionName"`

	// MapName is the name of the map to use for the new game session. Can be left blank for default level.
	MapName string `json:"mapName"`

	// StartingLocation is the location on the map where the session will start. Can be left blank for random location.
	StartingLocation string `json:"startingLocation"`

	// BSkipOnboarding specifies whether the onboarding process should be skipped.
	BSkipOnboarding bool `json:"bSkipOnboarding"`

	// AdvancedGameSettings contains advanced game rules and settings
	// to apply to the new game session. This field is optional.
	AdvancedGameSettings AppliedAdvancedGameSettings `json:"advancedGameSettings,omitempty"`

	// TODO: customOptionsOnlyForModding interface{}
}

// CreateNewGame sends a request to the Satisfactory dedicated server to create a new game session.
// It takes a CreateNewGameRequestData struct parameter to specify the
// session name, map, starting location, and advanced game settings.
func (c *GoFactoryClient) CreateNewGame(ctx context.Context, newGameData CreateNewGameRequestData) error {
	functionBody, err := json.Marshal(CreateNewGameRequest{
		Function: CreateNewGameFunction,
		Data: CreateNewGameData{
			GameData: newGameData,
		},
	})
	if err != nil {
		return err
	}

	request, err := c.CreatePostRequest(CreateNewGameFunction, functionBody)
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

// SaveGameRequest represents a request to save the current loaded save game.
type SaveGameRequest struct {
	// Function specifies the API function to call for saving the game.
	Function string `json:"function"`

	// Data contains the details of the save game request.
	Data SaveGameData `json:"data"`
}

// SaveGameData holds the name of the sav.
type SaveGameData struct {
	// SaveName is the name to apply to the save.
	SaveName string `json:"saveName"`
}

// SaveGame saves the current game session and applies `saveName` as the name.
func (c *GoFactoryClient) SaveGame(ctx context.Context, saveName string) error {
	functionBody, err := json.Marshal(SaveGameRequest{
		Function: SaveGameFunction,
		Data: SaveGameData{
			SaveName: saveName,
		},
	})
	if err != nil {
		return err
	}
	request, err := c.CreatePostRequest(SaveGameFunction, functionBody)
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

// DeleteSaveRequest represents a request to delete a specific save file
// from the Satisfactory dedicated server.
type DeleteSaveRequest struct {
	// Function specifies the API function to call for deleting a save file.
	Function string `json:"function"`

	// Data contains the details of the save file to delete.
	Data DeleteSaveData `json:"data"`
}

// DeleteSaveData holds the name of the save file to be deleted.
type DeleteSaveData struct {
	// SaveName is the name of the save file to delete.
	SaveName string `json:"saveName"`
}

// DeleteSave deletes the save file matching saveName
func (c *GoFactoryClient) DeleteSave(ctx context.Context, saveName string) error {
	functionBody, err := json.Marshal(DeleteSaveRequest{
		Function: DeleteSaveFileFunction,
		Data: DeleteSaveData{
			SaveName: saveName,
		},
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequest(DeleteSaveFileFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, req, &apiError)
	if err != nil {
		return err
	}

	if apiError != (APIError{}) {
		return &apiError
	}

	return nil
}

// DeleteSaveSessionRequest represents a request to delete a specific save session
// from the Satisfactory dedicated server. It includes the function name and the session data.
type DeleteSaveSessionRequest struct {
	// Function specifies the API function to call for deleting a save session.
	Function string `json:"function"`

	// Data contains the details of the save session to delete.
	Data DeleteSaveSessionData `json:"data"`
}

// DeleteSaveSessionData holds the name of the save session to be deleted.
type DeleteSaveSessionData struct {
	// SessionName is the name of the save session to delete.
	SessionName string `json:"sessionName"`
}

// DeleteSaveSession deletes all saves files that belong to a specific session.
func (c *GoFactoryClient) DeleteSaveSession(ctx context.Context, sessionName string) error {
	functionBody, err := json.Marshal(DeleteSaveSessionRequest{
		Function: DeleteSaveSessionFunction,
		Data: DeleteSaveSessionData{
			SessionName: sessionName,
		},
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequest(DeleteSaveSessionFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, req, &apiError)
	if err != nil {
		return err
	}

	if apiError != (APIError{}) {
		return &apiError
	}

	return nil
}

// EnumerateSessionsResponse represents the response from the Satisfactory server
// when requesting a list of available game sessions.
type EnumerateSessionsResponse struct {
	// Data contains the session data returned by the server.
	Data EnumerateSessionsResponseData `json:"data"`
}

// EnumerateSessionsResponseData contains the list of game sessions and
// the index of the currently active session.
type EnumerateSessionsResponseData struct {
	// Sessions is a slice of available game sessions.
	Sessions []EnumerateSessionsResponseDataArray `json:"sessions"`

	// CurrentSessionIndex is the index of the current session selected in the slice.
	CurrentSessionIndex int `json:"currentSessionIndex"`
}

// EnumerateSessionsResponseDataArray holds information about a specific game session.
type EnumerateSessionsResponseDataArray struct {
	// SessionName is the name of the game session.
	SessionName string `json:"sessionName"`

	// SaveHeaders contains detailed information about all save files belonging to the session.
	SaveHeaders []EnumerateSessionsSaveHeader `json:"saveHeaders"`
}

// EnumerateSessionsSaveHeader represents detailed metadata for a save file.
type EnumerateSessionsSaveHeader struct {
	// SaveVersion is the version of the save file format.
	SaveVersion int `json:"saveVersion"`

	// BuildVersion is the version of the game build used to create the save file.
	BuildVersion int `json:"buildVersion"`

	// SaveName is the name of the save file.
	SaveName string `json:"saveName"`

	// SaveLocationInfo is currently unknown on exactly what this is, not listed in official Ficsit documentation.
	SaveLocationInfo string `json:"saveLocationInfo"`

	// MapName is the name of the map used in this save.
	MapName string `json:"mapName"`

	// MapOptions contains any additional options used when the map was created.
	MapOptions string `json:"mapOptions"`

	// SessionName is the name of the session that owns this save.
	SessionName string `json:"sessionName"`

	// PlayDurationSeconds is the total play time of this save file in seconds.
	PlayDurationSeconds int `json:"playDurationSeconds"`

	// SaveDateTime is the date and time when this game file was saved.
	SaveDateTime string `json:"saveDateTime"`

	// IsModdedSave indicates whether this save was saved with mods.
	IsModdedSave bool `json:"isModdedSave"`

	// IsEditedSave indicates whether this save has been edited by third party tools.
	IsEditedSave bool `json:"isEditedSave"`

	// IsCreativeModeEnabled indicates whether Advanced Game settings is enabled for this save.
	IsCreativeModeEnabled bool `json:"isCreativeModeEnabled"`
}

// EnumerateSessions will return a complete slice of all sessions enumerated in the Satisfactory dedicated server.
func (c *GoFactoryClient) EnumerateSessions(ctx context.Context) (*EnumerateSessionsResponseData, error) {
	response, err := CreateAndSendPostRequest[EnumerateSessionsResponse](ctx, c,
		EnumerateSessionsFunction, CreateGenericFunctionBody(EnumerateSessionsFunction))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// LoadGameRequest represents a request to load a saved game session
// in the Satisfactory server API. It includes the function name
// and the data required to load the save.
type LoadGameRequest struct {
	// Function specifies the API function to call for loading a saved game.
	Function string `json:"function"`

	// Data contains the details of the save game to load.
	Data LoadGameRequestData `json:"data"`
}

// LoadGameRequestData holds the parameters for loading a saved game session.
type LoadGameRequestData struct {
	// SaveName is the name of the save file to load.
	SaveName string `json:"saveName"`

	// EnableAdvanceGameSettings specifies whether advanced game settings
	// should be applied when loading the save.
	EnableAdvanceGameSettings bool `json:"enableAdvanceGameSettings"`
}

// LoadGame will load the specified game save that matches `saveName` and an boolean to specify
// if Advanced Game Settings should be enabled when this save is loaded.
func (c *GoFactoryClient) LoadGame(ctx context.Context, saveName string, enableAdvancedSettings bool) error {
	functionBody, err := json.Marshal(LoadGameRequest{
		Function: LoadGameFunction,
		Data: LoadGameRequestData{
			SaveName:                  saveName,
			EnableAdvanceGameSettings: enableAdvancedSettings,
		},
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequest(LoadGameFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, req, &apiError)
	if err != nil {
		return err
	}

	if apiError != (APIError{}) {
		return &apiError
	}

	return nil
}

// UploadSaveGameRequest represents the request to upload a save game file
// to the Satisfactory server, including the data and file buffer.
type UploadSaveGameRequest struct {
	// Data contains the function name and save game details.
	Data UploadSaveGameData `json:"data"`

	// SaveFile is the binary data of the save game file to upload.
	SaveFile bytes.Buffer `json:"saveGameFile"`
}

// UploadSaveGameData contains the API function name and the details
// about how the save game should be handled after upload.
type UploadSaveGameData struct {
	// Function specifies the API function to call for uploading the save game.
	Function string `json:"function"`

	// SaveData contains the save game parameters.
	SaveData UploadSaveGameDataRequest `json:"data"`
}

// UploadSaveGameDataRequest holds the details for uploading a save game,
// including the save name and options for loading it immediately.
type UploadSaveGameDataRequest struct {
	// SaveName is the name to assign to the uploaded save file.
	SaveName string `json:"saveName"`

	// LoadImmediately specifies whether to load the save game immediately after upload.
	LoadImmediately bool `json:"loadSaveGame"`

	// EnableAdvanceGameSettings specifies whether advanced game settings
	// should be applied when loading the uploaded save.
	EnableAdvanceGameSettings bool `json:"enableAdvanceGameSettings"`
}

// UploadSaveGame uploads a save game file to the Satisfactory server.
// It streams the file from a Reader and informs the server what to do with it through the UploadSaveGameDataRequest
// parameter.
func (c *GoFactoryClient) UploadSaveGame(ctx context.Context, fileStream io.Reader, filename string, saveSettings UploadSaveGameDataRequest) error {
	var bodyBuffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&bodyBuffer)
	defer multipartWriter.Close()

	fileWriter, err := multipartWriter.CreateFormFile("save", filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, fileStream)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type": multipartWriter.FormDataContentType(),
	}

	functionBody, err := json.Marshal(UploadSaveGameRequest{
		Data: UploadSaveGameData{
			Function: UploadSaveGameFunction,
			SaveData: saveSettings,
		},
		SaveFile: bodyBuffer,
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequestWithHeaders(headers, UploadSaveGameFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, req, &apiError)
	if err != nil {
		return err
	}

	if apiError != (APIError{}) {
		return &apiError
	}

	return nil
}

// DownloadSaveGameRequest represents a request to download a save game file
// from the Satisfactory server, specifying the save file name.
type DownloadSaveGameRequest struct {
	// Function specifies the API function to call for downloading the save game.
	Function string `json:"function"`

	// Data contains the name of the save file to download.
	Data DownloadSaveGameRequestData `json:"data"`
}

// DownloadSaveGameRequestData holds the name of the save file to download.
type DownloadSaveGameRequestData struct {
	// SaveName is the name of the save file to download.
	SaveName string `json:"saveName"`
}

// DownloadSaveGame downloads a save game file from the Satisfactory dedicated server.
// It returns a binary []byte stream of the file from the returned request body.
func (c *GoFactoryClient) DownloadSaveGame(ctx context.Context, saveName string) ([]byte, error) {
	functionBody, err := json.Marshal(DownloadSaveGameRequest{
		Function: DownloadSaveGameFunction,
		Data: DownloadSaveGameRequestData{
			SaveName: saveName,
		},
	})
	if err != nil {
		return nil, err
	}

	req, err := c.CreatePostRequest(DownloadSaveGameFunction, functionBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileStream, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileStream, nil
}
