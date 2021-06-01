package main

import (
	"context"
	"fmt"
	"os"

	serviceregistrymgmt "github.com/redhat-developer/app-services-sdk-go/serviceregistrymgmt/apiv1"
)

func main() {
	apiClient := serviceregistrymgmt.NewAPIClient(&serviceregistrymgmt.Config{
		Debug: true,
	})

	registries, _, err := apiClient.DefaultApi.GetRegistries(context.TODO()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, len(registries))
}
