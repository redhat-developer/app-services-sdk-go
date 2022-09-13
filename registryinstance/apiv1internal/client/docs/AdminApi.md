# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleMapping**](AdminApi.md#CreateRoleMapping) | **Post** /admin/roleMappings | Create a new role mapping
[**DeleteRoleMapping**](AdminApi.md#DeleteRoleMapping) | **Delete** /admin/roleMappings/{principalId} | Delete a role mapping
[**ExportData**](AdminApi.md#ExportData) | **Get** /admin/export | Export registry data
[**GetConfigProperty**](AdminApi.md#GetConfigProperty) | **Get** /admin/config/properties/{propertyName} | Get configuration property value
[**GetLogConfiguration**](AdminApi.md#GetLogConfiguration) | **Get** /admin/loggers/{logger} | Get a single logger configuration
[**GetRoleMapping**](AdminApi.md#GetRoleMapping) | **Get** /admin/roleMappings/{principalId} | Return a single role mapping
[**ImportData**](AdminApi.md#ImportData) | **Post** /admin/import | Import registry data
[**ListConfigProperties**](AdminApi.md#ListConfigProperties) | **Get** /admin/config/properties | List all configuration properties
[**ListLogConfigurations**](AdminApi.md#ListLogConfigurations) | **Get** /admin/loggers | List logging configurations
[**ListRoleMappings**](AdminApi.md#ListRoleMappings) | **Get** /admin/roleMappings | List all role mappings
[**RemoveLogConfiguration**](AdminApi.md#RemoveLogConfiguration) | **Delete** /admin/loggers/{logger} | Removes logger configuration
[**ResetConfigProperty**](AdminApi.md#ResetConfigProperty) | **Delete** /admin/config/properties/{propertyName} | Reset a configuration property
[**SetLogConfiguration**](AdminApi.md#SetLogConfiguration) | **Put** /admin/loggers/{logger} | Set a logger&#39;s configuration
[**UpdateConfigProperty**](AdminApi.md#UpdateConfigProperty) | **Put** /admin/config/properties/{propertyName} | Update a configuration property
[**UpdateRoleMapping**](AdminApi.md#UpdateRoleMapping) | **Put** /admin/roleMappings/{principalId} | Update a role mapping



## CreateRoleMapping

> CreateRoleMapping(ctx).RoleMapping(roleMapping).Execute()

Create a new role mapping



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
    roleMapping := *openapiclient.NewRoleMapping("PrincipalId_example", openapiclient.RoleType("READ_ONLY")) // RoleMapping | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.CreateRoleMapping(context.Background()).RoleMapping(roleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.CreateRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleMapping** | [**RoleMapping**](RoleMapping.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleMapping

> DeleteRoleMapping(ctx, principalId).Execute()

Delete a role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.DeleteRoleMapping(context.Background(), principalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.DeleteRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportData

> *os.File ExportData(ctx).ForBrowser(forBrowser).Execute()

Export registry data



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
    forBrowser := true // bool | Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called `href`.  This `href` will be a single-use, naked download link suitable for use by a web browser to download the content. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.ExportData(context.Background()).ForBrowser(forBrowser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ExportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ExportData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forBrowser** | **bool** | Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called &#x60;href&#x60;.  This &#x60;href&#x60; will be a single-use, naked download link suitable for use by a web browser to download the content. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigProperty

> ConfigurationProperty GetConfigProperty(ctx, propertyName).Execute()

Get configuration property value



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
    propertyName := "propertyName_example" // string | The name of a configuration property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.GetConfigProperty(context.Background(), propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.GetConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigProperty`: ConfigurationProperty
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.GetConfigProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationProperty**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogConfiguration

> NamedLogConfiguration GetLogConfiguration(ctx, logger).Execute()

Get a single logger configuration



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
    logger := "logger_example" // string | The name of a single logger.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.GetLogConfiguration(context.Background(), logger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.GetLogConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogConfiguration`: NamedLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.GetLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logger** | **string** | The name of a single logger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMapping

> RoleMapping GetRoleMapping(ctx, principalId).Execute()

Return a single role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.GetRoleMapping(context.Background(), principalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.GetRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMapping`: RoleMapping
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.GetRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleMapping**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportData

> ImportData(ctx).Body(body).XRegistryPreserveGlobalId(xRegistryPreserveGlobalId).XRegistryPreserveContentId(xRegistryPreserveContentId).Execute()

Import registry data



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
    body := os.NewFile(1234, "some_file") // *os.File | The ZIP file representing the previously exported registry data.
    xRegistryPreserveGlobalId := true // bool | If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict. (optional)
    xRegistryPreserveContentId := true // bool | If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.ImportData(context.Background()).Body(body).XRegistryPreserveGlobalId(xRegistryPreserveGlobalId).XRegistryPreserveContentId(xRegistryPreserveContentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ImportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | The ZIP file representing the previously exported registry data. | 
 **xRegistryPreserveGlobalId** | **bool** | If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict. | 
 **xRegistryPreserveContentId** | **bool** | If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/zip
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigProperties

> []ConfigurationProperty ListConfigProperties(ctx).Execute()

List all configuration properties



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
    resp, r, err := api_client.AdminApi.ListConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ListConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigProperties`: []ConfigurationProperty
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ListConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigPropertiesRequest struct via the builder pattern


### Return type

[**[]ConfigurationProperty**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogConfigurations

> []NamedLogConfiguration ListLogConfigurations(ctx).Execute()

List logging configurations



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
    resp, r, err := api_client.AdminApi.ListLogConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ListLogConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogConfigurations`: []NamedLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ListLogConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogConfigurationsRequest struct via the builder pattern


### Return type

[**[]NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleMappings

> []RoleMapping ListRoleMappings(ctx).Execute()

List all role mappings



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
    resp, r, err := api_client.AdminApi.ListRoleMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ListRoleMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleMappings`: []RoleMapping
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.ListRoleMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleMappingsRequest struct via the builder pattern


### Return type

[**[]RoleMapping**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLogConfiguration

> NamedLogConfiguration RemoveLogConfiguration(ctx, logger).Execute()

Removes logger configuration



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
    logger := "logger_example" // string | The name of a single logger.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.RemoveLogConfiguration(context.Background(), logger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.RemoveLogConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLogConfiguration`: NamedLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.RemoveLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logger** | **string** | The name of a single logger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetConfigProperty

> ResetConfigProperty(ctx, propertyName).Execute()

Reset a configuration property



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
    propertyName := "propertyName_example" // string | The name of a configuration property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.ResetConfigProperty(context.Background(), propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.ResetConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConfigPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLogConfiguration

> NamedLogConfiguration SetLogConfiguration(ctx, logger).LogConfiguration(logConfiguration).Execute()

Set a logger's configuration



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
    logger := "logger_example" // string | The name of a single logger.
    logConfiguration := *openapiclient.NewLogConfiguration(openapiclient.LogLevel("DEBUG")) // LogConfiguration | The new logger configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.SetLogConfiguration(context.Background(), logger).LogConfiguration(logConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.SetLogConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLogConfiguration`: NamedLogConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.SetLogConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logger** | **string** | The name of a single logger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLogConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logConfiguration** | [**LogConfiguration**](LogConfiguration.md) | The new logger configuration. | 

### Return type

[**NamedLogConfiguration**](NamedLogConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigProperty

> UpdateConfigProperty(ctx, propertyName).UpdateConfigurationProperty(updateConfigurationProperty).Execute()

Update a configuration property



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
    propertyName := "propertyName_example" // string | The name of a configuration property.
    updateConfigurationProperty := *openapiclient.NewUpdateConfigurationProperty("Value_example") // UpdateConfigurationProperty | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.UpdateConfigProperty(context.Background(), propertyName).UpdateConfigurationProperty(updateConfigurationProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.UpdateConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConfigurationProperty** | [**UpdateConfigurationProperty**](UpdateConfigurationProperty.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleMapping

> UpdateRoleMapping(ctx, principalId).UpdateRole(updateRole).Execute()

Update a role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).
    updateRole := *openapiclient.NewUpdateRole(openapiclient.RoleType("READ_ONLY")) // UpdateRole | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApi.UpdateRoleMapping(context.Background(), principalId).UpdateRole(updateRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.UpdateRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRole** | [**UpdateRole**](UpdateRole.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

