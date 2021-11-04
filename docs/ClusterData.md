# ClusterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerServiceUrl** | Pointer to **string** | The broker service url (for produce and consume operations) | [optional] 
**BrokerServiceUrlTls** | Pointer to **string** | The secured broker service url (for produce and consume operations) | [optional] 
**PeerClusterNames** | Pointer to **[]string** | A set of peer cluster names | [optional] 
**ProxyProtocol** | Pointer to **string** | protocol to decide type of proxy routing eg: SNI-routing | [optional] 
**ProxyServiceUrl** | Pointer to **string** | Proxy-service url when client would like to connect to broker via proxy. | [optional] 
**ServiceUrl** | Pointer to **string** | The HTTP rest service URL (for admin operations) | [optional] 
**ServiceUrlTls** | Pointer to **string** | The HTTPS rest service URL (for admin operations) | [optional] 

## Methods

### NewClusterData

`func NewClusterData() *ClusterData`

NewClusterData instantiates a new ClusterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDataWithDefaults

`func NewClusterDataWithDefaults() *ClusterData`

NewClusterDataWithDefaults instantiates a new ClusterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerServiceUrl

`func (o *ClusterData) GetBrokerServiceUrl() string`

GetBrokerServiceUrl returns the BrokerServiceUrl field if non-nil, zero value otherwise.

### GetBrokerServiceUrlOk

`func (o *ClusterData) GetBrokerServiceUrlOk() (*string, bool)`

GetBrokerServiceUrlOk returns a tuple with the BrokerServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerServiceUrl

`func (o *ClusterData) SetBrokerServiceUrl(v string)`

SetBrokerServiceUrl sets BrokerServiceUrl field to given value.

### HasBrokerServiceUrl

`func (o *ClusterData) HasBrokerServiceUrl() bool`

HasBrokerServiceUrl returns a boolean if a field has been set.

### GetBrokerServiceUrlTls

`func (o *ClusterData) GetBrokerServiceUrlTls() string`

GetBrokerServiceUrlTls returns the BrokerServiceUrlTls field if non-nil, zero value otherwise.

### GetBrokerServiceUrlTlsOk

`func (o *ClusterData) GetBrokerServiceUrlTlsOk() (*string, bool)`

GetBrokerServiceUrlTlsOk returns a tuple with the BrokerServiceUrlTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerServiceUrlTls

`func (o *ClusterData) SetBrokerServiceUrlTls(v string)`

SetBrokerServiceUrlTls sets BrokerServiceUrlTls field to given value.

### HasBrokerServiceUrlTls

`func (o *ClusterData) HasBrokerServiceUrlTls() bool`

HasBrokerServiceUrlTls returns a boolean if a field has been set.

### GetPeerClusterNames

`func (o *ClusterData) GetPeerClusterNames() []string`

GetPeerClusterNames returns the PeerClusterNames field if non-nil, zero value otherwise.

### GetPeerClusterNamesOk

`func (o *ClusterData) GetPeerClusterNamesOk() (*[]string, bool)`

GetPeerClusterNamesOk returns a tuple with the PeerClusterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerClusterNames

`func (o *ClusterData) SetPeerClusterNames(v []string)`

SetPeerClusterNames sets PeerClusterNames field to given value.

### HasPeerClusterNames

`func (o *ClusterData) HasPeerClusterNames() bool`

HasPeerClusterNames returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *ClusterData) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *ClusterData) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *ClusterData) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *ClusterData) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetProxyServiceUrl

`func (o *ClusterData) GetProxyServiceUrl() string`

GetProxyServiceUrl returns the ProxyServiceUrl field if non-nil, zero value otherwise.

### GetProxyServiceUrlOk

`func (o *ClusterData) GetProxyServiceUrlOk() (*string, bool)`

GetProxyServiceUrlOk returns a tuple with the ProxyServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServiceUrl

`func (o *ClusterData) SetProxyServiceUrl(v string)`

SetProxyServiceUrl sets ProxyServiceUrl field to given value.

### HasProxyServiceUrl

`func (o *ClusterData) HasProxyServiceUrl() bool`

HasProxyServiceUrl returns a boolean if a field has been set.

### GetServiceUrl

`func (o *ClusterData) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ClusterData) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ClusterData) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ClusterData) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUrlTls

`func (o *ClusterData) GetServiceUrlTls() string`

GetServiceUrlTls returns the ServiceUrlTls field if non-nil, zero value otherwise.

### GetServiceUrlTlsOk

`func (o *ClusterData) GetServiceUrlTlsOk() (*string, bool)`

GetServiceUrlTlsOk returns a tuple with the ServiceUrlTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrlTls

`func (o *ClusterData) SetServiceUrlTls(v string)`

SetServiceUrlTls sets ServiceUrlTls field to given value.

### HasServiceUrlTls

`func (o *ClusterData) HasServiceUrlTls() bool`

HasServiceUrlTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


