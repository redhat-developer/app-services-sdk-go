# ConnectorClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Name-value string annotations for resource | [optional] 

## Methods

### NewConnectorClusterRequest

`func NewConnectorClusterRequest() *ConnectorClusterRequest`

NewConnectorClusterRequest instantiates a new ConnectorClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorClusterRequestWithDefaults

`func NewConnectorClusterRequestWithDefaults() *ConnectorClusterRequest`

NewConnectorClusterRequestWithDefaults instantiates a new ConnectorClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConnectorClusterRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorClusterRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorClusterRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorClusterRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


