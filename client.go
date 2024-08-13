package coinbase

import (
	"net/http"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

type Client struct {
	cfg *client.Configuration

	baseHTTPClient *http.Client
	apiKey         APIKey
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
		key, err := LoadAPIKeyFromFile(fileName)
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
	c.cfg.HTTPClient = c.baseHTTPClient
	c.cfg.HTTPClient.Transport = NewAuthedTransport(c.apiKey, c.baseHTTPClient.Transport)
	return c, nil
}
