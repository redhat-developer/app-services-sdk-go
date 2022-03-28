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

// ConnectorNamespacePatchRequest A connector namespace patch request
type ConnectorNamespacePatchRequest struct {

	Name *string `json:"name,omitempty"`

	Annotations *[]ConnectorNamespaceRequestMetaAnnotations `json:"annotations,omitempty"`

}

// NewConnectorNamespacePatchRequest instantiates a new ConnectorNamespacePatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespacePatchRequest() *ConnectorNamespacePatchRequest {
	this := ConnectorNamespacePatchRequest{}
	return &this
}

// NewConnectorNamespacePatchRequestWithDefaults instantiates a new ConnectorNamespacePatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespacePatchRequestWithDefaults() *ConnectorNamespacePatchRequest {
	this := ConnectorNamespacePatchRequest{}



	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorNamespacePatchRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespacePatchRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorNamespacePatchRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorNamespacePatchRequest) SetName(v string) {
	o.Name = &v
}


// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorNamespacePatchRequest) GetAnnotations() []ConnectorNamespaceRequestMetaAnnotations {
	if o == nil || o.Annotations == nil {
		var ret []ConnectorNamespaceRequestMetaAnnotations
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespacePatchRequest) GetAnnotationsOk() (*[]ConnectorNamespaceRequestMetaAnnotations, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorNamespacePatchRequest) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []ConnectorNamespaceRequestMetaAnnotations and assigns it to the Annotations field.
func (o *ConnectorNamespacePatchRequest) SetAnnotations(v []ConnectorNamespaceRequestMetaAnnotations) {
	o.Annotations = &v
}


func (o ConnectorNamespacePatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespacePatchRequest struct {
	value *ConnectorNamespacePatchRequest
	isSet bool
}

func (v NullableConnectorNamespacePatchRequest) Get() *ConnectorNamespacePatchRequest {
	return v.value
}

func (v *NullableConnectorNamespacePatchRequest) Set(val *ConnectorNamespacePatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespacePatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespacePatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespacePatchRequest(val *ConnectorNamespacePatchRequest) *NullableConnectorNamespacePatchRequest {
	return &NullableConnectorNamespacePatchRequest{value: val, isSet: true}
}

func (v NullableConnectorNamespacePatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespacePatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

