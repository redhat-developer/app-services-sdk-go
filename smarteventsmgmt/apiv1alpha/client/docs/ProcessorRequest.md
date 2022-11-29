# ProcessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Flows** | **map[string]interface{}** |  | 

## Methods

### NewProcessorRequest

`func NewProcessorRequest(name string, flows map[string]interface{}, ) *ProcessorRequest`

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


### GetFlows

`func (o *ProcessorRequest) GetFlows() map[string]interface{}`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ProcessorRequest) GetFlowsOk() (*map[string]interface{}, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ProcessorRequest) SetFlows(v map[string]interface{})`

SetFlows sets Flows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


