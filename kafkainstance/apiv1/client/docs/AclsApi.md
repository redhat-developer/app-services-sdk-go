# \AclsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAcl**](AclsApi.md#CreateAcl) | **Post** /api/v1/acls | Create ACL binding
[**DeleteAcls**](AclsApi.md#DeleteAcls) | **Delete** /api/v1/acls | Delete ACL bindings
[**GetAclResourceOperations**](AclsApi.md#GetAclResourceOperations) | **Get** /api/v1/acls/resource-operations | Retrieve allowed ACL resources and operations
[**GetAcls**](AclsApi.md#GetAcls) | **Get** /api/v1/acls | List ACL bindings



## CreateAcl

> CreateAcl(ctx).AclBinding(aclBinding).Execute()

Create ACL binding



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
    aclBinding := *openapiclient.NewAclBinding(openapiclient.AclResourceType("GROUP"), "ResourceName_example", openapiclient.AclPatternType("LITERAL"), "User:user-123-abc", openapiclient.AclOperation("ALL"), openapiclient.AclPermissionType("ALLOW")) // AclBinding | ACL to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclsApi.CreateAcl(context.Background()).AclBinding(aclBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.CreateAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclBinding** | [**AclBinding**](AclBinding.md) | ACL to create. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAcls

> AclBindingListPage DeleteAcls(ctx).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Operation(operation).Permission(permission).Execute()

Delete ACL bindings



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
    resourceType := openapiclient.AclResourceTypeFilter("ANY") // AclResourceTypeFilter | ACL Resource Type Filter (optional)
    resourceName := "resourceName_example" // string | ACL Resource Name Filter (optional)
    patternType := openapiclient.AclPatternTypeFilter("LITERAL") // AclPatternTypeFilter | ACL Pattern Type Filter (optional)
    principal := "User:*" // string | ACL Principal Filter. Either a specific user or the wildcard user `User:*` may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided `principal` exactly. (optional) (default to "")
    operation := openapiclient.AclOperationFilter("ALL") // AclOperationFilter | ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not `ANY`. (optional)
    permission := openapiclient.AclPermissionTypeFilter("ALLOW") // AclPermissionTypeFilter | ACL Permission Type Filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclsApi.DeleteAcls(context.Background()).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Operation(operation).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.DeleteAcls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAcls`: AclBindingListPage
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.DeleteAcls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | [**AclResourceTypeFilter**](AclResourceTypeFilter.md) | ACL Resource Type Filter | 
 **resourceName** | **string** | ACL Resource Name Filter | 
 **patternType** | [**AclPatternTypeFilter**](AclPatternTypeFilter.md) | ACL Pattern Type Filter | 
 **principal** | **string** | ACL Principal Filter. Either a specific user or the wildcard user &#x60;User:*&#x60; may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided &#x60;principal&#x60; exactly. | [default to &quot;&quot;]
 **operation** | [**AclOperationFilter**](AclOperationFilter.md) | ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not &#x60;ANY&#x60;. | 
 **permission** | [**AclPermissionTypeFilter**](AclPermissionTypeFilter.md) | ACL Permission Type Filter | 

### Return type

[**AclBindingListPage**](AclBindingListPage.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAclResourceOperations

> map[string][]string GetAclResourceOperations(ctx).Execute()

Retrieve allowed ACL resources and operations



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
    resp, r, err := api_client.AclsApi.GetAclResourceOperations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.GetAclResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAclResourceOperations`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.GetAclResourceOperations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAclResourceOperationsRequest struct via the builder pattern


### Return type

[**map[string][]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcls

> AclBindingListPage GetAcls(ctx).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Operation(operation).Permission(permission).Page(page).Size(size).Order(order).OrderKey(orderKey).Execute()

List ACL bindings



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
    resourceType := openapiclient.AclResourceTypeFilter("ANY") // AclResourceTypeFilter | ACL Resource Type Filter (optional)
    resourceName := "resourceName_example" // string | ACL Resource Name Filter (optional)
    patternType := openapiclient.AclPatternTypeFilter("LITERAL") // AclPatternTypeFilter | ACL Pattern Type Filter (optional)
    principal := "User:*" // string | ACL Principal Filter. Either a specific user or the wildcard user `User:*` may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided `principal` exactly. (optional) (default to "")
    operation := openapiclient.AclOperationFilter("ALL") // AclOperationFilter | ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not `ANY`. (optional)
    permission := openapiclient.AclPermissionTypeFilter("ALLOW") // AclPermissionTypeFilter | ACL Permission Type Filter (optional)
    page := int32(56) // int32 | Page number (optional)
    size := int32(56) // int32 | Number of records per page (optional)
    order := openapiclient.SortDirection("asc") // SortDirection | Order items are sorted (optional)
    orderKey := openapiclient.AclBindingOrderKey("resourceType") // AclBindingOrderKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclsApi.GetAcls(context.Background()).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Operation(operation).Permission(permission).Page(page).Size(size).Order(order).OrderKey(orderKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.GetAcls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcls`: AclBindingListPage
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.GetAcls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | [**AclResourceTypeFilter**](AclResourceTypeFilter.md) | ACL Resource Type Filter | 
 **resourceName** | **string** | ACL Resource Name Filter | 
 **patternType** | [**AclPatternTypeFilter**](AclPatternTypeFilter.md) | ACL Pattern Type Filter | 
 **principal** | **string** | ACL Principal Filter. Either a specific user or the wildcard user &#x60;User:*&#x60; may be provided. - When fetching by a specific user, the results will also include ACL bindings that apply to all users. - When deleting, ACL bindings to be delete must match the provided &#x60;principal&#x60; exactly. | [default to &quot;&quot;]
 **operation** | [**AclOperationFilter**](AclOperationFilter.md) | ACL Operation Filter. The ACL binding operation provided should be valid for the resource type in the request, if not &#x60;ANY&#x60;. | 
 **permission** | [**AclPermissionTypeFilter**](AclPermissionTypeFilter.md) | ACL Permission Type Filter | 
 **page** | **int32** | Page number | 
 **size** | **int32** | Number of records per page | 
 **order** | [**SortDirection**](SortDirection.md) | Order items are sorted | 
 **orderKey** | [**AclBindingOrderKey**](AclBindingOrderKey.md) |  | 

### Return type

[**AclBindingListPage**](AclBindingListPage.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

