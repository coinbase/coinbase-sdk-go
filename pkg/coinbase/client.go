package coinbase

import (
	"net/http"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/auth"
)

type Client struct {
	cfg    *client.Configuration
	client *client.APIClient

	baseHTTPClient *http.Client
	apiKey         auth.APIKey
}

type ClientOption func(*Client) error

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.cfg.Servers[0].URL = baseURL
		return nil
	}
}

func WithAPIKeyFromJSON(fileName string) ClientOption {
	return func(c *Client) error {
		key, err := auth.LoadAPIKeyFromFile(fileName)
		if err != nil {
			return err
		}

		c.apiKey = key
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.baseHTTPClient = httpClient
		return nil
	}
}

func NewClient(o ...ClientOption) (*Client, error) {
	c := &Client{
		cfg: client.NewConfiguration(),
	}
	for _, opt := range o {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	if c.baseHTTPClient == nil {
		c.baseHTTPClient = http.DefaultClient
	}
	c.cfg.HTTPClient = c.baseHTTPClient
	if c.cfg.HTTPClient.Transport == nil {
		c.cfg.HTTPClient.Transport = http.DefaultTransport
	}
	c.cfg.HTTPClient.Transport = auth.NewTransport(c.apiKey, c.cfg.HTTPClient.Transport)
	c.client = client.NewAPIClient(c.cfg)
	return c, nil
}
