package coinbase

import (
	"fmt"
	"net/http"
)

type authedTransport struct {
	transport http.RoundTripper
	apiKey    APIKey
}

func NewAuthedTransport(apiKey APIKey, transport http.RoundTripper) http.RoundTripper {
	return &authedTransport{
		transport: transport,
		apiKey:    apiKey,
	}
}

// RoundTrip implements the http.RoundTripper interface and wraps
// the base round tripper with logic to inject the API key auth-based HTTP headers
// into the request. Reference: https://pkg.go.dev/net/http#RoundTripper
func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	jwt, err := BuildJWT(
		t.apiKey,
		"cdp_service",
		fmt.Sprintf("%s %s%s", req.Method, req.URL.Host, req.URL.Path),
	)
	if err != nil {
		return nil, fmt.Errorf("error building data for auth header: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))

	return t.transport.RoundTrip(req)
}
