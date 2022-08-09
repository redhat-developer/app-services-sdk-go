# ErrorListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Items** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**Page** | **int64** |  | 
**Size** | **int64** |  | 
**Total** | **int64** |  | 

## Methods

### NewErrorListResponse

`func NewErrorListResponse(kind string, page int64, size int64, total int64, ) *ErrorListResponse`

NewErrorListResponse instantiates a new ErrorListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorListResponseWithDefaults

`func NewErrorListResponseWithDefaults() *ErrorListResponse`

NewErrorListResponseWithDefaults instantiates a new ErrorListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ErrorListResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorListResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorListResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetItems

`func (o *ErrorListResponse) GetItems() []Error`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorListResponse) GetItemsOk() (*[]Error, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorListResponse) SetItems(v []Error)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *ErrorListResponse) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ErrorListResponse) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ErrorListResponse) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *ErrorListResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ErrorListResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ErrorListResponse) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotal

`func (o *ErrorListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ErrorListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ErrorListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


