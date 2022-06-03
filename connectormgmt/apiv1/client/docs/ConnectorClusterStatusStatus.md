# ConnectorClusterStatusStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**ConnectorClusterState**](ConnectorClusterState.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectorClusterStatusStatus

`func NewConnectorClusterStatusStatus() *ConnectorClusterStatusStatus`

NewConnectorClusterStatusStatus instantiates a new ConnectorClusterStatusStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorClusterStatusStatusWithDefaults

`func NewConnectorClusterStatusStatusWithDefaults() *ConnectorClusterStatusStatus`

NewConnectorClusterStatusStatusWithDefaults instantiates a new ConnectorClusterStatusStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ConnectorClusterStatusStatus) GetState() ConnectorClusterState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectorClusterStatusStatus) GetStateOk() (*ConnectorClusterState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectorClusterStatusStatus) SetState(v ConnectorClusterState)`

SetState sets State field to given value.

### HasState

`func (o *ConnectorClusterStatusStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetError

`func (o *ConnectorClusterStatusStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectorClusterStatusStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectorClusterStatusStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectorClusterStatusStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


