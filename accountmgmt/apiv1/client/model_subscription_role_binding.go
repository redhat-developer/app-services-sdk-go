/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// SubscriptionRoleBinding struct for SubscriptionRoleBinding
type SubscriptionRoleBinding struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Account *AccountReference `json:"account,omitempty"`
	AccountEmail *string `json:"account_email,omitempty"`
	AccountUsername *string `json:"account_username,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Role *ObjectReference `json:"role,omitempty"`
	Subscription *ObjectReference `json:"subscription,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSubscriptionRoleBinding instantiates a new SubscriptionRoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRoleBinding() *SubscriptionRoleBinding {
	this := SubscriptionRoleBinding{}
	return &this
}

// NewSubscriptionRoleBindingWithDefaults instantiates a new SubscriptionRoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRoleBindingWithDefaults() *SubscriptionRoleBinding {
	this := SubscriptionRoleBinding{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SubscriptionRoleBinding) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionRoleBinding) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SubscriptionRoleBinding) SetKind(v string) {
	o.Kind = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetAccount() AccountReference {
	if o == nil || o.Account == nil {
		var ret AccountReference
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetAccountOk() (*AccountReference, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountReference and assigns it to the Account field.
func (o *SubscriptionRoleBinding) SetAccount(v AccountReference) {
	o.Account = &v
}

// GetAccountEmail returns the AccountEmail field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetAccountEmail() string {
	if o == nil || o.AccountEmail == nil {
		var ret string
		return ret
	}
	return *o.AccountEmail
}

// GetAccountEmailOk returns a tuple with the AccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetAccountEmailOk() (*string, bool) {
	if o == nil || o.AccountEmail == nil {
		return nil, false
	}
	return o.AccountEmail, true
}

// HasAccountEmail returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasAccountEmail() bool {
	if o != nil && o.AccountEmail != nil {
		return true
	}

	return false
}

// SetAccountEmail gets a reference to the given string and assigns it to the AccountEmail field.
func (o *SubscriptionRoleBinding) SetAccountEmail(v string) {
	o.AccountEmail = &v
}

// GetAccountUsername returns the AccountUsername field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetAccountUsername() string {
	if o == nil || o.AccountUsername == nil {
		var ret string
		return ret
	}
	return *o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetAccountUsernameOk() (*string, bool) {
	if o == nil || o.AccountUsername == nil {
		return nil, false
	}
	return o.AccountUsername, true
}

// HasAccountUsername returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasAccountUsername() bool {
	if o != nil && o.AccountUsername != nil {
		return true
	}

	return false
}

// SetAccountUsername gets a reference to the given string and assigns it to the AccountUsername field.
func (o *SubscriptionRoleBinding) SetAccountUsername(v string) {
	o.AccountUsername = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SubscriptionRoleBinding) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetRole() ObjectReference {
	if o == nil || o.Role == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetRoleOk() (*ObjectReference, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given ObjectReference and assigns it to the Role field.
func (o *SubscriptionRoleBinding) SetRole(v ObjectReference) {
	o.Role = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetSubscription() ObjectReference {
	if o == nil || o.Subscription == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetSubscriptionOk() (*ObjectReference, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ObjectReference and assigns it to the Subscription field.
func (o *SubscriptionRoleBinding) SetSubscription(v ObjectReference) {
	o.Subscription = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SubscriptionRoleBinding) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRoleBinding) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubscriptionRoleBinding) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SubscriptionRoleBinding) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SubscriptionRoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.AccountEmail != nil {
		toSerialize["account_email"] = o.AccountEmail
	}
	if o.AccountUsername != nil {
		toSerialize["account_username"] = o.AccountUsername
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionRoleBinding struct {
	value *SubscriptionRoleBinding
	isSet bool
}

func (v NullableSubscriptionRoleBinding) Get() *SubscriptionRoleBinding {
	return v.value
}

func (v *NullableSubscriptionRoleBinding) Set(val *SubscriptionRoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRoleBinding(val *SubscriptionRoleBinding) *NullableSubscriptionRoleBinding {
	return &NullableSubscriptionRoleBinding{value: val, isSet: true}
}

func (v NullableSubscriptionRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


