/*
 * Service Registry Fleet Manager
 *
 * Service Registry Fleet Manager is a REST API for managing Service Registry instances. Service Registry is a datastore for event schemas and API designs, which is based on the open source Apicurio Registry project.
 *
 * API version: 0.0.6
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
	"time"
)

// Registry struct for Registry
type Registry struct {

	Id string `json:"id"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

	Status RegistryStatusValue `json:"status"`

	RegistryUrl *string `json:"registryUrl,omitempty"`

	BrowserUrl *string `json:"browserUrl,omitempty"`

	// User-defined Registry instance name. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`

	// Registry instance owner.
	Owner *string `json:"owner,omitempty"`

	// Description of the Registry instance.
	Description *string `json:"description,omitempty"`

	// ISO 8601 UTC timestamp.
	CreatedAt time.Time `json:"created_at"`

	// ISO 8601 UTC timestamp.
	UpdatedAt time.Time `json:"updated_at"`

	InstanceType RegistryInstanceTypeValue `json:"instance_type"`

}

// NewRegistry instantiates a new Registry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistry(id string, status RegistryStatusValue, createdAt time.Time, updatedAt time.Time, instanceType RegistryInstanceTypeValue) *Registry {
	this := Registry{}
	this.Id = id
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.InstanceType = instanceType
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
func (o *Registry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Registry) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Registry) SetId(v string) {
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


// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *Registry) GetRegistryUrl() string {
	if o == nil || o.RegistryUrl == nil {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryUrlOk() (*string, bool) {
	if o == nil || o.RegistryUrl == nil {
		return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *Registry) HasRegistryUrl() bool {
	if o != nil && o.RegistryUrl != nil {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *Registry) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}


// GetBrowserUrl returns the BrowserUrl field value if set, zero value otherwise.
func (o *Registry) GetBrowserUrl() string {
	if o == nil || o.BrowserUrl == nil {
		var ret string
		return ret
	}
	return *o.BrowserUrl
}

// GetBrowserUrlOk returns a tuple with the BrowserUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetBrowserUrlOk() (*string, bool) {
	if o == nil || o.BrowserUrl == nil {
		return nil, false
	}
	return o.BrowserUrl, true
}

// HasBrowserUrl returns a boolean if a field has been set.
func (o *Registry) HasBrowserUrl() bool {
	if o != nil && o.BrowserUrl != nil {
		return true
	}

	return false
}

// SetBrowserUrl gets a reference to the given string and assigns it to the BrowserUrl field.
func (o *Registry) SetBrowserUrl(v string) {
	o.BrowserUrl = &v
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


// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Registry) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Registry) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Registry) SetOwner(v string) {
	o.Owner = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Registry) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Registry) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Registry) SetDescription(v string) {
	o.Description = &v
}


// GetCreatedAt returns the CreatedAt field value
func (o *Registry) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Registry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Registry) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}


// GetUpdatedAt returns the UpdatedAt field value
func (o *Registry) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Registry) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Registry) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}


// GetInstanceType returns the InstanceType field value
func (o *Registry) GetInstanceType() RegistryInstanceTypeValue {
	if o == nil {
		var ret RegistryInstanceTypeValue
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *Registry) GetInstanceTypeOk() (*RegistryInstanceTypeValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *Registry) SetInstanceType(v RegistryInstanceTypeValue) {
	o.InstanceType = v
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
    
	if o.RegistryUrl != nil {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
    
	if o.BrowserUrl != nil {
		toSerialize["browserUrl"] = o.BrowserUrl
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.RegistryDeploymentId != nil {
		toSerialize["registryDeploymentId"] = o.RegistryDeploymentId
	}
    
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
    
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
    
	if true {
		toSerialize["instance_type"] = o.InstanceType
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

