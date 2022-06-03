/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.8.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// MetricsRangeQueryList struct for MetricsRangeQueryList
type MetricsRangeQueryList struct {

	Kind *string `json:"kind,omitempty"`

	Id *string `json:"id,omitempty"`

	Items *[]RangeQuery `json:"items,omitempty"`

}

// NewMetricsRangeQueryList instantiates a new MetricsRangeQueryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRangeQueryList() *MetricsRangeQueryList {
	this := MetricsRangeQueryList{}
	return &this
}

// NewMetricsRangeQueryListWithDefaults instantiates a new MetricsRangeQueryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRangeQueryListWithDefaults() *MetricsRangeQueryList {
	this := MetricsRangeQueryList{}




	return &this
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MetricsRangeQueryList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MetricsRangeQueryList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MetricsRangeQueryList) SetKind(v string) {
	o.Kind = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsRangeQueryList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsRangeQueryList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsRangeQueryList) SetId(v string) {
	o.Id = &v
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *MetricsRangeQueryList) GetItems() []RangeQuery {
	if o == nil || o.Items == nil {
		var ret []RangeQuery
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryList) GetItemsOk() (*[]RangeQuery, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MetricsRangeQueryList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RangeQuery and assigns it to the Items field.
func (o *MetricsRangeQueryList) SetItems(v []RangeQuery) {
	o.Items = &v
}


func (o MetricsRangeQueryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
    
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableMetricsRangeQueryList struct {
	value *MetricsRangeQueryList
	isSet bool
}

func (v NullableMetricsRangeQueryList) Get() *MetricsRangeQueryList {
	return v.value
}

func (v *NullableMetricsRangeQueryList) Set(val *MetricsRangeQueryList) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRangeQueryList) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRangeQueryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRangeQueryList(val *MetricsRangeQueryList) *NullableMetricsRangeQueryList {
	return &NullableMetricsRangeQueryList{value: val, isSet: true}
}

func (v NullableMetricsRangeQueryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRangeQueryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

