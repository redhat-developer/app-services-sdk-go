# ResultListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | Total number of entries in the full result set | 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 
**Size** | Pointer to **float32** | Number of entries per page (returned for fetch requests) | [optional] 


## Methods

### NewResultListPage

`func NewResultListPage(total float32, ) *ResultListPage`

NewResultListPage instantiates a new ResultListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultListPageWithDefaults

`func NewResultListPageWithDefaults() *ResultListPage`

NewResultListPageWithDefaults instantiates a new ResultListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetTotal

`func (o *ResultListPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResultListPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResultListPage) SetTotal(v float32)`

SetTotal sets Total field to given value.



### GetPage

`func (o *ResultListPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResultListPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResultListPage) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ResultListPage) HasPage() bool`

HasPage returns a boolean if a field has been set.


### GetSize

`func (o *ResultListPage) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResultListPage) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResultListPage) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResultListPage) HasSize() bool`

HasSize returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

