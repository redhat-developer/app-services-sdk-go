# \ConnectorNamespacesApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorNamespace**](ConnectorNamespacesApi.md#CreateConnectorNamespace) | **Post** /api/connector_mgmt/v1/kafka_connector_namespaces | Create a new connector namespace
[**CreateEvaluationNamespace**](ConnectorNamespacesApi.md#CreateEvaluationNamespace) | **Post** /api/connector_mgmt/v1/kafka_connector_namespaces/eval | Create a new short lived evaluation connector namespace
[**DeleteConnectorNamespace**](ConnectorNamespacesApi.md#DeleteConnectorNamespace) | **Delete** /api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id} | Delete a connector namespace
[**GetConnectorNamespace**](ConnectorNamespacesApi.md#GetConnectorNamespace) | **Get** /api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id} | Get a connector namespace
[**ListConnectorNamespaces**](ConnectorNamespacesApi.md#ListConnectorNamespaces) | **Get** /api/connector_mgmt/v1/kafka_connector_namespaces | Returns a list of connector namespaces
[**UpdateConnectorNamespaceById**](ConnectorNamespacesApi.md#UpdateConnectorNamespaceById) | **Patch** /api/connector_mgmt/v1/kafka_connector_namespaces/{connector_namespace_id} | udpate a connector namespace



## CreateConnectorNamespace

> ConnectorNamespace CreateConnectorNamespace(ctx).ConnectorNamespaceRequest(connectorNamespaceRequest).Execute()

Create a new connector namespace



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
    connectorNamespaceRequest := *openapiclient.NewConnectorNamespaceRequest("Name_example", "ClusterId_example", openapiclient.ConnectorNamespaceTenantKind("user")) // ConnectorNamespaceRequest | Connector namespace data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.CreateConnectorNamespace(context.Background()).ConnectorNamespaceRequest(connectorNamespaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.CreateConnectorNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectorNamespace`: ConnectorNamespace
    fmt.Fprintf(os.Stdout, "Response from `ConnectorNamespacesApi.CreateConnectorNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorNamespaceRequest** | [**ConnectorNamespaceRequest**](ConnectorNamespaceRequest.md) | Connector namespace data | 

### Return type

[**ConnectorNamespace**](ConnectorNamespace.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvaluationNamespace

> ConnectorNamespace CreateEvaluationNamespace(ctx).ConnectorNamespaceEvalRequest(connectorNamespaceEvalRequest).Execute()

Create a new short lived evaluation connector namespace



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
    connectorNamespaceEvalRequest := *openapiclient.NewConnectorNamespaceEvalRequest("Name_example") // ConnectorNamespaceEvalRequest | Connector namespace data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.CreateEvaluationNamespace(context.Background()).ConnectorNamespaceEvalRequest(connectorNamespaceEvalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.CreateEvaluationNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvaluationNamespace`: ConnectorNamespace
    fmt.Fprintf(os.Stdout, "Response from `ConnectorNamespacesApi.CreateEvaluationNamespace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEvaluationNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorNamespaceEvalRequest** | [**ConnectorNamespaceEvalRequest**](ConnectorNamespaceEvalRequest.md) | Connector namespace data | 

### Return type

[**ConnectorNamespace**](ConnectorNamespace.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorNamespace

> Error DeleteConnectorNamespace(ctx, connectorNamespaceId).Execute()

Delete a connector namespace



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
    connectorNamespaceId := "connectorNamespaceId_example" // string | The id of the connector namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.DeleteConnectorNamespace(context.Background(), connectorNamespaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.DeleteConnectorNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnectorNamespace`: Error
    fmt.Fprintf(os.Stdout, "Response from `ConnectorNamespacesApi.DeleteConnectorNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorNamespaceId** | **string** | The id of the connector namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorNamespaceRequest struct via the builder pattern


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


## GetConnectorNamespace

> ConnectorNamespace GetConnectorNamespace(ctx, connectorNamespaceId).Execute()

Get a connector namespace



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
    connectorNamespaceId := "connectorNamespaceId_example" // string | The id of the connector namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.GetConnectorNamespace(context.Background(), connectorNamespaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.GetConnectorNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorNamespace`: ConnectorNamespace
    fmt.Fprintf(os.Stdout, "Response from `ConnectorNamespacesApi.GetConnectorNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorNamespaceId** | **string** | The id of the connector namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorNamespace**](ConnectorNamespace.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorNamespaces

> ConnectorNamespaceList ListConnectorNamespaces(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Returns a list of connector namespaces



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
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the `order by` clause of an SQL statement. Each query can be ordered by any of the `ConnectorType` fields. For example, to return all Connector types ordered by their name, use the following syntax:  ```sql name asc ```  To return all Connector types ordered by their name _and_ version, use the following syntax:  ```sql name asc, version asc ```  If the parameter isn't provided, or if the value is empty, then the results are ordered by name. (optional)
    search := "name = aws-sqs-source and channel = stable" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the `where` clause of a SQL statement. Allowed fields in the search are `name`, `description`, `version`, `label`, and `channel`. Allowed operators are `<>`, `=`, or `LIKE`. Allowed conjunctive operators are `AND` and `OR`. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name `aws-sqs-source` and the channel `stable`, use the following syntax:  ``` name = aws-sqs-source and channel = stable ```[p-]  To return a Kafka instance with a name that starts with `aws`, use the following syntax:  ``` name like aws%25 ```  If the parameter isn't provided, or if the value is empty, then all the Connector Type that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.ListConnectorNamespaces(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.ListConnectorNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectorNamespaces`: ConnectorNamespaceList
    fmt.Fprintf(os.Stdout, "Response from `ConnectorNamespacesApi.ListConnectorNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the &#x60;order by&#x60; clause of an SQL statement. Each query can be ordered by any of the &#x60;ConnectorType&#x60; fields. For example, to return all Connector types ordered by their name, use the following syntax:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  To return all Connector types ordered by their name _and_ version, use the following syntax:  &#x60;&#x60;&#x60;sql name asc, version asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results are ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the &#x60;where&#x60; clause of a SQL statement. Allowed fields in the search are &#x60;name&#x60;, &#x60;description&#x60;, &#x60;version&#x60;, &#x60;label&#x60;, and &#x60;channel&#x60;. Allowed operators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60;, or &#x60;LIKE&#x60;. Allowed conjunctive operators are &#x60;AND&#x60; and &#x60;OR&#x60;. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name &#x60;aws-sqs-source&#x60; and the channel &#x60;stable&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name &#x3D; aws-sqs-source and channel &#x3D; stable &#x60;&#x60;&#x60;[p-]  To return a Kafka instance with a name that starts with &#x60;aws&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name like aws%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the Connector Type that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  | 

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


## UpdateConnectorNamespaceById

> UpdateConnectorNamespaceById(ctx, connectorNamespaceId).ConnectorNamespacePatchRequest(connectorNamespacePatchRequest).Execute()

udpate a connector namespace



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
    connectorNamespaceId := "connectorNamespaceId_example" // string | The id of the connector namespace
    connectorNamespacePatchRequest := *openapiclient.NewConnectorNamespacePatchRequest() // ConnectorNamespacePatchRequest | Data to update namespace with

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorNamespacesApi.UpdateConnectorNamespaceById(context.Background(), connectorNamespaceId).ConnectorNamespacePatchRequest(connectorNamespacePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorNamespacesApi.UpdateConnectorNamespaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorNamespaceId** | **string** | The id of the connector namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorNamespaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorNamespacePatchRequest** | [**ConnectorNamespacePatchRequest**](ConnectorNamespacePatchRequest.md) | Data to update namespace with | 

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

