# ConnectorNamespaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**ConnectorNamespaceState**](ConnectorNamespaceState.md) |  | 
**Version** | Pointer to **string** |  | [optional] 
**ConnectorsDeployed** | **int32** |  | 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectorNamespaceStatus

`func NewConnectorNamespaceStatus(state ConnectorNamespaceState, connectorsDeployed int32, ) *ConnectorNamespaceStatus`

NewConnectorNamespaceStatus instantiates a new ConnectorNamespaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceStatusWithDefaults

`func NewConnectorNamespaceStatusWithDefaults() *ConnectorNamespaceStatus`

NewConnectorNamespaceStatusWithDefaults instantiates a new ConnectorNamespaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ConnectorNamespaceStatus) GetState() ConnectorNamespaceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectorNamespaceStatus) GetStateOk() (*ConnectorNamespaceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectorNamespaceStatus) SetState(v ConnectorNamespaceState)`

SetState sets State field to given value.


### GetVersion

`func (o *ConnectorNamespaceStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorNamespaceStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorNamespaceStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorNamespaceStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConnectorsDeployed

`func (o *ConnectorNamespaceStatus) GetConnectorsDeployed() int32`

GetConnectorsDeployed returns the ConnectorsDeployed field if non-nil, zero value otherwise.

### GetConnectorsDeployedOk

`func (o *ConnectorNamespaceStatus) GetConnectorsDeployedOk() (*int32, bool)`

GetConnectorsDeployedOk returns a tuple with the ConnectorsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorsDeployed

`func (o *ConnectorNamespaceStatus) SetConnectorsDeployed(v int32)`

SetConnectorsDeployed sets ConnectorsDeployed field to given value.


### GetError

`func (o *ConnectorNamespaceStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectorNamespaceStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectorNamespaceStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectorNamespaceStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


