/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorListAllOf struct for ConnectorListAllOf
type ConnectorListAllOf struct {

	Items *[]Connector `json:"items,omitempty"`

}

// NewConnectorListAllOf instantiates a new ConnectorListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorListAllOf() *ConnectorListAllOf {
	this := ConnectorListAllOf{}
	return &this
}

// NewConnectorListAllOfWithDefaults instantiates a new ConnectorListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorListAllOfWithDefaults() *ConnectorListAllOf {
	this := ConnectorListAllOf{}


	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectorListAllOf) GetItems() []Connector {
	if o == nil || o.Items == nil {
		var ret []Connector
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorListAllOf) GetItemsOk() (*[]Connector, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectorListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Connector and assigns it to the Items field.
func (o *ConnectorListAllOf) SetItems(v []Connector) {
	o.Items = &v
}


func (o ConnectorListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorListAllOf struct {
	value *ConnectorListAllOf
	isSet bool
}

func (v NullableConnectorListAllOf) Get() *ConnectorListAllOf {
	return v.value
}

func (v *NullableConnectorListAllOf) Set(val *ConnectorListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorListAllOf(val *ConnectorListAllOf) *NullableConnectorListAllOf {
	return &NullableConnectorListAllOf{value: val, isSet: true}
}

func (v NullableConnectorListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

