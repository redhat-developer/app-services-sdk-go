package kafkamgmt

import (
	"github.com/redhat-developer/app-services-sdk-go/internal"

	"github.com/redhat-developer/app-services-sdk-go/kafkaadmin/apiv1internal/client"
)

// Config defines the available configuration options
// to customise the API client settings
type Config = internal.APIConfig

// NewAPIClient returns a new KafkaAdmin v1 API client
// using a custom config
func NewAPIClient(cfg *Config) *kafkaadmin.APIClient {
	apiCfg := kafkaadmin.NewConfiguration()
	if cfg == nil {
		return kafkaadmin.NewAPIClient(apiCfg)
	}

	if cfg.HTTPClient == nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.ServerURL != nil {
		apiCfg.Host = cfg.ServerURL.Host
		apiCfg.Scheme = cfg.ServerURL.Scheme
	}

	apiCfg.Debug = cfg.Debug

	client := kafkaadmin.NewAPIClient(apiCfg)
	return client
}