# \EnterpriseDataplaneClustersApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnterpriseOsdClusters**](EnterpriseDataplaneClustersApi.md#GetEnterpriseOsdClusters) | **Get** /api/kafkas_mgmt/v1/clusters | 
[**RegisterEnterpriseOsdCluster**](EnterpriseDataplaneClustersApi.md#RegisterEnterpriseOsdCluster) | **Post** /api/kafkas_mgmt/v1/clusters | 



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

> EnterpriseClusterRegistrationResponse RegisterEnterpriseOsdCluster(ctx).EnterpriseOsdClusterPayload(enterpriseOsdClusterPayload).Execute()





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
    enterpriseOsdClusterPayload := *openapiclient.NewEnterpriseOsdClusterPayload("ClusterId_example", "ClusterExternalId_example", "ClusterIngressDnsName_example", int32(123)) // EnterpriseOsdClusterPayload | Enterprise data plane cluster details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster(context.Background()).EnterpriseOsdClusterPayload(enterpriseOsdClusterPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseDataplaneClustersApi.RegisterEnterpriseOsdCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEnterpriseOsdCluster`: EnterpriseClusterRegistrationResponse
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

[**EnterpriseClusterRegistrationResponse**](EnterpriseClusterRegistrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

