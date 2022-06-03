# ConsumerGroupMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LaggingPartitions** | Pointer to **int32** |  | [optional] 
**ActiveConsumers** | Pointer to **int32** |  | [optional] 
**UnassignedPartitions** | Pointer to **int32** |  | [optional] 

## Methods

### NewConsumerGroupMetrics

`func NewConsumerGroupMetrics() *ConsumerGroupMetrics`

NewConsumerGroupMetrics instantiates a new ConsumerGroupMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupMetricsWithDefaults

`func NewConsumerGroupMetricsWithDefaults() *ConsumerGroupMetrics`

NewConsumerGroupMetricsWithDefaults instantiates a new ConsumerGroupMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLaggingPartitions

`func (o *ConsumerGroupMetrics) GetLaggingPartitions() int32`

GetLaggingPartitions returns the LaggingPartitions field if non-nil, zero value otherwise.

### GetLaggingPartitionsOk

`func (o *ConsumerGroupMetrics) GetLaggingPartitionsOk() (*int32, bool)`

GetLaggingPartitionsOk returns a tuple with the LaggingPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaggingPartitions

`func (o *ConsumerGroupMetrics) SetLaggingPartitions(v int32)`

SetLaggingPartitions sets LaggingPartitions field to given value.

### HasLaggingPartitions

`func (o *ConsumerGroupMetrics) HasLaggingPartitions() bool`

HasLaggingPartitions returns a boolean if a field has been set.

### GetActiveConsumers

`func (o *ConsumerGroupMetrics) GetActiveConsumers() int32`

GetActiveConsumers returns the ActiveConsumers field if non-nil, zero value otherwise.

### GetActiveConsumersOk

`func (o *ConsumerGroupMetrics) GetActiveConsumersOk() (*int32, bool)`

GetActiveConsumersOk returns a tuple with the ActiveConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConsumers

`func (o *ConsumerGroupMetrics) SetActiveConsumers(v int32)`

SetActiveConsumers sets ActiveConsumers field to given value.

### HasActiveConsumers

`func (o *ConsumerGroupMetrics) HasActiveConsumers() bool`

HasActiveConsumers returns a boolean if a field has been set.

### GetUnassignedPartitions

`func (o *ConsumerGroupMetrics) GetUnassignedPartitions() int32`

GetUnassignedPartitions returns the UnassignedPartitions field if non-nil, zero value otherwise.

### GetUnassignedPartitionsOk

`func (o *ConsumerGroupMetrics) GetUnassignedPartitionsOk() (*int32, bool)`

GetUnassignedPartitionsOk returns a tuple with the UnassignedPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedPartitions

`func (o *ConsumerGroupMetrics) SetUnassignedPartitions(v int32)`

SetUnassignedPartitions sets UnassignedPartitions field to given value.

### HasUnassignedPartitions

`func (o *ConsumerGroupMetrics) HasUnassignedPartitions() bool`

HasUnassignedPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


