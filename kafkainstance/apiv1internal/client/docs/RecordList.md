# RecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Record**](Record.md) |  | [optional] 
**Total** | Pointer to **interface{}** | Total number of records returned in this request. This value does not indicate the total number of records in the topic. | [optional] 
**Size** | Pointer to **interface{}** | Not used | [optional] 
**Page** | Pointer to **interface{}** | Not used | [optional] 
**Kind** | Pointer to **string** |  | [optional] 


## Methods

### NewRecordList

`func NewRecordList() *RecordList`

NewRecordList instantiates a new RecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordListWithDefaults

`func NewRecordListWithDefaults() *RecordList`

NewRecordListWithDefaults instantiates a new RecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


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

### HasItems

`func (o *RecordList) HasItems() bool`

HasItems returns a boolean if a field has been set.


### GetTotal

`func (o *RecordList) GetTotal() interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RecordList) GetTotalOk() (*interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RecordList) SetTotal(v interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RecordList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *RecordList) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *RecordList) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

### GetSize

`func (o *RecordList) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RecordList) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RecordList) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *RecordList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *RecordList) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *RecordList) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

### GetPage

`func (o *RecordList) GetPage() interface{}`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RecordList) GetPageOk() (*interface{}, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RecordList) SetPage(v interface{})`

SetPage sets Page field to given value.

### HasPage

`func (o *RecordList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *RecordList) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *RecordList) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

