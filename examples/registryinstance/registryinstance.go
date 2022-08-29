package main

import (
	"context"
	"fmt"
	"os"

	rhoasAuth "github.com/redhat-developer/app-services-sdk-go/auth/apiv1"
	registryinstance "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal"
)

func main() {

	offlineToken := os.Getenv("OFFLINE_TOKEN")
	httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

	apiClient := registryinstance.NewAPIClient(&registryinstance.Config{
		HTTPClient: httpClient,
		BaseURL:    os.Getenv("API_URL"),
	})

	artifacts, _, err := apiClient.ArtifactsApi.ListArtifactsInGroup(context.TODO(), "default").Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(artifacts.GetArtifacts()))
}
