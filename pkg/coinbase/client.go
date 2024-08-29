package coinbase

import (
	"net/http"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/auth"
)

// Client is a coinbase client
type Client struct {
	cfg    *client.Configuration
	client *client.APIClient

	baseHTTPClient *http.Client
	apiKey         auth.APIKey
}

// ClientOption is a functional option for the client
type ClientOption func(*Client) error

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.cfg.Servers[0].URL = baseURL
		return nil
	}
}

// WithAPIKeyFromJSON sets the API key for the client loaded
// from a supplied json file
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

// WithAPIKey sets the explicit API key for the client loaded
func WithAPIKey(name, privateKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = auth.APIKey{
			Name:       name,
			PrivateKey: privateKey,
		}
		return nil
	}
}

// WithHTTPClient sets the http client for the client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.baseHTTPClient = httpClient
		return nil
	}
}

// WithDebug sets the debug flag for the client
func WithDebug() ClientOption {
	return func(c *Client) error {
		c.cfg.Debug = true
		return nil
	}
}

// NewClient creates a new coinbase client with the supplied options.
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
		c.baseHTTPClient = &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
			},
		}
	}

	c.cfg.HTTPClient = c.baseHTTPClient

	if c.cfg.HTTPClient.Transport == nil {
		c.cfg.HTTPClient.Transport = http.DefaultTransport
	}

	c.cfg.HTTPClient.Transport = auth.NewTransport(c.apiKey, c.cfg.HTTPClient.Transport)

	c.client = client.NewAPIClient(c.cfg)
	return c, nil
}
