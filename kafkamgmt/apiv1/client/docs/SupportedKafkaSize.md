# SupportedKafkaSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of this Kafka instance size. | [optional] 
**DisplayName** | Pointer to **string** | Display name of this Kafka instance size. | [optional] 
**IngressThroughputPerSec** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**EgressThroughputPerSec** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**TotalMaxConnections** | Pointer to **int32** | Maximum amount of total connections available to this Kafka instance size. | [optional] 
**MaxDataRetentionSize** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**MaxPartitions** | Pointer to **int32** | Maximum amount of total partitions available to this Kafka instance size. | [optional] 
**MaxDataRetentionPeriod** | Pointer to **string** | Maximum data retention period available to this Kafka instance size. | [optional] 
**MaxConnectionAttemptsPerSec** | Pointer to **int32** | Maximium connection attempts per second available to this Kafka instance size. | [optional] 
**MaxMessageSize** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**MinInSyncReplicas** | Pointer to **int32** | Minimum number of in-sync replicas. | [optional] 
**ReplicationFactor** | Pointer to **int32** | Replication factor available to this Kafka instance size. | [optional] 
**SupportedAzModes** | Pointer to **[]string** | List of Availability Zone modes that this Kafka instance size supports. The possible values are \&quot;single\&quot;, \&quot;multi\&quot;. | [optional] 
**LifespanSeconds** | Pointer to **NullableInt32** | The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires. | [optional] 
**QuotaConsumed** | Pointer to **int32** | Quota consumed by this Kafka instance size. | [optional] 
**QuotaType** | Pointer to **string** | Quota type used by this Kafka instance size. This is now deprecated, please refer to supported_billing_models at instance-type level instead. | [optional] 
**CapacityConsumed** | Pointer to **int32** | Data plane cluster capacity consumed by this Kafka instance size. | [optional] 
**MaturityStatus** | Pointer to **string** | Maturity level of the size. Can be \&quot;stable\&quot; or \&quot;preview\&quot;. | [optional] 

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

### GetDisplayName

`func (o *SupportedKafkaSize) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SupportedKafkaSize) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SupportedKafkaSize) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SupportedKafkaSize) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIngressThroughputPerSec

`func (o *SupportedKafkaSize) GetIngressThroughputPerSec() SupportedKafkaSizeBytesValueItem`

GetIngressThroughputPerSec returns the IngressThroughputPerSec field if non-nil, zero value otherwise.

### GetIngressThroughputPerSecOk

`func (o *SupportedKafkaSize) GetIngressThroughputPerSecOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressThroughputPerSec

`func (o *SupportedKafkaSize) SetIngressThroughputPerSec(v SupportedKafkaSizeBytesValueItem)`

SetIngressThroughputPerSec sets IngressThroughputPerSec field to given value.

### HasIngressThroughputPerSec

`func (o *SupportedKafkaSize) HasIngressThroughputPerSec() bool`

HasIngressThroughputPerSec returns a boolean if a field has been set.

### GetEgressThroughputPerSec

`func (o *SupportedKafkaSize) GetEgressThroughputPerSec() SupportedKafkaSizeBytesValueItem`

GetEgressThroughputPerSec returns the EgressThroughputPerSec field if non-nil, zero value otherwise.

### GetEgressThroughputPerSecOk

`func (o *SupportedKafkaSize) GetEgressThroughputPerSecOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressThroughputPerSec

`func (o *SupportedKafkaSize) SetEgressThroughputPerSec(v SupportedKafkaSizeBytesValueItem)`

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

`func (o *SupportedKafkaSize) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem`

GetMaxDataRetentionSize returns the MaxDataRetentionSize field if non-nil, zero value otherwise.

### GetMaxDataRetentionSizeOk

`func (o *SupportedKafkaSize) GetMaxDataRetentionSizeOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionSize

`func (o *SupportedKafkaSize) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem)`

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

### GetMaxMessageSize

`func (o *SupportedKafkaSize) GetMaxMessageSize() SupportedKafkaSizeBytesValueItem`

GetMaxMessageSize returns the MaxMessageSize field if non-nil, zero value otherwise.

### GetMaxMessageSizeOk

`func (o *SupportedKafkaSize) GetMaxMessageSizeOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetMaxMessageSizeOk returns a tuple with the MaxMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageSize

`func (o *SupportedKafkaSize) SetMaxMessageSize(v SupportedKafkaSizeBytesValueItem)`

SetMaxMessageSize sets MaxMessageSize field to given value.

### HasMaxMessageSize

`func (o *SupportedKafkaSize) HasMaxMessageSize() bool`

HasMaxMessageSize returns a boolean if a field has been set.

### GetMinInSyncReplicas

`func (o *SupportedKafkaSize) GetMinInSyncReplicas() int32`

GetMinInSyncReplicas returns the MinInSyncReplicas field if non-nil, zero value otherwise.

### GetMinInSyncReplicasOk

`func (o *SupportedKafkaSize) GetMinInSyncReplicasOk() (*int32, bool)`

GetMinInSyncReplicasOk returns a tuple with the MinInSyncReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInSyncReplicas

`func (o *SupportedKafkaSize) SetMinInSyncReplicas(v int32)`

SetMinInSyncReplicas sets MinInSyncReplicas field to given value.

### HasMinInSyncReplicas

`func (o *SupportedKafkaSize) HasMinInSyncReplicas() bool`

HasMinInSyncReplicas returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *SupportedKafkaSize) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *SupportedKafkaSize) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *SupportedKafkaSize) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *SupportedKafkaSize) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetSupportedAzModes

`func (o *SupportedKafkaSize) GetSupportedAzModes() []string`

GetSupportedAzModes returns the SupportedAzModes field if non-nil, zero value otherwise.

### GetSupportedAzModesOk

`func (o *SupportedKafkaSize) GetSupportedAzModesOk() (*[]string, bool)`

GetSupportedAzModesOk returns a tuple with the SupportedAzModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAzModes

`func (o *SupportedKafkaSize) SetSupportedAzModes(v []string)`

SetSupportedAzModes sets SupportedAzModes field to given value.

### HasSupportedAzModes

`func (o *SupportedKafkaSize) HasSupportedAzModes() bool`

HasSupportedAzModes returns a boolean if a field has been set.

### GetLifespanSeconds

`func (o *SupportedKafkaSize) GetLifespanSeconds() int32`

GetLifespanSeconds returns the LifespanSeconds field if non-nil, zero value otherwise.

### GetLifespanSecondsOk

`func (o *SupportedKafkaSize) GetLifespanSecondsOk() (*int32, bool)`

GetLifespanSecondsOk returns a tuple with the LifespanSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifespanSeconds

`func (o *SupportedKafkaSize) SetLifespanSeconds(v int32)`

SetLifespanSeconds sets LifespanSeconds field to given value.

### HasLifespanSeconds

`func (o *SupportedKafkaSize) HasLifespanSeconds() bool`

HasLifespanSeconds returns a boolean if a field has been set.

### SetLifespanSecondsNil

`func (o *SupportedKafkaSize) SetLifespanSecondsNil(b bool)`

 SetLifespanSecondsNil sets the value for LifespanSeconds to be an explicit nil

### UnsetLifespanSeconds
`func (o *SupportedKafkaSize) UnsetLifespanSeconds()`

UnsetLifespanSeconds ensures that no value is present for LifespanSeconds, not even an explicit nil
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

### GetMaturityStatus

`func (o *SupportedKafkaSize) GetMaturityStatus() string`

GetMaturityStatus returns the MaturityStatus field if non-nil, zero value otherwise.

### GetMaturityStatusOk

`func (o *SupportedKafkaSize) GetMaturityStatusOk() (*string, bool)`

GetMaturityStatusOk returns a tuple with the MaturityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityStatus

`func (o *SupportedKafkaSize) SetMaturityStatus(v string)`

SetMaturityStatus sets MaturityStatus field to given value.

### HasMaturityStatus

`func (o *SupportedKafkaSize) HasMaturityStatus() bool`

HasMaturityStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


