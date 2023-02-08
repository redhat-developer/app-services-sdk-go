/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.15.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// Values struct for Values
type Values struct {
	Timestamp *int64 `json:"timestamp,omitempty"`
	Value float64 `json:"value"`
}

// NewValues instantiates a new Values object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValues(value float64) *Values {
	this := Values{}
	this.Value = value
	return &this
}

// NewValuesWithDefaults instantiates a new Values object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesWithDefaults() *Values {
	this := Values{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Values) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Values) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Values) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetValue returns the Value field value
func (o *Values) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Values) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Values) SetValue(v float64) {
	o.Value = v
}

func (o Values) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValues struct {
	value *Values
	isSet bool
}

func (v NullableValues) Get() *Values {
	return v.value
}

func (v *NullableValues) Set(val *Values) {
	v.value = val
	v.isSet = true
}

func (v NullableValues) IsSet() bool {
	return v.isSet
}

func (v *NullableValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValues(val *Values) *NullableValues {
	return &NullableValues{value: val, isSet: true}
}

func (v NullableValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


