/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ObjectReference struct for ObjectReference
type ObjectReference struct {

	Id *string `json:"id,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

}

// NewObjectReference instantiates a new ObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectReference() *ObjectReference {
	this := ObjectReference{}
	return &this
}

// NewObjectReferenceWithDefaults instantiates a new ObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectReferenceWithDefaults() *ObjectReference {
	this := ObjectReference{}




	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectReference) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectReference) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectReference) SetId(v string) {
	o.Id = &v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ObjectReference) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ObjectReference) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ObjectReference) SetKind(v string) {
	o.Kind = &v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *ObjectReference) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ObjectReference) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ObjectReference) SetHref(v string) {
	o.Href = &v
}


func (o ObjectReference) MarshalJSON() ([]byte, error) {
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
    
	return json.Marshal(toSerialize)
}

type NullableObjectReference struct {
	value *ObjectReference
	isSet bool
}

func (v NullableObjectReference) Get() *ObjectReference {
	return v.value
}

func (v *NullableObjectReference) Set(val *ObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectReference(val *ObjectReference) *NullableObjectReference {
	return &NullableObjectReference{value: val, isSet: true}
}

func (v NullableObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

