package main

import (
	"context"
	"fmt"
	"os"

	registryinstance "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	apiClient := registryinstance.NewAPIClient(&registryinstance.Config{
		HTTPClient: tc,
		BaseURL:    os.Getenv("API_URL"),
	})

	artifacts, _, err := apiClient.ArtifactsApi.SearchArtifacts(context.TODO()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(artifacts.GetArtifacts()))
}
