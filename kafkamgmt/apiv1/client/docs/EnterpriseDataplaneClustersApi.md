# \EnterpriseDataplaneClustersApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnterpriseClusterById**](EnterpriseDataplaneClustersApi.md#DeleteEnterpriseClusterById) | **Delete** /api/kafkas_mgmt/v1/clusters/{id} | 
[**GetEnterpriseClusterById**](EnterpriseDataplaneClustersApi.md#GetEnterpriseClusterById) | **Get** /api/kafkas_mgmt/v1/clusters/{id} | 
[**GetEnterpriseClusterWithAddonParameters**](EnterpriseDataplaneClustersApi.md#GetEnterpriseClusterWithAddonParameters) | **Get** /api/kafkas_mgmt/v1/clusters/{id}/addon_parameters | 
[**GetEnterpriseOsdClusters**](EnterpriseDataplaneClustersApi.md#GetEnterpriseOsdClusters) | **Get** /api/kafkas_mgmt/v1/clusters | 
[**RegisterEnterpriseOsdCluster**](EnterpriseDataplaneClustersApi.md#RegisterEnterpriseOsdCluster) | **Post** /api/kafkas_mgmt/v1/clusters | 



## DeleteEnterpriseClusterById

> Error DeleteEnterpriseClusterById(ctx, id).Async(async).Force(force).Execute()



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
    async := true // bool | Perform the action in an asynchronous manner
    id := "id_example" // string | ID of the enterprise data plane cluster
    force := true // bool | When provided with value: true - enterprise cluster will be deleted alongside all kafkas present on the cluster. When skipped and enterprise cluster has any kafkas associated with it, the request will fail. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.DeleteEnterpriseClusterById(context.Background(), id).Async(async).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.DeleteEnterpriseClusterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnterpriseClusterById`: Error
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseDataplaneClustersApi.DeleteEnterpriseClusterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the enterprise data plane cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnterpriseClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Perform the action in an asynchronous manner | 

 **force** | **bool** | When provided with value: true - enterprise cluster will be deleted alongside all kafkas present on the cluster. When skipped and enterprise cluster has any kafkas associated with it, the request will fail. | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnterpriseClusterById

> EnterpriseCluster GetEnterpriseClusterById(ctx, id).Execute()





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
    id := "id_example" // string | ID of the enterprise data plane cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.GetEnterpriseClusterById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.GetEnterpriseClusterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnterpriseClusterById`: EnterpriseCluster
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseDataplaneClustersApi.GetEnterpriseClusterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the enterprise data plane cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnterpriseClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnterpriseCluster**](EnterpriseCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnterpriseClusterWithAddonParameters

> EnterpriseClusterWithAddonParameters GetEnterpriseClusterWithAddonParameters(ctx, id).Execute()





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
    id := "id_example" // string | ID of the enterprise data plane cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.GetEnterpriseClusterWithAddonParameters(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.GetEnterpriseClusterWithAddonParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnterpriseClusterWithAddonParameters`: EnterpriseClusterWithAddonParameters
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseDataplaneClustersApi.GetEnterpriseClusterWithAddonParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the enterprise data plane cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnterpriseClusterWithAddonParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnterpriseClusterWithAddonParameters**](EnterpriseClusterWithAddonParameters.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnterpriseOsdClusters

> EnterpriseClusterList GetEnterpriseOsdClusters(ctx).Execute()





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
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.GetEnterpriseOsdClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.GetEnterpriseOsdClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnterpriseOsdClusters`: EnterpriseClusterList
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseDataplaneClustersApi.GetEnterpriseOsdClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnterpriseOsdClustersRequest struct via the builder pattern


### Return type

[**EnterpriseClusterList**](EnterpriseClusterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterEnterpriseOsdCluster

> EnterpriseClusterWithAddonParameters RegisterEnterpriseOsdCluster(ctx).EnterpriseOsdClusterPayload(enterpriseOsdClusterPayload).Execute()





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
    enterpriseOsdClusterPayload := *openapiclient.NewEnterpriseOsdClusterPayload(false, "ClusterId_example", "ClusterExternalId_example", "ClusterIngressDnsName_example", int32(123)) // EnterpriseOsdClusterPayload | Enterprise data plane cluster details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster(context.Background()).EnterpriseOsdClusterPayload(enterpriseOsdClusterPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEnterpriseOsdCluster`: EnterpriseClusterWithAddonParameters
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterEnterpriseOsdClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enterpriseOsdClusterPayload** | [**EnterpriseOsdClusterPayload**](EnterpriseOsdClusterPayload.md) | Enterprise data plane cluster details | 

### Return type

[**EnterpriseClusterWithAddonParameters**](EnterpriseClusterWithAddonParameters.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

