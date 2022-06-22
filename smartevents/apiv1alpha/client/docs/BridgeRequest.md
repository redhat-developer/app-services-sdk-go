# BridgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ErrorHandler** | Pointer to [**Action**](Action.md) |  | [optional] 

## Methods

### NewBridgeRequest

`func NewBridgeRequest(name string, ) *BridgeRequest`

NewBridgeRequest instantiates a new BridgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeRequestWithDefaults

`func NewBridgeRequestWithDefaults() *BridgeRequest`

NewBridgeRequestWithDefaults instantiates a new BridgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BridgeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BridgeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BridgeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetErrorHandler

`func (o *BridgeRequest) GetErrorHandler() Action`

GetErrorHandler returns the ErrorHandler field if non-nil, zero value otherwise.

### GetErrorHandlerOk

`func (o *BridgeRequest) GetErrorHandlerOk() (*Action, bool)`

GetErrorHandlerOk returns a tuple with the ErrorHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHandler

`func (o *BridgeRequest) SetErrorHandler(v Action)`

SetErrorHandler sets ErrorHandler field to given value.

### HasErrorHandler

`func (o *BridgeRequest) HasErrorHandler() bool`

HasErrorHandler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


