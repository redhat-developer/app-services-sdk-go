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

// SkuRulesListAllOf struct for SkuRulesListAllOf
type SkuRulesListAllOf struct {
	Items *[]SkuRules `json:"items,omitempty"`
}

// NewSkuRulesListAllOf instantiates a new SkuRulesListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuRulesListAllOf() *SkuRulesListAllOf {
	this := SkuRulesListAllOf{}
	return &this
}

// NewSkuRulesListAllOfWithDefaults instantiates a new SkuRulesListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuRulesListAllOfWithDefaults() *SkuRulesListAllOf {
	this := SkuRulesListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SkuRulesListAllOf) GetItems() []SkuRules {
	if o == nil || o.Items == nil {
		var ret []SkuRules
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRulesListAllOf) GetItemsOk() (*[]SkuRules, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SkuRulesListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SkuRules and assigns it to the Items field.
func (o *SkuRulesListAllOf) SetItems(v []SkuRules) {
	o.Items = &v
}

func (o SkuRulesListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSkuRulesListAllOf struct {
	value *SkuRulesListAllOf
	isSet bool
}

func (v NullableSkuRulesListAllOf) Get() *SkuRulesListAllOf {
	return v.value
}

func (v *NullableSkuRulesListAllOf) Set(val *SkuRulesListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuRulesListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuRulesListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuRulesListAllOf(val *SkuRulesListAllOf) *NullableSkuRulesListAllOf {
	return &NullableSkuRulesListAllOf{value: val, isSet: true}
}

func (v NullableSkuRulesListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuRulesListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


