# Go API client for SSO service accounts API

Service Accounts API allows you to manage credentails for accessing instance SDKs.

## Installation

To install the package to your project use `go get`:

```shell
go get github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt
```

## Usage

### Importing the package

Import the `github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt/apiv1` package into your code:

```go
package main

import (
    "github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt/apiv1"
)
```

### Basic usage

If you do not need to customise any of the default configuration values you can pass `nil` to the `NewAPIClient` constructor:

```go
client := serviceAccounts.NewAPIClient(nil)

serviceaccountsList, resp, err := client.DefaultApi.GetServiceAccounts(context.Background()).Execute()
```

### Configuration

You can override the default configuration options:

```go
cfg := serviceaccounts.Config{
    HTTPClient: yourHTTPClient,
    Debug: true,
    BaseURL: "http://localhost:8080",
}

client := serviceaccounts.NewAPIClient(&cfg)
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

client := serviceaccounts.NewAPIClient(&connectormgmt.Config{
    HTTPClient: tc,
})
```

## Error handling

> NOTE: SDK use different type of error handling than other SDKs.

## Endpoints

For a full list of the available endpoints, see [Documentation for API Endpoints](./client/README.md#documentation-for-api-endpoints).