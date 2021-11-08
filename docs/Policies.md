# Policies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPolicies** | Pointer to [**AuthPolicies**](AuthPolicies.md) |  | [optional] 
**ReplicationClusters** | Pointer to **[]string** |  | [optional] 
**Bundles** | Pointer to [**BundlesData**](BundlesData.md) |  | [optional] 
**BacklogQuotaMap** | Pointer to [**map[string]BacklogQuota**](BacklogQuota.md) |  | [optional] 
**ClusterDispatchRate** | Pointer to [**map[string]DispatchRateImpl**](DispatchRateImpl.md) |  | [optional] 
**TopicDispatchRate** | Pointer to [**map[string]DispatchRateImpl**](DispatchRateImpl.md) |  | [optional] 
**SubscriptionDispatchRate** | Pointer to [**map[string]DispatchRateImpl**](DispatchRateImpl.md) |  | [optional] 
**ReplicatorDispatchRate** | Pointer to [**map[string]DispatchRateImpl**](DispatchRateImpl.md) |  | [optional] 
**ClusterSubscribeRate** | Pointer to [**map[string]SubscribeRate**](SubscribeRate.md) |  | [optional] 
**Persistence** | Pointer to [**PersistencePolicies**](PersistencePolicies.md) |  | [optional] 
**DeduplicationEnabled** | Pointer to **bool** |  | [optional] 
**AutoTopicCreationOverride** | Pointer to [**AutoTopicCreationOverride**](AutoTopicCreationOverride.md) |  | [optional] 
**AutoSubscriptionCreationOverride** | Pointer to [**AutoSubscriptionCreationOverride**](AutoSubscriptionCreationOverride.md) |  | [optional] 
**PublishMaxMessageRate** | Pointer to [**map[string]PublishRate**](PublishRate.md) |  | [optional] 
**LatencyStatsSampleRate** | Pointer to **map[string]int32** |  | [optional] 
**MessageTtlInSeconds** | Pointer to **int32** |  | [optional] 
**SubscriptionExpirationTimeMinutes** | Pointer to **int32** |  | [optional] 
**RetentionPolicies** | Pointer to [**RetentionPolicies**](RetentionPolicies.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EncryptionRequired** | Pointer to **bool** |  | [optional] 
**DelayedDeliveryPolicies** | Pointer to [**DelayedDeliveryPolicies**](DelayedDeliveryPolicies.md) |  | [optional] 
**InactiveTopicPolicies** | Pointer to [**InactiveTopicPolicies**](InactiveTopicPolicies.md) |  | [optional] 
**SubscriptionAuthMode** | Pointer to **string** |  | [optional] 
**MaxProducersPerTopic** | Pointer to **int32** |  | [optional] 
**MaxConsumersPerTopic** | Pointer to **int32** |  | [optional] 
**MaxConsumersPerSubscription** | Pointer to **int32** |  | [optional] 
**MaxUnackedMessagesPerConsumer** | Pointer to **int32** |  | [optional] 
**MaxUnackedMessagesPerSubscription** | Pointer to **int32** |  | [optional] 
**MaxSubscriptionsPerTopic** | Pointer to **int32** |  | [optional] 
**CompactionThreshold** | Pointer to **int64** |  | [optional] 
**OffloadThreshold** | Pointer to **int64** |  | [optional] 
**OffloadDeletionLagMs** | Pointer to **int64** |  | [optional] 
**MaxTopicsPerNamespace** | Pointer to **int32** |  | [optional] 
**SchemaAutoUpdateCompatibilityStrategy** | Pointer to **string** |  | [optional] 
**SchemaCompatibilityStrategy** | Pointer to **string** |  | [optional] 
**IsAllowAutoUpdateSchema** | Pointer to **bool** |  | [optional] 
**SchemaValidationEnforced** | Pointer to **bool** |  | [optional] 
**OffloadPolicies** | Pointer to [**OffloadPolicies**](OffloadPolicies.md) |  | [optional] 
**DeduplicationSnapshotIntervalSeconds** | Pointer to **int32** |  | [optional] 
**SubscriptionTypesEnabled** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**ResourceGroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicies

`func NewPolicies() *Policies`

NewPolicies instantiates a new Policies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesWithDefaults

`func NewPoliciesWithDefaults() *Policies`

NewPoliciesWithDefaults instantiates a new Policies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPolicies

`func (o *Policies) GetAuthPolicies() AuthPolicies`

GetAuthPolicies returns the AuthPolicies field if non-nil, zero value otherwise.

### GetAuthPoliciesOk

`func (o *Policies) GetAuthPoliciesOk() (*AuthPolicies, bool)`

GetAuthPoliciesOk returns a tuple with the AuthPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPolicies

`func (o *Policies) SetAuthPolicies(v AuthPolicies)`

SetAuthPolicies sets AuthPolicies field to given value.

### HasAuthPolicies

`func (o *Policies) HasAuthPolicies() bool`

HasAuthPolicies returns a boolean if a field has been set.

### GetReplicationClusters

`func (o *Policies) GetReplicationClusters() []string`

GetReplicationClusters returns the ReplicationClusters field if non-nil, zero value otherwise.

### GetReplicationClustersOk

`func (o *Policies) GetReplicationClustersOk() (*[]string, bool)`

GetReplicationClustersOk returns a tuple with the ReplicationClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationClusters

`func (o *Policies) SetReplicationClusters(v []string)`

SetReplicationClusters sets ReplicationClusters field to given value.

### HasReplicationClusters

`func (o *Policies) HasReplicationClusters() bool`

HasReplicationClusters returns a boolean if a field has been set.

### GetBundles

`func (o *Policies) GetBundles() BundlesData`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *Policies) GetBundlesOk() (*BundlesData, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *Policies) SetBundles(v BundlesData)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *Policies) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetBacklogQuotaMap

`func (o *Policies) GetBacklogQuotaMap() map[string]BacklogQuota`

GetBacklogQuotaMap returns the BacklogQuotaMap field if non-nil, zero value otherwise.

### GetBacklogQuotaMapOk

`func (o *Policies) GetBacklogQuotaMapOk() (*map[string]BacklogQuota, bool)`

GetBacklogQuotaMapOk returns a tuple with the BacklogQuotaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogQuotaMap

`func (o *Policies) SetBacklogQuotaMap(v map[string]BacklogQuota)`

SetBacklogQuotaMap sets BacklogQuotaMap field to given value.

### HasBacklogQuotaMap

`func (o *Policies) HasBacklogQuotaMap() bool`

HasBacklogQuotaMap returns a boolean if a field has been set.

### GetClusterDispatchRate

`func (o *Policies) GetClusterDispatchRate() map[string]DispatchRateImpl`

GetClusterDispatchRate returns the ClusterDispatchRate field if non-nil, zero value otherwise.

### GetClusterDispatchRateOk

`func (o *Policies) GetClusterDispatchRateOk() (*map[string]DispatchRateImpl, bool)`

GetClusterDispatchRateOk returns a tuple with the ClusterDispatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDispatchRate

`func (o *Policies) SetClusterDispatchRate(v map[string]DispatchRateImpl)`

SetClusterDispatchRate sets ClusterDispatchRate field to given value.

### HasClusterDispatchRate

`func (o *Policies) HasClusterDispatchRate() bool`

HasClusterDispatchRate returns a boolean if a field has been set.

### GetTopicDispatchRate

`func (o *Policies) GetTopicDispatchRate() map[string]DispatchRateImpl`

GetTopicDispatchRate returns the TopicDispatchRate field if non-nil, zero value otherwise.

### GetTopicDispatchRateOk

`func (o *Policies) GetTopicDispatchRateOk() (*map[string]DispatchRateImpl, bool)`

GetTopicDispatchRateOk returns a tuple with the TopicDispatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicDispatchRate

`func (o *Policies) SetTopicDispatchRate(v map[string]DispatchRateImpl)`

SetTopicDispatchRate sets TopicDispatchRate field to given value.

### HasTopicDispatchRate

`func (o *Policies) HasTopicDispatchRate() bool`

HasTopicDispatchRate returns a boolean if a field has been set.

### GetSubscriptionDispatchRate

`func (o *Policies) GetSubscriptionDispatchRate() map[string]DispatchRateImpl`

GetSubscriptionDispatchRate returns the SubscriptionDispatchRate field if non-nil, zero value otherwise.

### GetSubscriptionDispatchRateOk

`func (o *Policies) GetSubscriptionDispatchRateOk() (*map[string]DispatchRateImpl, bool)`

GetSubscriptionDispatchRateOk returns a tuple with the SubscriptionDispatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDispatchRate

`func (o *Policies) SetSubscriptionDispatchRate(v map[string]DispatchRateImpl)`

SetSubscriptionDispatchRate sets SubscriptionDispatchRate field to given value.

### HasSubscriptionDispatchRate

`func (o *Policies) HasSubscriptionDispatchRate() bool`

HasSubscriptionDispatchRate returns a boolean if a field has been set.

### GetReplicatorDispatchRate

`func (o *Policies) GetReplicatorDispatchRate() map[string]DispatchRateImpl`

GetReplicatorDispatchRate returns the ReplicatorDispatchRate field if non-nil, zero value otherwise.

### GetReplicatorDispatchRateOk

`func (o *Policies) GetReplicatorDispatchRateOk() (*map[string]DispatchRateImpl, bool)`

GetReplicatorDispatchRateOk returns a tuple with the ReplicatorDispatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatorDispatchRate

`func (o *Policies) SetReplicatorDispatchRate(v map[string]DispatchRateImpl)`

SetReplicatorDispatchRate sets ReplicatorDispatchRate field to given value.

### HasReplicatorDispatchRate

`func (o *Policies) HasReplicatorDispatchRate() bool`

HasReplicatorDispatchRate returns a boolean if a field has been set.

### GetClusterSubscribeRate

`func (o *Policies) GetClusterSubscribeRate() map[string]SubscribeRate`

GetClusterSubscribeRate returns the ClusterSubscribeRate field if non-nil, zero value otherwise.

### GetClusterSubscribeRateOk

`func (o *Policies) GetClusterSubscribeRateOk() (*map[string]SubscribeRate, bool)`

GetClusterSubscribeRateOk returns a tuple with the ClusterSubscribeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSubscribeRate

`func (o *Policies) SetClusterSubscribeRate(v map[string]SubscribeRate)`

SetClusterSubscribeRate sets ClusterSubscribeRate field to given value.

### HasClusterSubscribeRate

`func (o *Policies) HasClusterSubscribeRate() bool`

HasClusterSubscribeRate returns a boolean if a field has been set.

### GetPersistence

`func (o *Policies) GetPersistence() PersistencePolicies`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *Policies) GetPersistenceOk() (*PersistencePolicies, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *Policies) SetPersistence(v PersistencePolicies)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *Policies) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### GetDeduplicationEnabled

`func (o *Policies) GetDeduplicationEnabled() bool`

GetDeduplicationEnabled returns the DeduplicationEnabled field if non-nil, zero value otherwise.

### GetDeduplicationEnabledOk

`func (o *Policies) GetDeduplicationEnabledOk() (*bool, bool)`

GetDeduplicationEnabledOk returns a tuple with the DeduplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationEnabled

`func (o *Policies) SetDeduplicationEnabled(v bool)`

SetDeduplicationEnabled sets DeduplicationEnabled field to given value.

### HasDeduplicationEnabled

`func (o *Policies) HasDeduplicationEnabled() bool`

HasDeduplicationEnabled returns a boolean if a field has been set.

### GetAutoTopicCreationOverride

`func (o *Policies) GetAutoTopicCreationOverride() AutoTopicCreationOverride`

GetAutoTopicCreationOverride returns the AutoTopicCreationOverride field if non-nil, zero value otherwise.

### GetAutoTopicCreationOverrideOk

`func (o *Policies) GetAutoTopicCreationOverrideOk() (*AutoTopicCreationOverride, bool)`

GetAutoTopicCreationOverrideOk returns a tuple with the AutoTopicCreationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopicCreationOverride

`func (o *Policies) SetAutoTopicCreationOverride(v AutoTopicCreationOverride)`

SetAutoTopicCreationOverride sets AutoTopicCreationOverride field to given value.

### HasAutoTopicCreationOverride

`func (o *Policies) HasAutoTopicCreationOverride() bool`

HasAutoTopicCreationOverride returns a boolean if a field has been set.

### GetAutoSubscriptionCreationOverride

`func (o *Policies) GetAutoSubscriptionCreationOverride() AutoSubscriptionCreationOverride`

GetAutoSubscriptionCreationOverride returns the AutoSubscriptionCreationOverride field if non-nil, zero value otherwise.

### GetAutoSubscriptionCreationOverrideOk

`func (o *Policies) GetAutoSubscriptionCreationOverrideOk() (*AutoSubscriptionCreationOverride, bool)`

GetAutoSubscriptionCreationOverrideOk returns a tuple with the AutoSubscriptionCreationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscriptionCreationOverride

`func (o *Policies) SetAutoSubscriptionCreationOverride(v AutoSubscriptionCreationOverride)`

SetAutoSubscriptionCreationOverride sets AutoSubscriptionCreationOverride field to given value.

### HasAutoSubscriptionCreationOverride

`func (o *Policies) HasAutoSubscriptionCreationOverride() bool`

HasAutoSubscriptionCreationOverride returns a boolean if a field has been set.

### GetPublishMaxMessageRate

`func (o *Policies) GetPublishMaxMessageRate() map[string]PublishRate`

GetPublishMaxMessageRate returns the PublishMaxMessageRate field if non-nil, zero value otherwise.

### GetPublishMaxMessageRateOk

`func (o *Policies) GetPublishMaxMessageRateOk() (*map[string]PublishRate, bool)`

GetPublishMaxMessageRateOk returns a tuple with the PublishMaxMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishMaxMessageRate

`func (o *Policies) SetPublishMaxMessageRate(v map[string]PublishRate)`

SetPublishMaxMessageRate sets PublishMaxMessageRate field to given value.

### HasPublishMaxMessageRate

`func (o *Policies) HasPublishMaxMessageRate() bool`

HasPublishMaxMessageRate returns a boolean if a field has been set.

### GetLatencyStatsSampleRate

`func (o *Policies) GetLatencyStatsSampleRate() map[string]int32`

GetLatencyStatsSampleRate returns the LatencyStatsSampleRate field if non-nil, zero value otherwise.

### GetLatencyStatsSampleRateOk

`func (o *Policies) GetLatencyStatsSampleRateOk() (*map[string]int32, bool)`

GetLatencyStatsSampleRateOk returns a tuple with the LatencyStatsSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyStatsSampleRate

`func (o *Policies) SetLatencyStatsSampleRate(v map[string]int32)`

SetLatencyStatsSampleRate sets LatencyStatsSampleRate field to given value.

### HasLatencyStatsSampleRate

`func (o *Policies) HasLatencyStatsSampleRate() bool`

HasLatencyStatsSampleRate returns a boolean if a field has been set.

### GetMessageTtlInSeconds

`func (o *Policies) GetMessageTtlInSeconds() int32`

GetMessageTtlInSeconds returns the MessageTtlInSeconds field if non-nil, zero value otherwise.

### GetMessageTtlInSecondsOk

`func (o *Policies) GetMessageTtlInSecondsOk() (*int32, bool)`

GetMessageTtlInSecondsOk returns a tuple with the MessageTtlInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtlInSeconds

`func (o *Policies) SetMessageTtlInSeconds(v int32)`

SetMessageTtlInSeconds sets MessageTtlInSeconds field to given value.

### HasMessageTtlInSeconds

`func (o *Policies) HasMessageTtlInSeconds() bool`

HasMessageTtlInSeconds returns a boolean if a field has been set.

### GetSubscriptionExpirationTimeMinutes

`func (o *Policies) GetSubscriptionExpirationTimeMinutes() int32`

GetSubscriptionExpirationTimeMinutes returns the SubscriptionExpirationTimeMinutes field if non-nil, zero value otherwise.

### GetSubscriptionExpirationTimeMinutesOk

`func (o *Policies) GetSubscriptionExpirationTimeMinutesOk() (*int32, bool)`

GetSubscriptionExpirationTimeMinutesOk returns a tuple with the SubscriptionExpirationTimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpirationTimeMinutes

`func (o *Policies) SetSubscriptionExpirationTimeMinutes(v int32)`

SetSubscriptionExpirationTimeMinutes sets SubscriptionExpirationTimeMinutes field to given value.

### HasSubscriptionExpirationTimeMinutes

`func (o *Policies) HasSubscriptionExpirationTimeMinutes() bool`

HasSubscriptionExpirationTimeMinutes returns a boolean if a field has been set.

### GetRetentionPolicies

`func (o *Policies) GetRetentionPolicies() RetentionPolicies`

GetRetentionPolicies returns the RetentionPolicies field if non-nil, zero value otherwise.

### GetRetentionPoliciesOk

`func (o *Policies) GetRetentionPoliciesOk() (*RetentionPolicies, bool)`

GetRetentionPoliciesOk returns a tuple with the RetentionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicies

`func (o *Policies) SetRetentionPolicies(v RetentionPolicies)`

SetRetentionPolicies sets RetentionPolicies field to given value.

### HasRetentionPolicies

`func (o *Policies) HasRetentionPolicies() bool`

HasRetentionPolicies returns a boolean if a field has been set.

### GetDeleted

`func (o *Policies) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Policies) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Policies) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Policies) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEncryptionRequired

`func (o *Policies) GetEncryptionRequired() bool`

GetEncryptionRequired returns the EncryptionRequired field if non-nil, zero value otherwise.

### GetEncryptionRequiredOk

`func (o *Policies) GetEncryptionRequiredOk() (*bool, bool)`

GetEncryptionRequiredOk returns a tuple with the EncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionRequired

`func (o *Policies) SetEncryptionRequired(v bool)`

SetEncryptionRequired sets EncryptionRequired field to given value.

### HasEncryptionRequired

`func (o *Policies) HasEncryptionRequired() bool`

HasEncryptionRequired returns a boolean if a field has been set.

### GetDelayedDeliveryPolicies

`func (o *Policies) GetDelayedDeliveryPolicies() DelayedDeliveryPolicies`

GetDelayedDeliveryPolicies returns the DelayedDeliveryPolicies field if non-nil, zero value otherwise.

### GetDelayedDeliveryPoliciesOk

`func (o *Policies) GetDelayedDeliveryPoliciesOk() (*DelayedDeliveryPolicies, bool)`

GetDelayedDeliveryPoliciesOk returns a tuple with the DelayedDeliveryPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedDeliveryPolicies

`func (o *Policies) SetDelayedDeliveryPolicies(v DelayedDeliveryPolicies)`

SetDelayedDeliveryPolicies sets DelayedDeliveryPolicies field to given value.

### HasDelayedDeliveryPolicies

`func (o *Policies) HasDelayedDeliveryPolicies() bool`

HasDelayedDeliveryPolicies returns a boolean if a field has been set.

### GetInactiveTopicPolicies

`func (o *Policies) GetInactiveTopicPolicies() InactiveTopicPolicies`

GetInactiveTopicPolicies returns the InactiveTopicPolicies field if non-nil, zero value otherwise.

### GetInactiveTopicPoliciesOk

`func (o *Policies) GetInactiveTopicPoliciesOk() (*InactiveTopicPolicies, bool)`

GetInactiveTopicPoliciesOk returns a tuple with the InactiveTopicPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveTopicPolicies

`func (o *Policies) SetInactiveTopicPolicies(v InactiveTopicPolicies)`

SetInactiveTopicPolicies sets InactiveTopicPolicies field to given value.

### HasInactiveTopicPolicies

`func (o *Policies) HasInactiveTopicPolicies() bool`

HasInactiveTopicPolicies returns a boolean if a field has been set.

### GetSubscriptionAuthMode

`func (o *Policies) GetSubscriptionAuthMode() string`

GetSubscriptionAuthMode returns the SubscriptionAuthMode field if non-nil, zero value otherwise.

### GetSubscriptionAuthModeOk

`func (o *Policies) GetSubscriptionAuthModeOk() (*string, bool)`

GetSubscriptionAuthModeOk returns a tuple with the SubscriptionAuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAuthMode

`func (o *Policies) SetSubscriptionAuthMode(v string)`

SetSubscriptionAuthMode sets SubscriptionAuthMode field to given value.

### HasSubscriptionAuthMode

`func (o *Policies) HasSubscriptionAuthMode() bool`

HasSubscriptionAuthMode returns a boolean if a field has been set.

### GetMaxProducersPerTopic

`func (o *Policies) GetMaxProducersPerTopic() int32`

GetMaxProducersPerTopic returns the MaxProducersPerTopic field if non-nil, zero value otherwise.

### GetMaxProducersPerTopicOk

`func (o *Policies) GetMaxProducersPerTopicOk() (*int32, bool)`

GetMaxProducersPerTopicOk returns a tuple with the MaxProducersPerTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProducersPerTopic

`func (o *Policies) SetMaxProducersPerTopic(v int32)`

SetMaxProducersPerTopic sets MaxProducersPerTopic field to given value.

### HasMaxProducersPerTopic

`func (o *Policies) HasMaxProducersPerTopic() bool`

HasMaxProducersPerTopic returns a boolean if a field has been set.

### GetMaxConsumersPerTopic

`func (o *Policies) GetMaxConsumersPerTopic() int32`

GetMaxConsumersPerTopic returns the MaxConsumersPerTopic field if non-nil, zero value otherwise.

### GetMaxConsumersPerTopicOk

`func (o *Policies) GetMaxConsumersPerTopicOk() (*int32, bool)`

GetMaxConsumersPerTopicOk returns a tuple with the MaxConsumersPerTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsumersPerTopic

`func (o *Policies) SetMaxConsumersPerTopic(v int32)`

SetMaxConsumersPerTopic sets MaxConsumersPerTopic field to given value.

### HasMaxConsumersPerTopic

`func (o *Policies) HasMaxConsumersPerTopic() bool`

HasMaxConsumersPerTopic returns a boolean if a field has been set.

### GetMaxConsumersPerSubscription

`func (o *Policies) GetMaxConsumersPerSubscription() int32`

GetMaxConsumersPerSubscription returns the MaxConsumersPerSubscription field if non-nil, zero value otherwise.

### GetMaxConsumersPerSubscriptionOk

`func (o *Policies) GetMaxConsumersPerSubscriptionOk() (*int32, bool)`

GetMaxConsumersPerSubscriptionOk returns a tuple with the MaxConsumersPerSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsumersPerSubscription

`func (o *Policies) SetMaxConsumersPerSubscription(v int32)`

SetMaxConsumersPerSubscription sets MaxConsumersPerSubscription field to given value.

### HasMaxConsumersPerSubscription

`func (o *Policies) HasMaxConsumersPerSubscription() bool`

HasMaxConsumersPerSubscription returns a boolean if a field has been set.

### GetMaxUnackedMessagesPerConsumer

`func (o *Policies) GetMaxUnackedMessagesPerConsumer() int32`

GetMaxUnackedMessagesPerConsumer returns the MaxUnackedMessagesPerConsumer field if non-nil, zero value otherwise.

### GetMaxUnackedMessagesPerConsumerOk

`func (o *Policies) GetMaxUnackedMessagesPerConsumerOk() (*int32, bool)`

GetMaxUnackedMessagesPerConsumerOk returns a tuple with the MaxUnackedMessagesPerConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedMessagesPerConsumer

`func (o *Policies) SetMaxUnackedMessagesPerConsumer(v int32)`

SetMaxUnackedMessagesPerConsumer sets MaxUnackedMessagesPerConsumer field to given value.

### HasMaxUnackedMessagesPerConsumer

`func (o *Policies) HasMaxUnackedMessagesPerConsumer() bool`

HasMaxUnackedMessagesPerConsumer returns a boolean if a field has been set.

### GetMaxUnackedMessagesPerSubscription

`func (o *Policies) GetMaxUnackedMessagesPerSubscription() int32`

GetMaxUnackedMessagesPerSubscription returns the MaxUnackedMessagesPerSubscription field if non-nil, zero value otherwise.

### GetMaxUnackedMessagesPerSubscriptionOk

`func (o *Policies) GetMaxUnackedMessagesPerSubscriptionOk() (*int32, bool)`

GetMaxUnackedMessagesPerSubscriptionOk returns a tuple with the MaxUnackedMessagesPerSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedMessagesPerSubscription

`func (o *Policies) SetMaxUnackedMessagesPerSubscription(v int32)`

SetMaxUnackedMessagesPerSubscription sets MaxUnackedMessagesPerSubscription field to given value.

### HasMaxUnackedMessagesPerSubscription

`func (o *Policies) HasMaxUnackedMessagesPerSubscription() bool`

HasMaxUnackedMessagesPerSubscription returns a boolean if a field has been set.

### GetMaxSubscriptionsPerTopic

`func (o *Policies) GetMaxSubscriptionsPerTopic() int32`

GetMaxSubscriptionsPerTopic returns the MaxSubscriptionsPerTopic field if non-nil, zero value otherwise.

### GetMaxSubscriptionsPerTopicOk

`func (o *Policies) GetMaxSubscriptionsPerTopicOk() (*int32, bool)`

GetMaxSubscriptionsPerTopicOk returns a tuple with the MaxSubscriptionsPerTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptionsPerTopic

`func (o *Policies) SetMaxSubscriptionsPerTopic(v int32)`

SetMaxSubscriptionsPerTopic sets MaxSubscriptionsPerTopic field to given value.

### HasMaxSubscriptionsPerTopic

`func (o *Policies) HasMaxSubscriptionsPerTopic() bool`

HasMaxSubscriptionsPerTopic returns a boolean if a field has been set.

### GetCompactionThreshold

`func (o *Policies) GetCompactionThreshold() int64`

GetCompactionThreshold returns the CompactionThreshold field if non-nil, zero value otherwise.

### GetCompactionThresholdOk

`func (o *Policies) GetCompactionThresholdOk() (*int64, bool)`

GetCompactionThresholdOk returns a tuple with the CompactionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactionThreshold

`func (o *Policies) SetCompactionThreshold(v int64)`

SetCompactionThreshold sets CompactionThreshold field to given value.

### HasCompactionThreshold

`func (o *Policies) HasCompactionThreshold() bool`

HasCompactionThreshold returns a boolean if a field has been set.

### GetOffloadThreshold

`func (o *Policies) GetOffloadThreshold() int64`

GetOffloadThreshold returns the OffloadThreshold field if non-nil, zero value otherwise.

### GetOffloadThresholdOk

`func (o *Policies) GetOffloadThresholdOk() (*int64, bool)`

GetOffloadThresholdOk returns a tuple with the OffloadThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadThreshold

`func (o *Policies) SetOffloadThreshold(v int64)`

SetOffloadThreshold sets OffloadThreshold field to given value.

### HasOffloadThreshold

`func (o *Policies) HasOffloadThreshold() bool`

HasOffloadThreshold returns a boolean if a field has been set.

### GetOffloadDeletionLagMs

`func (o *Policies) GetOffloadDeletionLagMs() int64`

GetOffloadDeletionLagMs returns the OffloadDeletionLagMs field if non-nil, zero value otherwise.

### GetOffloadDeletionLagMsOk

`func (o *Policies) GetOffloadDeletionLagMsOk() (*int64, bool)`

GetOffloadDeletionLagMsOk returns a tuple with the OffloadDeletionLagMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadDeletionLagMs

`func (o *Policies) SetOffloadDeletionLagMs(v int64)`

SetOffloadDeletionLagMs sets OffloadDeletionLagMs field to given value.

### HasOffloadDeletionLagMs

`func (o *Policies) HasOffloadDeletionLagMs() bool`

HasOffloadDeletionLagMs returns a boolean if a field has been set.

### GetMaxTopicsPerNamespace

`func (o *Policies) GetMaxTopicsPerNamespace() int32`

GetMaxTopicsPerNamespace returns the MaxTopicsPerNamespace field if non-nil, zero value otherwise.

### GetMaxTopicsPerNamespaceOk

`func (o *Policies) GetMaxTopicsPerNamespaceOk() (*int32, bool)`

GetMaxTopicsPerNamespaceOk returns a tuple with the MaxTopicsPerNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopicsPerNamespace

`func (o *Policies) SetMaxTopicsPerNamespace(v int32)`

SetMaxTopicsPerNamespace sets MaxTopicsPerNamespace field to given value.

### HasMaxTopicsPerNamespace

`func (o *Policies) HasMaxTopicsPerNamespace() bool`

HasMaxTopicsPerNamespace returns a boolean if a field has been set.

### GetSchemaAutoUpdateCompatibilityStrategy

`func (o *Policies) GetSchemaAutoUpdateCompatibilityStrategy() string`

GetSchemaAutoUpdateCompatibilityStrategy returns the SchemaAutoUpdateCompatibilityStrategy field if non-nil, zero value otherwise.

### GetSchemaAutoUpdateCompatibilityStrategyOk

`func (o *Policies) GetSchemaAutoUpdateCompatibilityStrategyOk() (*string, bool)`

GetSchemaAutoUpdateCompatibilityStrategyOk returns a tuple with the SchemaAutoUpdateCompatibilityStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaAutoUpdateCompatibilityStrategy

`func (o *Policies) SetSchemaAutoUpdateCompatibilityStrategy(v string)`

SetSchemaAutoUpdateCompatibilityStrategy sets SchemaAutoUpdateCompatibilityStrategy field to given value.

### HasSchemaAutoUpdateCompatibilityStrategy

`func (o *Policies) HasSchemaAutoUpdateCompatibilityStrategy() bool`

HasSchemaAutoUpdateCompatibilityStrategy returns a boolean if a field has been set.

### GetSchemaCompatibilityStrategy

`func (o *Policies) GetSchemaCompatibilityStrategy() string`

GetSchemaCompatibilityStrategy returns the SchemaCompatibilityStrategy field if non-nil, zero value otherwise.

### GetSchemaCompatibilityStrategyOk

`func (o *Policies) GetSchemaCompatibilityStrategyOk() (*string, bool)`

GetSchemaCompatibilityStrategyOk returns a tuple with the SchemaCompatibilityStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaCompatibilityStrategy

`func (o *Policies) SetSchemaCompatibilityStrategy(v string)`

SetSchemaCompatibilityStrategy sets SchemaCompatibilityStrategy field to given value.

### HasSchemaCompatibilityStrategy

`func (o *Policies) HasSchemaCompatibilityStrategy() bool`

HasSchemaCompatibilityStrategy returns a boolean if a field has been set.

### GetIsAllowAutoUpdateSchema

`func (o *Policies) GetIsAllowAutoUpdateSchema() bool`

GetIsAllowAutoUpdateSchema returns the IsAllowAutoUpdateSchema field if non-nil, zero value otherwise.

### GetIsAllowAutoUpdateSchemaOk

`func (o *Policies) GetIsAllowAutoUpdateSchemaOk() (*bool, bool)`

GetIsAllowAutoUpdateSchemaOk returns a tuple with the IsAllowAutoUpdateSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowAutoUpdateSchema

`func (o *Policies) SetIsAllowAutoUpdateSchema(v bool)`

SetIsAllowAutoUpdateSchema sets IsAllowAutoUpdateSchema field to given value.

### HasIsAllowAutoUpdateSchema

`func (o *Policies) HasIsAllowAutoUpdateSchema() bool`

HasIsAllowAutoUpdateSchema returns a boolean if a field has been set.

### GetSchemaValidationEnforced

`func (o *Policies) GetSchemaValidationEnforced() bool`

GetSchemaValidationEnforced returns the SchemaValidationEnforced field if non-nil, zero value otherwise.

### GetSchemaValidationEnforcedOk

`func (o *Policies) GetSchemaValidationEnforcedOk() (*bool, bool)`

GetSchemaValidationEnforcedOk returns a tuple with the SchemaValidationEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaValidationEnforced

`func (o *Policies) SetSchemaValidationEnforced(v bool)`

SetSchemaValidationEnforced sets SchemaValidationEnforced field to given value.

### HasSchemaValidationEnforced

`func (o *Policies) HasSchemaValidationEnforced() bool`

HasSchemaValidationEnforced returns a boolean if a field has been set.

### GetOffloadPolicies

`func (o *Policies) GetOffloadPolicies() OffloadPolicies`

GetOffloadPolicies returns the OffloadPolicies field if non-nil, zero value otherwise.

### GetOffloadPoliciesOk

`func (o *Policies) GetOffloadPoliciesOk() (*OffloadPolicies, bool)`

GetOffloadPoliciesOk returns a tuple with the OffloadPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadPolicies

`func (o *Policies) SetOffloadPolicies(v OffloadPolicies)`

SetOffloadPolicies sets OffloadPolicies field to given value.

### HasOffloadPolicies

`func (o *Policies) HasOffloadPolicies() bool`

HasOffloadPolicies returns a boolean if a field has been set.

### GetDeduplicationSnapshotIntervalSeconds

`func (o *Policies) GetDeduplicationSnapshotIntervalSeconds() int32`

GetDeduplicationSnapshotIntervalSeconds returns the DeduplicationSnapshotIntervalSeconds field if non-nil, zero value otherwise.

### GetDeduplicationSnapshotIntervalSecondsOk

`func (o *Policies) GetDeduplicationSnapshotIntervalSecondsOk() (*int32, bool)`

GetDeduplicationSnapshotIntervalSecondsOk returns a tuple with the DeduplicationSnapshotIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationSnapshotIntervalSeconds

`func (o *Policies) SetDeduplicationSnapshotIntervalSeconds(v int32)`

SetDeduplicationSnapshotIntervalSeconds sets DeduplicationSnapshotIntervalSeconds field to given value.

### HasDeduplicationSnapshotIntervalSeconds

`func (o *Policies) HasDeduplicationSnapshotIntervalSeconds() bool`

HasDeduplicationSnapshotIntervalSeconds returns a boolean if a field has been set.

### GetSubscriptionTypesEnabled

`func (o *Policies) GetSubscriptionTypesEnabled() []string`

GetSubscriptionTypesEnabled returns the SubscriptionTypesEnabled field if non-nil, zero value otherwise.

### GetSubscriptionTypesEnabledOk

`func (o *Policies) GetSubscriptionTypesEnabledOk() (*[]string, bool)`

GetSubscriptionTypesEnabledOk returns a tuple with the SubscriptionTypesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypesEnabled

`func (o *Policies) SetSubscriptionTypesEnabled(v []string)`

SetSubscriptionTypesEnabled sets SubscriptionTypesEnabled field to given value.

### HasSubscriptionTypesEnabled

`func (o *Policies) HasSubscriptionTypesEnabled() bool`

HasSubscriptionTypesEnabled returns a boolean if a field has been set.

### GetProperties

`func (o *Policies) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Policies) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Policies) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Policies) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *Policies) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *Policies) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *Policies) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *Policies) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


