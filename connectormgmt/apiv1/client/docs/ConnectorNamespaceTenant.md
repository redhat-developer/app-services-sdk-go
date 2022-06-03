# ConnectorNamespaceTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ConnectorNamespaceTenantKind**](ConnectorNamespaceTenantKind.md) |  | 
**Id** | **string** | Either user or organisation id depending on the value of kind | 

## Methods

### NewConnectorNamespaceTenant

`func NewConnectorNamespaceTenant(kind ConnectorNamespaceTenantKind, id string, ) *ConnectorNamespaceTenant`

NewConnectorNamespaceTenant instantiates a new ConnectorNamespaceTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceTenantWithDefaults

`func NewConnectorNamespaceTenantWithDefaults() *ConnectorNamespaceTenant`

NewConnectorNamespaceTenantWithDefaults instantiates a new ConnectorNamespaceTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConnectorNamespaceTenant) GetKind() ConnectorNamespaceTenantKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorNamespaceTenant) GetKindOk() (*ConnectorNamespaceTenantKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorNamespaceTenant) SetKind(v ConnectorNamespaceTenantKind)`

SetKind sets Kind field to given value.


### GetId

`func (o *ConnectorNamespaceTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorNamespaceTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorNamespaceTenant) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


