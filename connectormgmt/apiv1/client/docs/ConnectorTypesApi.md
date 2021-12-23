# \ConnectorTypesApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectorTypeByID**](ConnectorTypesApi.md#GetConnectorTypeByID) | **Get** /api/connector_mgmt/v1/kafka_connector_types/{connector_type_id} | Get a connector type by id
[**ListConnectorTypes**](ConnectorTypesApi.md#ListConnectorTypes) | **Get** /api/connector_mgmt/v1/kafka_connector_types | Returns a list of connector types



## GetConnectorTypeByID

> ConnectorType GetConnectorTypeByID(ctx, connectorTypeId).Execute()

Get a connector type by id

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
    connectorTypeId := "connectorTypeId_example" // string | The id of the connector type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorTypesApi.GetConnectorTypeByID(context.Background(), connectorTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorTypesApi.GetConnectorTypeByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorTypeByID`: ConnectorType
    fmt.Fprintf(os.Stdout, "Response from `ConnectorTypesApi.GetConnectorTypeByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorTypeId** | **string** | The id of the connector type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorTypeByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorType**](ConnectorType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorTypes

> ConnectorTypeList ListConnectorTypes(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Returns a list of connector types

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
    resp, r, err := api_client.ConnectorTypesApi.ListConnectorTypes(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorTypesApi.ListConnectorTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectorTypes`: ConnectorTypeList
    fmt.Fprintf(os.Stdout, "Response from `ConnectorTypesApi.ListConnectorTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the &#x60;order by&#x60; clause of an SQL statement. Each query can be ordered by any of the &#x60;ConnectorType&#x60; fields. For example, to return all Connector types ordered by their name, use the following syntax:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  To return all Connector types ordered by their name _and_ version, use the following syntax:  &#x60;&#x60;&#x60;sql name asc, version asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results are ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the &#x60;where&#x60; clause of a SQL statement. Allowed fields in the search are &#x60;name&#x60;, &#x60;description&#x60;, &#x60;version&#x60;, &#x60;label&#x60;, and &#x60;channel&#x60;. Allowed operators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60;, or &#x60;LIKE&#x60;. Allowed conjunctive operators are &#x60;AND&#x60; and &#x60;OR&#x60;. However, you can use a maximum of 10 conjunctions in a search query.  Examples:  To return a Connector Type with the name &#x60;aws-sqs-source&#x60; and the channel &#x60;stable&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name &#x3D; aws-sqs-source and channel &#x3D; stable &#x60;&#x60;&#x60;[p-]  To return a Kafka instance with a name that starts with &#x60;aws&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name like aws%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the Connector Type that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  | 

### Return type

[**ConnectorTypeList**](ConnectorTypeList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

