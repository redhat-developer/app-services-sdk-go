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
	"time"
)

// ConnectorNamespaceMeta struct for ConnectorNamespaceMeta
type ConnectorNamespaceMeta struct {
	Owner *string `json:"owner,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
	Name *string `json:"name,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	ResourceVersion *int64 `json:"resource_version,omitempty"`
	Quota *ConnectorNamespaceQuota `json:"quota,omitempty"`
}

// NewConnectorNamespaceMeta instantiates a new ConnectorNamespaceMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceMeta() *ConnectorNamespaceMeta {
	this := ConnectorNamespaceMeta{}
	return &this
}

// NewConnectorNamespaceMetaWithDefaults instantiates a new ConnectorNamespaceMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceMetaWithDefaults() *ConnectorNamespaceMeta {
	this := ConnectorNamespaceMeta{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ConnectorNamespaceMeta) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectorNamespaceMeta) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ConnectorNamespaceMeta) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorNamespaceMeta) SetName(v string) {
	o.Name = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorNamespaceMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetResourceVersion() int64 {
	if o == nil || o.ResourceVersion == nil {
		var ret int64
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetResourceVersionOk() (*int64, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given int64 and assigns it to the ResourceVersion field.
func (o *ConnectorNamespaceMeta) SetResourceVersion(v int64) {
	o.ResourceVersion = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *ConnectorNamespaceMeta) GetQuota() ConnectorNamespaceQuota {
	if o == nil || o.Quota == nil {
		var ret ConnectorNamespaceQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceMeta) GetQuotaOk() (*ConnectorNamespaceQuota, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *ConnectorNamespaceMeta) HasQuota() bool {
	if o != nil && o.Quota != nil {
		return true
	}

	return false
}

// SetQuota gets a reference to the given ConnectorNamespaceQuota and assigns it to the Quota field.
func (o *ConnectorNamespaceMeta) SetQuota(v ConnectorNamespaceQuota) {
	o.Quota = &v
}

func (o ConnectorNamespaceMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.ResourceVersion != nil {
		toSerialize["resource_version"] = o.ResourceVersion
	}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceMeta struct {
	value *ConnectorNamespaceMeta
	isSet bool
}

func (v NullableConnectorNamespaceMeta) Get() *ConnectorNamespaceMeta {
	return v.value
}

func (v *NullableConnectorNamespaceMeta) Set(val *ConnectorNamespaceMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceMeta(val *ConnectorNamespaceMeta) *NullableConnectorNamespaceMeta {
	return &NullableConnectorNamespaceMeta{value: val, isSet: true}
}

func (v NullableConnectorNamespaceMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


