package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// ClaimRequest represents a request to claim or register a Satisfactory server.
type ClaimRequest struct {
	// Function specifies the API function to be called during POST.
	Function string `json:"function,omitempty"`

	// Data contains the information required to claim the server.
	Data ClaimRequestData `json:"data,omitempty"`
}

// ClaimRequestData holds the necessary data to claim a Satisfactory server.
type ClaimRequestData struct {
	// ServerName is what the claimed server will be called.
	ServerName string `json:"serverName,omitempty"`

	// AdminPassword is the administrator password to set onto the claimed server
	AdminPassword string `json:"adminPassword,omitempty"`
}

// ClaimResponse represents the response returned after successfully claiming a server.
type ClaimResponse struct {
	// Data contains the response body of the claimed servers authentication token
	Data ClaimResponseData `json:"data,omitempty"`
}

// ClaimResponseData holds the authentication token for the claimed server.
type ClaimResponseData struct {
	// AuthenticationToken is the `Administrator` privilege authentication token for the claimed server.
	AuthenticationToken string `json:"authenticationToken,omitempty"`
}

// ClaimServer claims a Satisfactory server using the provided ClaimRequestData struct.
// It updates the client's authentication token and privilege level upon success,
// and verifies the server state after the claim is completed.
func (c *GoFactoryClient) ClaimServer(ctx context.Context, claimData ClaimRequestData) error {
	if c.currentPrivilege != INITIAL_ADMIN_PRIVILEGE {
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
	c.currentPrivilege = ADMINISTRATOR_PRIVILEGE

	_, err = c.QueryServerState(ctx)
	if err != nil {
		return err
	}

	return nil
}
