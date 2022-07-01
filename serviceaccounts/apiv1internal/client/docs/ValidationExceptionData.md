# ValidationExceptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **map[string]string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationExceptionData

`func NewValidationExceptionData() *ValidationExceptionData`

NewValidationExceptionData instantiates a new ValidationExceptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationExceptionDataWithDefaults

`func NewValidationExceptionDataWithDefaults() *ValidationExceptionData`

NewValidationExceptionDataWithDefaults instantiates a new ValidationExceptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ValidationExceptionData) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ValidationExceptionData) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ValidationExceptionData) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ValidationExceptionData) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetError

`func (o *ValidationExceptionData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationExceptionData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationExceptionData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ValidationExceptionData) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDescription

`func (o *ValidationExceptionData) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ValidationExceptionData) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ValidationExceptionData) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ValidationExceptionData) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


