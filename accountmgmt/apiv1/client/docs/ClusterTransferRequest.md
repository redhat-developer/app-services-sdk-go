# ClusterTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterTransferRequest

`func NewClusterTransferRequest() *ClusterTransferRequest`

NewClusterTransferRequest instantiates a new ClusterTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTransferRequestWithDefaults

`func NewClusterTransferRequestWithDefaults() *ClusterTransferRequest`

NewClusterTransferRequestWithDefaults instantiates a new ClusterTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *ClusterTransferRequest) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *ClusterTransferRequest) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *ClusterTransferRequest) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *ClusterTransferRequest) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetOwner

`func (o *ClusterTransferRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ClusterTransferRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ClusterTransferRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ClusterTransferRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRecipient

`func (o *ClusterTransferRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ClusterTransferRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ClusterTransferRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *ClusterTransferRequest) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


