# ConnectorNamespaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ClusterId** | **string** |  | 
**Expiration** | Pointer to **string** | Namespace expiration timestamp in RFC 3339 format | [optional] 
**Tenant** | [**ConnectorNamespaceTenant**](ConnectorNamespaceTenant.md) |  | 
**Status** | [**ConnectorNamespaceStatus**](ConnectorNamespaceStatus.md) |  | 

## Methods

### NewConnectorNamespaceAllOf

`func NewConnectorNamespaceAllOf(name string, clusterId string, tenant ConnectorNamespaceTenant, status ConnectorNamespaceStatus, ) *ConnectorNamespaceAllOf`

NewConnectorNamespaceAllOf instantiates a new ConnectorNamespaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceAllOfWithDefaults

`func NewConnectorNamespaceAllOfWithDefaults() *ConnectorNamespaceAllOf`

NewConnectorNamespaceAllOfWithDefaults instantiates a new ConnectorNamespaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorNamespaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespaceAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetClusterId

`func (o *ConnectorNamespaceAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConnectorNamespaceAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConnectorNamespaceAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetExpiration

`func (o *ConnectorNamespaceAllOf) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ConnectorNamespaceAllOf) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ConnectorNamespaceAllOf) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ConnectorNamespaceAllOf) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetTenant

`func (o *ConnectorNamespaceAllOf) GetTenant() ConnectorNamespaceTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ConnectorNamespaceAllOf) GetTenantOk() (*ConnectorNamespaceTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ConnectorNamespaceAllOf) SetTenant(v ConnectorNamespaceTenant)`

SetTenant sets Tenant field to given value.


### GetStatus

`func (o *ConnectorNamespaceAllOf) GetStatus() ConnectorNamespaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorNamespaceAllOf) GetStatusOk() (*ConnectorNamespaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorNamespaceAllOf) SetStatus(v ConnectorNamespaceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


