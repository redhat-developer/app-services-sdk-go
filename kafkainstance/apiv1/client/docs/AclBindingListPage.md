# AclBindingListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Total** | **int32** | Total number of entries in the full result set | 
**Size** | Pointer to **int32** | Number of entries per page (returned for fetch requests) | [optional] 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 
**Items** | [**[]AclBinding**](AclBinding.md) |  | 

## Methods

### NewAclBindingListPage

`func NewAclBindingListPage(total int32, items []AclBinding, ) *AclBindingListPage`

NewAclBindingListPage instantiates a new AclBindingListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclBindingListPageWithDefaults

`func NewAclBindingListPageWithDefaults() *AclBindingListPage`

NewAclBindingListPageWithDefaults instantiates a new AclBindingListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AclBindingListPage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AclBindingListPage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AclBindingListPage) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AclBindingListPage) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTotal

`func (o *AclBindingListPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AclBindingListPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AclBindingListPage) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetSize

`func (o *AclBindingListPage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AclBindingListPage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AclBindingListPage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AclBindingListPage) HasSize() bool`

HasSize returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


