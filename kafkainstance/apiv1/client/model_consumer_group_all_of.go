/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.12.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ConsumerGroupAllOf A group of Kafka consumers
type ConsumerGroupAllOf struct {
	// Unique identifier for the consumer group
	GroupId string `json:"groupId"`
	State *ConsumerGroupState `json:"state,omitempty"`
	// The list of consumers associated with this consumer group
	Consumers []Consumer `json:"consumers"`
	Metrics *ConsumerGroupMetrics `json:"metrics,omitempty"`
}

// NewConsumerGroupAllOf instantiates a new ConsumerGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupAllOf(groupId string, consumers []Consumer) *ConsumerGroupAllOf {
	this := ConsumerGroupAllOf{}
	this.GroupId = groupId
	this.Consumers = consumers
	return &this
}

// NewConsumerGroupAllOfWithDefaults instantiates a new ConsumerGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupAllOfWithDefaults() *ConsumerGroupAllOf {
	this := ConsumerGroupAllOf{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ConsumerGroupAllOf) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupAllOf) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ConsumerGroupAllOf) SetGroupId(v string) {
	o.GroupId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ConsumerGroupAllOf) GetState() ConsumerGroupState {
	if o == nil || o.State == nil {
		var ret ConsumerGroupState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupAllOf) GetStateOk() (*ConsumerGroupState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ConsumerGroupAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ConsumerGroupState and assigns it to the State field.
func (o *ConsumerGroupAllOf) SetState(v ConsumerGroupState) {
	o.State = &v
}

// GetConsumers returns the Consumers field value
func (o *ConsumerGroupAllOf) GetConsumers() []Consumer {
	if o == nil {
		var ret []Consumer
		return ret
	}

	return o.Consumers
}

// GetConsumersOk returns a tuple with the Consumers field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupAllOf) GetConsumersOk() (*[]Consumer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Consumers, true
}

// SetConsumers sets field value
func (o *ConsumerGroupAllOf) SetConsumers(v []Consumer) {
	o.Consumers = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ConsumerGroupAllOf) GetMetrics() ConsumerGroupMetrics {
	if o == nil || o.Metrics == nil {
		var ret ConsumerGroupMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupAllOf) GetMetricsOk() (*ConsumerGroupMetrics, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ConsumerGroupAllOf) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given ConsumerGroupMetrics and assigns it to the Metrics field.
func (o *ConsumerGroupAllOf) SetMetrics(v ConsumerGroupMetrics) {
	o.Metrics = &v
}

func (o ConsumerGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["consumers"] = o.Consumers
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupAllOf struct {
	value *ConsumerGroupAllOf
	isSet bool
}

func (v NullableConsumerGroupAllOf) Get() *ConsumerGroupAllOf {
	return v.value
}

func (v *NullableConsumerGroupAllOf) Set(val *ConsumerGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupAllOf(val *ConsumerGroupAllOf) *NullableConsumerGroupAllOf {
	return &NullableConsumerGroupAllOf{value: val, isSet: true}
}

func (v NullableConsumerGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


