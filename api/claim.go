package api

import (
	"context"
	"encoding/json"
	"fmt"
)

type ClaimRequest struct {
	Function string           `json:"function,omitempty"`
	Data     ClaimRequestData `json:"data,omitempty"`
}

type ClaimRequestData struct {
	ServerName    string `json:"serverName,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
}

type ClaimResponse struct {
	Data ClaimResponseData `json:"data,omitempty"`
}

type ClaimResponseData struct {
	AuthenticationToken string `json:"authenticationToken,omitempty"`
}

func (c *GoFactoryClient) ClaimServer(ctx context.Context, claimData ClaimRequestData) error {
	if c.CurrentPrivilege != INITIAL_ADMIN_PRIVILEGE {
		return fmt.Errorf("privilege must be set to %s in order to claim the server", INITIAL_ADMIN_PRIVILEGE)
	}

	functionBody, err := json.Marshal(ClaimRequest{
		Function: ClaimServerFunction,
		Data:     claimData,
	})
	if err != nil {
		return err
	}
	newToken, err := CreateAndSendPostRequest[ClaimResponse](ctx, c, ClaimServerFunction, functionBody)
	if err != nil {
		return err
	}
	if newToken == nil {
		return fmt.Errorf("new authentication Token returned is empty")
	}
	c.Token = newToken.Data.AuthenticationToken

	_, err = c.QueryServerState(ctx)
	if err != nil {
		return err
	}

	return nil
}
