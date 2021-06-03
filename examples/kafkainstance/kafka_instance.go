package main

import (
	"context"
	"fmt"
	"os"

	kafkainstanceapi "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	apiClient := kafkainstanceapi.NewAPIClient(&kafkainstanceapi.Config{
		HTTPClient: tc,
		Debug:      true,
		BaseURL:    "http://localhost:9000/custom/path/to/rest",
	})

	res, _, err := apiClient.DefaultApi.GetTopics(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, len(res.GetItems()))
}
