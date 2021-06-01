package main

import (
	"context"
	"fmt"
	"os"

	kafkamgmt "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"
)

func main() {
	apiClient := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		Debug: true,
	})

	res, _, err := apiClient.DefaultApi.GetKafkas(context.Background()).Execute()
	if err != nil {

	}
	fmt.Fprintln(os.Stderr, res.Size)
}
