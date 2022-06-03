/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorClusterListAllOf struct for ConnectorClusterListAllOf
type ConnectorClusterListAllOf struct {
	Items *[]ConnectorCluster `json:"items,omitempty"`
}

// NewConnectorClusterListAllOf instantiates a new ConnectorClusterListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorClusterListAllOf() *ConnectorClusterListAllOf {
	this := ConnectorClusterListAllOf{}
	return &this
}

// NewConnectorClusterListAllOfWithDefaults instantiates a new ConnectorClusterListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterListAllOfWithDefaults() *ConnectorClusterListAllOf {
	this := ConnectorClusterListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectorClusterListAllOf) GetItems() []ConnectorCluster {
	if o == nil || o.Items == nil {
		var ret []ConnectorCluster
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterListAllOf) GetItemsOk() (*[]ConnectorCluster, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectorClusterListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConnectorCluster and assigns it to the Items field.
func (o *ConnectorClusterListAllOf) SetItems(v []ConnectorCluster) {
	o.Items = &v
}

func (o ConnectorClusterListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorClusterListAllOf struct {
	value *ConnectorClusterListAllOf
	isSet bool
}

func (v NullableConnectorClusterListAllOf) Get() *ConnectorClusterListAllOf {
	return v.value
}

func (v *NullableConnectorClusterListAllOf) Set(val *ConnectorClusterListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorClusterListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorClusterListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorClusterListAllOf(val *ConnectorClusterListAllOf) *NullableConnectorClusterListAllOf {
	return &NullableConnectorClusterListAllOf{value: val, isSet: true}
}

func (v NullableConnectorClusterListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorClusterListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


