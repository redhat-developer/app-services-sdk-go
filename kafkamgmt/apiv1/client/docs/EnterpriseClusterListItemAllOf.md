# EnterpriseClusterListItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKafkasViaPrivateNetwork** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | Pointer to **string** | The OCM&#39;s cluster id of the registered Enterprise cluster. | [optional] 
**Status** | Pointer to **string** | The status of Enterprise cluster registration | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider for this cluster. This valus will be used as the Kafka&#39;s cloud provider value when a Kafka is created on this cluster | [optional] 
**Region** | Pointer to **string** | The region of this cluster. This valus will be used as the Kafka&#39;s region value when a Kafka is created on this cluster | [optional] 
**MultiAz** | **bool** | A flag indicating whether this cluster is available on multiple availability zones or not | 

## Methods

### NewEnterpriseClusterListItemAllOf

`func NewEnterpriseClusterListItemAllOf(accessKafkasViaPrivateNetwork bool, multiAz bool, ) *EnterpriseClusterListItemAllOf`

NewEnterpriseClusterListItemAllOf instantiates a new EnterpriseClusterListItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterListItemAllOfWithDefaults

`func NewEnterpriseClusterListItemAllOfWithDefaults() *EnterpriseClusterListItemAllOf`

NewEnterpriseClusterListItemAllOfWithDefaults instantiates a new EnterpriseClusterListItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterListItemAllOf) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseClusterListItemAllOf) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterListItemAllOf) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseClusterListItemAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterListItemAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterListItemAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterListItemAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterListItemAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterListItemAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterListItemAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterListItemAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnterpriseClusterListItemAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnterpriseClusterListItemAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnterpriseClusterListItemAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *EnterpriseClusterListItemAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *EnterpriseClusterListItemAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnterpriseClusterListItemAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnterpriseClusterListItemAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnterpriseClusterListItemAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMultiAz

`func (o *EnterpriseClusterListItemAllOf) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *EnterpriseClusterListItemAllOf) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *EnterpriseClusterListItemAllOf) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


