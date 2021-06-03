package main

import (
	"context"
	"fmt"
	"os"

	kafkainstance "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	apiClient := kafkainstance.NewAPIClient(&kafkainstance.Config{
		HTTPClient: tc,
	})

	res, _, err := apiClient.DefaultApi.GetTopics(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, len(res.GetItems()))
}
