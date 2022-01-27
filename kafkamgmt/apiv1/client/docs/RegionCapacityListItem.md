# RegionCapacityListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | kafka instance type | [optional] 
**MaxCapacityReached** | **bool** | flag indicating whether the capacity for the instance type in the region is reached | 


## Methods

### NewRegionCapacityListItem

`func NewRegionCapacityListItem(maxCapacityReached bool, ) *RegionCapacityListItem`

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

### HasInstanceType

`func (o *RegionCapacityListItem) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


### GetMaxCapacityReached

`func (o *RegionCapacityListItem) GetMaxCapacityReached() bool`

GetMaxCapacityReached returns the MaxCapacityReached field if non-nil, zero value otherwise.

### GetMaxCapacityReachedOk

`func (o *RegionCapacityListItem) GetMaxCapacityReachedOk() (*bool, bool)`

GetMaxCapacityReachedOk returns a tuple with the MaxCapacityReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacityReached

`func (o *RegionCapacityListItem) SetMaxCapacityReached(v bool)`

SetMaxCapacityReached sets MaxCapacityReached field to given value.




[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

