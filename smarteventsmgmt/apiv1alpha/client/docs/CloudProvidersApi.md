# \CloudProvidersApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProviderAPIGetCloudProvider**](CloudProvidersApi.md#CloudProviderAPIGetCloudProvider) | **Get** /api/smartevents_mgmt/v2/cloud_providers/{id} | Get Cloud Provider.
[**CloudProviderAPIListCloudProviderRegions**](CloudProvidersApi.md#CloudProviderAPIListCloudProviderRegions) | **Get** /api/smartevents_mgmt/v2/cloud_providers/{id}/regions | List Supported Cloud Regions.
[**CloudProviderAPIListCloudProviders**](CloudProvidersApi.md#CloudProviderAPIListCloudProviders) | **Get** /api/smartevents_mgmt/v2/cloud_providers | List Supported Cloud Providers.



## CloudProviderAPIGetCloudProvider

> CloudProviderListResponse CloudProviderAPIGetCloudProvider(ctx, id).Execute()

Get Cloud Provider.



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.CloudProviderAPIGetCloudProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.CloudProviderAPIGetCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudProviderAPIGetCloudProvider`: CloudProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.CloudProviderAPIGetCloudProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderAPIGetCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProviderListResponse**](CloudProviderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProviderAPIListCloudProviderRegions

> CloudRegionListResponse CloudProviderAPIListCloudProviderRegions(ctx, id).Page(page).Size(size).Execute()

List Supported Cloud Regions.



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
    id := "id_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.CloudProviderAPIListCloudProviderRegions(context.Background(), id).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.CloudProviderAPIListCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudProviderAPIListCloudProviderRegions`: CloudRegionListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.CloudProviderAPIListCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderAPIListCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]

### Return type

[**CloudRegionListResponse**](CloudRegionListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProviderAPIListCloudProviders

> CloudProviderListResponse CloudProviderAPIListCloudProviders(ctx).Page(page).Size(size).Execute()

List Supported Cloud Providers.



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
    resp, r, err := api_client.CloudProvidersApi.CloudProviderAPIListCloudProviders(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.CloudProviderAPIListCloudProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudProviderAPIListCloudProviders`: CloudProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.CloudProviderAPIListCloudProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudProviderAPIListCloudProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]

### Return type

[**CloudProviderListResponse**](CloudProviderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

