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

// AccessTokenCfg struct for AccessTokenCfg
type AccessTokenCfg struct {
	Auths map[string]map[string]interface{} `json:"auths"`
}

// NewAccessTokenCfg instantiates a new AccessTokenCfg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenCfg(auths map[string]map[string]interface{}) *AccessTokenCfg {
	this := AccessTokenCfg{}
	this.Auths = auths
	return &this
}

// NewAccessTokenCfgWithDefaults instantiates a new AccessTokenCfg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenCfgWithDefaults() *AccessTokenCfg {
	this := AccessTokenCfg{}
	return &this
}

// GetAuths returns the Auths field value
func (o *AccessTokenCfg) GetAuths() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Auths
}

// GetAuthsOk returns a tuple with the Auths field value
// and a boolean to check if the value has been set.
func (o *AccessTokenCfg) GetAuthsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Auths, true
}

// SetAuths sets field value
func (o *AccessTokenCfg) SetAuths(v map[string]map[string]interface{}) {
	o.Auths = v
}

func (o AccessTokenCfg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auths"] = o.Auths
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTokenCfg struct {
	value *AccessTokenCfg
	isSet bool
}

func (v NullableAccessTokenCfg) Get() *AccessTokenCfg {
	return v.value
}

func (v *NullableAccessTokenCfg) Set(val *AccessTokenCfg) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenCfg) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenCfg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenCfg(val *AccessTokenCfg) *NullableAccessTokenCfg {
	return &NullableAccessTokenCfg{value: val, isSet: true}
}

func (v NullableAccessTokenCfg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenCfg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


