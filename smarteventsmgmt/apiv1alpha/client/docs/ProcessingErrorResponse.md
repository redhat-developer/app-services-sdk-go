# ProcessingErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordedAt** | Pointer to **time.Time** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProcessingErrorResponse

`func NewProcessingErrorResponse() *ProcessingErrorResponse`

NewProcessingErrorResponse instantiates a new ProcessingErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingErrorResponseWithDefaults

`func NewProcessingErrorResponseWithDefaults() *ProcessingErrorResponse`

NewProcessingErrorResponseWithDefaults instantiates a new ProcessingErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordedAt

`func (o *ProcessingErrorResponse) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *ProcessingErrorResponse) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *ProcessingErrorResponse) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *ProcessingErrorResponse) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetHeaders

`func (o *ProcessingErrorResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProcessingErrorResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProcessingErrorResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProcessingErrorResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPayload

`func (o *ProcessingErrorResponse) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProcessingErrorResponse) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProcessingErrorResponse) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ProcessingErrorResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


