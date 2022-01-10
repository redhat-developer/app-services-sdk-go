# ConnectorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the connector type. | [optional] 
**Version** | Pointer to **string** | Version of the connector type. | [optional] 
**Channels** | Pointer to **[]string** | Channels of the connector type. | [optional] 
**Description** | Pointer to **string** | A description of the connector. | [optional] 
**IconHref** | Pointer to **string** | URL to an icon of the connector. | [optional] 
**Labels** | Pointer to **[]string** | Labels used to categorize the connector | [optional] 
**Schema** | Pointer to **map[string]interface{}** | A json schema that can be used to validate a connectors connector_spec field. | [optional] 
**JsonSchema** | Pointer to **map[string]interface{}** | A json schema that can be used to validate a connectors connector_spec field. | [optional] 


## Methods

### NewConnectorType

`func NewConnectorType() *ConnectorType`

NewConnectorType instantiates a new ConnectorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTypeWithDefaults

`func NewConnectorTypeWithDefaults() *ConnectorType`

NewConnectorTypeWithDefaults instantiates a new ConnectorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetId

`func (o *ConnectorType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorType) HasId() bool`

HasId returns a boolean if a field has been set.


### GetKind

`func (o *ConnectorType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorType) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectorType) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetHref

`func (o *ConnectorType) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectorType) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectorType) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectorType) HasHref() bool`

HasHref returns a boolean if a field has been set.


### GetName

`func (o *ConnectorType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorType) HasName() bool`

HasName returns a boolean if a field has been set.


### GetVersion

`func (o *ConnectorType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorType) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


### GetChannels

`func (o *ConnectorType) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ConnectorType) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ConnectorType) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ConnectorType) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


### GetDescription

`func (o *ConnectorType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


### GetIconHref

`func (o *ConnectorType) GetIconHref() string`

GetIconHref returns the IconHref field if non-nil, zero value otherwise.

### GetIconHrefOk

`func (o *ConnectorType) GetIconHrefOk() (*string, bool)`

GetIconHrefOk returns a tuple with the IconHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconHref

`func (o *ConnectorType) SetIconHref(v string)`

SetIconHref sets IconHref field to given value.

### HasIconHref

`func (o *ConnectorType) HasIconHref() bool`

HasIconHref returns a boolean if a field has been set.


### GetLabels

`func (o *ConnectorType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConnectorType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConnectorType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ConnectorType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


### GetSchema

`func (o *ConnectorType) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ConnectorType) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ConnectorType) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ConnectorType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


### GetJsonSchema

`func (o *ConnectorType) GetJsonSchema() map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *ConnectorType) GetJsonSchemaOk() (*map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *ConnectorType) SetJsonSchema(v map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *ConnectorType) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

