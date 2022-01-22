# \ConnectorClustersApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorCluster**](ConnectorClustersApi.md#CreateConnectorCluster) | **Post** /api/connector_mgmt/v1/kafka_connector_clusters | Create a new connector cluster
[**DeleteConnectorCluster**](ConnectorClustersApi.md#DeleteConnectorCluster) | **Delete** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Delete a connector cluster
[**GetConnectorCluster**](ConnectorClustersApi.md#GetConnectorCluster) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Get a connector cluster
[**GetConnectorClusterAddonParameters**](ConnectorClustersApi.md#GetConnectorClusterAddonParameters) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id}/addon_parameters | Get a connector cluster&#39;s addon parameters
[**ListConnectorClusters**](ConnectorClustersApi.md#ListConnectorClusters) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters | Returns a list of connector clusters
[**UpdateConnectorClusterById**](ConnectorClustersApi.md#UpdateConnectorClusterById) | **Put** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | udpate a connector cluster



## CreateConnectorCluster

> ConnectorCluster CreateConnectorCluster(ctx).Async(async).ConnectorClusterRequest(connectorClusterRequest).Execute()

Create a new connector cluster



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
    connectorClusterRequest := *openapiclient.NewConnectorClusterRequest() // ConnectorClusterRequest | Connector cluster data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.CreateConnectorCluster(context.Background()).Async(async).ConnectorClusterRequest(connectorClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.CreateConnectorCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectorCluster`: ConnectorCluster
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.CreateConnectorCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Perform the action in an asynchronous manner | 
 **connectorClusterRequest** | [**ConnectorClusterRequest**](ConnectorClusterRequest.md) | Connector cluster data | 

### Return type

[**ConnectorCluster**](ConnectorCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorCluster

> Error DeleteConnectorCluster(ctx, connectorClusterId).Execute()

Delete a connector cluster



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
    connectorClusterId := "connectorClusterId_example" // string | The id of the connector cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.DeleteConnectorCluster(context.Background(), connectorClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.DeleteConnectorCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnectorCluster`: Error
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.DeleteConnectorCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorClusterId** | **string** | The id of the connector cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetConnectorCluster

> ConnectorCluster GetConnectorCluster(ctx, connectorClusterId).Execute()

Get a connector cluster



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
    connectorClusterId := "connectorClusterId_example" // string | The id of the connector cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.GetConnectorCluster(context.Background(), connectorClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.GetConnectorCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorCluster`: ConnectorCluster
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.GetConnectorCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorClusterId** | **string** | The id of the connector cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorCluster**](ConnectorCluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorClusterAddonParameters

> []AddonParameter GetConnectorClusterAddonParameters(ctx, connectorClusterId).Execute()

Get a connector cluster's addon parameters



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
    connectorClusterId := "connectorClusterId_example" // string | The id of the connector cluster

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.GetConnectorClusterAddonParameters(context.Background(), connectorClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.GetConnectorClusterAddonParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorClusterAddonParameters`: []AddonParameter
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.GetConnectorClusterAddonParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorClusterId** | **string** | The id of the connector cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorClusterAddonParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AddonParameter**](AddonParameter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorClusters

> ConnectorClusterList ListConnectorClusters(ctx).Page(page).Size(size).Execute()

Returns a list of connector clusters



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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.ListConnectorClusters(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.ListConnectorClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectorClusters`: ConnectorClusterList
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.ListConnectorClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**ConnectorClusterList**](ConnectorClusterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectorClusterById

> UpdateConnectorClusterById(ctx, connectorClusterId).ConnectorClusterRequest(connectorClusterRequest).Execute()

udpate a connector cluster



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
    connectorClusterId := "connectorClusterId_example" // string | The id of the connector cluster
    connectorClusterRequest := *openapiclient.NewConnectorClusterRequest() // ConnectorClusterRequest | Data to updated connector with

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.UpdateConnectorClusterById(context.Background(), connectorClusterId).ConnectorClusterRequest(connectorClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.UpdateConnectorClusterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorClusterId** | **string** | The id of the connector cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorClusterRequest** | [**ConnectorClusterRequest**](ConnectorClusterRequest.md) | Data to updated connector with | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

