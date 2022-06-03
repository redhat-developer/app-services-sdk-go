# ConnectorNamespaceRequestMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Namespace name must match pattern &#x60;^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$&#x60;, or it may be empty to be auto-generated. | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConnectorNamespaceRequestMeta

`func NewConnectorNamespaceRequestMeta() *ConnectorNamespaceRequestMeta`

NewConnectorNamespaceRequestMeta instantiates a new ConnectorNamespaceRequestMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceRequestMetaWithDefaults

`func NewConnectorNamespaceRequestMetaWithDefaults() *ConnectorNamespaceRequestMeta`

NewConnectorNamespaceRequestMetaWithDefaults instantiates a new ConnectorNamespaceRequestMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorNamespaceRequestMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespaceRequestMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespaceRequestMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorNamespaceRequestMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConnectorNamespaceRequestMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorNamespaceRequestMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorNamespaceRequestMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorNamespaceRequestMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


