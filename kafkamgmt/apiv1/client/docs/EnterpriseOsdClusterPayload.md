# EnterpriseOsdClusterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKafkasViaPrivateNetwork** | **bool** | Sets whether Kafkas created on this data plane cluster have to be accessed via private network | 
**ClusterId** | **string** | The data plane cluster ID. This is the ID of the cluster obtained from OpenShift Cluster Manager (OCM) API | 
**ClusterIngressDnsName** | **string** | dns name of the cluster. Can be obtained from the response JSON of the /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt;/ingresses (dns_name) | 
**KafkaMachinePoolNodeCount** | **int32** | The node count given to the created kafka machine pool.  The machine pool must be created via /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt;/machine_pools prior to passing this value. The created machine pool must have a &#x60;bf2.org/kafkaInstanceProfileType&#x3D;standard&#x60; label and a &#x60;bf2.org/kafkaInstanceProfileType&#x3D;standard:NoExecute&#x60; taint. The name of the machine pool must be &#x60;kafka-standard&#x60;  The node count value has to be a multiple of 3 with a minimum of 3 nodes. | 

## Methods

### NewEnterpriseOsdClusterPayload

`func NewEnterpriseOsdClusterPayload(accessKafkasViaPrivateNetwork bool, clusterId string, clusterIngressDnsName string, kafkaMachinePoolNodeCount int32, ) *EnterpriseOsdClusterPayload`

NewEnterpriseOsdClusterPayload instantiates a new EnterpriseOsdClusterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseOsdClusterPayloadWithDefaults

`func NewEnterpriseOsdClusterPayloadWithDefaults() *EnterpriseOsdClusterPayload`

NewEnterpriseOsdClusterPayloadWithDefaults instantiates a new EnterpriseOsdClusterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseOsdClusterPayload) GetAccessKafkasViaPrivateNetwork() bool`

GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field if non-nil, zero value otherwise.

### GetAccessKafkasViaPrivateNetworkOk

`func (o *EnterpriseOsdClusterPayload) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool)`

GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKafkasViaPrivateNetwork

`func (o *EnterpriseOsdClusterPayload) SetAccessKafkasViaPrivateNetwork(v bool)`

SetAccessKafkasViaPrivateNetwork sets AccessKafkasViaPrivateNetwork field to given value.


### GetClusterId

`func (o *EnterpriseOsdClusterPayload) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseOsdClusterPayload) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseOsdClusterPayload) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterIngressDnsName

`func (o *EnterpriseOsdClusterPayload) GetClusterIngressDnsName() string`

GetClusterIngressDnsName returns the ClusterIngressDnsName field if non-nil, zero value otherwise.

### GetClusterIngressDnsNameOk

`func (o *EnterpriseOsdClusterPayload) GetClusterIngressDnsNameOk() (*string, bool)`

GetClusterIngressDnsNameOk returns a tuple with the ClusterIngressDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIngressDnsName

`func (o *EnterpriseOsdClusterPayload) SetClusterIngressDnsName(v string)`

SetClusterIngressDnsName sets ClusterIngressDnsName field to given value.


### GetKafkaMachinePoolNodeCount

`func (o *EnterpriseOsdClusterPayload) GetKafkaMachinePoolNodeCount() int32`

GetKafkaMachinePoolNodeCount returns the KafkaMachinePoolNodeCount field if non-nil, zero value otherwise.

### GetKafkaMachinePoolNodeCountOk

`func (o *EnterpriseOsdClusterPayload) GetKafkaMachinePoolNodeCountOk() (*int32, bool)`

GetKafkaMachinePoolNodeCountOk returns a tuple with the KafkaMachinePoolNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaMachinePoolNodeCount

`func (o *EnterpriseOsdClusterPayload) SetKafkaMachinePoolNodeCount(v int32)`

SetKafkaMachinePoolNodeCount sets KafkaMachinePoolNodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


