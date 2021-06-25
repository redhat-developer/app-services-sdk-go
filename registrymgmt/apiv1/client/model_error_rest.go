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

// ErrorRest struct for ErrorRest
type ErrorRest struct {

	Id *string `json:"id,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

	Code *string `json:"code,omitempty"`

	Reason *string `json:"reason,omitempty"`

	OperationId *string `json:"operation_id,omitempty"`

}

// NewErrorRest instantiates a new ErrorRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorRest() *ErrorRest {
	this := ErrorRest{}
	return &this
}

// NewErrorRestWithDefaults instantiates a new ErrorRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorRestWithDefaults() *ErrorRest {
	this := ErrorRest{}







	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *ErrorRest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ErrorRest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ErrorRest) SetId(v string) {
	o.Id = &v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ErrorRest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ErrorRest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ErrorRest) SetKind(v string) {
	o.Kind = &v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *ErrorRest) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ErrorRest) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ErrorRest) SetHref(v string) {
	o.Href = &v
}


// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorRest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorRest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorRest) SetCode(v string) {
	o.Code = &v
}


// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorRest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorRest) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorRest) SetReason(v string) {
	o.Reason = &v
}


// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *ErrorRest) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorRest) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *ErrorRest) HasOperationId() bool {
	if o != nil && o.OperationId != nil {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *ErrorRest) SetOperationId(v string) {
	o.OperationId = &v
}


func (o ErrorRest) MarshalJSON() ([]byte, error) {
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

type NullableErrorRest struct {
	value *ErrorRest
	isSet bool
}

func (v NullableErrorRest) Get() *ErrorRest {
	return v.value
}

func (v *NullableErrorRest) Set(val *ErrorRest) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorRest) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorRest(val *ErrorRest) *NullableErrorRest {
	return &NullableErrorRest{value: val, isSet: true}
}

func (v NullableErrorRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

