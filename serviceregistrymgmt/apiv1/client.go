package serviceregistrymgmt

import (
	"github.com/redhat-developer/app-services-sdk-go/internal"

	apiClient "github.com/redhat-developer/app-services-sdk-go/serviceregistrymgmt/apiv1/client"
)

// Config defines the available configuration options
// to customise the API client settings
type Config = internal.APIConfig

// NewAPIClient returns a new KafkaManagement v1 API client
// using a custom config
func NewAPIClient(cfg *Config) *apiClient.APIClient {
	apiCfg := apiClient.NewConfiguration()
	if cfg == nil {
		return apiClient.NewAPIClient(apiCfg)
	}

	if cfg.HTTPClient != nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.BaseURL != "" {
		apiCfg.Servers = []apiClient.ServerConfiguration{
			{
				URL: cfg.BaseURL,
			},
		}
	}

	apiCfg.Debug = cfg.Debug

	client := apiClient.NewAPIClient(apiCfg)

	return client
}
