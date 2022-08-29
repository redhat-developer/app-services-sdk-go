package main

import (
	"context"
	"os"

	rhoasAuth "github.com/redhat-developer/app-services-sdk-go/auth/apiv1"
	connectormgmt "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1"
)

func main() {
	offlineToken := os.Getenv("OFFLINE_TOKEN")
	httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

	ctx := context.Background()

	baseURL := os.Getenv("API_URL")
	client := connectormgmt.NewAPIClient(&connectormgmt.Config{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
	})

	_, _, err := client.ConnectorsApi.GetConnector(ctx, "id").Execute()

	if err != nil {
		panic(err)
	}
}
