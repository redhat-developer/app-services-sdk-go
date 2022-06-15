# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the object. Not supported for all object kinds. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Link path to request the object. Not supported for all object kinds. | [optional] 
**Reason** | Pointer to **string** | General reason for the error. Does not change between specific occurrences. | [optional] 
**Detail** | Pointer to **string** | Detail specific to an error occurrence. May be different depending on the condition(s) that trigger the error. | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Error) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Error) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Error) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Error) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Error) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Error) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Error) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Error) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *Error) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Error) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Error) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Error) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetReason

`func (o *Error) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Error) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Error) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Error) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDetail

`func (o *Error) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Error) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Error) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Error) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCode

`func (o *Error) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Error) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Error) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Error) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Error) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetClass

`func (o *Error) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Error) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Error) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Error) HasClass() bool`

HasClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


