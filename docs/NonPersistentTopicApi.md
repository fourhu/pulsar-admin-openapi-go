# \NonPersistentTopicApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Compact**](NonPersistentTopicApi.md#Compact) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Trigger a compaction operation on a topic.
[**CompactionStatus**](NonPersistentTopicApi.md#CompactionStatus) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Get the status of a compaction operation for a topic.
[**CreateMissedPartitions**](NonPersistentTopicApi.md#CreateMissedPartitions) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/createMissedPartitions | Create missed partitions of an existing partitioned topic.
[**CreateNonPartitionedTopic**](NonPersistentTopicApi.md#CreateNonPartitionedTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic} | Create a non-partitioned topic.
[**CreatePartitionedTopic**](NonPersistentTopicApi.md#CreatePartitionedTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Create a partitioned topic.
[**CreateSubscription**](NonPersistentTopicApi.md#CreateSubscription) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subscriptionName} | Create a subscription on the topic.
[**DeleteDeduplicationSnapshotInterval**](NonPersistentTopicApi.md#DeleteDeduplicationSnapshotInterval) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Delete deduplicationSnapshotInterval config on a topic.
[**DeleteDelayedDeliveryPolicies**](NonPersistentTopicApi.md#DeleteDelayedDeliveryPolicies) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**DeleteInactiveTopicPolicies**](NonPersistentTopicApi.md#DeleteInactiveTopicPolicies) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Delete inactive topic policies on a topic.
[**DeleteMaxUnackedMessagesOnConsumer**](NonPersistentTopicApi.md#DeleteMaxUnackedMessagesOnConsumer) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Delete max unacked messages per consumer config on a topic.
[**DeleteMaxUnackedMessagesOnSubscription**](NonPersistentTopicApi.md#DeleteMaxUnackedMessagesOnSubscription) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Delete max unacked messages per subscription config on a topic.
[**DeletePartitionedTopic**](NonPersistentTopicApi.md#DeletePartitionedTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Delete a partitioned topic.
[**DeleteSubscription**](NonPersistentTopicApi.md#DeleteSubscription) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName} | Delete a subscription.
[**DeleteTopic**](NonPersistentTopicApi.md#DeleteTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic} | Delete a topic.
[**ExamineMessage**](NonPersistentTopicApi.md#ExamineMessage) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/examinemessage | Examine a specific message on a topic by position relative to the earliest or the latest message.
[**ExpireMessagesForAllSubscriptions**](NonPersistentTopicApi.md#ExpireMessagesForAllSubscriptions) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/all_subscription/expireMessages/{expireTimeInSeconds} | Expiry messages on all subscriptions of topic.
[**ExpireTopicMessages**](NonPersistentTopicApi.md#ExpireTopicMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages | Expiry messages on a topic subscription.
[**ExpireTopicMessages_0**](NonPersistentTopicApi.md#ExpireTopicMessages_0) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages/{expireTimeInSeconds} | Expiry messages on a topic subscription.
[**GetBacklog**](NonPersistentTopicApi.md#GetBacklog) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/backlog | Get estimated backlog for offline topic.
[**GetBacklogQuotaMap**](NonPersistentTopicApi.md#GetBacklogQuotaMap) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuotaMap | Get backlog quota map on a topic.
[**GetCompactionThreshold**](NonPersistentTopicApi.md#GetCompactionThreshold) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Get compaction threshold configuration for specified topic.
[**GetDeduplication**](NonPersistentTopicApi.md#GetDeduplication) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Get deduplication configuration of a topic.
[**GetDeduplicationSnapshotInterval**](NonPersistentTopicApi.md#GetDeduplicationSnapshotInterval) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Get deduplicationSnapshotInterval config on a topic.
[**GetDelayedDeliveryPolicies**](NonPersistentTopicApi.md#GetDelayedDeliveryPolicies) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Get delayed delivery messages config on a topic.
[**GetDispatchRate**](NonPersistentTopicApi.md#GetDispatchRate) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Get dispatch rate configuration for specified topic.
[**GetInactiveTopicPolicies**](NonPersistentTopicApi.md#GetInactiveTopicPolicies) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Get inactive topic policies on a topic.
[**GetInternalStats**](NonPersistentTopicApi.md#GetInternalStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internalStats | Get the internal stats for the topic.
[**GetLastMessageId**](NonPersistentTopicApi.md#GetLastMessageId) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/lastMessageId | Return the last commit message id of topic
[**GetList**](NonPersistentTopicApi.md#GetList) | **Get** /non-persistent/{tenant}/{namespace} | Get the list of non-persistent topics under a namespace.
[**GetListFromBundle**](NonPersistentTopicApi.md#GetListFromBundle) | **Get** /non-persistent/{tenant}/{namespace}/{bundle} | Get the list of non-persistent topics under a namespace bundle.
[**GetManagedLedgerInfo**](NonPersistentTopicApi.md#GetManagedLedgerInfo) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internal-info | Get the stored topic metadata.
[**GetMaxConsumers**](NonPersistentTopicApi.md#GetMaxConsumers) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Get maxConsumers config for specified topic.
[**GetMaxConsumersPerSubscription**](NonPersistentTopicApi.md#GetMaxConsumersPerSubscription) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Get max consumers per subscription configuration for specified topic.
[**GetMaxMessageSize**](NonPersistentTopicApi.md#GetMaxMessageSize) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Get maxMessageSize config for specified topic.
[**GetMaxProducers**](NonPersistentTopicApi.md#GetMaxProducers) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Get maxProducers config for specified topic.
[**GetMaxSubscriptionsPerTopic**](NonPersistentTopicApi.md#GetMaxSubscriptionsPerTopic) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Get maxSubscriptionsPerTopic config for specified topic.
[**GetMaxUnackedMessagesOnConsumer**](NonPersistentTopicApi.md#GetMaxUnackedMessagesOnConsumer) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Get max unacked messages per consumer config on a topic.
[**GetMaxUnackedMessagesOnSubscription**](NonPersistentTopicApi.md#GetMaxUnackedMessagesOnSubscription) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Get max unacked messages per subscription config on a topic.
[**GetMessageById**](NonPersistentTopicApi.md#GetMessageById) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/ledger/{ledgerId}/entry/{entryId} | Get message by its messageId.
[**GetMessageTTL**](NonPersistentTopicApi.md#GetMessageTTL) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Get message TTL in seconds for a topic
[**GetOffloadPolicies**](NonPersistentTopicApi.md#GetOffloadPolicies) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Get offload policies on a topic.
[**GetPartitionedMetadata**](NonPersistentTopicApi.md#GetPartitionedMetadata) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Get partitioned topic metadata.
[**GetPartitionedStats**](NonPersistentTopicApi.md#GetPartitionedStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitioned-stats | Get the stats for the partitioned topic.
[**GetPartitionedTopicList**](NonPersistentTopicApi.md#GetPartitionedTopicList) | **Get** /non-persistent/{tenant}/{namespace}/partitioned | Get the list of partitioned topics under a namespace.
[**GetPermissionsOnTopic**](NonPersistentTopicApi.md#GetPermissionsOnTopic) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/permissions | Get permissions on a topic.
[**GetPersistence**](NonPersistentTopicApi.md#GetPersistence) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Get configuration of persistence policies for specified topic.
[**GetPublishRate**](NonPersistentTopicApi.md#GetPublishRate) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Get publish rate configuration for specified topic.
[**GetReplicatorDispatchRate**](NonPersistentTopicApi.md#GetReplicatorDispatchRate) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Get replicatorDispatchRate config for specified topic.
[**GetRetention**](NonPersistentTopicApi.md#GetRetention) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/retention | Get retention configuration for specified topic.
[**GetStats**](NonPersistentTopicApi.md#GetStats) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/stats | Get the stats for the topic.
[**GetSubscribeRate**](NonPersistentTopicApi.md#GetSubscribeRate) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Get subscribe rate configuration for specified topic.
[**GetSubscriptionDispatchRate**](NonPersistentTopicApi.md#GetSubscriptionDispatchRate) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Get subscription message dispatch rate configuration for specified topic.
[**GetSubscriptionTypesEnabled**](NonPersistentTopicApi.md#GetSubscriptionTypesEnabled) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionTypesEnabled | Get is enable sub type fors specified topic.
[**GetSubscriptions**](NonPersistentTopicApi.md#GetSubscriptions) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptions | Get the list of persistent subscriptions for a given topic.
[**GrantPermissionsOnTopic**](NonPersistentTopicApi.md#GrantPermissionsOnTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Grant a new permission to a role on a single topic.
[**OffloadStatus**](NonPersistentTopicApi.md#OffloadStatus) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**PeekNthMessage**](NonPersistentTopicApi.md#PeekNthMessage) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/position/{messagePosition} | Peek nth message on a topic subscription.
[**RemoveBacklogQuota**](NonPersistentTopicApi.md#RemoveBacklogQuota) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuota | Remove a backlog quota policy from a topic.
[**RemoveCompactionThreshold**](NonPersistentTopicApi.md#RemoveCompactionThreshold) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Remove compaction threshold configuration for specified topic.
[**RemoveDeduplication**](NonPersistentTopicApi.md#RemoveDeduplication) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Remove deduplication configuration for specified topic.
[**RemoveDispatchRate**](NonPersistentTopicApi.md#RemoveDispatchRate) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Remove message dispatch rate configuration for specified topic.
[**RemoveMaxConsumers**](NonPersistentTopicApi.md#RemoveMaxConsumers) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Remove maxConsumers config for specified topic.
[**RemoveMaxConsumersPerSubscription**](NonPersistentTopicApi.md#RemoveMaxConsumersPerSubscription) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Remove max consumers per subscription configuration for specified topic.
[**RemoveMaxMessageSize**](NonPersistentTopicApi.md#RemoveMaxMessageSize) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Remove maxMessageSize config for specified topic.
[**RemoveMaxProducers**](NonPersistentTopicApi.md#RemoveMaxProducers) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Remove maxProducers config for specified topic.
[**RemoveMaxSubscriptionsPerTopic**](NonPersistentTopicApi.md#RemoveMaxSubscriptionsPerTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Remove maxSubscriptionsPerTopic config for specified topic.
[**RemoveMessageTTL**](NonPersistentTopicApi.md#RemoveMessageTTL) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Remove message TTL in seconds for a topic
[**RemoveOffloadPolicies**](NonPersistentTopicApi.md#RemoveOffloadPolicies) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Delete offload policies on a topic.
[**RemovePersistence**](NonPersistentTopicApi.md#RemovePersistence) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Remove configuration of persistence policies for specified topic.
[**RemovePublishRate**](NonPersistentTopicApi.md#RemovePublishRate) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Remove message publish rate configuration for specified topic.
[**RemoveReplicatorDispatchRate**](NonPersistentTopicApi.md#RemoveReplicatorDispatchRate) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Remove replicatorDispatchRate config for specified topic.
[**RemoveRetention**](NonPersistentTopicApi.md#RemoveRetention) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/retention | Remove retention configuration for specified topic.
[**RemoveSubscribeRate**](NonPersistentTopicApi.md#RemoveSubscribeRate) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Remove subscribe rate configuration for specified topic.
[**RemoveSubscriptionDispatchRate**](NonPersistentTopicApi.md#RemoveSubscriptionDispatchRate) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Remove subscription message dispatch rate configuration for specified topic.
[**ResetCursor**](NonPersistentTopicApi.md#ResetCursor) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor/{timestamp} | Reset subscription to message position closest to absolute timestamp (in ms).
[**ResetCursorOnPosition**](NonPersistentTopicApi.md#ResetCursorOnPosition) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor | Reset subscription to message position closest to given position.
[**RevokePermissionsOnTopic**](NonPersistentTopicApi.md#RevokePermissionsOnTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Revoke permissions on a topic.
[**SetBacklogQuota**](NonPersistentTopicApi.md#SetBacklogQuota) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuota | Set a backlog quota for a topic.
[**SetCompactionThreshold**](NonPersistentTopicApi.md#SetCompactionThreshold) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Set compaction threshold configuration for specified topic.
[**SetDeduplication**](NonPersistentTopicApi.md#SetDeduplication) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Set deduplication enabled on a topic.
[**SetDeduplicationSnapshotInterval**](NonPersistentTopicApi.md#SetDeduplicationSnapshotInterval) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Set deduplicationSnapshotInterval config on a topic.
[**SetDelayedDeliveryPolicies**](NonPersistentTopicApi.md#SetDelayedDeliveryPolicies) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**SetDispatchRate**](NonPersistentTopicApi.md#SetDispatchRate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Set message dispatch rate configuration for specified topic.
[**SetInactiveTopicPolicies**](NonPersistentTopicApi.md#SetInactiveTopicPolicies) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Set inactive topic policies on a topic.
[**SetMaxConsumers**](NonPersistentTopicApi.md#SetMaxConsumers) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Set maxConsumers config for specified topic.
[**SetMaxConsumersPerSubscription**](NonPersistentTopicApi.md#SetMaxConsumersPerSubscription) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Set max consumers per subscription configuration for specified topic.
[**SetMaxMessageSize**](NonPersistentTopicApi.md#SetMaxMessageSize) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Set maxMessageSize config for specified topic.
[**SetMaxProducers**](NonPersistentTopicApi.md#SetMaxProducers) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Set maxProducers config for specified topic.
[**SetMaxSubscriptionsPerTopic**](NonPersistentTopicApi.md#SetMaxSubscriptionsPerTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Set maxSubscriptionsPerTopic config for specified topic.
[**SetMaxUnackedMessagesOnConsumer**](NonPersistentTopicApi.md#SetMaxUnackedMessagesOnConsumer) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Set max unacked messages per consumer config on a topic.
[**SetMaxUnackedMessagesOnSubscription**](NonPersistentTopicApi.md#SetMaxUnackedMessagesOnSubscription) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Set max unacked messages per subscription config on a topic.
[**SetMessageTTL**](NonPersistentTopicApi.md#SetMessageTTL) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Set message TTL in seconds for a topic
[**SetOffloadPolicies**](NonPersistentTopicApi.md#SetOffloadPolicies) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Set offload policies on a topic.
[**SetPersistence**](NonPersistentTopicApi.md#SetPersistence) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Set configuration of persistence policies for specified topic.
[**SetPublishRate**](NonPersistentTopicApi.md#SetPublishRate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Set message publish rate configuration for specified topic.
[**SetReplicatedSubscriptionStatus**](NonPersistentTopicApi.md#SetReplicatedSubscriptionStatus) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/replicatedSubscriptionStatus | Enable or disable a replicated subscription on a topic.
[**SetReplicatorDispatchRate**](NonPersistentTopicApi.md#SetReplicatorDispatchRate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Set replicatorDispatchRate config for specified topic.
[**SetRetention**](NonPersistentTopicApi.md#SetRetention) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/retention | Set retention configuration for specified topic.
[**SetSubscribeRate**](NonPersistentTopicApi.md#SetSubscribeRate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Set subscribe rate configuration for specified topic.
[**SetSubscriptionDispatchRate**](NonPersistentTopicApi.md#SetSubscriptionDispatchRate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Set subscription message dispatch rate configuration for specified topic.
[**SetSubscriptionTypesEnabled**](NonPersistentTopicApi.md#SetSubscriptionTypesEnabled) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionTypesEnabled | Set is enable sub types for specified topic
[**SkipAllMessages**](NonPersistentTopicApi.md#SkipAllMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip_all | Skip all messages on a topic subscription.
[**SkipMessages**](NonPersistentTopicApi.md#SkipMessages) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip/{numMessages} | Skipping messages on a topic subscription.
[**Terminate**](NonPersistentTopicApi.md#Terminate) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/terminate | Terminate a topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog
[**TerminatePartitionedTopic**](NonPersistentTopicApi.md#TerminatePartitionedTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/terminate/partitions | Terminate all partitioned topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog
[**TriggerOffload**](NonPersistentTopicApi.md#TriggerOffload) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**TruncateTopic**](NonPersistentTopicApi.md#TruncateTopic) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/truncate | Truncate a topic.
[**UnloadTopic**](NonPersistentTopicApi.md#UnloadTopic) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/unload | Unload a topic
[**UpdatePartitionedTopic**](NonPersistentTopicApi.md#UpdatePartitionedTopic) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Increment partitions of an existing partitioned topic.



## Compact

> Compact(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Trigger a compaction operation on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.Compact(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.Compact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompactionStatus

> LongRunningProcessStatus CompactionStatus(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get the status of a compaction operation for a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CompactionStatus(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CompactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompactionStatus`: LongRunningProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.CompactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

[**LongRunningProcessStatus**](LongRunningProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMissedPartitions

> CreateMissedPartitions(ctx, tenant, namespace, topic).Execute()

Create missed partitions of an existing partitioned topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CreateMissedPartitions(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateMissedPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMissedPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonPartitionedTopic

> CreateNonPartitionedTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Create a non-partitioned topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CreateNonPartitionedTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateNonPartitionedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonPartitionedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartitionedTopic

> CreatePartitionedTopic(ctx, tenant, namespace, topic).Body(body).CreateLocalTopicOnly(createLocalTopicOnly).Execute()

Create a partitioned topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    body := int32(56) // int32 | The number of partitions for the topic
    createLocalTopicOnly := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CreatePartitionedTopic(context.Background(), tenant, namespace, topic).Body(body).CreateLocalTopicOnly(createLocalTopicOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreatePartitionedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartitionedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **int32** | The number of partitions for the topic | 
 **createLocalTopicOnly** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> CreateSubscription(ctx, tenant, namespace, topic, subscriptionName).Authoritative(authoritative).Replicated(replicated).MessageId(messageId).Execute()

Create a subscription on the topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subscriptionName := "subscriptionName_example" // string | Subscription to create position on
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    replicated := true // bool | Is replicated required to perform this operation (optional)
    messageId := *openapiclient.NewMessageIdImpl() // MessageIdImpl | messageId where to create the subscription. It can be 'latest', 'earliest' or (ledgerId:entryId) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CreateSubscription(context.Background(), tenant, namespace, topic, subscriptionName).Authoritative(authoritative).Replicated(replicated).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subscriptionName** | **string** | Subscription to create position on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **replicated** | **bool** | Is replicated required to perform this operation | 
 **messageId** | [**MessageIdImpl**](MessageIdImpl.md) | messageId where to create the subscription. It can be &#39;latest&#39;, &#39;earliest&#39; or (ledgerId:entryId) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeduplicationSnapshotInterval

> DeleteDeduplicationSnapshotInterval(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete deduplicationSnapshotInterval config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteDeduplicationSnapshotInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeduplicationSnapshotIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelayedDeliveryPolicies

> DeleteDelayedDeliveryPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Set delayed delivery messages config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteDelayedDeliveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelayedDeliveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInactiveTopicPolicies

> DeleteInactiveTopicPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete inactive topic policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteInactiveTopicPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInactiveTopicPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMaxUnackedMessagesOnConsumer

> DeleteMaxUnackedMessagesOnConsumer(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete max unacked messages per consumer config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteMaxUnackedMessagesOnConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaxUnackedMessagesOnConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMaxUnackedMessagesOnSubscription

> DeleteMaxUnackedMessagesOnSubscription(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete max unacked messages per subscription config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteMaxUnackedMessagesOnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMaxUnackedMessagesOnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartitionedTopic

> DeletePartitionedTopic(ctx, tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()

Delete a partitioned topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    force := true // bool | Stop all producer/consumer/replicator and delete topic forcefully (optional) (default to false)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    deleteSchema := true // bool | Delete the topic's schema storage (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeletePartitionedTopic(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeletePartitionedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartitionedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Stop all producer/consumer/replicator and delete topic forcefully | [default to false]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **deleteSchema** | **bool** | Delete the topic&#39;s schema storage | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, tenant, namespace, topic, subName).Force(force).Authoritative(authoritative).Execute()

Delete a subscription.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscription to be deleted
    force := true // bool | Disconnect and close all consumers and delete subscription forcefully (optional) (default to false)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteSubscription(context.Background(), tenant, namespace, topic, subName).Force(force).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscription to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **force** | **bool** | Disconnect and close all consumers and delete subscription forcefully | [default to false]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopic

> DeleteTopic(ctx, tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()

Delete a topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    force := true // bool | Stop all producer/consumer/replicator and delete topic forcefully (optional) (default to false)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    deleteSchema := true // bool | Delete the topic's schema storage (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteTopic(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Stop all producer/consumer/replicator and delete topic forcefully | [default to false]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **deleteSchema** | **bool** | Delete the topic&#39;s schema storage | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExamineMessage

> ExamineMessage(ctx, tenant, namespace, topic).InitialPosition(initialPosition).MessagePosition(messagePosition).Authoritative(authoritative).Execute()

Examine a specific message on a topic by position relative to the earliest or the latest message.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    initialPosition := "initialPosition_example" // string | Relative start position to examine message.It can be 'latest' or 'earliest' (optional) (default to "latest")
    messagePosition := int64(789) // int64 | The position of messages (default 1) (optional) (default to 1)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ExamineMessage(context.Background(), tenant, namespace, topic).InitialPosition(initialPosition).MessagePosition(messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExamineMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiExamineMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **initialPosition** | **string** | Relative start position to examine message.It can be &#39;latest&#39; or &#39;earliest&#39; | [default to &quot;latest&quot;]
 **messagePosition** | **int64** | The position of messages (default 1) | [default to 1]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireMessagesForAllSubscriptions

> ExpireMessagesForAllSubscriptions(ctx, tenant, namespace, topic, expireTimeInSeconds).Authoritative(authoritative).Execute()

Expiry messages on all subscriptions of topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    expireTimeInSeconds := int32(56) // int32 | Expires beyond the specified number of seconds (default to 0)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ExpireMessagesForAllSubscriptions(context.Background(), tenant, namespace, topic, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExpireMessagesForAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**expireTimeInSeconds** | **int32** | Expires beyond the specified number of seconds | [default to 0]

### Other Parameters

Other parameters are passed through a pointer to a apiExpireMessagesForAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireTopicMessages

> ExpireTopicMessages(ctx, tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()

Expiry messages on a topic subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscription to be Expiry messages on
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    messageId := *openapiclient.NewResetCursorData() // ResetCursorData | messageId to reset back to (ledgerId:entryId) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ExpireTopicMessages(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExpireTopicMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscription to be Expiry messages on | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireTopicMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **messageId** | [**ResetCursorData**](ResetCursorData.md) | messageId to reset back to (ledgerId:entryId) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireTopicMessages_0

> ExpireTopicMessages_0(ctx, tenant, namespace, topic, subName, expireTimeInSeconds).Authoritative(authoritative).Execute()

Expiry messages on a topic subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscription to be Expiry messages on
    expireTimeInSeconds := int32(56) // int32 | Expires beyond the specified number of seconds (default to 0)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ExpireTopicMessages_0(context.Background(), tenant, namespace, topic, subName, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExpireTopicMessages_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscription to be Expiry messages on | 
**expireTimeInSeconds** | **int32** | Expires beyond the specified number of seconds | [default to 0]

### Other Parameters

Other parameters are passed through a pointer to a apiExpireTopicMessages_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBacklog

> PersistentOfflineTopicStats GetBacklog(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get estimated backlog for offline topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetBacklog(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetBacklog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBacklog`: PersistentOfflineTopicStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetBacklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBacklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

[**PersistentOfflineTopicStats**](PersistentOfflineTopicStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBacklogQuotaMap

> GetBacklogQuotaMap(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get backlog quota map on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetBacklogQuotaMap(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetBacklogQuotaMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBacklogQuotaMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompactionThreshold

> GetCompactionThreshold(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get compaction threshold configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetCompactionThreshold(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetCompactionThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompactionThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeduplication

> GetDeduplication(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get deduplication configuration of a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDeduplication(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDeduplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeduplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeduplicationSnapshotInterval

> GetDeduplicationSnapshotInterval(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get deduplicationSnapshotInterval config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDeduplicationSnapshotInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeduplicationSnapshotIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelayedDeliveryPolicies

> GetDelayedDeliveryPolicies(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get delayed delivery messages config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDelayedDeliveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelayedDeliveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispatchRate

> GetDispatchRate(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInactiveTopicPolicies

> GetInactiveTopicPolicies(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get inactive topic policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetInactiveTopicPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInactiveTopicPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInternalStats

> PersistentTopicInternalStats GetInternalStats(ctx, tenant, namespace, topic).Authoritative(authoritative).Metadata(metadata).Execute()

Get the internal stats for the topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    metadata := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetInternalStats(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetInternalStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInternalStats`: PersistentTopicInternalStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetInternalStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **metadata** | **bool** |  | [default to false]

### Return type

[**PersistentTopicInternalStats**](PersistentTopicInternalStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastMessageId

> GetLastMessageId(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Return the last commit message id of topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetLastMessageId(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetLastMessageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastMessageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList

> []string GetList(ctx, tenant, namespace).Execute()

Get the list of non-persistent topics under a namespace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetList(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetList`: []string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListFromBundle

> []string GetListFromBundle(ctx, tenant, namespace, bundle).Execute()

Get the list of non-persistent topics under a namespace bundle.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    bundle := "bundle_example" // string | Bundle range of a topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetListFromBundle(context.Background(), tenant, namespace, bundle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetListFromBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListFromBundle`: []string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetListFromBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**bundle** | **string** | Bundle range of a topic | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListFromBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedLedgerInfo

> GetManagedLedgerInfo(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get the stored topic metadata.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetManagedLedgerInfo(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetManagedLedgerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedLedgerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxConsumers

> GetMaxConsumers(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get maxConsumers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxConsumers(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxConsumersPerSubscription

> GetMaxConsumersPerSubscription(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get max consumers per subscription configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxConsumersPerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxConsumersPerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxMessageSize

> GetMaxMessageSize(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get maxMessageSize config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxMessageSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxMessageSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxProducers

> GetMaxProducers(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get maxProducers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxProducers(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxProducers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxProducersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxSubscriptionsPerTopic

> GetMaxSubscriptionsPerTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get maxSubscriptionsPerTopic config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxSubscriptionsPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxSubscriptionsPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxUnackedMessagesOnConsumer

> GetMaxUnackedMessagesOnConsumer(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get max unacked messages per consumer config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxUnackedMessagesOnConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxUnackedMessagesOnConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxUnackedMessagesOnSubscription

> GetMaxUnackedMessagesOnSubscription(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get max unacked messages per subscription config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxUnackedMessagesOnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxUnackedMessagesOnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageById

> GetMessageById(ctx, tenant, namespace, topic, ledgerId, entryId).Authoritative(authoritative).Execute()

Get message by its messageId.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    ledgerId := int64(789) // int64 | The ledger id
    entryId := int64(789) // int64 | The entry id
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMessageById(context.Background(), tenant, namespace, topic, ledgerId, entryId).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMessageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**ledgerId** | **int64** | The ledger id | 
**entryId** | **int64** | The entry id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageTTL

> GetMessageTTL(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get message TTL in seconds for a topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMessageTTL(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMessageTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadPolicies

> GetOffloadPolicies(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get offload policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetOffloadPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetOffloadPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffloadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartitionedMetadata

> PartitionedTopicMetadata GetPartitionedMetadata(ctx, tenant, namespace, topic).Authoritative(authoritative).CheckAllowAutoCreation(checkAllowAutoCreation).Execute()

Get partitioned topic metadata.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    checkAllowAutoCreation := true // bool | Is check configuration required to automatically create topic (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedMetadata(context.Background(), tenant, namespace, topic).Authoritative(authoritative).CheckAllowAutoCreation(checkAllowAutoCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedMetadata`: PartitionedTopicMetadata
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPartitionedMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionedMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **checkAllowAutoCreation** | **bool** | Is check configuration required to automatically create topic | [default to false]

### Return type

[**PartitionedTopicMetadata**](PartitionedTopicMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartitionedStats

> GetPartitionedStats(ctx, tenant, namespace, topic).PerPartition(perPartition).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()

Get the stats for the partitioned topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    perPartition := true // bool | Get per partition stats (optional) (default to true)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    getPreciseBacklog := true // bool | If return precise backlog or imprecise backlog (optional) (default to false)
    subscriptionBacklogSize := true // bool | If return backlog size for each subscription, require locking on ledger so be careful not to use when there's heavy traffic. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedStats(context.Background(), tenant, namespace, topic).PerPartition(perPartition).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPartition** | **bool** | Get per partition stats | [default to true]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **getPreciseBacklog** | **bool** | If return precise backlog or imprecise backlog | [default to false]
 **subscriptionBacklogSize** | **bool** | If return backlog size for each subscription, require locking on ledger so be careful not to use when there&#39;s heavy traffic. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartitionedTopicList

> []string GetPartitionedTopicList(ctx, tenant, namespace).Execute()

Get the list of partitioned topics under a namespace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedTopicList(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedTopicList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedTopicList`: []string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPartitionedTopicList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionedTopicListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionsOnTopic

> map[string][]string GetPermissionsOnTopic(ctx, tenant, namespace, topic).Execute()

Get permissions on a topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPermissionsOnTopic(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPermissionsOnTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionsOnTopic`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPermissionsOnTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsOnTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**map[string][]string**](set.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersistence

> GetPersistence(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get configuration of persistence policies for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPersistence(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPersistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublishRate

> GetPublishRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get publish rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPublishRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublishRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicatorDispatchRate

> GetReplicatorDispatchRate(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get replicatorDispatchRate config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetReplicatorDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicatorDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> GetRetention(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get retention configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetRetention(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStats

> NonPersistentTopicStats GetStats(ctx, tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()

Get the stats for the topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    getPreciseBacklog := true // bool | If return precise backlog or imprecise backlog (optional) (default to false)
    subscriptionBacklogSize := true // bool | If return backlog size for each subscription, require locking on ledger so be careful not to use when there's heavy traffic. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetStats(context.Background(), tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStats`: NonPersistentTopicStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **getPreciseBacklog** | **bool** | If return precise backlog or imprecise backlog | [default to false]
 **subscriptionBacklogSize** | **bool** | If return backlog size for each subscription, require locking on ledger so be careful not to use when there&#39;s heavy traffic. | [default to false]

### Return type

[**NonPersistentTopicStats**](NonPersistentTopicStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribeRate

> GetSubscribeRate(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get subscribe rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscribeRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscribeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscribeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionDispatchRate

> GetSubscriptionDispatchRate(ctx, tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()

Get subscription message dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    applied := true // bool |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscriptionDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applied** | **bool** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionTypesEnabled

> GetSubscriptionTypesEnabled(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get is enable sub type fors specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscriptionTypesEnabled(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscriptionTypesEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionTypesEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> GetSubscriptions(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Get the list of persistent subscriptions for a given topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscriptions(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantPermissionsOnTopic

> GrantPermissionsOnTopic(ctx, tenant, namespace, topic, role).Body(body).Execute()

Grant a new permission to a role on a single topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    role := "role_example" // string | Client role to which grant permissions
    body := []string{"Property_example"} // []string | Actions to be granted (produce,functions,consume) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GrantPermissionsOnTopic(context.Background(), tenant, namespace, topic, role).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GrantPermissionsOnTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**role** | **string** | Client role to which grant permissions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPermissionsOnTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **[]string** | Actions to be granted (produce,functions,consume) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OffloadStatus

> OffloadProcessStatus OffloadStatus(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Offload a prefix of a topic to long term storage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.OffloadStatus(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.OffloadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OffloadStatus`: OffloadProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.OffloadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffloadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

[**OffloadProcessStatus**](OffloadProcessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeekNthMessage

> PeekNthMessage(ctx, tenant, namespace, topic, subName, messagePosition).Authoritative(authoritative).Execute()

Peek nth message on a topic subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscribed message expired
    messagePosition := int32(56) // int32 | The number of messages (default 1) (default to 1)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.PeekNthMessage(context.Background(), tenant, namespace, topic, subName, messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.PeekNthMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscribed message expired | 
**messagePosition** | **int32** | The number of messages (default 1) | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiPeekNthMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBacklogQuota

> RemoveBacklogQuota(ctx, tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Authoritative(authoritative).Execute()

Remove a backlog quota policy from a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    backlogQuotaType := "backlogQuotaType_example" // string |  (optional)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveBacklogQuota(context.Background(), tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveBacklogQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBacklogQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backlogQuotaType** | **string** |  | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompactionThreshold

> RemoveCompactionThreshold(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove compaction threshold configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveCompactionThreshold(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveCompactionThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCompactionThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDeduplication

> RemoveDeduplication(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove deduplication configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveDeduplication(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveDeduplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeduplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDispatchRate

> RemoveDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove message dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaxConsumers

> RemoveMaxConsumers(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove maxConsumers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxConsumers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaxConsumersPerSubscription

> RemoveMaxConsumersPerSubscription(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove max consumers per subscription configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxConsumersPerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxConsumersPerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaxMessageSize

> RemoveMaxMessageSize(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove maxMessageSize config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxMessageSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxMessageSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaxProducers

> RemoveMaxProducers(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove maxProducers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxProducers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxProducers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxProducersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaxSubscriptionsPerTopic

> RemoveMaxSubscriptionsPerTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove maxSubscriptionsPerTopic config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxSubscriptionsPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxSubscriptionsPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMessageTTL

> RemoveMessageTTL(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove message TTL in seconds for a topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMessageTTL(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMessageTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMessageTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOffloadPolicies

> RemoveOffloadPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Delete offload policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveOffloadPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveOffloadPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOffloadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePersistence

> RemovePersistence(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove configuration of persistence policies for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemovePersistence(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemovePersistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePublishRate

> RemovePublishRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove message publish rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemovePublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemovePublishRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePublishRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveReplicatorDispatchRate

> RemoveReplicatorDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove replicatorDispatchRate config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveReplicatorDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReplicatorDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRetention

> RemoveRetention(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove retention configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveRetention(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSubscribeRate

> RemoveSubscribeRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Remove subscribe rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool |  (optional) (default to false)
    body := *openapiclient.NewSubscribeRate() // SubscribeRate | Subscribe rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveSubscribeRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveSubscribeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscribeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]
 **body** | [**SubscribeRate**](SubscribeRate.md) | Subscribe rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSubscriptionDispatchRate

> RemoveSubscriptionDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Remove subscription message dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveSubscriptionDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscriptionDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCursor

> ResetCursor(ctx, tenant, namespace, topic, subName, timestamp).Authoritative(authoritative).Execute()

Reset subscription to message position closest to absolute timestamp (in ms).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscription to reset position on
    timestamp := int64(789) // int64 | the timestamp to reset back
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ResetCursor(context.Background(), tenant, namespace, topic, subName, timestamp).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ResetCursor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscription to reset position on | 
**timestamp** | **int64** | the timestamp to reset back | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCursorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCursorOnPosition

> ResetCursorOnPosition(ctx, tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()

Reset subscription to message position closest to given position.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Subscription to reset position on
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    messageId := *openapiclient.NewResetCursorData() // ResetCursorData | messageId to reset back to (ledgerId:entryId) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ResetCursorOnPosition(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ResetCursorOnPosition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Subscription to reset position on | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCursorOnPositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **messageId** | [**ResetCursorData**](ResetCursorData.md) | messageId to reset back to (ledgerId:entryId) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePermissionsOnTopic

> RevokePermissionsOnTopic(ctx, tenant, namespace, topic, role).Execute()

Revoke permissions on a topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    role := "role_example" // string | Client role to which grant permissions

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RevokePermissionsOnTopic(context.Background(), tenant, namespace, topic, role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RevokePermissionsOnTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**role** | **string** | Client role to which grant permissions | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsOnTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBacklogQuota

> SetBacklogQuota(ctx, tenant, namespace, topic).Authoritative(authoritative).BacklogQuotaType(backlogQuotaType).Execute()

Set a backlog quota for a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    backlogQuotaType := "backlogQuotaType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetBacklogQuota(context.Background(), tenant, namespace, topic).Authoritative(authoritative).BacklogQuotaType(backlogQuotaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetBacklogQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBacklogQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **backlogQuotaType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCompactionThreshold

> SetCompactionThreshold(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set compaction threshold configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int64(789) // int64 | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetCompactionThreshold(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetCompactionThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCompactionThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int64** | Dispatch rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeduplication

> SetDeduplication(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set deduplication enabled on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := true // bool | DeduplicationEnabled policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDeduplication(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDeduplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeduplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **bool** | DeduplicationEnabled policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeduplicationSnapshotInterval

> SetDeduplicationSnapshotInterval(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set deduplicationSnapshotInterval config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | Interval to take deduplication snapshot for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDeduplicationSnapshotInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeduplicationSnapshotIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | Interval to take deduplication snapshot for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDelayedDeliveryPolicies

> SetDelayedDeliveryPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set delayed delivery messages config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewDelayedDeliveryPolicies() // DelayedDeliveryPolicies | Delayed delivery policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDelayedDeliveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDelayedDeliveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**DelayedDeliveryPolicies**](DelayedDeliveryPolicies.md) | Delayed delivery policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDispatchRate

> SetDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set message dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Dispatch rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInactiveTopicPolicies

> SetInactiveTopicPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set inactive topic policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewInactiveTopicPolicies() // InactiveTopicPolicies | inactive topic policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetInactiveTopicPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInactiveTopicPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**InactiveTopicPolicies**](InactiveTopicPolicies.md) | inactive topic policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxConsumers

> SetMaxConsumers(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set maxConsumers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | The max consumers of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxConsumers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | The max consumers of the topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxConsumersPerSubscription

> SetMaxConsumersPerSubscription(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set max consumers per subscription configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxConsumersPerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxConsumersPerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | Dispatch rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxMessageSize

> SetMaxMessageSize(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set maxMessageSize config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | The max message size of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxMessageSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxMessageSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | The max message size of the topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxProducers

> SetMaxProducers(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set maxProducers config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | The max producers of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxProducers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxProducers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxProducersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | The max producers of the topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxSubscriptionsPerTopic

> SetMaxSubscriptionsPerTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set maxSubscriptionsPerTopic config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | The max subscriptions of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxSubscriptionsPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxSubscriptionsPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | The max subscriptions of the topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxUnackedMessagesOnConsumer

> SetMaxUnackedMessagesOnConsumer(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set max unacked messages per consumer config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | Max unacked messages on consumer policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxUnackedMessagesOnConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesOnConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | Max unacked messages on consumer policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxUnackedMessagesOnSubscription

> SetMaxUnackedMessagesOnSubscription(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set max unacked messages per subscription config on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := int32(56) // int32 | Max unacked messages on subscription policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxUnackedMessagesOnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesOnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **int32** | Max unacked messages on subscription policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMessageTTL

> SetMessageTTL(ctx, tenant, namespace, topic).MessageTTL(messageTTL).Authoritative(authoritative).Execute()

Set message TTL in seconds for a topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    messageTTL := int32(56) // int32 | TTL in seconds for the specified namespace
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMessageTTL(context.Background(), tenant, namespace, topic).MessageTTL(messageTTL).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMessageTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMessageTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **messageTTL** | **int32** | TTL in seconds for the specified namespace | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOffloadPolicies

> SetOffloadPolicies(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set offload policies on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewOffloadPoliciesImpl() // OffloadPoliciesImpl | Offload policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetOffloadPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetOffloadPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOffloadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**OffloadPoliciesImpl**](OffloadPoliciesImpl.md) | Offload policies for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPersistence

> SetPersistence(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set configuration of persistence policies for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewPersistencePolicies() // PersistencePolicies | Bookkeeper persistence policies for specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetPersistence(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetPersistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**PersistencePolicies**](PersistencePolicies.md) | Bookkeeper persistence policies for specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPublishRate

> SetPublishRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set message publish rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewPublishRate() // PublishRate | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetPublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetPublishRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPublishRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**PublishRate**](PublishRate.md) | Dispatch rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReplicatedSubscriptionStatus

> SetReplicatedSubscriptionStatus(ctx, tenant, namespace, topic, subName).Body(body).Authoritative(authoritative).Execute()

Enable or disable a replicated subscription on a topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Name of subscription
    body := true // bool | Whether to enable replicated subscription
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetReplicatedSubscriptionStatus(context.Background(), tenant, namespace, topic, subName).Body(body).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetReplicatedSubscriptionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Name of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReplicatedSubscriptionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **bool** | Whether to enable replicated subscription | 
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReplicatorDispatchRate

> SetReplicatorDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set replicatorDispatchRate config for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Replicator dispatch rate of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetReplicatorDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReplicatorDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Replicator dispatch rate of the topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetention

> SetRetention(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set retention configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewRetentionPolicies() // RetentionPolicies | Retention policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetRetention(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**RetentionPolicies**](RetentionPolicies.md) | Retention policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscribeRate

> SetSubscribeRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set subscribe rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewSubscribeRate() // SubscribeRate | Subscribe rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetSubscribeRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetSubscribeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscribeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**SubscribeRate**](SubscribeRate.md) | Subscribe rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionDispatchRate

> SetSubscriptionDispatchRate(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set subscription message dispatch rate configuration for specified topic.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Subscription message dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetSubscriptionDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Subscription message dispatch rate for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionTypesEnabled

> SetSubscriptionTypesEnabled(ctx, tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()

Set is enable sub types for specified topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | 
    namespace := "namespace_example" // string | 
    topic := "topic_example" // string | 
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)
    body := []string{"Property_example"} // []string | Enable sub types for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetSubscriptionTypesEnabled(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetSubscriptionTypesEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionTypesEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **body** | **[]string** | Enable sub types for the specified topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipAllMessages

> SkipAllMessages(ctx, tenant, namespace, topic, subName).Authoritative(authoritative).Execute()

Skip all messages on a topic subscription.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Name of subscription
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SkipAllMessages(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SkipAllMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Name of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipAllMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipMessages

> SkipMessages(ctx, tenant, namespace, topic, subName, numMessages).Authoritative(authoritative).Execute()

Skipping messages on a topic subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    subName := "subName_example" // string | Name of subscription
    numMessages := int32(56) // int32 | The number of messages to skip (default to 0)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SkipMessages(context.Background(), tenant, namespace, topic, subName, numMessages).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SkipMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 
**subName** | **string** | Name of subscription | 
**numMessages** | **int32** | The number of messages to skip | [default to 0]

### Other Parameters

Other parameters are passed through a pointer to a apiSkipMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate

> map[string]interface{} Terminate(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Terminate a topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.Terminate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.Terminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Terminate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.Terminate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminatePartitionedTopic

> TerminatePartitionedTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Terminate all partitioned topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.TerminatePartitionedTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.TerminatePartitionedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminatePartitionedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerOffload

> TriggerOffload(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Offload a prefix of a topic to long term storage

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.TriggerOffload(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.TriggerOffload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerOffloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TruncateTopic

> TruncateTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Truncate a topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.TruncateTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.TruncateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTruncateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnloadTopic

> UnloadTopic(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

Unload a topic

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.UnloadTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.UnloadTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnloadTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartitionedTopic

> UpdatePartitionedTopic(ctx, tenant, namespace, topic).Body(body).UpdateLocalTopicOnly(updateLocalTopicOnly).Authoritative(authoritative).Execute()

Increment partitions of an existing partitioned topic.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenant := "tenant_example" // string | Specify the tenant
    namespace := "namespace_example" // string | Specify the namespace
    topic := "topic_example" // string | Specify topic name
    body := int32(56) // int32 | The number of partitions for the topic
    updateLocalTopicOnly := true // bool |  (optional) (default to false)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.UpdatePartitionedTopic(context.Background(), tenant, namespace, topic).Body(body).UpdateLocalTopicOnly(updateLocalTopicOnly).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.UpdatePartitionedTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 
**topic** | **string** | Specify topic name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartitionedTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **int32** | The number of partitions for the topic | 
 **updateLocalTopicOnly** | **bool** |  | [default to false]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

