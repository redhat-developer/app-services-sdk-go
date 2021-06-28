/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.5
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// ErrorListRest struct for ErrorListRest
type ErrorListRest struct {

	Kind string `json:"kind"`

	Page int32 `json:"page"`

	Size int32 `json:"size"`

	Total int32 `json:"total"`

	Items []ErrorRest `json:"items"`

}

// NewErrorListRest instantiates a new ErrorListRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorListRest(kind string, page int32, size int32, total int32, items []ErrorRest) *ErrorListRest {
	this := ErrorListRest{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewErrorListRestWithDefaults instantiates a new ErrorListRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorListRestWithDefaults() *ErrorListRest {
	this := ErrorListRest{}






	return &this
}


// GetKind returns the Kind field value
func (o *ErrorListRest) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ErrorListRest) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ErrorListRest) SetKind(v string) {
	o.Kind = v
}


// GetPage returns the Page field value
func (o *ErrorListRest) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ErrorListRest) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ErrorListRest) SetPage(v int32) {
	o.Page = v
}


// GetSize returns the Size field value
func (o *ErrorListRest) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ErrorListRest) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ErrorListRest) SetSize(v int32) {
	o.Size = v
}


// GetTotal returns the Total field value
func (o *ErrorListRest) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ErrorListRest) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ErrorListRest) SetTotal(v int32) {
	o.Total = v
}


// GetItems returns the Items field value
func (o *ErrorListRest) GetItems() []ErrorRest {
	if o == nil {
		var ret []ErrorRest
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ErrorListRest) GetItemsOk() (*[]ErrorRest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ErrorListRest) SetItems(v []ErrorRest) {
	o.Items = v
}


func (o ErrorListRest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["kind"] = o.Kind
	}
    
	if true {
		toSerialize["page"] = o.Page
	}
    
	if true {
		toSerialize["size"] = o.Size
	}
    
	if true {
		toSerialize["total"] = o.Total
	}
    
	if true {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableErrorListRest struct {
	value *ErrorListRest
	isSet bool
}

func (v NullableErrorListRest) Get() *ErrorListRest {
	return v.value
}

func (v *NullableErrorListRest) Set(val *ErrorListRest) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorListRest) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorListRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorListRest(val *ErrorListRest) *NullableErrorListRest {
	return &NullableErrorListRest{value: val, isSet: true}
}

func (v NullableErrorListRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorListRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

