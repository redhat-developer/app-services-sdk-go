package internal

import (
	"net/http"
	"net/url"
)

// APIConfig defines the available configuration options
// to customise the API client settings
type APIConfig struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug      bool
	// ServerURL sets a custom API server base URL
	ServerURL  *url.URL
}