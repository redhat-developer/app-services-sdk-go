# ConnectorNamespaceMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | Namespace name must match pattern &#x60;^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$&#x60;, or it may be empty to be auto-generated. | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ResourceVersion** | Pointer to **int64** |  | [optional] 
**Quota** | Pointer to [**ConnectorNamespaceQuota**](ConnectorNamespaceQuota.md) |  | [optional] 

## Methods

### NewConnectorNamespaceMeta

`func NewConnectorNamespaceMeta() *ConnectorNamespaceMeta`

NewConnectorNamespaceMeta instantiates a new ConnectorNamespaceMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceMetaWithDefaults

`func NewConnectorNamespaceMetaWithDefaults() *ConnectorNamespaceMeta`

NewConnectorNamespaceMetaWithDefaults instantiates a new ConnectorNamespaceMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ConnectorNamespaceMeta) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConnectorNamespaceMeta) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConnectorNamespaceMeta) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConnectorNamespaceMeta) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorNamespaceMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorNamespaceMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorNamespaceMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorNamespaceMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ConnectorNamespaceMeta) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorNamespaceMeta) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorNamespaceMeta) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ConnectorNamespaceMeta) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *ConnectorNamespaceMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespaceMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespaceMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorNamespaceMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConnectorNamespaceMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorNamespaceMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorNamespaceMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorNamespaceMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ConnectorNamespaceMeta) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ConnectorNamespaceMeta) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ConnectorNamespaceMeta) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ConnectorNamespaceMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetQuota

`func (o *ConnectorNamespaceMeta) GetQuota() ConnectorNamespaceQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ConnectorNamespaceMeta) GetQuotaOk() (*ConnectorNamespaceQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ConnectorNamespaceMeta) SetQuota(v ConnectorNamespaceQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ConnectorNamespaceMeta) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


