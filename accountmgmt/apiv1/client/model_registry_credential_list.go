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

// RegistryCredentialList struct for RegistryCredentialList
type RegistryCredentialList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []RegistryCredential `json:"items"`
}

// NewRegistryCredentialList instantiates a new RegistryCredentialList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCredentialList(kind string, page int32, size int32, total int32, items []RegistryCredential) *RegistryCredentialList {
	this := RegistryCredentialList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewRegistryCredentialListWithDefaults instantiates a new RegistryCredentialList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCredentialListWithDefaults() *RegistryCredentialList {
	this := RegistryCredentialList{}
	return &this
}

// GetKind returns the Kind field value
func (o *RegistryCredentialList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *RegistryCredentialList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *RegistryCredentialList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *RegistryCredentialList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *RegistryCredentialList) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *RegistryCredentialList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *RegistryCredentialList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *RegistryCredentialList) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *RegistryCredentialList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *RegistryCredentialList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RegistryCredentialList) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *RegistryCredentialList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *RegistryCredentialList) GetItems() []RegistryCredential {
	if o == nil {
		var ret []RegistryCredential
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RegistryCredentialList) GetItemsOk() (*[]RegistryCredential, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *RegistryCredentialList) SetItems(v []RegistryCredential) {
	o.Items = v
}

func (o RegistryCredentialList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryCredentialList struct {
	value *RegistryCredentialList
	isSet bool
}

func (v NullableRegistryCredentialList) Get() *RegistryCredentialList {
	return v.value
}

func (v *NullableRegistryCredentialList) Set(val *RegistryCredentialList) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCredentialList) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCredentialList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCredentialList(val *RegistryCredentialList) *NullableRegistryCredentialList {
	return &NullableRegistryCredentialList{value: val, isSet: true}
}

func (v NullableRegistryCredentialList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCredentialList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


