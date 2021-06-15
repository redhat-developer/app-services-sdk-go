package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/redhat-developer/app-services-sdk-go/internal/apigen/openapi"
)

var accessToken string
var downloadURL string
var downloadLocation string

func init() {
	flag.StringVar(&accessToken, "token", "", "Access token")
	flag.StringVar(&downloadURL, "download-url", "", "Download URL")
	flag.StringVar(&downloadLocation, "download-location", ".", "The location to download the file to")
}

func main() {
	flag.Parse()
	if downloadURL == "" {
		log.Fatalln("Missing required flag: --download-url")
	}

	fileName, err := openapi.GetFileName(downloadURL)
	if err != nil {
		log.Fatalln(err)
	}

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	req, err := http.NewRequest("GET", downloadURL, nil)
	req.Header.Set("Authorization", "token "+accessToken)
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalln("could not download", downloadURL, "-", resp.Status)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	os.Rename(fileName, path.Join(downloadLocation, fileName))

	defer file.Close()
}
