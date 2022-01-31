# Go API client for Service Registry Service

Managed Service Registry API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs.

## Installation

To install the package to your project use `go get`:

```shell
go get github.com/redhat-developer/app-services-sdk-go/registrymgmt
```

## Usage

### Importing the package

Import the `github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1` package into your code:

```go
package main

import (
    "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1"
)
```

### Basic usage

If you do not need to customise any of the default configuration values you can pass `nil` to the `NewAPIClient` constructor:

```go
client := registrymgmt.NewAPIClient(nil)

registries, resp, err := client.RegistriesApi.GetRegistries(context.Background()).Execute()
```

### Configuration

You can override the default configuration options:

```go
cfg := registrymgmt.Config{
    HTTPClient: yourHTTPClient,
    Debug: true,
    BaseURL: "http://localhost:8080",
}

client := registrymgmt.NewAPIClient(&cfg)
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

client := registrymgmt.NewAPIClient(&registrymgmt.Config{
    HTTPClient: tc,
})

registries, _, err := client.RegistriesApi.GetRegistries(ctx).Execute()
```

## Error handling

Checking specific error codes

```go
import "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/error"

if errors.IsAPIError(err, registrymgmt.ERROR_4){
 // Do something
}
```

Obtaining API error 

```go
import "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/error"

apiError := errors.GetAPIError(err)
fmt.Println(apiError.Code)
```

## Endpoints

For a full list of the available endpoints, see [Documentation for API Endpoints](./client/README.md#documentation-for-api-endpoints).