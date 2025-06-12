package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

// Version is the current version of the GoFactory API client.
const Version = "1.0.0"

// GoFactoryClient is a client for interacting with the Satisfactory dedicated server API.
type GoFactoryClient struct {
	// URL is the base URL (& port, if necessary) of the Satisfactory dedicated server API.
	URL string

	// Token is the authentication token used for API requests.
	Token string

	// currentPrivilege represents the current privilege level of the client.
	currentPrivilege string

	// Client is the underlying HTTP client used for API requests.
	Client *http.Client
}

// ApiResponse is an empty interface used as a placeholder
// for API responses returned by the Satisfactory dedicated server.
type ApiResponse any

// NewGoFactoryClient creates a new GoFactoryClient with the specified URL,
// authentication token, and an option to skip TLS verification.
func NewGoFactoryClient(url string, token string, skipVerify bool) *GoFactoryClient {
	return &GoFactoryClient{
		URL:   url,
		Token: token,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
			},
		},
	}
}

// CreatePostRequest creates a HTTP POST request to call the specified API function.
func (c *GoFactoryClient) CreatePostRequest(functionName string, apiFunction []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.URL+"/api/v1/?function="+functionName, bytes.NewBuffer(apiFunction))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.Token)
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

// CreatePostRequestWithHeaders creates a HTTP POST request to call the specified API function,
// using a map of custom headers.
func (c *GoFactoryClient) CreatePostRequestWithHeaders(headers map[string]string, functionName string, apiFunction []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.URL+"/api/v1/?function="+functionName, bytes.NewBuffer(apiFunction))
	if err != nil {
		return nil, err
	}

	for header, headerValue := range headers {
		request.Header.Add(header, headerValue)
	}

	return request, nil
}

// SendPostRequest sends the provided HTTP request to the server and decodes the response
// into the given ApiResponse. It also handles API errors returned by the server.
func (c *GoFactoryClient) SendPostRequest(ctx context.Context, request *http.Request, response ApiResponse) error {
	resp, err := c.Client.Do(request.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if berr := resp.Body.Close(); err != nil {
			err = berr
		}
	}()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		var apiError APIError
		err := json.NewDecoder(resp.Body).Decode(&apiError)
		if err != nil {
			return err
		}
		return &apiError
	}

	if resp.StatusCode == 204 {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

// CreateAndSendPostRequest creates a HTTP POST request for the given API function,
// sends it, and decodes the response into the generic type Resp.
func CreateAndSendPostRequest[Resp any](ctx context.Context, c *GoFactoryClient, functionName string, apiFunction []byte) (*Resp, error) {
	request, err := c.CreatePostRequest(functionName, apiFunction)
	if err != nil {
		return nil, err
	}
	var resp Resp
	err = c.SendPostRequest(ctx, request, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAndSendPostRequestWithHeaders creates a HTTP POST request with custom headers
// for the given API function, sends it, and decodes the response into the generic type Resp.
func CreateAndSendPostRequestWithHeaders[Resp any](ctx context.Context, c *GoFactoryClient, headers map[string]string, functionName string, apiFunction []byte) (*Resp, error) {
	request, err := c.CreatePostRequestWithHeaders(headers, functionName, apiFunction)
	if err != nil {
		return nil, err
	}
	var resp Resp
	err = c.SendPostRequest(ctx, request, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
