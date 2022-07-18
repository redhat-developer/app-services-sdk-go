# ConsumerGroupResetOffsetParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | [**OffsetType**](OffsetType.md) |  | 
**Value** | Pointer to **string** | Value associated with the given &#x60;offset&#x60;. Not used for &#x60;offset&#x60; values &#x60;earliest&#x60; and &#x60;latest&#x60;. When &#x60;offset&#x60; is &#x60;timestamp&#x60; then &#x60;value&#x60; must be a valid timestamp representing the point in time to reset the consumer group. When &#x60;offset&#x60; is &#x60;absolute&#x60; then &#x60;value&#x60; must be the integer offset to which the consumer group will be reset. | [optional] 
**Topics** | Pointer to [**[]TopicsToResetOffset**](TopicsToResetOffset.md) |  | [optional] 

## Methods

### NewConsumerGroupResetOffsetParameters

`func NewConsumerGroupResetOffsetParameters(offset OffsetType, ) *ConsumerGroupResetOffsetParameters`

NewConsumerGroupResetOffsetParameters instantiates a new ConsumerGroupResetOffsetParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupResetOffsetParametersWithDefaults

`func NewConsumerGroupResetOffsetParametersWithDefaults() *ConsumerGroupResetOffsetParameters`

NewConsumerGroupResetOffsetParametersWithDefaults instantiates a new ConsumerGroupResetOffsetParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ConsumerGroupResetOffsetParameters) GetOffset() OffsetType`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConsumerGroupResetOffsetParameters) GetOffsetOk() (*OffsetType, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConsumerGroupResetOffsetParameters) SetOffset(v OffsetType)`

SetOffset sets Offset field to given value.


### GetValue

`func (o *ConsumerGroupResetOffsetParameters) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsumerGroupResetOffsetParameters) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsumerGroupResetOffsetParameters) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsumerGroupResetOffsetParameters) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTopics

`func (o *ConsumerGroupResetOffsetParameters) GetTopics() []TopicsToResetOffset`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ConsumerGroupResetOffsetParameters) GetTopicsOk() (*[]TopicsToResetOffset, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ConsumerGroupResetOffsetParameters) SetTopics(v []TopicsToResetOffset)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ConsumerGroupResetOffsetParameters) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


