/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmt

import (
	"encoding/json"
)

// VersionMetadataAllOf struct for VersionMetadataAllOf
type VersionMetadataAllOf struct {
	Collections *[]ObjectReference `json:"collections,omitempty"`
}

// NewVersionMetadataAllOf instantiates a new VersionMetadataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionMetadataAllOf() *VersionMetadataAllOf {
	this := VersionMetadataAllOf{}
	return &this
}

// NewVersionMetadataAllOfWithDefaults instantiates a new VersionMetadataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionMetadataAllOfWithDefaults() *VersionMetadataAllOf {
	this := VersionMetadataAllOf{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *VersionMetadataAllOf) GetCollections() []ObjectReference {
	if o == nil || o.Collections == nil {
		var ret []ObjectReference
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadataAllOf) GetCollectionsOk() (*[]ObjectReference, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *VersionMetadataAllOf) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []ObjectReference and assigns it to the Collections field.
func (o *VersionMetadataAllOf) SetCollections(v []ObjectReference) {
	o.Collections = &v
}

func (o VersionMetadataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	return json.Marshal(toSerialize)
}

type NullableVersionMetadataAllOf struct {
	value *VersionMetadataAllOf
	isSet bool
}

func (v NullableVersionMetadataAllOf) Get() *VersionMetadataAllOf {
	return v.value
}

func (v *NullableVersionMetadataAllOf) Set(val *VersionMetadataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionMetadataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionMetadataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionMetadataAllOf(val *VersionMetadataAllOf) *NullableVersionMetadataAllOf {
	return &NullableVersionMetadataAllOf{value: val, isSet: true}
}

func (v NullableVersionMetadataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionMetadataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


