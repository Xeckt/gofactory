package api

import (
	"encoding/json"
	"log"
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

func (c *GoFactoryClient) GetServerHealth(customData string) (*HealthCheckResponse, *APIError, error) {
	functionBody, err := json.Marshal(healthCheckRequest{
		Function: HealthCheckFunction,
		Data: healthCheckCustomData{
			CustomData: customData,
		},
	})

	req, err := c.createPostRequest(HealthCheckFunction, functionBody)
	if err != nil {
		return nil, nil, err
	}

	var healthCheckResponse healthCheckResponseData
	apiErr, err := c.sendPostRequest(req, &healthCheckResponse)
	if err != nil {
		log.Fatal(err)
	}

	if apiErr != nil {
		return nil, apiErr, nil
	}

	return &healthCheckResponse.Data, nil, nil
}
