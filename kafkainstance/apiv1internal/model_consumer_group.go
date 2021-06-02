/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstance

import (
	"encoding/json"
)

// ConsumerGroup A group of Kafka consumers
type ConsumerGroup struct {
	// Unique identifier for the consumer group
	GroupId string `json:"groupId"`
	// The list of consumers associated with this consumer group
	Consumers []Consumer `json:"consumers"`
}

// NewConsumerGroup instantiates a new ConsumerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroup(groupId string, consumers []Consumer) *ConsumerGroup {
	this := ConsumerGroup{}
	this.GroupId = groupId
	this.Consumers = consumers
	return &this
}

// NewConsumerGroupWithDefaults instantiates a new ConsumerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupWithDefaults() *ConsumerGroup {
	this := ConsumerGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ConsumerGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ConsumerGroup) SetGroupId(v string) {
	o.GroupId = v
}

// GetConsumers returns the Consumers field value
func (o *ConsumerGroup) GetConsumers() []Consumer {
	if o == nil {
		var ret []Consumer
		return ret
	}

	return o.Consumers
}

// GetConsumersOk returns a tuple with the Consumers field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetConsumersOk() (*[]Consumer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Consumers, true
}

// SetConsumers sets field value
func (o *ConsumerGroup) SetConsumers(v []Consumer) {
	o.Consumers = v
}

func (o ConsumerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["consumers"] = o.Consumers
	}
	return json.Marshal(toSerialize)
}

type NullableConsumerGroup struct {
	value *ConsumerGroup
	isSet bool
}

func (v NullableConsumerGroup) Get() *ConsumerGroup {
	return v.value
}

func (v *NullableConsumerGroup) Set(val *ConsumerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroup(val *ConsumerGroup) *NullableConsumerGroup {
	return &NullableConsumerGroup{value: val, isSet: true}
}

func (v NullableConsumerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


