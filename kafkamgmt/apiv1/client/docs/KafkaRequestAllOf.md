# KafkaRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting]  | [optional] 
**CloudProvider** | Pointer to **string** | Name of Cloud used to deploy. For example AWS | [optional] 
**MultiAz** | **bool** |  | 
**Region** | Pointer to **string** | Values will be regions of specific cloud provider. For example: us-east-1 for AWS | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BootstrapServerHost** | Pointer to **string** |  | [optional] 
**AdminApiServerUrl** | Pointer to **string** | The kafka admin server url to perform kafka admin operations e.g acl management etc. The value will be available when the Kafka has been fully provisioned i.e it reaches a &#39;ready&#39; state | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**FailedReason** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceTypeName** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**ReauthenticationEnabled** | **bool** |  | 
**KafkaStorageSize** | Pointer to **string** | Maximum data storage available to this Kafka. This is now deprecated, please use max_data_retention_size instead. | [optional] 
**MaxDataRetentionSize** | Pointer to [**SupportedKafkaSizeBytesValueItem**](SupportedKafkaSizeBytesValueItem.md) |  | [optional] 
**BrowserUrl** | Pointer to **string** |  | [optional] 
**SizeId** | Pointer to **string** |  | [optional] 
**IngressThroughputPerSec** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**EgressThroughputPerSec** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**TotalMaxConnections** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxPartitions** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxDataRetentionPeriod** | Pointer to **string** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**MaxConnectionAttemptsPerSec** | Pointer to **int32** | This field is now deprecated, please use the /api/kafkas_mgmt/v1/instance_types/{cloud_provider}/{cloud_region} endpoint to retrieve the field instead. | [optional] 
**BillingCloudAccountId** | Pointer to **string** |  | [optional] 
**Marketplace** | Pointer to **string** |  | [optional] 
**BillingModel** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaRequestAllOf

`func NewKafkaRequestAllOf(multiAz bool, reauthenticationEnabled bool, ) *KafkaRequestAllOf`

NewKafkaRequestAllOf instantiates a new KafkaRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestAllOfWithDefaults

`func NewKafkaRequestAllOfWithDefaults() *KafkaRequestAllOf`

NewKafkaRequestAllOfWithDefaults instantiates a new KafkaRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *KafkaRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *KafkaRequestAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequestAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequestAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequestAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequestAllOf) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequestAllOf) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequestAllOf) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.


### GetRegion

`func (o *KafkaRequestAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequestAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequestAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequestAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *KafkaRequestAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KafkaRequestAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KafkaRequestAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KafkaRequestAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootstrapServerHost

`func (o *KafkaRequestAllOf) GetBootstrapServerHost() string`

GetBootstrapServerHost returns the BootstrapServerHost field if non-nil, zero value otherwise.

### GetBootstrapServerHostOk

`func (o *KafkaRequestAllOf) GetBootstrapServerHostOk() (*string, bool)`

GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServerHost

`func (o *KafkaRequestAllOf) SetBootstrapServerHost(v string)`

SetBootstrapServerHost sets BootstrapServerHost field to given value.

### HasBootstrapServerHost

`func (o *KafkaRequestAllOf) HasBootstrapServerHost() bool`

HasBootstrapServerHost returns a boolean if a field has been set.

### GetAdminApiServerUrl

`func (o *KafkaRequestAllOf) GetAdminApiServerUrl() string`

GetAdminApiServerUrl returns the AdminApiServerUrl field if non-nil, zero value otherwise.

### GetAdminApiServerUrlOk

`func (o *KafkaRequestAllOf) GetAdminApiServerUrlOk() (*string, bool)`

GetAdminApiServerUrlOk returns a tuple with the AdminApiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminApiServerUrl

`func (o *KafkaRequestAllOf) SetAdminApiServerUrl(v string)`

SetAdminApiServerUrl sets AdminApiServerUrl field to given value.

### HasAdminApiServerUrl

`func (o *KafkaRequestAllOf) HasAdminApiServerUrl() bool`

HasAdminApiServerUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KafkaRequestAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KafkaRequestAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KafkaRequestAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KafkaRequestAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *KafkaRequestAllOf) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KafkaRequestAllOf) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KafkaRequestAllOf) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *KafkaRequestAllOf) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *KafkaRequestAllOf) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *KafkaRequestAllOf) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetUpdatedAt

`func (o *KafkaRequestAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KafkaRequestAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KafkaRequestAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KafkaRequestAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFailedReason

`func (o *KafkaRequestAllOf) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *KafkaRequestAllOf) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *KafkaRequestAllOf) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *KafkaRequestAllOf) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetVersion

`func (o *KafkaRequestAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KafkaRequestAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KafkaRequestAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KafkaRequestAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetInstanceType

`func (o *KafkaRequestAllOf) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *KafkaRequestAllOf) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *KafkaRequestAllOf) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *KafkaRequestAllOf) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceTypeName

`func (o *KafkaRequestAllOf) GetInstanceTypeName() string`

GetInstanceTypeName returns the InstanceTypeName field if non-nil, zero value otherwise.

### GetInstanceTypeNameOk

`func (o *KafkaRequestAllOf) GetInstanceTypeNameOk() (*string, bool)`

GetInstanceTypeNameOk returns a tuple with the InstanceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeName

`func (o *KafkaRequestAllOf) SetInstanceTypeName(v string)`

SetInstanceTypeName sets InstanceTypeName field to given value.

### HasInstanceTypeName

`func (o *KafkaRequestAllOf) HasInstanceTypeName() bool`

HasInstanceTypeName returns a boolean if a field has been set.

### GetReauthenticationEnabled

`func (o *KafkaRequestAllOf) GetReauthenticationEnabled() bool`

GetReauthenticationEnabled returns the ReauthenticationEnabled field if non-nil, zero value otherwise.

### GetReauthenticationEnabledOk

`func (o *KafkaRequestAllOf) GetReauthenticationEnabledOk() (*bool, bool)`

GetReauthenticationEnabledOk returns a tuple with the ReauthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticationEnabled

`func (o *KafkaRequestAllOf) SetReauthenticationEnabled(v bool)`

SetReauthenticationEnabled sets ReauthenticationEnabled field to given value.


### GetKafkaStorageSize

`func (o *KafkaRequestAllOf) GetKafkaStorageSize() string`

GetKafkaStorageSize returns the KafkaStorageSize field if non-nil, zero value otherwise.

### GetKafkaStorageSizeOk

`func (o *KafkaRequestAllOf) GetKafkaStorageSizeOk() (*string, bool)`

GetKafkaStorageSizeOk returns a tuple with the KafkaStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaStorageSize

`func (o *KafkaRequestAllOf) SetKafkaStorageSize(v string)`

SetKafkaStorageSize sets KafkaStorageSize field to given value.

### HasKafkaStorageSize

`func (o *KafkaRequestAllOf) HasKafkaStorageSize() bool`

HasKafkaStorageSize returns a boolean if a field has been set.

### GetMaxDataRetentionSize

`func (o *KafkaRequestAllOf) GetMaxDataRetentionSize() SupportedKafkaSizeBytesValueItem`

GetMaxDataRetentionSize returns the MaxDataRetentionSize field if non-nil, zero value otherwise.

### GetMaxDataRetentionSizeOk

`func (o *KafkaRequestAllOf) GetMaxDataRetentionSizeOk() (*SupportedKafkaSizeBytesValueItem, bool)`

GetMaxDataRetentionSizeOk returns a tuple with the MaxDataRetentionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionSize

`func (o *KafkaRequestAllOf) SetMaxDataRetentionSize(v SupportedKafkaSizeBytesValueItem)`

SetMaxDataRetentionSize sets MaxDataRetentionSize field to given value.

### HasMaxDataRetentionSize

`func (o *KafkaRequestAllOf) HasMaxDataRetentionSize() bool`

HasMaxDataRetentionSize returns a boolean if a field has been set.

### GetBrowserUrl

`func (o *KafkaRequestAllOf) GetBrowserUrl() string`

GetBrowserUrl returns the BrowserUrl field if non-nil, zero value otherwise.

### GetBrowserUrlOk

`func (o *KafkaRequestAllOf) GetBrowserUrlOk() (*string, bool)`

GetBrowserUrlOk returns a tuple with the BrowserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUrl

`func (o *KafkaRequestAllOf) SetBrowserUrl(v string)`

SetBrowserUrl sets BrowserUrl field to given value.

### HasBrowserUrl

`func (o *KafkaRequestAllOf) HasBrowserUrl() bool`

HasBrowserUrl returns a boolean if a field has been set.

### GetSizeId

`func (o *KafkaRequestAllOf) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *KafkaRequestAllOf) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *KafkaRequestAllOf) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *KafkaRequestAllOf) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### GetIngressThroughputPerSec

`func (o *KafkaRequestAllOf) GetIngressThroughputPerSec() string`

GetIngressThroughputPerSec returns the IngressThroughputPerSec field if non-nil, zero value otherwise.

### GetIngressThroughputPerSecOk

`func (o *KafkaRequestAllOf) GetIngressThroughputPerSecOk() (*string, bool)`

GetIngressThroughputPerSecOk returns a tuple with the IngressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressThroughputPerSec

`func (o *KafkaRequestAllOf) SetIngressThroughputPerSec(v string)`

SetIngressThroughputPerSec sets IngressThroughputPerSec field to given value.

### HasIngressThroughputPerSec

`func (o *KafkaRequestAllOf) HasIngressThroughputPerSec() bool`

HasIngressThroughputPerSec returns a boolean if a field has been set.

### GetEgressThroughputPerSec

`func (o *KafkaRequestAllOf) GetEgressThroughputPerSec() string`

GetEgressThroughputPerSec returns the EgressThroughputPerSec field if non-nil, zero value otherwise.

### GetEgressThroughputPerSecOk

`func (o *KafkaRequestAllOf) GetEgressThroughputPerSecOk() (*string, bool)`

GetEgressThroughputPerSecOk returns a tuple with the EgressThroughputPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressThroughputPerSec

`func (o *KafkaRequestAllOf) SetEgressThroughputPerSec(v string)`

SetEgressThroughputPerSec sets EgressThroughputPerSec field to given value.

### HasEgressThroughputPerSec

`func (o *KafkaRequestAllOf) HasEgressThroughputPerSec() bool`

HasEgressThroughputPerSec returns a boolean if a field has been set.

### GetTotalMaxConnections

`func (o *KafkaRequestAllOf) GetTotalMaxConnections() int32`

GetTotalMaxConnections returns the TotalMaxConnections field if non-nil, zero value otherwise.

### GetTotalMaxConnectionsOk

`func (o *KafkaRequestAllOf) GetTotalMaxConnectionsOk() (*int32, bool)`

GetTotalMaxConnectionsOk returns a tuple with the TotalMaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaxConnections

`func (o *KafkaRequestAllOf) SetTotalMaxConnections(v int32)`

SetTotalMaxConnections sets TotalMaxConnections field to given value.

### HasTotalMaxConnections

`func (o *KafkaRequestAllOf) HasTotalMaxConnections() bool`

HasTotalMaxConnections returns a boolean if a field has been set.

### GetMaxPartitions

`func (o *KafkaRequestAllOf) GetMaxPartitions() int32`

GetMaxPartitions returns the MaxPartitions field if non-nil, zero value otherwise.

### GetMaxPartitionsOk

`func (o *KafkaRequestAllOf) GetMaxPartitionsOk() (*int32, bool)`

GetMaxPartitionsOk returns a tuple with the MaxPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPartitions

`func (o *KafkaRequestAllOf) SetMaxPartitions(v int32)`

SetMaxPartitions sets MaxPartitions field to given value.

### HasMaxPartitions

`func (o *KafkaRequestAllOf) HasMaxPartitions() bool`

HasMaxPartitions returns a boolean if a field has been set.

### GetMaxDataRetentionPeriod

`func (o *KafkaRequestAllOf) GetMaxDataRetentionPeriod() string`

GetMaxDataRetentionPeriod returns the MaxDataRetentionPeriod field if non-nil, zero value otherwise.

### GetMaxDataRetentionPeriodOk

`func (o *KafkaRequestAllOf) GetMaxDataRetentionPeriodOk() (*string, bool)`

GetMaxDataRetentionPeriodOk returns a tuple with the MaxDataRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataRetentionPeriod

`func (o *KafkaRequestAllOf) SetMaxDataRetentionPeriod(v string)`

SetMaxDataRetentionPeriod sets MaxDataRetentionPeriod field to given value.

### HasMaxDataRetentionPeriod

`func (o *KafkaRequestAllOf) HasMaxDataRetentionPeriod() bool`

HasMaxDataRetentionPeriod returns a boolean if a field has been set.

### GetMaxConnectionAttemptsPerSec

`func (o *KafkaRequestAllOf) GetMaxConnectionAttemptsPerSec() int32`

GetMaxConnectionAttemptsPerSec returns the MaxConnectionAttemptsPerSec field if non-nil, zero value otherwise.

### GetMaxConnectionAttemptsPerSecOk

`func (o *KafkaRequestAllOf) GetMaxConnectionAttemptsPerSecOk() (*int32, bool)`

GetMaxConnectionAttemptsPerSecOk returns a tuple with the MaxConnectionAttemptsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionAttemptsPerSec

`func (o *KafkaRequestAllOf) SetMaxConnectionAttemptsPerSec(v int32)`

SetMaxConnectionAttemptsPerSec sets MaxConnectionAttemptsPerSec field to given value.

### HasMaxConnectionAttemptsPerSec

`func (o *KafkaRequestAllOf) HasMaxConnectionAttemptsPerSec() bool`

HasMaxConnectionAttemptsPerSec returns a boolean if a field has been set.

### GetBillingCloudAccountId

`func (o *KafkaRequestAllOf) GetBillingCloudAccountId() string`

GetBillingCloudAccountId returns the BillingCloudAccountId field if non-nil, zero value otherwise.

### GetBillingCloudAccountIdOk

`func (o *KafkaRequestAllOf) GetBillingCloudAccountIdOk() (*string, bool)`

GetBillingCloudAccountIdOk returns a tuple with the BillingCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCloudAccountId

`func (o *KafkaRequestAllOf) SetBillingCloudAccountId(v string)`

SetBillingCloudAccountId sets BillingCloudAccountId field to given value.

### HasBillingCloudAccountId

`func (o *KafkaRequestAllOf) HasBillingCloudAccountId() bool`

HasBillingCloudAccountId returns a boolean if a field has been set.

### GetMarketplace

`func (o *KafkaRequestAllOf) GetMarketplace() string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *KafkaRequestAllOf) GetMarketplaceOk() (*string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *KafkaRequestAllOf) SetMarketplace(v string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *KafkaRequestAllOf) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetBillingModel

`func (o *KafkaRequestAllOf) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *KafkaRequestAllOf) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *KafkaRequestAllOf) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *KafkaRequestAllOf) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


