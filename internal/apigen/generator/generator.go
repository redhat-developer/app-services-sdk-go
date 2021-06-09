package generator

import (
	"fmt"
	"os/exec"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/openapi"
)

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

type Options struct {
	Generator      string
	DownloadURL    string
	AccessToken    string
	ClientID       string
	TemplatesDir   string
	ClientMetadata *metadata.SdkEntry
}

const cmdName = "@openapitools/openapi-generator-cli"

func Generate(opts *Options) error {
	cfg := Config{
		ClientID:           opts.ClientID,
		Version:            "5.1.1",
		Input:              opts.DownloadURL,
		TemplatesDir:       opts.TemplatesDir,
		IgnoreFileOverride: ".openapi-generator-ignore",
		ClientMetadata:     opts.ClientMetadata,
	}
	err := exec.Command("npx", cmdName, "version-manager", "set", cfg.Version).Run()
	if err != nil {
		return err
	}
	switch opts.Generator {
	case "go":
		return generateGo(&cfg)
	}
	return nil
}

func generateGo(cfg *Config) error {
	packageName := fmt.Sprintf("%vclient", cfg.ClientMetadata.APIGroup)
	apiVersion := fmt.Sprintf("api%v", cfg.ClientMetadata.APIVersion)
	outputPath := fmt.Sprintf("%v/%v/client", cfg.ClientMetadata.APIGroup, apiVersion)
	inputFile, err := openapi.GetFileName(cfg.Input)
	if err != nil {
		return err
	}

	cmdArgs := []string{
		cmdName,
		"generate",
		"-g", "go",
		"-i", inputFile,
		"-o", outputPath,
		"-p", "generateInterfaces=true",
		"--ignore-file-override", cfg.IgnoreFileOverride,
		"--git-user-id", "redhat-developer",
		"--git-repo-id", fmt.Sprintf("%v/%v", "app-services-sdk-go", packageName),
		"--package-name", packageName,
	}
	if cfg.TemplatesDir != "" {
		cmdArgs = append(cmdArgs, "-t", cfg.TemplatesDir)
	}

	err = exec.Command("npx", cmdArgs...).Run()
	if err != nil {
		return err
	}
	return nil
}
