/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmt

import (
	"encoding/json"
)

// QuotaCostListAllOf struct for QuotaCostListAllOf
type QuotaCostListAllOf struct {
	Items *[]QuotaCost `json:"items,omitempty"`
}

// NewQuotaCostListAllOf instantiates a new QuotaCostListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaCostListAllOf() *QuotaCostListAllOf {
	this := QuotaCostListAllOf{}
	return &this
}

// NewQuotaCostListAllOfWithDefaults instantiates a new QuotaCostListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaCostListAllOfWithDefaults() *QuotaCostListAllOf {
	this := QuotaCostListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QuotaCostListAllOf) GetItems() []QuotaCost {
	if o == nil || o.Items == nil {
		var ret []QuotaCost
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaCostListAllOf) GetItemsOk() (*[]QuotaCost, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QuotaCostListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QuotaCost and assigns it to the Items field.
func (o *QuotaCostListAllOf) SetItems(v []QuotaCost) {
	o.Items = &v
}

func (o QuotaCostListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaCostListAllOf struct {
	value *QuotaCostListAllOf
	isSet bool
}

func (v NullableQuotaCostListAllOf) Get() *QuotaCostListAllOf {
	return v.value
}

func (v *NullableQuotaCostListAllOf) Set(val *QuotaCostListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaCostListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaCostListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaCostListAllOf(val *QuotaCostListAllOf) *NullableQuotaCostListAllOf {
	return &NullableQuotaCostListAllOf{value: val, isSet: true}
}

func (v NullableQuotaCostListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaCostListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


