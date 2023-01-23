# SourceConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind (type) of this resource | 
**Id** | **string** | The unique identifier of this resource | 
**Name** | **string** | The name of this resource | 
**Href** | **string** | The URL of this resource, without the protocol | 
**SubmittedAt** | **time.Time** |  | 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | [**ManagedResourceStatus**](ManagedResourceStatus.md) |  | 
**Owner** | **string** | The user that owns this resource | 
**ConnectorTypeId** | **string** | The connector type | 
**Connector** | **map[string]interface{}** | The Connector configuration payload | 
**StatusMessage** | Pointer to **string** | A detailed status message in case there is a problem with the connector | [optional] 

## Methods

### NewSourceConnectorResponse

`func NewSourceConnectorResponse(kind string, id string, name string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, connectorTypeId string, connector map[string]interface{}, ) *SourceConnectorResponse`

NewSourceConnectorResponse instantiates a new SourceConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceConnectorResponseWithDefaults

`func NewSourceConnectorResponseWithDefaults() *SourceConnectorResponse`

NewSourceConnectorResponseWithDefaults instantiates a new SourceConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SourceConnectorResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SourceConnectorResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SourceConnectorResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *SourceConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SourceConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHref

`func (o *SourceConnectorResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SourceConnectorResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SourceConnectorResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetSubmittedAt

`func (o *SourceConnectorResponse) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *SourceConnectorResponse) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *SourceConnectorResponse) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.


### GetPublishedAt

`func (o *SourceConnectorResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *SourceConnectorResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *SourceConnectorResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *SourceConnectorResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SourceConnectorResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SourceConnectorResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SourceConnectorResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SourceConnectorResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SourceConnectorResponse) GetStatus() ManagedResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceConnectorResponse) GetStatusOk() (*ManagedResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceConnectorResponse) SetStatus(v ManagedResourceStatus)`

SetStatus sets Status field to given value.


### GetOwner

`func (o *SourceConnectorResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SourceConnectorResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SourceConnectorResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetConnectorTypeId

`func (o *SourceConnectorResponse) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *SourceConnectorResponse) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *SourceConnectorResponse) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.


### GetConnector

`func (o *SourceConnectorResponse) GetConnector() map[string]interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SourceConnectorResponse) GetConnectorOk() (*map[string]interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SourceConnectorResponse) SetConnector(v map[string]interface{})`

SetConnector sets Connector field to given value.


### GetStatusMessage

`func (o *SourceConnectorResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SourceConnectorResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SourceConnectorResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SourceConnectorResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


