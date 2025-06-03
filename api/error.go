package api

import "fmt"

type APIError struct {
	StatusCode string      `json:"errorCode"`
	Message    string      `json:"errorMessage"`
	Data       interface{} `json:"errorData,omitempty"`
}

func (e *APIError) Error() string {
	if e.StatusCode == "invalid_token" {
		return fmt.Sprintf("Invalid token for the Satisfactory API: %s", e.Message)
	}
	return fmt.Sprintf("%s: %s", e.StatusCode, e.Message)
}
