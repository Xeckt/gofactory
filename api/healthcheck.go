package api

import (
	"context"
	"encoding/json"
)

type HealthCheckResponse struct {
	Health     string `json:"health"`
	CustomData string `json:"serverCustomData"`
}

type HealthCheckResponseData struct {
	Data HealthCheckResponse `json:"data"`
}

type HealthCheckCustomData struct {
	CustomData string `json:"clientCustomData"`
}

type HealthCheckRequest struct {
	Function string                `json:"function"`
	Data     HealthCheckCustomData `json:"data"`
}

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
