# CloudResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Category** | Pointer to **string** |  | [optional] 
**CategoryPretty** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GenericName** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to **int64** |  | [optional] 
**MemoryPretty** | Pointer to **string** |  | [optional] 
**NamePretty** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**SizePretty** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudResource

`func NewCloudResource() *CloudResource`

NewCloudResource instantiates a new CloudResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResourceWithDefaults

`func NewCloudResourceWithDefaults() *CloudResource`

NewCloudResourceWithDefaults instantiates a new CloudResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CloudResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudResource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudResource) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *CloudResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetActive

`func (o *CloudResource) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudResource) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudResource) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudResource) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCategory

`func (o *CloudResource) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudResource) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudResource) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudResource) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryPretty

`func (o *CloudResource) GetCategoryPretty() string`

GetCategoryPretty returns the CategoryPretty field if non-nil, zero value otherwise.

### GetCategoryPrettyOk

`func (o *CloudResource) GetCategoryPrettyOk() (*string, bool)`

GetCategoryPrettyOk returns a tuple with the CategoryPretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPretty

`func (o *CloudResource) SetCategoryPretty(v string)`

SetCategoryPretty sets CategoryPretty field to given value.

### HasCategoryPretty

`func (o *CloudResource) HasCategoryPretty() bool`

HasCategoryPretty returns a boolean if a field has been set.

### GetCloudProvider

`func (o *CloudResource) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CloudResource) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CloudResource) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CloudResource) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCpuCores

`func (o *CloudResource) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *CloudResource) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *CloudResource) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *CloudResource) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGenericName

`func (o *CloudResource) GetGenericName() string`

GetGenericName returns the GenericName field if non-nil, zero value otherwise.

### GetGenericNameOk

`func (o *CloudResource) GetGenericNameOk() (*string, bool)`

GetGenericNameOk returns a tuple with the GenericName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericName

`func (o *CloudResource) SetGenericName(v string)`

SetGenericName sets GenericName field to given value.

### HasGenericName

`func (o *CloudResource) HasGenericName() bool`

HasGenericName returns a boolean if a field has been set.

### GetMemory

`func (o *CloudResource) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudResource) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudResource) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudResource) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMemoryPretty

`func (o *CloudResource) GetMemoryPretty() string`

GetMemoryPretty returns the MemoryPretty field if non-nil, zero value otherwise.

### GetMemoryPrettyOk

`func (o *CloudResource) GetMemoryPrettyOk() (*string, bool)`

GetMemoryPrettyOk returns a tuple with the MemoryPretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPretty

`func (o *CloudResource) SetMemoryPretty(v string)`

SetMemoryPretty sets MemoryPretty field to given value.

### HasMemoryPretty

`func (o *CloudResource) HasMemoryPretty() bool`

HasMemoryPretty returns a boolean if a field has been set.

### GetNamePretty

`func (o *CloudResource) GetNamePretty() string`

GetNamePretty returns the NamePretty field if non-nil, zero value otherwise.

### GetNamePrettyOk

`func (o *CloudResource) GetNamePrettyOk() (*string, bool)`

GetNamePrettyOk returns a tuple with the NamePretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePretty

`func (o *CloudResource) SetNamePretty(v string)`

SetNamePretty sets NamePretty field to given value.

### HasNamePretty

`func (o *CloudResource) HasNamePretty() bool`

HasNamePretty returns a boolean if a field has been set.

### GetResourceType

`func (o *CloudResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CloudResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CloudResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CloudResource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSizePretty

`func (o *CloudResource) GetSizePretty() string`

GetSizePretty returns the SizePretty field if non-nil, zero value otherwise.

### GetSizePrettyOk

`func (o *CloudResource) GetSizePrettyOk() (*string, bool)`

GetSizePrettyOk returns a tuple with the SizePretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePretty

`func (o *CloudResource) SetSizePretty(v string)`

SetSizePretty sets SizePretty field to given value.

### HasSizePretty

`func (o *CloudResource) HasSizePretty() bool`

HasSizePretty returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


