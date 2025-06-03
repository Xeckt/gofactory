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

type PasswordlessLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel,omitempty"`
}

type PasswordlessLoginResponse struct {
	Data PasswordlessLoginResponseData `json:"data,omitempty"`
}

type PasswordlessLoginResponseData struct {
	AuthToken string `json:"authenticationToken,omitempty"`
}

func (c *GoFactoryClient) PasswordlessLogin(ctx context.Context, privilege string) (*PasswordlessLoginResponseData, error) {
	functionBody, err := json.Marshal(PasswordlessLoginRequest{
		Function: PasswordlessLoginFunction,
		Data: PasswordlessLoginRequestData{
			MinimumPrivilegeLevel: privilege,
		},
	})
	if err != nil {
		return nil, err
	}

	tokenResponse, err := CreateAndSendPostRequest[PasswordlessLoginResponse](ctx, c, PasswordlessLoginFunction, functionBody)
	if err != nil {
		return nil, err
	}

	return &tokenResponse.Data, nil
}
