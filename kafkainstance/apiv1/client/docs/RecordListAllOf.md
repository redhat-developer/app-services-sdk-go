# RecordListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Record**](Record.md) |  | 
**Total** | Pointer to **int32** | Total number of records returned in this request. This value does not indicate the total number of records in the topic. | [optional] 
**Size** | Pointer to **int32** | Not used | [optional] 
**Page** | Pointer to **int32** | Not used | [optional] 

## Methods

### NewRecordListAllOf

`func NewRecordListAllOf(items []Record, ) *RecordListAllOf`

NewRecordListAllOf instantiates a new RecordListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordListAllOfWithDefaults

`func NewRecordListAllOfWithDefaults() *RecordListAllOf`

NewRecordListAllOfWithDefaults instantiates a new RecordListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RecordListAllOf) GetItems() []Record`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecordListAllOf) GetItemsOk() (*[]Record, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecordListAllOf) SetItems(v []Record)`

SetItems sets Items field to given value.


### GetTotal

`func (o *RecordListAllOf) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RecordListAllOf) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RecordListAllOf) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RecordListAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSize

`func (o *RecordListAllOf) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RecordListAllOf) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RecordListAllOf) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RecordListAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *RecordListAllOf) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RecordListAllOf) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RecordListAllOf) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *RecordListAllOf) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


