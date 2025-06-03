package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

const Version = "0.2.0"

type GoFactoryClient struct {
	URL              string
	Token            string
	CurrentPrivilege string
	Client           *http.Client
}

type ApiResponse interface{}

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

func (c *GoFactoryClient) CreatePostRequest(functionName string, apiFunction []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.URL+"/api/v1/?function="+functionName, bytes.NewBuffer(apiFunction))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.Token)
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

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

func (c *GoFactoryClient) SendPostRequest(ctx context.Context, request *http.Request, response ApiResponse) (*APIError, error) {
	resp, err := c.Client.Do(request.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		var apiError APIError
		err := json.NewDecoder(resp.Body).Decode(&apiError)
		if err != nil {
			return nil, err
		}
		return &apiError, nil
	}

	return nil, json.NewDecoder(resp.Body).Decode(response)
}

func CreateAndSendPostRequest[Resp any](ctx context.Context, c *GoFactoryClient, functionName string, apiFunction []byte) (*Resp, error) {
	request, err := c.CreatePostRequest(functionName, apiFunction)
	if err != nil {
		return nil, err
	}
	var resp Resp
	apiError, err := c.SendPostRequest(ctx, request, &resp)
	if err != nil {
		return nil, err
	}
	if apiError != nil {
		return nil, apiError
	}
	return &resp, nil
}

func CreateAndSendPostRequestWithHeaders[Resp any](ctx context.Context, c *GoFactoryClient, headers map[string]string, functionName string, apiFunction []byte) (*Resp, error) {
	request, err := c.CreatePostRequestWithHeaders(headers, functionName, apiFunction)
	if err != nil {
		return nil, err
	}
	var resp Resp
	apiError, err := c.SendPostRequest(ctx, request, &resp)
	if err != nil {
		return nil, err
	}
	if apiError != nil {
		return nil, apiError
	}
	return &resp, nil
}
