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

// Topic Kafka Topic (A feed where records are stored and published)
type Topic struct {
	// The name of the topic.
	Name *string `json:"name,omitempty"`
	// Topic configuration entry.
	Config *[]ConfigEntry `json:"config,omitempty"`
	// Partitions for this topic.
	Partitions *[]Partition `json:"partitions,omitempty"`
}

// NewTopic instantiates a new Topic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopic() *Topic {
	this := Topic{}
	return &this
}

// NewTopicWithDefaults instantiates a new Topic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicWithDefaults() *Topic {
	this := Topic{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Topic) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Topic) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Topic) SetName(v string) {
	o.Name = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Topic) GetConfig() []ConfigEntry {
	if o == nil || o.Config == nil {
		var ret []ConfigEntry
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetConfigOk() (*[]ConfigEntry, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Topic) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []ConfigEntry and assigns it to the Config field.
func (o *Topic) SetConfig(v []ConfigEntry) {
	o.Config = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *Topic) GetPartitions() []Partition {
	if o == nil || o.Partitions == nil {
		var ret []Partition
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetPartitionsOk() (*[]Partition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *Topic) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []Partition and assigns it to the Partitions field.
func (o *Topic) SetPartitions(v []Partition) {
	o.Partitions = &v
}

func (o Topic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	return json.Marshal(toSerialize)
}

type NullableTopic struct {
	value *Topic
	isSet bool
}

func (v NullableTopic) Get() *Topic {
	return v.value
}

func (v *NullableTopic) Set(val *Topic) {
	v.value = val
	v.isSet = true
}

func (v NullableTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopic(val *Topic) *NullableTopic {
	return &NullableTopic{value: val, isSet: true}
}

func (v NullableTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
