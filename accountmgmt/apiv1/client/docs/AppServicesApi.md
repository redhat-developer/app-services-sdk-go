# \AppServicesApi

All URIs are relative to *http://localhost:14321*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsMgmtV1AccessTokenPost**](AppServicesApi.md#ApiAccountsMgmtV1AccessTokenPost) | **Post** /api/accounts_mgmt/v1/access_token | Return access token generated from registries in docker format
[**ApiAccountsMgmtV1CurrentAccountGet**](AppServicesApi.md#ApiAccountsMgmtV1CurrentAccountGet) | **Get** /api/accounts_mgmt/v1/current_account | Get the authenticated account
[**ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet**](AppServicesApi.md#ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/quota_cost | Returns a summary of quota cost
[**ApiAuthorizationsV1SelfTermsReviewPost**](AppServicesApi.md#ApiAuthorizationsV1SelfTermsReviewPost) | **Post** /api/authorizations/v1/self_terms_review | Review your status of Terms



## ApiAccountsMgmtV1AccessTokenPost

> AccessTokenCfg ApiAccountsMgmtV1AccessTokenPost(ctx).Execute()

Return access token generated from registries in docker format

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
    resp, r, err := api_client.AppServicesApi.ApiAccountsMgmtV1AccessTokenPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServicesApi.ApiAccountsMgmtV1AccessTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccessTokenPost`: AccessTokenCfg
    fmt.Fprintf(os.Stdout, "Response from `AppServicesApi.ApiAccountsMgmtV1AccessTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccessTokenPostRequest struct via the builder pattern


### Return type

[**AccessTokenCfg**](AccessTokenCfg.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1CurrentAccountGet

> Account ApiAccountsMgmtV1CurrentAccountGet(ctx).FetchLabels(fetchLabels).Execute()

Get the authenticated account

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
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppServicesApi.ApiAccountsMgmtV1CurrentAccountGet(context.Background()).FetchLabels(fetchLabels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServicesApi.ApiAccountsMgmtV1CurrentAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1CurrentAccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `AppServicesApi.ApiAccountsMgmtV1CurrentAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1CurrentAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 

### Return type

[**Account**](Account.md)

### Authorization

[AccessToken](../README.md#AccessToken), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet

> QuotaCostList ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).Search(search).FetchRelatedResources(fetchRelatedResources).ForceRecalc(forceRecalc).FetchCloudAccounts(fetchCloudAccounts).Execute()

Returns a summary of quota cost

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
    orgId := "orgId_example" // string | The id of organization
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    fetchRelatedResources := true // bool | If true, includes the related resources in the output. Could slow request response time. (optional)
    forceRecalc := true // bool | If true, includes that ConsumedQuota should be recalculated. (optional)
    fetchCloudAccounts := true // bool | If true, includes the marketplace cloud accounts in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppServicesApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(context.Background(), orgId).Search(search).FetchRelatedResources(fetchRelatedResources).ForceRecalc(forceRecalc).FetchCloudAccounts(fetchCloudAccounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServicesApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet`: QuotaCostList
    fmt.Fprintf(os.Stdout, "Response from `AppServicesApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **fetchRelatedResources** | **bool** | If true, includes the related resources in the output. Could slow request response time. | 
 **forceRecalc** | **bool** | If true, includes that ConsumedQuota should be recalculated. | 
 **fetchCloudAccounts** | **bool** | If true, includes the marketplace cloud accounts in the output. Could slow request response time. | 

### Return type

[**QuotaCostList**](QuotaCostList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1SelfTermsReviewPost

> TermsReviewResponse ApiAuthorizationsV1SelfTermsReviewPost(ctx).SelfTermsReview(selfTermsReview).Execute()

Review your status of Terms

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
    selfTermsReview := *openapiclient.NewSelfTermsReview() // SelfTermsReview | Data to check self terms for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppServicesApi.ApiAuthorizationsV1SelfTermsReviewPost(context.Background()).SelfTermsReview(selfTermsReview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServicesApi.ApiAuthorizationsV1SelfTermsReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1SelfTermsReviewPost`: TermsReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServicesApi.ApiAuthorizationsV1SelfTermsReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1SelfTermsReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfTermsReview** | [**SelfTermsReview**](SelfTermsReview.md) | Data to check self terms for | 

### Return type

[**TermsReviewResponse**](TermsReviewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

