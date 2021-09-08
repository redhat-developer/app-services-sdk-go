# \GroupsApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConsumerGroupById**](GroupsApi.md#DeleteConsumerGroupById) | **Delete** /consumer-groups/{consumerGroupId} | Delete a consumer group.
[**GetConsumerGroupById**](GroupsApi.md#GetConsumerGroupById) | **Get** /consumer-groups/{consumerGroupId} | Get a single consumer group by its unique ID.
[**GetConsumerGroups**](GroupsApi.md#GetConsumerGroups) | **Get** /consumer-groups | List of consumer groups in the Kafka instance.
[**ResetConsumerGroupOffset**](GroupsApi.md#ResetConsumerGroupOffset) | **Post** /consumer-groups/{consumerGroupId}/reset-offset | Reset the offset for a consumer group.



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
    resp, r, err := api_client.GroupsApi.DeleteConsumerGroupById(context.Background(), consumerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteConsumerGroupById``: %v\n", err)
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


## GetConsumerGroupById

> ConsumerGroup GetConsumerGroupById(ctx, consumerGroupId).Order(order).OrderKey(orderKey).PartitionFilter(partitionFilter).Topic(topic).Execute()

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
    order := "order_example" // string | Order of the items sorting. Ascending order is used as default. (optional)
    orderKey := "orderKey_example" // string | Order key to sort the topics by. (optional)
    partitionFilter := int32(56) // int32 | Value of partition to include. Value -1 means filter is not active. (optional)
    topic := "{"groupId":"consumer_group_1","consumers":[{"groupId":"consumer_group_1","topic":"topic-1","partition":0,"memberId":"consumer_group_member1","offset":5,"lag":0,"logEndOffset":5},{"groupId":"consumer_group_1","topic":"topic-1","partition":1,"memberId":"consumer_group_member2","offset":3,"lag":0,"logEndOffset":3},{"groupId":"consumer_group_1","topic":"topic-1","partition":2,"memberId":"consumer_group_member3","offset":6,"lag":1,"logEndOffset":5}]}" // string | Filter consumer groups for a specific topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetConsumerGroupById(context.Background(), consumerGroupId).Order(order).OrderKey(orderKey).PartitionFilter(partitionFilter).Topic(topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetConsumerGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroupById`: ConsumerGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetConsumerGroupById`: %v\n", resp)
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

 **order** | **string** | Order of the items sorting. Ascending order is used as default. | 
 **orderKey** | **string** | Order key to sort the topics by. | 
 **partitionFilter** | **int32** | Value of partition to include. Value -1 means filter is not active. | 
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

> ConsumerGroupList GetConsumerGroups(ctx).Offset(offset).Limit(limit).Size(size).Page(page).Topic(topic).GroupIdFilter(groupIdFilter).Order(order).OrderKey(orderKey).Execute()

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
    offset := int32(56) // int32 | The page offset (optional)
    limit := int32(56) // int32 | Maximum number of consumer groups to return (optional)
    size := int32(56) // int32 | Maximum number of consumer groups to return on single page (optional)
    page := int32(56) // int32 | The page when returning the list of consumer groups (optional)
    topic := "topic_example" // string | Return consumer groups where the topic name contains with this value (optional)
    groupIdFilter := "groupIdFilter_example" // string | Return the consumer groups where the ID contains with this value (optional)
    order := "order_example" // string | Order of the consumer groups sorting. Ascending order is used as default. (optional)
    orderKey := "orderKey_example" // string | Order key to sort the items by. Only the value 'name' is currently applicable. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetConsumerGroups(context.Background()).Offset(offset).Limit(limit).Size(size).Page(page).Topic(topic).GroupIdFilter(groupIdFilter).Order(order).OrderKey(orderKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetConsumerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroups`: ConsumerGroupList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetConsumerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The page offset | 
 **limit** | **int32** | Maximum number of consumer groups to return | 
 **size** | **int32** | Maximum number of consumer groups to return on single page | 
 **page** | **int32** | The page when returning the list of consumer groups | 
 **topic** | **string** | Return consumer groups where the topic name contains with this value | 
 **groupIdFilter** | **string** | Return the consumer groups where the ID contains with this value | 
 **order** | **string** | Order of the consumer groups sorting. Ascending order is used as default. | 
 **orderKey** | **string** | Order key to sort the items by. Only the value &#39;name&#39; is currently applicable. | 

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


## ResetConsumerGroupOffset

> ConsumerGroupResetOffsetResult ResetConsumerGroupOffset(ctx, consumerGroupId).ConsumerGroupResetOffsetParameters(consumerGroupResetOffsetParameters).Execute()

Reset the offset for a consumer group.



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
    consumerGroupId := "consumerGroupId_example" // string | The ID of the consumer group.
    consumerGroupResetOffsetParameters := *openapiclient.NewConsumerGroupResetOffsetParameters("Offset_example") // ConsumerGroupResetOffsetParameters | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.ResetConsumerGroupOffset(context.Background(), consumerGroupId).ConsumerGroupResetOffsetParameters(consumerGroupResetOffsetParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.ResetConsumerGroupOffset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetConsumerGroupOffset`: ConsumerGroupResetOffsetResult
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.ResetConsumerGroupOffset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The ID of the consumer group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConsumerGroupOffsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerGroupResetOffsetParameters** | [**ConsumerGroupResetOffsetParameters**](ConsumerGroupResetOffsetParameters.md) |  | 

### Return type

[**ConsumerGroupResetOffsetResult**](ConsumerGroupResetOffsetResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

