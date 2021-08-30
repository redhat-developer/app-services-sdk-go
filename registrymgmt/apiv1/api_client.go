package registrymgmt

import (
	"net/http"

	apiv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

// Config defines the available configuration options
// to customise the API client settings
type Config struct {
	// HTTPClient is a custom HTTP client
	HTTPClient *http.Client
	// Debug enables debug-level logging
	Debug bool
	// BaseURL sets a custom API server base URL
	BaseURL string
	// UserAgent sets a custom user-agent
	UserAgent string
}

// NewAPIClient returns a new v1 API client
// using a custom config
func NewAPIClient(cfg *Config) *apiv1.APIClient {
	apiCfg := apiv1.NewConfiguration()
	if cfg == nil {
		return apiv1.NewAPIClient(apiCfg)
	}

	if cfg.HTTPClient != nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.BaseURL != "" {
		apiCfg.Servers = []apiv1.ServerConfiguration{
			{
				URL: cfg.BaseURL,
			},
		}
	}

	apiCfg.Debug = cfg.Debug
	apiCfg.UserAgent = cfg.UserAgent

	client := apiv1.NewAPIClient(apiCfg)

	return client
}
