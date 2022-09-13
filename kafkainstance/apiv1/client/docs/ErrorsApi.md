# \ErrorsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetError**](ErrorsApi.md#GetError) | **Get** /api/v1/errors/{errorId} | Get an error by its unique ID
[**GetErrors**](ErrorsApi.md#GetErrors) | **Get** /api/v1/errors | Get list of errors



## GetError

> Error GetError(ctx, errorId).Execute()

Get an error by its unique ID

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
    errorId := "errorId_example" // string | Error identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrorsApi.GetError(context.Background(), errorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrorsApi.GetError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetError`: Error
    fmt.Fprintf(os.Stdout, "Response from `ErrorsApi.GetError`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**errorId** | **string** | Error identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetErrors

> ErrorList GetErrors(ctx).Execute()

Get list of errors

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrorsApi.GetErrors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrorsApi.GetErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetErrors`: ErrorList
    fmt.Fprintf(os.Stdout, "Response from `ErrorsApi.GetErrors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorsRequest struct via the builder pattern


### Return type

[**ErrorList**](ErrorList.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

