# SupportedKafkaSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;SupportedKafkaSize&#39; link. | [optional] 
**Id** | Pointer to **string** | Unique identifier of this Kafka instance size. | [optional] 
**IngressThroughputPerSec** | Pointer to **string** | Ingress throughput per second available to this Kafka instance size. | [optional] 
**EgressThroughputPerSec** | Pointer to **string** | Egress throughput per second available to this Kafka instance size. | [optional] 
**TotalMaxConnections** | Pointer to **int32** | Total amount of connections available to this Kafka instance size. | [optional] 
**MaxDataRetentionSize** | Pointer to **string** | Maximum data storage available to this Kafka instance size. | [optional] 
**MaxPartitions** | Pointer to **int32** | Total amount of partitions available to this Kafka instance size. | [optional] 
**MaxDataRetentionPeriod** | Pointer to **string** | Maximum data retention period available to this Kafka instance size. | [optional] 
**MaxConnectionAttemptsPerSec** | Pointer to **int32** | Maximium connection attempts per second available to this Kafka instance size. | [optional] 
**QuotaConsumed** | Pointer to **int32** | Quota consumed by this Kafka instance size. | [optional] 
**QuotaType** | Pointer to **string** | Quota type used by this Kafka instance size. | [optional] 
**CapacityConsumed** | Pointer to **int32** | Data plane cluster capacity consumed by this Kafka instance size. | [optional] 


## Methods

### NewSupportedKafkaSize

`func NewSupportedKafkaSize() *SupportedKafkaSize`

NewSupportedKafkaSize instantiates a new SupportedKafkaSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedKafkaSizeWithDefaults

`func NewSupportedKafkaSizeWithDefaults() *SupportedKafkaSize`

NewSupportedKafkaSizeWithDefaults instantiates a new SupportedKafkaSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKind

`func (o *SupportedKafkaSize) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SupportedKafkaSize) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SupportedKafkaSize) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SupportedKafkaSize) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetId

`func (o *SupportedKafkaSize) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedKafkaSize) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedKafkaSize) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedKafkaSize) HasId() bool`

HasId returns a boolean if a field has been set.


### GetIngressThroughputPerSec

`func (o *SupportedKafkaSize) GetIngressThroughputPerSec() string`

GetIngressThroughputPerSec returns the IngressThroughputPerSec field if non-nil, zero value otherwise.

### GetIngressThroughputPerSecOk

`func (o *SupportedKafkaSize) GetIngressThroughputPerSecOk() (*string, bool)`

GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressThroughputPerSec

`func (o *SupportedKafkaSize) SetIngressThroughputPerSec(v string)`

SetIngressThroughputPerSec sets IngressThroughputPerSec field to given value.

### HasIngressThroughputPerSec

`func (o *SupportedKafkaSize) HasIngressThroughputPerSec() bool`

HasIngressThroughputPerSec returns a boolean if a field has been set.


### GetEgressThroughputPerSec

`func (o *SupportedKafkaSize) GetEgressThroughputPerSec() string`

GetEgressThroughputPerSec returns the EgressThroughputPerSec field if non-nil, zero value otherwise.

### GetEgressThroughputPerSecOk

`func (o *SupportedKafkaSize) GetEgressThroughputPerSecOk() (*string, bool)`

GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressThroughputPerSec

`func (o *SupportedKafkaSize) SetEgressThroughputPerSec(v string)`

SetEgressThroughputPerSec sets EgressThroughputPerSec field to given value.

### HasEgressThroughputPerSec

`func (o *SupportedKafkaSize) HasEgressThroughputPerSec() bool`

HasEgressThroughputPerSec returns a boolean if a field has been set.


### GetTotalMaxConnections

`func (o *SupportedKafkaSize) GetTotalMaxConnections() int32`

GetTotalMaxConnections returns the TotalMaxConnections field if non-nil, zero value otherwise.

### GetTotalMaxConnectionsOk

`func (o *SupportedKafkaSize) GetTotalMaxConnectionsOk() (*int32, bool)`

GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaxConnections

`func (o *SupportedKafkaSize) SetTotalMaxConnections(v int32)`

SetTotalMaxConnections sets TotalMaxConnections field to given value.

### HasTotalMaxConnections

`func (o *SupportedKafkaSize) HasTotalMaxConnections() bool`

HasTotalMaxConnections returns a boolean if a field has been set.


### GetMaxDataRetentionSize

`func (o *SupportedKafkaSize) GetMaxDataRetentionSize() string`

GetMaxDataRetentionSize returns the MaxDataRetentionSize field if non-nil, zero value otherwise.

### GetMaxDataRetentionSizeOk

`func (o *SupportedKafkaSize) GetMaxDataRetentionSizeOk() (*string, bool)`

GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionSize

`func (o *SupportedKafkaSize) SetMaxDataRetentionSize(v string)`

SetMaxDataRetentionSize sets MaxDataRetentionSize field to given value.

### HasMaxDataRetentionSize

`func (o *SupportedKafkaSize) HasMaxDataRetentionSize() bool`

HasMaxDataRetentionSize returns a boolean if a field has been set.


### GetMaxPartitions

`func (o *SupportedKafkaSize) GetMaxPartitions() int32`

GetMaxPartitions returns the MaxPartitions field if non-nil, zero value otherwise.

### GetMaxPartitionsOk

`func (o *SupportedKafkaSize) GetMaxPartitionsOk() (*int32, bool)`

GetMaxPartitionsOk returns a tuple with the MaxPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPartitions

`func (o *SupportedKafkaSize) SetMaxPartitions(v int32)`

SetMaxPartitions sets MaxPartitions field to given value.

### HasMaxPartitions

`func (o *SupportedKafkaSize) HasMaxPartitions() bool`

HasMaxPartitions returns a boolean if a field has been set.


### GetMaxDataRetentionPeriod

`func (o *SupportedKafkaSize) GetMaxDataRetentionPeriod() string`

GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field if non-nil, zero value otherwise.

### GetMaxDataRetentionPeriodOk

`func (o *SupportedKafkaSize) GetMaxDataRetentionPeriodOk() (*string, bool)`

GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionPeriod

`func (o *SupportedKafkaSize) SetMaxDataRetentionPeriod(v string)`

SetMaxDataRetentionPeriod sets MaxDataRetentionPeriod field to given value.

### HasMaxDataRetentionPeriod

`func (o *SupportedKafkaSize) HasMaxDataRetentionPeriod() bool`

HasMaxDataRetentionPeriod returns a boolean if a field has been set.


### GetMaxConnectionAttemptsPerSec

`func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSec() int32`

GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field if non-nil, zero value otherwise.

### GetMaxConnectionAttemptsPerSecOk

`func (o *SupportedKafkaSize) GetMaxConnectionAttemptsPerSecOk() (*int32, bool)`

GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAttemptsPerSec

`func (o *SupportedKafkaSize) SetMaxConnectionAttemptsPerSec(v int32)`

SetMaxConnectionAttemptsPerSec sets MaxConnectionAttemptsPerSec field to given value.

### HasMaxConnectionAttemptsPerSec

`func (o *SupportedKafkaSize) HasMaxConnectionAttemptsPerSec() bool`

HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.


### GetQuotaConsumed

`func (o *SupportedKafkaSize) GetQuotaConsumed() int32`

GetQuotaConsumed returns the QuotaConsumed field if non-nil, zero value otherwise.

### GetQuotaConsumedOk

`func (o *SupportedKafkaSize) GetQuotaConsumedOk() (*int32, bool)`

GetQuotaConsumedOk returns a tuple with the QuotaConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaConsumed

`func (o *SupportedKafkaSize) SetQuotaConsumed(v int32)`

SetQuotaConsumed sets QuotaConsumed field to given value.

### HasQuotaConsumed

`func (o *SupportedKafkaSize) HasQuotaConsumed() bool`

HasQuotaConsumed returns a boolean if a field has been set.


### GetQuotaType

`func (o *SupportedKafkaSize) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *SupportedKafkaSize) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *SupportedKafkaSize) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *SupportedKafkaSize) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.


### GetCapacityConsumed

`func (o *SupportedKafkaSize) GetCapacityConsumed() int32`

GetCapacityConsumed returns the CapacityConsumed field if non-nil, zero value otherwise.

### GetCapacityConsumedOk

`func (o *SupportedKafkaSize) GetCapacityConsumedOk() (*int32, bool)`

GetCapacityConsumedOk returns a tuple with the CapacityConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityConsumed

`func (o *SupportedKafkaSize) SetCapacityConsumed(v int32)`

SetCapacityConsumed sets CapacityConsumed field to given value.

### HasCapacityConsumed

`func (o *SupportedKafkaSize) HasCapacityConsumed() bool`

HasCapacityConsumed returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

