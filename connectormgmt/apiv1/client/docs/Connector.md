# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**ConnectorTypeId** | **string** |  | 
**NamespaceId** | **string** |  | 
**Channel** | Pointer to [**Channel**](Channel.md) |  | [optional] [default to CHANNEL_STABLE]
**DesiredState** | [**ConnectorDesiredState**](ConnectorDesiredState.md) |  | 
**Annotations** | Pointer to **map[string]string** | Name-value string annotations for resource | [optional] 
**ResourceVersion** | Pointer to **int64** |  | [optional] 
**Kafka** | [**KafkaConnectionSettings**](KafkaConnectionSettings.md) |  | 
**ServiceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 
**SchemaRegistry** | Pointer to [**SchemaRegistryConnectionSettings**](SchemaRegistryConnectionSettings.md) |  | [optional] 
**Connector** | **map[string]interface{}** |  | 
**Status** | Pointer to [**ConnectorStatusStatus**](ConnectorStatusStatus.md) |  | [optional] 

## Methods

### NewConnector

`func NewConnector(name string, connectorTypeId string, namespaceId string, desiredState ConnectorDesiredState, kafka KafkaConnectionSettings, serviceAccount ServiceAccount, connector map[string]interface{}, ) *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Connector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Connector) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Connector) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Connector) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Connector) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *Connector) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Connector) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Connector) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Connector) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetOwner

`func (o *Connector) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Connector) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Connector) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Connector) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Connector) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Connector) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Connector) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Connector) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Connector) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Connector) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Connector) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Connector) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Connector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connector) SetName(v string)`

SetName sets Name field to given value.


### GetConnectorTypeId

`func (o *Connector) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *Connector) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *Connector) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.


### GetNamespaceId

`func (o *Connector) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *Connector) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *Connector) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.


### GetChannel

`func (o *Connector) GetChannel() Channel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Connector) GetChannelOk() (*Channel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Connector) SetChannel(v Channel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Connector) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDesiredState

`func (o *Connector) GetDesiredState() ConnectorDesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *Connector) GetDesiredStateOk() (*ConnectorDesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *Connector) SetDesiredState(v ConnectorDesiredState)`

SetDesiredState sets DesiredState field to given value.


### GetAnnotations

`func (o *Connector) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Connector) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Connector) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Connector) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetResourceVersion

`func (o *Connector) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *Connector) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *Connector) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *Connector) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetKafka

`func (o *Connector) GetKafka() KafkaConnectionSettings`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *Connector) GetKafkaOk() (*KafkaConnectionSettings, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *Connector) SetKafka(v KafkaConnectionSettings)`

SetKafka sets Kafka field to given value.


### GetServiceAccount

`func (o *Connector) GetServiceAccount() ServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *Connector) GetServiceAccountOk() (*ServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *Connector) SetServiceAccount(v ServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetSchemaRegistry

`func (o *Connector) GetSchemaRegistry() SchemaRegistryConnectionSettings`

GetSchemaRegistry returns the SchemaRegistry field if non-nil, zero value otherwise.

### GetSchemaRegistryOk

`func (o *Connector) GetSchemaRegistryOk() (*SchemaRegistryConnectionSettings, bool)`

GetSchemaRegistryOk returns a tuple with the SchemaRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistry

`func (o *Connector) SetSchemaRegistry(v SchemaRegistryConnectionSettings)`

SetSchemaRegistry sets SchemaRegistry field to given value.

### HasSchemaRegistry

`func (o *Connector) HasSchemaRegistry() bool`

HasSchemaRegistry returns a boolean if a field has been set.

### GetConnector

`func (o *Connector) GetConnector() map[string]interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *Connector) GetConnectorOk() (*map[string]interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *Connector) SetConnector(v map[string]interface{})`

SetConnector sets Connector field to given value.


### GetStatus

`func (o *Connector) GetStatus() ConnectorStatusStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Connector) GetStatusOk() (*ConnectorStatusStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Connector) SetStatus(v ConnectorStatusStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Connector) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


