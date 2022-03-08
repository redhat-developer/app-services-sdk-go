# TopicsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Topic**](Topic.md) |  | 
**Total** | **int32** | Total number of entries in the full result set | 
**Size** | Pointer to **int32** | Number of entries per page (returned for fetch requests) | [optional] 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

