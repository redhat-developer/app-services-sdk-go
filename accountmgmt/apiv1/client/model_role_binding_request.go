/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmt

import (
	"encoding/json"
)

// RoleBindingRequest struct for RoleBindingRequest
type RoleBindingRequest struct {
	AccountId *string `json:"account_id,omitempty"`
	ConfigManaged *bool `json:"config_managed,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	RoleId *string `json:"role_id,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRoleBindingRequest instantiates a new RoleBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBindingRequest() *RoleBindingRequest {
	this := RoleBindingRequest{}
	return &this
}

// NewRoleBindingRequestWithDefaults instantiates a new RoleBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingRequestWithDefaults() *RoleBindingRequest {
	this := RoleBindingRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *RoleBindingRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetConfigManaged returns the ConfigManaged field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetConfigManaged() bool {
	if o == nil || o.ConfigManaged == nil {
		var ret bool
		return ret
	}
	return *o.ConfigManaged
}

// GetConfigManagedOk returns a tuple with the ConfigManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetConfigManagedOk() (*bool, bool) {
	if o == nil || o.ConfigManaged == nil {
		return nil, false
	}
	return o.ConfigManaged, true
}

// HasConfigManaged returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasConfigManaged() bool {
	if o != nil && o.ConfigManaged != nil {
		return true
	}

	return false
}

// SetConfigManaged gets a reference to the given bool and assigns it to the ConfigManaged field.
func (o *RoleBindingRequest) SetConfigManaged(v bool) {
	o.ConfigManaged = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *RoleBindingRequest) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleBindingRequest) SetRoleId(v string) {
	o.RoleId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *RoleBindingRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleBindingRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBindingRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleBindingRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleBindingRequest) SetType(v string) {
	o.Type = &v
}

func (o RoleBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ConfigManaged != nil {
		toSerialize["config_managed"] = o.ConfigManaged
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.SubscriptionId != nil {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBindingRequest struct {
	value *RoleBindingRequest
	isSet bool
}

func (v NullableRoleBindingRequest) Get() *RoleBindingRequest {
	return v.value
}

func (v *NullableRoleBindingRequest) Set(val *RoleBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBindingRequest(val *RoleBindingRequest) *NullableRoleBindingRequest {
	return &NullableRoleBindingRequest{value: val, isSet: true}
}

func (v NullableRoleBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


