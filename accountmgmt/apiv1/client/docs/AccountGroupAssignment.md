# AccountGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AccountGroupId** | **string** |  | 
**AccountId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountGroupAssignment

`func NewAccountGroupAssignment(accountGroupId string, accountId string, ) *AccountGroupAssignment`

NewAccountGroupAssignment instantiates a new AccountGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupAssignmentWithDefaults

`func NewAccountGroupAssignmentWithDefaults() *AccountGroupAssignment`

NewAccountGroupAssignmentWithDefaults instantiates a new AccountGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AccountGroupAssignment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccountGroupAssignment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccountGroupAssignment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AccountGroupAssignment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AccountGroupAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AccountGroupAssignment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountGroupAssignment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountGroupAssignment) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AccountGroupAssignment) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAccountGroupId

`func (o *AccountGroupAssignment) GetAccountGroupId() string`

GetAccountGroupId returns the AccountGroupId field if non-nil, zero value otherwise.

### GetAccountGroupIdOk

`func (o *AccountGroupAssignment) GetAccountGroupIdOk() (*string, bool)`

GetAccountGroupIdOk returns a tuple with the AccountGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupId

`func (o *AccountGroupAssignment) SetAccountGroupId(v string)`

SetAccountGroupId sets AccountGroupId field to given value.


### GetAccountId

`func (o *AccountGroupAssignment) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountGroupAssignment) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountGroupAssignment) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *AccountGroupAssignment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountGroupAssignment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountGroupAssignment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountGroupAssignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountGroupAssignment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountGroupAssignment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountGroupAssignment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountGroupAssignment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


