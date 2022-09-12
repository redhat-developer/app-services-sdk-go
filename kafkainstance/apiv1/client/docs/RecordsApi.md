# \RecordsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumeRecords**](RecordsApi.md#ConsumeRecords) | **Get** /api/v1/topics/{topicName}/records | Consume records from a topic
[**ProduceRecord**](RecordsApi.md#ProduceRecord) | **Post** /api/v1/topics/{topicName}/records | Send a record to a topic



## ConsumeRecords

> RecordList ConsumeRecords(ctx, topicName).Include(include).Limit(limit).MaxValueLength(maxValueLength).Offset(offset).Partition(partition).Timestamp(timestamp).Execute()

Consume records from a topic



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
    topicName := "topicName_example" // string | Topic name
    include := []openapiclient.RecordIncludedProperty{openapiclient.RecordIncludedProperty("partition")} // []RecordIncludedProperty | List of properties to include for each record in the response (optional)
    limit := int32(56) // int32 | Limit the number of records fetched and returned (optional)
    maxValueLength := int32(56) // int32 | Maximum length of string values returned in the response. Values with a length that exceeds this parameter will be truncated. When this parameter is not included in the request, the full string values will be returned. (optional)
    offset := int32(56) // int32 | Retrieve messages with an offset equal to or greater than this offset. If both `timestamp` and `offset` are requested, `timestamp` is given preference. (optional)
    partition := int32(56) // int32 | Retrieve messages only from this partition (optional)
    timestamp := TODO // interface{} | Retrieve messages with a timestamp equal to or later than this timestamp. If both `timestamp` and `offset` are requested, `timestamp` is given preference. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsApi.ConsumeRecords(context.Background(), topicName).Include(include).Limit(limit).MaxValueLength(maxValueLength).Offset(offset).Partition(partition).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ConsumeRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConsumeRecords`: RecordList
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ConsumeRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumeRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | [**[]RecordIncludedProperty**](RecordIncludedProperty.md) | List of properties to include for each record in the response | 
 **limit** | **int32** | Limit the number of records fetched and returned | 
 **maxValueLength** | **int32** | Maximum length of string values returned in the response. Values with a length that exceeds this parameter will be truncated. When this parameter is not included in the request, the full string values will be returned. | 
 **offset** | **int32** | Retrieve messages with an offset equal to or greater than this offset. If both &#x60;timestamp&#x60; and &#x60;offset&#x60; are requested, &#x60;timestamp&#x60; is given preference. | 
 **partition** | **int32** | Retrieve messages only from this partition | 
 **timestamp** | [**interface{}**](interface{}.md) | Retrieve messages with a timestamp equal to or later than this timestamp. If both &#x60;timestamp&#x60; and &#x60;offset&#x60; are requested, &#x60;timestamp&#x60; is given preference. | 

### Return type

[**RecordList**](RecordList.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProduceRecord

> Record ProduceRecord(ctx, topicName).Record(record).Execute()

Send a record to a topic



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
    topicName := "topicName_example" // string | Topic name
    record := *openapiclient.NewRecord("Value_example") // Record | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsApi.ProduceRecord(context.Background(), topicName).Record(record).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ProduceRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProduceRecord`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ProduceRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | Topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiProduceRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **record** | [**Record**](Record.md) |  | 

### Return type

[**Record**](Record.md)

### Authorization

[Bearer](../README.md#Bearer), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

