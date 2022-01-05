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

// ReservedResourceAllOf struct for ReservedResourceAllOf
type ReservedResourceAllOf struct {
	AvailabilityZoneType *string `json:"availability_zone_type,omitempty"`
	BillingModel *string `json:"billing_model,omitempty"`
	Byoc bool `json:"byoc"`
	Cluster *bool `json:"cluster,omitempty"`
	Count *int32 `json:"count,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Subscription *ObjectReference `json:"subscription,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewReservedResourceAllOf instantiates a new ReservedResourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedResourceAllOf(byoc bool) *ReservedResourceAllOf {
	this := ReservedResourceAllOf{}
	this.Byoc = byoc
	return &this
}

// NewReservedResourceAllOfWithDefaults instantiates a new ReservedResourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedResourceAllOfWithDefaults() *ReservedResourceAllOf {
	this := ReservedResourceAllOf{}
	return &this
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetAvailabilityZoneType() string {
	if o == nil || o.AvailabilityZoneType == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil || o.AvailabilityZoneType == nil {
		return nil, false
	}
	return o.AvailabilityZoneType, true
}

// HasAvailabilityZoneType returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasAvailabilityZoneType() bool {
	if o != nil && o.AvailabilityZoneType != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneType gets a reference to the given string and assigns it to the AvailabilityZoneType field.
func (o *ReservedResourceAllOf) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = &v
}

// GetBillingModel returns the BillingModel field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetBillingModel() string {
	if o == nil || o.BillingModel == nil {
		var ret string
		return ret
	}
	return *o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetBillingModelOk() (*string, bool) {
	if o == nil || o.BillingModel == nil {
		return nil, false
	}
	return o.BillingModel, true
}

// HasBillingModel returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasBillingModel() bool {
	if o != nil && o.BillingModel != nil {
		return true
	}

	return false
}

// SetBillingModel gets a reference to the given string and assigns it to the BillingModel field.
func (o *ReservedResourceAllOf) SetBillingModel(v string) {
	o.BillingModel = &v
}

// GetByoc returns the Byoc field value
func (o *ReservedResourceAllOf) GetByoc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetByocOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Byoc, true
}

// SetByoc sets field value
func (o *ReservedResourceAllOf) SetByoc(v bool) {
	o.Byoc = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCluster() bool {
	if o == nil || o.Cluster == nil {
		var ret bool
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetClusterOk() (*bool, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given bool and assigns it to the Cluster field.
func (o *ReservedResourceAllOf) SetCluster(v bool) {
	o.Cluster = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ReservedResourceAllOf) SetCount(v int32) {
	o.Count = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ReservedResourceAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ReservedResourceAllOf) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ReservedResourceAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetSubscription() ObjectReference {
	if o == nil || o.Subscription == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetSubscriptionOk() (*ObjectReference, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ObjectReference and assigns it to the Subscription field.
func (o *ReservedResourceAllOf) SetSubscription(v ObjectReference) {
	o.Subscription = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReservedResourceAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedResourceAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReservedResourceAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ReservedResourceAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ReservedResourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityZoneType != nil {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if o.BillingModel != nil {
		toSerialize["billing_model"] = o.BillingModel
	}
	if true {
		toSerialize["byoc"] = o.Byoc
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableReservedResourceAllOf struct {
	value *ReservedResourceAllOf
	isSet bool
}

func (v NullableReservedResourceAllOf) Get() *ReservedResourceAllOf {
	return v.value
}

func (v *NullableReservedResourceAllOf) Set(val *ReservedResourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedResourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedResourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedResourceAllOf(val *ReservedResourceAllOf) *NullableReservedResourceAllOf {
	return &NullableReservedResourceAllOf{value: val, isSet: true}
}

func (v NullableReservedResourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedResourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


