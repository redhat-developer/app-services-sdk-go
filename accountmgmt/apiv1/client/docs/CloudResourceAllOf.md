# CloudResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewCloudResourceAllOf

`func NewCloudResourceAllOf() *CloudResourceAllOf`

NewCloudResourceAllOf instantiates a new CloudResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResourceAllOfWithDefaults

`func NewCloudResourceAllOfWithDefaults() *CloudResourceAllOf`

NewCloudResourceAllOfWithDefaults instantiates a new CloudResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudResourceAllOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudResourceAllOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudResourceAllOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudResourceAllOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCategory

`func (o *CloudResourceAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudResourceAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudResourceAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudResourceAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryPretty

`func (o *CloudResourceAllOf) GetCategoryPretty() string`

GetCategoryPretty returns the CategoryPretty field if non-nil, zero value otherwise.

### GetCategoryPrettyOk

`func (o *CloudResourceAllOf) GetCategoryPrettyOk() (*string, bool)`

GetCategoryPrettyOk returns a tuple with the CategoryPretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPretty

`func (o *CloudResourceAllOf) SetCategoryPretty(v string)`

SetCategoryPretty sets CategoryPretty field to given value.

### HasCategoryPretty

`func (o *CloudResourceAllOf) HasCategoryPretty() bool`

HasCategoryPretty returns a boolean if a field has been set.

### GetCloudProvider

`func (o *CloudResourceAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CloudResourceAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CloudResourceAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CloudResourceAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCpuCores

`func (o *CloudResourceAllOf) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *CloudResourceAllOf) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *CloudResourceAllOf) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *CloudResourceAllOf) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudResourceAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudResourceAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudResourceAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudResourceAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGenericName

`func (o *CloudResourceAllOf) GetGenericName() string`

GetGenericName returns the GenericName field if non-nil, zero value otherwise.

### GetGenericNameOk

`func (o *CloudResourceAllOf) GetGenericNameOk() (*string, bool)`

GetGenericNameOk returns a tuple with the GenericName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericName

`func (o *CloudResourceAllOf) SetGenericName(v string)`

SetGenericName sets GenericName field to given value.

### HasGenericName

`func (o *CloudResourceAllOf) HasGenericName() bool`

HasGenericName returns a boolean if a field has been set.

### GetMemory

`func (o *CloudResourceAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CloudResourceAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CloudResourceAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CloudResourceAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMemoryPretty

`func (o *CloudResourceAllOf) GetMemoryPretty() string`

GetMemoryPretty returns the MemoryPretty field if non-nil, zero value otherwise.

### GetMemoryPrettyOk

`func (o *CloudResourceAllOf) GetMemoryPrettyOk() (*string, bool)`

GetMemoryPrettyOk returns a tuple with the MemoryPretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPretty

`func (o *CloudResourceAllOf) SetMemoryPretty(v string)`

SetMemoryPretty sets MemoryPretty field to given value.

### HasMemoryPretty

`func (o *CloudResourceAllOf) HasMemoryPretty() bool`

HasMemoryPretty returns a boolean if a field has been set.

### GetNamePretty

`func (o *CloudResourceAllOf) GetNamePretty() string`

GetNamePretty returns the NamePretty field if non-nil, zero value otherwise.

### GetNamePrettyOk

`func (o *CloudResourceAllOf) GetNamePrettyOk() (*string, bool)`

GetNamePrettyOk returns a tuple with the NamePretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePretty

`func (o *CloudResourceAllOf) SetNamePretty(v string)`

SetNamePretty sets NamePretty field to given value.

### HasNamePretty

`func (o *CloudResourceAllOf) HasNamePretty() bool`

HasNamePretty returns a boolean if a field has been set.

### GetResourceType

`func (o *CloudResourceAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CloudResourceAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CloudResourceAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CloudResourceAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSizePretty

`func (o *CloudResourceAllOf) GetSizePretty() string`

GetSizePretty returns the SizePretty field if non-nil, zero value otherwise.

### GetSizePrettyOk

`func (o *CloudResourceAllOf) GetSizePrettyOk() (*string, bool)`

GetSizePrettyOk returns a tuple with the SizePretty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizePretty

`func (o *CloudResourceAllOf) SetSizePretty(v string)`

SetSizePretty sets SizePretty field to given value.

### HasSizePretty

`func (o *CloudResourceAllOf) HasSizePretty() bool`

HasSizePretty returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudResourceAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudResourceAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudResourceAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudResourceAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


