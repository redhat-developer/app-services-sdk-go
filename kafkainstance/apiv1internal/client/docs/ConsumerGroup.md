# ConsumerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique identifier for the consumer group | 
**State** | Pointer to [**ConsumerGroupState**](ConsumerGroupState.md) |  | [optional] 
**Consumers** | [**[]Consumer**](Consumer.md) | The list of consumers associated with this consumer group | 
**Metrics** | Pointer to [**ConsumerGroupMetrics**](ConsumerGroupMetrics.md) |  | [optional] 


## Methods

### NewConsumerGroup

`func NewConsumerGroup(groupId string, consumers []Consumer, ) *ConsumerGroup`

NewConsumerGroup instantiates a new ConsumerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupWithDefaults

`func NewConsumerGroupWithDefaults() *ConsumerGroup`

NewConsumerGroupWithDefaults instantiates a new ConsumerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetGroupId

`func (o *ConsumerGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConsumerGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConsumerGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



### GetState

`func (o *ConsumerGroup) GetState() ConsumerGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConsumerGroup) GetStateOk() (*ConsumerGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConsumerGroup) SetState(v ConsumerGroupState)`

SetState sets State field to given value.

### HasState

`func (o *ConsumerGroup) HasState() bool`

HasState returns a boolean if a field has been set.


### GetConsumers

`func (o *ConsumerGroup) GetConsumers() []Consumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ConsumerGroup) GetConsumersOk() (*[]Consumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ConsumerGroup) SetConsumers(v []Consumer)`

SetConsumers sets Consumers field to given value.



### GetMetrics

`func (o *ConsumerGroup) GetMetrics() ConsumerGroupMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ConsumerGroup) GetMetricsOk() (*ConsumerGroupMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ConsumerGroup) SetMetrics(v ConsumerGroupMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ConsumerGroup) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

