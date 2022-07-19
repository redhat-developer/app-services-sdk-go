# \GlobalRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalRule**](GlobalRulesApi.md#CreateGlobalRule) | **Post** /admin/rules | Create global rule
[**DeleteAllGlobalRules**](GlobalRulesApi.md#DeleteAllGlobalRules) | **Delete** /admin/rules | Delete all global rules
[**DeleteGlobalRule**](GlobalRulesApi.md#DeleteGlobalRule) | **Delete** /admin/rules/{rule} | Delete global rule
[**GetGlobalRuleConfig**](GlobalRulesApi.md#GetGlobalRuleConfig) | **Get** /admin/rules/{rule} | Get global rule configuration
[**ListGlobalRules**](GlobalRulesApi.md#ListGlobalRules) | **Get** /admin/rules | List global rules
[**UpdateGlobalRuleConfig**](GlobalRulesApi.md#UpdateGlobalRuleConfig) | **Put** /admin/rules/{rule} | Update global rule configuration



## CreateGlobalRule

> CreateGlobalRule(ctx).Rule(rule).Execute()

Create global rule



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
    rule := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRulesApi.CreateGlobalRule(context.Background()).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.CreateGlobalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rule** | [**Rule**](Rule.md) |  | 

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


## DeleteAllGlobalRules

> DeleteAllGlobalRules(ctx).Execute()

Delete all global rules



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
    resp, r, err := api_client.GlobalRulesApi.DeleteAllGlobalRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.DeleteAllGlobalRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllGlobalRulesRequest struct via the builder pattern


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


## DeleteGlobalRule

> DeleteGlobalRule(ctx, rule).Execute()

Delete global rule



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRulesApi.DeleteGlobalRule(context.Background(), rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.DeleteGlobalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalRuleRequest struct via the builder pattern


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


## GetGlobalRuleConfig

> Rule GetGlobalRuleConfig(ctx, rule).Execute()

Get global rule configuration



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRulesApi.GetGlobalRuleConfig(context.Background(), rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.GetGlobalRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `GlobalRulesApi.GetGlobalRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalRules

> []RuleType ListGlobalRules(ctx).Execute()

List global rules



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
    resp, r, err := api_client.GlobalRulesApi.ListGlobalRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.ListGlobalRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalRules`: []RuleType
    fmt.Fprintf(os.Stdout, "Response from `GlobalRulesApi.ListGlobalRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalRulesRequest struct via the builder pattern


### Return type

[**[]RuleType**](RuleType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalRuleConfig

> Rule UpdateGlobalRuleConfig(ctx, rule).Rule2(rule2).Execute()

Update global rule configuration



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.
    rule2 := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalRulesApi.UpdateGlobalRuleConfig(context.Background(), rule).Rule2(rule2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalRulesApi.UpdateGlobalRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `GlobalRulesApi.UpdateGlobalRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rule2** | [**Rule**](Rule.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

