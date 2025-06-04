package auth

import (
	"fmt"
	"net/http"
)

type transport struct {
	transport http.RoundTripper
	apiKey    APIKey
}

func NewTransport(apiKey APIKey, t http.RoundTripper) http.RoundTripper {
	return &transport{
		transport: t,
		apiKey:    apiKey,
	}
}

// RoundTrip implements the http.RoundTripper interface and wraps
// the base round tripper with logic to inject the API key auth-based HTTP headers
// into the request. Reference: https://pkg.go.dev/net/http#RoundTripper
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	jwt, err := BuildJWT(
		t.apiKey,
		"cdp_service",
		fmt.Sprintf("%s %s%s", req.Method, req.URL.Host, req.URL.Path),
	)
	if err != nil {
		return nil, fmt.Errorf("error building data for auth header: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))

	req.Header.Set(
		"Correlation-Context",
		fmt.Sprintf(
			"%s,%s",
			fmt.Sprintf("%s=%s", "sdk_version", "0.0.28"),
			fmt.Sprintf("%s=%s", "sdk_language", "go"),
		),
	)

	return t.transport.RoundTrip(req)
}
