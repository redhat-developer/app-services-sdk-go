/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// KafkaConnectionSettings Holds the configuration to connect to a Kafka Instance.
type KafkaConnectionSettings struct {
	Id string `json:"id"`
	Url string `json:"url"`
}

// NewKafkaConnectionSettings instantiates a new KafkaConnectionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaConnectionSettings(id string, url string) *KafkaConnectionSettings {
	this := KafkaConnectionSettings{}
	this.Id = id
	this.Url = url
	return &this
}

// NewKafkaConnectionSettingsWithDefaults instantiates a new KafkaConnectionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaConnectionSettingsWithDefaults() *KafkaConnectionSettings {
	this := KafkaConnectionSettings{}
	return &this
}

// GetId returns the Id field value
func (o *KafkaConnectionSettings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KafkaConnectionSettings) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KafkaConnectionSettings) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *KafkaConnectionSettings) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *KafkaConnectionSettings) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *KafkaConnectionSettings) SetUrl(v string) {
	o.Url = v
}

func (o KafkaConnectionSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaConnectionSettings struct {
	value *KafkaConnectionSettings
	isSet bool
}

func (v NullableKafkaConnectionSettings) Get() *KafkaConnectionSettings {
	return v.value
}

func (v *NullableKafkaConnectionSettings) Set(val *KafkaConnectionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaConnectionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaConnectionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaConnectionSettings(val *KafkaConnectionSettings) *NullableKafkaConnectionSettings {
	return &NullableKafkaConnectionSettings{value: val, isSet: true}
}

func (v NullableKafkaConnectionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaConnectionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


