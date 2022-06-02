# RecordListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Record**](Record.md) |  | [optional] 
**Total** | Pointer to **interface{}** | Total number of records returned in this request. This value does not indicate the total number of records in the topic. | [optional] 
**Size** | Pointer to **interface{}** | Not used | [optional] 
**Page** | Pointer to **interface{}** | Not used | [optional] 


## Methods

### NewRecordListAllOf

`func NewRecordListAllOf() *RecordListAllOf`

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

### HasItems

`func (o *RecordListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


### GetTotal

`func (o *RecordListAllOf) GetTotal() interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RecordListAllOf) GetTotalOk() (*interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RecordListAllOf) SetTotal(v interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RecordListAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *RecordListAllOf) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *RecordListAllOf) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

### GetSize

`func (o *RecordListAllOf) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RecordListAllOf) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RecordListAllOf) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *RecordListAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *RecordListAllOf) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *RecordListAllOf) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

### GetPage

`func (o *RecordListAllOf) GetPage() interface{}`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RecordListAllOf) GetPageOk() (*interface{}, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RecordListAllOf) SetPage(v interface{})`

SetPage sets Page field to given value.

### HasPage

`func (o *RecordListAllOf) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *RecordListAllOf) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *RecordListAllOf) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

