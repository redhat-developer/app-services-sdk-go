/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.8.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// ServiceAccount Service Account created in MAS-SSO for the Kafka Cluster for authentication
type ServiceAccount struct {
	// server generated unique id of the service account
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	Owner *string `json:"owner,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewServiceAccount instantiates a new ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccount() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// NewServiceAccountWithDefaults instantiates a new ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountWithDefaults() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ServiceAccount) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ServiceAccount) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ServiceAccount) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ServiceAccount) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ServiceAccount) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ServiceAccount) SetHref(v string) {
	o.Href = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceAccount) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceAccount) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceAccount) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ServiceAccount) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ServiceAccount) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ServiceAccount) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ServiceAccount) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ServiceAccount) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ServiceAccount) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServiceAccount) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServiceAccount) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServiceAccount) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccount) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccount) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o ServiceAccount) MarshalJSON() ([]byte, error) {
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
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccount struct {
	value *ServiceAccount
	isSet bool
}

func (v NullableServiceAccount) Get() *ServiceAccount {
	return v.value
}

func (v *NullableServiceAccount) Set(val *ServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccount(val *ServiceAccount) *NullableServiceAccount {
	return &NullableServiceAccount{value: val, isSet: true}
}

func (v NullableServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


