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
)

// ClusterRegistrationResponse struct for ClusterRegistrationResponse
type ClusterRegistrationResponse struct {
	AccountId *string `json:"account_id,omitempty"`
	AuthorizationToken *string `json:"authorization_token,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	// Cluster Registration expiration in Unix time
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// NewClusterRegistrationResponse instantiates a new ClusterRegistrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRegistrationResponse() *ClusterRegistrationResponse {
	this := ClusterRegistrationResponse{}
	return &this
}

// NewClusterRegistrationResponseWithDefaults instantiates a new ClusterRegistrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRegistrationResponseWithDefaults() *ClusterRegistrationResponse {
	this := ClusterRegistrationResponse{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ClusterRegistrationResponse) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRegistrationResponse) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ClusterRegistrationResponse) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ClusterRegistrationResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAuthorizationToken returns the AuthorizationToken field value if set, zero value otherwise.
func (o *ClusterRegistrationResponse) GetAuthorizationToken() string {
	if o == nil || o.AuthorizationToken == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationToken
}

// GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRegistrationResponse) GetAuthorizationTokenOk() (*string, bool) {
	if o == nil || o.AuthorizationToken == nil {
		return nil, false
	}
	return o.AuthorizationToken, true
}

// HasAuthorizationToken returns a boolean if a field has been set.
func (o *ClusterRegistrationResponse) HasAuthorizationToken() bool {
	if o != nil && o.AuthorizationToken != nil {
		return true
	}

	return false
}

// SetAuthorizationToken gets a reference to the given string and assigns it to the AuthorizationToken field.
func (o *ClusterRegistrationResponse) SetAuthorizationToken(v string) {
	o.AuthorizationToken = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ClusterRegistrationResponse) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRegistrationResponse) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ClusterRegistrationResponse) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ClusterRegistrationResponse) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ClusterRegistrationResponse) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRegistrationResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ClusterRegistrationResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ClusterRegistrationResponse) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o ClusterRegistrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AuthorizationToken != nil {
		toSerialize["authorization_token"] = o.AuthorizationToken
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRegistrationResponse struct {
	value *ClusterRegistrationResponse
	isSet bool
}

func (v NullableClusterRegistrationResponse) Get() *ClusterRegistrationResponse {
	return v.value
}

func (v *NullableClusterRegistrationResponse) Set(val *ClusterRegistrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRegistrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRegistrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRegistrationResponse(val *ClusterRegistrationResponse) *NullableClusterRegistrationResponse {
	return &NullableClusterRegistrationResponse{value: val, isSet: true}
}

func (v NullableClusterRegistrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRegistrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

