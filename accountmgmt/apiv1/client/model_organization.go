/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmt

import (
	"encoding/json"
	"time"
)

// Organization struct for Organization
type Organization struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Capabilities *[]Capability `json:"capabilities,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	EbsAccountId *string `json:"ebs_account_id,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Labels *[]Label `json:"labels,omitempty"`
	Name *string `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Organization) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Organization) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Organization) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Organization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Organization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Organization) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Organization) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Organization) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Organization) SetKind(v string) {
	o.Kind = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Organization) GetCapabilities() []Capability {
	if o == nil || o.Capabilities == nil {
		var ret []Capability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCapabilitiesOk() (*[]Capability, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Organization) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []Capability and assigns it to the Capabilities field.
func (o *Organization) SetCapabilities(v []Capability) {
	o.Capabilities = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEbsAccountId returns the EbsAccountId field value if set, zero value otherwise.
func (o *Organization) GetEbsAccountId() string {
	if o == nil || o.EbsAccountId == nil {
		var ret string
		return ret
	}
	return *o.EbsAccountId
}

// GetEbsAccountIdOk returns a tuple with the EbsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetEbsAccountIdOk() (*string, bool) {
	if o == nil || o.EbsAccountId == nil {
		return nil, false
	}
	return o.EbsAccountId, true
}

// HasEbsAccountId returns a boolean if a field has been set.
func (o *Organization) HasEbsAccountId() bool {
	if o != nil && o.EbsAccountId != nil {
		return true
	}

	return false
}

// SetEbsAccountId gets a reference to the given string and assigns it to the EbsAccountId field.
func (o *Organization) SetEbsAccountId(v string) {
	o.EbsAccountId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Organization) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Organization) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Organization) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Organization) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Organization) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *Organization) SetLabels(v []Label) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organization) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organization) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organization) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Organization) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Organization) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Organization) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.EbsAccountId != nil {
		toSerialize["ebs_account_id"] = o.EbsAccountId
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


