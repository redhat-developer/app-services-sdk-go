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

// FeatureReview struct for FeatureReview
type FeatureReview struct {
	AccountUsername string `json:"account_username"`
	Feature string `json:"feature"`
}

// NewFeatureReview instantiates a new FeatureReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureReview(accountUsername string, feature string) *FeatureReview {
	this := FeatureReview{}
	this.AccountUsername = accountUsername
	this.Feature = feature
	return &this
}

// NewFeatureReviewWithDefaults instantiates a new FeatureReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureReviewWithDefaults() *FeatureReview {
	this := FeatureReview{}
	return &this
}

// GetAccountUsername returns the AccountUsername field value
func (o *FeatureReview) GetAccountUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value
// and a boolean to check if the value has been set.
func (o *FeatureReview) GetAccountUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountUsername, true
}

// SetAccountUsername sets field value
func (o *FeatureReview) SetAccountUsername(v string) {
	o.AccountUsername = v
}

// GetFeature returns the Feature field value
func (o *FeatureReview) GetFeature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value
// and a boolean to check if the value has been set.
func (o *FeatureReview) GetFeatureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Feature, true
}

// SetFeature sets field value
func (o *FeatureReview) SetFeature(v string) {
	o.Feature = v
}

func (o FeatureReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_username"] = o.AccountUsername
	}
	if true {
		toSerialize["feature"] = o.Feature
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureReview struct {
	value *FeatureReview
	isSet bool
}

func (v NullableFeatureReview) Get() *FeatureReview {
	return v.value
}

func (v *NullableFeatureReview) Set(val *FeatureReview) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureReview) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureReview(val *FeatureReview) *NullableFeatureReview {
	return &NullableFeatureReview{value: val, isSet: true}
}

func (v NullableFeatureReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


