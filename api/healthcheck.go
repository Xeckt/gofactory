package api

import (
	"context"
	"encoding/json"
)

// HealthCheckResponse represents the response returned by the server's health check,
// including health status and any custom server data.
type HealthCheckResponse struct {
	// Health indicates the current health status of the server.
	Health string `json:"health,omitempty"`

	// CustomData contains any custom server-specific data.
	CustomData string `json:"serverCustomData,omitempty"`
}

// HealthCheckResponseData wraps the HealthCheckResponse
// within a data field, matching the server's API response format.
type HealthCheckResponseData struct {
	// Data holds the actual health check response data.
	Data HealthCheckResponse `json:"data,omitempty"`
}

// HealthCheckCustomData represents optional custom data
// that the client can send to the server as part of the health check.
type HealthCheckCustomData struct {
	// CustomData is any custom data sent by the client.
	CustomData string `json:"clientCustomData,omitempty"`
}

// HealthCheckRequest represents the request body for performing
// a server health check with optional custom client data.
type HealthCheckRequest struct {
	// Function specifies the API function to call for the health check.
	Function string `json:"function"`

	// Data contains any custom client data for the health check.
	Data HealthCheckCustomData `json:"data"`
}

// GetServerHealth sends a health check request to the Satisfactory server,
// including any custom client data. It returns the server's health status
// and custom server data wrapped in a HealthCheckResponse.
func (c *GoFactoryClient) GetServerHealth(ctx context.Context, customData string) (*HealthCheckResponse, error) {
	functionBody, err := json.Marshal(HealthCheckRequest{
		Function: HealthCheckFunction,
		Data: HealthCheckCustomData{
			CustomData: customData,
		},
	})
	if err != nil {
		return nil, err
	}

	healthCheckResponse, err := CreateAndSendPostRequest[HealthCheckResponseData](ctx, c, HealthCheckFunction, functionBody)
	if err != nil {
		return nil, err
	}
	return &healthCheckResponse.Data, nil
}
