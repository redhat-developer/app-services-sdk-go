# ErrorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | General reason for the error. Does not change between specific occurrences. | [optional] 
**Detail** | Pointer to **string** | Detail specific to an error occurrence. May be different depending on the condition(s) that trigger the error. | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorAllOf

`func NewErrorAllOf() *ErrorAllOf`

NewErrorAllOf instantiates a new ErrorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorAllOfWithDefaults

`func NewErrorAllOfWithDefaults() *ErrorAllOf`

NewErrorAllOfWithDefaults instantiates a new ErrorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ErrorAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorAllOf) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorAllOf) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorAllOf) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorAllOf) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCode

`func (o *ErrorAllOf) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorAllOf) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorAllOf) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ErrorAllOf) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorAllOf) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorAllOf) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorAllOf) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetClass

`func (o *ErrorAllOf) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ErrorAllOf) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ErrorAllOf) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *ErrorAllOf) HasClass() bool`

HasClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


