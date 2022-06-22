# RegionCapacityListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | **string** | kafka instance type | 
**AvailableSizes** | **[]string** | list of available Kafka instance sizes that can be created in this region when taking account current capacity and regional limits | 

## Methods

### NewRegionCapacityListItem

`func NewRegionCapacityListItem(instanceType string, availableSizes []string, ) *RegionCapacityListItem`

NewRegionCapacityListItem instantiates a new RegionCapacityListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionCapacityListItemWithDefaults

`func NewRegionCapacityListItemWithDefaults() *RegionCapacityListItem`

NewRegionCapacityListItemWithDefaults instantiates a new RegionCapacityListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *RegionCapacityListItem) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RegionCapacityListItem) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RegionCapacityListItem) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetAvailableSizes

`func (o *RegionCapacityListItem) GetAvailableSizes() []string`

GetAvailableSizes returns the AvailableSizes field if non-nil, zero value otherwise.

### GetAvailableSizesOk

`func (o *RegionCapacityListItem) GetAvailableSizesOk() (*[]string, bool)`

GetAvailableSizesOk returns a tuple with the AvailableSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSizes

`func (o *RegionCapacityListItem) SetAvailableSizes(v []string)`

SetAvailableSizes sets AvailableSizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


