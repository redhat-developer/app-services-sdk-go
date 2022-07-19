# QuotaRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**BillingModel** | Pointer to **string** |  | [optional] 
**Byoc** | Pointer to **string** |  | [optional] 
**Cloud** | Pointer to **string** |  | [optional] 
**Cost** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**QuotaId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewQuotaRules

`func NewQuotaRules(cost int32, ) *QuotaRules`

NewQuotaRules instantiates a new QuotaRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaRulesWithDefaults

`func NewQuotaRulesWithDefaults() *QuotaRules`

NewQuotaRulesWithDefaults instantiates a new QuotaRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *QuotaRules) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *QuotaRules) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *QuotaRules) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *QuotaRules) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *QuotaRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuotaRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *QuotaRules) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotaRules) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotaRules) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QuotaRules) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *QuotaRules) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *QuotaRules) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *QuotaRules) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *QuotaRules) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetBillingModel

`func (o *QuotaRules) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *QuotaRules) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *QuotaRules) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *QuotaRules) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetByoc

`func (o *QuotaRules) GetByoc() string`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *QuotaRules) GetByocOk() (*string, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *QuotaRules) SetByoc(v string)`

SetByoc sets Byoc field to given value.

### HasByoc

`func (o *QuotaRules) HasByoc() bool`

HasByoc returns a boolean if a field has been set.

### GetCloud

`func (o *QuotaRules) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *QuotaRules) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *QuotaRules) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *QuotaRules) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCost

`func (o *QuotaRules) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *QuotaRules) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *QuotaRules) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetName

`func (o *QuotaRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuotaRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuotaRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuotaRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProduct

`func (o *QuotaRules) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *QuotaRules) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *QuotaRules) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *QuotaRules) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetQuotaId

`func (o *QuotaRules) GetQuotaId() string`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *QuotaRules) GetQuotaIdOk() (*string, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *QuotaRules) SetQuotaId(v string)`

SetQuotaId sets QuotaId field to given value.

### HasQuotaId

`func (o *QuotaRules) HasQuotaId() bool`

HasQuotaId returns a boolean if a field has been set.

### GetType

`func (o *QuotaRules) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotaRules) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotaRules) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QuotaRules) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


