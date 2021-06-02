/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmt

import (
	"encoding/json"
)

// KafkaRequestListAllOf struct for KafkaRequestListAllOf
type KafkaRequestListAllOf struct {
	Items *[]KafkaRequest `json:"items,omitempty"`
}

// NewKafkaRequestListAllOf instantiates a new KafkaRequestListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequestListAllOf() *KafkaRequestListAllOf {
	this := KafkaRequestListAllOf{}
	return &this
}

// NewKafkaRequestListAllOfWithDefaults instantiates a new KafkaRequestListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestListAllOfWithDefaults() *KafkaRequestListAllOf {
	this := KafkaRequestListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *KafkaRequestListAllOf) GetItems() []KafkaRequest {
	if o == nil || o.Items == nil {
		var ret []KafkaRequest
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestListAllOf) GetItemsOk() (*[]KafkaRequest, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *KafkaRequestListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KafkaRequest and assigns it to the Items field.
func (o *KafkaRequestListAllOf) SetItems(v []KafkaRequest) {
	o.Items = &v
}

func (o KafkaRequestListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaRequestListAllOf struct {
	value *KafkaRequestListAllOf
	isSet bool
}

func (v NullableKafkaRequestListAllOf) Get() *KafkaRequestListAllOf {
	return v.value
}

func (v *NullableKafkaRequestListAllOf) Set(val *KafkaRequestListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequestListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequestListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequestListAllOf(val *KafkaRequestListAllOf) *NullableKafkaRequestListAllOf {
	return &NullableKafkaRequestListAllOf{value: val, isSet: true}
}

func (v NullableKafkaRequestListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequestListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

