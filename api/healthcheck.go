package api

import (
	"encoding/json"
)

type HealthCheckResponse struct {
	Health     string `json:"health"`
	CustomData string `json:"serverCustomData"`
}

type healthCheckResponseData struct {
	Data HealthCheckResponse `json:"data"`
}

type healthCheckCustomData struct {
	CustomData string `json:"clientCustomData"`
}

type healthCheckRequest struct {
	Function string                `json:"function"`
	Data     healthCheckCustomData `json:"data"`
}

func (c *GoFactoryClient) GetServerHealth(customData string) (*HealthCheckResponse, error) {
	functionBody, err := json.Marshal(healthCheckRequest{
		Function: HealthCheckFunction,
		Data: healthCheckCustomData{
			CustomData: customData,
		},
	})
	if err != nil {
		return nil, err
	}

	healthCheckResponse, err := CreateAndSendPostRequest[healthCheckResponseData](c, HealthCheckFunction, functionBody)
	if err != nil {
		return nil, err
	}
	return &healthCheckResponse.Data, nil
}
