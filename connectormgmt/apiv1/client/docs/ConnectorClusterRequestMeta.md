# ConnectorClusterRequestMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Name-value string annotations for resource | [optional] 

## Methods

### NewConnectorClusterRequestMeta

`func NewConnectorClusterRequestMeta() *ConnectorClusterRequestMeta`

NewConnectorClusterRequestMeta instantiates a new ConnectorClusterRequestMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorClusterRequestMetaWithDefaults

`func NewConnectorClusterRequestMetaWithDefaults() *ConnectorClusterRequestMeta`

NewConnectorClusterRequestMetaWithDefaults instantiates a new ConnectorClusterRequestMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorClusterRequestMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorClusterRequestMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorClusterRequestMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorClusterRequestMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConnectorClusterRequestMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorClusterRequestMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorClusterRequestMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorClusterRequestMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


