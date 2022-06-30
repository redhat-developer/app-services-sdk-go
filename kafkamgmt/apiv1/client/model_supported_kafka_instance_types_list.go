/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// SupportedKafkaInstanceTypesList struct for SupportedKafkaInstanceTypesList
type SupportedKafkaInstanceTypesList struct {
	InstanceTypes *[]SupportedKafkaInstanceType `json:"instance_types,omitempty"`
}

// NewSupportedKafkaInstanceTypesList instantiates a new SupportedKafkaInstanceTypesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedKafkaInstanceTypesList() *SupportedKafkaInstanceTypesList {
	this := SupportedKafkaInstanceTypesList{}
	return &this
}

// NewSupportedKafkaInstanceTypesListWithDefaults instantiates a new SupportedKafkaInstanceTypesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedKafkaInstanceTypesListWithDefaults() *SupportedKafkaInstanceTypesList {
	this := SupportedKafkaInstanceTypesList{}
	return &this
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceTypesList) GetInstanceTypes() []SupportedKafkaInstanceType {
	if o == nil || o.InstanceTypes == nil {
		var ret []SupportedKafkaInstanceType
		return ret
	}
	return *o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceTypesList) GetInstanceTypesOk() (*[]SupportedKafkaInstanceType, bool) {
	if o == nil || o.InstanceTypes == nil {
		return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceTypesList) HasInstanceTypes() bool {
	if o != nil && o.InstanceTypes != nil {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []SupportedKafkaInstanceType and assigns it to the InstanceTypes field.
func (o *SupportedKafkaInstanceTypesList) SetInstanceTypes(v []SupportedKafkaInstanceType) {
	o.InstanceTypes = &v
}

func (o SupportedKafkaInstanceTypesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceTypes != nil {
		toSerialize["instance_types"] = o.InstanceTypes
	}
	return json.Marshal(toSerialize)
}

type NullableSupportedKafkaInstanceTypesList struct {
	value *SupportedKafkaInstanceTypesList
	isSet bool
}

func (v NullableSupportedKafkaInstanceTypesList) Get() *SupportedKafkaInstanceTypesList {
	return v.value
}

func (v *NullableSupportedKafkaInstanceTypesList) Set(val *SupportedKafkaInstanceTypesList) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedKafkaInstanceTypesList) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedKafkaInstanceTypesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedKafkaInstanceTypesList(val *SupportedKafkaInstanceTypesList) *NullableSupportedKafkaInstanceTypesList {
	return &NullableSupportedKafkaInstanceTypesList{value: val, isSet: true}
}

func (v NullableSupportedKafkaInstanceTypesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedKafkaInstanceTypesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


