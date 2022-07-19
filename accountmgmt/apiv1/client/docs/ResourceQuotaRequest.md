# ResourceQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** |  | 
**SkuCount** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceQuotaRequest

`func NewResourceQuotaRequest(sku string, skuCount int32, ) *ResourceQuotaRequest`

NewResourceQuotaRequest instantiates a new ResourceQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuotaRequestWithDefaults

`func NewResourceQuotaRequestWithDefaults() *ResourceQuotaRequest`

NewResourceQuotaRequestWithDefaults instantiates a new ResourceQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *ResourceQuotaRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ResourceQuotaRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ResourceQuotaRequest) SetSku(v string)`

SetSku sets Sku field to given value.


### GetSkuCount

`func (o *ResourceQuotaRequest) GetSkuCount() int32`

GetSkuCount returns the SkuCount field if non-nil, zero value otherwise.

### GetSkuCountOk

`func (o *ResourceQuotaRequest) GetSkuCountOk() (*int32, bool)`

GetSkuCountOk returns a tuple with the SkuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCount

`func (o *ResourceQuotaRequest) SetSkuCount(v int32)`

SetSkuCount sets SkuCount field to given value.


### GetType

`func (o *ResourceQuotaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceQuotaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceQuotaRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceQuotaRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


