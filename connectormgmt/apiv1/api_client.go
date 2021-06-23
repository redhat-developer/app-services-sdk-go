package connectormgmt

import (
	"github.com/redhat-developer/app-services-sdk-go/internal"

	apiv1 "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

// APIConfig defines the available configuration options
// to customise the API client settings
type Config = internal.APIConfig

// NewAPIClient returns a new ConnectorManagement v1 API client
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

	client := apiv1.NewAPIClient(apiCfg)

	return client
}
