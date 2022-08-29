package main

import (
	"context"
	"fmt"
	"os"

	rhoasAuth "github.com/redhat-developer/app-services-sdk-go/auth/apiv1"
	kafkainstanceapi "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1"
)

func main() {

	offlineToken := os.Getenv("OFFLINE_TOKEN")
	httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

	apiClient := kafkainstanceapi.NewAPIClient(&kafkainstanceapi.Config{
		HTTPClient: httpClient,
		Debug:      true,
		BaseURL:    os.Getenv("API_URL"),
	})

	res, _, err := apiClient.TopicsApi.GetTopics(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, len(res.GetItems()))
}
