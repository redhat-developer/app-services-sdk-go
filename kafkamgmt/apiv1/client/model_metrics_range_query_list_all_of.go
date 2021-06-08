/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances and connectors.
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// MetricsRangeQueryListAllOf struct for MetricsRangeQueryListAllOf
type MetricsRangeQueryListAllOf struct {

	Kind *string `json:"kind,omitempty"`

	Id *string `json:"id,omitempty"`

	Items *[]RangeQuery `json:"items,omitempty"`

}

// NewMetricsRangeQueryListAllOf instantiates a new MetricsRangeQueryListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRangeQueryListAllOf() *MetricsRangeQueryListAllOf {
	this := MetricsRangeQueryListAllOf{}
	return &this
}

// NewMetricsRangeQueryListAllOfWithDefaults instantiates a new MetricsRangeQueryListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRangeQueryListAllOfWithDefaults() *MetricsRangeQueryListAllOf {
	this := MetricsRangeQueryListAllOf{}




	return &this
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MetricsRangeQueryListAllOf) SetKind(v string) {
	o.Kind = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsRangeQueryListAllOf) SetId(v string) {
	o.Id = &v
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *MetricsRangeQueryListAllOf) GetItems() []RangeQuery {
	if o == nil || o.Items == nil {
		var ret []RangeQuery
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRangeQueryListAllOf) GetItemsOk() (*[]RangeQuery, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MetricsRangeQueryListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RangeQuery and assigns it to the Items field.
func (o *MetricsRangeQueryListAllOf) SetItems(v []RangeQuery) {
	o.Items = &v
}


func (o MetricsRangeQueryListAllOf) MarshalJSON() ([]byte, error) {
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

type NullableMetricsRangeQueryListAllOf struct {
	value *MetricsRangeQueryListAllOf
	isSet bool
}

func (v NullableMetricsRangeQueryListAllOf) Get() *MetricsRangeQueryListAllOf {
	return v.value
}

func (v *NullableMetricsRangeQueryListAllOf) Set(val *MetricsRangeQueryListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRangeQueryListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRangeQueryListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRangeQueryListAllOf(val *MetricsRangeQueryListAllOf) *NullableMetricsRangeQueryListAllOf {
	return &NullableMetricsRangeQueryListAllOf{value: val, isSet: true}
}

func (v NullableMetricsRangeQueryListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRangeQueryListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

