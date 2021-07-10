# UpdateTopicInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | Topic configuration entries. | [optional] 
**NumPartitions** | Pointer to **int32** | Number of partitions (only increasing supported) | [optional] 


## Methods

### NewUpdateTopicInput

`func NewUpdateTopicInput() *UpdateTopicInput`

NewUpdateTopicInput instantiates a new UpdateTopicInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTopicInputWithDefaults

`func NewUpdateTopicInputWithDefaults() *UpdateTopicInput`

NewUpdateTopicInputWithDefaults instantiates a new UpdateTopicInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetConfig

`func (o *UpdateTopicInput) GetConfig() []ConfigEntry`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateTopicInput) GetConfigOk() (*[]ConfigEntry, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateTopicInput) SetConfig(v []ConfigEntry)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateTopicInput) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


### GetNumPartitions

`func (o *UpdateTopicInput) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *UpdateTopicInput) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *UpdateTopicInput) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.

### HasNumPartitions

`func (o *UpdateTopicInput) HasNumPartitions() bool`

HasNumPartitions returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

