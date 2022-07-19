# CloudProviderListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]CloudProviderResponse**](CloudProviderResponse.md) |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCloudProviderListResponse

`func NewCloudProviderListResponse() *CloudProviderListResponse`

NewCloudProviderListResponse instantiates a new CloudProviderListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderListResponseWithDefaults

`func NewCloudProviderListResponseWithDefaults() *CloudProviderListResponse`

NewCloudProviderListResponseWithDefaults instantiates a new CloudProviderListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudProviderListResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProviderListResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProviderListResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudProviderListResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *CloudProviderListResponse) GetItems() []CloudProviderResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudProviderListResponse) GetItemsOk() (*[]CloudProviderResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudProviderListResponse) SetItems(v []CloudProviderResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudProviderListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *CloudProviderListResponse) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CloudProviderListResponse) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CloudProviderListResponse) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *CloudProviderListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *CloudProviderListResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudProviderListResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudProviderListResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudProviderListResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *CloudProviderListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudProviderListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudProviderListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudProviderListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


