package api

import (
	"context"
	"encoding/json"
	"fmt"
)

const NOT_AUTHENTICATED_PRIVILEGE string = "NotAuthenticated"
const CLIENT_PRIVILEGE string = "ClientPrivilege"
const ADMINISTRATOR_PRIVILEGE string = "Administrator"
const INITIAL_ADMIN_PRIVILEGE string = "InitialAdmin"
const API_TOKEN_PRIVILEGE string = "ApiToken"

type PasswordlessLoginRequest struct {
	Function string                       `json:"function"`
	Data     PasswordlessLoginRequestData `json:"data"`
}

type PasswordLoginRequest struct {
	Function string                   `json:"function"`
	Data     PasswordLoginRequestData `json:"data"`
}

type PasswordLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`
	Password              string `json:"password"`
}

type PasswordlessLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`
}

type LoginResponse struct {
	Data PasswordLoginResponseData `json:"data,omitempty"`
}

type PasswordLoginResponseData struct {
	AuthToken string `json:"authenticationToken,omitempty"`
}

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
		fmt.Println("Error is here", tokenResponse, err)
		return err
	}

	c.currentPrivilege = privilege
	c.Token = tokenResponse.Data.AuthToken
	return nil
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

type ClientPasswordRequest struct {
	Function string                    `json:"function,omitempty"`
	Data     ClientPasswordRequestData `json:"data,omitempty"`
}

type ClientPasswordRequestData struct {
	Password string `json:"password,omitempty"`
}

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

type AdminPasswordRequest struct {
	Function string                   `json:"function,omitempty"`
	Data     AdminPasswordRequestData `json:"data,omitempty"`
}

type AdminPasswordRequestData struct {
	Password string `json:"password,omitempty"`
}

type AdminPasswordResponseData struct {
	Data AdminPasswordResponse `json:"data,omitempty"`
}

type AdminPasswordResponse struct {
	AuthToken string `json:"authenticationToken,omitempty"`
}

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
