# ErrorListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Error**](Error.md) |  | 
**Total** | Pointer to **int32** | Total number of errors returned in this request | [optional] 

## Methods

### NewErrorListAllOf

`func NewErrorListAllOf(items []Error, ) *ErrorListAllOf`

NewErrorListAllOf instantiates a new ErrorListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorListAllOfWithDefaults

`func NewErrorListAllOfWithDefaults() *ErrorListAllOf`

NewErrorListAllOfWithDefaults instantiates a new ErrorListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ErrorListAllOf) GetItems() []Error`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorListAllOf) GetItemsOk() (*[]Error, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorListAllOf) SetItems(v []Error)`

SetItems sets Items field to given value.


### GetTotal

`func (o *ErrorListAllOf) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ErrorListAllOf) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ErrorListAllOf) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ErrorListAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


