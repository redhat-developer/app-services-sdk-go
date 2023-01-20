# \SourceConnectorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourceConnectorsAPICreateSourceConnector**](SourceConnectorsApi.md#SourceConnectorsAPICreateSourceConnector) | **Post** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources | Create a Source Connector for a Bridge instance
[**SourceConnectorsAPIDeleteSourceConnector**](SourceConnectorsApi.md#SourceConnectorsAPIDeleteSourceConnector) | **Delete** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Delete a Source Connector of a Bridge instance
[**SourceConnectorsAPIGetSourceConnector**](SourceConnectorsApi.md#SourceConnectorsAPIGetSourceConnector) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Get a Source Connector of a Bridge instance
[**SourceConnectorsAPIGetSourceConnectors**](SourceConnectorsApi.md#SourceConnectorsAPIGetSourceConnectors) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources | Get the list of Sink Connectors for a Bridge
[**SourceConnectorsAPIUpdateSourceConnector**](SourceConnectorsApi.md#SourceConnectorsAPIUpdateSourceConnector) | **Put** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sources/{sourceId} | Update a Source Connector instance.



## SourceConnectorsAPICreateSourceConnector

> SourceConnectorResponse SourceConnectorsAPICreateSourceConnector(ctx, bridgeId).ConnectorRequest(connectorRequest).Execute()

Create a Source Connector for a Bridge instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bridgeId := "bridgeId_example" // string | 
    connectorRequest := *openapiclient.NewConnectorRequest("my-connector", "slack_sink_0.1", map[string]interface{}(123)) // ConnectorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourceConnectorsApi.SourceConnectorsAPICreateSourceConnector(context.Background(), bridgeId).ConnectorRequest(connectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceConnectorsApi.SourceConnectorsAPICreateSourceConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SourceConnectorsAPICreateSourceConnector`: SourceConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceConnectorsApi.SourceConnectorsAPICreateSourceConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceConnectorsAPICreateSourceConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorRequest** | [**ConnectorRequest**](ConnectorRequest.md) |  | 

### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceConnectorsAPIDeleteSourceConnector

> SourceConnectorsAPIDeleteSourceConnector(ctx, bridgeId, sourceId).Execute()

Delete a Source Connector of a Bridge instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bridgeId := "bridgeId_example" // string | 
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourceConnectorsApi.SourceConnectorsAPIDeleteSourceConnector(context.Background(), bridgeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceConnectorsApi.SourceConnectorsAPIDeleteSourceConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceConnectorsAPIDeleteSourceConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceConnectorsAPIGetSourceConnector

> SourceConnectorResponse SourceConnectorsAPIGetSourceConnector(ctx, bridgeId, sourceId).Execute()

Get a Source Connector of a Bridge instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bridgeId := "bridgeId_example" // string | 
    sourceId := "sourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourceConnectorsApi.SourceConnectorsAPIGetSourceConnector(context.Background(), bridgeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceConnectorsApi.SourceConnectorsAPIGetSourceConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SourceConnectorsAPIGetSourceConnector`: SourceConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceConnectorsApi.SourceConnectorsAPIGetSourceConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceConnectorsAPIGetSourceConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceConnectorsAPIGetSourceConnectors

> SourceConnectorListResponse SourceConnectorsAPIGetSourceConnectors(ctx, bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()

Get the list of Sink Connectors for a Bridge



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bridgeId := "bridgeId_example" // string | 
    name := "name_example" // string |  (optional)
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    status := []openapiclient.ManagedResourceStatus{openapiclient.ManagedResourceStatus("accepted")} // []ManagedResourceStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourceConnectorsApi.SourceConnectorsAPIGetSourceConnectors(context.Background(), bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceConnectorsApi.SourceConnectorsAPIGetSourceConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SourceConnectorsAPIGetSourceConnectors`: SourceConnectorListResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceConnectorsApi.SourceConnectorsAPIGetSourceConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceConnectorsAPIGetSourceConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 

### Return type

[**SourceConnectorListResponse**](SourceConnectorListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceConnectorsAPIUpdateSourceConnector

> SourceConnectorResponse SourceConnectorsAPIUpdateSourceConnector(ctx, bridgeId, sourceId).ConnectorRequest(connectorRequest).Execute()

Update a Source Connector instance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bridgeId := "bridgeId_example" // string | 
    sourceId := "sourceId_example" // string | 
    connectorRequest := *openapiclient.NewConnectorRequest("my-connector", "slack_sink_0.1", map[string]interface{}(123)) // ConnectorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SourceConnectorsApi.SourceConnectorsAPIUpdateSourceConnector(context.Background(), bridgeId, sourceId).ConnectorRequest(connectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceConnectorsApi.SourceConnectorsAPIUpdateSourceConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SourceConnectorsAPIUpdateSourceConnector`: SourceConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceConnectorsApi.SourceConnectorsAPIUpdateSourceConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceConnectorsAPIUpdateSourceConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectorRequest** | [**ConnectorRequest**](ConnectorRequest.md) |  | 

### Return type

[**SourceConnectorResponse**](SourceConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

