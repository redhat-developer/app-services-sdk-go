# \BlahApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRegistry**](BlahApi.md#DeleteRegistry) | **Delete** /api/serviceregistry_mgmt/v1/{registryId} | Delete a Registry



## DeleteRegistry

> DeleteRegistry(ctx, registryId).Execute()

Delete a Registry



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
    registryId := int32(56) // int32 | A unique identifier for a `Registry`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlahApi.DeleteRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlahApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **int32** | A unique identifier for a &#x60;Registry&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

