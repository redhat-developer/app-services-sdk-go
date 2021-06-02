package kafkainstance

import (
	"github.com/redhat-developer/app-services-sdk-go/internal"

	apiv1 "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

// APIConfig defines the available configuration options
// to customise the API client settings
type Config = internal.APIConfig

// NewAPIClient returns a new KafkaManagement v1 API client
// using a custom config
func NewAPIClient(cfg *Config) *apiv1.APIClient {
	apiCfg := apiv1.NewConfiguration()
	if cfg == nil {
		return apiv1.NewAPIClient(apiCfg)
	}

	if cfg.HTTPClient != nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.ServerURL != nil {
		apiCfg.Host = cfg.ServerURL.Host
		apiCfg.Scheme = cfg.ServerURL.Scheme
	}

	apiCfg.Debug = cfg.Debug

	client := apiv1.NewAPIClient(apiCfg)

	return client
}
