/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicList A list of topics.
type TopicList struct {
	Items *[]Topic `json:"items,omitempty"`
}

// NewTopicList instantiates a new TopicList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicList() *TopicList {
	this := TopicList{}
	return &this
}

// NewTopicListWithDefaults instantiates a new TopicList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicListWithDefaults() *TopicList {
	this := TopicList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TopicList) GetItems() []Topic {
	if o == nil || o.Items == nil {
		var ret []Topic
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicList) GetItemsOk() (*[]Topic, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TopicList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Topic and assigns it to the Items field.
func (o *TopicList) SetItems(v []Topic) {
	o.Items = &v
}

func (o TopicList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableTopicList struct {
	value *TopicList
	isSet bool
}

func (v NullableTopicList) Get() *TopicList {
	return v.value
}

func (v *NullableTopicList) Set(val *TopicList) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicList) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicList(val *TopicList) *NullableTopicList {
	return &NullableTopicList{value: val, isSet: true}
}

func (v NullableTopicList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

