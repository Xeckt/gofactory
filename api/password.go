package api

import (
	"context"
	"encoding/json"
	"fmt"
)

// Privilege level constants used for authentication.
const (
	// NOT_AUTHENTICATED_PRIVILEGE represents a user who is not authenticated.
	NOT_AUTHENTICATED_PRIVILEGE string = "NotAuthenticated"

	// CLIENT_PRIVILEGE represents a standard client privilege level.
	CLIENT_PRIVILEGE string = "Client"

	// ADMINISTRATOR_PRIVILEGE represents administrator-level privileges.
	ADMINISTRATOR_PRIVILEGE string = "Administrator"

	// INITIAL_ADMIN_PRIVILEGE represents the initial admin privilege level during claim.
	INITIAL_ADMIN_PRIVILEGE string = "InitialAdmin"

	// API_TOKEN_PRIVILEGE represents an API token-based privilege level.
	API_TOKEN_PRIVILEGE string = "ApiToken"
)

// PasswordlessLoginRequest represents a request to authenticate using
// passwordless login with a minimum required privilege level.
type PasswordlessLoginRequest struct {
	// Function specifies the API function to call for passwordless login.
	Function string `json:"function"`

	// Data contains the required privilege level for login.
	Data PasswordlessLoginRequestData `json:"data"`
}

// PasswordLoginRequest represents a request to authenticate using
// a password and a required privilege level.
type PasswordLoginRequest struct {
	// Function specifies the API function to call for password-based login.
	Function string `json:"function"`

	// Data contains the required privilege level and password.
	Data PasswordLoginRequestData `json:"data"`
}

// PasswordLoginRequestData contains the minimum privilege level
// and password for password-based authentication.
type PasswordLoginRequestData struct {
	// MinimumPrivilegeLevel specifies the minimum required privilege for authentication.
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`

	// Password is the password used for authentication.
	Password string `json:"password"`
}

// PasswordlessLoginRequestData contains the minimum privilege level
// required for passwordless login.
type PasswordlessLoginRequestData struct {
	// MinimumPrivilegeLevel specifies the minimum required privilege for authentication.
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`
}

// LoginResponse represents the API response for authentication requests.
type LoginResponse struct {
	// Data contains the authentication token returned by the server.
	Data PasswordLoginResponseData `json:"data,omitempty"`
}

// PasswordLoginResponseData holds the authentication token
// returned after a successful authentication request.
type PasswordLoginResponseData struct {
	// AuthToken is the token used for subsequent API requests.
	AuthToken string `json:"authenticationToken,omitempty"`
}

// PasswordlessLogin authenticates the client using passwordless login on an unclaimed server or
// when client protection password is not set. Updates the *GoFactoryClient.Token field with the new token.
func (c *GoFactoryClient) PasswordlessLogin(ctx context.Context, privilege string) error {
	functionBody, err := json.Marshal(PasswordlessLoginRequest{
		Function: PasswordlessLoginFunction,
		Data: PasswordlessLoginRequestData{
			MinimumPrivilegeLevel: privilege,
		},
	})
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	tokenResponse, err := CreateAndSendPostRequestWithHeaders[LoginResponse](ctx, c, headers, PasswordlessLoginFunction, functionBody)
	if err != nil {
		return err
	}

	c.currentPrivilege = privilege
	c.Token = tokenResponse.Data.AuthToken

	return nil
}

// PasswordLogin authenticates the client using a password for the specified privilege level.
// Updates the *GoFactoryClient.Token field with the new token.
func (c *GoFactoryClient) PasswordLogin(ctx context.Context, privilege string, password string) error {
	functionBody, err := json.Marshal(PasswordLoginRequest{
		Function: PasswordLoginFunction,
		Data: PasswordLoginRequestData{
			MinimumPrivilegeLevel: privilege,
			Password:              password,
		},
	})
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	tokenResponse, err := CreateAndSendPostRequestWithHeaders[LoginResponse](ctx, c, headers, PasswordLoginFunction, functionBody)
	if err != nil {
		return err
	}

	c.currentPrivilege = privilege
	c.Token = tokenResponse.Data.AuthToken
	return nil
}

// ClientPasswordRequest represents a request to set the client password.
type ClientPasswordRequest struct {
	// Function specifies the API function to call for setting the client password.
	Function string `json:"function,omitempty"`

	// Data contains the new client password.
	Data ClientPasswordRequestData `json:"data,omitempty"`
}

// ClientPasswordRequestData holds the new client password.
type ClientPasswordRequestData struct {
	// Password is the new client password.
	Password string `json:"password,omitempty"`
}

// SetClientPassword sets a new client password for the Satisfactory dedicated server.
func (c *GoFactoryClient) SetClientPassword(ctx context.Context, newPassword string) error {
	functionBody, err := json.Marshal(ClientPasswordRequest{
		Function: SetClientPasswordFunction,
		Data: ClientPasswordRequestData{
			Password: newPassword,
		},
	})
	if err != nil {
		return err
	}

	req, err := c.CreatePostRequest(SetClientPasswordFunction, functionBody)
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

// AdminPasswordRequest represents a request to set the administrator password.
type AdminPasswordRequest struct {
	// Function specifies the API function to call for setting the admin password.
	Function string `json:"function,omitempty"`

	// Data contains the new administrator password.
	Data AdminPasswordRequestData `json:"data,omitempty"`
}

// AdminPasswordRequestData holds the new administrator password.
type AdminPasswordRequestData struct {
	// Password is the new administrator password.
	Password string `json:"password,omitempty"`
}

// AdminPasswordResponseData wraps the admin password response from the server.
type AdminPasswordResponseData struct {
	// Data contains the authentication token after setting the admin password.
	Data AdminPasswordResponse `json:"data,omitempty"`
}

// AdminPasswordResponse represents the response containing the new authentication token.
type AdminPasswordResponse struct {
	// AuthToken is the new authentication token after setting the admin password.
	AuthToken string `json:"authenticationToken,omitempty"`
}

// SetAdminPassword sets a new administrator password on the Satisfactory server
// and updates the *GoFactoryClient.Token with the new token.
// This POST requests invalidates all previous Client and Admin tokens.
func (c *GoFactoryClient) SetAdminPassword(ctx context.Context, newPassword string) error {
	functionBody, err := json.Marshal(AdminPasswordRequest{
		Function: SetAdminPasswordFunction,
		Data: AdminPasswordRequestData{
			Password: newPassword,
		},
	})
	if err != nil {
		return err
	}

	resp, err := CreateAndSendPostRequest[AdminPasswordResponseData](ctx, c, SetAdminPasswordFunction, functionBody)
	if err != nil {
		return err
	}

	c.Token = resp.Data.AuthToken
	return nil
}
