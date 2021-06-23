# CloudProviderClusterTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**CloudProvider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**MultiAz** | Pointer to **bool** |  | [optional] 


## Methods

### NewCloudProviderClusterTarget

`func NewCloudProviderClusterTarget(kind string, ) *CloudProviderClusterTarget`

NewCloudProviderClusterTarget instantiates a new CloudProviderClusterTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderClusterTargetWithDefaults

`func NewCloudProviderClusterTargetWithDefaults() *CloudProviderClusterTarget`

NewCloudProviderClusterTargetWithDefaults instantiates a new CloudProviderClusterTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKind

`func (o *CloudProviderClusterTarget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProviderClusterTarget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProviderClusterTarget) SetKind(v string)`

SetKind sets Kind field to given value.



### GetCloudProvider

`func (o *CloudProviderClusterTarget) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CloudProviderClusterTarget) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CloudProviderClusterTarget) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CloudProviderClusterTarget) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.


### GetRegion

`func (o *CloudProviderClusterTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudProviderClusterTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudProviderClusterTarget) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudProviderClusterTarget) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


### GetMultiAz

`func (o *CloudProviderClusterTarget) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *CloudProviderClusterTarget) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *CloudProviderClusterTarget) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *CloudProviderClusterTarget) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

