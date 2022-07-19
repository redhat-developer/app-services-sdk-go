/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
	"time"
)

// ConnectorNamespace A connector namespace
type ConnectorNamespace struct {
	Id string `json:"id"`
	Kind *string `json:"kind,omitempty"`
	Href *string `json:"href,omitempty"`
	Owner *string `json:"owner,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name string `json:"name"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	ResourceVersion int64 `json:"resource_version"`
	Quota *ConnectorNamespaceQuota `json:"quota,omitempty"`
	ClusterId string `json:"cluster_id"`
	// Namespace expiration timestamp in RFC 3339 format
	Expiration *string `json:"expiration,omitempty"`
	Tenant ConnectorNamespaceTenant `json:"tenant"`
	Status ConnectorNamespaceStatus `json:"status"`
}

// NewConnectorNamespace instantiates a new ConnectorNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespace(id string, name string, resourceVersion int64, clusterId string, tenant ConnectorNamespaceTenant, status ConnectorNamespaceStatus) *ConnectorNamespace {
	this := ConnectorNamespace{}
	this.Id = id
	this.Name = name
	this.ResourceVersion = resourceVersion
	this.ClusterId = clusterId
	this.Tenant = tenant
	this.Status = status
	return &this
}

// NewConnectorNamespaceWithDefaults instantiates a new ConnectorNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceWithDefaults() *ConnectorNamespace {
	this := ConnectorNamespace{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectorNamespace) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorNamespace) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConnectorNamespace) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConnectorNamespace) SetHref(v string) {
	o.Href = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ConnectorNamespace) SetOwner(v string) {
	o.Owner = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectorNamespace) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ConnectorNamespace) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value
func (o *ConnectorNamespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorNamespace) SetName(v string) {
	o.Name = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *ConnectorNamespace) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetResourceVersion returns the ResourceVersion field value
func (o *ConnectorNamespace) GetResourceVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetResourceVersionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceVersion, true
}

// SetResourceVersion sets field value
func (o *ConnectorNamespace) SetResourceVersion(v int64) {
	o.ResourceVersion = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetQuota() ConnectorNamespaceQuota {
	if o == nil || o.Quota == nil {
		var ret ConnectorNamespaceQuota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetQuotaOk() (*ConnectorNamespaceQuota, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasQuota() bool {
	if o != nil && o.Quota != nil {
		return true
	}

	return false
}

// SetQuota gets a reference to the given ConnectorNamespaceQuota and assigns it to the Quota field.
func (o *ConnectorNamespace) SetQuota(v ConnectorNamespaceQuota) {
	o.Quota = &v
}

// GetClusterId returns the ClusterId field value
func (o *ConnectorNamespace) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetClusterIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConnectorNamespace) SetClusterId(v string) {
	o.ClusterId = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ConnectorNamespace) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ConnectorNamespace) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *ConnectorNamespace) SetExpiration(v string) {
	o.Expiration = &v
}

// GetTenant returns the Tenant field value
func (o *ConnectorNamespace) GetTenant() ConnectorNamespaceTenant {
	if o == nil {
		var ret ConnectorNamespaceTenant
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetTenantOk() (*ConnectorNamespaceTenant, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *ConnectorNamespace) SetTenant(v ConnectorNamespaceTenant) {
	o.Tenant = v
}

// GetStatus returns the Status field value
func (o *ConnectorNamespace) GetStatus() ConnectorNamespaceStatus {
	if o == nil {
		var ret ConnectorNamespaceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespace) GetStatusOk() (*ConnectorNamespaceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConnectorNamespace) SetStatus(v ConnectorNamespaceStatus) {
	o.Status = v
}

func (o ConnectorNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if true {
		toSerialize["resource_version"] = o.ResourceVersion
	}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespace struct {
	value *ConnectorNamespace
	isSet bool
}

func (v NullableConnectorNamespace) Get() *ConnectorNamespace {
	return v.value
}

func (v *NullableConnectorNamespace) Set(val *ConnectorNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespace(val *ConnectorNamespace) *NullableConnectorNamespace {
	return &NullableConnectorNamespace{value: val, isSet: true}
}

func (v NullableConnectorNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


