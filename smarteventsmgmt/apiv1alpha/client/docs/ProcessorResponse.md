# ProcessorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind (type) of this resource | 
**Id** | **string** | The unique identifier of this resource | 
**Href** | **string** | The URL of this resource, without the protocol | 
**SubmittedAt** | **time.Time** |  | 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | [**ManagedResourceStatus**](ManagedResourceStatus.md) |  | 
**Owner** | **string** | The user that owns this resource | 
**Name** | **string** | The name of the processor | 
**Flows** | **map[string]interface{}** | The Camel YAML DSL code, formatted as JSON, that defines the flows in the processor | 
**StatusMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewProcessorResponse

`func NewProcessorResponse(kind string, id string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, name string, flows map[string]interface{}, ) *ProcessorResponse`

NewProcessorResponse instantiates a new ProcessorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorResponseWithDefaults

`func NewProcessorResponseWithDefaults() *ProcessorResponse`

NewProcessorResponseWithDefaults instantiates a new ProcessorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ProcessorResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProcessorResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProcessorResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *ProcessorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *ProcessorResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProcessorResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProcessorResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetSubmittedAt

`func (o *ProcessorResponse) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *ProcessorResponse) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *ProcessorResponse) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.


### GetPublishedAt

`func (o *ProcessorResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ProcessorResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ProcessorResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *ProcessorResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ProcessorResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ProcessorResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ProcessorResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ProcessorResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessorResponse) GetStatus() ManagedResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessorResponse) GetStatusOk() (*ManagedResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessorResponse) SetStatus(v ManagedResourceStatus)`

SetStatus sets Status field to given value.


### GetOwner

`func (o *ProcessorResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ProcessorResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ProcessorResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetName

`func (o *ProcessorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetFlows

`func (o *ProcessorResponse) GetFlows() map[string]interface{}`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ProcessorResponse) GetFlowsOk() (*map[string]interface{}, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ProcessorResponse) SetFlows(v map[string]interface{})`

SetFlows sets Flows field to given value.


### GetStatusMessage

`func (o *ProcessorResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ProcessorResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ProcessorResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ProcessorResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


