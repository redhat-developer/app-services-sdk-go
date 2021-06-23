# ConnectorCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ConnectorClusterAllOfMetadata**](ConnectorClusterAllOfMetadata.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 


## Methods

### NewConnectorCluster

`func NewConnectorCluster() *ConnectorCluster`

NewConnectorCluster instantiates a new ConnectorCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorClusterWithDefaults

`func NewConnectorClusterWithDefaults() *ConnectorCluster`

NewConnectorClusterWithDefaults instantiates a new ConnectorCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetId

`func (o *ConnectorCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorCluster) HasId() bool`

HasId returns a boolean if a field has been set.


### GetKind

`func (o *ConnectorCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnectorCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetHref

`func (o *ConnectorCluster) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectorCluster) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectorCluster) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectorCluster) HasHref() bool`

HasHref returns a boolean if a field has been set.


### GetMetadata

`func (o *ConnectorCluster) GetMetadata() ConnectorClusterAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectorCluster) GetMetadataOk() (*ConnectorClusterAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectorCluster) SetMetadata(v ConnectorClusterAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectorCluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


### GetStatus

`func (o *ConnectorCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

