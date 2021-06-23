# ClusterTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**ClusterId** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**MultiAz** | Pointer to **bool** |  | [optional] 


## Methods

### NewClusterTarget

`func NewClusterTarget(kind string, ) *ClusterTarget`

NewClusterTarget instantiates a new ClusterTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTargetWithDefaults

`func NewClusterTargetWithDefaults() *ClusterTarget`

NewClusterTargetWithDefaults instantiates a new ClusterTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKind

`func (o *ClusterTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterTarget) SetKind(v string)`

SetKind sets Kind field to given value.



### GetClusterId

`func (o *ClusterTarget) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterTarget) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterTarget) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterTarget) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.


### GetCloudProvider

`func (o *ClusterTarget) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterTarget) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterTarget) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ClusterTarget) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.


### GetRegion

`func (o *ClusterTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterTarget) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterTarget) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


### GetMultiAz

`func (o *ClusterTarget) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *ClusterTarget) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *ClusterTarget) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *ClusterTarget) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

