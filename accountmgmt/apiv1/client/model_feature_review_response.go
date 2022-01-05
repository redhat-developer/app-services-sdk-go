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

// FeatureReviewResponse struct for FeatureReviewResponse
type FeatureReviewResponse struct {
	Enabled bool `json:"enabled"`
	FeatureId string `json:"feature_id"`
}

// NewFeatureReviewResponse instantiates a new FeatureReviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureReviewResponse(enabled bool, featureId string) *FeatureReviewResponse {
	this := FeatureReviewResponse{}
	this.Enabled = enabled
	this.FeatureId = featureId
	return &this
}

// NewFeatureReviewResponseWithDefaults instantiates a new FeatureReviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureReviewResponseWithDefaults() *FeatureReviewResponse {
	this := FeatureReviewResponse{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *FeatureReviewResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FeatureReviewResponse) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FeatureReviewResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFeatureId returns the FeatureId field value
func (o *FeatureReviewResponse) GetFeatureId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value
// and a boolean to check if the value has been set.
func (o *FeatureReviewResponse) GetFeatureIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FeatureId, true
}

// SetFeatureId sets field value
func (o *FeatureReviewResponse) SetFeatureId(v string) {
	o.FeatureId = v
}

func (o FeatureReviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["feature_id"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureReviewResponse struct {
	value *FeatureReviewResponse
	isSet bool
}

func (v NullableFeatureReviewResponse) Get() *FeatureReviewResponse {
	return v.value
}

func (v *NullableFeatureReviewResponse) Set(val *FeatureReviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureReviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureReviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureReviewResponse(val *FeatureReviewResponse) *NullableFeatureReviewResponse {
	return &NullableFeatureReviewResponse{value: val, isSet: true}
}

func (v NullableFeatureReviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureReviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


