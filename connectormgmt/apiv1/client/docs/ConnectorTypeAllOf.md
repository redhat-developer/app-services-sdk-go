# ConnectorTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector type. | [optional] 
**Version** | Pointer to **string** | Version of the connector type. | [optional] 
**Channels** | Pointer to **[]string** | Channels of the connector type. | [optional] 
**Description** | Pointer to **string** | A description of the connector. | [optional] 
**IconHref** | Pointer to **string** | URL to an icon of the connector. | [optional] 
**Labels** | Pointer to **[]string** | Labels used to categorize the connector | [optional] 
**Schema** | Pointer to **map[string]interface{}** | A json schema that can be used to validate a connectors connector_spec field. | [optional] 
**JsonSchema** | Pointer to **map[string]interface{}** | A json schema that can be used to validate a connectors connector_spec field. | [optional] 


## Methods

### NewConnectorTypeAllOf

`func NewConnectorTypeAllOf() *ConnectorTypeAllOf`

NewConnectorTypeAllOf instantiates a new ConnectorTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTypeAllOfWithDefaults

`func NewConnectorTypeAllOfWithDefaults() *ConnectorTypeAllOf`

NewConnectorTypeAllOfWithDefaults instantiates a new ConnectorTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *ConnectorTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


### GetVersion

`func (o *ConnectorTypeAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorTypeAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorTypeAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorTypeAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


### GetChannels

`func (o *ConnectorTypeAllOf) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ConnectorTypeAllOf) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ConnectorTypeAllOf) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ConnectorTypeAllOf) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


### GetDescription

`func (o *ConnectorTypeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorTypeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorTypeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorTypeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


### GetIconHref

`func (o *ConnectorTypeAllOf) GetIconHref() string`

GetIconHref returns the IconHref field if non-nil, zero value otherwise.

### GetIconHrefOk

`func (o *ConnectorTypeAllOf) GetIconHrefOk() (*string, bool)`

GetIconHrefOk returns a tuple with the IconHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconHref

`func (o *ConnectorTypeAllOf) SetIconHref(v string)`

SetIconHref sets IconHref field to given value.

### HasIconHref

`func (o *ConnectorTypeAllOf) HasIconHref() bool`

HasIconHref returns a boolean if a field has been set.


### GetLabels

`func (o *ConnectorTypeAllOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConnectorTypeAllOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConnectorTypeAllOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ConnectorTypeAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


### GetSchema

`func (o *ConnectorTypeAllOf) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ConnectorTypeAllOf) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ConnectorTypeAllOf) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ConnectorTypeAllOf) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


### GetJsonSchema

`func (o *ConnectorTypeAllOf) GetJsonSchema() map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *ConnectorTypeAllOf) GetJsonSchemaOk() (*map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *ConnectorTypeAllOf) SetJsonSchema(v map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *ConnectorTypeAllOf) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

