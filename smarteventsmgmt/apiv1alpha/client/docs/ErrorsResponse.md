# ErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ErrorResponse**](ErrorResponse.md) |  | [optional] 

## Methods

### NewErrorsResponse

`func NewErrorsResponse() *ErrorsResponse`

NewErrorsResponse instantiates a new ErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResponseWithDefaults

`func NewErrorsResponseWithDefaults() *ErrorsResponse`

NewErrorsResponseWithDefaults instantiates a new ErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ErrorsResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorsResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorsResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorsResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *ErrorsResponse) GetItems() []ErrorResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorsResponse) GetItemsOk() (*[]ErrorResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorsResponse) SetItems(v []ErrorResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


