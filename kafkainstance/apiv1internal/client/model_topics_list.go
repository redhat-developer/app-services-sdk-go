/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicsList A list of topics.
type TopicsList struct {
	// The page offset
	Offset int32 `json:"offset"`
	// number of entries per page
	Limit int32 `json:"limit"`
	// Total number of topics
	Count int32 `json:"count"`
	// List of topics
	Items []Topic `json:"items"`
}

// NewTopicsList instantiates a new TopicsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicsList(offset int32, limit int32, count int32, items []Topic) *TopicsList {
	this := TopicsList{}
	this.Offset = offset
	this.Limit = limit
	this.Count = count
	this.Items = items
	return &this
}

// NewTopicsListWithDefaults instantiates a new TopicsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicsListWithDefaults() *TopicsList {
	this := TopicsList{}
	return &this
}

// GetOffset returns the Offset field value
func (o *TopicsList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *TopicsList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *TopicsList) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *TopicsList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TopicsList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *TopicsList) SetLimit(v int32) {
	o.Limit = v
}

// GetCount returns the Count field value
func (o *TopicsList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TopicsList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *TopicsList) SetCount(v int32) {
	o.Count = v
}

// GetItems returns the Items field value
func (o *TopicsList) GetItems() []Topic {
	if o == nil {
		var ret []Topic
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TopicsList) GetItemsOk() (*[]Topic, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *TopicsList) SetItems(v []Topic) {
	o.Items = v
}

func (o TopicsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableTopicsList struct {
	value *TopicsList
	isSet bool
}

func (v NullableTopicsList) Get() *TopicsList {
	return v.value
}

func (v *NullableTopicsList) Set(val *TopicsList) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicsList) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicsList(val *TopicsList) *NullableTopicsList {
	return &NullableTopicsList{value: val, isSet: true}
}

func (v NullableTopicsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
