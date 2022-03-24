/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.8.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// ConsumerGroupState the model 'ConsumerGroupState'
type ConsumerGroupState string

// List of ConsumerGroupState
const (
	CONSUMERGROUPSTATE_UNKNOWN ConsumerGroupState = "UNKNOWN"
	CONSUMERGROUPSTATE_PREPARING_REBALANCE ConsumerGroupState = "PREPARING_REBALANCE"
	CONSUMERGROUPSTATE_COMPLETING_REBALANCE ConsumerGroupState = "COMPLETING_REBALANCE"
	CONSUMERGROUPSTATE_STABLE ConsumerGroupState = "STABLE"
	CONSUMERGROUPSTATE_DEAD ConsumerGroupState = "DEAD"
	CONSUMERGROUPSTATE_EMPTY ConsumerGroupState = "EMPTY"
)

var allowedConsumerGroupStateEnumValues = []ConsumerGroupState{
	"UNKNOWN",
	"PREPARING_REBALANCE",
	"COMPLETING_REBALANCE",
	"STABLE",
	"DEAD",
	"EMPTY",
}

func (v *ConsumerGroupState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsumerGroupState(value)
	for _, existing := range allowedConsumerGroupStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsumerGroupState", value)
}

// NewConsumerGroupStateFromValue returns a pointer to a valid ConsumerGroupState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsumerGroupStateFromValue(v string) (*ConsumerGroupState, error) {
	ev := ConsumerGroupState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsumerGroupState: valid values are %v", v, allowedConsumerGroupStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsumerGroupState) IsValid() bool {
	for _, existing := range allowedConsumerGroupStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsumerGroupState value
func (v ConsumerGroupState) Ptr() *ConsumerGroupState {
	return &v
}

type NullableConsumerGroupState struct {
	value *ConsumerGroupState
	isSet bool
}

func (v NullableConsumerGroupState) Get() *ConsumerGroupState {
	return v.value
}

func (v *NullableConsumerGroupState) Set(val *ConsumerGroupState) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupState) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupState(val *ConsumerGroupState) *NullableConsumerGroupState {
	return &NullableConsumerGroupState{value: val, isSet: true}
}

func (v NullableConsumerGroupState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
