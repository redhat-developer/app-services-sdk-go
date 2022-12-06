# \ConnectorClustersApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorCluster**](ConnectorClustersApi.md#CreateConnectorCluster) | **Post** /api/connector_mgmt/v1/kafka_connector_clusters | Create a new connector cluster
[**DeleteConnectorCluster**](ConnectorClustersApi.md#DeleteConnectorCluster) | **Delete** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Delete a connector cluster
[**GetConnectorCluster**](ConnectorClustersApi.md#GetConnectorCluster) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Get a connector cluster
[**GetConnectorClusterAddonParameters**](ConnectorClustersApi.md#GetConnectorClusterAddonParameters) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id}/addon_parameters | Get a connector cluster&#39;s addon parameters
[**GetConnectorClusterNamespaces**](ConnectorClustersApi.md#GetConnectorClusterNamespaces) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id}/namespaces | Get a connector cluster&#39;s namespaces
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

> []AddonParameter GetConnectorClusterAddonParameters(ctx, connectorClusterId).ResetCredentials(resetCredentials).Execute()

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
    resetCredentials := true // bool | Resets cluster service account credentials when true (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.GetConnectorClusterAddonParameters(context.Background(), connectorClusterId).ResetCredentials(resetCredentials).Execute()
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

 **resetCredentials** | **bool** | Resets cluster service account credentials when true | 

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


## GetConnectorClusterNamespaces

> ConnectorNamespaceList GetConnectorClusterNamespaces(ctx, connectorClusterId).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Get a connector cluster's namespaces



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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the `order by` clause of an SQL statement. Each query can be ordered by any of the underlying resource fields supported in the search parameter. For example, to return all Connector types ordered by their name, use the following syntax:  ```sql name asc ```  To return all Connector types ordered by their name _and_ version, use the following syntax:  ```sql name asc, version asc ```  If the parameter isn't provided, or if the value is empty, then the results are ordered by name. (optional)
    search := "name = aws-sqs-source and channel = stable" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the `where` clause of a SQL statement.  Allowed fields in the search depend on the resource type:  * Cluster: id, created_at, updated_at, owner, organisation_id, name, state, client_id * Namespace: id, created_at, updated_at, name, cluster_id, owner, expiration, tenant_user_id, tenant_organisation_id, state * Connector Types: id, created_at, updated_at, version, name, description, label, channel, featured_rank, pricing_tier * Connectors: id, created_at, updated_at, name, owner, organisation_id, connector_type_id, desired_state, state, channel, namespace_id, kafka_id, kafka_bootstrap_server, service_account_client_id, schema_registry_id, schema_registry_url  Allowed operators are `<>`, `=`, `LIKE`, or `ILIKE`. Allowed conjunctive operators are `AND` and `OR`. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name `aws-sqs-source` and the channel `stable`, use the following syntax:  ``` name = aws-sqs-source and channel = stable ```[p-]  To return a connector instance with a name that starts with `aws`, use the following syntax:  ``` name like aws%25 ```  To return a connector type with a name containing `aws` matching any character case combination, use the following syntax:  ``` name ilike %25aws%25 ```  If the parameter isn't provided, or if the value is empty, then all the resources that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.GetConnectorClusterNamespaces(context.Background(), connectorClusterId).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorClustersApi.GetConnectorClusterNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorClusterNamespaces`: ConnectorNamespaceList
    fmt.Fprintf(os.Stdout, "Response from `ConnectorClustersApi.GetConnectorClusterNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorClusterId** | **string** | The id of the connector cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorClusterNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the &#x60;order by&#x60; clause of an SQL statement. Each query can be ordered by any of the underlying resource fields supported in the search parameter. For example, to return all Connector types ordered by their name, use the following syntax:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  To return all Connector types ordered by their name _and_ version, use the following syntax:  &#x60;&#x60;&#x60;sql name asc, version asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results are ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the &#x60;where&#x60; clause of a SQL statement.  Allowed fields in the search depend on the resource type:  * Cluster: id, created_at, updated_at, owner, organisation_id, name, state, client_id * Namespace: id, created_at, updated_at, name, cluster_id, owner, expiration, tenant_user_id, tenant_organisation_id, state * Connector Types: id, created_at, updated_at, version, name, description, label, channel, featured_rank, pricing_tier * Connectors: id, created_at, updated_at, name, owner, organisation_id, connector_type_id, desired_state, state, channel, namespace_id, kafka_id, kafka_bootstrap_server, service_account_client_id, schema_registry_id, schema_registry_url  Allowed operators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;LIKE&#x60;, or &#x60;ILIKE&#x60;. Allowed conjunctive operators are &#x60;AND&#x60; and &#x60;OR&#x60;. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name &#x60;aws-sqs-source&#x60; and the channel &#x60;stable&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name &#x3D; aws-sqs-source and channel &#x3D; stable &#x60;&#x60;&#x60;[p-]  To return a connector instance with a name that starts with &#x60;aws&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name like aws%25 &#x60;&#x60;&#x60;  To return a connector type with a name containing &#x60;aws&#x60; matching any character case combination, use the following syntax:  &#x60;&#x60;&#x60; name ilike %25aws%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the resources that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  | 

### Return type

[**ConnectorNamespaceList**](ConnectorNamespaceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorClusters

> ConnectorClusterList ListConnectorClusters(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

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
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the `order by` clause of an SQL statement. Each query can be ordered by any of the underlying resource fields supported in the search parameter. For example, to return all Connector types ordered by their name, use the following syntax:  ```sql name asc ```  To return all Connector types ordered by their name _and_ version, use the following syntax:  ```sql name asc, version asc ```  If the parameter isn't provided, or if the value is empty, then the results are ordered by name. (optional)
    search := "name = aws-sqs-source and channel = stable" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the `where` clause of a SQL statement.  Allowed fields in the search depend on the resource type:  * Cluster: id, created_at, updated_at, owner, organisation_id, name, state, client_id * Namespace: id, created_at, updated_at, name, cluster_id, owner, expiration, tenant_user_id, tenant_organisation_id, state * Connector Types: id, created_at, updated_at, version, name, description, label, channel, featured_rank, pricing_tier * Connectors: id, created_at, updated_at, name, owner, organisation_id, connector_type_id, desired_state, state, channel, namespace_id, kafka_id, kafka_bootstrap_server, service_account_client_id, schema_registry_id, schema_registry_url  Allowed operators are `<>`, `=`, `LIKE`, or `ILIKE`. Allowed conjunctive operators are `AND` and `OR`. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name `aws-sqs-source` and the channel `stable`, use the following syntax:  ``` name = aws-sqs-source and channel = stable ```[p-]  To return a connector instance with a name that starts with `aws`, use the following syntax:  ``` name like aws%25 ```  To return a connector type with a name containing `aws` matching any character case combination, use the following syntax:  ``` name ilike %25aws%25 ```  If the parameter isn't provided, or if the value is empty, then all the resources that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorClustersApi.ListConnectorClusters(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
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
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the &#x60;order by&#x60; clause of an SQL statement. Each query can be ordered by any of the underlying resource fields supported in the search parameter. For example, to return all Connector types ordered by their name, use the following syntax:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  To return all Connector types ordered by their name _and_ version, use the following syntax:  &#x60;&#x60;&#x60;sql name asc, version asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results are ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the &#x60;where&#x60; clause of a SQL statement.  Allowed fields in the search depend on the resource type:  * Cluster: id, created_at, updated_at, owner, organisation_id, name, state, client_id * Namespace: id, created_at, updated_at, name, cluster_id, owner, expiration, tenant_user_id, tenant_organisation_id, state * Connector Types: id, created_at, updated_at, version, name, description, label, channel, featured_rank, pricing_tier * Connectors: id, created_at, updated_at, name, owner, organisation_id, connector_type_id, desired_state, state, channel, namespace_id, kafka_id, kafka_bootstrap_server, service_account_client_id, schema_registry_id, schema_registry_url  Allowed operators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60;, &#x60;LIKE&#x60;, or &#x60;ILIKE&#x60;. Allowed conjunctive operators are &#x60;AND&#x60; and &#x60;OR&#x60;. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name &#x60;aws-sqs-source&#x60; and the channel &#x60;stable&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name &#x3D; aws-sqs-source and channel &#x3D; stable &#x60;&#x60;&#x60;[p-]  To return a connector instance with a name that starts with &#x60;aws&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name like aws%25 &#x60;&#x60;&#x60;  To return a connector type with a name containing &#x60;aws&#x60; matching any character case combination, use the following syntax:  &#x60;&#x60;&#x60; name ilike %25aws%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the resources that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  | 

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

