package coinbase

import (
	"net/http"
	"os"
	"testing"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/stretchr/testify/assert"
)

func TestWithBaseURL(t *testing.T) {
	tests := []struct {
		name         string
		baseURL      string
		expectations func(t *testing.T, c *Client, err error)
	}{
		{
			name:    "should eq test_url",
			baseURL: "test_url",
			expectations: func(t *testing.T, c *Client, err error) {
				assert.Equal(t, "test_url", c.cfg.Servers[0].URL)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{cfg: client.NewConfiguration()}
			err := WithBaseURL(tt.baseURL)(c)
			tt.expectations(t, c, err)
		})
	}
}

func TestWithAPIKeyFromJSON(t *testing.T) {
	tests := []struct {
		name         string
		contents     string
		expectations func(t *testing.T, c *Client, err error)
	}{
		{
			name:     "with a valid json file",
			contents: `{"privateKey": "test_key", "name": "some-name"}`,
			expectations: func(t *testing.T, c *Client, err error) {
				assert.Equal(t, "test_key", c.apiKey.PrivateKey)
				assert.Equal(t, "some-name", c.apiKey.Name)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{cfg: client.NewConfiguration()}
			f := createTempJSONFile(t, tt.contents)
			defer f.Close()
			err := WithAPIKeyFromJSON(f.filePath)(c)
			tt.expectations(t, c, err)
		})
	}
}

func TestWithHTTPClient(t *testing.T) {
	testClient := &http.Client{}
	tests := []struct {
		name         string
		httpClient   *http.Client
		expectations func(t *testing.T, c *Client, err error)
	}{
		{
			name:       "should eq test_http_client",
			httpClient: testClient,
			expectations: func(t *testing.T, c *Client, err error) {
				assert.Equal(t, testClient, c.baseHTTPClient)
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{cfg: client.NewConfiguration()}
			err := WithHTTPClient(tt.httpClient)(c)
			tt.expectations(t, c, err)
		})
	}
}

func TestWithDebug(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{name: "should eq true", expected: true},
		{name: "should eq false", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{cfg: client.NewConfiguration()}
			if tt.expected {
				err := WithDebug()(c)
				assert.NoError(t, err)
				assert.True(t, c.cfg.Debug)
			} else {
				assert.False(t, c.cfg.Debug)
			}
		})
	}
}

type f struct {
	t        *testing.T
	filePath string
}

func (f *f) Close() {
	f.t.Helper()
	os.Remove(f.filePath)
}

func createTempJSONFile(t *testing.T, content string) *f {
	t.Helper()
	file, err := os.CreateTemp("", "coinbase_api_key_test_*.json")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer file.Close()

	filePath := file.Name()
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}
	return &f{t: t, filePath: filePath}
}
