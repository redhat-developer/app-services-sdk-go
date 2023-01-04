# BridgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the bridge | 
**CloudProvider** | **string** | The cloud provider where the bridge resides | 
**Region** | **string** | The cloud provider region where the bridge resides | 

## Methods

### NewBridgeRequest

`func NewBridgeRequest(name string, cloudProvider string, region string, ) *BridgeRequest`

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


### GetCloudProvider

`func (o *BridgeRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *BridgeRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *BridgeRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *BridgeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BridgeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BridgeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


