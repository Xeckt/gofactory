package api

import (
	"context"
	"encoding/json"
)

const NOT_AUTHENTICATED_PRIVILEGE string = "NotAuthenticated"
const CLIENT_PRIVILEGE string = "ClientPrivilege"
const ADMINISTRATOR_PRIVILEGE string = "Administrator"
const INITIAL_ADMIN_PRIVILEGE string = "InitialAdmin"
const API_TOKEN_PRIVILEGE string = "ApiToken"

type PasswordlessLoginRequest struct {
	Function string                       `json:"function,omitempty"`
	Data     PasswordlessLoginRequestData `json:"data,omitempty"`
}

type PasswordLoginRequest struct {
	Function string                   `json:"function,omitempty"`
	Data     PasswordLoginRequestData `json:"data,omitempty"`
}

type PasswordLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel,omitempty"`
	Password              string `json:"password,omitempty"`
}

type PasswordlessLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel,omitempty"`
}

type LoginResponse struct {
	Data PasswordLoginResponseData `json:"data,omitempty"`
}

type PasswordLoginResponseData struct {
	AuthToken string `json:"authenticationToken,omitempty"`
}

func (c *GoFactoryClient) PasswordlessLogin(ctx context.Context, privilege string) (*PasswordLoginResponseData, error) {
	functionBody, err := json.Marshal(PasswordlessLoginRequest{
		Function: PasswordlessLoginFunction,
		Data: PasswordlessLoginRequestData{
			MinimumPrivilegeLevel: privilege,
		},
	})
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	tokenResponse, err := CreateAndSendPostRequestWithHeaders[LoginResponse](ctx, c, headers, PasswordlessLoginFunction, functionBody)
	if err != nil {
		return nil, err
	}
	c.currentPrivilege = privilege
	return &tokenResponse.Data, nil
}

func (c *GoFactoryClient) PasswordLogin(ctx context.Context, privilege string, password string) (*PasswordLoginResponseData, error) {
	functionBody, err := json.Marshal(PasswordLoginRequest{
		Function: PasswordLoginFunction,
		Data: PasswordLoginRequestData{
			MinimumPrivilegeLevel: privilege,
			Password:              password,
		},
	})

	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	tokenResponse, err := CreateAndSendPostRequestWithHeaders[LoginResponse](ctx, c, headers, PasswordLoginFunction, functionBody)
	if err != nil {
		return nil, err
	}
	c.currentPrivilege = privilege
	return &tokenResponse.Data, nil
}
