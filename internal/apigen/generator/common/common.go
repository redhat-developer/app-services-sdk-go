package common

import "github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"

type Config struct {
	ClientID             string
	Version              string
	Generator            string
	Input                string
	TemplatesDir         string
	AdditionalProperties string
	IgnoreFileOverride   string
	ClientMetadata       *metadata.SdkEntry
}

type Generator interface {
	Generate(cfg *Config) error
}

const CmdName = "@openapitools/openapi-generator-cli"
