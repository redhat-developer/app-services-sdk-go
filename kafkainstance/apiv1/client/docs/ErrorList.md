# ErrorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Total** | **int32** | Total number of errors returned in this request | 
**Size** | Pointer to **int32** | Number of entries per page (returned for fetch requests) | [optional] 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 
**Items** | [**[]Error**](Error.md) |  | 

## Methods

### NewErrorList

`func NewErrorList(total int32, items []Error, ) *ErrorList`

NewErrorList instantiates a new ErrorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorListWithDefaults

`func NewErrorListWithDefaults() *ErrorList`

NewErrorListWithDefaults instantiates a new ErrorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ErrorList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTotal

`func (o *ErrorList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ErrorList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ErrorList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSize

`func (o *ErrorList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ErrorList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ErrorList) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ErrorList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *ErrorList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ErrorList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ErrorList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ErrorList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetItems

`func (o *ErrorList) GetItems() []Error`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorList) GetItemsOk() (*[]Error, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorList) SetItems(v []Error)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


