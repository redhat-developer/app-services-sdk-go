# AccountGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewAccountGroupRequest

`func NewAccountGroupRequest(description string, name string, ) *AccountGroupRequest`

NewAccountGroupRequest instantiates a new AccountGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupRequestWithDefaults

`func NewAccountGroupRequestWithDefaults() *AccountGroupRequest`

NewAccountGroupRequestWithDefaults instantiates a new AccountGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AccountGroupRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccountGroupRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccountGroupRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AccountGroupRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AccountGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AccountGroupRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountGroupRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountGroupRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AccountGroupRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *AccountGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *AccountGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountGroupRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


