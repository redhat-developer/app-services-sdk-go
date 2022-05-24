/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// ServiceAccountAllOf struct for ServiceAccountAllOf
type ServiceAccountAllOf struct {

	// server generated unique id of the service account
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	ClientId *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

}

// NewServiceAccountAllOf instantiates a new ServiceAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountAllOf() *ServiceAccountAllOf {
	this := ServiceAccountAllOf{}
	return &this
}

// NewServiceAccountAllOfWithDefaults instantiates a new ServiceAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountAllOfWithDefaults() *ServiceAccountAllOf {
	this := ServiceAccountAllOf{}








	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountAllOf) SetId(v string) {
	o.Id = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountAllOf) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountAllOf) SetDescription(v string) {
	o.Description = &v
}


// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceAccountAllOf) SetClientId(v string) {
	o.ClientId = &v
}


// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ServiceAccountAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}


// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServiceAccountAllOf) SetCreatedBy(v string) {
	o.CreatedBy = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


func (o ServiceAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
    
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
    
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
    
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
    
	return json.Marshal(toSerialize)
}

type NullableServiceAccountAllOf struct {
	value *ServiceAccountAllOf
	isSet bool
}

func (v NullableServiceAccountAllOf) Get() *ServiceAccountAllOf {
	return v.value
}

func (v *NullableServiceAccountAllOf) Set(val *ServiceAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountAllOf(val *ServiceAccountAllOf) *NullableServiceAccountAllOf {
	return &NullableServiceAccountAllOf{value: val, isSet: true}
}

func (v NullableServiceAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

