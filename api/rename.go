package api

import (
	"context"
	"encoding/json"
)

// RenameRequest represents a request to rename the Satisfactory server.
type RenameRequest struct {
	// Function specifies the API function to call for renaming the server.
	Function string `json:"function,omitempty"`

	// Data contains the new server name.
	Data RenameRequestData `json:"data,omitempty"`
}

// RenameRequestData holds the new server name to be set.
type RenameRequestData struct {
	// ServerName is the new name for the Satisfactory server.
	ServerName string `json:"serverName,omitempty"`
}

// RenameServer renames the Satisfactory dedicated server to the specified serverName.
func (c *GoFactoryClient) RenameServer(ctx context.Context, serverName string) error {
	functionBody, err := json.Marshal(RenameRequest{
		Function: RenameServerFunction,
		Data: RenameRequestData{
			ServerName: serverName,
		},
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequest(RenameServerFunction, functionBody)
	if err != nil {
		return err
	}

	var apiError APIError
	err = c.SendPostRequest(ctx, req, &apiError)

	if err != nil {
		return err
	}

	return nil
}
