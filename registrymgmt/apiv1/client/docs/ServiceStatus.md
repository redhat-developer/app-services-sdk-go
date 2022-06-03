# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxInstancesReached** | Pointer to **bool** | Boolean property indicating if the maximum number of total Registry instances have been reached, therefore creation of more instances should not be allowed. | [optional] 

## Methods

### NewServiceStatus

`func NewServiceStatus() *ServiceStatus`

NewServiceStatus instantiates a new ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusWithDefaults

`func NewServiceStatusWithDefaults() *ServiceStatus`

NewServiceStatusWithDefaults instantiates a new ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxInstancesReached

`func (o *ServiceStatus) GetMaxInstancesReached() bool`

GetMaxInstancesReached returns the MaxInstancesReached field if non-nil, zero value otherwise.

### GetMaxInstancesReachedOk

`func (o *ServiceStatus) GetMaxInstancesReachedOk() (*bool, bool)`

GetMaxInstancesReachedOk returns a tuple with the MaxInstancesReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstancesReached

`func (o *ServiceStatus) SetMaxInstancesReached(v bool)`

SetMaxInstancesReached sets MaxInstancesReached field to given value.

### HasMaxInstancesReached

`func (o *ServiceStatus) HasMaxInstancesReached() bool`

HasMaxInstancesReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


