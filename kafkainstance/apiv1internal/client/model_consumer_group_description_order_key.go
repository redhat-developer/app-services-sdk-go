/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.11.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// ConsumerGroupDescriptionOrderKey the model 'ConsumerGroupDescriptionOrderKey'
type ConsumerGroupDescriptionOrderKey string

// List of ConsumerGroupDescriptionOrderKey
const (
	CONSUMERGROUPDESCRIPTIONORDERKEY_OFFSET ConsumerGroupDescriptionOrderKey = "offset"
	CONSUMERGROUPDESCRIPTIONORDERKEY_END_OFFSET ConsumerGroupDescriptionOrderKey = "endOffset"
	CONSUMERGROUPDESCRIPTIONORDERKEY_LAG ConsumerGroupDescriptionOrderKey = "lag"
	CONSUMERGROUPDESCRIPTIONORDERKEY_PARTITION ConsumerGroupDescriptionOrderKey = "partition"
)

var allowedConsumerGroupDescriptionOrderKeyEnumValues = []ConsumerGroupDescriptionOrderKey{
	"offset",
	"endOffset",
	"lag",
	"partition",
}

func (v *ConsumerGroupDescriptionOrderKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsumerGroupDescriptionOrderKey(value)
	for _, existing := range allowedConsumerGroupDescriptionOrderKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsumerGroupDescriptionOrderKey", value)
}

// NewConsumerGroupDescriptionOrderKeyFromValue returns a pointer to a valid ConsumerGroupDescriptionOrderKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsumerGroupDescriptionOrderKeyFromValue(v string) (*ConsumerGroupDescriptionOrderKey, error) {
	ev := ConsumerGroupDescriptionOrderKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsumerGroupDescriptionOrderKey: valid values are %v", v, allowedConsumerGroupDescriptionOrderKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsumerGroupDescriptionOrderKey) IsValid() bool {
	for _, existing := range allowedConsumerGroupDescriptionOrderKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsumerGroupDescriptionOrderKey value
func (v ConsumerGroupDescriptionOrderKey) Ptr() *ConsumerGroupDescriptionOrderKey {
	return &v
}

type NullableConsumerGroupDescriptionOrderKey struct {
	value *ConsumerGroupDescriptionOrderKey
	isSet bool
}

func (v NullableConsumerGroupDescriptionOrderKey) Get() *ConsumerGroupDescriptionOrderKey {
	return v.value
}

func (v *NullableConsumerGroupDescriptionOrderKey) Set(val *ConsumerGroupDescriptionOrderKey) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupDescriptionOrderKey) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupDescriptionOrderKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupDescriptionOrderKey(val *ConsumerGroupDescriptionOrderKey) *NullableConsumerGroupDescriptionOrderKey {
	return &NullableConsumerGroupDescriptionOrderKey{value: val, isSet: true}
}

func (v NullableConsumerGroupDescriptionOrderKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupDescriptionOrderKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

