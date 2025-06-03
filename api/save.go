package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
)

type CreateNewGameRequest struct {
	Function string            `json:"function"`
	Data     CreateNewGameData `json:"data"`
}

type CreateNewGameData struct {
	GameData CreateNewGameRequestData `json:"newGameData"`
}

type CreateNewGameRequestData struct {
	SessionName          string                      `json:"sessionName"`
	MapName              string                      `json:"mapName"`
	StartingLocation     string                      `json:"startingLocation"`
	BSkipOnboarding      bool                        `json:"bSkipOnboarding"`
	AdvancedGameSettings AppliedAdvancedGameSettings `json:"advancedGameSettings,omitempty"`
	// TODO: customOptionsOnlyForModding interface{}
}

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

type SaveGameRequest struct {
	Function string       `json:"function"`
	Data     SaveGameData `json:"data"`
}

type SaveGameData struct {
	SaveName string `json:"saveName"`
}

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

type DeleteSaveRequest struct {
	Function string         `json:"function"`
	Data     DeleteSaveData `json:"data"`
}

type DeleteSaveData struct {
	SaveName string `json:"saveName"`
}

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

type DeleteSaveSessionRequest struct {
	Function string                `json:"function"`
	Data     DeleteSaveSessionData `json:"data"`
}

type DeleteSaveSessionData struct {
	SessionName string `json:"sessionName"`
}

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

type EnumerateSessionsResponse struct {
	Data EnumerateSessionsResponseData `json:"data"`
}

type EnumerateSessionsResponseData struct {
	Sessions            []EnumerateSessionsResponseDataArray `json:"sessions"`
	CurrentSessionIndex int                                  `json:"currentSessionIndex"`
}

type EnumerateSessionsResponseDataArray struct {
	SessionName string                        `json:"sessionName"`
	SaveHeaders []EnumerateSessionsSaveHeader `json:"saveHeaders"`
}

type EnumerateSessionsSaveHeader struct {
	SaveVersion           int    `json:"saveVersion"`
	BuildVersion          int    `json:"buildVersion"`
	SaveName              string `json:"saveName"`
	SaveLocationInfo      string `json:"saveLocationInfo"`
	MapName               string `json:"mapName"`
	MapOptions            string `json:"mapOptions"`
	SessionName           string `json:"sessionName"`
	PlayDurationSeconds   int    `json:"playDurationSeconds"`
	SaveDateTime          string `json:"saveDateTime"`
	IsModdedSave          bool   `json:"isModdedSave"`
	IsEditedSave          bool   `json:"isEditedSave"`
	IsCreativeModeEnabled bool   `json:"isCreativeModeEnabled"`
}

func (c *GoFactoryClient) EnumerateSessions(ctx context.Context) (*EnumerateSessionsResponseData, error) {
	response, err := CreateAndSendPostRequest[EnumerateSessionsResponse](ctx, c,
		EnumerateSessionsFunction, CreateGenericFunctionBody(EnumerateSessionsFunction))
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

type LoadGameRequest struct {
	Function string              `json:"function"`
	Data     LoadGameRequestData `json:"data"`
}

type LoadGameRequestData struct {
	SaveName                  string `json:"saveName"`
	EnableAdvanceGameSettings bool   `json:"enableAdvanceGameSettings"`
}

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

type UploadSaveGameRequest struct {
	Data     UploadSaveGameData `json:"data"`
	SaveFile bytes.Buffer       `json:"saveGameFile"`
}

type UploadSaveGameData struct {
	Function string                    `json:"function"`
	SaveData UploadSaveGameDataRequest `json:"data"`
}

type UploadSaveGameDataRequest struct {
	SaveName                  string `json:"saveName"`
	LoadImmediately           bool   `json:"loadSaveGame"`
	EnableAdvanceGameSettings bool   `json:"enableAdvanceGameSettings"`
}

func (c *GoFactoryClient) UploadSaveGame(ctx context.Context, fileStream io.Reader, filename string, saveData UploadSaveGameDataRequest) error {
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
			SaveData: saveData,
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
