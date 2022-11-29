# \ErrorCatalogApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ErrorsAPIGetError**](ErrorCatalogApi.md#ErrorsAPIGetError) | **Get** /api/smartevents_mgmt/v2/errors/{id} | Get an error from the error catalog.
[**ErrorsAPIGetErrors**](ErrorCatalogApi.md#ErrorsAPIGetErrors) | **Get** /api/smartevents_mgmt/v2/errors | Get the list of errors.



## ErrorsAPIGetError

> BridgeError ErrorsAPIGetError(ctx, id).Execute()

Get an error from the error catalog.



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrorCatalogApi.ErrorsAPIGetError(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrorCatalogApi.ErrorsAPIGetError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ErrorsAPIGetError`: BridgeError
    fmt.Fprintf(os.Stdout, "Response from `ErrorCatalogApi.ErrorsAPIGetError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiErrorsAPIGetErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BridgeError**](BridgeError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ErrorsAPIGetErrors

> ErrorListResponse ErrorsAPIGetErrors(ctx).Page(page).Size(size).Execute()

Get the list of errors.



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
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrorCatalogApi.ErrorsAPIGetErrors(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrorCatalogApi.ErrorsAPIGetErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ErrorsAPIGetErrors`: ErrorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ErrorCatalogApi.ErrorsAPIGetErrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiErrorsAPIGetErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]

### Return type

[**ErrorListResponse**](ErrorListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

