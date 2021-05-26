package kafkamgmt

import (
	"github.com/redhat-developer/app-services-sdk-go/common"

	apiv1 "github.com/redhat-developer/app-services-sdk-go/kafkaadmin/apiv1internal"
)

// APIConfig defines the available configuration options
// to customise the API client settings
type APIConfig = common.APIConfig

// NewAPIClient returns a new KafkaManagement v1 API client
// using a custom config
func NewAPIClient(cfg *APIConfig) apiv1.DefaultApi {
	apiCfg := apiv1.NewConfiguration()
	if cfg == nil {
		client := apiv1.NewAPIClient(apiCfg)
		return client.DefaultApi
	}

	if cfg.HTTPClient == nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.ServerURL != nil {
		apiCfg.Host = cfg.ServerURL.Host
		apiCfg.Scheme = cfg.ServerURL.Scheme
	}

	apiCfg.Debug = cfg.Debug

	client := apiv1.NewAPIClient(apiCfg)

	return client.DefaultApi
}

// NewAPIClientWithDefaults returns a new KafkaManagement v1 API client
// using the default configuration values
func NewAPIClientWithDefaults(config *APIConfig) apiv1.DefaultApi {
	return NewAPIClient(nil)
}
