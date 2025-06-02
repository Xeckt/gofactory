package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func (c *GoFactoryClient) createPostRequest(functionName string, apiFunction []byte) (*http.Request, error) {
	fmt.Println("Debug: ", c.url+"/?function="+functionName)
	request, err := http.NewRequest(http.MethodPost, c.url+"/?function="+functionName, bytes.NewBuffer(apiFunction))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

func (c *GoFactoryClient) sendPostRequest(request *http.Request, response APIResponse) (*APIError, error) {
	resp, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		var apiError APIError
		return nil, json.NewDecoder(resp.Body).Decode(&apiError)
	}

	fmt.Println(resp.Body)
	return nil, json.NewDecoder(resp.Body).Decode(response)
}
