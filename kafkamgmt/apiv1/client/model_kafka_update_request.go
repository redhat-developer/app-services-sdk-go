/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// KafkaUpdateRequest struct for KafkaUpdateRequest
type KafkaUpdateRequest struct {

	Owner string `json:"owner"`

}

// NewKafkaUpdateRequest instantiates a new KafkaUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaUpdateRequest(owner string) *KafkaUpdateRequest {
	this := KafkaUpdateRequest{}
	this.Owner = owner
	return &this
}

// NewKafkaUpdateRequestWithDefaults instantiates a new KafkaUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaUpdateRequestWithDefaults() *KafkaUpdateRequest {
	this := KafkaUpdateRequest{}


	return &this
}


// GetOwner returns the Owner field value
func (o *KafkaUpdateRequest) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *KafkaUpdateRequest) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *KafkaUpdateRequest) SetOwner(v string) {
	o.Owner = v
}


func (o KafkaUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["owner"] = o.Owner
	}
    
	return json.Marshal(toSerialize)
}

type NullableKafkaUpdateRequest struct {
	value *KafkaUpdateRequest
	isSet bool
}

func (v NullableKafkaUpdateRequest) Get() *KafkaUpdateRequest {
	return v.value
}

func (v *NullableKafkaUpdateRequest) Set(val *KafkaUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaUpdateRequest(val *KafkaUpdateRequest) *NullableKafkaUpdateRequest {
	return &NullableKafkaUpdateRequest{value: val, isSet: true}
}

func (v NullableKafkaUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
