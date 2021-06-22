package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
)

var repoMetadataPath string
var accessToken string
var clientPayloadStr string
var generatorInput string
var templatesDir string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&generatorInput, "generator", "", "Language: go, java, ts")
	flag.StringVar(&templatesDir, "templates-dir", "", "Templates directory")
}

func main() {
	flag.Parse()
	if repoMetadataPath == "" {
		log.Fatalln("Missing required flag: --repo-metadata")
	}
	if generatorInput == "" {
		log.Fatalln("Missing required flag: --generator")
	}
	if generatorInput != "go" && generatorInput != "java" && generatorInput != "js" {
		log.Fatalln("Invalid valuer for --generator:", generatorInput)
	}

	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	clientConfig, err := metadata.GetClientMetadata(repoMetadataPath)
	if err != nil && generatorInput == "go" {
		log.Fatalln("no metadata found:", err)
	}

	fullTemplatesDirPath := path.Join(root, templatesDir)

	for id, metadata := range *clientConfig {
		if metadata.Disabled {
			fmt.Fprintf(os.Stderr, "Skipping generation of '%v': generation disabled\n", id)
			continue
		}
		genOpts := generator.Options{
			Generator:      generatorInput,
			ClientID:       id,
			AccessToken:    accessToken,
			ClientMetadata: &metadata,
			TemplatesDir:   fullTemplatesDirPath,
		}
		err = generator.Generate(&genOpts)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
