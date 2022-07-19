# ProcessorCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ProcessorSchemaEntryResponse**](ProcessorSchemaEntryResponse.md) |  | [optional] 

## Methods

### NewProcessorCatalogResponse

`func NewProcessorCatalogResponse() *ProcessorCatalogResponse`

NewProcessorCatalogResponse instantiates a new ProcessorCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorCatalogResponseWithDefaults

`func NewProcessorCatalogResponseWithDefaults() *ProcessorCatalogResponse`

NewProcessorCatalogResponseWithDefaults instantiates a new ProcessorCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ProcessorCatalogResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProcessorCatalogResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProcessorCatalogResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProcessorCatalogResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *ProcessorCatalogResponse) GetItems() []ProcessorSchemaEntryResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProcessorCatalogResponse) GetItemsOk() (*[]ProcessorSchemaEntryResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProcessorCatalogResponse) SetItems(v []ProcessorSchemaEntryResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProcessorCatalogResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


