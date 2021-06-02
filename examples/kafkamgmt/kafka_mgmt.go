package main

import (
	"context"
	"os"

	kafkamgmt "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: os.Getenv("ACCESS_TOKEN"),
		},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		HTTPClient: tc,
	})

	_, _, err := client.DefaultApi.GetKafkas(ctx).Execute()
	if err != nil {
		panic(err)
	}
}
