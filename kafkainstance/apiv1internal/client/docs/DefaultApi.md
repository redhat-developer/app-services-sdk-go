# \DefaultApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](DefaultApi.md#CreateTopic) | **Post** /topics | Creates a new topic
[**DeleteConsumerGroupById**](DefaultApi.md#DeleteConsumerGroupById) | **Delete** /consumer-groups/{consumerGroupId} | Delete a consumer group.
[**DeleteTopic**](DefaultApi.md#DeleteTopic) | **Delete** /topics/{topicName} | Deletes a  topic
[**GetConsumerGroupById**](DefaultApi.md#GetConsumerGroupById) | **Get** /consumer-groups/{consumerGroupId} | Get a single consumer group by its unique ID.
[**GetConsumerGroups**](DefaultApi.md#GetConsumerGroups) | **Get** /consumer-groups | List of consumer groups in the Kafka instance.
[**GetTopic**](DefaultApi.md#GetTopic) | **Get** /topics/{topicName} | Retrieves the topic with the specified name.
[**GetTopics**](DefaultApi.md#GetTopics) | **Get** /topics | List of topics
[**UpdateTopic**](DefaultApi.md#UpdateTopic) | **Patch** /topics/{topicName} | Updates the topic with the specified name.



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
    resp, r, err := api_client.DefaultApi.CreateTopic(context.Background()).NewTopicInput(newTopicInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTopic`: %v\n", resp)
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


## DeleteConsumerGroupById

> DeleteConsumerGroupById(ctx, consumerGroupId).Execute()

Delete a consumer group.



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
    consumerGroupId := "consumerGroupId_example" // string | The unique ID of the cobsumer group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConsumerGroupById(context.Background(), consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConsumerGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The unique ID of the cobsumer group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := api_client.DefaultApi.DeleteTopic(context.Background(), topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTopic``: %v\n", err)
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


## GetConsumerGroupById

> ConsumerGroup GetConsumerGroupById(ctx, consumerGroupId).Topic(topic).Execute()

Get a single consumer group by its unique ID.

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
    consumerGroupId := "consumerGroupId_example" // string | The unique ID of the consumer group
    topic := "{"groupId":"consumer_group_1","consumers":[{"groupId":"consumer_group_1","topic":"topic-1","partition":0,"memberId":"consumer_group_member1","offset":5,"lag":0,"logEndOffset":5},{"groupId":"consumer_group_1","topic":"topic-1","partition":1,"memberId":"consumer_group_member2","offset":3,"lag":0,"logEndOffset":3},{"groupId":"consumer_group_1","topic":"topic-1","partition":2,"memberId":"consumer_group_member3","offset":6,"lag":1,"logEndOffset":5}]}" // string | Filter consumer groups for a specific topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConsumerGroupById(context.Background(), consumerGroupId).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConsumerGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroupById`: ConsumerGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConsumerGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The unique ID of the consumer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topic** | **string** | Filter consumer groups for a specific topic | 

### Return type

[**ConsumerGroup**](ConsumerGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerGroups

> ConsumerGroupList GetConsumerGroups(ctx).Limit(limit).Offset(offset).Topic(topic).GroupIdFilter(groupIdFilter).Execute()

List of consumer groups in the Kafka instance.



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
    limit := int32(56) // int32 | Maximum number of consumer groups to returnd (optional)
    offset := int32(56) // int32 | The page offset when returning the list of consumer groups (optional)
    topic := "topic_example" // string | Return consumer groups for this topic (optional)
    groupIdFilter := "groupIdFilter_example" // string | Return the consumer groups where the ID begins with this value (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConsumerGroups(context.Background()).Limit(limit).Offset(offset).Topic(topic).GroupIdFilter(groupIdFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConsumerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroups`: ConsumerGroupList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConsumerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of consumer groups to returnd | 
 **offset** | **int32** | The page offset when returning the list of consumer groups | 
 **topic** | **string** | Return consumer groups for this topic | 
 **groupIdFilter** | **string** | Return the consumer groups where the ID begins with this value | 

### Return type

[**ConsumerGroupList**](ConsumerGroupList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    resp, r, err := api_client.DefaultApi.GetTopic(context.Background(), topicName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopic`: %v\n", resp)
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

> TopicsList GetTopics(ctx).Limit(limit).Filter(filter).Offset(offset).Order(order).Execute()

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
    limit := int32(56) // int32 | Maximum number of topics to return (optional)
    filter := "filter_example" // string | Filter to apply when returning the list of topics (optional)
    offset := int32(56) // int32 | The page offset when returning the limit of requested topics. (optional)
    order := "order_example" // string | Order of the items sorting. If \"asc\" is set as a value, ascending order is used, descending otherwise. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopics(context.Background()).Limit(limit).Filter(filter).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopics`: TopicsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of topics to return | 
 **filter** | **string** | Filter to apply when returning the list of topics | 
 **offset** | **int32** | The page offset when returning the limit of requested topics. | 
 **order** | **string** | Order of the items sorting. If \&quot;asc\&quot; is set as a value, ascending order is used, descending otherwise. | 

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
    resp, r, err := api_client.DefaultApi.UpdateTopic(context.Background(), topicName).UpdateTopicInput(updateTopicInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTopic`: %v\n", resp)
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

