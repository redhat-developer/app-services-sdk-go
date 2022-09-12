# \TopicsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](TopicsApi.md#CreateTopic) | **Post** /api/v1/topics | Creates a new topic
[**DeleteTopic**](TopicsApi.md#DeleteTopic) | **Delete** /api/v1/topics/{topicName} | Deletes a topic
[**GetTopic**](TopicsApi.md#GetTopic) | **Get** /api/v1/topics/{topicName} | Retrieves a single topic
[**GetTopics**](TopicsApi.md#GetTopics) | **Get** /api/v1/topics | Retrieves a list of topics
[**UpdateTopic**](TopicsApi.md#UpdateTopic) | **Patch** /api/v1/topics/{topicName} | Updates a single topic



## CreateTopic

> Topic CreateTopic(ctx).NewTopicInput(newTopicInput).Execute()

Creates a new topic



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
    newTopicInput := *openapiclient.NewNewTopicInput("Name_example", *openapiclient.NewTopicSettings()) // NewTopicInput | Topic to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.CreateTopic(context.Background()).NewTopicInput(newTopicInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.CreateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.CreateTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newTopicInput** | [**NewTopicInput**](NewTopicInput.md) | Topic to create. | 

### Return type

[**Topic**](Topic.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopic

> DeleteTopic(ctx, topicName).Execute()

Deletes a topic



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
    topicName := "topicName_example" // string | Name of the topic to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.DeleteTopic(context.Background(), topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.DeleteTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Name of the topic to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopic

> Topic GetTopic(ctx, topicName).Execute()

Retrieves a single topic



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
    topicName := "topicName_example" // string | Name of the topic to describe

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.GetTopic(context.Background(), topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Name of the topic to describe | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Topic**](Topic.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopics

> TopicsList GetTopics(ctx).Offset(offset).Limit(limit).Size(size).Filter(filter).Page(page).Order(order).OrderKey(orderKey).Execute()

Retrieves a list of topics



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
    offset := int32(56) // int32 | Offset of the first record to return, zero-based (optional)
    limit := int32(56) // int32 | Maximum number of records to return (optional)
    size := int32(56) // int32 | Number of records per page (optional)
    filter := "filter_example" // string | Filter to apply when returning the list of topics (optional)
    page := int32(56) // int32 | Page number (optional)
    order := openapiclient.SortDirection("asc") // SortDirection | Order items are sorted (optional)
    orderKey := openapiclient.TopicOrderKey("name") // TopicOrderKey | Order key to sort the topics by. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.GetTopics(context.Background()).Offset(offset).Limit(limit).Size(size).Filter(filter).Page(page).Order(order).OrderKey(orderKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.GetTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopics`: TopicsList
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.GetTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset of the first record to return, zero-based | 
 **limit** | **int32** | Maximum number of records to return | 
 **size** | **int32** | Number of records per page | 
 **filter** | **string** | Filter to apply when returning the list of topics | 
 **page** | **int32** | Page number | 
 **order** | [**SortDirection**](SortDirection.md) | Order items are sorted | 
 **orderKey** | [**TopicOrderKey**](TopicOrderKey.md) | Order key to sort the topics by. | 

### Return type

[**TopicsList**](TopicsList.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopic

> Topic UpdateTopic(ctx, topicName).TopicSettings(topicSettings).Execute()

Updates a single topic



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
    topicName := "topicName_example" // string | Name of the topic to update
    topicSettings := *openapiclient.NewTopicSettings() // TopicSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.UpdateTopic(context.Background(), topicName).TopicSettings(topicSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.UpdateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.UpdateTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Name of the topic to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topicSettings** | [**TopicSettings**](TopicSettings.md) |  | 

### Return type

[**Topic**](Topic.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

