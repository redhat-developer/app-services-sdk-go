# EnterpriseClusterRegistrationResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKafkasViaPrivateNetwork** | **bool** | Indicates whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | Pointer to **string** | OCM cluster id of the registered Enterprise cluster | [optional] 
**Status** | Pointer to **string** | status of registered Enterprise cluster | [optional] 
**FleetshardParameters** | Pointer to [**[]FleetshardParameter**](FleetshardParameter.md) |  | [optional] 

## Methods

### NewEnterpriseClusterRegistrationResponseAllOf

`func NewEnterpriseClusterRegistrationResponseAllOf(accessKafkasViaPrivateNetwork bool, ) *EnterpriseClusterRegistrationResponseAllOf`

NewEnterpriseClusterRegistrationResponseAllOf instantiates a new EnterpriseClusterRegistrationResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterRegistrationResponseAllOfWithDefaults

`func NewEnterpriseClusterRegistrationResponseAllOfWithDefaults() *EnterpriseClusterRegistrationResponseAllOf`

NewEnterpriseClusterRegistrationResponseAllOfWithDefaults instantiates a new EnterpriseClusterRegistrationResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseClusterRegistrationResponseAllOf) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterRegistrationResponseAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterRegistrationResponseAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterRegistrationResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterRegistrationResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetFleetshardParameters() []FleetshardParameter`

GetFleetshardParameters returns the FleetshardParameters field if non-nil, zero value otherwise.

### GetFleetshardParametersOk

`func (o *EnterpriseClusterRegistrationResponseAllOf) GetFleetshardParametersOk() (*[]FleetshardParameter, bool)`

GetFleetshardParametersOk returns a tuple with the FleetshardParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponseAllOf) SetFleetshardParameters(v []FleetshardParameter)`

SetFleetshardParameters sets FleetshardParameters field to given value.

### HasFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponseAllOf) HasFleetshardParameters() bool`

HasFleetshardParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


