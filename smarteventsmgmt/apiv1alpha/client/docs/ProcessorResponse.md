# ProcessorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Href** | **string** |  | 
**SubmittedAt** | **time.Time** |  | 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | [**ManagedResourceStatus**](ManagedResourceStatus.md) |  | 
**Owner** | **string** |  | 
**Type** | [**ProcessorType**](ProcessorType.md) |  | 
**Filters** | Pointer to [**[]BaseFilter**](BaseFilter.md) |  | [optional] 
**TransformationTemplate** | Pointer to **string** |  | [optional] 
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewProcessorResponse

`func NewProcessorResponse(kind string, id string, href string, submittedAt time.Time, status ManagedResourceStatus, owner string, type_ ProcessorType, ) *ProcessorResponse`

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

### HasName

`func (o *ProcessorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

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


### GetType

`func (o *ProcessorResponse) GetType() ProcessorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorResponse) GetTypeOk() (*ProcessorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorResponse) SetType(v ProcessorType)`

SetType sets Type field to given value.


### GetFilters

`func (o *ProcessorResponse) GetFilters() []BaseFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProcessorResponse) GetFiltersOk() (*[]BaseFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProcessorResponse) SetFilters(v []BaseFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ProcessorResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTransformationTemplate

`func (o *ProcessorResponse) GetTransformationTemplate() string`

GetTransformationTemplate returns the TransformationTemplate field if non-nil, zero value otherwise.

### GetTransformationTemplateOk

`func (o *ProcessorResponse) GetTransformationTemplateOk() (*string, bool)`

GetTransformationTemplateOk returns a tuple with the TransformationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationTemplate

`func (o *ProcessorResponse) SetTransformationTemplate(v string)`

SetTransformationTemplate sets TransformationTemplate field to given value.

### HasTransformationTemplate

`func (o *ProcessorResponse) HasTransformationTemplate() bool`

HasTransformationTemplate returns a boolean if a field has been set.

### GetAction

`func (o *ProcessorResponse) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProcessorResponse) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProcessorResponse) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProcessorResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSource

`func (o *ProcessorResponse) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProcessorResponse) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProcessorResponse) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProcessorResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

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


