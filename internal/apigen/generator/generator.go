package generator

import (
	"os/exec"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/common"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/golang"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
)

type Options struct {
	Generator      string
	DownloadURL    string
	AccessToken    string
	ClientID       string
	TemplatesDir   string
	ClientMetadata *metadata.SdkEntry
}

// Generate generates an API client and mocks
func Generate(opts *Options) error {
	cfg := common.Config{
		ClientID:           opts.ClientID,
		Version:            "5.1.1",
		Input:              opts.DownloadURL,
		TemplatesDir:       opts.TemplatesDir,
		IgnoreFileOverride: ".openapi-generator-ignore",
		ClientMetadata:     opts.ClientMetadata,
	}
	err := exec.Command("npx", common.CmdName, "version-manager", "set", cfg.Version).Run()
	if err != nil {
		return err
	}
	switch opts.Generator {
	case "go":
		goGen := golang.GoGen{
			Config: &cfg,
		}
		return goGen.Generate()
	}
	return nil
}
