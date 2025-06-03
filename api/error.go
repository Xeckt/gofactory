package api

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	StatusCode string      `json:"errorCode"`
	Message    string      `json:"errorMessage"`
	Data       interface{} `json:"errorData,omitempty"`
}

func (e *APIError) Error() string {
	if e.StatusCode == "invalid_token" {
		return fmt.Sprintf("API Error - Invalid Token for the Satisfactory API: %s", e.Message)
	}
	if e.Data != nil {
		b, err := json.MarshalIndent(e.Data, "", "  ")
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("API error - Status code: %s\nAPI error - Message: %s\nData: %s\n", string(b))
	}
	return fmt.Sprintf("API error - Status code: %s\nAPI error - Message: %s", e.StatusCode, e.Message)
}
