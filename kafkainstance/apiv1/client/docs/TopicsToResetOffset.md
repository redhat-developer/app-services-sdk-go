# TopicsToResetOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** |  | 
**Partitions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewTopicsToResetOffset

`func NewTopicsToResetOffset(topic string, ) *TopicsToResetOffset`

NewTopicsToResetOffset instantiates a new TopicsToResetOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicsToResetOffsetWithDefaults

`func NewTopicsToResetOffsetWithDefaults() *TopicsToResetOffset`

NewTopicsToResetOffsetWithDefaults instantiates a new TopicsToResetOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *TopicsToResetOffset) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TopicsToResetOffset) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TopicsToResetOffset) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetPartitions

`func (o *TopicsToResetOffset) GetPartitions() []int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *TopicsToResetOffset) GetPartitionsOk() (*[]int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *TopicsToResetOffset) SetPartitions(v []int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *TopicsToResetOffset) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


