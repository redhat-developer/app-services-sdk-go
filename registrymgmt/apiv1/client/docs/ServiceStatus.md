# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEvalInstancesReached** | Pointer to **bool** | Boolean property indicating if the maximum number of Trial instances have been reached, therefore creation of more eval instances should not be allowed. | [optional] 


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


### GetMaxEvalInstancesReached

`func (o *ServiceStatus) GetMaxEvalInstancesReached() bool`

GetMaxEvalInstancesReached returns the MaxEvalInstancesReached field if non-nil, zero value otherwise.

### GetMaxEvalInstancesReachedOk

`func (o *ServiceStatus) GetMaxEvalInstancesReachedOk() (*bool, bool)`

GetMaxEvalInstancesReachedOk returns a tuple with the MaxEvalInstancesReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEvalInstancesReached

`func (o *ServiceStatus) SetMaxEvalInstancesReached(v bool)`

SetMaxEvalInstancesReached sets MaxEvalInstancesReached field to given value.

### HasMaxEvalInstancesReached

`func (o *ServiceStatus) HasMaxEvalInstancesReached() bool`

HasMaxEvalInstancesReached returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

