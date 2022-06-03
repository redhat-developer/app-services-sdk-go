# Record1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** | The record&#39;s partition within the topic | [optional] 
**Offset** | Pointer to **int64** | The record&#39;s offset within the topic partition | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | Timestamp associated with the record. The type is indicated by &#x60;timestampType&#x60;. When producing a record, this value will be used as the record&#39;s &#x60;CREATE_TIME&#x60;. | [optional] 
**TimestampType** | Pointer to **string** | Type of timestamp associated with the record | [optional] [readonly] 
**Headers** | Pointer to **map[string]string** | Record headers, key/value pairs | [optional] 
**Key** | Pointer to **string** | Record key | [optional] 
**Value** | **string** | Record value | 

## Methods

### NewRecord1

`func NewRecord1(value string, ) *Record1`

NewRecord1 instantiates a new Record1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecord1WithDefaults

`func NewRecord1WithDefaults() *Record1`

NewRecord1WithDefaults instantiates a new Record1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *Record1) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Record1) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Record1) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Record1) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *Record1) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Record1) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Record1) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Record1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *Record1) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Record1) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Record1) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Record1) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampType

`func (o *Record1) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *Record1) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *Record1) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *Record1) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### GetHeaders

`func (o *Record1) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Record1) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Record1) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Record1) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetKey

`func (o *Record1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Record1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Record1) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Record1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Record1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Record1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Record1) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


