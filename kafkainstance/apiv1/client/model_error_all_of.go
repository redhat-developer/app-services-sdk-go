/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.12.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ErrorAllOf General error response
type ErrorAllOf struct {
	// General reason for the error. Does not change between specific occurrences.
	Reason *string `json:"reason,omitempty"`
	// Detail specific to an error occurrence. May be different depending on the condition(s) that trigger the error.
	Detail *string `json:"detail,omitempty"`
	Code *int32 `json:"code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Class *string `json:"class,omitempty"`
}

// NewErrorAllOf instantiates a new ErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorAllOf() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// NewErrorAllOfWithDefaults instantiates a new ErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorAllOfWithDefaults() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ErrorAllOf) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorAllOf) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ErrorAllOf) SetDetail(v string) {
	o.Detail = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorAllOf) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ErrorAllOf) SetCode(v int32) {
	o.Code = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ErrorAllOf) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ErrorAllOf) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ErrorAllOf) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *ErrorAllOf) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *ErrorAllOf) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *ErrorAllOf) SetClass(v string) {
	o.Class = &v
}

func (o ErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	return json.Marshal(toSerialize)
}

type NullableErrorAllOf struct {
	value *ErrorAllOf
	isSet bool
}

func (v NullableErrorAllOf) Get() *ErrorAllOf {
	return v.value
}

func (v *NullableErrorAllOf) Set(val *ErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorAllOf(val *ErrorAllOf) *NullableErrorAllOf {
	return &NullableErrorAllOf{value: val, isSet: true}
}

func (v NullableErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


