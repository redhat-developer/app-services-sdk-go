# CreateGroupMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** | User-defined name-value pairs. Name and value must be strings. | [optional] 
**Id** | **string** |  | 

## Methods

### NewCreateGroupMetaData

`func NewCreateGroupMetaData(id string, ) *CreateGroupMetaData`

NewCreateGroupMetaData instantiates a new CreateGroupMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupMetaDataWithDefaults

`func NewCreateGroupMetaDataWithDefaults() *CreateGroupMetaData`

NewCreateGroupMetaDataWithDefaults instantiates a new CreateGroupMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateGroupMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProperties

`func (o *CreateGroupMetaData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateGroupMetaData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateGroupMetaData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateGroupMetaData) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetId

`func (o *CreateGroupMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGroupMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGroupMetaData) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


