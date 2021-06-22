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

// Registry struct for Registry
type Registry struct {

	Id int32 `json:"id"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

	Status RegistryStatusValue `json:"status"`

	RegistryUrl string `json:"registryUrl"`

	// User-defined Registry name. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`

}

// NewRegistry instantiates a new Registry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistry(id int32, status RegistryStatusValue, registryUrl string) *Registry {
	this := Registry{}
	this.Id = id
	this.Status = status
	this.RegistryUrl = registryUrl
	return &this
}

// NewRegistryWithDefaults instantiates a new Registry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryWithDefaults() *Registry {
	this := Registry{}








	return &this
}


// GetId returns the Id field value
func (o *Registry) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Registry) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Registry) SetId(v int32) {
	o.Id = v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Registry) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Registry) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Registry) SetKind(v string) {
	o.Kind = &v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *Registry) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Registry) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Registry) SetHref(v string) {
	o.Href = &v
}


// GetStatus returns the Status field value
func (o *Registry) GetStatus() RegistryStatusValue {
	if o == nil {
		var ret RegistryStatusValue
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Registry) GetStatusOk() (*RegistryStatusValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Registry) SetStatus(v RegistryStatusValue) {
	o.Status = v
}


// GetRegistryUrl returns the RegistryUrl field value
func (o *Registry) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *Registry) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *Registry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Registry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Registry) SetName(v string) {
	o.Name = &v
}


// GetRegistryDeploymentId returns the RegistryDeploymentId field value if set, zero value otherwise.
func (o *Registry) GetRegistryDeploymentId() int32 {
	if o == nil || o.RegistryDeploymentId == nil {
		var ret int32
		return ret
	}
	return *o.RegistryDeploymentId
}

// GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryDeploymentIdOk() (*int32, bool) {
	if o == nil || o.RegistryDeploymentId == nil {
		return nil, false
	}
	return o.RegistryDeploymentId, true
}

// HasRegistryDeploymentId returns a boolean if a field has been set.
func (o *Registry) HasRegistryDeploymentId() bool {
	if o != nil && o.RegistryDeploymentId != nil {
		return true
	}

	return false
}

// SetRegistryDeploymentId gets a reference to the given int32 and assigns it to the RegistryDeploymentId field.
func (o *Registry) SetRegistryDeploymentId(v int32) {
	o.RegistryDeploymentId = &v
}


func (o Registry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["id"] = o.Id
	}
    
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
    
	if o.Href != nil {
		toSerialize["href"] = o.Href
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

type NullableRegistry struct {
	value *Registry
	isSet bool
}

func (v NullableRegistry) Get() *Registry {
	return v.value
}

func (v *NullableRegistry) Set(val *Registry) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistry(val *Registry) *NullableRegistry {
	return &NullableRegistry{value: val, isSet: true}
}

func (v NullableRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

