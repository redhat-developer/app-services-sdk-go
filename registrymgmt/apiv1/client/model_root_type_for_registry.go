/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.1
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// RootTypeForRegistry Service Registry instance within a multi-tenant deployment.
type RootTypeForRegistry struct {

	Id string `json:"id"`

	Status RegistryStatusValue `json:"status"`

	RegistryUrl string `json:"registryUrl"`

	// User-defined Registry name. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`

}

// NewRootTypeForRegistry instantiates a new RootTypeForRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootTypeForRegistry(id string, status RegistryStatusValue, registryUrl string) *RootTypeForRegistry {
	this := RootTypeForRegistry{}
	this.Id = id
	this.Status = status
	this.RegistryUrl = registryUrl
	return &this
}

// NewRootTypeForRegistryWithDefaults instantiates a new RootTypeForRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTypeForRegistryWithDefaults() *RootTypeForRegistry {
	this := RootTypeForRegistry{}






	return &this
}


// GetId returns the Id field value
func (o *RootTypeForRegistry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RootTypeForRegistry) SetId(v string) {
	o.Id = v
}


// GetStatus returns the Status field value
func (o *RootTypeForRegistry) GetStatus() RegistryStatusValue {
	if o == nil {
		var ret RegistryStatusValue
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetStatusOk() (*RegistryStatusValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RootTypeForRegistry) SetStatus(v RegistryStatusValue) {
	o.Status = v
}


// GetRegistryUrl returns the RegistryUrl field value
func (o *RootTypeForRegistry) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetRegistryUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *RootTypeForRegistry) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RootTypeForRegistry) SetName(v string) {
	o.Name = &v
}


// GetRegistryDeploymentId returns the RegistryDeploymentId field value if set, zero value otherwise.
func (o *RootTypeForRegistry) GetRegistryDeploymentId() int32 {
	if o == nil || o.RegistryDeploymentId == nil {
		var ret int32
		return ret
	}
	return *o.RegistryDeploymentId
}

// GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootTypeForRegistry) GetRegistryDeploymentIdOk() (*int32, bool) {
	if o == nil || o.RegistryDeploymentId == nil {
		return nil, false
	}
	return o.RegistryDeploymentId, true
}

// HasRegistryDeploymentId returns a boolean if a field has been set.
func (o *RootTypeForRegistry) HasRegistryDeploymentId() bool {
	if o != nil && o.RegistryDeploymentId != nil {
		return true
	}

	return false
}

// SetRegistryDeploymentId gets a reference to the given int32 and assigns it to the RegistryDeploymentId field.
func (o *RootTypeForRegistry) SetRegistryDeploymentId(v int32) {
	o.RegistryDeploymentId = &v
}


func (o RootTypeForRegistry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["id"] = o.Id
	}
    
	if true {
		toSerialize["status"] = o.Status
	}
    
	if true {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.RegistryDeploymentId != nil {
		toSerialize["registryDeploymentId"] = o.RegistryDeploymentId
	}
    
	return json.Marshal(toSerialize)
}

type NullableRootTypeForRegistry struct {
	value *RootTypeForRegistry
	isSet bool
}

func (v NullableRootTypeForRegistry) Get() *RootTypeForRegistry {
	return v.value
}

func (v *NullableRootTypeForRegistry) Set(val *RootTypeForRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableRootTypeForRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableRootTypeForRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootTypeForRegistry(val *RootTypeForRegistry) *NullableRootTypeForRegistry {
	return &NullableRootTypeForRegistry{value: val, isSet: true}
}

func (v NullableRootTypeForRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootTypeForRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

