package main

import (
	"context"
	"fmt"
	"os"

	serviceregistrymgmt "github.com/redhat-developer/app-services-sdk-go/serviceregistrymgmt/apiv1"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	apiClient := serviceregistrymgmt.NewAPIClient(&serviceregistrymgmt.Config{
		HTTPClient: tc,
	})

	registries, _, err := apiClient.DefaultApi.GetRegistries(context.TODO()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(registries))
}
