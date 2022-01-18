package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/MakeNowJust/heredoc"
	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/metadata"
)

var repoMetadataPath string
var repoErrorsPath string
var accessToken string
var clientPayloadStr string

func init() {
	flag.StringVar(&repoMetadataPath, "repo-metadata", "", "path to repo metadata JSON")
	flag.StringVar(&repoErrorsPath, "errors-folder", "", "path to repo errors JSON files")
}

func main() {
	flag.Parse()
	if repoMetadataPath == "" {
		log.Fatalln("Missing required flag: --repo-metadata")
	}
	if repoErrorsPath == "" {
		log.Fatalln("Missing required flag: --errors-folder")
	}

	clientConfig, err := metadata.GetClientMetadata(repoMetadataPath)
	if err != nil {
		log.Fatalln("no metadata found:", err)
	}

	for id, metadata := range *clientConfig {
		if metadata.Disabled {
			fmt.Fprintf(os.Stderr, "Skipping generation of '%v': generation disabled\n", id)
			continue
		}

		templateString := heredoc.Doc(`
		package {{.Package}}
		
		// {{.Name}} represents error code returned by Kafka Management API
		type ServiceErrorCode string
		  
		const (
			{{range .Errors}}
			// {{.Reason}}
			ERROR_{{.ErrorID}} ServiceErrorCode = "{{.ErrorCode}}"
			{{end}}
		)
		`)
		data := struct {
			Package string
			Errors []interface
		}{
			Package: metadata.APIGroup,
		}
		errorsSDKFileTemplate, _ := template.New("errorsSDK").Parse(templateString)
		if err != nil {
			log.Fatalln(err)
		}

		file, err := os.Create(metadata.ErrorsTarget))
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		errorsSDKFileTemplate.Execute(file, data)
)

		fmt.Fprintf(os.Stderr, "Generated %v errors SDK ", metadata.APIGroup)

	}
}
