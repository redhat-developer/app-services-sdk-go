package main

import (
	"github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal"
	"context"
	"fmt"
	"os"
)

func main() {
	apiClient := kafkainstance.NewAPIClient(&kafkainstance.Config{
		Debug: true,
	})

	res, _, err := apiClient.DefaultApi.GetTopics(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, len(res.GetItems()))
}
