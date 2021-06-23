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

> ConnectorTypeList ListConnectorTypes(ctx).Page(page).Size(size).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectorTypesApi.ListConnectorTypes(context.Background()).Page(page).Size(size).Execute()
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

