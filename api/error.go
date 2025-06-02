package api

import "fmt"

type APIError struct {
	StatusCode string `json:"errorCode"`
	Message    string `json:"errorMessage"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.StatusCode, e.Message)
}
