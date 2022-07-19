# RecordAllOf

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

### NewRecordAllOf

`func NewRecordAllOf(value string, ) *RecordAllOf`

NewRecordAllOf instantiates a new RecordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAllOfWithDefaults

`func NewRecordAllOfWithDefaults() *RecordAllOf`

NewRecordAllOfWithDefaults instantiates a new RecordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *RecordAllOf) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *RecordAllOf) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *RecordAllOf) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *RecordAllOf) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetOffset

`func (o *RecordAllOf) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RecordAllOf) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RecordAllOf) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RecordAllOf) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *RecordAllOf) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordAllOf) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordAllOf) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RecordAllOf) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampType

`func (o *RecordAllOf) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *RecordAllOf) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *RecordAllOf) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.

### HasTimestampType

`func (o *RecordAllOf) HasTimestampType() bool`

HasTimestampType returns a boolean if a field has been set.

### GetHeaders

`func (o *RecordAllOf) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RecordAllOf) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RecordAllOf) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RecordAllOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetKey

`func (o *RecordAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RecordAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RecordAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RecordAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *RecordAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RecordAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RecordAllOf) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


