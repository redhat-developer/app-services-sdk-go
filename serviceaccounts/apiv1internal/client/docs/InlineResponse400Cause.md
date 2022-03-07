# InlineResponse400Cause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackTrace** | Pointer to [**[]InlineResponse400CauseStackTrace**](InlineResponse400CauseStackTrace.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**LocalizedMessage** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to [**[]InlineResponse400CauseSuppressed**](InlineResponse400CauseSuppressed.md) |  | [optional] 


## Methods

### NewInlineResponse400Cause

`func NewInlineResponse400Cause() *InlineResponse400Cause`

NewInlineResponse400Cause instantiates a new InlineResponse400Cause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse400CauseWithDefaults

`func NewInlineResponse400CauseWithDefaults() *InlineResponse400Cause`

NewInlineResponse400CauseWithDefaults instantiates a new InlineResponse400Cause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetStackTrace

`func (o *InlineResponse400Cause) GetStackTrace() []InlineResponse400CauseStackTrace`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *InlineResponse400Cause) GetStackTraceOk() (*[]InlineResponse400CauseStackTrace, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *InlineResponse400Cause) SetStackTrace(v []InlineResponse400CauseStackTrace)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *InlineResponse400Cause) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.


### GetMessage

`func (o *InlineResponse400Cause) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse400Cause) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse400Cause) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse400Cause) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


### GetLocalizedMessage

`func (o *InlineResponse400Cause) GetLocalizedMessage() string`

GetLocalizedMessage returns the LocalizedMessage field if non-nil, zero value otherwise.

### GetLocalizedMessageOk

`func (o *InlineResponse400Cause) GetLocalizedMessageOk() (*string, bool)`

GetLocalizedMessageOk returns a tuple with the LocalizedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMessage

`func (o *InlineResponse400Cause) SetLocalizedMessage(v string)`

SetLocalizedMessage sets LocalizedMessage field to given value.

### HasLocalizedMessage

`func (o *InlineResponse400Cause) HasLocalizedMessage() bool`

HasLocalizedMessage returns a boolean if a field has been set.


### GetSuppressed

`func (o *InlineResponse400Cause) GetSuppressed() []InlineResponse400CauseSuppressed`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *InlineResponse400Cause) GetSuppressedOk() (*[]InlineResponse400CauseSuppressed, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *InlineResponse400Cause) SetSuppressed(v []InlineResponse400CauseSuppressed)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *InlineResponse400Cause) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

