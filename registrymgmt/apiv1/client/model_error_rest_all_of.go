/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.4
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// ErrorRestAllOf struct for ErrorRestAllOf
type ErrorRestAllOf struct {

	Code *string `json:"code,omitempty"`

	Reason *string `json:"reason,omitempty"`

	OperationId *string `json:"operation_id,omitempty"`

}

// NewErrorRestAllOf instantiates a new ErrorRestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorRestAllOf() *ErrorRestAllOf {
	this := ErrorRestAllOf{}
	return &this
}

// NewErrorRestAllOfWithDefaults instantiates a new ErrorRestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorRestAllOfWithDefaults() *ErrorRestAllOf {
	this := ErrorRestAllOf{}




	return &this
}


// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorRestAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRestAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorRestAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorRestAllOf) SetCode(v string) {
	o.Code = &v
}


// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorRestAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRestAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorRestAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorRestAllOf) SetReason(v string) {
	o.Reason = &v
}


// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *ErrorRestAllOf) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRestAllOf) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *ErrorRestAllOf) HasOperationId() bool {
	if o != nil && o.OperationId != nil {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *ErrorRestAllOf) SetOperationId(v string) {
	o.OperationId = &v
}


func (o ErrorRestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
    
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
    
	if o.OperationId != nil {
		toSerialize["operation_id"] = o.OperationId
	}
    
	return json.Marshal(toSerialize)
}

type NullableErrorRestAllOf struct {
	value *ErrorRestAllOf
	isSet bool
}

func (v NullableErrorRestAllOf) Get() *ErrorRestAllOf {
	return v.value
}

func (v *NullableErrorRestAllOf) Set(val *ErrorRestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorRestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorRestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorRestAllOf(val *ErrorRestAllOf) *NullableErrorRestAllOf {
	return &NullableErrorRestAllOf{value: val, isSet: true}
}

func (v NullableErrorRestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorRestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
