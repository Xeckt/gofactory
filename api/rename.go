package api

import (
	"context"
	"encoding/json"
)

type RenameRequest struct {
	Function string            `json:"function,omitempty"`
	Data     RenameRequestData `json:"data,omitempty"`
}

type RenameRequestData struct {
	ServerName string `json:"serverName,omitempty"`
}

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
