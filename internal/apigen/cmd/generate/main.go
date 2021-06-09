package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/generator"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
)

var repoMetadataPath string
var accessToken string
var clientID string
var downloadURL string
var clientPayloadStr string
var generatorInput string
var templatesDir string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&clientID, "client-id", "", "Client ID")
	flag.StringVar(&downloadURL, "download-url", "", "Download URL")
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
	if downloadURL == "" {
		log.Fatalln("Missing required flag: --download-url")
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
	fullMetadataPath := path.Join(root, repoMetadataPath)
	f, err := ioutil.ReadFile(fullMetadataPath)
	if err != nil {
		log.Fatalln(err)
	}
	var repoMetadata metadata.RepoMetadata
	err = json.Unmarshal([]byte(f), &repoMetadata)
	if err != nil {
		log.Fatalln(err)
	}

	clientConfig, ok := repoMetadata[clientID]
	if !ok && generatorInput == "go" {
		log.Fatalln("no config found for client:", clientID)
	}

	fullTemplatesDirPath := path.Join(root, templatesDir)

	genOpts := generator.Options{
		Generator:      generatorInput,
		DownloadURL:    downloadURL,
		ClientID:       clientID,
		AccessToken:    accessToken,
		ClientMetadata: &clientConfig,
		TemplatesDir:   fullTemplatesDirPath,
	}
	err = generator.Generate(&genOpts)
	if err != nil {
		log.Fatalln(err)
	}
}
