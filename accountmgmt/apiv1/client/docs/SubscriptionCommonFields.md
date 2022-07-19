# SubscriptionCommonFields

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

## Methods

### NewSubscriptionCommonFields

`func NewSubscriptionCommonFields(managed bool, ) *SubscriptionCommonFields`

NewSubscriptionCommonFields instantiates a new SubscriptionCommonFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCommonFieldsWithDefaults

`func NewSubscriptionCommonFieldsWithDefaults() *SubscriptionCommonFields`

NewSubscriptionCommonFieldsWithDefaults instantiates a new SubscriptionCommonFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SubscriptionCommonFields) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SubscriptionCommonFields) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SubscriptionCommonFields) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SubscriptionCommonFields) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionCommonFields) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionCommonFields) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionCommonFields) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionCommonFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SubscriptionCommonFields) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SubscriptionCommonFields) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SubscriptionCommonFields) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SubscriptionCommonFields) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBillingExpirationDate

`func (o *SubscriptionCommonFields) GetBillingExpirationDate() time.Time`

GetBillingExpirationDate returns the BillingExpirationDate field if non-nil, zero value otherwise.

### GetBillingExpirationDateOk

`func (o *SubscriptionCommonFields) GetBillingExpirationDateOk() (*time.Time, bool)`

GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingExpirationDate

`func (o *SubscriptionCommonFields) SetBillingExpirationDate(v time.Time)`

SetBillingExpirationDate sets BillingExpirationDate field to given value.

### HasBillingExpirationDate

`func (o *SubscriptionCommonFields) HasBillingExpirationDate() bool`

HasBillingExpirationDate returns a boolean if a field has been set.

### GetBillingMarketplaceAccount

`func (o *SubscriptionCommonFields) GetBillingMarketplaceAccount() string`

GetBillingMarketplaceAccount returns the BillingMarketplaceAccount field if non-nil, zero value otherwise.

### GetBillingMarketplaceAccountOk

`func (o *SubscriptionCommonFields) GetBillingMarketplaceAccountOk() (*string, bool)`

GetBillingMarketplaceAccountOk returns a tuple with the BillingMarketplaceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMarketplaceAccount

`func (o *SubscriptionCommonFields) SetBillingMarketplaceAccount(v string)`

SetBillingMarketplaceAccount sets BillingMarketplaceAccount field to given value.

### HasBillingMarketplaceAccount

`func (o *SubscriptionCommonFields) HasBillingMarketplaceAccount() bool`

HasBillingMarketplaceAccount returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *SubscriptionCommonFields) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *SubscriptionCommonFields) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *SubscriptionCommonFields) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *SubscriptionCommonFields) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *SubscriptionCommonFields) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *SubscriptionCommonFields) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *SubscriptionCommonFields) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *SubscriptionCommonFields) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *SubscriptionCommonFields) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *SubscriptionCommonFields) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *SubscriptionCommonFields) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *SubscriptionCommonFields) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *SubscriptionCommonFields) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SubscriptionCommonFields) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SubscriptionCommonFields) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SubscriptionCommonFields) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *SubscriptionCommonFields) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *SubscriptionCommonFields) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *SubscriptionCommonFields) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *SubscriptionCommonFields) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *SubscriptionCommonFields) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *SubscriptionCommonFields) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *SubscriptionCommonFields) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *SubscriptionCommonFields) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *SubscriptionCommonFields) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *SubscriptionCommonFields) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *SubscriptionCommonFields) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *SubscriptionCommonFields) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatorId

`func (o *SubscriptionCommonFields) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SubscriptionCommonFields) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SubscriptionCommonFields) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *SubscriptionCommonFields) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionCommonFields) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionCommonFields) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionCommonFields) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionCommonFields) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *SubscriptionCommonFields) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *SubscriptionCommonFields) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *SubscriptionCommonFields) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *SubscriptionCommonFields) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetLastReconcileDate

`func (o *SubscriptionCommonFields) GetLastReconcileDate() time.Time`

GetLastReconcileDate returns the LastReconcileDate field if non-nil, zero value otherwise.

### GetLastReconcileDateOk

`func (o *SubscriptionCommonFields) GetLastReconcileDateOk() (*time.Time, bool)`

GetLastReconcileDateOk returns a tuple with the LastReconcileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileDate

`func (o *SubscriptionCommonFields) SetLastReconcileDate(v time.Time)`

SetLastReconcileDate sets LastReconcileDate field to given value.

### HasLastReconcileDate

`func (o *SubscriptionCommonFields) HasLastReconcileDate() bool`

HasLastReconcileDate returns a boolean if a field has been set.

### GetLastReleasedAt

`func (o *SubscriptionCommonFields) GetLastReleasedAt() time.Time`

GetLastReleasedAt returns the LastReleasedAt field if non-nil, zero value otherwise.

### GetLastReleasedAtOk

`func (o *SubscriptionCommonFields) GetLastReleasedAtOk() (*time.Time, bool)`

GetLastReleasedAtOk returns a tuple with the LastReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleasedAt

`func (o *SubscriptionCommonFields) SetLastReleasedAt(v time.Time)`

SetLastReleasedAt sets LastReleasedAt field to given value.

### HasLastReleasedAt

`func (o *SubscriptionCommonFields) HasLastReleasedAt() bool`

HasLastReleasedAt returns a boolean if a field has been set.

### GetLastTelemetryDate

`func (o *SubscriptionCommonFields) GetLastTelemetryDate() time.Time`

GetLastTelemetryDate returns the LastTelemetryDate field if non-nil, zero value otherwise.

### GetLastTelemetryDateOk

`func (o *SubscriptionCommonFields) GetLastTelemetryDateOk() (*time.Time, bool)`

GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTelemetryDate

`func (o *SubscriptionCommonFields) SetLastTelemetryDate(v time.Time)`

SetLastTelemetryDate sets LastTelemetryDate field to given value.

### HasLastTelemetryDate

`func (o *SubscriptionCommonFields) HasLastTelemetryDate() bool`

HasLastTelemetryDate returns a boolean if a field has been set.

### GetManaged

`func (o *SubscriptionCommonFields) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SubscriptionCommonFields) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SubscriptionCommonFields) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOrganizationId

`func (o *SubscriptionCommonFields) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionCommonFields) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionCommonFields) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionCommonFields) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionCommonFields) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionCommonFields) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionCommonFields) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionCommonFields) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductBundle

`func (o *SubscriptionCommonFields) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *SubscriptionCommonFields) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *SubscriptionCommonFields) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *SubscriptionCommonFields) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *SubscriptionCommonFields) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *SubscriptionCommonFields) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *SubscriptionCommonFields) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *SubscriptionCommonFields) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *SubscriptionCommonFields) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SubscriptionCommonFields) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SubscriptionCommonFields) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SubscriptionCommonFields) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *SubscriptionCommonFields) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *SubscriptionCommonFields) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *SubscriptionCommonFields) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *SubscriptionCommonFields) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *SubscriptionCommonFields) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *SubscriptionCommonFields) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *SubscriptionCommonFields) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *SubscriptionCommonFields) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *SubscriptionCommonFields) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *SubscriptionCommonFields) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *SubscriptionCommonFields) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *SubscriptionCommonFields) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionCommonFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCommonFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCommonFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionCommonFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *SubscriptionCommonFields) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *SubscriptionCommonFields) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *SubscriptionCommonFields) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *SubscriptionCommonFields) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *SubscriptionCommonFields) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *SubscriptionCommonFields) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *SubscriptionCommonFields) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *SubscriptionCommonFields) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *SubscriptionCommonFields) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *SubscriptionCommonFields) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *SubscriptionCommonFields) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *SubscriptionCommonFields) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsage

`func (o *SubscriptionCommonFields) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionCommonFields) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionCommonFields) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SubscriptionCommonFields) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


