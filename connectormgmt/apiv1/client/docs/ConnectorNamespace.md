# ConnectorNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ResourceVersion** | **int64** |  | 
**Quota** | Pointer to [**ConnectorNamespaceQuota**](ConnectorNamespaceQuota.md) |  | [optional] 
**ClusterId** | **string** |  | 
**Expiration** | Pointer to **string** | Namespace expiration timestamp in RFC 3339 format | [optional] 
**Tenant** | [**ConnectorNamespaceTenant**](ConnectorNamespaceTenant.md) |  | 
**Status** | [**ConnectorNamespaceStatus**](ConnectorNamespaceStatus.md) |  | 

## Methods

### NewConnectorNamespace

`func NewConnectorNamespace(id string, name string, resourceVersion int64, clusterId string, tenant ConnectorNamespaceTenant, status ConnectorNamespaceStatus, ) *ConnectorNamespace`

NewConnectorNamespace instantiates a new ConnectorNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceWithDefaults

`func NewConnectorNamespaceWithDefaults() *ConnectorNamespace`

NewConnectorNamespaceWithDefaults instantiates a new ConnectorNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorNamespace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorNamespace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorNamespace) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ConnectorNamespace) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorNamespace) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorNamespace) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectorNamespace) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *ConnectorNamespace) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectorNamespace) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectorNamespace) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectorNamespace) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetOwner

`func (o *ConnectorNamespace) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConnectorNamespace) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConnectorNamespace) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConnectorNamespace) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorNamespace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorNamespace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorNamespace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorNamespace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ConnectorNamespace) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorNamespace) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorNamespace) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ConnectorNamespace) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *ConnectorNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespace) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *ConnectorNamespace) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorNamespace) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorNamespace) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorNamespace) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ConnectorNamespace) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ConnectorNamespace) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ConnectorNamespace) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.


### GetQuota

`func (o *ConnectorNamespace) GetQuota() ConnectorNamespaceQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ConnectorNamespace) GetQuotaOk() (*ConnectorNamespaceQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ConnectorNamespace) SetQuota(v ConnectorNamespaceQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ConnectorNamespace) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetClusterId

`func (o *ConnectorNamespace) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConnectorNamespace) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConnectorNamespace) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetExpiration

`func (o *ConnectorNamespace) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ConnectorNamespace) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ConnectorNamespace) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ConnectorNamespace) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetTenant

`func (o *ConnectorNamespace) GetTenant() ConnectorNamespaceTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ConnectorNamespace) GetTenantOk() (*ConnectorNamespaceTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ConnectorNamespace) SetTenant(v ConnectorNamespaceTenant)`

SetTenant sets Tenant field to given value.


### GetStatus

`func (o *ConnectorNamespace) GetStatus() ConnectorNamespaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorNamespace) GetStatusOk() (*ConnectorNamespaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorNamespace) SetStatus(v ConnectorNamespaceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


