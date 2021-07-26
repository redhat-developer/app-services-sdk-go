# \TopicsApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](TopicsApi.md#CreateTopic) | **Post** /topics | Creates a new topic
[**DeleteTopic**](TopicsApi.md#DeleteTopic) | **Delete** /topics/{topicName} | Deletes a  topic
[**GetTopic**](TopicsApi.md#GetTopic) | **Get** /topics/{topicName} | Retrieves the topic with the specified name.
[**GetTopics**](TopicsApi.md#GetTopics) | **Get** /topics | List of topics
[**UpdateTopic**](TopicsApi.md#UpdateTopic) | **Patch** /topics/{topicName} | Updates the topic with the specified name.



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
    newTopicInput := *openapiclient.NewNewTopicInput("Name_example", *openapiclient.NewTopicSettings(int32(123))) // NewTopicInput | Topic to create.

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

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopic

> DeleteTopic(ctx, topicName).Execute()

Deletes a  topic



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
    topicName := "topicName_example" // string | The topic name to delete.

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
**topicName** | **string** | The topic name to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopic

> Topic GetTopic(ctx, topicName).Execute()

Retrieves the topic with the specified name.



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
    topicName := "topicName_example" // string | The topic name to retrieve.

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
**topicName** | **string** | The topic name to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Topic**](Topic.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopics

> TopicsList GetTopics(ctx).Offset(offset).Limit(limit).Size(size).Filter(filter).Page(page).Order(order).OrderKey(orderKey).Execute()

List of topics



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
    offset := int32(56) // int32 | The page offset (optional)
    limit := int32(56) // int32 | Maximum number of topics to return (optional)
    size := int32(56) // int32 | Maximum number of topics to return on single page (optional)
    filter := "filter_example" // string | Filter to apply when returning the list of topics (optional)
    page := int32(56) // int32 | The page when returning the limit of requested topics. (optional)
    order := "order_example" // string | Order of the items sorting. Ascending order is used as default. (optional)
    orderKey := "orderKey_example" // string | Order key to sort the topics by. (optional)

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
 **offset** | **int32** | The page offset | 
 **limit** | **int32** | Maximum number of topics to return | 
 **size** | **int32** | Maximum number of topics to return on single page | 
 **filter** | **string** | Filter to apply when returning the list of topics | 
 **page** | **int32** | The page when returning the limit of requested topics. | 
 **order** | **string** | Order of the items sorting. Ascending order is used as default. | 
 **orderKey** | **string** | Order key to sort the topics by. | 

### Return type

[**TopicsList**](TopicsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopic

> Topic UpdateTopic(ctx, topicName).UpdateTopicInput(updateTopicInput).Execute()

Updates the topic with the specified name.



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
    topicName := "topicName_example" // string | The topic name which is its unique id.
    updateTopicInput := *openapiclient.NewUpdateTopicInput() // UpdateTopicInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TopicsApi.UpdateTopic(context.Background(), topicName).UpdateTopicInput(updateTopicInput).Execute()
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
**topicName** | **string** | The topic name which is its unique id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTopicInput** | [**UpdateTopicInput**](UpdateTopicInput.md) |  | 

### Return type

[**Topic**](Topic.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

