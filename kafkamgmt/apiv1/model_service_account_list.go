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

// ServiceAccountList struct for ServiceAccountList
type ServiceAccountList struct {
	Kind string `json:"kind"`
	Items []ServiceAccountListItem `json:"items"`
}

// NewServiceAccountList instantiates a new ServiceAccountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountList(kind string, items []ServiceAccountListItem) *ServiceAccountList {
	this := ServiceAccountList{}
	this.Kind = kind
	this.Items = items
	return &this
}

// NewServiceAccountListWithDefaults instantiates a new ServiceAccountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListWithDefaults() *ServiceAccountList {
	this := ServiceAccountList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ServiceAccountList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceAccountList) SetKind(v string) {
	o.Kind = v
}

// GetItems returns the Items field value
func (o *ServiceAccountList) GetItems() []ServiceAccountListItem {
	if o == nil {
		var ret []ServiceAccountListItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountList) GetItemsOk() (*[]ServiceAccountListItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ServiceAccountList) SetItems(v []ServiceAccountListItem) {
	o.Items = v
}

func (o ServiceAccountList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountList struct {
	value *ServiceAccountList
	isSet bool
}

func (v NullableServiceAccountList) Get() *ServiceAccountList {
	return v.value
}

func (v *NullableServiceAccountList) Set(val *ServiceAccountList) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountList) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountList(val *ServiceAccountList) *NullableServiceAccountList {
	return &NullableServiceAccountList{value: val, isSet: true}
}

func (v NullableServiceAccountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

