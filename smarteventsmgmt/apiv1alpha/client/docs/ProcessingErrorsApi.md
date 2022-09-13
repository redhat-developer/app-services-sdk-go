# \ProcessingErrorsApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessingErrorsAPIGetBridgeErrors**](ProcessingErrorsApi.md#ProcessingErrorsAPIGetBridgeErrors) | **Get** /api/smartevents_mgmt/v1/bridges/{bridgeId}/errors | Get the list of errors for a particular Bridge instance



## ProcessingErrorsAPIGetBridgeErrors

> ProcessingErrorListResponse ProcessingErrorsAPIGetBridgeErrors(ctx, bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()

Get the list of errors for a particular Bridge instance



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
    resp, r, err := api_client.ProcessingErrorsApi.ProcessingErrorsAPIGetBridgeErrors(context.Background(), bridgeId).Name(name).Page(page).Size(size).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessingErrorsApi.ProcessingErrorsAPIGetBridgeErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessingErrorsAPIGetBridgeErrors`: ProcessingErrorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessingErrorsApi.ProcessingErrorsAPIGetBridgeErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessingErrorsAPIGetBridgeErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 

### Return type

[**ProcessingErrorListResponse**](ProcessingErrorListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

