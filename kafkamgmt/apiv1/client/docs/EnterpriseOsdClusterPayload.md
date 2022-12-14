# EnterpriseOsdClusterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | OSD cluster ID | 
**ClusterExternalId** | **string** | external cluster ID. Can be obtained from the response JSON of ocm get /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt; | 
**ClusterIngressDnsName** | **string** | dns name of the cluster. Can be obtained from the response JSON of the /api/clusters_mgmt/v1/clusters/&lt;cluster_id&gt;/ingresses (dns_name) | 

## Methods

### NewEnterpriseOsdClusterPayload

`func NewEnterpriseOsdClusterPayload(clusterId string, clusterExternalId string, clusterIngressDnsName string, ) *EnterpriseOsdClusterPayload`

NewEnterpriseOsdClusterPayload instantiates a new EnterpriseOsdClusterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseOsdClusterPayloadWithDefaults

`func NewEnterpriseOsdClusterPayloadWithDefaults() *EnterpriseOsdClusterPayload`

NewEnterpriseOsdClusterPayloadWithDefaults instantiates a new EnterpriseOsdClusterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetClusterExternalId

`func (o *EnterpriseOsdClusterPayload) GetClusterExternalId() string`

GetClusterExternalId returns the ClusterExternalId field if non-nil, zero value otherwise.

### GetClusterExternalIdOk

`func (o *EnterpriseOsdClusterPayload) GetClusterExternalIdOk() (*string, bool)`

GetClusterExternalIdOk returns a tuple with the ClusterExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExternalId

`func (o *EnterpriseOsdClusterPayload) SetClusterExternalId(v string)`

SetClusterExternalId sets ClusterExternalId field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


