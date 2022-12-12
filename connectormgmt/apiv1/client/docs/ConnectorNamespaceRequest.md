# ConnectorNamespaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Namespace name must match pattern &#x60;^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$&#x60;, or it may be empty to be auto-generated. | 
**Annotations** | Pointer to **map[string]string** | Name-value string annotations for resource | [optional] 
**ClusterId** | **string** |  | 
**Kind** | [**ConnectorNamespaceTenantKind**](ConnectorNamespaceTenantKind.md) |  | 

## Methods

### NewConnectorNamespaceRequest

`func NewConnectorNamespaceRequest(name string, clusterId string, kind ConnectorNamespaceTenantKind, ) *ConnectorNamespaceRequest`

NewConnectorNamespaceRequest instantiates a new ConnectorNamespaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceRequestWithDefaults

`func NewConnectorNamespaceRequestWithDefaults() *ConnectorNamespaceRequest`

NewConnectorNamespaceRequestWithDefaults instantiates a new ConnectorNamespaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorNamespaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorNamespaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorNamespaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *ConnectorNamespaceRequest) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConnectorNamespaceRequest) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConnectorNamespaceRequest) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConnectorNamespaceRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetClusterId

`func (o *ConnectorNamespaceRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConnectorNamespaceRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConnectorNamespaceRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetKind

`func (o *ConnectorNamespaceRequest) GetKind() ConnectorNamespaceTenantKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorNamespaceRequest) GetKindOk() (*ConnectorNamespaceTenantKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorNamespaceRequest) SetKind(v ConnectorNamespaceTenantKind)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


