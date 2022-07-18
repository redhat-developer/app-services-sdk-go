# TopicsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | [**[]Topic**](Topic.md) |  | 
**Total** | **int32** | Total number of entries in the full result set | 
**Size** | Pointer to **int32** | Number of entries per page (returned for fetch requests) | [optional] 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 
**Offset** | Pointer to **int32** | Offset of the first record returned, zero-based | [optional] 
**Limit** | Pointer to **int32** | Maximum number of records to return, from request | [optional] 
**Count** | Pointer to **int32** | Total number of entries in the full result set | [optional] 

## Methods

### NewTopicsList

`func NewTopicsList(items []Topic, total int32, ) *TopicsList`

NewTopicsList instantiates a new TopicsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicsListWithDefaults

`func NewTopicsListWithDefaults() *TopicsList`

NewTopicsListWithDefaults instantiates a new TopicsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TopicsList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TopicsList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TopicsList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TopicsList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *TopicsList) GetItems() []Topic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TopicsList) GetItemsOk() (*[]Topic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TopicsList) SetItems(v []Topic)`

SetItems sets Items field to given value.


### GetTotal

`func (o *TopicsList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TopicsList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TopicsList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSize

`func (o *TopicsList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TopicsList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TopicsList) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TopicsList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *TopicsList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TopicsList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TopicsList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TopicsList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetOffset

`func (o *TopicsList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TopicsList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TopicsList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TopicsList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *TopicsList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TopicsList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TopicsList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TopicsList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *TopicsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TopicsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TopicsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TopicsList) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


