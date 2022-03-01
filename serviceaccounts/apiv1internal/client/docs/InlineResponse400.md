# InlineResponse400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to [**InlineResponse400Cause**](InlineResponse400Cause.md) |  | [optional] 
**StackTrace** | Pointer to [**[]InlineResponse400CauseStackTrace**](InlineResponse400CauseStackTrace.md) |  | [optional] 
**Response** | Pointer to [**InlineResponse400Response**](InlineResponse400Response.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**LocalizedMessage** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to [**[]InlineResponse400CauseSuppressed**](InlineResponse400CauseSuppressed.md) |  | [optional] 


## Methods

### NewInlineResponse400

`func NewInlineResponse400() *InlineResponse400`

NewInlineResponse400 instantiates a new InlineResponse400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse400WithDefaults

`func NewInlineResponse400WithDefaults() *InlineResponse400`

NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCause

`func (o *InlineResponse400) GetCause() InlineResponse400Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *InlineResponse400) GetCauseOk() (*InlineResponse400Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *InlineResponse400) SetCause(v InlineResponse400Cause)`

SetCause sets Cause field to given value.

### HasCause

`func (o *InlineResponse400) HasCause() bool`

HasCause returns a boolean if a field has been set.


### GetStackTrace

`func (o *InlineResponse400) GetStackTrace() []InlineResponse400CauseStackTrace`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *InlineResponse400) GetStackTraceOk() (*[]InlineResponse400CauseStackTrace, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *InlineResponse400) SetStackTrace(v []InlineResponse400CauseStackTrace)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *InlineResponse400) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.


### GetResponse

`func (o *InlineResponse400) GetResponse() InlineResponse400Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InlineResponse400) GetResponseOk() (*InlineResponse400Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InlineResponse400) SetResponse(v InlineResponse400Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InlineResponse400) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


### GetMessage

`func (o *InlineResponse400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse400) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse400) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


### GetLocalizedMessage

`func (o *InlineResponse400) GetLocalizedMessage() string`

GetLocalizedMessage returns the LocalizedMessage field if non-nil, zero value otherwise.

### GetLocalizedMessageOk

`func (o *InlineResponse400) GetLocalizedMessageOk() (*string, bool)`

GetLocalizedMessageOk returns a tuple with the LocalizedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMessage

`func (o *InlineResponse400) SetLocalizedMessage(v string)`

SetLocalizedMessage sets LocalizedMessage field to given value.

### HasLocalizedMessage

`func (o *InlineResponse400) HasLocalizedMessage() bool`

HasLocalizedMessage returns a boolean if a field has been set.


### GetSuppressed

`func (o *InlineResponse400) GetSuppressed() []InlineResponse400CauseSuppressed`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *InlineResponse400) GetSuppressedOk() (*[]InlineResponse400CauseSuppressed, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *InlineResponse400) SetSuppressed(v []InlineResponse400CauseSuppressed)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *InlineResponse400) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

