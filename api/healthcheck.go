package api

import (
	"encoding/json"
	"log"
)

type healthCheckCustomData struct {
	CustomData string `json:"clientCustomData"`
}

type healthCheckRequest struct {
	Function string                `json:"function"`
	Data     healthCheckCustomData `json:"data"`
}

type HealthCheckResponse struct {
	Data struct {
		Health     string `json:"health"`
		CustomData string `json:"serverCustomData"`
	} `json:"data"`
}

func (c *GoFactoryClient) GetServerHealth(customData string) (*HealthCheckResponse, *APIError, error) {
	request, err := json.Marshal(healthCheckRequest{
		Function: HealthCheckFunction,
		Data: healthCheckCustomData{
			CustomData: customData,
		},
	})

	req, err := c.createPostRequest(HealthCheckFunction, request)
	if err != nil {
		return nil, nil, err
	}

	var healthCheckResponse HealthCheckResponse
	apiErr, err := c.sendPostRequest(req, &healthCheckResponse)
	if err != nil {
		log.Fatal(err)
	}

	if apiErr != nil {
		return nil, apiErr, nil
	}

	return &healthCheckResponse, nil, nil
}
