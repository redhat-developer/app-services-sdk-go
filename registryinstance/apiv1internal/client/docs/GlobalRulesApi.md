# \GlobalRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGlobalRule**](GlobalRulesApi.md#DeleteGlobalRule) | **Delete** /admin/rules/{rule} | Delete global rule
[**ListGlobalRules**](GlobalRulesApi.md#ListGlobalRules) | **Get** /admin/rules | List global rules



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

