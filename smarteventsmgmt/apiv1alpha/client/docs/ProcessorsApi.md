# \ProcessorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessorsAPIAddProcessorToBridge**](ProcessorsApi.md#ProcessorsAPIAddProcessorToBridge) | **Post** /api/smartevents_mgmt/v1/bridges/{bridgeId}/processors | Create a Processor of a Bridge instance
[**ProcessorsAPIDeleteProcessor**](ProcessorsApi.md#ProcessorsAPIDeleteProcessor) | **Delete** /api/smartevents_mgmt/v1/bridges/{bridgeId}/processors/{processorId} | Delete a Processor of a Bridge instance
[**ProcessorsAPIGetProcessor**](ProcessorsApi.md#ProcessorsAPIGetProcessor) | **Get** /api/smartevents_mgmt/v1/bridges/{bridgeId}/processors/{processorId} | Get a Processor of a Bridge instance
[**ProcessorsAPIListProcessors**](ProcessorsApi.md#ProcessorsAPIListProcessors) | **Get** /api/smartevents_mgmt/v1/bridges/{bridgeId}/processors | Get the list of Processors of a Bridge instance
[**ProcessorsAPIUpdateProcessor**](ProcessorsApi.md#ProcessorsAPIUpdateProcessor) | **Put** /api/smartevents_mgmt/v1/bridges/{bridgeId}/processors/{processorId} | Update a Processor instance Filter definition or Transformation template.



## ProcessorsAPIAddProcessorToBridge

> ProcessorResponse ProcessorsAPIAddProcessorToBridge(ctx, bridgeId).ProcessorRequest(processorRequest).Execute()

Create a Processor of a Bridge instance



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
    processorRequest := *openapiclient.NewProcessorRequest("Name_example") // ProcessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIAddProcessorToBridge(context.Background(), bridgeId).ProcessorRequest(processorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIAddProcessorToBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPIAddProcessorToBridge`: ProcessorResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPIAddProcessorToBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIAddProcessorToBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processorRequest** | [**ProcessorRequest**](ProcessorRequest.md) |  | 

### Return type

[**ProcessorResponse**](ProcessorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorsAPIDeleteProcessor

> ProcessorsAPIDeleteProcessor(ctx, bridgeId, processorId).Execute()

Delete a Processor of a Bridge instance



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
    processorId := "processorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIDeleteProcessor(context.Background(), bridgeId, processorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIDeleteProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**processorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIDeleteProcessorRequest struct via the builder pattern


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


## ProcessorsAPIGetProcessor

> ProcessorResponse ProcessorsAPIGetProcessor(ctx, bridgeId, processorId).Execute()

Get a Processor of a Bridge instance



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
    processorId := "processorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIGetProcessor(context.Background(), bridgeId, processorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIGetProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPIGetProcessor`: ProcessorResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPIGetProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**processorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIGetProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProcessorResponse**](ProcessorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorsAPIListProcessors

> ProcessorListResponse ProcessorsAPIListProcessors(ctx, bridgeId).Name(name).Page(page).Size(size).Status(status).Type_(type_).Execute()

Get the list of Processors of a Bridge instance



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
    type_ := openapiclient.ProcessorType("source") // ProcessorType |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIListProcessors(context.Background(), bridgeId).Name(name).Page(page).Size(size).Status(status).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIListProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPIListProcessors`: ProcessorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPIListProcessors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIListProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 
 **type_** | [**ProcessorType**](ProcessorType.md) |  | 

### Return type

[**ProcessorListResponse**](ProcessorListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorsAPIUpdateProcessor

> ProcessorResponse ProcessorsAPIUpdateProcessor(ctx, bridgeId, processorId).ProcessorRequest(processorRequest).Execute()

Update a Processor instance Filter definition or Transformation template.



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
    processorId := "processorId_example" // string | 
    processorRequest := *openapiclient.NewProcessorRequest("Name_example") // ProcessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIUpdateProcessor(context.Background(), bridgeId, processorId).ProcessorRequest(processorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIUpdateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPIUpdateProcessor`: ProcessorResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPIUpdateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 
**processorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIUpdateProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **processorRequest** | [**ProcessorRequest**](ProcessorRequest.md) |  | 

### Return type

[**ProcessorResponse**](ProcessorResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

