/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorNamespaceRequestMeta struct for ConnectorNamespaceRequestMeta
type ConnectorNamespaceRequestMeta struct {
	// Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
	Name *string `json:"name,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewConnectorNamespaceRequestMeta instantiates a new ConnectorNamespaceRequestMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceRequestMeta() *ConnectorNamespaceRequestMeta {
	this := ConnectorNamespaceRequestMeta{}
	return &this
}

// NewConnectorNamespaceRequestMetaWithDefaults instantiates a new ConnectorNamespaceRequestMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceRequestMetaWithDefaults() *ConnectorNamespaceRequestMeta {
	this := ConnectorNamespaceRequestMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorNamespaceRequestMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceRequestMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorNamespaceRequestMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorNamespaceRequestMeta) SetName(v string) {
	o.Name = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorNamespaceRequestMeta) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceRequestMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorNamespaceRequestMeta) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorNamespaceRequestMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o ConnectorNamespaceRequestMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceRequestMeta struct {
	value *ConnectorNamespaceRequestMeta
	isSet bool
}

func (v NullableConnectorNamespaceRequestMeta) Get() *ConnectorNamespaceRequestMeta {
	return v.value
}

func (v *NullableConnectorNamespaceRequestMeta) Set(val *ConnectorNamespaceRequestMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceRequestMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceRequestMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceRequestMeta(val *ConnectorNamespaceRequestMeta) *NullableConnectorNamespaceRequestMeta {
	return &NullableConnectorNamespaceRequestMeta{value: val, isSet: true}
}

func (v NullableConnectorNamespaceRequestMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceRequestMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


