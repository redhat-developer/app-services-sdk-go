# ConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConnectorTypeId** | **string** |  | 
**NamespaceId** | **string** |  | 
**Channel** | Pointer to [**Channel**](Channel.md) |  | [optional] [default to CHANNEL_STABLE]
**DesiredState** | [**ConnectorDesiredState**](ConnectorDesiredState.md) |  | 
**Kafka** | [**KafkaConnectionSettings**](KafkaConnectionSettings.md) |  | 
**ServiceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 
**SchemaRegistry** | Pointer to [**SchemaRegistryConnectionSettings**](SchemaRegistryConnectionSettings.md) |  | [optional] 
**Connector** | **map[string]interface{}** |  | 

## Methods

### NewConnectorRequest

`func NewConnectorRequest(name string, connectorTypeId string, namespaceId string, desiredState ConnectorDesiredState, kafka KafkaConnectionSettings, serviceAccount ServiceAccount, connector map[string]interface{}, ) *ConnectorRequest`

NewConnectorRequest instantiates a new ConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorRequestWithDefaults

`func NewConnectorRequestWithDefaults() *ConnectorRequest`

NewConnectorRequestWithDefaults instantiates a new ConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConnectorTypeId

`func (o *ConnectorRequest) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *ConnectorRequest) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *ConnectorRequest) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.


### GetNamespaceId

`func (o *ConnectorRequest) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *ConnectorRequest) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *ConnectorRequest) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.


### GetChannel

`func (o *ConnectorRequest) GetChannel() Channel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConnectorRequest) GetChannelOk() (*Channel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConnectorRequest) SetChannel(v Channel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ConnectorRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDesiredState

`func (o *ConnectorRequest) GetDesiredState() ConnectorDesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *ConnectorRequest) GetDesiredStateOk() (*ConnectorDesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *ConnectorRequest) SetDesiredState(v ConnectorDesiredState)`

SetDesiredState sets DesiredState field to given value.


### GetKafka

`func (o *ConnectorRequest) GetKafka() KafkaConnectionSettings`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *ConnectorRequest) GetKafkaOk() (*KafkaConnectionSettings, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *ConnectorRequest) SetKafka(v KafkaConnectionSettings)`

SetKafka sets Kafka field to given value.


### GetServiceAccount

`func (o *ConnectorRequest) GetServiceAccount() ServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ConnectorRequest) GetServiceAccountOk() (*ServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ConnectorRequest) SetServiceAccount(v ServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetSchemaRegistry

`func (o *ConnectorRequest) GetSchemaRegistry() SchemaRegistryConnectionSettings`

GetSchemaRegistry returns the SchemaRegistry field if non-nil, zero value otherwise.

### GetSchemaRegistryOk

`func (o *ConnectorRequest) GetSchemaRegistryOk() (*SchemaRegistryConnectionSettings, bool)`

GetSchemaRegistryOk returns a tuple with the SchemaRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistry

`func (o *ConnectorRequest) SetSchemaRegistry(v SchemaRegistryConnectionSettings)`

SetSchemaRegistry sets SchemaRegistry field to given value.

### HasSchemaRegistry

`func (o *ConnectorRequest) HasSchemaRegistry() bool`

HasSchemaRegistry returns a boolean if a field has been set.

### GetConnector

`func (o *ConnectorRequest) GetConnector() map[string]interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectorRequest) GetConnectorOk() (*map[string]interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectorRequest) SetConnector(v map[string]interface{})`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


