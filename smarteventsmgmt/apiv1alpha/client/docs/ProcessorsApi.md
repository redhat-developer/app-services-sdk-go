# \ProcessorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessorsAPICreateProcessor**](ProcessorsApi.md#ProcessorsAPICreateProcessor) | **Post** /api/smartevents_mgmt/v2/bridges/{bridgeId}/processors | Create a Processor of a Bridge instance
[**ProcessorsAPIDeleteProcessor**](ProcessorsApi.md#ProcessorsAPIDeleteProcessor) | **Delete** /api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId} | Delete a Processor of a Bridge instance
[**ProcessorsAPIGetProcessor**](ProcessorsApi.md#ProcessorsAPIGetProcessor) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId} | Get a Processor of a Bridge instance
[**ProcessorsAPIGetProcessors**](ProcessorsApi.md#ProcessorsAPIGetProcessors) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId}/processors | Get the list of Processors of a Bridge instance
[**ProcessorsAPIUpdateProcessor**](ProcessorsApi.md#ProcessorsAPIUpdateProcessor) | **Put** /api/smartevents_mgmt/v2/bridges/{bridgeId}/processors/{processorId} | Update a Processor instance.



## ProcessorsAPICreateProcessor

> ProcessorResponse ProcessorsAPICreateProcessor(ctx, bridgeId).ProcessorRequest(processorRequest).Execute()

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
    processorRequest := *openapiclient.NewProcessorRequest("processor1", map[string]interface{}(123)) // ProcessorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPICreateProcessor(context.Background(), bridgeId).ProcessorRequest(processorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPICreateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPICreateProcessor`: ProcessorResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPICreateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPICreateProcessorRequest struct via the builder pattern


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


## ProcessorsAPIGetProcessors

> ProcessorListResponse ProcessorsAPIGetProcessors(ctx, bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProcessorsApi.ProcessorsAPIGetProcessors(context.Background(), bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessorsApi.ProcessorsAPIGetProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorsAPIGetProcessors`: ProcessorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessorsApi.ProcessorsAPIGetProcessors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessorsAPIGetProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 

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

Update a Processor instance.



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
    processorRequest := *openapiclient.NewProcessorRequest("processor1", map[string]interface{}(123)) // ProcessorRequest |  (optional)

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

