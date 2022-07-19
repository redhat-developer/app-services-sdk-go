# ConsumedQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AvailabilityZoneType** | Pointer to **string** |  | [optional] 
**BillingModel** | Pointer to **string** |  | [optional] 
**Byoc** | **bool** |  | 
**CloudProviderId** | Pointer to **string** |  | [optional] 
**Count** | **int32** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewConsumedQuota

`func NewConsumedQuota(byoc bool, count int32, ) *ConsumedQuota`

NewConsumedQuota instantiates a new ConsumedQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumedQuotaWithDefaults

`func NewConsumedQuotaWithDefaults() *ConsumedQuota`

NewConsumedQuotaWithDefaults instantiates a new ConsumedQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ConsumedQuota) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConsumedQuota) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConsumedQuota) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConsumedQuota) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ConsumedQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumedQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumedQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumedQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ConsumedQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumedQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumedQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConsumedQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAvailabilityZoneType

`func (o *ConsumedQuota) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *ConsumedQuota) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *ConsumedQuota) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *ConsumedQuota) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetBillingModel

`func (o *ConsumedQuota) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *ConsumedQuota) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *ConsumedQuota) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *ConsumedQuota) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetByoc

`func (o *ConsumedQuota) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ConsumedQuota) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ConsumedQuota) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetCloudProviderId

`func (o *ConsumedQuota) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *ConsumedQuota) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *ConsumedQuota) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *ConsumedQuota) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetCount

`func (o *ConsumedQuota) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConsumedQuota) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConsumedQuota) SetCount(v int32)`

SetCount sets Count field to given value.


### GetOrganizationId

`func (o *ConsumedQuota) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ConsumedQuota) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ConsumedQuota) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ConsumedQuota) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlanId

`func (o *ConsumedQuota) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ConsumedQuota) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ConsumedQuota) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ConsumedQuota) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetResourceName

`func (o *ConsumedQuota) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ConsumedQuota) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ConsumedQuota) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ConsumedQuota) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *ConsumedQuota) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ConsumedQuota) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ConsumedQuota) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ConsumedQuota) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetVersion

`func (o *ConsumedQuota) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsumedQuota) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsumedQuota) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsumedQuota) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


