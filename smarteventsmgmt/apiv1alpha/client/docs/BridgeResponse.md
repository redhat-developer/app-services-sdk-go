# BridgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**ManagedResourceStatus**](ManagedResourceStatus.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**ErrorHandler** | Pointer to [**Action**](Action.md) |  | [optional] 

## Methods

### NewBridgeResponse

`func NewBridgeResponse() *BridgeResponse`

NewBridgeResponse instantiates a new BridgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeResponseWithDefaults

`func NewBridgeResponseWithDefaults() *BridgeResponse`

NewBridgeResponseWithDefaults instantiates a new BridgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BridgeResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BridgeResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BridgeResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BridgeResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *BridgeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BridgeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BridgeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BridgeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BridgeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BridgeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BridgeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BridgeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHref

`func (o *BridgeResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BridgeResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BridgeResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BridgeResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *BridgeResponse) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *BridgeResponse) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *BridgeResponse) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *BridgeResponse) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *BridgeResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *BridgeResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *BridgeResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *BridgeResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStatus

`func (o *BridgeResponse) GetStatus() ManagedResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BridgeResponse) GetStatusOk() (*ManagedResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BridgeResponse) SetStatus(v ManagedResourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BridgeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOwner

`func (o *BridgeResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BridgeResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BridgeResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BridgeResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEndpoint

`func (o *BridgeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BridgeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BridgeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *BridgeResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetErrorHandler

`func (o *BridgeResponse) GetErrorHandler() Action`

GetErrorHandler returns the ErrorHandler field if non-nil, zero value otherwise.

### GetErrorHandlerOk

`func (o *BridgeResponse) GetErrorHandlerOk() (*Action, bool)`

GetErrorHandlerOk returns a tuple with the ErrorHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHandler

`func (o *BridgeResponse) SetErrorHandler(v Action)`

SetErrorHandler sets ErrorHandler field to given value.

### HasErrorHandler

`func (o *BridgeResponse) HasErrorHandler() bool`

HasErrorHandler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


