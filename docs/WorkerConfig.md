# WorkerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** |  | [optional] 
**WorkerHostname** | Pointer to **string** |  | [optional] 
**WorkerPort** | Pointer to **int32** |  | [optional] 
**WorkerPortTls** | Pointer to **int32** |  | [optional] 
**AuthenticateMetricsEndpoint** | Pointer to **bool** |  | [optional] 
**IncludeStandardPrometheusMetrics** | Pointer to **bool** |  | [optional] 
**JvmGCMetricsLoggerClassName** | Pointer to **string** |  | [optional] 
**NumHttpServerThreads** | Pointer to **int32** |  | [optional] 
**HttpRequestsLimitEnabled** | Pointer to **bool** |  | [optional] 
**HttpRequestsMaxPerSecond** | Pointer to **float64** |  | [optional] 
**ConfigurationStoreServers** | Pointer to **string** |  | [optional] 
**ZooKeeperSessionTimeoutMillis** | Pointer to **int64** |  | [optional] 
**ZooKeeperOperationTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**ZooKeeperCacheExpirySeconds** | Pointer to **int32** |  | [optional] 
**ConnectorsDirectory** | Pointer to **string** |  | [optional] 
**NarExtractionDirectory** | Pointer to **string** |  | [optional] 
**ValidateConnectorConfig** | Pointer to **bool** |  | [optional] 
**FunctionsDirectory** | Pointer to **string** |  | [optional] 
**FunctionMetadataTopicName** | Pointer to **string** |  | [optional] 
**UseCompactedMetadataTopic** | Pointer to **bool** |  | [optional] 
**FunctionWebServiceUrl** | Pointer to **string** |  | [optional] 
**PulsarServiceUrl** | Pointer to **string** |  | [optional] 
**PulsarWebServiceUrl** | Pointer to **string** |  | [optional] 
**ClusterCoordinationTopicName** | Pointer to **string** |  | [optional] 
**PulsarFunctionsNamespace** | Pointer to **string** |  | [optional] 
**PulsarFunctionsCluster** | Pointer to **string** |  | [optional] 
**NumFunctionPackageReplicas** | Pointer to **int32** |  | [optional] 
**DownloadDirectory** | Pointer to **string** |  | [optional] 
**StateStorageServiceUrl** | Pointer to **string** |  | [optional] 
**FunctionAssignmentTopicName** | Pointer to **string** |  | [optional] 
**SchedulerClassName** | Pointer to **string** |  | [optional] 
**FailureCheckFreqMs** | Pointer to **int64** |  | [optional] 
**RescheduleTimeoutMs** | Pointer to **int64** |  | [optional] 
**RebalanceCheckFreqSec** | Pointer to **int64** |  | [optional] 
**InitialBrokerReconnectMaxRetries** | Pointer to **int32** |  | [optional] 
**AssignmentWriteMaxRetries** | Pointer to **int32** |  | [optional] 
**InstanceLivenessCheckFreqMs** | Pointer to **int64** |  | [optional] 
**BrokerClientAuthenticationEnabled** | Pointer to **bool** |  | [optional] 
**BrokerClientAuthenticationPlugin** | Pointer to **string** |  | [optional] 
**BrokerClientAuthenticationParameters** | Pointer to **string** |  | [optional] 
**BookkeeperClientAuthenticationPlugin** | Pointer to **string** |  | [optional] 
**BookkeeperClientAuthenticationParametersName** | Pointer to **string** |  | [optional] 
**BookkeeperClientAuthenticationParameters** | Pointer to **string** |  | [optional] 
**TopicCompactionFrequencySec** | Pointer to **int64** |  | [optional] 
**TlsEnabled** | Pointer to **bool** |  | [optional] 
**TlsCertificateFilePath** | Pointer to **string** |  | [optional] 
**TlsKeyFilePath** | Pointer to **string** |  | [optional] 
**TlsTrustCertsFilePath** | Pointer to **string** |  | [optional] 
**TlsAllowInsecureConnection** | Pointer to **bool** |  | [optional] 
**TlsRequireTrustedClientCertOnConnect** | Pointer to **bool** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 
**TlsEnableHostnameVerification** | Pointer to **bool** |  | [optional] 
**TlsCertRefreshCheckDurationSec** | Pointer to **int64** |  | [optional] 
**AuthenticationEnabled** | Pointer to **bool** |  | [optional] 
**AuthenticationProviders** | Pointer to **[]string** |  | [optional] 
**AuthorizationEnabled** | Pointer to **bool** |  | [optional] 
**AuthorizationProvider** | Pointer to **string** |  | [optional] 
**SuperUserRoles** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**InitializedDlogMetadata** | Pointer to **bool** |  | [optional] 
**BrokerClientTrustCertsFilePath** | Pointer to **string** |  | [optional] 
**FunctionRuntimeFactoryClassName** | Pointer to **string** |  | [optional] 
**FunctionRuntimeFactoryConfigs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SecretsProviderConfiguratorClassName** | Pointer to **string** |  | [optional] 
**SecretsProviderConfiguratorConfig** | Pointer to **map[string]string** |  | [optional] 
**FunctionInstanceMinResources** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**FunctionInstanceMaxResources** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**FunctionInstanceResourceGranularities** | Pointer to [**Resources**](Resources.md) |  | [optional] 
**FunctionInstanceResourceChangeInLockStep** | Pointer to **bool** |  | [optional] 
**FunctionAuthProviderClassName** | Pointer to **string** |  | [optional] 
**RuntimeCustomizerClassName** | Pointer to **string** |  | [optional] 
**RuntimeCustomizerConfig** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MaxPendingAsyncRequests** | Pointer to **int32** |  | [optional] 
**ForwardSourceMessageProperty** | Pointer to **bool** |  | [optional] 
**FunctionsWorkerServiceNarPackage** | Pointer to **string** |  | [optional] 
**FunctionsWorkerServiceCustomConfigs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ExposeAdminClientEnabled** | Pointer to **bool** |  | [optional] 
**ThreadContainerFactory** | Pointer to [**ThreadContainerFactory**](ThreadContainerFactory.md) |  | [optional] 
**ProcessContainerFactory** | Pointer to [**ProcessContainerFactory**](ProcessContainerFactory.md) |  | [optional] 
**KubernetesContainerFactory** | Pointer to [**KubernetesContainerFactory**](KubernetesContainerFactory.md) |  | [optional] 
**ClientAuthenticationParameters** | Pointer to **string** |  | [optional] 
**ClientAuthenticationPlugin** | Pointer to **string** |  | [optional] 
**FunctionMetadataTopic** | Pointer to **string** |  | [optional] 
**ClusterCoordinationTopic** | Pointer to **string** |  | [optional] 
**FunctionAssignmentTopic** | Pointer to **string** |  | [optional] 
**TlsTrustChainBytes** | Pointer to **[]string** |  | [optional] 
**WorkerWebAddress** | Pointer to **string** |  | [optional] 
**WorkerWebAddressTls** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkerConfig

`func NewWorkerConfig() *WorkerConfig`

NewWorkerConfig instantiates a new WorkerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerConfigWithDefaults

`func NewWorkerConfigWithDefaults() *WorkerConfig`

NewWorkerConfigWithDefaults instantiates a new WorkerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *WorkerConfig) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *WorkerConfig) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *WorkerConfig) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *WorkerConfig) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkerHostname

`func (o *WorkerConfig) GetWorkerHostname() string`

GetWorkerHostname returns the WorkerHostname field if non-nil, zero value otherwise.

### GetWorkerHostnameOk

`func (o *WorkerConfig) GetWorkerHostnameOk() (*string, bool)`

GetWorkerHostnameOk returns a tuple with the WorkerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerHostname

`func (o *WorkerConfig) SetWorkerHostname(v string)`

SetWorkerHostname sets WorkerHostname field to given value.

### HasWorkerHostname

`func (o *WorkerConfig) HasWorkerHostname() bool`

HasWorkerHostname returns a boolean if a field has been set.

### GetWorkerPort

`func (o *WorkerConfig) GetWorkerPort() int32`

GetWorkerPort returns the WorkerPort field if non-nil, zero value otherwise.

### GetWorkerPortOk

`func (o *WorkerConfig) GetWorkerPortOk() (*int32, bool)`

GetWorkerPortOk returns a tuple with the WorkerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerPort

`func (o *WorkerConfig) SetWorkerPort(v int32)`

SetWorkerPort sets WorkerPort field to given value.

### HasWorkerPort

`func (o *WorkerConfig) HasWorkerPort() bool`

HasWorkerPort returns a boolean if a field has been set.

### GetWorkerPortTls

`func (o *WorkerConfig) GetWorkerPortTls() int32`

GetWorkerPortTls returns the WorkerPortTls field if non-nil, zero value otherwise.

### GetWorkerPortTlsOk

`func (o *WorkerConfig) GetWorkerPortTlsOk() (*int32, bool)`

GetWorkerPortTlsOk returns a tuple with the WorkerPortTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerPortTls

`func (o *WorkerConfig) SetWorkerPortTls(v int32)`

SetWorkerPortTls sets WorkerPortTls field to given value.

### HasWorkerPortTls

`func (o *WorkerConfig) HasWorkerPortTls() bool`

HasWorkerPortTls returns a boolean if a field has been set.

### GetAuthenticateMetricsEndpoint

`func (o *WorkerConfig) GetAuthenticateMetricsEndpoint() bool`

GetAuthenticateMetricsEndpoint returns the AuthenticateMetricsEndpoint field if non-nil, zero value otherwise.

### GetAuthenticateMetricsEndpointOk

`func (o *WorkerConfig) GetAuthenticateMetricsEndpointOk() (*bool, bool)`

GetAuthenticateMetricsEndpointOk returns a tuple with the AuthenticateMetricsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticateMetricsEndpoint

`func (o *WorkerConfig) SetAuthenticateMetricsEndpoint(v bool)`

SetAuthenticateMetricsEndpoint sets AuthenticateMetricsEndpoint field to given value.

### HasAuthenticateMetricsEndpoint

`func (o *WorkerConfig) HasAuthenticateMetricsEndpoint() bool`

HasAuthenticateMetricsEndpoint returns a boolean if a field has been set.

### GetIncludeStandardPrometheusMetrics

`func (o *WorkerConfig) GetIncludeStandardPrometheusMetrics() bool`

GetIncludeStandardPrometheusMetrics returns the IncludeStandardPrometheusMetrics field if non-nil, zero value otherwise.

### GetIncludeStandardPrometheusMetricsOk

`func (o *WorkerConfig) GetIncludeStandardPrometheusMetricsOk() (*bool, bool)`

GetIncludeStandardPrometheusMetricsOk returns a tuple with the IncludeStandardPrometheusMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStandardPrometheusMetrics

`func (o *WorkerConfig) SetIncludeStandardPrometheusMetrics(v bool)`

SetIncludeStandardPrometheusMetrics sets IncludeStandardPrometheusMetrics field to given value.

### HasIncludeStandardPrometheusMetrics

`func (o *WorkerConfig) HasIncludeStandardPrometheusMetrics() bool`

HasIncludeStandardPrometheusMetrics returns a boolean if a field has been set.

### GetJvmGCMetricsLoggerClassName

`func (o *WorkerConfig) GetJvmGCMetricsLoggerClassName() string`

GetJvmGCMetricsLoggerClassName returns the JvmGCMetricsLoggerClassName field if non-nil, zero value otherwise.

### GetJvmGCMetricsLoggerClassNameOk

`func (o *WorkerConfig) GetJvmGCMetricsLoggerClassNameOk() (*string, bool)`

GetJvmGCMetricsLoggerClassNameOk returns a tuple with the JvmGCMetricsLoggerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJvmGCMetricsLoggerClassName

`func (o *WorkerConfig) SetJvmGCMetricsLoggerClassName(v string)`

SetJvmGCMetricsLoggerClassName sets JvmGCMetricsLoggerClassName field to given value.

### HasJvmGCMetricsLoggerClassName

`func (o *WorkerConfig) HasJvmGCMetricsLoggerClassName() bool`

HasJvmGCMetricsLoggerClassName returns a boolean if a field has been set.

### GetNumHttpServerThreads

`func (o *WorkerConfig) GetNumHttpServerThreads() int32`

GetNumHttpServerThreads returns the NumHttpServerThreads field if non-nil, zero value otherwise.

### GetNumHttpServerThreadsOk

`func (o *WorkerConfig) GetNumHttpServerThreadsOk() (*int32, bool)`

GetNumHttpServerThreadsOk returns a tuple with the NumHttpServerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHttpServerThreads

`func (o *WorkerConfig) SetNumHttpServerThreads(v int32)`

SetNumHttpServerThreads sets NumHttpServerThreads field to given value.

### HasNumHttpServerThreads

`func (o *WorkerConfig) HasNumHttpServerThreads() bool`

HasNumHttpServerThreads returns a boolean if a field has been set.

### GetHttpRequestsLimitEnabled

`func (o *WorkerConfig) GetHttpRequestsLimitEnabled() bool`

GetHttpRequestsLimitEnabled returns the HttpRequestsLimitEnabled field if non-nil, zero value otherwise.

### GetHttpRequestsLimitEnabledOk

`func (o *WorkerConfig) GetHttpRequestsLimitEnabledOk() (*bool, bool)`

GetHttpRequestsLimitEnabledOk returns a tuple with the HttpRequestsLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestsLimitEnabled

`func (o *WorkerConfig) SetHttpRequestsLimitEnabled(v bool)`

SetHttpRequestsLimitEnabled sets HttpRequestsLimitEnabled field to given value.

### HasHttpRequestsLimitEnabled

`func (o *WorkerConfig) HasHttpRequestsLimitEnabled() bool`

HasHttpRequestsLimitEnabled returns a boolean if a field has been set.

### GetHttpRequestsMaxPerSecond

`func (o *WorkerConfig) GetHttpRequestsMaxPerSecond() float64`

GetHttpRequestsMaxPerSecond returns the HttpRequestsMaxPerSecond field if non-nil, zero value otherwise.

### GetHttpRequestsMaxPerSecondOk

`func (o *WorkerConfig) GetHttpRequestsMaxPerSecondOk() (*float64, bool)`

GetHttpRequestsMaxPerSecondOk returns a tuple with the HttpRequestsMaxPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestsMaxPerSecond

`func (o *WorkerConfig) SetHttpRequestsMaxPerSecond(v float64)`

SetHttpRequestsMaxPerSecond sets HttpRequestsMaxPerSecond field to given value.

### HasHttpRequestsMaxPerSecond

`func (o *WorkerConfig) HasHttpRequestsMaxPerSecond() bool`

HasHttpRequestsMaxPerSecond returns a boolean if a field has been set.

### GetConfigurationStoreServers

`func (o *WorkerConfig) GetConfigurationStoreServers() string`

GetConfigurationStoreServers returns the ConfigurationStoreServers field if non-nil, zero value otherwise.

### GetConfigurationStoreServersOk

`func (o *WorkerConfig) GetConfigurationStoreServersOk() (*string, bool)`

GetConfigurationStoreServersOk returns a tuple with the ConfigurationStoreServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStoreServers

`func (o *WorkerConfig) SetConfigurationStoreServers(v string)`

SetConfigurationStoreServers sets ConfigurationStoreServers field to given value.

### HasConfigurationStoreServers

`func (o *WorkerConfig) HasConfigurationStoreServers() bool`

HasConfigurationStoreServers returns a boolean if a field has been set.

### GetZooKeeperSessionTimeoutMillis

`func (o *WorkerConfig) GetZooKeeperSessionTimeoutMillis() int64`

GetZooKeeperSessionTimeoutMillis returns the ZooKeeperSessionTimeoutMillis field if non-nil, zero value otherwise.

### GetZooKeeperSessionTimeoutMillisOk

`func (o *WorkerConfig) GetZooKeeperSessionTimeoutMillisOk() (*int64, bool)`

GetZooKeeperSessionTimeoutMillisOk returns a tuple with the ZooKeeperSessionTimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZooKeeperSessionTimeoutMillis

`func (o *WorkerConfig) SetZooKeeperSessionTimeoutMillis(v int64)`

SetZooKeeperSessionTimeoutMillis sets ZooKeeperSessionTimeoutMillis field to given value.

### HasZooKeeperSessionTimeoutMillis

`func (o *WorkerConfig) HasZooKeeperSessionTimeoutMillis() bool`

HasZooKeeperSessionTimeoutMillis returns a boolean if a field has been set.

### GetZooKeeperOperationTimeoutSeconds

`func (o *WorkerConfig) GetZooKeeperOperationTimeoutSeconds() int32`

GetZooKeeperOperationTimeoutSeconds returns the ZooKeeperOperationTimeoutSeconds field if non-nil, zero value otherwise.

### GetZooKeeperOperationTimeoutSecondsOk

`func (o *WorkerConfig) GetZooKeeperOperationTimeoutSecondsOk() (*int32, bool)`

GetZooKeeperOperationTimeoutSecondsOk returns a tuple with the ZooKeeperOperationTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZooKeeperOperationTimeoutSeconds

`func (o *WorkerConfig) SetZooKeeperOperationTimeoutSeconds(v int32)`

SetZooKeeperOperationTimeoutSeconds sets ZooKeeperOperationTimeoutSeconds field to given value.

### HasZooKeeperOperationTimeoutSeconds

`func (o *WorkerConfig) HasZooKeeperOperationTimeoutSeconds() bool`

HasZooKeeperOperationTimeoutSeconds returns a boolean if a field has been set.

### GetZooKeeperCacheExpirySeconds

`func (o *WorkerConfig) GetZooKeeperCacheExpirySeconds() int32`

GetZooKeeperCacheExpirySeconds returns the ZooKeeperCacheExpirySeconds field if non-nil, zero value otherwise.

### GetZooKeeperCacheExpirySecondsOk

`func (o *WorkerConfig) GetZooKeeperCacheExpirySecondsOk() (*int32, bool)`

GetZooKeeperCacheExpirySecondsOk returns a tuple with the ZooKeeperCacheExpirySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZooKeeperCacheExpirySeconds

`func (o *WorkerConfig) SetZooKeeperCacheExpirySeconds(v int32)`

SetZooKeeperCacheExpirySeconds sets ZooKeeperCacheExpirySeconds field to given value.

### HasZooKeeperCacheExpirySeconds

`func (o *WorkerConfig) HasZooKeeperCacheExpirySeconds() bool`

HasZooKeeperCacheExpirySeconds returns a boolean if a field has been set.

### GetConnectorsDirectory

`func (o *WorkerConfig) GetConnectorsDirectory() string`

GetConnectorsDirectory returns the ConnectorsDirectory field if non-nil, zero value otherwise.

### GetConnectorsDirectoryOk

`func (o *WorkerConfig) GetConnectorsDirectoryOk() (*string, bool)`

GetConnectorsDirectoryOk returns a tuple with the ConnectorsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorsDirectory

`func (o *WorkerConfig) SetConnectorsDirectory(v string)`

SetConnectorsDirectory sets ConnectorsDirectory field to given value.

### HasConnectorsDirectory

`func (o *WorkerConfig) HasConnectorsDirectory() bool`

HasConnectorsDirectory returns a boolean if a field has been set.

### GetNarExtractionDirectory

`func (o *WorkerConfig) GetNarExtractionDirectory() string`

GetNarExtractionDirectory returns the NarExtractionDirectory field if non-nil, zero value otherwise.

### GetNarExtractionDirectoryOk

`func (o *WorkerConfig) GetNarExtractionDirectoryOk() (*string, bool)`

GetNarExtractionDirectoryOk returns a tuple with the NarExtractionDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarExtractionDirectory

`func (o *WorkerConfig) SetNarExtractionDirectory(v string)`

SetNarExtractionDirectory sets NarExtractionDirectory field to given value.

### HasNarExtractionDirectory

`func (o *WorkerConfig) HasNarExtractionDirectory() bool`

HasNarExtractionDirectory returns a boolean if a field has been set.

### GetValidateConnectorConfig

`func (o *WorkerConfig) GetValidateConnectorConfig() bool`

GetValidateConnectorConfig returns the ValidateConnectorConfig field if non-nil, zero value otherwise.

### GetValidateConnectorConfigOk

`func (o *WorkerConfig) GetValidateConnectorConfigOk() (*bool, bool)`

GetValidateConnectorConfigOk returns a tuple with the ValidateConnectorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateConnectorConfig

`func (o *WorkerConfig) SetValidateConnectorConfig(v bool)`

SetValidateConnectorConfig sets ValidateConnectorConfig field to given value.

### HasValidateConnectorConfig

`func (o *WorkerConfig) HasValidateConnectorConfig() bool`

HasValidateConnectorConfig returns a boolean if a field has been set.

### GetFunctionsDirectory

`func (o *WorkerConfig) GetFunctionsDirectory() string`

GetFunctionsDirectory returns the FunctionsDirectory field if non-nil, zero value otherwise.

### GetFunctionsDirectoryOk

`func (o *WorkerConfig) GetFunctionsDirectoryOk() (*string, bool)`

GetFunctionsDirectoryOk returns a tuple with the FunctionsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsDirectory

`func (o *WorkerConfig) SetFunctionsDirectory(v string)`

SetFunctionsDirectory sets FunctionsDirectory field to given value.

### HasFunctionsDirectory

`func (o *WorkerConfig) HasFunctionsDirectory() bool`

HasFunctionsDirectory returns a boolean if a field has been set.

### GetFunctionMetadataTopicName

`func (o *WorkerConfig) GetFunctionMetadataTopicName() string`

GetFunctionMetadataTopicName returns the FunctionMetadataTopicName field if non-nil, zero value otherwise.

### GetFunctionMetadataTopicNameOk

`func (o *WorkerConfig) GetFunctionMetadataTopicNameOk() (*string, bool)`

GetFunctionMetadataTopicNameOk returns a tuple with the FunctionMetadataTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMetadataTopicName

`func (o *WorkerConfig) SetFunctionMetadataTopicName(v string)`

SetFunctionMetadataTopicName sets FunctionMetadataTopicName field to given value.

### HasFunctionMetadataTopicName

`func (o *WorkerConfig) HasFunctionMetadataTopicName() bool`

HasFunctionMetadataTopicName returns a boolean if a field has been set.

### GetUseCompactedMetadataTopic

`func (o *WorkerConfig) GetUseCompactedMetadataTopic() bool`

GetUseCompactedMetadataTopic returns the UseCompactedMetadataTopic field if non-nil, zero value otherwise.

### GetUseCompactedMetadataTopicOk

`func (o *WorkerConfig) GetUseCompactedMetadataTopicOk() (*bool, bool)`

GetUseCompactedMetadataTopicOk returns a tuple with the UseCompactedMetadataTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCompactedMetadataTopic

`func (o *WorkerConfig) SetUseCompactedMetadataTopic(v bool)`

SetUseCompactedMetadataTopic sets UseCompactedMetadataTopic field to given value.

### HasUseCompactedMetadataTopic

`func (o *WorkerConfig) HasUseCompactedMetadataTopic() bool`

HasUseCompactedMetadataTopic returns a boolean if a field has been set.

### GetFunctionWebServiceUrl

`func (o *WorkerConfig) GetFunctionWebServiceUrl() string`

GetFunctionWebServiceUrl returns the FunctionWebServiceUrl field if non-nil, zero value otherwise.

### GetFunctionWebServiceUrlOk

`func (o *WorkerConfig) GetFunctionWebServiceUrlOk() (*string, bool)`

GetFunctionWebServiceUrlOk returns a tuple with the FunctionWebServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionWebServiceUrl

`func (o *WorkerConfig) SetFunctionWebServiceUrl(v string)`

SetFunctionWebServiceUrl sets FunctionWebServiceUrl field to given value.

### HasFunctionWebServiceUrl

`func (o *WorkerConfig) HasFunctionWebServiceUrl() bool`

HasFunctionWebServiceUrl returns a boolean if a field has been set.

### GetPulsarServiceUrl

`func (o *WorkerConfig) GetPulsarServiceUrl() string`

GetPulsarServiceUrl returns the PulsarServiceUrl field if non-nil, zero value otherwise.

### GetPulsarServiceUrlOk

`func (o *WorkerConfig) GetPulsarServiceUrlOk() (*string, bool)`

GetPulsarServiceUrlOk returns a tuple with the PulsarServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarServiceUrl

`func (o *WorkerConfig) SetPulsarServiceUrl(v string)`

SetPulsarServiceUrl sets PulsarServiceUrl field to given value.

### HasPulsarServiceUrl

`func (o *WorkerConfig) HasPulsarServiceUrl() bool`

HasPulsarServiceUrl returns a boolean if a field has been set.

### GetPulsarWebServiceUrl

`func (o *WorkerConfig) GetPulsarWebServiceUrl() string`

GetPulsarWebServiceUrl returns the PulsarWebServiceUrl field if non-nil, zero value otherwise.

### GetPulsarWebServiceUrlOk

`func (o *WorkerConfig) GetPulsarWebServiceUrlOk() (*string, bool)`

GetPulsarWebServiceUrlOk returns a tuple with the PulsarWebServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarWebServiceUrl

`func (o *WorkerConfig) SetPulsarWebServiceUrl(v string)`

SetPulsarWebServiceUrl sets PulsarWebServiceUrl field to given value.

### HasPulsarWebServiceUrl

`func (o *WorkerConfig) HasPulsarWebServiceUrl() bool`

HasPulsarWebServiceUrl returns a boolean if a field has been set.

### GetClusterCoordinationTopicName

`func (o *WorkerConfig) GetClusterCoordinationTopicName() string`

GetClusterCoordinationTopicName returns the ClusterCoordinationTopicName field if non-nil, zero value otherwise.

### GetClusterCoordinationTopicNameOk

`func (o *WorkerConfig) GetClusterCoordinationTopicNameOk() (*string, bool)`

GetClusterCoordinationTopicNameOk returns a tuple with the ClusterCoordinationTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCoordinationTopicName

`func (o *WorkerConfig) SetClusterCoordinationTopicName(v string)`

SetClusterCoordinationTopicName sets ClusterCoordinationTopicName field to given value.

### HasClusterCoordinationTopicName

`func (o *WorkerConfig) HasClusterCoordinationTopicName() bool`

HasClusterCoordinationTopicName returns a boolean if a field has been set.

### GetPulsarFunctionsNamespace

`func (o *WorkerConfig) GetPulsarFunctionsNamespace() string`

GetPulsarFunctionsNamespace returns the PulsarFunctionsNamespace field if non-nil, zero value otherwise.

### GetPulsarFunctionsNamespaceOk

`func (o *WorkerConfig) GetPulsarFunctionsNamespaceOk() (*string, bool)`

GetPulsarFunctionsNamespaceOk returns a tuple with the PulsarFunctionsNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarFunctionsNamespace

`func (o *WorkerConfig) SetPulsarFunctionsNamespace(v string)`

SetPulsarFunctionsNamespace sets PulsarFunctionsNamespace field to given value.

### HasPulsarFunctionsNamespace

`func (o *WorkerConfig) HasPulsarFunctionsNamespace() bool`

HasPulsarFunctionsNamespace returns a boolean if a field has been set.

### GetPulsarFunctionsCluster

`func (o *WorkerConfig) GetPulsarFunctionsCluster() string`

GetPulsarFunctionsCluster returns the PulsarFunctionsCluster field if non-nil, zero value otherwise.

### GetPulsarFunctionsClusterOk

`func (o *WorkerConfig) GetPulsarFunctionsClusterOk() (*string, bool)`

GetPulsarFunctionsClusterOk returns a tuple with the PulsarFunctionsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsarFunctionsCluster

`func (o *WorkerConfig) SetPulsarFunctionsCluster(v string)`

SetPulsarFunctionsCluster sets PulsarFunctionsCluster field to given value.

### HasPulsarFunctionsCluster

`func (o *WorkerConfig) HasPulsarFunctionsCluster() bool`

HasPulsarFunctionsCluster returns a boolean if a field has been set.

### GetNumFunctionPackageReplicas

`func (o *WorkerConfig) GetNumFunctionPackageReplicas() int32`

GetNumFunctionPackageReplicas returns the NumFunctionPackageReplicas field if non-nil, zero value otherwise.

### GetNumFunctionPackageReplicasOk

`func (o *WorkerConfig) GetNumFunctionPackageReplicasOk() (*int32, bool)`

GetNumFunctionPackageReplicasOk returns a tuple with the NumFunctionPackageReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFunctionPackageReplicas

`func (o *WorkerConfig) SetNumFunctionPackageReplicas(v int32)`

SetNumFunctionPackageReplicas sets NumFunctionPackageReplicas field to given value.

### HasNumFunctionPackageReplicas

`func (o *WorkerConfig) HasNumFunctionPackageReplicas() bool`

HasNumFunctionPackageReplicas returns a boolean if a field has been set.

### GetDownloadDirectory

`func (o *WorkerConfig) GetDownloadDirectory() string`

GetDownloadDirectory returns the DownloadDirectory field if non-nil, zero value otherwise.

### GetDownloadDirectoryOk

`func (o *WorkerConfig) GetDownloadDirectoryOk() (*string, bool)`

GetDownloadDirectoryOk returns a tuple with the DownloadDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDirectory

`func (o *WorkerConfig) SetDownloadDirectory(v string)`

SetDownloadDirectory sets DownloadDirectory field to given value.

### HasDownloadDirectory

`func (o *WorkerConfig) HasDownloadDirectory() bool`

HasDownloadDirectory returns a boolean if a field has been set.

### GetStateStorageServiceUrl

`func (o *WorkerConfig) GetStateStorageServiceUrl() string`

GetStateStorageServiceUrl returns the StateStorageServiceUrl field if non-nil, zero value otherwise.

### GetStateStorageServiceUrlOk

`func (o *WorkerConfig) GetStateStorageServiceUrlOk() (*string, bool)`

GetStateStorageServiceUrlOk returns a tuple with the StateStorageServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateStorageServiceUrl

`func (o *WorkerConfig) SetStateStorageServiceUrl(v string)`

SetStateStorageServiceUrl sets StateStorageServiceUrl field to given value.

### HasStateStorageServiceUrl

`func (o *WorkerConfig) HasStateStorageServiceUrl() bool`

HasStateStorageServiceUrl returns a boolean if a field has been set.

### GetFunctionAssignmentTopicName

`func (o *WorkerConfig) GetFunctionAssignmentTopicName() string`

GetFunctionAssignmentTopicName returns the FunctionAssignmentTopicName field if non-nil, zero value otherwise.

### GetFunctionAssignmentTopicNameOk

`func (o *WorkerConfig) GetFunctionAssignmentTopicNameOk() (*string, bool)`

GetFunctionAssignmentTopicNameOk returns a tuple with the FunctionAssignmentTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAssignmentTopicName

`func (o *WorkerConfig) SetFunctionAssignmentTopicName(v string)`

SetFunctionAssignmentTopicName sets FunctionAssignmentTopicName field to given value.

### HasFunctionAssignmentTopicName

`func (o *WorkerConfig) HasFunctionAssignmentTopicName() bool`

HasFunctionAssignmentTopicName returns a boolean if a field has been set.

### GetSchedulerClassName

`func (o *WorkerConfig) GetSchedulerClassName() string`

GetSchedulerClassName returns the SchedulerClassName field if non-nil, zero value otherwise.

### GetSchedulerClassNameOk

`func (o *WorkerConfig) GetSchedulerClassNameOk() (*string, bool)`

GetSchedulerClassNameOk returns a tuple with the SchedulerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerClassName

`func (o *WorkerConfig) SetSchedulerClassName(v string)`

SetSchedulerClassName sets SchedulerClassName field to given value.

### HasSchedulerClassName

`func (o *WorkerConfig) HasSchedulerClassName() bool`

HasSchedulerClassName returns a boolean if a field has been set.

### GetFailureCheckFreqMs

`func (o *WorkerConfig) GetFailureCheckFreqMs() int64`

GetFailureCheckFreqMs returns the FailureCheckFreqMs field if non-nil, zero value otherwise.

### GetFailureCheckFreqMsOk

`func (o *WorkerConfig) GetFailureCheckFreqMsOk() (*int64, bool)`

GetFailureCheckFreqMsOk returns a tuple with the FailureCheckFreqMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCheckFreqMs

`func (o *WorkerConfig) SetFailureCheckFreqMs(v int64)`

SetFailureCheckFreqMs sets FailureCheckFreqMs field to given value.

### HasFailureCheckFreqMs

`func (o *WorkerConfig) HasFailureCheckFreqMs() bool`

HasFailureCheckFreqMs returns a boolean if a field has been set.

### GetRescheduleTimeoutMs

`func (o *WorkerConfig) GetRescheduleTimeoutMs() int64`

GetRescheduleTimeoutMs returns the RescheduleTimeoutMs field if non-nil, zero value otherwise.

### GetRescheduleTimeoutMsOk

`func (o *WorkerConfig) GetRescheduleTimeoutMsOk() (*int64, bool)`

GetRescheduleTimeoutMsOk returns a tuple with the RescheduleTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescheduleTimeoutMs

`func (o *WorkerConfig) SetRescheduleTimeoutMs(v int64)`

SetRescheduleTimeoutMs sets RescheduleTimeoutMs field to given value.

### HasRescheduleTimeoutMs

`func (o *WorkerConfig) HasRescheduleTimeoutMs() bool`

HasRescheduleTimeoutMs returns a boolean if a field has been set.

### GetRebalanceCheckFreqSec

`func (o *WorkerConfig) GetRebalanceCheckFreqSec() int64`

GetRebalanceCheckFreqSec returns the RebalanceCheckFreqSec field if non-nil, zero value otherwise.

### GetRebalanceCheckFreqSecOk

`func (o *WorkerConfig) GetRebalanceCheckFreqSecOk() (*int64, bool)`

GetRebalanceCheckFreqSecOk returns a tuple with the RebalanceCheckFreqSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebalanceCheckFreqSec

`func (o *WorkerConfig) SetRebalanceCheckFreqSec(v int64)`

SetRebalanceCheckFreqSec sets RebalanceCheckFreqSec field to given value.

### HasRebalanceCheckFreqSec

`func (o *WorkerConfig) HasRebalanceCheckFreqSec() bool`

HasRebalanceCheckFreqSec returns a boolean if a field has been set.

### GetInitialBrokerReconnectMaxRetries

`func (o *WorkerConfig) GetInitialBrokerReconnectMaxRetries() int32`

GetInitialBrokerReconnectMaxRetries returns the InitialBrokerReconnectMaxRetries field if non-nil, zero value otherwise.

### GetInitialBrokerReconnectMaxRetriesOk

`func (o *WorkerConfig) GetInitialBrokerReconnectMaxRetriesOk() (*int32, bool)`

GetInitialBrokerReconnectMaxRetriesOk returns a tuple with the InitialBrokerReconnectMaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBrokerReconnectMaxRetries

`func (o *WorkerConfig) SetInitialBrokerReconnectMaxRetries(v int32)`

SetInitialBrokerReconnectMaxRetries sets InitialBrokerReconnectMaxRetries field to given value.

### HasInitialBrokerReconnectMaxRetries

`func (o *WorkerConfig) HasInitialBrokerReconnectMaxRetries() bool`

HasInitialBrokerReconnectMaxRetries returns a boolean if a field has been set.

### GetAssignmentWriteMaxRetries

`func (o *WorkerConfig) GetAssignmentWriteMaxRetries() int32`

GetAssignmentWriteMaxRetries returns the AssignmentWriteMaxRetries field if non-nil, zero value otherwise.

### GetAssignmentWriteMaxRetriesOk

`func (o *WorkerConfig) GetAssignmentWriteMaxRetriesOk() (*int32, bool)`

GetAssignmentWriteMaxRetriesOk returns a tuple with the AssignmentWriteMaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentWriteMaxRetries

`func (o *WorkerConfig) SetAssignmentWriteMaxRetries(v int32)`

SetAssignmentWriteMaxRetries sets AssignmentWriteMaxRetries field to given value.

### HasAssignmentWriteMaxRetries

`func (o *WorkerConfig) HasAssignmentWriteMaxRetries() bool`

HasAssignmentWriteMaxRetries returns a boolean if a field has been set.

### GetInstanceLivenessCheckFreqMs

`func (o *WorkerConfig) GetInstanceLivenessCheckFreqMs() int64`

GetInstanceLivenessCheckFreqMs returns the InstanceLivenessCheckFreqMs field if non-nil, zero value otherwise.

### GetInstanceLivenessCheckFreqMsOk

`func (o *WorkerConfig) GetInstanceLivenessCheckFreqMsOk() (*int64, bool)`

GetInstanceLivenessCheckFreqMsOk returns a tuple with the InstanceLivenessCheckFreqMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLivenessCheckFreqMs

`func (o *WorkerConfig) SetInstanceLivenessCheckFreqMs(v int64)`

SetInstanceLivenessCheckFreqMs sets InstanceLivenessCheckFreqMs field to given value.

### HasInstanceLivenessCheckFreqMs

`func (o *WorkerConfig) HasInstanceLivenessCheckFreqMs() bool`

HasInstanceLivenessCheckFreqMs returns a boolean if a field has been set.

### GetBrokerClientAuthenticationEnabled

`func (o *WorkerConfig) GetBrokerClientAuthenticationEnabled() bool`

GetBrokerClientAuthenticationEnabled returns the BrokerClientAuthenticationEnabled field if non-nil, zero value otherwise.

### GetBrokerClientAuthenticationEnabledOk

`func (o *WorkerConfig) GetBrokerClientAuthenticationEnabledOk() (*bool, bool)`

GetBrokerClientAuthenticationEnabledOk returns a tuple with the BrokerClientAuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientAuthenticationEnabled

`func (o *WorkerConfig) SetBrokerClientAuthenticationEnabled(v bool)`

SetBrokerClientAuthenticationEnabled sets BrokerClientAuthenticationEnabled field to given value.

### HasBrokerClientAuthenticationEnabled

`func (o *WorkerConfig) HasBrokerClientAuthenticationEnabled() bool`

HasBrokerClientAuthenticationEnabled returns a boolean if a field has been set.

### GetBrokerClientAuthenticationPlugin

`func (o *WorkerConfig) GetBrokerClientAuthenticationPlugin() string`

GetBrokerClientAuthenticationPlugin returns the BrokerClientAuthenticationPlugin field if non-nil, zero value otherwise.

### GetBrokerClientAuthenticationPluginOk

`func (o *WorkerConfig) GetBrokerClientAuthenticationPluginOk() (*string, bool)`

GetBrokerClientAuthenticationPluginOk returns a tuple with the BrokerClientAuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientAuthenticationPlugin

`func (o *WorkerConfig) SetBrokerClientAuthenticationPlugin(v string)`

SetBrokerClientAuthenticationPlugin sets BrokerClientAuthenticationPlugin field to given value.

### HasBrokerClientAuthenticationPlugin

`func (o *WorkerConfig) HasBrokerClientAuthenticationPlugin() bool`

HasBrokerClientAuthenticationPlugin returns a boolean if a field has been set.

### GetBrokerClientAuthenticationParameters

`func (o *WorkerConfig) GetBrokerClientAuthenticationParameters() string`

GetBrokerClientAuthenticationParameters returns the BrokerClientAuthenticationParameters field if non-nil, zero value otherwise.

### GetBrokerClientAuthenticationParametersOk

`func (o *WorkerConfig) GetBrokerClientAuthenticationParametersOk() (*string, bool)`

GetBrokerClientAuthenticationParametersOk returns a tuple with the BrokerClientAuthenticationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientAuthenticationParameters

`func (o *WorkerConfig) SetBrokerClientAuthenticationParameters(v string)`

SetBrokerClientAuthenticationParameters sets BrokerClientAuthenticationParameters field to given value.

### HasBrokerClientAuthenticationParameters

`func (o *WorkerConfig) HasBrokerClientAuthenticationParameters() bool`

HasBrokerClientAuthenticationParameters returns a boolean if a field has been set.

### GetBookkeeperClientAuthenticationPlugin

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationPlugin() string`

GetBookkeeperClientAuthenticationPlugin returns the BookkeeperClientAuthenticationPlugin field if non-nil, zero value otherwise.

### GetBookkeeperClientAuthenticationPluginOk

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationPluginOk() (*string, bool)`

GetBookkeeperClientAuthenticationPluginOk returns a tuple with the BookkeeperClientAuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperClientAuthenticationPlugin

`func (o *WorkerConfig) SetBookkeeperClientAuthenticationPlugin(v string)`

SetBookkeeperClientAuthenticationPlugin sets BookkeeperClientAuthenticationPlugin field to given value.

### HasBookkeeperClientAuthenticationPlugin

`func (o *WorkerConfig) HasBookkeeperClientAuthenticationPlugin() bool`

HasBookkeeperClientAuthenticationPlugin returns a boolean if a field has been set.

### GetBookkeeperClientAuthenticationParametersName

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationParametersName() string`

GetBookkeeperClientAuthenticationParametersName returns the BookkeeperClientAuthenticationParametersName field if non-nil, zero value otherwise.

### GetBookkeeperClientAuthenticationParametersNameOk

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationParametersNameOk() (*string, bool)`

GetBookkeeperClientAuthenticationParametersNameOk returns a tuple with the BookkeeperClientAuthenticationParametersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperClientAuthenticationParametersName

`func (o *WorkerConfig) SetBookkeeperClientAuthenticationParametersName(v string)`

SetBookkeeperClientAuthenticationParametersName sets BookkeeperClientAuthenticationParametersName field to given value.

### HasBookkeeperClientAuthenticationParametersName

`func (o *WorkerConfig) HasBookkeeperClientAuthenticationParametersName() bool`

HasBookkeeperClientAuthenticationParametersName returns a boolean if a field has been set.

### GetBookkeeperClientAuthenticationParameters

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationParameters() string`

GetBookkeeperClientAuthenticationParameters returns the BookkeeperClientAuthenticationParameters field if non-nil, zero value otherwise.

### GetBookkeeperClientAuthenticationParametersOk

`func (o *WorkerConfig) GetBookkeeperClientAuthenticationParametersOk() (*string, bool)`

GetBookkeeperClientAuthenticationParametersOk returns a tuple with the BookkeeperClientAuthenticationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookkeeperClientAuthenticationParameters

`func (o *WorkerConfig) SetBookkeeperClientAuthenticationParameters(v string)`

SetBookkeeperClientAuthenticationParameters sets BookkeeperClientAuthenticationParameters field to given value.

### HasBookkeeperClientAuthenticationParameters

`func (o *WorkerConfig) HasBookkeeperClientAuthenticationParameters() bool`

HasBookkeeperClientAuthenticationParameters returns a boolean if a field has been set.

### GetTopicCompactionFrequencySec

`func (o *WorkerConfig) GetTopicCompactionFrequencySec() int64`

GetTopicCompactionFrequencySec returns the TopicCompactionFrequencySec field if non-nil, zero value otherwise.

### GetTopicCompactionFrequencySecOk

`func (o *WorkerConfig) GetTopicCompactionFrequencySecOk() (*int64, bool)`

GetTopicCompactionFrequencySecOk returns a tuple with the TopicCompactionFrequencySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCompactionFrequencySec

`func (o *WorkerConfig) SetTopicCompactionFrequencySec(v int64)`

SetTopicCompactionFrequencySec sets TopicCompactionFrequencySec field to given value.

### HasTopicCompactionFrequencySec

`func (o *WorkerConfig) HasTopicCompactionFrequencySec() bool`

HasTopicCompactionFrequencySec returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *WorkerConfig) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *WorkerConfig) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *WorkerConfig) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *WorkerConfig) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetTlsCertificateFilePath

`func (o *WorkerConfig) GetTlsCertificateFilePath() string`

GetTlsCertificateFilePath returns the TlsCertificateFilePath field if non-nil, zero value otherwise.

### GetTlsCertificateFilePathOk

`func (o *WorkerConfig) GetTlsCertificateFilePathOk() (*string, bool)`

GetTlsCertificateFilePathOk returns a tuple with the TlsCertificateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificateFilePath

`func (o *WorkerConfig) SetTlsCertificateFilePath(v string)`

SetTlsCertificateFilePath sets TlsCertificateFilePath field to given value.

### HasTlsCertificateFilePath

`func (o *WorkerConfig) HasTlsCertificateFilePath() bool`

HasTlsCertificateFilePath returns a boolean if a field has been set.

### GetTlsKeyFilePath

`func (o *WorkerConfig) GetTlsKeyFilePath() string`

GetTlsKeyFilePath returns the TlsKeyFilePath field if non-nil, zero value otherwise.

### GetTlsKeyFilePathOk

`func (o *WorkerConfig) GetTlsKeyFilePathOk() (*string, bool)`

GetTlsKeyFilePathOk returns a tuple with the TlsKeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsKeyFilePath

`func (o *WorkerConfig) SetTlsKeyFilePath(v string)`

SetTlsKeyFilePath sets TlsKeyFilePath field to given value.

### HasTlsKeyFilePath

`func (o *WorkerConfig) HasTlsKeyFilePath() bool`

HasTlsKeyFilePath returns a boolean if a field has been set.

### GetTlsTrustCertsFilePath

`func (o *WorkerConfig) GetTlsTrustCertsFilePath() string`

GetTlsTrustCertsFilePath returns the TlsTrustCertsFilePath field if non-nil, zero value otherwise.

### GetTlsTrustCertsFilePathOk

`func (o *WorkerConfig) GetTlsTrustCertsFilePathOk() (*string, bool)`

GetTlsTrustCertsFilePathOk returns a tuple with the TlsTrustCertsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTrustCertsFilePath

`func (o *WorkerConfig) SetTlsTrustCertsFilePath(v string)`

SetTlsTrustCertsFilePath sets TlsTrustCertsFilePath field to given value.

### HasTlsTrustCertsFilePath

`func (o *WorkerConfig) HasTlsTrustCertsFilePath() bool`

HasTlsTrustCertsFilePath returns a boolean if a field has been set.

### GetTlsAllowInsecureConnection

`func (o *WorkerConfig) GetTlsAllowInsecureConnection() bool`

GetTlsAllowInsecureConnection returns the TlsAllowInsecureConnection field if non-nil, zero value otherwise.

### GetTlsAllowInsecureConnectionOk

`func (o *WorkerConfig) GetTlsAllowInsecureConnectionOk() (*bool, bool)`

GetTlsAllowInsecureConnectionOk returns a tuple with the TlsAllowInsecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowInsecureConnection

`func (o *WorkerConfig) SetTlsAllowInsecureConnection(v bool)`

SetTlsAllowInsecureConnection sets TlsAllowInsecureConnection field to given value.

### HasTlsAllowInsecureConnection

`func (o *WorkerConfig) HasTlsAllowInsecureConnection() bool`

HasTlsAllowInsecureConnection returns a boolean if a field has been set.

### GetTlsRequireTrustedClientCertOnConnect

`func (o *WorkerConfig) GetTlsRequireTrustedClientCertOnConnect() bool`

GetTlsRequireTrustedClientCertOnConnect returns the TlsRequireTrustedClientCertOnConnect field if non-nil, zero value otherwise.

### GetTlsRequireTrustedClientCertOnConnectOk

`func (o *WorkerConfig) GetTlsRequireTrustedClientCertOnConnectOk() (*bool, bool)`

GetTlsRequireTrustedClientCertOnConnectOk returns a tuple with the TlsRequireTrustedClientCertOnConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRequireTrustedClientCertOnConnect

`func (o *WorkerConfig) SetTlsRequireTrustedClientCertOnConnect(v bool)`

SetTlsRequireTrustedClientCertOnConnect sets TlsRequireTrustedClientCertOnConnect field to given value.

### HasTlsRequireTrustedClientCertOnConnect

`func (o *WorkerConfig) HasTlsRequireTrustedClientCertOnConnect() bool`

HasTlsRequireTrustedClientCertOnConnect returns a boolean if a field has been set.

### GetUseTls

`func (o *WorkerConfig) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *WorkerConfig) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *WorkerConfig) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *WorkerConfig) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetTlsEnableHostnameVerification

`func (o *WorkerConfig) GetTlsEnableHostnameVerification() bool`

GetTlsEnableHostnameVerification returns the TlsEnableHostnameVerification field if non-nil, zero value otherwise.

### GetTlsEnableHostnameVerificationOk

`func (o *WorkerConfig) GetTlsEnableHostnameVerificationOk() (*bool, bool)`

GetTlsEnableHostnameVerificationOk returns a tuple with the TlsEnableHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnableHostnameVerification

`func (o *WorkerConfig) SetTlsEnableHostnameVerification(v bool)`

SetTlsEnableHostnameVerification sets TlsEnableHostnameVerification field to given value.

### HasTlsEnableHostnameVerification

`func (o *WorkerConfig) HasTlsEnableHostnameVerification() bool`

HasTlsEnableHostnameVerification returns a boolean if a field has been set.

### GetTlsCertRefreshCheckDurationSec

`func (o *WorkerConfig) GetTlsCertRefreshCheckDurationSec() int64`

GetTlsCertRefreshCheckDurationSec returns the TlsCertRefreshCheckDurationSec field if non-nil, zero value otherwise.

### GetTlsCertRefreshCheckDurationSecOk

`func (o *WorkerConfig) GetTlsCertRefreshCheckDurationSecOk() (*int64, bool)`

GetTlsCertRefreshCheckDurationSecOk returns a tuple with the TlsCertRefreshCheckDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertRefreshCheckDurationSec

`func (o *WorkerConfig) SetTlsCertRefreshCheckDurationSec(v int64)`

SetTlsCertRefreshCheckDurationSec sets TlsCertRefreshCheckDurationSec field to given value.

### HasTlsCertRefreshCheckDurationSec

`func (o *WorkerConfig) HasTlsCertRefreshCheckDurationSec() bool`

HasTlsCertRefreshCheckDurationSec returns a boolean if a field has been set.

### GetAuthenticationEnabled

`func (o *WorkerConfig) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *WorkerConfig) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *WorkerConfig) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *WorkerConfig) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthenticationProviders

`func (o *WorkerConfig) GetAuthenticationProviders() []string`

GetAuthenticationProviders returns the AuthenticationProviders field if non-nil, zero value otherwise.

### GetAuthenticationProvidersOk

`func (o *WorkerConfig) GetAuthenticationProvidersOk() (*[]string, bool)`

GetAuthenticationProvidersOk returns a tuple with the AuthenticationProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviders

`func (o *WorkerConfig) SetAuthenticationProviders(v []string)`

SetAuthenticationProviders sets AuthenticationProviders field to given value.

### HasAuthenticationProviders

`func (o *WorkerConfig) HasAuthenticationProviders() bool`

HasAuthenticationProviders returns a boolean if a field has been set.

### GetAuthorizationEnabled

`func (o *WorkerConfig) GetAuthorizationEnabled() bool`

GetAuthorizationEnabled returns the AuthorizationEnabled field if non-nil, zero value otherwise.

### GetAuthorizationEnabledOk

`func (o *WorkerConfig) GetAuthorizationEnabledOk() (*bool, bool)`

GetAuthorizationEnabledOk returns a tuple with the AuthorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEnabled

`func (o *WorkerConfig) SetAuthorizationEnabled(v bool)`

SetAuthorizationEnabled sets AuthorizationEnabled field to given value.

### HasAuthorizationEnabled

`func (o *WorkerConfig) HasAuthorizationEnabled() bool`

HasAuthorizationEnabled returns a boolean if a field has been set.

### GetAuthorizationProvider

`func (o *WorkerConfig) GetAuthorizationProvider() string`

GetAuthorizationProvider returns the AuthorizationProvider field if non-nil, zero value otherwise.

### GetAuthorizationProviderOk

`func (o *WorkerConfig) GetAuthorizationProviderOk() (*string, bool)`

GetAuthorizationProviderOk returns a tuple with the AuthorizationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProvider

`func (o *WorkerConfig) SetAuthorizationProvider(v string)`

SetAuthorizationProvider sets AuthorizationProvider field to given value.

### HasAuthorizationProvider

`func (o *WorkerConfig) HasAuthorizationProvider() bool`

HasAuthorizationProvider returns a boolean if a field has been set.

### GetSuperUserRoles

`func (o *WorkerConfig) GetSuperUserRoles() []string`

GetSuperUserRoles returns the SuperUserRoles field if non-nil, zero value otherwise.

### GetSuperUserRolesOk

`func (o *WorkerConfig) GetSuperUserRolesOk() (*[]string, bool)`

GetSuperUserRolesOk returns a tuple with the SuperUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperUserRoles

`func (o *WorkerConfig) SetSuperUserRoles(v []string)`

SetSuperUserRoles sets SuperUserRoles field to given value.

### HasSuperUserRoles

`func (o *WorkerConfig) HasSuperUserRoles() bool`

HasSuperUserRoles returns a boolean if a field has been set.

### GetProperties

`func (o *WorkerConfig) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkerConfig) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkerConfig) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkerConfig) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetInitializedDlogMetadata

`func (o *WorkerConfig) GetInitializedDlogMetadata() bool`

GetInitializedDlogMetadata returns the InitializedDlogMetadata field if non-nil, zero value otherwise.

### GetInitializedDlogMetadataOk

`func (o *WorkerConfig) GetInitializedDlogMetadataOk() (*bool, bool)`

GetInitializedDlogMetadataOk returns a tuple with the InitializedDlogMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializedDlogMetadata

`func (o *WorkerConfig) SetInitializedDlogMetadata(v bool)`

SetInitializedDlogMetadata sets InitializedDlogMetadata field to given value.

### HasInitializedDlogMetadata

`func (o *WorkerConfig) HasInitializedDlogMetadata() bool`

HasInitializedDlogMetadata returns a boolean if a field has been set.

### GetBrokerClientTrustCertsFilePath

`func (o *WorkerConfig) GetBrokerClientTrustCertsFilePath() string`

GetBrokerClientTrustCertsFilePath returns the BrokerClientTrustCertsFilePath field if non-nil, zero value otherwise.

### GetBrokerClientTrustCertsFilePathOk

`func (o *WorkerConfig) GetBrokerClientTrustCertsFilePathOk() (*string, bool)`

GetBrokerClientTrustCertsFilePathOk returns a tuple with the BrokerClientTrustCertsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerClientTrustCertsFilePath

`func (o *WorkerConfig) SetBrokerClientTrustCertsFilePath(v string)`

SetBrokerClientTrustCertsFilePath sets BrokerClientTrustCertsFilePath field to given value.

### HasBrokerClientTrustCertsFilePath

`func (o *WorkerConfig) HasBrokerClientTrustCertsFilePath() bool`

HasBrokerClientTrustCertsFilePath returns a boolean if a field has been set.

### GetFunctionRuntimeFactoryClassName

`func (o *WorkerConfig) GetFunctionRuntimeFactoryClassName() string`

GetFunctionRuntimeFactoryClassName returns the FunctionRuntimeFactoryClassName field if non-nil, zero value otherwise.

### GetFunctionRuntimeFactoryClassNameOk

`func (o *WorkerConfig) GetFunctionRuntimeFactoryClassNameOk() (*string, bool)`

GetFunctionRuntimeFactoryClassNameOk returns a tuple with the FunctionRuntimeFactoryClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionRuntimeFactoryClassName

`func (o *WorkerConfig) SetFunctionRuntimeFactoryClassName(v string)`

SetFunctionRuntimeFactoryClassName sets FunctionRuntimeFactoryClassName field to given value.

### HasFunctionRuntimeFactoryClassName

`func (o *WorkerConfig) HasFunctionRuntimeFactoryClassName() bool`

HasFunctionRuntimeFactoryClassName returns a boolean if a field has been set.

### GetFunctionRuntimeFactoryConfigs

`func (o *WorkerConfig) GetFunctionRuntimeFactoryConfigs() map[string]map[string]interface{}`

GetFunctionRuntimeFactoryConfigs returns the FunctionRuntimeFactoryConfigs field if non-nil, zero value otherwise.

### GetFunctionRuntimeFactoryConfigsOk

`func (o *WorkerConfig) GetFunctionRuntimeFactoryConfigsOk() (*map[string]map[string]interface{}, bool)`

GetFunctionRuntimeFactoryConfigsOk returns a tuple with the FunctionRuntimeFactoryConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionRuntimeFactoryConfigs

`func (o *WorkerConfig) SetFunctionRuntimeFactoryConfigs(v map[string]map[string]interface{})`

SetFunctionRuntimeFactoryConfigs sets FunctionRuntimeFactoryConfigs field to given value.

### HasFunctionRuntimeFactoryConfigs

`func (o *WorkerConfig) HasFunctionRuntimeFactoryConfigs() bool`

HasFunctionRuntimeFactoryConfigs returns a boolean if a field has been set.

### GetSecretsProviderConfiguratorClassName

`func (o *WorkerConfig) GetSecretsProviderConfiguratorClassName() string`

GetSecretsProviderConfiguratorClassName returns the SecretsProviderConfiguratorClassName field if non-nil, zero value otherwise.

### GetSecretsProviderConfiguratorClassNameOk

`func (o *WorkerConfig) GetSecretsProviderConfiguratorClassNameOk() (*string, bool)`

GetSecretsProviderConfiguratorClassNameOk returns a tuple with the SecretsProviderConfiguratorClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsProviderConfiguratorClassName

`func (o *WorkerConfig) SetSecretsProviderConfiguratorClassName(v string)`

SetSecretsProviderConfiguratorClassName sets SecretsProviderConfiguratorClassName field to given value.

### HasSecretsProviderConfiguratorClassName

`func (o *WorkerConfig) HasSecretsProviderConfiguratorClassName() bool`

HasSecretsProviderConfiguratorClassName returns a boolean if a field has been set.

### GetSecretsProviderConfiguratorConfig

`func (o *WorkerConfig) GetSecretsProviderConfiguratorConfig() map[string]string`

GetSecretsProviderConfiguratorConfig returns the SecretsProviderConfiguratorConfig field if non-nil, zero value otherwise.

### GetSecretsProviderConfiguratorConfigOk

`func (o *WorkerConfig) GetSecretsProviderConfiguratorConfigOk() (*map[string]string, bool)`

GetSecretsProviderConfiguratorConfigOk returns a tuple with the SecretsProviderConfiguratorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsProviderConfiguratorConfig

`func (o *WorkerConfig) SetSecretsProviderConfiguratorConfig(v map[string]string)`

SetSecretsProviderConfiguratorConfig sets SecretsProviderConfiguratorConfig field to given value.

### HasSecretsProviderConfiguratorConfig

`func (o *WorkerConfig) HasSecretsProviderConfiguratorConfig() bool`

HasSecretsProviderConfiguratorConfig returns a boolean if a field has been set.

### GetFunctionInstanceMinResources

`func (o *WorkerConfig) GetFunctionInstanceMinResources() Resources`

GetFunctionInstanceMinResources returns the FunctionInstanceMinResources field if non-nil, zero value otherwise.

### GetFunctionInstanceMinResourcesOk

`func (o *WorkerConfig) GetFunctionInstanceMinResourcesOk() (*Resources, bool)`

GetFunctionInstanceMinResourcesOk returns a tuple with the FunctionInstanceMinResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionInstanceMinResources

`func (o *WorkerConfig) SetFunctionInstanceMinResources(v Resources)`

SetFunctionInstanceMinResources sets FunctionInstanceMinResources field to given value.

### HasFunctionInstanceMinResources

`func (o *WorkerConfig) HasFunctionInstanceMinResources() bool`

HasFunctionInstanceMinResources returns a boolean if a field has been set.

### GetFunctionInstanceMaxResources

`func (o *WorkerConfig) GetFunctionInstanceMaxResources() Resources`

GetFunctionInstanceMaxResources returns the FunctionInstanceMaxResources field if non-nil, zero value otherwise.

### GetFunctionInstanceMaxResourcesOk

`func (o *WorkerConfig) GetFunctionInstanceMaxResourcesOk() (*Resources, bool)`

GetFunctionInstanceMaxResourcesOk returns a tuple with the FunctionInstanceMaxResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionInstanceMaxResources

`func (o *WorkerConfig) SetFunctionInstanceMaxResources(v Resources)`

SetFunctionInstanceMaxResources sets FunctionInstanceMaxResources field to given value.

### HasFunctionInstanceMaxResources

`func (o *WorkerConfig) HasFunctionInstanceMaxResources() bool`

HasFunctionInstanceMaxResources returns a boolean if a field has been set.

### GetFunctionInstanceResourceGranularities

`func (o *WorkerConfig) GetFunctionInstanceResourceGranularities() Resources`

GetFunctionInstanceResourceGranularities returns the FunctionInstanceResourceGranularities field if non-nil, zero value otherwise.

### GetFunctionInstanceResourceGranularitiesOk

`func (o *WorkerConfig) GetFunctionInstanceResourceGranularitiesOk() (*Resources, bool)`

GetFunctionInstanceResourceGranularitiesOk returns a tuple with the FunctionInstanceResourceGranularities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionInstanceResourceGranularities

`func (o *WorkerConfig) SetFunctionInstanceResourceGranularities(v Resources)`

SetFunctionInstanceResourceGranularities sets FunctionInstanceResourceGranularities field to given value.

### HasFunctionInstanceResourceGranularities

`func (o *WorkerConfig) HasFunctionInstanceResourceGranularities() bool`

HasFunctionInstanceResourceGranularities returns a boolean if a field has been set.

### GetFunctionInstanceResourceChangeInLockStep

`func (o *WorkerConfig) GetFunctionInstanceResourceChangeInLockStep() bool`

GetFunctionInstanceResourceChangeInLockStep returns the FunctionInstanceResourceChangeInLockStep field if non-nil, zero value otherwise.

### GetFunctionInstanceResourceChangeInLockStepOk

`func (o *WorkerConfig) GetFunctionInstanceResourceChangeInLockStepOk() (*bool, bool)`

GetFunctionInstanceResourceChangeInLockStepOk returns a tuple with the FunctionInstanceResourceChangeInLockStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionInstanceResourceChangeInLockStep

`func (o *WorkerConfig) SetFunctionInstanceResourceChangeInLockStep(v bool)`

SetFunctionInstanceResourceChangeInLockStep sets FunctionInstanceResourceChangeInLockStep field to given value.

### HasFunctionInstanceResourceChangeInLockStep

`func (o *WorkerConfig) HasFunctionInstanceResourceChangeInLockStep() bool`

HasFunctionInstanceResourceChangeInLockStep returns a boolean if a field has been set.

### GetFunctionAuthProviderClassName

`func (o *WorkerConfig) GetFunctionAuthProviderClassName() string`

GetFunctionAuthProviderClassName returns the FunctionAuthProviderClassName field if non-nil, zero value otherwise.

### GetFunctionAuthProviderClassNameOk

`func (o *WorkerConfig) GetFunctionAuthProviderClassNameOk() (*string, bool)`

GetFunctionAuthProviderClassNameOk returns a tuple with the FunctionAuthProviderClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAuthProviderClassName

`func (o *WorkerConfig) SetFunctionAuthProviderClassName(v string)`

SetFunctionAuthProviderClassName sets FunctionAuthProviderClassName field to given value.

### HasFunctionAuthProviderClassName

`func (o *WorkerConfig) HasFunctionAuthProviderClassName() bool`

HasFunctionAuthProviderClassName returns a boolean if a field has been set.

### GetRuntimeCustomizerClassName

`func (o *WorkerConfig) GetRuntimeCustomizerClassName() string`

GetRuntimeCustomizerClassName returns the RuntimeCustomizerClassName field if non-nil, zero value otherwise.

### GetRuntimeCustomizerClassNameOk

`func (o *WorkerConfig) GetRuntimeCustomizerClassNameOk() (*string, bool)`

GetRuntimeCustomizerClassNameOk returns a tuple with the RuntimeCustomizerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCustomizerClassName

`func (o *WorkerConfig) SetRuntimeCustomizerClassName(v string)`

SetRuntimeCustomizerClassName sets RuntimeCustomizerClassName field to given value.

### HasRuntimeCustomizerClassName

`func (o *WorkerConfig) HasRuntimeCustomizerClassName() bool`

HasRuntimeCustomizerClassName returns a boolean if a field has been set.

### GetRuntimeCustomizerConfig

`func (o *WorkerConfig) GetRuntimeCustomizerConfig() map[string]map[string]interface{}`

GetRuntimeCustomizerConfig returns the RuntimeCustomizerConfig field if non-nil, zero value otherwise.

### GetRuntimeCustomizerConfigOk

`func (o *WorkerConfig) GetRuntimeCustomizerConfigOk() (*map[string]map[string]interface{}, bool)`

GetRuntimeCustomizerConfigOk returns a tuple with the RuntimeCustomizerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCustomizerConfig

`func (o *WorkerConfig) SetRuntimeCustomizerConfig(v map[string]map[string]interface{})`

SetRuntimeCustomizerConfig sets RuntimeCustomizerConfig field to given value.

### HasRuntimeCustomizerConfig

`func (o *WorkerConfig) HasRuntimeCustomizerConfig() bool`

HasRuntimeCustomizerConfig returns a boolean if a field has been set.

### GetMaxPendingAsyncRequests

`func (o *WorkerConfig) GetMaxPendingAsyncRequests() int32`

GetMaxPendingAsyncRequests returns the MaxPendingAsyncRequests field if non-nil, zero value otherwise.

### GetMaxPendingAsyncRequestsOk

`func (o *WorkerConfig) GetMaxPendingAsyncRequestsOk() (*int32, bool)`

GetMaxPendingAsyncRequestsOk returns a tuple with the MaxPendingAsyncRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPendingAsyncRequests

`func (o *WorkerConfig) SetMaxPendingAsyncRequests(v int32)`

SetMaxPendingAsyncRequests sets MaxPendingAsyncRequests field to given value.

### HasMaxPendingAsyncRequests

`func (o *WorkerConfig) HasMaxPendingAsyncRequests() bool`

HasMaxPendingAsyncRequests returns a boolean if a field has been set.

### GetForwardSourceMessageProperty

`func (o *WorkerConfig) GetForwardSourceMessageProperty() bool`

GetForwardSourceMessageProperty returns the ForwardSourceMessageProperty field if non-nil, zero value otherwise.

### GetForwardSourceMessagePropertyOk

`func (o *WorkerConfig) GetForwardSourceMessagePropertyOk() (*bool, bool)`

GetForwardSourceMessagePropertyOk returns a tuple with the ForwardSourceMessageProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardSourceMessageProperty

`func (o *WorkerConfig) SetForwardSourceMessageProperty(v bool)`

SetForwardSourceMessageProperty sets ForwardSourceMessageProperty field to given value.

### HasForwardSourceMessageProperty

`func (o *WorkerConfig) HasForwardSourceMessageProperty() bool`

HasForwardSourceMessageProperty returns a boolean if a field has been set.

### GetFunctionsWorkerServiceNarPackage

`func (o *WorkerConfig) GetFunctionsWorkerServiceNarPackage() string`

GetFunctionsWorkerServiceNarPackage returns the FunctionsWorkerServiceNarPackage field if non-nil, zero value otherwise.

### GetFunctionsWorkerServiceNarPackageOk

`func (o *WorkerConfig) GetFunctionsWorkerServiceNarPackageOk() (*string, bool)`

GetFunctionsWorkerServiceNarPackageOk returns a tuple with the FunctionsWorkerServiceNarPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsWorkerServiceNarPackage

`func (o *WorkerConfig) SetFunctionsWorkerServiceNarPackage(v string)`

SetFunctionsWorkerServiceNarPackage sets FunctionsWorkerServiceNarPackage field to given value.

### HasFunctionsWorkerServiceNarPackage

`func (o *WorkerConfig) HasFunctionsWorkerServiceNarPackage() bool`

HasFunctionsWorkerServiceNarPackage returns a boolean if a field has been set.

### GetFunctionsWorkerServiceCustomConfigs

`func (o *WorkerConfig) GetFunctionsWorkerServiceCustomConfigs() map[string]map[string]interface{}`

GetFunctionsWorkerServiceCustomConfigs returns the FunctionsWorkerServiceCustomConfigs field if non-nil, zero value otherwise.

### GetFunctionsWorkerServiceCustomConfigsOk

`func (o *WorkerConfig) GetFunctionsWorkerServiceCustomConfigsOk() (*map[string]map[string]interface{}, bool)`

GetFunctionsWorkerServiceCustomConfigsOk returns a tuple with the FunctionsWorkerServiceCustomConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsWorkerServiceCustomConfigs

`func (o *WorkerConfig) SetFunctionsWorkerServiceCustomConfigs(v map[string]map[string]interface{})`

SetFunctionsWorkerServiceCustomConfigs sets FunctionsWorkerServiceCustomConfigs field to given value.

### HasFunctionsWorkerServiceCustomConfigs

`func (o *WorkerConfig) HasFunctionsWorkerServiceCustomConfigs() bool`

HasFunctionsWorkerServiceCustomConfigs returns a boolean if a field has been set.

### GetExposeAdminClientEnabled

`func (o *WorkerConfig) GetExposeAdminClientEnabled() bool`

GetExposeAdminClientEnabled returns the ExposeAdminClientEnabled field if non-nil, zero value otherwise.

### GetExposeAdminClientEnabledOk

`func (o *WorkerConfig) GetExposeAdminClientEnabledOk() (*bool, bool)`

GetExposeAdminClientEnabledOk returns a tuple with the ExposeAdminClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeAdminClientEnabled

`func (o *WorkerConfig) SetExposeAdminClientEnabled(v bool)`

SetExposeAdminClientEnabled sets ExposeAdminClientEnabled field to given value.

### HasExposeAdminClientEnabled

`func (o *WorkerConfig) HasExposeAdminClientEnabled() bool`

HasExposeAdminClientEnabled returns a boolean if a field has been set.

### GetThreadContainerFactory

`func (o *WorkerConfig) GetThreadContainerFactory() ThreadContainerFactory`

GetThreadContainerFactory returns the ThreadContainerFactory field if non-nil, zero value otherwise.

### GetThreadContainerFactoryOk

`func (o *WorkerConfig) GetThreadContainerFactoryOk() (*ThreadContainerFactory, bool)`

GetThreadContainerFactoryOk returns a tuple with the ThreadContainerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadContainerFactory

`func (o *WorkerConfig) SetThreadContainerFactory(v ThreadContainerFactory)`

SetThreadContainerFactory sets ThreadContainerFactory field to given value.

### HasThreadContainerFactory

`func (o *WorkerConfig) HasThreadContainerFactory() bool`

HasThreadContainerFactory returns a boolean if a field has been set.

### GetProcessContainerFactory

`func (o *WorkerConfig) GetProcessContainerFactory() ProcessContainerFactory`

GetProcessContainerFactory returns the ProcessContainerFactory field if non-nil, zero value otherwise.

### GetProcessContainerFactoryOk

`func (o *WorkerConfig) GetProcessContainerFactoryOk() (*ProcessContainerFactory, bool)`

GetProcessContainerFactoryOk returns a tuple with the ProcessContainerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessContainerFactory

`func (o *WorkerConfig) SetProcessContainerFactory(v ProcessContainerFactory)`

SetProcessContainerFactory sets ProcessContainerFactory field to given value.

### HasProcessContainerFactory

`func (o *WorkerConfig) HasProcessContainerFactory() bool`

HasProcessContainerFactory returns a boolean if a field has been set.

### GetKubernetesContainerFactory

`func (o *WorkerConfig) GetKubernetesContainerFactory() KubernetesContainerFactory`

GetKubernetesContainerFactory returns the KubernetesContainerFactory field if non-nil, zero value otherwise.

### GetKubernetesContainerFactoryOk

`func (o *WorkerConfig) GetKubernetesContainerFactoryOk() (*KubernetesContainerFactory, bool)`

GetKubernetesContainerFactoryOk returns a tuple with the KubernetesContainerFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesContainerFactory

`func (o *WorkerConfig) SetKubernetesContainerFactory(v KubernetesContainerFactory)`

SetKubernetesContainerFactory sets KubernetesContainerFactory field to given value.

### HasKubernetesContainerFactory

`func (o *WorkerConfig) HasKubernetesContainerFactory() bool`

HasKubernetesContainerFactory returns a boolean if a field has been set.

### GetClientAuthenticationParameters

`func (o *WorkerConfig) GetClientAuthenticationParameters() string`

GetClientAuthenticationParameters returns the ClientAuthenticationParameters field if non-nil, zero value otherwise.

### GetClientAuthenticationParametersOk

`func (o *WorkerConfig) GetClientAuthenticationParametersOk() (*string, bool)`

GetClientAuthenticationParametersOk returns a tuple with the ClientAuthenticationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationParameters

`func (o *WorkerConfig) SetClientAuthenticationParameters(v string)`

SetClientAuthenticationParameters sets ClientAuthenticationParameters field to given value.

### HasClientAuthenticationParameters

`func (o *WorkerConfig) HasClientAuthenticationParameters() bool`

HasClientAuthenticationParameters returns a boolean if a field has been set.

### GetClientAuthenticationPlugin

`func (o *WorkerConfig) GetClientAuthenticationPlugin() string`

GetClientAuthenticationPlugin returns the ClientAuthenticationPlugin field if non-nil, zero value otherwise.

### GetClientAuthenticationPluginOk

`func (o *WorkerConfig) GetClientAuthenticationPluginOk() (*string, bool)`

GetClientAuthenticationPluginOk returns a tuple with the ClientAuthenticationPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationPlugin

`func (o *WorkerConfig) SetClientAuthenticationPlugin(v string)`

SetClientAuthenticationPlugin sets ClientAuthenticationPlugin field to given value.

### HasClientAuthenticationPlugin

`func (o *WorkerConfig) HasClientAuthenticationPlugin() bool`

HasClientAuthenticationPlugin returns a boolean if a field has been set.

### GetFunctionMetadataTopic

`func (o *WorkerConfig) GetFunctionMetadataTopic() string`

GetFunctionMetadataTopic returns the FunctionMetadataTopic field if non-nil, zero value otherwise.

### GetFunctionMetadataTopicOk

`func (o *WorkerConfig) GetFunctionMetadataTopicOk() (*string, bool)`

GetFunctionMetadataTopicOk returns a tuple with the FunctionMetadataTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionMetadataTopic

`func (o *WorkerConfig) SetFunctionMetadataTopic(v string)`

SetFunctionMetadataTopic sets FunctionMetadataTopic field to given value.

### HasFunctionMetadataTopic

`func (o *WorkerConfig) HasFunctionMetadataTopic() bool`

HasFunctionMetadataTopic returns a boolean if a field has been set.

### GetClusterCoordinationTopic

`func (o *WorkerConfig) GetClusterCoordinationTopic() string`

GetClusterCoordinationTopic returns the ClusterCoordinationTopic field if non-nil, zero value otherwise.

### GetClusterCoordinationTopicOk

`func (o *WorkerConfig) GetClusterCoordinationTopicOk() (*string, bool)`

GetClusterCoordinationTopicOk returns a tuple with the ClusterCoordinationTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCoordinationTopic

`func (o *WorkerConfig) SetClusterCoordinationTopic(v string)`

SetClusterCoordinationTopic sets ClusterCoordinationTopic field to given value.

### HasClusterCoordinationTopic

`func (o *WorkerConfig) HasClusterCoordinationTopic() bool`

HasClusterCoordinationTopic returns a boolean if a field has been set.

### GetFunctionAssignmentTopic

`func (o *WorkerConfig) GetFunctionAssignmentTopic() string`

GetFunctionAssignmentTopic returns the FunctionAssignmentTopic field if non-nil, zero value otherwise.

### GetFunctionAssignmentTopicOk

`func (o *WorkerConfig) GetFunctionAssignmentTopicOk() (*string, bool)`

GetFunctionAssignmentTopicOk returns a tuple with the FunctionAssignmentTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAssignmentTopic

`func (o *WorkerConfig) SetFunctionAssignmentTopic(v string)`

SetFunctionAssignmentTopic sets FunctionAssignmentTopic field to given value.

### HasFunctionAssignmentTopic

`func (o *WorkerConfig) HasFunctionAssignmentTopic() bool`

HasFunctionAssignmentTopic returns a boolean if a field has been set.

### GetTlsTrustChainBytes

`func (o *WorkerConfig) GetTlsTrustChainBytes() []string`

GetTlsTrustChainBytes returns the TlsTrustChainBytes field if non-nil, zero value otherwise.

### GetTlsTrustChainBytesOk

`func (o *WorkerConfig) GetTlsTrustChainBytesOk() (*[]string, bool)`

GetTlsTrustChainBytesOk returns a tuple with the TlsTrustChainBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTrustChainBytes

`func (o *WorkerConfig) SetTlsTrustChainBytes(v []string)`

SetTlsTrustChainBytes sets TlsTrustChainBytes field to given value.

### HasTlsTrustChainBytes

`func (o *WorkerConfig) HasTlsTrustChainBytes() bool`

HasTlsTrustChainBytes returns a boolean if a field has been set.

### GetWorkerWebAddress

`func (o *WorkerConfig) GetWorkerWebAddress() string`

GetWorkerWebAddress returns the WorkerWebAddress field if non-nil, zero value otherwise.

### GetWorkerWebAddressOk

`func (o *WorkerConfig) GetWorkerWebAddressOk() (*string, bool)`

GetWorkerWebAddressOk returns a tuple with the WorkerWebAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerWebAddress

`func (o *WorkerConfig) SetWorkerWebAddress(v string)`

SetWorkerWebAddress sets WorkerWebAddress field to given value.

### HasWorkerWebAddress

`func (o *WorkerConfig) HasWorkerWebAddress() bool`

HasWorkerWebAddress returns a boolean if a field has been set.

### GetWorkerWebAddressTls

`func (o *WorkerConfig) GetWorkerWebAddressTls() string`

GetWorkerWebAddressTls returns the WorkerWebAddressTls field if non-nil, zero value otherwise.

### GetWorkerWebAddressTlsOk

`func (o *WorkerConfig) GetWorkerWebAddressTlsOk() (*string, bool)`

GetWorkerWebAddressTlsOk returns a tuple with the WorkerWebAddressTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerWebAddressTls

`func (o *WorkerConfig) SetWorkerWebAddressTls(v string)`

SetWorkerWebAddressTls sets WorkerWebAddressTls field to given value.

### HasWorkerWebAddressTls

`func (o *WorkerConfig) HasWorkerWebAddressTls() bool`

HasWorkerWebAddressTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


