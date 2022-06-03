# ConnectorNamespaceMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersion** | Pointer to **int64** |  | [optional] 
**Quota** | Pointer to [**ConnectorNamespaceQuota**](ConnectorNamespaceQuota.md) |  | [optional] 

## Methods

### NewConnectorNamespaceMetaAllOf

`func NewConnectorNamespaceMetaAllOf() *ConnectorNamespaceMetaAllOf`

NewConnectorNamespaceMetaAllOf instantiates a new ConnectorNamespaceMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceMetaAllOfWithDefaults

`func NewConnectorNamespaceMetaAllOfWithDefaults() *ConnectorNamespaceMetaAllOf`

NewConnectorNamespaceMetaAllOfWithDefaults instantiates a new ConnectorNamespaceMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVersion

`func (o *ConnectorNamespaceMetaAllOf) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ConnectorNamespaceMetaAllOf) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ConnectorNamespaceMetaAllOf) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ConnectorNamespaceMetaAllOf) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetQuota

`func (o *ConnectorNamespaceMetaAllOf) GetQuota() ConnectorNamespaceQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ConnectorNamespaceMetaAllOf) GetQuotaOk() (*ConnectorNamespaceQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ConnectorNamespaceMetaAllOf) SetQuota(v ConnectorNamespaceQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ConnectorNamespaceMetaAllOf) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


