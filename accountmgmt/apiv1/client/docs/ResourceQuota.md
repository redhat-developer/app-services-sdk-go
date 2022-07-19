# ResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**SkuCount** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResourceQuota

`func NewResourceQuota(skuCount int32, ) *ResourceQuota`

NewResourceQuota instantiates a new ResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuotaWithDefaults

`func NewResourceQuotaWithDefaults() *ResourceQuota`

NewResourceQuotaWithDefaults instantiates a new ResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ResourceQuota) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceQuota) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceQuota) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceQuota) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ResourceQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ResourceQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ResourceQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceQuota) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceQuota) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceQuota) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceQuota) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ResourceQuota) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ResourceQuota) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ResourceQuota) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ResourceQuota) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSku

`func (o *ResourceQuota) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ResourceQuota) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ResourceQuota) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ResourceQuota) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetSkuCount

`func (o *ResourceQuota) GetSkuCount() int32`

GetSkuCount returns the SkuCount field if non-nil, zero value otherwise.

### GetSkuCountOk

`func (o *ResourceQuota) GetSkuCountOk() (*int32, bool)`

GetSkuCountOk returns a tuple with the SkuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCount

`func (o *ResourceQuota) SetSkuCount(v int32)`

SetSkuCount sets SkuCount field to given value.


### GetType

`func (o *ResourceQuota) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceQuota) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceQuota) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceQuota) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceQuota) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceQuota) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceQuota) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceQuota) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


