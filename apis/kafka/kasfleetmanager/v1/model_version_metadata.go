/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasfleetmanager

import (
	"encoding/json"
)

// VersionMetadata struct for VersionMetadata
type VersionMetadata struct {
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
	Collections *[]ObjectReference `json:"collections,omitempty"`
}

// NewVersionMetadata instantiates a new VersionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionMetadata() *VersionMetadata {
	this := VersionMetadata{}
	return &this
}

// NewVersionMetadataWithDefaults instantiates a new VersionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionMetadataWithDefaults() *VersionMetadata {
	this := VersionMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VersionMetadata) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VersionMetadata) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VersionMetadata) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *VersionMetadata) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *VersionMetadata) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *VersionMetadata) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VersionMetadata) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VersionMetadata) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VersionMetadata) SetHref(v string) {
	o.Href = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *VersionMetadata) GetCollections() []ObjectReference {
	if o == nil || o.Collections == nil {
		var ret []ObjectReference
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionMetadata) GetCollectionsOk() (*[]ObjectReference, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *VersionMetadata) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []ObjectReference and assigns it to the Collections field.
func (o *VersionMetadata) SetCollections(v []ObjectReference) {
	o.Collections = &v
}

func (o VersionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	return json.Marshal(toSerialize)
}

type NullableVersionMetadata struct {
	value *VersionMetadata
	isSet bool
}

func (v NullableVersionMetadata) Get() *VersionMetadata {
	return v.value
}

func (v *NullableVersionMetadata) Set(val *VersionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionMetadata(val *VersionMetadata) *NullableVersionMetadata {
	return &NullableVersionMetadata{value: val, isSet: true}
}

func (v NullableVersionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


