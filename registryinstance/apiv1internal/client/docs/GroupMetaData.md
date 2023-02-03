# GroupMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An ID of a single artifact group. | 
**Description** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedOn** | **string** |  | 
**ModifiedBy** | **string** |  | 
**ModifiedOn** | **string** |  | 
**Properties** | **map[string]string** | User-defined name-value pairs. Name and value must be strings. | 

## Methods

### NewGroupMetaData

`func NewGroupMetaData(id string, description string, createdBy string, createdOn string, modifiedBy string, modifiedOn string, properties map[string]string, ) *GroupMetaData`

NewGroupMetaData instantiates a new GroupMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMetaDataWithDefaults

`func NewGroupMetaDataWithDefaults() *GroupMetaData`

NewGroupMetaDataWithDefaults instantiates a new GroupMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMetaData) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *GroupMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedBy

`func (o *GroupMetaData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GroupMetaData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GroupMetaData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedOn

`func (o *GroupMetaData) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *GroupMetaData) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *GroupMetaData) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedBy

`func (o *GroupMetaData) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GroupMetaData) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GroupMetaData) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifiedOn

`func (o *GroupMetaData) GetModifiedOn() string`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *GroupMetaData) GetModifiedOnOk() (*string, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *GroupMetaData) SetModifiedOn(v string)`

SetModifiedOn sets ModifiedOn field to given value.


### GetProperties

`func (o *GroupMetaData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupMetaData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupMetaData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


