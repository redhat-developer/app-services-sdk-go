# ConnectorRequestMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConnectorTypeId** | **string** |  | 
**Channel** | Pointer to [**Channel**](Channel.md) |  | [optional] [default to CHANNEL_STABLE]
**DeploymentLocation** | [**DeploymentLocation**](DeploymentLocation.md) |  | 
**DesiredState** | [**ConnectorDesiredState**](ConnectorDesiredState.md) |  | 


## Methods

### NewConnectorRequestMeta

`func NewConnectorRequestMeta(name string, connectorTypeId string, deploymentLocation DeploymentLocation, desiredState ConnectorDesiredState, ) *ConnectorRequestMeta`

NewConnectorRequestMeta instantiates a new ConnectorRequestMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRequestMetaWithDefaults

`func NewConnectorRequestMetaWithDefaults() *ConnectorRequestMeta`

NewConnectorRequestMetaWithDefaults instantiates a new ConnectorRequestMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *ConnectorRequestMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorRequestMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorRequestMeta) SetName(v string)`

SetName sets Name field to given value.



### GetConnectorTypeId

`func (o *ConnectorRequestMeta) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *ConnectorRequestMeta) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *ConnectorRequestMeta) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.



### GetChannel

`func (o *ConnectorRequestMeta) GetChannel() Channel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConnectorRequestMeta) GetChannelOk() (*Channel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConnectorRequestMeta) SetChannel(v Channel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ConnectorRequestMeta) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


### GetDeploymentLocation

`func (o *ConnectorRequestMeta) GetDeploymentLocation() DeploymentLocation`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *ConnectorRequestMeta) GetDeploymentLocationOk() (*DeploymentLocation, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *ConnectorRequestMeta) SetDeploymentLocation(v DeploymentLocation)`

SetDeploymentLocation sets DeploymentLocation field to given value.



### GetDesiredState

`func (o *ConnectorRequestMeta) GetDesiredState() ConnectorDesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *ConnectorRequestMeta) GetDesiredStateOk() (*ConnectorDesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *ConnectorRequestMeta) SetDesiredState(v ConnectorDesiredState)`

SetDesiredState sets DesiredState field to given value.




[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

