package main

import (
	"context"
	"fmt"
	"os"

	srsmgmt "github.com/redhat-developer/app-services-sdk-go/srsmgmt/apiv1"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	apiClient := srsmgmt.NewAPIClient(&srsmgmt.Config{
		HTTPClient: tc,
	})

	registries, _, err := apiClient.RegistriesApi.GetRegistries(context.TODO()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(registries.GetItems()))
}
