package main

import (
	"context"
	"os"

	connectormgmt "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	baseURL := os.Getenv("API_URL")
	client := connectormgmt.NewAPIClient(&connectormgmt.Config{
		HTTPClient: tc,
		BaseURL:    baseURL,
	})

	_, _, err := client.ConnectorsApi.GetConnector(ctx, "id").id("kafka-id").Execute()
	if err != nil {
		panic(err)
	}
}
