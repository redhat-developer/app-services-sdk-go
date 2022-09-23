# \ArtifactsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifact**](ArtifactsApi.md#CreateArtifact) | **Post** /groups/{groupId}/artifacts | Create artifact
[**DeleteArtifact**](ArtifactsApi.md#DeleteArtifact) | **Delete** /groups/{groupId}/artifacts/{artifactId} | Delete artifact
[**DeleteArtifactsInGroup**](ArtifactsApi.md#DeleteArtifactsInGroup) | **Delete** /groups/{groupId}/artifacts | Delete artifacts in group
[**GetContentByGlobalId**](ArtifactsApi.md#GetContentByGlobalId) | **Get** /ids/globalIds/{globalId} | Get artifact by global ID
[**GetContentByHash**](ArtifactsApi.md#GetContentByHash) | **Get** /ids/contentHashes/{contentHash}/ | Get artifact content by SHA-256 hash
[**GetContentById**](ArtifactsApi.md#GetContentById) | **Get** /ids/contentIds/{contentId}/ | Get artifact content by ID
[**GetLatestArtifact**](ArtifactsApi.md#GetLatestArtifact) | **Get** /groups/{groupId}/artifacts/{artifactId} | Get latest artifact
[**ListArtifactsInGroup**](ArtifactsApi.md#ListArtifactsInGroup) | **Get** /groups/{groupId}/artifacts | List artifacts in group
[**ReferencesByContentHash**](ArtifactsApi.md#ReferencesByContentHash) | **Get** /ids/contentHashes/{contentHash}/references | List artifact references by hash
[**ReferencesByContentId**](ArtifactsApi.md#ReferencesByContentId) | **Get** /ids/contentIds/{contentId}/references | List artifact references by content ID
[**ReferencesByGlobalId**](ArtifactsApi.md#ReferencesByGlobalId) | **Get** /ids/globalIds/{globalId}/references | Returns a list with all the references for the artifact with the given global id.
[**UpdateArtifact**](ArtifactsApi.md#UpdateArtifact) | **Put** /groups/{groupId}/artifacts/{artifactId} | Update artifact
[**UpdateArtifactState**](ArtifactsApi.md#UpdateArtifactState) | **Put** /groups/{groupId}/artifacts/{artifactId}/state | Update artifact state



## CreateArtifact

> ArtifactMetaData CreateArtifact(ctx, groupId).Body(body).XRegistryArtifactType(xRegistryArtifactType).XRegistryArtifactId(xRegistryArtifactId).XRegistryVersion(xRegistryVersion).IfExists(ifExists).Canonical(canonical).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryContentHash(xRegistryContentHash).XRegistryHashAlgorithm(xRegistryHashAlgorithm).Execute()

Create artifact



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
    groupId := "groupId_example" // string | Unique ID of an artifact group.
    body := os.NewFile(1234, "some_file") // *os.File | The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    xRegistryArtifactType := openapiclient.ArtifactType("AVRO") // ArtifactType | Specifies the type of the artifact being added. Possible values include:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) (optional)
    xRegistryArtifactId := "xRegistryArtifactId_example" // string | A client-provided, globally unique identifier for the new artifact. (optional)
    xRegistryVersion := "xRegistryVersion_example" // string | Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version `1`). (optional)
    ifExists := openapiclient.IfExists("FAIL") // IfExists | Set this option to instruct the server on what to do if the artifact already exists. (optional)
    canonical := true // bool | Used only when the `ifExists` query parameter is set to `RETURN_OR_UPDATE`, this parameter can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. (optional)
    xRegistryDescription := "xRegistryDescription_example" // string | Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryDescriptionEncoded := "xRegistryDescriptionEncoded_example" // string | Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryName := "xRegistryName_example" // string | Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryNameEncoded := "xRegistryNameEncoded_example" // string | Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryContentHash := "xRegistryContentHash_example" // string | Specifies the (optional) hash of the artifact to be verified. (optional)
    xRegistryHashAlgorithm := "xRegistryHashAlgorithm_example" // string | The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.CreateArtifact(context.Background(), groupId).Body(body).XRegistryArtifactType(xRegistryArtifactType).XRegistryArtifactId(xRegistryArtifactId).XRegistryVersion(xRegistryVersion).IfExists(ifExists).Canonical(canonical).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryContentHash(xRegistryContentHash).XRegistryHashAlgorithm(xRegistryHashAlgorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.CreateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifact`: ArtifactMetaData
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.CreateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique ID of an artifact group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 
 **xRegistryArtifactType** | [**ArtifactType**](ArtifactType.md) | Specifies the type of the artifact being added. Possible values include:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;) | 
 **xRegistryArtifactId** | **string** | A client-provided, globally unique identifier for the new artifact. | 
 **xRegistryVersion** | **string** | Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version &#x60;1&#x60;). | 
 **ifExists** | [**IfExists**](IfExists.md) | Set this option to instruct the server on what to do if the artifact already exists. | 
 **canonical** | **bool** | Used only when the &#x60;ifExists&#x60; query parameter is set to &#x60;RETURN_OR_UPDATE&#x60;, this parameter can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. | 
 **xRegistryDescription** | **string** | Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryDescriptionEncoded** | **string** | Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryName** | **string** | Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryNameEncoded** | **string** | Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryContentHash** | **string** | Specifies the (optional) hash of the artifact to be verified. | 
 **xRegistryHashAlgorithm** | **string** | The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) | 

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifact

> DeleteArtifact(ctx, groupId, artifactId).Execute()

Delete artifact



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
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.DeleteArtifact(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.DeleteArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRequest struct via the builder pattern


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


## DeleteArtifactsInGroup

> DeleteArtifactsInGroup(ctx, groupId).Execute()

Delete artifacts in group



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
    groupId := "groupId_example" // string | Unique ID of an artifact group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.DeleteArtifactsInGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.DeleteArtifactsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique ID of an artifact group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactsInGroupRequest struct via the builder pattern


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


## GetContentByGlobalId

> *os.File GetContentByGlobalId(ctx, globalId).Dereference(dereference).Execute()

Get artifact by global ID



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
    globalId := int64(789) // int64 | Global identifier for an artifact version.
    dereference := true // bool | Allows the user to specify if the content should be dereferenced when being returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.GetContentByGlobalId(context.Background(), globalId).Dereference(dereference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetContentByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentByGlobalId`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetContentByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalId** | **int64** | Global identifier for an artifact version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByGlobalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dereference** | **bool** | Allows the user to specify if the content should be dereferenced when being returned | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentByHash

> *os.File GetContentByHash(ctx, contentHash).Execute()

Get artifact content by SHA-256 hash



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
    contentHash := "contentHash_example" // string | SHA-256 content hash for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.GetContentByHash(context.Background(), contentHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetContentByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentByHash`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetContentByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentHash** | **string** | SHA-256 content hash for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentById

> *os.File GetContentById(ctx, contentId).Execute()

Get artifact content by ID



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
    contentId := int64(789) // int64 | Global identifier for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.GetContentById(context.Background(), contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetContentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentById`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **int64** | Global identifier for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestArtifact

> *os.File GetLatestArtifact(ctx, groupId, artifactId).Dereference(dereference).Execute()

Get latest artifact



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
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    dereference := true // bool | Allows the user to specify if the content should be dereferenced when being returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.GetLatestArtifact(context.Background(), groupId, artifactId).Dereference(dereference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetLatestArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestArtifact`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetLatestArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dereference** | **bool** | Allows the user to specify if the content should be dereferenced when being returned | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactsInGroup

> ArtifactSearchResults ListArtifactsInGroup(ctx, groupId).Limit(limit).Offset(offset).Order(order).Orderby(orderby).Execute()

List artifacts in group



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
    groupId := "groupId_example" // string | Unique ID of an artifact group.
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting the result set.  Defaults to 0. (optional)
    order := openapiclient.SortOrder("asc") // SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := openapiclient.SortBy("name") // SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ListArtifactsInGroup(context.Background(), groupId).Limit(limit).Offset(offset).Order(order).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ListArtifactsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactsInGroup`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ListArtifactsInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique ID of an artifact group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactsInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | 
 **offset** | **int32** | The number of artifacts to skip before starting the result set.  Defaults to 0. | 
 **order** | [**SortOrder**](SortOrder.md) | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | [**SortBy**](SortBy.md) | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 

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


## ReferencesByContentHash

> []ArtifactReference ReferencesByContentHash(ctx, contentHash).Execute()

List artifact references by hash



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
    contentHash := "contentHash_example" // string | SHA-256 content hash for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ReferencesByContentHash(context.Background(), contentHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ReferencesByContentHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentHash`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ReferencesByContentHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentHash** | **string** | SHA-256 content hash for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByContentHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ArtifactReference**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencesByContentId

> []ArtifactReference ReferencesByContentId(ctx, contentId).Execute()

List artifact references by content ID



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
    contentId := int64(789) // int64 | Global identifier for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ReferencesByContentId(context.Background(), contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ReferencesByContentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ReferencesByContentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **int64** | Global identifier for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByContentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ArtifactReference**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencesByGlobalId

> []ArtifactReference ReferencesByGlobalId(ctx, globalId).Execute()

Returns a list with all the references for the artifact with the given global id.



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
    globalId := int64(789) // int64 | Global identifier for an artifact version.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ReferencesByGlobalId(context.Background(), globalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ReferencesByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByGlobalId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ReferencesByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalId** | **int64** | Global identifier for an artifact version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByGlobalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ArtifactReference**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifact

> ArtifactMetaData UpdateArtifact(ctx, groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).Execute()

Update artifact



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
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    body := interface{}({"openapi":"3.0.2","info":{"title":"Empty API","version":"1.0.7","description":"An example API design using OpenAPI."},"paths":{"/widgets":{"get":{"responses":{"200":{"content":{"application/json":{"schema":{"type":"array","items":{"type":"string"}}}},"description":"All widgets"}},"summary":"Get widgets"}}},"components":{"schemas":{"Widget":{"title":"Root Type for Widget","description":"A sample data type.","type":"object","properties":{"property-1":{"type":"string"},"property-2":{"type":"boolean"}},"example":{"property-1":"value1","property-2":true}}}}}) // interface{} | The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    xRegistryVersion := "xRegistryVersion_example" // string | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. (optional)
    xRegistryName := "xRegistryName_example" // string | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryNameEncoded := "xRegistryNameEncoded_example" // string | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryDescription := "xRegistryDescription_example" // string | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryDescriptionEncoded := "xRegistryDescriptionEncoded_example" // string | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.UpdateArtifact(context.Background(), groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.UpdateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifact`: ArtifactMetaData
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.UpdateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **interface{}** | The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 
 **xRegistryVersion** | **string** | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. | 
 **xRegistryName** | **string** | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryNameEncoded** | **string** | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryDescription** | **string** | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryDescriptionEncoded** | **string** | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | 

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactState

> UpdateArtifactState(ctx, groupId, artifactId).UpdateState(updateState).Execute()

Update artifact state



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
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    updateState := *openapiclient.NewUpdateState(openapiclient.ArtifactState("ENABLED")) // UpdateState | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.UpdateArtifactState(context.Background(), groupId, artifactId).UpdateState(updateState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.UpdateArtifactState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateState** | [**UpdateState**](UpdateState.md) |  | 

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

