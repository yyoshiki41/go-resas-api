package resas

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	Endpoint = "https://opendata.resas-portal.go.jp"
	APIV1    = "/api/v1"
)

var (
	httpClient = &http.Client{}
)

// Config is a setting for RESAS API.
type Config struct {
	APIKey   string
	Endpoint string
}

// Client represents an API client for RESAS API.
type Client struct {
	httpClient *http.Client
	config     *Config
}

// NewClient returns a new pushwoosh API client.
func NewClient(apiKey string) (*Client, error) {
	return NewClientWithConfig(
		&Config{
			APIKey:   apiKey,
			Endpoint: Endpoint,
		})
}

func NewClientWithConfig(config *Config) (*Client, error) {
	if httpClient == nil {
		return nil, errors.New("httpClient is nil")
	}
	if config == nil {
		return nil, errors.New("config is nil")
	}

	return &Client{
		httpClient: httpClient,
		config:     config,
	}, nil
}

func (c *Client) call(ctx context.Context, method string, apiPath string, urlParams url.Values, res interface{}) error {
	u, err := url.Parse(c.config.Endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, apiPath)
	if urlParams != nil {
		u.RawQuery = urlParams.Encode()
	}

	fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-KEY", c.config.APIKey)
	req = req.WithContext(ctx)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if res == nil {
		return nil
	}
	return json.NewDecoder(response.Body).Decode(&res)
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}
