# ProcessorListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ProcessorResponse**](ProcessorResponse.md) |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewProcessorListResponse

`func NewProcessorListResponse() *ProcessorListResponse`

NewProcessorListResponse instantiates a new ProcessorListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorListResponseWithDefaults

`func NewProcessorListResponseWithDefaults() *ProcessorListResponse`

NewProcessorListResponseWithDefaults instantiates a new ProcessorListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ProcessorListResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProcessorListResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProcessorListResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProcessorListResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *ProcessorListResponse) GetItems() []ProcessorResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProcessorListResponse) GetItemsOk() (*[]ProcessorResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProcessorListResponse) SetItems(v []ProcessorResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProcessorListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *ProcessorListResponse) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProcessorListResponse) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProcessorListResponse) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProcessorListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *ProcessorListResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProcessorListResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProcessorListResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ProcessorListResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *ProcessorListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProcessorListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProcessorListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProcessorListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


