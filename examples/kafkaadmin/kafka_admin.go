package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-sdk-go/kafkaadmin/apiv1internal"
)

func main() {
	apiClient := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		Debug: true,
	})

	res, _, err := apiClient.DefaultApi.GetTopics(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, len(res.GetItems()))
}
