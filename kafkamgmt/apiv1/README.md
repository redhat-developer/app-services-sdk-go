# Go API client for Kafka Service Fleet Manager

Kafka Service Fleet Manager is a REST API to manage Kafka instances and connectors.

## Installation

To install the package to your project use `go get`:
```shell
go get github.com/redhat-developer/app-services-sdk-go
```

## Usage

### Importing the package

Import the `github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1` package into your code:

```go
package main

import (
    "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"
)
```

### Basic usage

If you do not need to customise any of the default configuration values you can pass `nil` to the `NewAPIClient` constructor:

```go
client := kafkamgmt.NewAPIClient(nil)

kafkas, resp, err := client.DefaultApi.GetKafkas(context.Background()).Execute()
```

### Configuration

You can override the default configuration options:

```go
baseURL, _ := url.Parse("http://localhost:8001")
cfg := kafkamgmt.Config{
    HTTPClient: yourHTTPClient,
    Debug: true,
    ServerURL: baseURL,
}

client := kafkamgmt.NewAPIClient(&cfg)
```

## Authentication

The SDK does not directly handle authentication. When creating a new client, pass a [`http.Client`](https://golang.org/pkg/net/http/#Client) that can handle authentication for you.

```go
ctx := context.Background()
ts := oauth2.StaticTokenSource(
    &oauth2.Token{
        AccessToken: ".. your access token ..",
    },
)

tc := oauth2.NewClient(ctx, ts)

client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
    HTTPClient: tc,
})

kafkas, _, err := client.DefaultApi.GetKafkas(ctx).Execute()
```