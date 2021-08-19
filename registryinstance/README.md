# Go API client for Service Registry Instance

Managed Service Registry API Instance API that lets you interact with registry instances. Registry is a datastore for standard event schemas and API designs.

## Installation

To install the package to your project use `go get`:

```shell
go get github.com/redhat-developer/app-services-sdk-go/registryinstance
```

## Usage

### Importing the package

Import the `github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal` package into your code:

```go
package main

import (
    "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal"
)
```

### Basic usage

If you do not need to customise any of the default configuration values you can pass `nil` to the `NewAPIClient` constructor:

```go
client := registryinstance.NewAPIClient(nil)

registries, resp, err := client.ArtifactsAPI.GetArtifacts(context.Background()).Execute()
```

### Configuration

You can override the default configuration options:

```go
cfg := registryinstance.Config{
    HTTPClient: yourHTTPClient,
    Debug: true,
    BaseURL: "http://localhost:8080",
}

client := registryinstance.NewAPIClient(&cfg)
```

## Endpoints

For a full list of the available endpoints, see [Documentation for API Endpoints](./client/README.md#documentation-for-api-endpoints).

## Local development

Registry API requires local patch in order to be valid. 
Please execute manually:

```bash
./scripts/patch_apis.sh
```

to patch API
