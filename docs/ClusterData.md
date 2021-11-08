# ClusterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceUrl** | Pointer to **string** | The HTTP rest service URL (for admin operations) | [optional] 
**ServiceUrlTls** | Pointer to **string** | The HTTPS rest service URL (for admin operations) | [optional] 
**BrokerServiceUrl** | Pointer to **string** | The broker service url (for produce and consume operations) | [optional] 
**BrokerServiceUrlTls** | Pointer to **string** | The secured broker service url (for produce and consume operations) | [optional] 
**ProxyServiceUrl** | Pointer to **string** | Proxy-service url when client would like to connect to broker via proxy. | [optional] 
**AuthenticationPlugin** | Pointer to **string** | Authentication plugin when client would like to connect to cluster. | [optional] 
**AuthenticationParameters** | Pointer to **string** | Authentication parameters when client would like to connect to cluster. | [optional] 
**ProxyProtocol** | Pointer to **string** | protocol to decide type of proxy routing eg: SNI-routing | [optional] 
**PeerClusterNames** | Pointer to **[]string** | A set of peer cluster names | [optional] 
**BrokerClientTlsEnabled** | Pointer to **bool** | Enable TLS when talking with other brokers in the same cluster (admin operation) or different clusters (replication) | [optional] 
**TlsAllowInsecureConnection** | Pointer to **bool** | Allow TLS connections to servers whose certificate cannot be be verified to have been signed by a trusted certificate authority. | [optional] 
**BrokerClientTlsEnabledWithKeyStore** | Pointer to **bool** | Whether internal client use KeyStore type to authenticate with other Pulsar brokers | [optional] 
**BrokerClientTlsTrustStoreType** | Pointer to **string** | TLS TrustStore type configuration for internal client: JKS, PKCS12 used by the internal client to authenticate with Pulsar brokers | [optional] 
**BrokerClientTlsTrustStore** | Pointer to **string** | TLS TrustStore path for internal client used by the internal client to authenticate with Pulsar brokers | [optional] 
**BrokerClientTlsTrustStorePassword** | Pointer to **string** | TLS TrustStore password for internal client used by the internal client to authenticate with Pulsar brokers | [optional] 
**BrokerClientTrustCertsFilePath** | Pointer to **string** | Path for the trusted TLS certificate file for outgoing connection to a server (broker) | [optional] 
**ListenerName** | Pointer to **string** | listenerName when client would like to connect to cluster | [optional] 

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

### GetAuthenticationPlugin

`func (o *ClusterData) GetAuthenticationPlugin() string`

GetAuthenticationPlugin returns the AuthenticationPlugin field if non-nil, zero value otherwise.

### GetAuthenticationPluginOk

`func (o *ClusterData) GetAuthenticationPluginOk() (*string, bool)`

GetAuthenticationPluginOk returns a tuple with the AuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPlugin

`func (o *ClusterData) SetAuthenticationPlugin(v string)`

SetAuthenticationPlugin sets AuthenticationPlugin field to given value.

### HasAuthenticationPlugin

`func (o *ClusterData) HasAuthenticationPlugin() bool`

HasAuthenticationPlugin returns a boolean if a field has been set.

### GetAuthenticationParameters

`func (o *ClusterData) GetAuthenticationParameters() string`

GetAuthenticationParameters returns the AuthenticationParameters field if non-nil, zero value otherwise.

### GetAuthenticationParametersOk

`func (o *ClusterData) GetAuthenticationParametersOk() (*string, bool)`

GetAuthenticationParametersOk returns a tuple with the AuthenticationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationParameters

`func (o *ClusterData) SetAuthenticationParameters(v string)`

SetAuthenticationParameters sets AuthenticationParameters field to given value.

### HasAuthenticationParameters

`func (o *ClusterData) HasAuthenticationParameters() bool`

HasAuthenticationParameters returns a boolean if a field has been set.

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

### GetBrokerClientTlsEnabled

`func (o *ClusterData) GetBrokerClientTlsEnabled() bool`

GetBrokerClientTlsEnabled returns the BrokerClientTlsEnabled field if non-nil, zero value otherwise.

### GetBrokerClientTlsEnabledOk

`func (o *ClusterData) GetBrokerClientTlsEnabledOk() (*bool, bool)`

GetBrokerClientTlsEnabledOk returns a tuple with the BrokerClientTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTlsEnabled

`func (o *ClusterData) SetBrokerClientTlsEnabled(v bool)`

SetBrokerClientTlsEnabled sets BrokerClientTlsEnabled field to given value.

### HasBrokerClientTlsEnabled

`func (o *ClusterData) HasBrokerClientTlsEnabled() bool`

HasBrokerClientTlsEnabled returns a boolean if a field has been set.

### GetTlsAllowInsecureConnection

`func (o *ClusterData) GetTlsAllowInsecureConnection() bool`

GetTlsAllowInsecureConnection returns the TlsAllowInsecureConnection field if non-nil, zero value otherwise.

### GetTlsAllowInsecureConnectionOk

`func (o *ClusterData) GetTlsAllowInsecureConnectionOk() (*bool, bool)`

GetTlsAllowInsecureConnectionOk returns a tuple with the TlsAllowInsecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowInsecureConnection

`func (o *ClusterData) SetTlsAllowInsecureConnection(v bool)`

SetTlsAllowInsecureConnection sets TlsAllowInsecureConnection field to given value.

### HasTlsAllowInsecureConnection

`func (o *ClusterData) HasTlsAllowInsecureConnection() bool`

HasTlsAllowInsecureConnection returns a boolean if a field has been set.

### GetBrokerClientTlsEnabledWithKeyStore

`func (o *ClusterData) GetBrokerClientTlsEnabledWithKeyStore() bool`

GetBrokerClientTlsEnabledWithKeyStore returns the BrokerClientTlsEnabledWithKeyStore field if non-nil, zero value otherwise.

### GetBrokerClientTlsEnabledWithKeyStoreOk

`func (o *ClusterData) GetBrokerClientTlsEnabledWithKeyStoreOk() (*bool, bool)`

GetBrokerClientTlsEnabledWithKeyStoreOk returns a tuple with the BrokerClientTlsEnabledWithKeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTlsEnabledWithKeyStore

`func (o *ClusterData) SetBrokerClientTlsEnabledWithKeyStore(v bool)`

SetBrokerClientTlsEnabledWithKeyStore sets BrokerClientTlsEnabledWithKeyStore field to given value.

### HasBrokerClientTlsEnabledWithKeyStore

`func (o *ClusterData) HasBrokerClientTlsEnabledWithKeyStore() bool`

HasBrokerClientTlsEnabledWithKeyStore returns a boolean if a field has been set.

### GetBrokerClientTlsTrustStoreType

`func (o *ClusterData) GetBrokerClientTlsTrustStoreType() string`

GetBrokerClientTlsTrustStoreType returns the BrokerClientTlsTrustStoreType field if non-nil, zero value otherwise.

### GetBrokerClientTlsTrustStoreTypeOk

`func (o *ClusterData) GetBrokerClientTlsTrustStoreTypeOk() (*string, bool)`

GetBrokerClientTlsTrustStoreTypeOk returns a tuple with the BrokerClientTlsTrustStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTlsTrustStoreType

`func (o *ClusterData) SetBrokerClientTlsTrustStoreType(v string)`

SetBrokerClientTlsTrustStoreType sets BrokerClientTlsTrustStoreType field to given value.

### HasBrokerClientTlsTrustStoreType

`func (o *ClusterData) HasBrokerClientTlsTrustStoreType() bool`

HasBrokerClientTlsTrustStoreType returns a boolean if a field has been set.

### GetBrokerClientTlsTrustStore

`func (o *ClusterData) GetBrokerClientTlsTrustStore() string`

GetBrokerClientTlsTrustStore returns the BrokerClientTlsTrustStore field if non-nil, zero value otherwise.

### GetBrokerClientTlsTrustStoreOk

`func (o *ClusterData) GetBrokerClientTlsTrustStoreOk() (*string, bool)`

GetBrokerClientTlsTrustStoreOk returns a tuple with the BrokerClientTlsTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTlsTrustStore

`func (o *ClusterData) SetBrokerClientTlsTrustStore(v string)`

SetBrokerClientTlsTrustStore sets BrokerClientTlsTrustStore field to given value.

### HasBrokerClientTlsTrustStore

`func (o *ClusterData) HasBrokerClientTlsTrustStore() bool`

HasBrokerClientTlsTrustStore returns a boolean if a field has been set.

### GetBrokerClientTlsTrustStorePassword

`func (o *ClusterData) GetBrokerClientTlsTrustStorePassword() string`

GetBrokerClientTlsTrustStorePassword returns the BrokerClientTlsTrustStorePassword field if non-nil, zero value otherwise.

### GetBrokerClientTlsTrustStorePasswordOk

`func (o *ClusterData) GetBrokerClientTlsTrustStorePasswordOk() (*string, bool)`

GetBrokerClientTlsTrustStorePasswordOk returns a tuple with the BrokerClientTlsTrustStorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTlsTrustStorePassword

`func (o *ClusterData) SetBrokerClientTlsTrustStorePassword(v string)`

SetBrokerClientTlsTrustStorePassword sets BrokerClientTlsTrustStorePassword field to given value.

### HasBrokerClientTlsTrustStorePassword

`func (o *ClusterData) HasBrokerClientTlsTrustStorePassword() bool`

HasBrokerClientTlsTrustStorePassword returns a boolean if a field has been set.

### GetBrokerClientTrustCertsFilePath

`func (o *ClusterData) GetBrokerClientTrustCertsFilePath() string`

GetBrokerClientTrustCertsFilePath returns the BrokerClientTrustCertsFilePath field if non-nil, zero value otherwise.

### GetBrokerClientTrustCertsFilePathOk

`func (o *ClusterData) GetBrokerClientTrustCertsFilePathOk() (*string, bool)`

GetBrokerClientTrustCertsFilePathOk returns a tuple with the BrokerClientTrustCertsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTrustCertsFilePath

`func (o *ClusterData) SetBrokerClientTrustCertsFilePath(v string)`

SetBrokerClientTrustCertsFilePath sets BrokerClientTrustCertsFilePath field to given value.

### HasBrokerClientTrustCertsFilePath

`func (o *ClusterData) HasBrokerClientTrustCertsFilePath() bool`

HasBrokerClientTrustCertsFilePath returns a boolean if a field has been set.

### GetListenerName

`func (o *ClusterData) GetListenerName() string`

GetListenerName returns the ListenerName field if non-nil, zero value otherwise.

### GetListenerNameOk

`func (o *ClusterData) GetListenerNameOk() (*string, bool)`

GetListenerNameOk returns a tuple with the ListenerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerName

`func (o *ClusterData) SetListenerName(v string)`

SetListenerName sets ListenerName field to given value.

### HasListenerName

`func (o *ClusterData) HasListenerName() bool`

HasListenerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


