# EnterpriseClusterWithAddonParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**AccessKafkasViaPrivateNetwork** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | Pointer to **string** | OCM cluster id of the registered Enterprise cluster | [optional] 
**Status** | Pointer to **string** | status of registered Enterprise cluster | [optional] 
**FleetshardParameters** | Pointer to [**[]FleetshardParameter**](FleetshardParameter.md) |  | [optional] 

## Methods

### NewEnterpriseClusterWithAddonParameters

`func NewEnterpriseClusterWithAddonParameters(id string, kind string, href string, accessKafkasViaPrivateNetwork bool, ) *EnterpriseClusterWithAddonParameters`

NewEnterpriseClusterWithAddonParameters instantiates a new EnterpriseClusterWithAddonParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterWithAddonParametersWithDefaults

`func NewEnterpriseClusterWithAddonParametersWithDefaults() *EnterpriseClusterWithAddonParameters`

NewEnterpriseClusterWithAddonParametersWithDefaults instantiates a new EnterpriseClusterWithAddonParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnterpriseClusterWithAddonParameters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnterpriseClusterWithAddonParameters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnterpriseClusterWithAddonParameters) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *EnterpriseClusterWithAddonParameters) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnterpriseClusterWithAddonParameters) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnterpriseClusterWithAddonParameters) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *EnterpriseClusterWithAddonParameters) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EnterpriseClusterWithAddonParameters) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EnterpriseClusterWithAddonParameters) SetHref(v string)`

SetHref sets Href field to given value.


### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterWithAddonParameters) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseClusterWithAddonParameters) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterWithAddonParameters) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseClusterWithAddonParameters) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterWithAddonParameters) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterWithAddonParameters) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterWithAddonParameters) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterWithAddonParameters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterWithAddonParameters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterWithAddonParameters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterWithAddonParameters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFleetshardParameters

`func (o *EnterpriseClusterWithAddonParameters) GetFleetshardParameters() []FleetshardParameter`

GetFleetshardParameters returns the FleetshardParameters field if non-nil, zero value otherwise.

### GetFleetshardParametersOk

`func (o *EnterpriseClusterWithAddonParameters) GetFleetshardParametersOk() (*[]FleetshardParameter, bool)`

GetFleetshardParametersOk returns a tuple with the FleetshardParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetshardParameters

`func (o *EnterpriseClusterWithAddonParameters) SetFleetshardParameters(v []FleetshardParameter)`

SetFleetshardParameters sets FleetshardParameters field to given value.

### HasFleetshardParameters

`func (o *EnterpriseClusterWithAddonParameters) HasFleetshardParameters() bool`

HasFleetshardParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


