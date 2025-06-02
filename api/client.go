package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

const Version = "0.1.0"

type GoFactoryClient struct {
	url    string
	token  string
	client *http.Client
}

type APIResponse interface{}

func NewGoFactoryClient(url string, token string) *GoFactoryClient {
	return &GoFactoryClient{
		url:   url,
		token: token,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (c *GoFactoryClient) VerifyToken() (*APIError, error) {
	request, err := c.CreatePostRequest(VerifyAuthTokenFunction, createGenericFunctionBody(VerifyAuthTokenFunction))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil, nil
	}

	var apiError APIError
	err = json.NewDecoder(resp.Body).Decode(&apiError)
	if err != nil {
		return nil, err
	}

	return &apiError, nil
}

func (c *GoFactoryClient) CreatePostRequest(functionName string, apiFunction []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.url+"/api/v1/?function="+functionName, bytes.NewBuffer(apiFunction))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

func (c *GoFactoryClient) SendPostRequest(request *http.Request, response APIResponse) (*APIError, error) {
	resp, err := c.client.Do(request)
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

func CreateAndSendPostRequest[Resp any](c *GoFactoryClient, functionName string, apiFunction []byte) (*Resp, error) {
	request, err := c.CreatePostRequest(functionName, apiFunction)
	if err != nil {
		return nil, err
	}
	var resp Resp
	apiError, err := c.SendPostRequest(request, &resp)
	if err != nil {
		return nil, err
	}
	if apiError != nil {
		return nil, apiError
	}
	return &resp, nil
}
