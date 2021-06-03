/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// ServiceAccountListItemAllOf struct for ServiceAccountListItemAllOf
type ServiceAccountListItemAllOf struct {

	// server generated unique id of the service account
	Id *string `json:"id,omitempty"`

	// client id of the service account
	ClientId *string `json:"client_id,omitempty"`

	// name of the service account
	Name *string `json:"name,omitempty"`

	// owner of the service account
	Owner *string `json:"owner,omitempty"`

	// service account creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// description of the service account
	Description *string `json:"description,omitempty"`
}

// NewServiceAccountListItemAllOf instantiates a new ServiceAccountListItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountListItemAllOf() *ServiceAccountListItemAllOf {
	this := ServiceAccountListItemAllOf{}
	return &this
}

// NewServiceAccountListItemAllOfWithDefaults instantiates a new ServiceAccountListItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountListItemAllOfWithDefaults() *ServiceAccountListItemAllOf {
	this := ServiceAccountListItemAllOf{}

	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountListItemAllOf) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceAccountListItemAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountListItemAllOf) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServiceAccountListItemAllOf) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountListItemAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountListItemAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountListItemAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountListItemAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountListItemAllOf) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceAccountListItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}

	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	return json.Marshal(toSerialize)
}

type NullableServiceAccountListItemAllOf struct {
	value *ServiceAccountListItemAllOf
	isSet bool
}

func (v NullableServiceAccountListItemAllOf) Get() *ServiceAccountListItemAllOf {
	return v.value
}

func (v *NullableServiceAccountListItemAllOf) Set(val *ServiceAccountListItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountListItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountListItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountListItemAllOf(val *ServiceAccountListItemAllOf) *NullableServiceAccountListItemAllOf {
	return &NullableServiceAccountListItemAllOf{value: val, isSet: true}
}

func (v NullableServiceAccountListItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountListItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
