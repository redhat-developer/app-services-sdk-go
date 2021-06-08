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

// MetricsInstantQueryList struct for MetricsInstantQueryList
type MetricsInstantQueryList struct {

	Kind *string `json:"kind,omitempty"`

	Id *string `json:"id,omitempty"`

	Items *[]InstantQuery `json:"items,omitempty"`

}

// NewMetricsInstantQueryList instantiates a new MetricsInstantQueryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsInstantQueryList() *MetricsInstantQueryList {
	this := MetricsInstantQueryList{}
	return &this
}

// NewMetricsInstantQueryListWithDefaults instantiates a new MetricsInstantQueryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsInstantQueryListWithDefaults() *MetricsInstantQueryList {
	this := MetricsInstantQueryList{}




	return &this
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MetricsInstantQueryList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInstantQueryList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MetricsInstantQueryList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MetricsInstantQueryList) SetKind(v string) {
	o.Kind = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsInstantQueryList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInstantQueryList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsInstantQueryList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsInstantQueryList) SetId(v string) {
	o.Id = &v
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *MetricsInstantQueryList) GetItems() []InstantQuery {
	if o == nil || o.Items == nil {
		var ret []InstantQuery
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInstantQueryList) GetItemsOk() (*[]InstantQuery, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MetricsInstantQueryList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InstantQuery and assigns it to the Items field.
func (o *MetricsInstantQueryList) SetItems(v []InstantQuery) {
	o.Items = &v
}


func (o MetricsInstantQueryList) MarshalJSON() ([]byte, error) {
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

type NullableMetricsInstantQueryList struct {
	value *MetricsInstantQueryList
	isSet bool
}

func (v NullableMetricsInstantQueryList) Get() *MetricsInstantQueryList {
	return v.value
}

func (v *NullableMetricsInstantQueryList) Set(val *MetricsInstantQueryList) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsInstantQueryList) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsInstantQueryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsInstantQueryList(val *MetricsInstantQueryList) *NullableMetricsInstantQueryList {
	return &NullableMetricsInstantQueryList{value: val, isSet: true}
}

func (v NullableMetricsInstantQueryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsInstantQueryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

