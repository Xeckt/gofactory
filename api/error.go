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
		return fmt.Sprintf("gofactory error: invalid token for the satisfactory api: %s", e.Message)
	}
	if e.Data != nil {
		b, err := json.MarshalIndent(e.Data, "", "  ")
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("gofactory error: status code: %s\napi error - message: %s\ndata: %s\n",
			e.StatusCode, e.Message, string(b))
	}
	return fmt.Sprintf("gofactory error: status code: %s\napi error - message: %s\n", e.StatusCode, e.Message)
}
