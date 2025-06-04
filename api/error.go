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
	switch e.StatusCode {
	case "invalid_token":
		return e.invalidToken()
	default:
		return e.defaultMessage()
	}
}

func (e *APIError) invalidToken() string {
	return fmt.Sprintf(
		"gofactory api error: invalid token for the Satisfactory API: %s",
		e.Message,
	)
}

func (e *APIError) defaultMessage() string {
	if e.Data != nil {
		dataStr, err := e.marshalErrorData()
		if err != nil {
			return fmt.Sprintf("gofactory error | cannot marshal errData: %s",
				err)
		}
		return fmt.Sprintf(
			"gofactory api error | status code: %s | message: %s\ndata: %s\n",
			e.StatusCode, e.Message, dataStr,
		)
	}
	return fmt.Sprintf("gofactory api error | status code: %s | message: %s\n",
		e.StatusCode, e.Message)
}

// marshalErrorData marshals the error data to a JSON string.
func (e *APIError) marshalErrorData() (string, error) {
	b, err := json.MarshalIndent(e.Data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
