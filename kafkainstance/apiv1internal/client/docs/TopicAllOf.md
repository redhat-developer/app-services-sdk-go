# TopicAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the topic. | [optional] 
**IsInternal** | Pointer to **bool** |  | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) | Partitions for this topic. | [optional] 
**Config** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | Topic configuration entry. | [optional] 

## Methods

### NewTopicAllOf

`func NewTopicAllOf() *TopicAllOf`

NewTopicAllOf instantiates a new TopicAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicAllOfWithDefaults

`func NewTopicAllOfWithDefaults() *TopicAllOf`

NewTopicAllOfWithDefaults instantiates a new TopicAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopicAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopicAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopicAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopicAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsInternal

`func (o *TopicAllOf) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *TopicAllOf) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *TopicAllOf) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *TopicAllOf) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### GetPartitions

`func (o *TopicAllOf) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *TopicAllOf) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *TopicAllOf) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *TopicAllOf) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetConfig

`func (o *TopicAllOf) GetConfig() []ConfigEntry`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TopicAllOf) GetConfigOk() (*[]ConfigEntry, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TopicAllOf) SetConfig(v []ConfigEntry)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TopicAllOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


