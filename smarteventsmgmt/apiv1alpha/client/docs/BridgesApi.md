# \BridgesApi

All URIs are relative to *https://api.stage.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BridgesAPICreateBridge**](BridgesApi.md#BridgesAPICreateBridge) | **Post** /api/smartevents_mgmt/v2/bridges | Create a Bridge instance
[**BridgesAPIDeleteBridge**](BridgesApi.md#BridgesAPIDeleteBridge) | **Delete** /api/smartevents_mgmt/v2/bridges/{bridgeId} | Delete a Bridge instance
[**BridgesAPIGetBridge**](BridgesApi.md#BridgesAPIGetBridge) | **Get** /api/smartevents_mgmt/v2/bridges/{bridgeId} | Get a Bridge instance
[**BridgesAPIGetBridges**](BridgesApi.md#BridgesAPIGetBridges) | **Get** /api/smartevents_mgmt/v2/bridges | Get the list of Bridge instances
[**BridgesAPIUpdateBridge**](BridgesApi.md#BridgesAPIUpdateBridge) | **Put** /api/smartevents_mgmt/v2/bridges/{bridgeId} | Update a Bridge instance



## BridgesAPICreateBridge

> BridgeResponse BridgesAPICreateBridge(ctx).BridgeRequest(bridgeRequest).Execute()

Create a Bridge instance



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
    bridgeRequest := *openapiclient.NewBridgeRequest("Name_example", "CloudProvider_example", "Region_example") // BridgeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgesApi.BridgesAPICreateBridge(context.Background()).BridgeRequest(bridgeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgesApi.BridgesAPICreateBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BridgesAPICreateBridge`: BridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgesApi.BridgesAPICreateBridge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBridgesAPICreateBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridgeRequest** | [**BridgeRequest**](BridgeRequest.md) |  | 

### Return type

[**BridgeResponse**](BridgeResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BridgesAPIDeleteBridge

> BridgesAPIDeleteBridge(ctx, bridgeId).Execute()

Delete a Bridge instance



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgesApi.BridgesAPIDeleteBridge(context.Background(), bridgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgesApi.BridgesAPIDeleteBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBridgesAPIDeleteBridgeRequest struct via the builder pattern


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


## BridgesAPIGetBridge

> BridgeResponse BridgesAPIGetBridge(ctx, bridgeId).Execute()

Get a Bridge instance



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgesApi.BridgesAPIGetBridge(context.Background(), bridgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgesApi.BridgesAPIGetBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BridgesAPIGetBridge`: BridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgesApi.BridgesAPIGetBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBridgesAPIGetBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BridgeResponse**](BridgeResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BridgesAPIGetBridges

> BridgeListResponse BridgesAPIGetBridges(ctx).Name(name).Page(page).Size(size).Status(status).Execute()

Get the list of Bridge instances



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
    name := "name_example" // string |  (optional)
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    status := []openapiclient.ManagedResourceStatus{openapiclient.ManagedResourceStatus("accepted")} // []ManagedResourceStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgesApi.BridgesAPIGetBridges(context.Background()).Name(name).Page(page).Size(size).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgesApi.BridgesAPIGetBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BridgesAPIGetBridges`: BridgeListResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgesApi.BridgesAPIGetBridges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBridgesAPIGetBridgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **status** | [**[]ManagedResourceStatus**](ManagedResourceStatus.md) |  | 

### Return type

[**BridgeListResponse**](BridgeListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BridgesAPIUpdateBridge

> BridgeResponse BridgesAPIUpdateBridge(ctx, bridgeId).BridgeRequest(bridgeRequest).Execute()

Update a Bridge instance



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
    bridgeRequest := *openapiclient.NewBridgeRequest("Name_example", "CloudProvider_example", "Region_example") // BridgeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgesApi.BridgesAPIUpdateBridge(context.Background(), bridgeId).BridgeRequest(bridgeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgesApi.BridgesAPIUpdateBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BridgesAPIUpdateBridge`: BridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgesApi.BridgesAPIUpdateBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bridgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBridgesAPIUpdateBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bridgeRequest** | [**BridgeRequest**](BridgeRequest.md) |  | 

### Return type

[**BridgeResponse**](BridgeResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

