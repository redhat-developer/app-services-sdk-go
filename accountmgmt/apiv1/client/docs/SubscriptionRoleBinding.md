# SubscriptionRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AccountReference**](AccountReference.md) |  | [optional] 
**AccountEmail** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Role** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriptionRoleBinding

`func NewSubscriptionRoleBinding() *SubscriptionRoleBinding`

NewSubscriptionRoleBinding instantiates a new SubscriptionRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRoleBindingWithDefaults

`func NewSubscriptionRoleBindingWithDefaults() *SubscriptionRoleBinding`

NewSubscriptionRoleBindingWithDefaults instantiates a new SubscriptionRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SubscriptionRoleBinding) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SubscriptionRoleBinding) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SubscriptionRoleBinding) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SubscriptionRoleBinding) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionRoleBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionRoleBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionRoleBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionRoleBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SubscriptionRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SubscriptionRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SubscriptionRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SubscriptionRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAccount

`func (o *SubscriptionRoleBinding) GetAccount() AccountReference`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubscriptionRoleBinding) GetAccountOk() (*AccountReference, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubscriptionRoleBinding) SetAccount(v AccountReference)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SubscriptionRoleBinding) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountEmail

`func (o *SubscriptionRoleBinding) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *SubscriptionRoleBinding) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *SubscriptionRoleBinding) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *SubscriptionRoleBinding) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### GetAccountUsername

`func (o *SubscriptionRoleBinding) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *SubscriptionRoleBinding) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *SubscriptionRoleBinding) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *SubscriptionRoleBinding) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionRoleBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionRoleBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionRoleBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionRoleBinding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRole

`func (o *SubscriptionRoleBinding) GetRole() ObjectReference`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SubscriptionRoleBinding) GetRoleOk() (*ObjectReference, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SubscriptionRoleBinding) SetRole(v ObjectReference)`

SetRole sets Role field to given value.

### HasRole

`func (o *SubscriptionRoleBinding) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscription

`func (o *SubscriptionRoleBinding) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionRoleBinding) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionRoleBinding) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SubscriptionRoleBinding) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionRoleBinding) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionRoleBinding) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionRoleBinding) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionRoleBinding) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


