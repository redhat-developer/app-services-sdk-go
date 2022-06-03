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

// ConnectorClusterStatusStatus struct for ConnectorClusterStatusStatus
type ConnectorClusterStatusStatus struct {
	State *ConnectorClusterState `json:"state,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewConnectorClusterStatusStatus instantiates a new ConnectorClusterStatusStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorClusterStatusStatus() *ConnectorClusterStatusStatus {
	this := ConnectorClusterStatusStatus{}
	return &this
}

// NewConnectorClusterStatusStatusWithDefaults instantiates a new ConnectorClusterStatusStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterStatusStatusWithDefaults() *ConnectorClusterStatusStatus {
	this := ConnectorClusterStatusStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ConnectorClusterStatusStatus) GetState() ConnectorClusterState {
	if o == nil || o.State == nil {
		var ret ConnectorClusterState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterStatusStatus) GetStateOk() (*ConnectorClusterState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ConnectorClusterStatusStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ConnectorClusterState and assigns it to the State field.
func (o *ConnectorClusterStatusStatus) SetState(v ConnectorClusterState) {
	o.State = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConnectorClusterStatusStatus) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterStatusStatus) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConnectorClusterStatusStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ConnectorClusterStatusStatus) SetError(v string) {
	o.Error = &v
}

func (o ConnectorClusterStatusStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorClusterStatusStatus struct {
	value *ConnectorClusterStatusStatus
	isSet bool
}

func (v NullableConnectorClusterStatusStatus) Get() *ConnectorClusterStatusStatus {
	return v.value
}

func (v *NullableConnectorClusterStatusStatus) Set(val *ConnectorClusterStatusStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorClusterStatusStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorClusterStatusStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorClusterStatusStatus(val *ConnectorClusterStatusStatus) *NullableConnectorClusterStatusStatus {
	return &NullableConnectorClusterStatusStatus{value: val, isSet: true}
}

func (v NullableConnectorClusterStatusStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorClusterStatusStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


