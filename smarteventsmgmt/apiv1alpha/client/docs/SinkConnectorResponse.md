# SinkConnectorResponse

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
**UriDsl** | **string** | The URI to be used in Camel DSL to send data to this sink | 

## Methods

### NewSinkConnectorResponse

`func NewSinkConnectorResponse(kind string, id string, name string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, connectorTypeId string, connector map[string]interface{}, uriDsl string, ) *SinkConnectorResponse`

NewSinkConnectorResponse instantiates a new SinkConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinkConnectorResponseWithDefaults

`func NewSinkConnectorResponseWithDefaults() *SinkConnectorResponse`

NewSinkConnectorResponseWithDefaults instantiates a new SinkConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SinkConnectorResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SinkConnectorResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SinkConnectorResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *SinkConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SinkConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SinkConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SinkConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SinkConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SinkConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHref

`func (o *SinkConnectorResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SinkConnectorResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SinkConnectorResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetSubmittedAt

`func (o *SinkConnectorResponse) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *SinkConnectorResponse) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *SinkConnectorResponse) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.


### GetPublishedAt

`func (o *SinkConnectorResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *SinkConnectorResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *SinkConnectorResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *SinkConnectorResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SinkConnectorResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SinkConnectorResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SinkConnectorResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SinkConnectorResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SinkConnectorResponse) GetStatus() ManagedResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SinkConnectorResponse) GetStatusOk() (*ManagedResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SinkConnectorResponse) SetStatus(v ManagedResourceStatus)`

SetStatus sets Status field to given value.


### GetOwner

`func (o *SinkConnectorResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SinkConnectorResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SinkConnectorResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetConnectorTypeId

`func (o *SinkConnectorResponse) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *SinkConnectorResponse) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *SinkConnectorResponse) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.


### GetConnector

`func (o *SinkConnectorResponse) GetConnector() map[string]interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *SinkConnectorResponse) GetConnectorOk() (*map[string]interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *SinkConnectorResponse) SetConnector(v map[string]interface{})`

SetConnector sets Connector field to given value.


### GetStatusMessage

`func (o *SinkConnectorResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SinkConnectorResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SinkConnectorResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SinkConnectorResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUriDsl

`func (o *SinkConnectorResponse) GetUriDsl() string`

GetUriDsl returns the UriDsl field if non-nil, zero value otherwise.

### GetUriDslOk

`func (o *SinkConnectorResponse) GetUriDslOk() (*string, bool)`

GetUriDslOk returns a tuple with the UriDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriDsl

`func (o *SinkConnectorResponse) SetUriDsl(v string)`

SetUriDsl sets UriDsl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


