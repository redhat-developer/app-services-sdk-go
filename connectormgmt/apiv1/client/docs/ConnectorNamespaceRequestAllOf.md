# ConnectorNamespaceRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to [**ConnectorNamespaceTenantKind**](ConnectorNamespaceTenantKind.md) |  | [optional] 

## Methods

### NewConnectorNamespaceRequestAllOf

`func NewConnectorNamespaceRequestAllOf() *ConnectorNamespaceRequestAllOf`

NewConnectorNamespaceRequestAllOf instantiates a new ConnectorNamespaceRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceRequestAllOfWithDefaults

`func NewConnectorNamespaceRequestAllOfWithDefaults() *ConnectorNamespaceRequestAllOf`

NewConnectorNamespaceRequestAllOfWithDefaults instantiates a new ConnectorNamespaceRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ConnectorNamespaceRequestAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConnectorNamespaceRequestAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConnectorNamespaceRequestAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ConnectorNamespaceRequestAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetKind

`func (o *ConnectorNamespaceRequestAllOf) GetKind() ConnectorNamespaceTenantKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorNamespaceRequestAllOf) GetKindOk() (*ConnectorNamespaceTenantKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorNamespaceRequestAllOf) SetKind(v ConnectorNamespaceTenantKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectorNamespaceRequestAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


