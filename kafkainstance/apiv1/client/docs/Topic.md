# Topic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the object. Not supported for all object kinds. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Link path to request the object. Not supported for all object kinds. | [optional] 
**Name** | Pointer to **string** | The name of the topic. | [optional] 
**IsInternal** | Pointer to **bool** |  | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) | Partitions for this topic. | [optional] 
**Config** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | Topic configuration entry. | [optional] 

## Methods

### NewTopic

`func NewTopic() *Topic`

NewTopic instantiates a new Topic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicWithDefaults

`func NewTopicWithDefaults() *Topic`

NewTopicWithDefaults instantiates a new Topic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Topic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Topic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Topic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Topic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Topic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Topic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Topic) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Topic) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *Topic) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Topic) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Topic) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Topic) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *Topic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Topic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Topic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Topic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsInternal

`func (o *Topic) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *Topic) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *Topic) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *Topic) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### GetPartitions

`func (o *Topic) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Topic) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Topic) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *Topic) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetConfig

`func (o *Topic) GetConfig() []ConfigEntry`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Topic) GetConfigOk() (*[]ConfigEntry, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Topic) SetConfig(v []ConfigEntry)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Topic) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


