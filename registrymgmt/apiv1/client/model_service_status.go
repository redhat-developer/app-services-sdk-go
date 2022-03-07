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
)

// ServiceStatus Schema for the service status response body
type ServiceStatus struct {

	// Boolean property indicating if the maximum number of total Registry instances have been reached, therefore creation of more instances should not be allowed.
	MaxInstancesReached *bool `json:"max_instances_reached,omitempty"`

}

// NewServiceStatus instantiates a new ServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStatus() *ServiceStatus {
	this := ServiceStatus{}
	return &this
}

// NewServiceStatusWithDefaults instantiates a new ServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStatusWithDefaults() *ServiceStatus {
	this := ServiceStatus{}


	return &this
}


// GetMaxInstancesReached returns the MaxInstancesReached field value if set, zero value otherwise.
func (o *ServiceStatus) GetMaxInstancesReached() bool {
	if o == nil || o.MaxInstancesReached == nil {
		var ret bool
		return ret
	}
	return *o.MaxInstancesReached
}

// GetMaxInstancesReachedOk returns a tuple with the MaxInstancesReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetMaxInstancesReachedOk() (*bool, bool) {
	if o == nil || o.MaxInstancesReached == nil {
		return nil, false
	}
	return o.MaxInstancesReached, true
}

// HasMaxInstancesReached returns a boolean if a field has been set.
func (o *ServiceStatus) HasMaxInstancesReached() bool {
	if o != nil && o.MaxInstancesReached != nil {
		return true
	}

	return false
}

// SetMaxInstancesReached gets a reference to the given bool and assigns it to the MaxInstancesReached field.
func (o *ServiceStatus) SetMaxInstancesReached(v bool) {
	o.MaxInstancesReached = &v
}


func (o ServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.MaxInstancesReached != nil {
		toSerialize["max_instances_reached"] = o.MaxInstancesReached
	}
    
	return json.Marshal(toSerialize)
}

type NullableServiceStatus struct {
	value *ServiceStatus
	isSet bool
}

func (v NullableServiceStatus) Get() *ServiceStatus {
	return v.value
}

func (v *NullableServiceStatus) Set(val *ServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStatus(val *ServiceStatus) *NullableServiceStatus {
	return &NullableServiceStatus{value: val, isSet: true}
}

func (v NullableServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

