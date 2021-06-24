/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.4
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// RegistryCreate Information used to create a new Service Registry instance within a multi-tenant deployment.
type RegistryCreate struct {

	// User-defined Registry name. Does not have to be unique.
	Name *string `json:"name,omitempty"`

}

// NewRegistryCreate instantiates a new RegistryCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCreate() *RegistryCreate {
	this := RegistryCreate{}
	return &this
}

// NewRegistryCreateWithDefaults instantiates a new RegistryCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCreateWithDefaults() *RegistryCreate {
	this := RegistryCreate{}


	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistryCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistryCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistryCreate) SetName(v string) {
	o.Name = &v
}


func (o RegistryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	return json.Marshal(toSerialize)
}

type NullableRegistryCreate struct {
	value *RegistryCreate
	isSet bool
}

func (v NullableRegistryCreate) Get() *RegistryCreate {
	return v.value
}

func (v *NullableRegistryCreate) Set(val *RegistryCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCreate(val *RegistryCreate) *NullableRegistryCreate {
	return &NullableRegistryCreate{value: val, isSet: true}
}

func (v NullableRegistryCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

