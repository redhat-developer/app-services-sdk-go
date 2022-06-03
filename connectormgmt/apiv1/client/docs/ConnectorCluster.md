# ConnectorCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ConnectorClusterStatusStatus**](ConnectorClusterStatusStatus.md) |  | [optional] 

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

### GetOwner

`func (o *ConnectorCluster) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConnectorCluster) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConnectorCluster) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConnectorCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConnectorCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConnectorCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ConnectorCluster) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorCluster) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorCluster) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ConnectorCluster) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *ConnectorCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectorCluster) GetStatus() ConnectorClusterStatusStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorCluster) GetStatusOk() (*ConnectorClusterStatusStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorCluster) SetStatus(v ConnectorClusterStatusStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


