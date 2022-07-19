# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchArtifacts**](SearchApi.md#SearchArtifacts) | **Get** /search/artifacts | Search for artifacts
[**SearchArtifactsByContent**](SearchApi.md#SearchArtifactsByContent) | **Post** /search/artifacts | Search for artifacts by content



## SearchArtifacts

> ArtifactSearchResults SearchArtifacts(ctx).Name(name).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Labels(labels).Properties(properties).Description(description).Group(group).GlobalId(globalId).ContentId(contentId).Execute()

Search for artifacts



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
    name := "name_example" // string | Filter by artifact name. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. (optional) (default to 0)
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional) (default to 20)
    order := openapiclient.SortOrder("asc") // SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := openapiclient.SortBy("name") // SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)
    labels := []string{"Inner_example"} // []string | Filter by label.  Include one or more label to only return artifacts containing all of the specified labels. (optional)
    properties := []string{"Inner_example"} // []string | Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example `properties=foo:bar` will return only artifacts with a custom property named `foo` and value `bar`. (optional)
    description := "description_example" // string | Filter by description. (optional)
    group := "group_example" // string | Filter by artifact group. (optional)
    globalId := int64(789) // int64 | Filter by globalId. (optional)
    contentId := int64(789) // int64 | Filter by contentId. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchArtifacts(context.Background()).Name(name).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Labels(labels).Properties(properties).Description(description).Group(group).GlobalId(globalId).ContentId(contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchArtifacts`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by artifact name. | 
 **offset** | **int32** | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. | [default to 0]
 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | [default to 20]
 **order** | [**SortOrder**](SortOrder.md) | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | [**SortBy**](SortBy.md) | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 
 **labels** | **[]string** | Filter by label.  Include one or more label to only return artifacts containing all of the specified labels. | 
 **properties** | **[]string** | Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example &#x60;properties&#x3D;foo:bar&#x60; will return only artifacts with a custom property named &#x60;foo&#x60; and value &#x60;bar&#x60;. | 
 **description** | **string** | Filter by description. | 
 **group** | **string** | Filter by artifact group. | 
 **globalId** | **int64** | Filter by globalId. | 
 **contentId** | **int64** | Filter by contentId. | 

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchArtifactsByContent

> ArtifactSearchResults SearchArtifactsByContent(ctx).Body(body).Canonical(canonical).ArtifactType(artifactType).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Execute()

Search for artifacts by content



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
    body := os.NewFile(1234, "some_file") // *os.File | The content to search for.
    canonical := true // bool | Parameter that can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the `artifactType` query parameter. (optional)
    artifactType := openapiclient.ArtifactType("AVRO") // ArtifactType | Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the `canonical` query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. (optional) (default to 0)
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional) (default to 20)
    order := "order_example" // string | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := "orderby_example" // string | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchArtifactsByContent(context.Background()).Body(body).Canonical(canonical).ArtifactType(artifactType).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchArtifactsByContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchArtifactsByContent`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchArtifactsByContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtifactsByContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | The content to search for. | 
 **canonical** | **bool** | Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the &#x60;artifactType&#x60; query parameter. | 
 **artifactType** | [**ArtifactType**](ArtifactType.md) | Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the &#x60;canonical&#x60; query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts. | 
 **offset** | **int32** | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. | [default to 0]
 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | [default to 20]
 **order** | **string** | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | **string** | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

