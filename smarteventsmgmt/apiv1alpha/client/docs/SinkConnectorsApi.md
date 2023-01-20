# \SinkConnectorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SinkConnectorsAPICreateSinkConnector**](SinkConnectorsApi.md#SinkConnectorsAPICreateSinkConnector) | **Post** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks | Create a Sink Connector for a Bridge instance
[**SinkConnectorsAPIDeleteSinkConnector**](SinkConnectorsApi.md#SinkConnectorsAPIDeleteSinkConnector) | **Delete** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Delete a Sink Connector of a Bridge instance
[**SinkConnectorsAPIGetSinkConnector**](SinkConnectorsApi.md#SinkConnectorsAPIGetSinkConnector) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Get a Sink Connector of a Bridge instance
[**SinkConnectorsAPIGetSinkConnectors**](SinkConnectorsApi.md#SinkConnectorsAPIGetSinkConnectors) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks | Get the list of Sink Connectors for a Bridge
[**SinkConnectorsAPIUpdateSinkConnector**](SinkConnectorsApi.md#SinkConnectorsAPIUpdateSinkConnector) | **Put** /api/smartevents_mgmt/v2/bridges/{bridgeId}/sinks/{sinkId} | Update a Sink Connector instance.



## SinkConnectorsAPICreateSinkConnector

> SinkConnectorResponse SinkConnectorsAPICreateSinkConnector(ctx, bridgeId).ConnectorRequest(connectorRequest).Execute()

Create a Sink Connector for a Bridge instance



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
    resp, r, err := api_client.SinkConnectorsApi.SinkConnectorsAPICreateSinkConnector(context.Background(), bridgeId).ConnectorRequest(connectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkConnectorsApi.SinkConnectorsAPICreateSinkConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SinkConnectorsAPICreateSinkConnector`: SinkConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SinkConnectorsApi.SinkConnectorsAPICreateSinkConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSinkConnectorsAPICreateSinkConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorRequest** | [**ConnectorRequest**](ConnectorRequest.md) |  | 

### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SinkConnectorsAPIDeleteSinkConnector

> SinkConnectorsAPIDeleteSinkConnector(ctx, bridgeId, sinkId).Execute()

Delete a Sink Connector of a Bridge instance



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
    sinkId := "sinkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SinkConnectorsApi.SinkConnectorsAPIDeleteSinkConnector(context.Background(), bridgeId, sinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkConnectorsApi.SinkConnectorsAPIDeleteSinkConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sinkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSinkConnectorsAPIDeleteSinkConnectorRequest struct via the builder pattern


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


## SinkConnectorsAPIGetSinkConnector

> SinkConnectorResponse SinkConnectorsAPIGetSinkConnector(ctx, bridgeId, sinkId).Execute()

Get a Sink Connector of a Bridge instance



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
    sinkId := "sinkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SinkConnectorsApi.SinkConnectorsAPIGetSinkConnector(context.Background(), bridgeId, sinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkConnectorsApi.SinkConnectorsAPIGetSinkConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SinkConnectorsAPIGetSinkConnector`: SinkConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SinkConnectorsApi.SinkConnectorsAPIGetSinkConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sinkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSinkConnectorsAPIGetSinkConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SinkConnectorsAPIGetSinkConnectors

> SinkConnectorListResponse SinkConnectorsAPIGetSinkConnectors(ctx, bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()

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
    resp, r, err := api_client.SinkConnectorsApi.SinkConnectorsAPIGetSinkConnectors(context.Background(), bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkConnectorsApi.SinkConnectorsAPIGetSinkConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SinkConnectorsAPIGetSinkConnectors`: SinkConnectorListResponse
    fmt.Fprintf(os.Stdout, "Response from `SinkConnectorsApi.SinkConnectorsAPIGetSinkConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSinkConnectorsAPIGetSinkConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 

### Return type

[**SinkConnectorListResponse**](SinkConnectorListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SinkConnectorsAPIUpdateSinkConnector

> SinkConnectorResponse SinkConnectorsAPIUpdateSinkConnector(ctx, bridgeId, sinkId).ConnectorRequest(connectorRequest).Execute()

Update a Sink Connector instance.



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
    sinkId := "sinkId_example" // string | 
    connectorRequest := *openapiclient.NewConnectorRequest("my-connector", "slack_sink_0.1", map[string]interface{}(123)) // ConnectorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SinkConnectorsApi.SinkConnectorsAPIUpdateSinkConnector(context.Background(), bridgeId, sinkId).ConnectorRequest(connectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SinkConnectorsApi.SinkConnectorsAPIUpdateSinkConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SinkConnectorsAPIUpdateSinkConnector`: SinkConnectorResponse
    fmt.Fprintf(os.Stdout, "Response from `SinkConnectorsApi.SinkConnectorsAPIUpdateSinkConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**sinkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSinkConnectorsAPIUpdateSinkConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectorRequest** | [**ConnectorRequest**](ConnectorRequest.md) |  | 

### Return type

[**SinkConnectorResponse**](SinkConnectorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

