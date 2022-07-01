/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// RegionCapacityListItem schema for a kafka instance type capacity in region
type RegionCapacityListItem struct {
	// kafka instance type
	InstanceType string `json:"instance_type"`
	// list of available Kafka instance sizes that can be created in this region when taking account current capacity and regional limits
	AvailableSizes []string `json:"available_sizes"`
}

// NewRegionCapacityListItem instantiates a new RegionCapacityListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionCapacityListItem(instanceType string, availableSizes []string) *RegionCapacityListItem {
	this := RegionCapacityListItem{}
	this.InstanceType = instanceType
	this.AvailableSizes = availableSizes
	return &this
}

// NewRegionCapacityListItemWithDefaults instantiates a new RegionCapacityListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionCapacityListItemWithDefaults() *RegionCapacityListItem {
	this := RegionCapacityListItem{}
	return &this
}

// GetInstanceType returns the InstanceType field value
func (o *RegionCapacityListItem) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *RegionCapacityListItem) GetInstanceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *RegionCapacityListItem) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetAvailableSizes returns the AvailableSizes field value
func (o *RegionCapacityListItem) GetAvailableSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableSizes
}

// GetAvailableSizesOk returns a tuple with the AvailableSizes field value
// and a boolean to check if the value has been set.
func (o *RegionCapacityListItem) GetAvailableSizesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableSizes, true
}

// SetAvailableSizes sets field value
func (o *RegionCapacityListItem) SetAvailableSizes(v []string) {
	o.AvailableSizes = v
}

func (o RegionCapacityListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instance_type"] = o.InstanceType
	}
	if true {
		toSerialize["available_sizes"] = o.AvailableSizes
	}
	return json.Marshal(toSerialize)
}

type NullableRegionCapacityListItem struct {
	value *RegionCapacityListItem
	isSet bool
}

func (v NullableRegionCapacityListItem) Get() *RegionCapacityListItem {
	return v.value
}

func (v *NullableRegionCapacityListItem) Set(val *RegionCapacityListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionCapacityListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionCapacityListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionCapacityListItem(val *RegionCapacityListItem) *NullableRegionCapacityListItem {
	return &NullableRegionCapacityListItem{value: val, isSet: true}
}

func (v NullableRegionCapacityListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionCapacityListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


