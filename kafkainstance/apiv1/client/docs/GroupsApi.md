# \GroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConsumerGroupById**](GroupsApi.md#DeleteConsumerGroupById) | **Delete** /api/v1/consumer-groups/{consumerGroupId} | Delete a consumer group.
[**GetConsumerGroupById**](GroupsApi.md#GetConsumerGroupById) | **Get** /api/v1/consumer-groups/{consumerGroupId} | Get a single consumer group by its unique ID.
[**GetConsumerGroups**](GroupsApi.md#GetConsumerGroups) | **Get** /api/v1/consumer-groups | List of consumer groups in the Kafka instance.
[**ResetConsumerGroupOffset**](GroupsApi.md#ResetConsumerGroupOffset) | **Post** /api/v1/consumer-groups/{consumerGroupId}/reset-offset | Reset the offset for a consumer group.



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
    consumerGroupId := "consumerGroupId_example" // string | Consumer group identifier

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
**consumerGroupId** | **string** | Consumer group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerGroupByIdRequest struct via the builder pattern


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
    consumerGroupId := "consumerGroupId_example" // string | Consumer group identifier
    order := openapiclient.SortDirection("asc") // SortDirection | Order items are sorted (optional)
    orderKey := openapiclient.ConsumerGroupDescriptionOrderKey("offset") // ConsumerGroupDescriptionOrderKey |  (optional)
    partitionFilter := int32(56) // int32 | Value of partition to include. Value -1 means filter is not active. (optional)
    topic := "topic_example" // string | Filter consumer groups for a specific topic (optional)

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
**consumerGroupId** | **string** | Consumer group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**SortDirection**](SortDirection.md) | Order items are sorted | 
 **orderKey** | [**ConsumerGroupDescriptionOrderKey**](ConsumerGroupDescriptionOrderKey.md) |  | 
 **partitionFilter** | **int32** | Value of partition to include. Value -1 means filter is not active. | 
 **topic** | **string** | Filter consumer groups for a specific topic | 

### Return type

[**ConsumerGroup**](ConsumerGroup.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

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
    offset := int32(56) // int32 | Offset of the first record to return, zero-based (optional)
    limit := int32(56) // int32 | Maximum number of records to return (optional)
    size := int32(56) // int32 | Number of records per page (optional)
    page := int32(56) // int32 | Page number (optional)
    topic := "topic_example" // string | Return consumer groups where the topic name contains this value (optional)
    groupIdFilter := "groupIdFilter_example" // string | Return the consumer groups where the ID contains this value (optional)
    order := openapiclient.SortDirection("asc") // SortDirection | Order items are sorted (optional)
    orderKey := openapiclient.ConsumerGroupOrderKey("name") // ConsumerGroupOrderKey |  (optional)

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
 **offset** | **int32** | Offset of the first record to return, zero-based | 
 **limit** | **int32** | Maximum number of records to return | 
 **size** | **int32** | Number of records per page | 
 **page** | **int32** | Page number | 
 **topic** | **string** | Return consumer groups where the topic name contains this value | 
 **groupIdFilter** | **string** | Return the consumer groups where the ID contains this value | 
 **order** | [**SortDirection**](SortDirection.md) | Order items are sorted | 
 **orderKey** | [**ConsumerGroupOrderKey**](ConsumerGroupOrderKey.md) |  | 

### Return type

[**ConsumerGroupList**](ConsumerGroupList.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

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
    consumerGroupId := "consumerGroupId_example" // string | Consumer group identifier
    consumerGroupResetOffsetParameters := *openapiclient.NewConsumerGroupResetOffsetParameters(openapiclient.OffsetType("timestamp")) // ConsumerGroupResetOffsetParameters | 

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
**consumerGroupId** | **string** | Consumer group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConsumerGroupOffsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerGroupResetOffsetParameters** | [**ConsumerGroupResetOffsetParameters**](ConsumerGroupResetOffsetParameters.md) |  | 

### Return type

[**ConsumerGroupResetOffsetResult**](ConsumerGroupResetOffsetResult.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

