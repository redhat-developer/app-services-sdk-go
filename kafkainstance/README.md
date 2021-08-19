# Go API client for kafkainstance

An API to provide REST endpoints for query Kafka for admin operations

## Installation

To install the package to your project use `go get`:

```shell
go get github.com/redhat-developer/app-services-sdk-go/kafkainstance
```

## Usage

### Importing the package

Import the `github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal` package into your code:

```go
package main

import (
    "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal"
)
```

### Basic usage

If you do not need to customise any of the default configuration values you can pass `nil` to the `NewAPIClient` constructor:

```go
client := kafkainstance.NewAPIClient(nil)

kafkas, resp, err := client.DefaultApi.GetTopics(context.Background()).Execute()
```

### Configuration

You can override the default configuration options:

```go
cfg := kafkainstance.Config{
    HTTPClient: yourHTTPClient,
    Debug: true,
    BaseURL: "http://localhost:8001/rest",
}

client := kafkainstance.NewAPIClient(&cfg)
```

## Authentication

The SDK does not directly handle authentication. When creating a new client, pass a [`http.Client`](https://golang.org/pkg/net/http/#Client) that can handle authentication for you. The recommended way to do this is using the [`oauth2`](https://pkg.go.dev/golang.org/x/oauth2) library. If you have an OAuth2 access token you can use it with the `oauth2` library using:

```go
ctx := context.Background()
ts := oauth2.StaticTokenSource(
    &oauth2.Token{
        AccessToken: ".. your access token ..",
    },
)

tc := oauth2.NewClient(ctx, ts)

client := kafkainstance.NewAPIClient(&kafkainstance.Config{
    HTTPClient: tc,
})

topics, _, err := client.DefaultApi.GetTopics(ctx).Execute()
```

## Endpoints

For a full list of the available endpoints, see [Documentation for API Endpoints](./client/README.md#documentation-for-api-endpoints).