# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafka**](DefaultApi.md#CreateKafka) | **Post** /api/kafkas_mgmt/v1/kafkas | Creates a Kafka request
[**DeleteKafkaById**](DefaultApi.md#DeleteKafkaById) | **Delete** /api/kafkas_mgmt/v1/kafkas/{id} | Deletes a Kafka request by ID
[**FederateMetrics**](DefaultApi.md#FederateMetrics) | **Get** /api/kafkas_mgmt/v1/kafkas/{id}/metrics/federate | Returns all metrics in scrapeable format for a given kafka id
[**GetCloudProviderRegions**](DefaultApi.md#GetCloudProviderRegions) | **Get** /api/kafkas_mgmt/v1/cloud_providers/{id}/regions | Returns the list of supported regions of the supported cloud provider
[**GetCloudProviders**](DefaultApi.md#GetCloudProviders) | **Get** /api/kafkas_mgmt/v1/cloud_providers | Returns the list of supported cloud providers
[**GetKafkaById**](DefaultApi.md#GetKafkaById) | **Get** /api/kafkas_mgmt/v1/kafkas/{id} | Returns a Kafka request by ID
[**GetKafkas**](DefaultApi.md#GetKafkas) | **Get** /api/kafkas_mgmt/v1/kafkas | Returns a list of Kafka requests
[**GetMetricsByInstantQuery**](DefaultApi.md#GetMetricsByInstantQuery) | **Get** /api/kafkas_mgmt/v1/kafkas/{id}/metrics/query | Returns metrics with instant query by Kafka ID
[**GetMetricsByRangeQuery**](DefaultApi.md#GetMetricsByRangeQuery) | **Get** /api/kafkas_mgmt/v1/kafkas/{id}/metrics/query_range | Returns metrics with timeseries range query by Kafka ID
[**GetServiceStatus**](DefaultApi.md#GetServiceStatus) | **Get** /api/kafkas_mgmt/v1/status | Returns the status of resources, such as whether maximum service capacity has been reached
[**GetVersionMetadata**](DefaultApi.md#GetVersionMetadata) | **Get** /api/kafkas_mgmt/v1 | Returns the version metadata
[**UpdateKafkaById**](DefaultApi.md#UpdateKafkaById) | **Patch** /api/kafkas_mgmt/v1/kafkas/{id} | Update a Kafka instance by id



## CreateKafka

> KafkaRequest CreateKafka(ctx).Async(async).KafkaRequestPayload(kafkaRequestPayload).Execute()

Creates a Kafka request

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
    async := true // bool | Perform the action in an asynchronous manner
    kafkaRequestPayload := *openapiclient.NewKafkaRequestPayload("Name_example") // KafkaRequestPayload | Kafka data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateKafka(context.Background()).Async(async).KafkaRequestPayload(kafkaRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateKafka``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafka`: KafkaRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateKafka`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Perform the action in an asynchronous manner | 
 **kafkaRequestPayload** | [**KafkaRequestPayload**](KafkaRequestPayload.md) | Kafka data | 

### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaById

> Error DeleteKafkaById(ctx, id).Async(async).Execute()

Deletes a Kafka request by ID

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
    id := "id_example" // string | The ID of record
    async := true // bool | Perform the action in an asynchronous manner

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKafkaById(context.Background(), id).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKafkaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKafkaById`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteKafkaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Perform the action in an asynchronous manner | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FederateMetrics

> string FederateMetrics(ctx, id).Execute()

Returns all metrics in scrapeable format for a given kafka id

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
    id := "id_example" // string | The ID of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FederateMetrics(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FederateMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FederateMetrics`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FederateMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiFederateMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviderRegions

> CloudRegionList GetCloudProviderRegions(ctx, id).Page(page).Size(size).InstanceType(instanceType).Execute()

Returns the list of supported regions of the supported cloud provider

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
    id := "id_example" // string | The ID of record
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)
    instanceType := "eval" // string | The Kafka instance type to filter the results by (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetCloudProviderRegions(context.Background(), id).Page(page).Size(size).InstanceType(instanceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviderRegions`: CloudRegionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **instanceType** | **string** | The Kafka instance type to filter the results by | 

### Return type

[**CloudRegionList**](CloudRegionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudProviders

> CloudProviderList GetCloudProviders(ctx).Page(page).Size(size).Execute()

Returns the list of supported cloud providers

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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetCloudProviders(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCloudProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudProviders`: CloudProviderList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCloudProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**CloudProviderList**](CloudProviderList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaById

> KafkaRequest GetKafkaById(ctx, id).Execute()

Returns a Kafka request by ID

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
    id := "id_example" // string | The ID of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKafkaById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKafkaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaById`: KafkaRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKafkaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkas

> KafkaRequestList GetKafkas(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Returns a list of Kafka requests

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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the `order by` clause of an SQL statement. Each query can be ordered by any of the following `kafkaRequests` fields:  * bootstrap_server_host * cloud_provider * cluster_id * created_at * href * id * instance_type * multi_az * name * organisation_id * owner * reauthentication_enabled * region * status * updated_at * version  For example, to return all Kafka instances ordered by their name, use the following syntax:  ```sql name asc ```  To return all Kafka instances ordered by their name _and_ created date, use the following syntax:  ```sql name asc, created_at asc ```  If the parameter isn't provided, or if the value is empty, then the results are ordered by name. (optional)
    search := "name = my-kafka and cloud_provider = aws" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the `where` clause of an SQL statement. Allowed fields in the search are `cloud_provider`, `name`, `owner`, `region`, and `status`. Allowed comparators are `<>`, `=`, or `LIKE`. Allowed joins are `AND` and `OR`. However, you can use a maximum of 10 joins in a search query.  Examples:  To return a Kafka instance with the name `my-kafka` and the region `aws`, use the following syntax:  ``` name = my-kafka and cloud_provider = aws ```[p-]  To return a Kafka instance with a name that starts with `my`, use the following syntax:  ``` name like my%25 ```  If the parameter isn't provided, or if the value is empty, then all the Kafka instances that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKafkas(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKafkas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkas`: KafkaRequestList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKafkas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the &#x60;order by&#x60; clause of an SQL statement. Each query can be ordered by any of the following &#x60;kafkaRequests&#x60; fields:  * bootstrap_server_host * cloud_provider * cluster_id * created_at * href * id * instance_type * multi_az * name * organisation_id * owner * reauthentication_enabled * region * status * updated_at * version  For example, to return all Kafka instances ordered by their name, use the following syntax:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  To return all Kafka instances ordered by their name _and_ created date, use the following syntax:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results are ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the &#x60;where&#x60; clause of an SQL statement. Allowed fields in the search are &#x60;cloud_provider&#x60;, &#x60;name&#x60;, &#x60;owner&#x60;, &#x60;region&#x60;, and &#x60;status&#x60;. Allowed comparators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60;, or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;. However, you can use a maximum of 10 joins in a search query.  Examples:  To return a Kafka instance with the name &#x60;my-kafka&#x60; and the region &#x60;aws&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name &#x3D; my-kafka and cloud_provider &#x3D; aws &#x60;&#x60;&#x60;[p-]  To return a Kafka instance with a name that starts with &#x60;my&#x60;, use the following syntax:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the Kafka instances that the user has permission to see are returned.  Note. If the query is invalid, an error is returned.  | 

### Return type

[**KafkaRequestList**](KafkaRequestList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByInstantQuery

> MetricsInstantQueryList GetMetricsByInstantQuery(ctx, id).Filters(filters).Execute()

Returns metrics with instant query by Kafka ID

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
    id := "id_example" // string | The ID of record
    filters := []string{"Inner_example"} // []string | List of metrics to fetch. Fetch all metrics when empty. List entries are Kafka internal metric names. (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetricsByInstantQuery(context.Background(), id).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetricsByInstantQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsByInstantQuery`: MetricsInstantQueryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetricsByInstantQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsByInstantQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **[]string** | List of metrics to fetch. Fetch all metrics when empty. List entries are Kafka internal metric names. | [default to []]

### Return type

[**MetricsInstantQueryList**](MetricsInstantQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByRangeQuery

> MetricsRangeQueryList GetMetricsByRangeQuery(ctx, id).Duration(duration).Interval(interval).Filters(filters).Execute()

Returns metrics with timeseries range query by Kafka ID

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
    id := "id_example" // string | The ID of record
    duration := int64(5) // int64 | The length of time in minutes for which to return the metrics (default to 5)
    interval := int64(30) // int64 | The interval in seconds between data points (default to 30)
    filters := []string{"Inner_example"} // []string | List of metrics to fetch. Fetch all metrics when empty. List entries are Kafka internal metric names. (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetricsByRangeQuery(context.Background(), id).Duration(duration).Interval(interval).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetricsByRangeQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsByRangeQuery`: MetricsRangeQueryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetricsByRangeQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsByRangeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **int64** | The length of time in minutes for which to return the metrics | [default to 5]
 **interval** | **int64** | The interval in seconds between data points | [default to 30]
 **filters** | **[]string** | List of metrics to fetch. Fetch all metrics when empty. List entries are Kafka internal metric names. | [default to []]

### Return type

[**MetricsRangeQueryList**](MetricsRangeQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceStatus

> ServiceStatus GetServiceStatus(ctx).Execute()

Returns the status of resources, such as whether maximum service capacity has been reached



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
    resp, r, err := api_client.DefaultApi.GetServiceStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetServiceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceStatus`: ServiceStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceStatusRequest struct via the builder pattern


### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionMetadata

> VersionMetadata GetVersionMetadata(ctx).Execute()

Returns the version metadata

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
    resp, r, err := api_client.DefaultApi.GetVersionMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVersionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersionMetadata`: VersionMetadata
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVersionMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionMetadataRequest struct via the builder pattern


### Return type

[**VersionMetadata**](VersionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaById

> KafkaRequest UpdateKafkaById(ctx, id).KafkaUpdateRequest(kafkaUpdateRequest).Execute()

Update a Kafka instance by id

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
    id := "id_example" // string | The ID of record
    kafkaUpdateRequest := *openapiclient.NewKafkaUpdateRequest() // KafkaUpdateRequest | Update owner of kafka

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateKafkaById(context.Background(), id).KafkaUpdateRequest(kafkaUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateKafkaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKafkaById`: KafkaRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateKafkaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaUpdateRequest** | [**KafkaUpdateRequest**](KafkaUpdateRequest.md) | Update owner of kafka | 

### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

