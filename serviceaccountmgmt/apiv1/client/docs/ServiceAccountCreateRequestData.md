# ServiceAccountCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceAccountCreateRequestData

`func NewServiceAccountCreateRequestData(name string, ) *ServiceAccountCreateRequestData`

NewServiceAccountCreateRequestData instantiates a new ServiceAccountCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountCreateRequestDataWithDefaults

`func NewServiceAccountCreateRequestDataWithDefaults() *ServiceAccountCreateRequestData`

NewServiceAccountCreateRequestDataWithDefaults instantiates a new ServiceAccountCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceAccountCreateRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountCreateRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountCreateRequestData) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceAccountCreateRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountCreateRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountCreateRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountCreateRequestData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


