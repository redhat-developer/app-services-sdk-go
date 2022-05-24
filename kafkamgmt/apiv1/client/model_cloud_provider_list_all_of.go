/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// CloudProviderListAllOf struct for CloudProviderListAllOf
type CloudProviderListAllOf struct {

	Items *[]CloudProvider `json:"items,omitempty"`

}

// NewCloudProviderListAllOf instantiates a new CloudProviderListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderListAllOf() *CloudProviderListAllOf {
	this := CloudProviderListAllOf{}
	return &this
}

// NewCloudProviderListAllOfWithDefaults instantiates a new CloudProviderListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderListAllOfWithDefaults() *CloudProviderListAllOf {
	this := CloudProviderListAllOf{}


	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *CloudProviderListAllOf) GetItems() []CloudProvider {
	if o == nil || o.Items == nil {
		var ret []CloudProvider
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderListAllOf) GetItemsOk() (*[]CloudProvider, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CloudProviderListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CloudProvider and assigns it to the Items field.
func (o *CloudProviderListAllOf) SetItems(v []CloudProvider) {
	o.Items = &v
}


func (o CloudProviderListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableCloudProviderListAllOf struct {
	value *CloudProviderListAllOf
	isSet bool
}

func (v NullableCloudProviderListAllOf) Get() *CloudProviderListAllOf {
	return v.value
}

func (v *NullableCloudProviderListAllOf) Set(val *CloudProviderListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderListAllOf(val *CloudProviderListAllOf) *NullableCloudProviderListAllOf {
	return &NullableCloudProviderListAllOf{value: val, isSet: true}
}

func (v NullableCloudProviderListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

