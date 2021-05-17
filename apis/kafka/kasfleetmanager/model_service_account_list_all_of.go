/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ServiceAccountListAllOf struct for ServiceAccountListAllOf
type ServiceAccountListAllOf struct {
	Kind string `json:"kind"`
	Items []ServiceAccountListItem `json:"items"`
}

// NewServiceAccountListAllOf instantiates a new ServiceAccountListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountListAllOf(kind string, items []ServiceAccountListItem) *ServiceAccountListAllOf {
	this := ServiceAccountListAllOf{}
	this.Kind = kind
	this.Items = items
	return &this
}

// NewServiceAccountListAllOfWithDefaults instantiates a new ServiceAccountListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListAllOfWithDefaults() *ServiceAccountListAllOf {
	this := ServiceAccountListAllOf{}
	return &this
}

// GetKind returns the Kind field value
func (o *ServiceAccountListAllOf) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOf) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceAccountListAllOf) SetKind(v string) {
	o.Kind = v
}

// GetItems returns the Items field value
func (o *ServiceAccountListAllOf) GetItems() []ServiceAccountListItem {
	if o == nil {
		var ret []ServiceAccountListItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountListAllOf) GetItemsOk() (*[]ServiceAccountListItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ServiceAccountListAllOf) SetItems(v []ServiceAccountListItem) {
	o.Items = v
}

func (o ServiceAccountListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountListAllOf struct {
	value *ServiceAccountListAllOf
	isSet bool
}

func (v NullableServiceAccountListAllOf) Get() *ServiceAccountListAllOf {
	return v.value
}

func (v *NullableServiceAccountListAllOf) Set(val *ServiceAccountListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountListAllOf(val *ServiceAccountListAllOf) *NullableServiceAccountListAllOf {
	return &NullableServiceAccountListAllOf{value: val, isSet: true}
}

func (v NullableServiceAccountListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


