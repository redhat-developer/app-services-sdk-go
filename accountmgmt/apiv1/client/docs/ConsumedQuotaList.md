# ConsumedQuotaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Page** | **int32** |  | 
**Size** | **int32** |  | 
**Total** | **int32** |  | 
**Items** | [**[]ConsumedQuota**](ConsumedQuota.md) |  | 

## Methods

### NewConsumedQuotaList

`func NewConsumedQuotaList(kind string, page int32, size int32, total int32, items []ConsumedQuota, ) *ConsumedQuotaList`

NewConsumedQuotaList instantiates a new ConsumedQuotaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumedQuotaListWithDefaults

`func NewConsumedQuotaListWithDefaults() *ConsumedQuotaList`

NewConsumedQuotaListWithDefaults instantiates a new ConsumedQuotaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConsumedQuotaList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConsumedQuotaList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConsumedQuotaList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPage

`func (o *ConsumedQuotaList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ConsumedQuotaList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ConsumedQuotaList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *ConsumedQuotaList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ConsumedQuotaList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ConsumedQuotaList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *ConsumedQuotaList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConsumedQuotaList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConsumedQuotaList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ConsumedQuotaList) GetItems() []ConsumedQuota`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConsumedQuotaList) GetItemsOk() (*[]ConsumedQuota, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConsumedQuotaList) SetItems(v []ConsumedQuota)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


