# DeletedSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
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
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Metrics** | Pointer to **string** |  | [optional] 
**QueryTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeletedSubscription

`func NewDeletedSubscription(managed bool, ) *DeletedSubscription`

NewDeletedSubscription instantiates a new DeletedSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedSubscriptionWithDefaults

`func NewDeletedSubscriptionWithDefaults() *DeletedSubscription`

NewDeletedSubscriptionWithDefaults instantiates a new DeletedSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *DeletedSubscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DeletedSubscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DeletedSubscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DeletedSubscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *DeletedSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DeletedSubscription) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeletedSubscription) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeletedSubscription) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeletedSubscription) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBillingExpirationDate

`func (o *DeletedSubscription) GetBillingExpirationDate() time.Time`

GetBillingExpirationDate returns the BillingExpirationDate field if non-nil, zero value otherwise.

### GetBillingExpirationDateOk

`func (o *DeletedSubscription) GetBillingExpirationDateOk() (*time.Time, bool)`

GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingExpirationDate

`func (o *DeletedSubscription) SetBillingExpirationDate(v time.Time)`

SetBillingExpirationDate sets BillingExpirationDate field to given value.

### HasBillingExpirationDate

`func (o *DeletedSubscription) HasBillingExpirationDate() bool`

HasBillingExpirationDate returns a boolean if a field has been set.

### GetBillingMarketplaceAccount

`func (o *DeletedSubscription) GetBillingMarketplaceAccount() string`

GetBillingMarketplaceAccount returns the BillingMarketplaceAccount field if non-nil, zero value otherwise.

### GetBillingMarketplaceAccountOk

`func (o *DeletedSubscription) GetBillingMarketplaceAccountOk() (*string, bool)`

GetBillingMarketplaceAccountOk returns a tuple with the BillingMarketplaceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMarketplaceAccount

`func (o *DeletedSubscription) SetBillingMarketplaceAccount(v string)`

SetBillingMarketplaceAccount sets BillingMarketplaceAccount field to given value.

### HasBillingMarketplaceAccount

`func (o *DeletedSubscription) HasBillingMarketplaceAccount() bool`

HasBillingMarketplaceAccount returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *DeletedSubscription) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *DeletedSubscription) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *DeletedSubscription) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *DeletedSubscription) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DeletedSubscription) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DeletedSubscription) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DeletedSubscription) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *DeletedSubscription) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *DeletedSubscription) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *DeletedSubscription) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *DeletedSubscription) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *DeletedSubscription) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *DeletedSubscription) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DeletedSubscription) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DeletedSubscription) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DeletedSubscription) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *DeletedSubscription) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *DeletedSubscription) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *DeletedSubscription) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *DeletedSubscription) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *DeletedSubscription) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *DeletedSubscription) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *DeletedSubscription) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *DeletedSubscription) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *DeletedSubscription) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *DeletedSubscription) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *DeletedSubscription) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *DeletedSubscription) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatorId

`func (o *DeletedSubscription) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DeletedSubscription) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DeletedSubscription) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *DeletedSubscription) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisplayName

`func (o *DeletedSubscription) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeletedSubscription) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeletedSubscription) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeletedSubscription) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *DeletedSubscription) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *DeletedSubscription) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *DeletedSubscription) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *DeletedSubscription) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetLastReconcileDate

`func (o *DeletedSubscription) GetLastReconcileDate() time.Time`

GetLastReconcileDate returns the LastReconcileDate field if non-nil, zero value otherwise.

### GetLastReconcileDateOk

`func (o *DeletedSubscription) GetLastReconcileDateOk() (*time.Time, bool)`

GetLastReconcileDateOk returns a tuple with the LastReconcileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileDate

`func (o *DeletedSubscription) SetLastReconcileDate(v time.Time)`

SetLastReconcileDate sets LastReconcileDate field to given value.

### HasLastReconcileDate

`func (o *DeletedSubscription) HasLastReconcileDate() bool`

HasLastReconcileDate returns a boolean if a field has been set.

### GetLastReleasedAt

`func (o *DeletedSubscription) GetLastReleasedAt() time.Time`

GetLastReleasedAt returns the LastReleasedAt field if non-nil, zero value otherwise.

### GetLastReleasedAtOk

`func (o *DeletedSubscription) GetLastReleasedAtOk() (*time.Time, bool)`

GetLastReleasedAtOk returns a tuple with the LastReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleasedAt

`func (o *DeletedSubscription) SetLastReleasedAt(v time.Time)`

SetLastReleasedAt sets LastReleasedAt field to given value.

### HasLastReleasedAt

`func (o *DeletedSubscription) HasLastReleasedAt() bool`

HasLastReleasedAt returns a boolean if a field has been set.

### GetLastTelemetryDate

`func (o *DeletedSubscription) GetLastTelemetryDate() time.Time`

GetLastTelemetryDate returns the LastTelemetryDate field if non-nil, zero value otherwise.

### GetLastTelemetryDateOk

`func (o *DeletedSubscription) GetLastTelemetryDateOk() (*time.Time, bool)`

GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTelemetryDate

`func (o *DeletedSubscription) SetLastTelemetryDate(v time.Time)`

SetLastTelemetryDate sets LastTelemetryDate field to given value.

### HasLastTelemetryDate

`func (o *DeletedSubscription) HasLastTelemetryDate() bool`

HasLastTelemetryDate returns a boolean if a field has been set.

### GetManaged

`func (o *DeletedSubscription) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeletedSubscription) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeletedSubscription) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOrganizationId

`func (o *DeletedSubscription) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeletedSubscription) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeletedSubscription) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeletedSubscription) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlanId

`func (o *DeletedSubscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DeletedSubscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DeletedSubscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DeletedSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductBundle

`func (o *DeletedSubscription) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *DeletedSubscription) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *DeletedSubscription) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *DeletedSubscription) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *DeletedSubscription) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *DeletedSubscription) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *DeletedSubscription) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *DeletedSubscription) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *DeletedSubscription) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *DeletedSubscription) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *DeletedSubscription) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *DeletedSubscription) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *DeletedSubscription) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *DeletedSubscription) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *DeletedSubscription) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *DeletedSubscription) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *DeletedSubscription) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *DeletedSubscription) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *DeletedSubscription) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *DeletedSubscription) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *DeletedSubscription) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *DeletedSubscription) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *DeletedSubscription) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *DeletedSubscription) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *DeletedSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeletedSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeletedSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeletedSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *DeletedSubscription) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *DeletedSubscription) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *DeletedSubscription) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *DeletedSubscription) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *DeletedSubscription) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *DeletedSubscription) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *DeletedSubscription) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *DeletedSubscription) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *DeletedSubscription) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *DeletedSubscription) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *DeletedSubscription) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *DeletedSubscription) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsage

`func (o *DeletedSubscription) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DeletedSubscription) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DeletedSubscription) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DeletedSubscription) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeletedSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeletedSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeletedSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeletedSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetrics

`func (o *DeletedSubscription) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DeletedSubscription) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DeletedSubscription) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DeletedSubscription) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetQueryTimestamp

`func (o *DeletedSubscription) GetQueryTimestamp() time.Time`

GetQueryTimestamp returns the QueryTimestamp field if non-nil, zero value otherwise.

### GetQueryTimestampOk

`func (o *DeletedSubscription) GetQueryTimestampOk() (*time.Time, bool)`

GetQueryTimestampOk returns a tuple with the QueryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTimestamp

`func (o *DeletedSubscription) SetQueryTimestamp(v time.Time)`

SetQueryTimestamp sets QueryTimestamp field to given value.

### HasQueryTimestamp

`func (o *DeletedSubscription) HasQueryTimestamp() bool`

HasQueryTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


