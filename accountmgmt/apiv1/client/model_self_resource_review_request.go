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

// SelfResourceReviewRequest struct for SelfResourceReviewRequest
type SelfResourceReviewRequest struct {
	Action *string `json:"action,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
}

// NewSelfResourceReviewRequest instantiates a new SelfResourceReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfResourceReviewRequest() *SelfResourceReviewRequest {
	this := SelfResourceReviewRequest{}
	return &this
}

// NewSelfResourceReviewRequestWithDefaults instantiates a new SelfResourceReviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfResourceReviewRequestWithDefaults() *SelfResourceReviewRequest {
	this := SelfResourceReviewRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SelfResourceReviewRequest) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfResourceReviewRequest) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SelfResourceReviewRequest) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SelfResourceReviewRequest) SetAction(v string) {
	o.Action = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *SelfResourceReviewRequest) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfResourceReviewRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *SelfResourceReviewRequest) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *SelfResourceReviewRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o SelfResourceReviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	return json.Marshal(toSerialize)
}

type NullableSelfResourceReviewRequest struct {
	value *SelfResourceReviewRequest
	isSet bool
}

func (v NullableSelfResourceReviewRequest) Get() *SelfResourceReviewRequest {
	return v.value
}

func (v *NullableSelfResourceReviewRequest) Set(val *SelfResourceReviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfResourceReviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfResourceReviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfResourceReviewRequest(val *SelfResourceReviewRequest) *NullableSelfResourceReviewRequest {
	return &NullableSelfResourceReviewRequest{value: val, isSet: true}
}

func (v NullableSelfResourceReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfResourceReviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


