# ErrorsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Items** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewErrorsList

`func NewErrorsList(kind string, ) *ErrorsList`

NewErrorsList instantiates a new ErrorsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsListWithDefaults

`func NewErrorsListWithDefaults() *ErrorsList`

NewErrorsListWithDefaults instantiates a new ErrorsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ErrorsList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorsList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorsList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetItems

`func (o *ErrorsList) GetItems() []Error`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorsList) GetItemsOk() (*[]Error, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorsList) SetItems(v []Error)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorsList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


