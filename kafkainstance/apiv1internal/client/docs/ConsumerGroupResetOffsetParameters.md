# ConsumerGroupResetOffsetParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Offset** | **string** |  | 
**Topics** | Pointer to [**[]TopicsToResetOffset**](TopicsToResetOffset.md) |  | [optional] 


## Methods

### NewConsumerGroupResetOffsetParameters

`func NewConsumerGroupResetOffsetParameters(offset string, ) *ConsumerGroupResetOffsetParameters`

NewConsumerGroupResetOffsetParameters instantiates a new ConsumerGroupResetOffsetParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupResetOffsetParametersWithDefaults

`func NewConsumerGroupResetOffsetParametersWithDefaults() *ConsumerGroupResetOffsetParameters`

NewConsumerGroupResetOffsetParametersWithDefaults instantiates a new ConsumerGroupResetOffsetParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


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


### GetOffset

`func (o *ConsumerGroupResetOffsetParameters) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConsumerGroupResetOffsetParameters) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConsumerGroupResetOffsetParameters) SetOffset(v string)`

SetOffset sets Offset field to given value.



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

