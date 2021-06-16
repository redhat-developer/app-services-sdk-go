package generator

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/common"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator/golang"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
)

type Options struct {
	Generator      string
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
		Input:              opts.ClientMetadata.OpenApiFile,
		TemplatesDir:       opts.TemplatesDir,
		IgnoreFileOverride: ".openapi-generator-ignore",
		ClientMetadata:     opts.ClientMetadata,
	}
	if out, err := exec.Command("npx", common.CmdName, "version-manager", "set", cfg.Version).CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		return err
	}
	switch opts.Generator {
	case "go":
		goGen := golang.GoGen{
			Config: &cfg,
		}
		if err := goGen.Generate(); err != nil {
			return err
		}
	}
	return nil
}
