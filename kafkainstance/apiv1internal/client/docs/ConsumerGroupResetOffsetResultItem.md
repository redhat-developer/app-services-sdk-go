# ConsumerGroupResetOffsetResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 

## Methods

### NewConsumerGroupResetOffsetResultItem

`func NewConsumerGroupResetOffsetResultItem() *ConsumerGroupResetOffsetResultItem`

NewConsumerGroupResetOffsetResultItem instantiates a new ConsumerGroupResetOffsetResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupResetOffsetResultItemWithDefaults

`func NewConsumerGroupResetOffsetResultItemWithDefaults() *ConsumerGroupResetOffsetResultItem`

NewConsumerGroupResetOffsetResultItemWithDefaults instantiates a new ConsumerGroupResetOffsetResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *ConsumerGroupResetOffsetResultItem) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConsumerGroupResetOffsetResultItem) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ConsumerGroupResetOffsetResultItem) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ConsumerGroupResetOffsetResultItem) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPartition

`func (o *ConsumerGroupResetOffsetResultItem) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ConsumerGroupResetOffsetResultItem) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ConsumerGroupResetOffsetResultItem) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ConsumerGroupResetOffsetResultItem) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *ConsumerGroupResetOffsetResultItem) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConsumerGroupResetOffsetResultItem) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConsumerGroupResetOffsetResultItem) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ConsumerGroupResetOffsetResultItem) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


