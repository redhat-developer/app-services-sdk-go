# AclBindingListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AclBinding**](AclBinding.md) |  | [optional] 
**Total** | Pointer to **float32** | Total number of entries in the full result set | [optional] 
**Page** | Pointer to **int32** | Current page number | [optional] 
**Size** | Pointer to **float32** | Number of entries per page | [optional] 


## Methods

### NewAclBindingListPage

`func NewAclBindingListPage() *AclBindingListPage`

NewAclBindingListPage instantiates a new AclBindingListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclBindingListPageWithDefaults

`func NewAclBindingListPageWithDefaults() *AclBindingListPage`

NewAclBindingListPageWithDefaults instantiates a new AclBindingListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetItems

`func (o *AclBindingListPage) GetItems() []AclBinding`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AclBindingListPage) GetItemsOk() (*[]AclBinding, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AclBindingListPage) SetItems(v []AclBinding)`

SetItems sets Items field to given value.

### HasItems

`func (o *AclBindingListPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


### GetTotal

`func (o *AclBindingListPage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AclBindingListPage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AclBindingListPage) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AclBindingListPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


### GetPage

`func (o *AclBindingListPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AclBindingListPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AclBindingListPage) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AclBindingListPage) HasPage() bool`

HasPage returns a boolean if a field has been set.


### GetSize

`func (o *AclBindingListPage) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AclBindingListPage) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AclBindingListPage) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AclBindingListPage) HasSize() bool`

HasSize returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

