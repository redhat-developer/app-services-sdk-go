# ListDeprecated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | **[]map[string]interface{}** |  | 
**Total** | **int32** | Total number of entries in the full result set | 
**Size** | Pointer to **int32** | Number of entries per page (returned for fetch requests) | [optional] 
**Page** | Pointer to **int32** | Current page number (returned for fetch requests) | [optional] 


## Methods

### NewListDeprecated

`func NewListDeprecated(items []map[string]interface{}, total int32, ) *ListDeprecated`

NewListDeprecated instantiates a new ListDeprecated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeprecatedWithDefaults

`func NewListDeprecatedWithDefaults() *ListDeprecated`

NewListDeprecatedWithDefaults instantiates a new ListDeprecated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKind

`func (o *ListDeprecated) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListDeprecated) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListDeprecated) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ListDeprecated) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetItems

`func (o *ListDeprecated) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListDeprecated) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListDeprecated) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.



### GetTotal

`func (o *ListDeprecated) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListDeprecated) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListDeprecated) SetTotal(v int32)`

SetTotal sets Total field to given value.



### GetSize

`func (o *ListDeprecated) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListDeprecated) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListDeprecated) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListDeprecated) HasSize() bool`

HasSize returns a boolean if a field has been set.


### GetPage

`func (o *ListDeprecated) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListDeprecated) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListDeprecated) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListDeprecated) HasPage() bool`

HasPage returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

