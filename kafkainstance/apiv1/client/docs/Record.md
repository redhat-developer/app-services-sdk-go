# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the object. Not supported for all object kinds. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Link path to request the object. Not supported for all object kinds. | [optional] 
**Partition** | Pointer to **int32** | The record&#39;s partition within the topic | [optional] 
**Offset** | Pointer to **int64** | The record&#39;s offset within the topic partition | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | Timestamp associated with the record. The type is indicated by &#x60;timestampType&#x60;. When producing a record, this value will be used as the record&#39;s &#x60;CREATE_TIME&#x60;. | [optional] 
**TimestampType** | Pointer to **string** | Type of timestamp associated with the record | [optional] [readonly] 
**Headers** | Pointer to **map[string]string** | Record headers, key/value pairs | [optional] 
**Key** | Pointer to **string** | Record key | [optional] 
**Value** | **string** | Record value | 

## Methods

### NewRecord

`func NewRecord(value string, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Record) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Record) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Record) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Record) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Record) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Record) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Record) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Record) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *Record) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Record) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Record) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Record) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetPartition

`func (o *Record) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Record) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Record) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Record) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *Record) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Record) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Record) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Record) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *Record) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Record) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Record) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Record) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampType

`func (o *Record) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *Record) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *Record) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *Record) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### GetHeaders

`func (o *Record) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Record) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Record) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Record) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetKey

`func (o *Record) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Record) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Record) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Record) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Record) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Record) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Record) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


