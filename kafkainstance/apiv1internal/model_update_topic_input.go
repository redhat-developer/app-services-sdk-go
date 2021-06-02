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

// UpdateTopicInput Kafka Topic (A feed where records are stored and published)
type UpdateTopicInput struct {
	// Topic configuration entry.
	Config *[]ConfigEntry `json:"config,omitempty"`
}

// NewUpdateTopicInput instantiates a new UpdateTopicInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTopicInput() *UpdateTopicInput {
	this := UpdateTopicInput{}
	return &this
}

// NewUpdateTopicInputWithDefaults instantiates a new UpdateTopicInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTopicInputWithDefaults() *UpdateTopicInput {
	this := UpdateTopicInput{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateTopicInput) GetConfig() []ConfigEntry {
	if o == nil || o.Config == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTopicInput) GetConfigOk() (*[]ConfigEntry, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateTopicInput) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *UpdateTopicInput) SetConfig(v []ConfigEntry) {
	o.Config = &v
}

func (o UpdateTopicInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTopicInput struct {
	value *UpdateTopicInput
	isSet bool
}

func (v NullableUpdateTopicInput) Get() *UpdateTopicInput {
	return v.value
}

func (v *NullableUpdateTopicInput) Set(val *UpdateTopicInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTopicInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTopicInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTopicInput(val *UpdateTopicInput) *NullableUpdateTopicInput {
	return &NullableUpdateTopicInput{value: val, isSet: true}
}

func (v NullableUpdateTopicInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTopicInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


