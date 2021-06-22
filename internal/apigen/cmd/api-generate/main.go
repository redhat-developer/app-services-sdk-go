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
var clientID string
var clientPayloadStr string
var generatorInput string
var templatesDir string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&clientID, "client-id", "", "Client ID")
	flag.StringVar(&generatorInput, "generator", "", "Language: go, java, ts")
	flag.StringVar(&templatesDir, "templates-dir", "", "Templates directory")
}

func main() {
	flag.Parse()
	if repoMetadataPath == "" {
		log.Fatalln("Missing required flag: --repo-metadata")
	}
	if clientID == "" {
		log.Fatalln("Missing required flag: --client-id")
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
	clientConfig, err := metadata.GetSdkEntry(clientID, repoMetadataPath)
	if err != nil && generatorInput == "go" {
		log.Fatalln("no config found for client:", clientID)
	}
	if clientConfig.Disabled {
		fmt.Fprintf(os.Stderr, "Skipping generation of '%v': generation disabled\n", clientID)
		return
	}

	fullTemplatesDirPath := path.Join(root, templatesDir)

	genOpts := generator.Options{
		Generator:      generatorInput,
		ClientID:       clientID,
		AccessToken:    accessToken,
		ClientMetadata: clientConfig,
		TemplatesDir:   fullTemplatesDirPath,
	}
	err = generator.Generate(&genOpts)
	if err != nil {
		log.Fatalln(err)
	}
}
