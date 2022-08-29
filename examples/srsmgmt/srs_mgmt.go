package main

import (
	"context"
	"fmt"
	"os"

	rhoasAuth "github.com/redhat-developer/app-services-sdk-go/auth/apiv1"
	registrymgmt "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1"
)

func main() {
	offlineToken := os.Getenv("OFFLINE_TOKEN")
	httpClient := rhoasAuth.BuildAuthenticatedHTTPClient(offlineToken)

	apiClient := registrymgmt.NewAPIClient(&registrymgmt.Config{
		HTTPClient: httpClient,
	})

	registries, _, err := apiClient.RegistriesApi.GetRegistries(context.TODO()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(registries.GetItems()))
}
