# ProcessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Filters** | Pointer to [**[]BaseFilter**](BaseFilter.md) |  | [optional] 
**TransformationTemplate** | Pointer to **string** |  | [optional] 
**Action** | Pointer to [**Action**](Action.md) |  | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 

## Methods

### NewProcessorRequest

`func NewProcessorRequest(name string, ) *ProcessorRequest`

NewProcessorRequest instantiates a new ProcessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorRequestWithDefaults

`func NewProcessorRequestWithDefaults() *ProcessorRequest`

NewProcessorRequestWithDefaults instantiates a new ProcessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProcessorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFilters

`func (o *ProcessorRequest) GetFilters() []BaseFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ProcessorRequest) GetFiltersOk() (*[]BaseFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ProcessorRequest) SetFilters(v []BaseFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ProcessorRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTransformationTemplate

`func (o *ProcessorRequest) GetTransformationTemplate() string`

GetTransformationTemplate returns the TransformationTemplate field if non-nil, zero value otherwise.

### GetTransformationTemplateOk

`func (o *ProcessorRequest) GetTransformationTemplateOk() (*string, bool)`

GetTransformationTemplateOk returns a tuple with the TransformationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationTemplate

`func (o *ProcessorRequest) SetTransformationTemplate(v string)`

SetTransformationTemplate sets TransformationTemplate field to given value.

### HasTransformationTemplate

`func (o *ProcessorRequest) HasTransformationTemplate() bool`

HasTransformationTemplate returns a boolean if a field has been set.

### GetAction

`func (o *ProcessorRequest) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProcessorRequest) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProcessorRequest) SetAction(v Action)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProcessorRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSource

`func (o *ProcessorRequest) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProcessorRequest) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProcessorRequest) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProcessorRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


