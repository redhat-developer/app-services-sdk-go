# ConnectorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kafka** | [**KafkaConnectionSettings**](KafkaConnectionSettings.md) |  | 
**ServiceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 
**SchemaRegistry** | Pointer to [**SchemaRegistryConnectionSettings**](SchemaRegistryConnectionSettings.md) |  | [optional] 
**Connector** | **map[string]interface{}** |  | 

## Methods

### NewConnectorConfiguration

`func NewConnectorConfiguration(kafka KafkaConnectionSettings, serviceAccount ServiceAccount, connector map[string]interface{}, ) *ConnectorConfiguration`

NewConnectorConfiguration instantiates a new ConnectorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConfigurationWithDefaults

`func NewConnectorConfigurationWithDefaults() *ConnectorConfiguration`

NewConnectorConfigurationWithDefaults instantiates a new ConnectorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafka

`func (o *ConnectorConfiguration) GetKafka() KafkaConnectionSettings`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *ConnectorConfiguration) GetKafkaOk() (*KafkaConnectionSettings, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *ConnectorConfiguration) SetKafka(v KafkaConnectionSettings)`

SetKafka sets Kafka field to given value.


### GetServiceAccount

`func (o *ConnectorConfiguration) GetServiceAccount() ServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ConnectorConfiguration) GetServiceAccountOk() (*ServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ConnectorConfiguration) SetServiceAccount(v ServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetSchemaRegistry

`func (o *ConnectorConfiguration) GetSchemaRegistry() SchemaRegistryConnectionSettings`

GetSchemaRegistry returns the SchemaRegistry field if non-nil, zero value otherwise.

### GetSchemaRegistryOk

`func (o *ConnectorConfiguration) GetSchemaRegistryOk() (*SchemaRegistryConnectionSettings, bool)`

GetSchemaRegistryOk returns a tuple with the SchemaRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistry

`func (o *ConnectorConfiguration) SetSchemaRegistry(v SchemaRegistryConnectionSettings)`

SetSchemaRegistry sets SchemaRegistry field to given value.

### HasSchemaRegistry

`func (o *ConnectorConfiguration) HasSchemaRegistry() bool`

HasSchemaRegistry returns a boolean if a field has been set.

### GetConnector

`func (o *ConnectorConfiguration) GetConnector() map[string]interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectorConfiguration) GetConnectorOk() (*map[string]interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectorConfiguration) SetConnector(v map[string]interface{})`

SetConnector sets Connector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


