package api

import (
	"crypto/tls"
	"net/http"
)

type GoFactoryClient struct {
	url    string
	token  string
	client *http.Client
}

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

func (c *GoFactoryClient) createPostRequest(function string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.url+"/?function="+function, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}
