# ClusterTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ClusterUuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClusterTransfer

`func NewClusterTransfer() *ClusterTransfer`

NewClusterTransfer instantiates a new ClusterTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTransferWithDefaults

`func NewClusterTransferWithDefaults() *ClusterTransfer`

NewClusterTransferWithDefaults instantiates a new ClusterTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ClusterTransfer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterTransfer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterTransfer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ClusterTransfer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ClusterTransfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterTransfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterTransfer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterTransfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ClusterTransfer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterTransfer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterTransfer) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterTransfer) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetClusterUuid

`func (o *ClusterTransfer) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *ClusterTransfer) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *ClusterTransfer) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *ClusterTransfer) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClusterTransfer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterTransfer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterTransfer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterTransfer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ClusterTransfer) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ClusterTransfer) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ClusterTransfer) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ClusterTransfer) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetOwner

`func (o *ClusterTransfer) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ClusterTransfer) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ClusterTransfer) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ClusterTransfer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRecipient

`func (o *ClusterTransfer) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ClusterTransfer) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ClusterTransfer) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ClusterTransfer) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSecret

`func (o *ClusterTransfer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClusterTransfer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClusterTransfer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ClusterTransfer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterTransfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterTransfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterTransfer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterTransfer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClusterTransfer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterTransfer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterTransfer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClusterTransfer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


