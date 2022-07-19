# SubscriptionCommonFieldsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingExpirationDate** | Pointer to **time.Time** | If set, the date the subscription expires based on the billing model | [optional] 
**BillingMarketplaceAccount** | Pointer to **string** |  | [optional] 
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudProviderId** | Pointer to **string** |  | [optional] 
**ClusterBillingModel** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ConsoleUrl** | Pointer to **string** |  | [optional] 
**ConsumerUuid** | Pointer to **string** |  | [optional] 
**CpuTotal** | Pointer to **int32** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalClusterId** | Pointer to **string** |  | [optional] 
**LastReconcileDate** | Pointer to **time.Time** | Last time this subscription were reconciled about cluster usage | [optional] 
**LastReleasedAt** | Pointer to **time.Time** | Last time status was set to Released for this cluster/subscription in Unix time | [optional] 
**LastTelemetryDate** | Pointer to **time.Time** | Last telemetry authorization request for this cluster/subscription in Unix time | [optional] 
**Managed** | **bool** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**ProductBundle** | Pointer to **string** |  | [optional] 
**Provenance** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Released** | Pointer to **bool** |  | [optional] 
**ServiceLevel** | Pointer to **string** |  | [optional] 
**SocketTotal** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportLevel** | Pointer to **string** |  | [optional] 
**SystemUnits** | Pointer to **string** |  | [optional] 
**TrialEndDate** | Pointer to **time.Time** | If the subscription is a trial, date the trial ends | [optional] 
**Usage** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionCommonFieldsAllOf

`func NewSubscriptionCommonFieldsAllOf(managed bool, ) *SubscriptionCommonFieldsAllOf`

NewSubscriptionCommonFieldsAllOf instantiates a new SubscriptionCommonFieldsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCommonFieldsAllOfWithDefaults

`func NewSubscriptionCommonFieldsAllOfWithDefaults() *SubscriptionCommonFieldsAllOf`

NewSubscriptionCommonFieldsAllOfWithDefaults instantiates a new SubscriptionCommonFieldsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingExpirationDate

`func (o *SubscriptionCommonFieldsAllOf) GetBillingExpirationDate() time.Time`

GetBillingExpirationDate returns the BillingExpirationDate field if non-nil, zero value otherwise.

### GetBillingExpirationDateOk

`func (o *SubscriptionCommonFieldsAllOf) GetBillingExpirationDateOk() (*time.Time, bool)`

GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingExpirationDate

`func (o *SubscriptionCommonFieldsAllOf) SetBillingExpirationDate(v time.Time)`

SetBillingExpirationDate sets BillingExpirationDate field to given value.

### HasBillingExpirationDate

`func (o *SubscriptionCommonFieldsAllOf) HasBillingExpirationDate() bool`

HasBillingExpirationDate returns a boolean if a field has been set.

### GetBillingMarketplaceAccount

`func (o *SubscriptionCommonFieldsAllOf) GetBillingMarketplaceAccount() string`

GetBillingMarketplaceAccount returns the BillingMarketplaceAccount field if non-nil, zero value otherwise.

### GetBillingMarketplaceAccountOk

`func (o *SubscriptionCommonFieldsAllOf) GetBillingMarketplaceAccountOk() (*string, bool)`

GetBillingMarketplaceAccountOk returns a tuple with the BillingMarketplaceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMarketplaceAccount

`func (o *SubscriptionCommonFieldsAllOf) SetBillingMarketplaceAccount(v string)`

SetBillingMarketplaceAccount sets BillingMarketplaceAccount field to given value.

### HasBillingMarketplaceAccount

`func (o *SubscriptionCommonFieldsAllOf) HasBillingMarketplaceAccount() bool`

HasBillingMarketplaceAccount returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *SubscriptionCommonFieldsAllOf) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *SubscriptionCommonFieldsAllOf) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *SubscriptionCommonFieldsAllOf) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *SubscriptionCommonFieldsAllOf) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *SubscriptionCommonFieldsAllOf) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *SubscriptionCommonFieldsAllOf) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *SubscriptionCommonFieldsAllOf) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *SubscriptionCommonFieldsAllOf) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *SubscriptionCommonFieldsAllOf) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *SubscriptionCommonFieldsAllOf) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *SubscriptionCommonFieldsAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SubscriptionCommonFieldsAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SubscriptionCommonFieldsAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *SubscriptionCommonFieldsAllOf) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *SubscriptionCommonFieldsAllOf) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *SubscriptionCommonFieldsAllOf) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *SubscriptionCommonFieldsAllOf) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *SubscriptionCommonFieldsAllOf) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *SubscriptionCommonFieldsAllOf) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *SubscriptionCommonFieldsAllOf) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *SubscriptionCommonFieldsAllOf) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *SubscriptionCommonFieldsAllOf) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *SubscriptionCommonFieldsAllOf) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *SubscriptionCommonFieldsAllOf) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *SubscriptionCommonFieldsAllOf) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatorId

`func (o *SubscriptionCommonFieldsAllOf) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SubscriptionCommonFieldsAllOf) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *SubscriptionCommonFieldsAllOf) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionCommonFieldsAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionCommonFieldsAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionCommonFieldsAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionCommonFieldsAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *SubscriptionCommonFieldsAllOf) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *SubscriptionCommonFieldsAllOf) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *SubscriptionCommonFieldsAllOf) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetLastReconcileDate

`func (o *SubscriptionCommonFieldsAllOf) GetLastReconcileDate() time.Time`

GetLastReconcileDate returns the LastReconcileDate field if non-nil, zero value otherwise.

### GetLastReconcileDateOk

`func (o *SubscriptionCommonFieldsAllOf) GetLastReconcileDateOk() (*time.Time, bool)`

GetLastReconcileDateOk returns a tuple with the LastReconcileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileDate

`func (o *SubscriptionCommonFieldsAllOf) SetLastReconcileDate(v time.Time)`

SetLastReconcileDate sets LastReconcileDate field to given value.

### HasLastReconcileDate

`func (o *SubscriptionCommonFieldsAllOf) HasLastReconcileDate() bool`

HasLastReconcileDate returns a boolean if a field has been set.

### GetLastReleasedAt

`func (o *SubscriptionCommonFieldsAllOf) GetLastReleasedAt() time.Time`

GetLastReleasedAt returns the LastReleasedAt field if non-nil, zero value otherwise.

### GetLastReleasedAtOk

`func (o *SubscriptionCommonFieldsAllOf) GetLastReleasedAtOk() (*time.Time, bool)`

GetLastReleasedAtOk returns a tuple with the LastReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleasedAt

`func (o *SubscriptionCommonFieldsAllOf) SetLastReleasedAt(v time.Time)`

SetLastReleasedAt sets LastReleasedAt field to given value.

### HasLastReleasedAt

`func (o *SubscriptionCommonFieldsAllOf) HasLastReleasedAt() bool`

HasLastReleasedAt returns a boolean if a field has been set.

### GetLastTelemetryDate

`func (o *SubscriptionCommonFieldsAllOf) GetLastTelemetryDate() time.Time`

GetLastTelemetryDate returns the LastTelemetryDate field if non-nil, zero value otherwise.

### GetLastTelemetryDateOk

`func (o *SubscriptionCommonFieldsAllOf) GetLastTelemetryDateOk() (*time.Time, bool)`

GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTelemetryDate

`func (o *SubscriptionCommonFieldsAllOf) SetLastTelemetryDate(v time.Time)`

SetLastTelemetryDate sets LastTelemetryDate field to given value.

### HasLastTelemetryDate

`func (o *SubscriptionCommonFieldsAllOf) HasLastTelemetryDate() bool`

HasLastTelemetryDate returns a boolean if a field has been set.

### GetManaged

`func (o *SubscriptionCommonFieldsAllOf) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SubscriptionCommonFieldsAllOf) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SubscriptionCommonFieldsAllOf) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOrganizationId

`func (o *SubscriptionCommonFieldsAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionCommonFieldsAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionCommonFieldsAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionCommonFieldsAllOf) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionCommonFieldsAllOf) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionCommonFieldsAllOf) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductBundle

`func (o *SubscriptionCommonFieldsAllOf) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *SubscriptionCommonFieldsAllOf) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *SubscriptionCommonFieldsAllOf) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *SubscriptionCommonFieldsAllOf) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *SubscriptionCommonFieldsAllOf) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *SubscriptionCommonFieldsAllOf) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *SubscriptionCommonFieldsAllOf) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *SubscriptionCommonFieldsAllOf) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *SubscriptionCommonFieldsAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SubscriptionCommonFieldsAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SubscriptionCommonFieldsAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SubscriptionCommonFieldsAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *SubscriptionCommonFieldsAllOf) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *SubscriptionCommonFieldsAllOf) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *SubscriptionCommonFieldsAllOf) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *SubscriptionCommonFieldsAllOf) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *SubscriptionCommonFieldsAllOf) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *SubscriptionCommonFieldsAllOf) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *SubscriptionCommonFieldsAllOf) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *SubscriptionCommonFieldsAllOf) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *SubscriptionCommonFieldsAllOf) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *SubscriptionCommonFieldsAllOf) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *SubscriptionCommonFieldsAllOf) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *SubscriptionCommonFieldsAllOf) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionCommonFieldsAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCommonFieldsAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCommonFieldsAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionCommonFieldsAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *SubscriptionCommonFieldsAllOf) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *SubscriptionCommonFieldsAllOf) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *SubscriptionCommonFieldsAllOf) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *SubscriptionCommonFieldsAllOf) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *SubscriptionCommonFieldsAllOf) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *SubscriptionCommonFieldsAllOf) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *SubscriptionCommonFieldsAllOf) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *SubscriptionCommonFieldsAllOf) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *SubscriptionCommonFieldsAllOf) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *SubscriptionCommonFieldsAllOf) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *SubscriptionCommonFieldsAllOf) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *SubscriptionCommonFieldsAllOf) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsage

`func (o *SubscriptionCommonFieldsAllOf) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionCommonFieldsAllOf) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionCommonFieldsAllOf) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SubscriptionCommonFieldsAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


