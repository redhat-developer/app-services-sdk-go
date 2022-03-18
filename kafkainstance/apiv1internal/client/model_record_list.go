/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.8.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// RecordList A page of records consumed from a topic
type RecordList struct {

	// Total number of records returned in this request. This value does not indicate the total number of records in the topic.
	Total int32 `json:"total"`

	// Not used
	Size *int32 `json:"size,omitempty"`

	// Not used
	Page *int32 `json:"page,omitempty"`

	Items []Record `json:"items"`

}

// NewRecordList instantiates a new RecordList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordList(total int32, items []Record) *RecordList {
	this := RecordList{}
	this.Total = total
	this.Items = items
	return &this
}

// NewRecordListWithDefaults instantiates a new RecordList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordListWithDefaults() *RecordList {
	this := RecordList{}





	return &this
}


// GetTotal returns the Total field value
func (o *RecordList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RecordList) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *RecordList) SetTotal(v int32) {
	o.Total = v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *RecordList) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordList) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RecordList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *RecordList) SetSize(v int32) {
	o.Size = &v
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *RecordList) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordList) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *RecordList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *RecordList) SetPage(v int32) {
	o.Page = &v
}


// GetItems returns the Items field value
func (o *RecordList) GetItems() []Record {
	if o == nil {
		var ret []Record
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RecordList) GetItemsOk() (*[]Record, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *RecordList) SetItems(v []Record) {
	o.Items = v
}


func (o RecordList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["total"] = o.Total
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	if true {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableRecordList struct {
	value *RecordList
	isSet bool
}

func (v NullableRecordList) Get() *RecordList {
	return v.value
}

func (v *NullableRecordList) Set(val *RecordList) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordList) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordList(val *RecordList) *NullableRecordList {
	return &NullableRecordList{value: val, isSet: true}
}

func (v NullableRecordList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

