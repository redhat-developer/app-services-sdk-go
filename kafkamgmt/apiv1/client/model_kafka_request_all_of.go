/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.1.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// KafkaRequestAllOf struct for KafkaRequestAllOf
type KafkaRequestAllOf struct {

	// Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting] 
	Status *string `json:"status,omitempty"`

	// Name of Cloud used to deploy. For example AWS
	CloudProvider *string `json:"cloud_provider,omitempty"`

	MultiAz *bool `json:"multi_az,omitempty"`

	// Values will be regions of specific cloud provider. For example: us-east-1 for AWS
	Region *string `json:"region,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Name *string `json:"name,omitempty"`

	BootstrapServerHost *string `json:"bootstrap_server_host,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	Version *string `json:"version,omitempty"`

	InstanceType *string `json:"instance_type,omitempty"`

}

// NewKafkaRequestAllOf instantiates a new KafkaRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequestAllOf() *KafkaRequestAllOf {
	this := KafkaRequestAllOf{}
	return &this
}

// NewKafkaRequestAllOfWithDefaults instantiates a new KafkaRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestAllOfWithDefaults() *KafkaRequestAllOf {
	this := KafkaRequestAllOf{}













	return &this
}


// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KafkaRequestAllOf) SetStatus(v string) {
	o.Status = &v
}


// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KafkaRequestAllOf) SetCloudProvider(v string) {
	o.CloudProvider = &v
}


// GetMultiAz returns the MultiAz field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetMultiAz() bool {
	if o == nil || o.MultiAz == nil {
		var ret bool
		return ret
	}
	return *o.MultiAz
}

// GetMultiAzOk returns a tuple with the MultiAz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetMultiAzOk() (*bool, bool) {
	if o == nil || o.MultiAz == nil {
		return nil, false
	}
	return o.MultiAz, true
}

// HasMultiAz returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasMultiAz() bool {
	if o != nil && o.MultiAz != nil {
		return true
	}

	return false
}

// SetMultiAz gets a reference to the given bool and assigns it to the MultiAz field.
func (o *KafkaRequestAllOf) SetMultiAz(v bool) {
	o.MultiAz = &v
}


// GetRegion returns the Region field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *KafkaRequestAllOf) SetRegion(v string) {
	o.Region = &v
}


// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *KafkaRequestAllOf) SetOwner(v string) {
	o.Owner = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KafkaRequestAllOf) SetName(v string) {
	o.Name = &v
}


// GetBootstrapServerHost returns the BootstrapServerHost field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetBootstrapServerHost() string {
	if o == nil || o.BootstrapServerHost == nil {
		var ret string
		return ret
	}
	return *o.BootstrapServerHost
}

// GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetBootstrapServerHostOk() (*string, bool) {
	if o == nil || o.BootstrapServerHost == nil {
		return nil, false
	}
	return o.BootstrapServerHost, true
}

// HasBootstrapServerHost returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasBootstrapServerHost() bool {
	if o != nil && o.BootstrapServerHost != nil {
		return true
	}

	return false
}

// SetBootstrapServerHost gets a reference to the given string and assigns it to the BootstrapServerHost field.
func (o *KafkaRequestAllOf) SetBootstrapServerHost(v string) {
	o.BootstrapServerHost = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *KafkaRequestAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *KafkaRequestAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}


// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetFailedReason() string {
	if o == nil || o.FailedReason == nil {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetFailedReasonOk() (*string, bool) {
	if o == nil || o.FailedReason == nil {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasFailedReason() bool {
	if o != nil && o.FailedReason != nil {
		return true
	}

	return false
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *KafkaRequestAllOf) SetFailedReason(v string) {
	o.FailedReason = &v
}


// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KafkaRequestAllOf) SetVersion(v string) {
	o.Version = &v
}


// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *KafkaRequestAllOf) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequestAllOf) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *KafkaRequestAllOf) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *KafkaRequestAllOf) SetInstanceType(v string) {
	o.InstanceType = &v
}


func (o KafkaRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
    
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
    
	if o.MultiAz != nil {
		toSerialize["multi_az"] = o.MultiAz
	}
    
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
    
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.BootstrapServerHost != nil {
		toSerialize["bootstrap_server_host"] = o.BootstrapServerHost
	}
    
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
    
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
    
	if o.FailedReason != nil {
		toSerialize["failed_reason"] = o.FailedReason
	}
    
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
    
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
    
	return json.Marshal(toSerialize)
}

type NullableKafkaRequestAllOf struct {
	value *KafkaRequestAllOf
	isSet bool
}

func (v NullableKafkaRequestAllOf) Get() *KafkaRequestAllOf {
	return v.value
}

func (v *NullableKafkaRequestAllOf) Set(val *KafkaRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequestAllOf(val *KafkaRequestAllOf) *NullableKafkaRequestAllOf {
	return &NullableKafkaRequestAllOf{value: val, isSet: true}
}

func (v NullableKafkaRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

