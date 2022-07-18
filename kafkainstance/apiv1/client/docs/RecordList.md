# RecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | [**[]Record**](Record.md) |  | 
**Total** | **int32** | Total number of records returned in this request. This value does not indicate the total number of records in the topic. | 
**Size** | Pointer to **int32** | Not used | [optional] 
**Page** | Pointer to **int32** | Not used | [optional] 

## Methods

### NewRecordList

`func NewRecordList(items []Record, total int32, ) *RecordList`

NewRecordList instantiates a new RecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordListWithDefaults

`func NewRecordListWithDefaults() *RecordList`

NewRecordListWithDefaults instantiates a new RecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RecordList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RecordList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RecordList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RecordList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *RecordList) GetItems() []Record`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecordList) GetItemsOk() (*[]Record, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecordList) SetItems(v []Record)`

SetItems sets Items field to given value.


### GetTotal

`func (o *RecordList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RecordList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RecordList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSize

`func (o *RecordList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RecordList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RecordList) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RecordList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *RecordList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RecordList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RecordList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *RecordList) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


