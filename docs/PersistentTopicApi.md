# \PersistentTopicApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Compact**](PersistentTopicApi.md#Compact) | **Put** /persistent/{tenant}/{namespace}/{topic}/compaction | Trigger a compaction operation on a topic.
[**CompactionStatus**](PersistentTopicApi.md#CompactionStatus) | **Get** /persistent/{tenant}/{namespace}/{topic}/compaction | Get the status of a compaction operation for a topic.
[**CreateMissedPartitions**](PersistentTopicApi.md#CreateMissedPartitions) | **Post** /persistent/{tenant}/{namespace}/{topic}/createMissedPartitions | Create missed partitions of an existing partitioned topic.
[**CreateNonPartitionedTopic**](PersistentTopicApi.md#CreateNonPartitionedTopic) | **Put** /persistent/{tenant}/{namespace}/{topic} | Create a non-partitioned topic.
[**CreatePartitionedTopic**](PersistentTopicApi.md#CreatePartitionedTopic) | **Put** /persistent/{tenant}/{namespace}/{topic}/partitions | Create a partitioned topic.
[**CreateSubscription**](PersistentTopicApi.md#CreateSubscription) | **Put** /persistent/{tenant}/{namespace}/{topic}/subscription/{subscriptionName} | Create a subscription on the topic.
[**DeleteDeduplicationSnapshotInterval**](PersistentTopicApi.md#DeleteDeduplicationSnapshotInterval) | **Delete** /persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Delete deduplicationSnapshotInterval config on a topic.
[**DeleteDelayedDeliveryPolicies**](PersistentTopicApi.md#DeleteDelayedDeliveryPolicies) | **Delete** /persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**DeleteInactiveTopicPolicies**](PersistentTopicApi.md#DeleteInactiveTopicPolicies) | **Delete** /persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Delete inactive topic policies on a topic.
[**DeleteMaxUnackedMessagesOnConsumer**](PersistentTopicApi.md#DeleteMaxUnackedMessagesOnConsumer) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Delete max unacked messages per consumer config on a topic.
[**DeleteMaxUnackedMessagesOnSubscription**](PersistentTopicApi.md#DeleteMaxUnackedMessagesOnSubscription) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Delete max unacked messages per subscription config on a topic.
[**DeletePartitionedTopic**](PersistentTopicApi.md#DeletePartitionedTopic) | **Delete** /persistent/{tenant}/{namespace}/{topic}/partitions | Delete a partitioned topic.
[**DeleteSubscription**](PersistentTopicApi.md#DeleteSubscription) | **Delete** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName} | Delete a subscription.
[**DeleteTopic**](PersistentTopicApi.md#DeleteTopic) | **Delete** /persistent/{tenant}/{namespace}/{topic} | Delete a topic.
[**ExamineMessage**](PersistentTopicApi.md#ExamineMessage) | **Get** /persistent/{tenant}/{namespace}/{topic}/examinemessage | Examine a specific message on a topic by position relative to the earliest or the latest message.
[**ExpireMessagesForAllSubscriptions**](PersistentTopicApi.md#ExpireMessagesForAllSubscriptions) | **Post** /persistent/{tenant}/{namespace}/{topic}/all_subscription/expireMessages/{expireTimeInSeconds} | Expiry messages on all subscriptions of topic.
[**ExpireTopicMessages**](PersistentTopicApi.md#ExpireTopicMessages) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages | Expiry messages on a topic subscription.
[**ExpireTopicMessages_0**](PersistentTopicApi.md#ExpireTopicMessages_0) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages/{expireTimeInSeconds} | Expiry messages on a topic subscription.
[**GetBacklog**](PersistentTopicApi.md#GetBacklog) | **Get** /persistent/{tenant}/{namespace}/{topic}/backlog | Get estimated backlog for offline topic.
[**GetBacklogQuotaMap**](PersistentTopicApi.md#GetBacklogQuotaMap) | **Get** /persistent/{tenant}/{namespace}/{topic}/backlogQuotaMap | Get backlog quota map on a topic.
[**GetCompactionThreshold**](PersistentTopicApi.md#GetCompactionThreshold) | **Get** /persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Get compaction threshold configuration for specified topic.
[**GetDeduplication**](PersistentTopicApi.md#GetDeduplication) | **Get** /persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Get deduplication configuration of a topic.
[**GetDeduplicationSnapshotInterval**](PersistentTopicApi.md#GetDeduplicationSnapshotInterval) | **Get** /persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Get deduplicationSnapshotInterval config on a topic.
[**GetDelayedDeliveryPolicies**](PersistentTopicApi.md#GetDelayedDeliveryPolicies) | **Get** /persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Get delayed delivery messages config on a topic.
[**GetDispatchRate**](PersistentTopicApi.md#GetDispatchRate) | **Get** /persistent/{tenant}/{namespace}/{topic}/dispatchRate | Get dispatch rate configuration for specified topic.
[**GetInactiveTopicPolicies**](PersistentTopicApi.md#GetInactiveTopicPolicies) | **Get** /persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Get inactive topic policies on a topic.
[**GetInternalStats**](PersistentTopicApi.md#GetInternalStats) | **Get** /persistent/{tenant}/{namespace}/{topic}/internalStats | Get the internal stats for the topic.
[**GetLastMessageId**](PersistentTopicApi.md#GetLastMessageId) | **Get** /persistent/{tenant}/{namespace}/{topic}/lastMessageId | Return the last commit message id of topic
[**GetList**](PersistentTopicApi.md#GetList) | **Get** /persistent/{tenant}/{namespace} | Get the list of topics under a namespace.
[**GetManagedLedgerInfo**](PersistentTopicApi.md#GetManagedLedgerInfo) | **Get** /persistent/{tenant}/{namespace}/{topic}/internal-info | Get the stored topic metadata.
[**GetMaxConsumers**](PersistentTopicApi.md#GetMaxConsumers) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxConsumers | Get maxConsumers config for specified topic.
[**GetMaxConsumersPerSubscription**](PersistentTopicApi.md#GetMaxConsumersPerSubscription) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Get max consumers per subscription configuration for specified topic.
[**GetMaxMessageSize**](PersistentTopicApi.md#GetMaxMessageSize) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Get maxMessageSize config for specified topic.
[**GetMaxProducers**](PersistentTopicApi.md#GetMaxProducers) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxProducers | Get maxProducers config for specified topic.
[**GetMaxSubscriptionsPerTopic**](PersistentTopicApi.md#GetMaxSubscriptionsPerTopic) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Get maxSubscriptionsPerTopic config for specified topic.
[**GetMaxUnackedMessagesOnConsumer**](PersistentTopicApi.md#GetMaxUnackedMessagesOnConsumer) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Get max unacked messages per consumer config on a topic.
[**GetMaxUnackedMessagesOnSubscription**](PersistentTopicApi.md#GetMaxUnackedMessagesOnSubscription) | **Get** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Get max unacked messages per subscription config on a topic.
[**GetMessageById**](PersistentTopicApi.md#GetMessageById) | **Get** /persistent/{tenant}/{namespace}/{topic}/ledger/{ledgerId}/entry/{entryId} | Get message by its messageId.
[**GetMessageTTL**](PersistentTopicApi.md#GetMessageTTL) | **Get** /persistent/{tenant}/{namespace}/{topic}/messageTTL | Get message TTL in seconds for a topic
[**GetOffloadPolicies**](PersistentTopicApi.md#GetOffloadPolicies) | **Get** /persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Get offload policies on a topic.
[**GetPartitionedMetadata**](PersistentTopicApi.md#GetPartitionedMetadata) | **Get** /persistent/{tenant}/{namespace}/{topic}/partitions | Get partitioned topic metadata.
[**GetPartitionedStats**](PersistentTopicApi.md#GetPartitionedStats) | **Get** /persistent/{tenant}/{namespace}/{topic}/partitioned-stats | Get the stats for the partitioned topic.
[**GetPartitionedTopicList**](PersistentTopicApi.md#GetPartitionedTopicList) | **Get** /persistent/{tenant}/{namespace}/partitioned | Get the list of partitioned topics under a namespace.
[**GetPermissionsOnTopic**](PersistentTopicApi.md#GetPermissionsOnTopic) | **Get** /persistent/{tenant}/{namespace}/{topic}/permissions | Get permissions on a topic.
[**GetPersistence**](PersistentTopicApi.md#GetPersistence) | **Get** /persistent/{tenant}/{namespace}/{topic}/persistence | Get configuration of persistence policies for specified topic.
[**GetPublishRate**](PersistentTopicApi.md#GetPublishRate) | **Get** /persistent/{tenant}/{namespace}/{topic}/publishRate | Get publish rate configuration for specified topic.
[**GetReplicatorDispatchRate**](PersistentTopicApi.md#GetReplicatorDispatchRate) | **Get** /persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Get replicatorDispatchRate config for specified topic.
[**GetRetention**](PersistentTopicApi.md#GetRetention) | **Get** /persistent/{tenant}/{namespace}/{topic}/retention | Get retention configuration for specified topic.
[**GetStats**](PersistentTopicApi.md#GetStats) | **Get** /persistent/{tenant}/{namespace}/{topic}/stats | Get the stats for the topic.
[**GetSubscribeRate**](PersistentTopicApi.md#GetSubscribeRate) | **Get** /persistent/{tenant}/{namespace}/{topic}/subscribeRate | Get subscribe rate configuration for specified topic.
[**GetSubscriptionDispatchRate**](PersistentTopicApi.md#GetSubscriptionDispatchRate) | **Get** /persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Get subscription message dispatch rate configuration for specified topic.
[**GetSubscriptionTypesEnabled**](PersistentTopicApi.md#GetSubscriptionTypesEnabled) | **Get** /persistent/{tenant}/{namespace}/{topic}/subscriptionTypesEnabled | Get is enable sub type fors specified topic.
[**GetSubscriptions**](PersistentTopicApi.md#GetSubscriptions) | **Get** /persistent/{tenant}/{namespace}/{topic}/subscriptions | Get the list of persistent subscriptions for a given topic.
[**GrantPermissionsOnTopic**](PersistentTopicApi.md#GrantPermissionsOnTopic) | **Post** /persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Grant a new permission to a role on a single topic.
[**OffloadStatus**](PersistentTopicApi.md#OffloadStatus) | **Get** /persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**PeekNthMessage**](PersistentTopicApi.md#PeekNthMessage) | **Get** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/position/{messagePosition} | Peek nth message on a topic subscription.
[**RemoveBacklogQuota**](PersistentTopicApi.md#RemoveBacklogQuota) | **Delete** /persistent/{tenant}/{namespace}/{topic}/backlogQuota | Remove a backlog quota policy from a topic.
[**RemoveCompactionThreshold**](PersistentTopicApi.md#RemoveCompactionThreshold) | **Delete** /persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Remove compaction threshold configuration for specified topic.
[**RemoveDeduplication**](PersistentTopicApi.md#RemoveDeduplication) | **Delete** /persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Remove deduplication configuration for specified topic.
[**RemoveDispatchRate**](PersistentTopicApi.md#RemoveDispatchRate) | **Delete** /persistent/{tenant}/{namespace}/{topic}/dispatchRate | Remove message dispatch rate configuration for specified topic.
[**RemoveMaxConsumers**](PersistentTopicApi.md#RemoveMaxConsumers) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxConsumers | Remove maxConsumers config for specified topic.
[**RemoveMaxConsumersPerSubscription**](PersistentTopicApi.md#RemoveMaxConsumersPerSubscription) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Remove max consumers per subscription configuration for specified topic.
[**RemoveMaxMessageSize**](PersistentTopicApi.md#RemoveMaxMessageSize) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Remove maxMessageSize config for specified topic.
[**RemoveMaxProducers**](PersistentTopicApi.md#RemoveMaxProducers) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxProducers | Remove maxProducers config for specified topic.
[**RemoveMaxSubscriptionsPerTopic**](PersistentTopicApi.md#RemoveMaxSubscriptionsPerTopic) | **Delete** /persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Remove maxSubscriptionsPerTopic config for specified topic.
[**RemoveMessageTTL**](PersistentTopicApi.md#RemoveMessageTTL) | **Delete** /persistent/{tenant}/{namespace}/{topic}/messageTTL | Remove message TTL in seconds for a topic
[**RemoveOffloadPolicies**](PersistentTopicApi.md#RemoveOffloadPolicies) | **Delete** /persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Delete offload policies on a topic.
[**RemovePersistence**](PersistentTopicApi.md#RemovePersistence) | **Delete** /persistent/{tenant}/{namespace}/{topic}/persistence | Remove configuration of persistence policies for specified topic.
[**RemovePublishRate**](PersistentTopicApi.md#RemovePublishRate) | **Delete** /persistent/{tenant}/{namespace}/{topic}/publishRate | Remove message publish rate configuration for specified topic.
[**RemoveReplicatorDispatchRate**](PersistentTopicApi.md#RemoveReplicatorDispatchRate) | **Delete** /persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Remove replicatorDispatchRate config for specified topic.
[**RemoveRetention**](PersistentTopicApi.md#RemoveRetention) | **Delete** /persistent/{tenant}/{namespace}/{topic}/retention | Remove retention configuration for specified topic.
[**RemoveSubscribeRate**](PersistentTopicApi.md#RemoveSubscribeRate) | **Delete** /persistent/{tenant}/{namespace}/{topic}/subscribeRate | Remove subscribe rate configuration for specified topic.
[**RemoveSubscriptionDispatchRate**](PersistentTopicApi.md#RemoveSubscriptionDispatchRate) | **Delete** /persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Remove subscription message dispatch rate configuration for specified topic.
[**ResetCursor**](PersistentTopicApi.md#ResetCursor) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor/{timestamp} | Reset subscription to message position closest to absolute timestamp (in ms).
[**ResetCursorOnPosition**](PersistentTopicApi.md#ResetCursorOnPosition) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor | Reset subscription to message position closest to given position.
[**RevokePermissionsOnTopic**](PersistentTopicApi.md#RevokePermissionsOnTopic) | **Delete** /persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Revoke permissions on a topic.
[**SetBacklogQuota**](PersistentTopicApi.md#SetBacklogQuota) | **Post** /persistent/{tenant}/{namespace}/{topic}/backlogQuota | Set a backlog quota for a topic.
[**SetCompactionThreshold**](PersistentTopicApi.md#SetCompactionThreshold) | **Post** /persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Set compaction threshold configuration for specified topic.
[**SetDeduplication**](PersistentTopicApi.md#SetDeduplication) | **Post** /persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Set deduplication enabled on a topic.
[**SetDeduplicationSnapshotInterval**](PersistentTopicApi.md#SetDeduplicationSnapshotInterval) | **Post** /persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Set deduplicationSnapshotInterval config on a topic.
[**SetDelayedDeliveryPolicies**](PersistentTopicApi.md#SetDelayedDeliveryPolicies) | **Post** /persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**SetDispatchRate**](PersistentTopicApi.md#SetDispatchRate) | **Post** /persistent/{tenant}/{namespace}/{topic}/dispatchRate | Set message dispatch rate configuration for specified topic.
[**SetInactiveTopicPolicies**](PersistentTopicApi.md#SetInactiveTopicPolicies) | **Post** /persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Set inactive topic policies on a topic.
[**SetMaxConsumers**](PersistentTopicApi.md#SetMaxConsumers) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxConsumers | Set maxConsumers config for specified topic.
[**SetMaxConsumersPerSubscription**](PersistentTopicApi.md#SetMaxConsumersPerSubscription) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Set max consumers per subscription configuration for specified topic.
[**SetMaxMessageSize**](PersistentTopicApi.md#SetMaxMessageSize) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxMessageSize | Set maxMessageSize config for specified topic.
[**SetMaxProducers**](PersistentTopicApi.md#SetMaxProducers) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxProducers | Set maxProducers config for specified topic.
[**SetMaxSubscriptionsPerTopic**](PersistentTopicApi.md#SetMaxSubscriptionsPerTopic) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxSubscriptionsPerTopic | Set maxSubscriptionsPerTopic config for specified topic.
[**SetMaxUnackedMessagesOnConsumer**](PersistentTopicApi.md#SetMaxUnackedMessagesOnConsumer) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Set max unacked messages per consumer config on a topic.
[**SetMaxUnackedMessagesOnSubscription**](PersistentTopicApi.md#SetMaxUnackedMessagesOnSubscription) | **Post** /persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Set max unacked messages per subscription config on a topic.
[**SetMessageTTL**](PersistentTopicApi.md#SetMessageTTL) | **Post** /persistent/{tenant}/{namespace}/{topic}/messageTTL | Set message TTL in seconds for a topic
[**SetOffloadPolicies**](PersistentTopicApi.md#SetOffloadPolicies) | **Post** /persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Set offload policies on a topic.
[**SetPersistence**](PersistentTopicApi.md#SetPersistence) | **Post** /persistent/{tenant}/{namespace}/{topic}/persistence | Set configuration of persistence policies for specified topic.
[**SetPublishRate**](PersistentTopicApi.md#SetPublishRate) | **Post** /persistent/{tenant}/{namespace}/{topic}/publishRate | Set message publish rate configuration for specified topic.
[**SetReplicatedSubscriptionStatus**](PersistentTopicApi.md#SetReplicatedSubscriptionStatus) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/replicatedSubscriptionStatus | Enable or disable a replicated subscription on a topic.
[**SetReplicatorDispatchRate**](PersistentTopicApi.md#SetReplicatorDispatchRate) | **Post** /persistent/{tenant}/{namespace}/{topic}/replicatorDispatchRate | Set replicatorDispatchRate config for specified topic.
[**SetRetention**](PersistentTopicApi.md#SetRetention) | **Post** /persistent/{tenant}/{namespace}/{topic}/retention | Set retention configuration for specified topic.
[**SetSubscribeRate**](PersistentTopicApi.md#SetSubscribeRate) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscribeRate | Set subscribe rate configuration for specified topic.
[**SetSubscriptionDispatchRate**](PersistentTopicApi.md#SetSubscriptionDispatchRate) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Set subscription message dispatch rate configuration for specified topic.
[**SetSubscriptionTypesEnabled**](PersistentTopicApi.md#SetSubscriptionTypesEnabled) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscriptionTypesEnabled | Set is enable sub types for specified topic
[**SkipAllMessages**](PersistentTopicApi.md#SkipAllMessages) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip_all | Skip all messages on a topic subscription.
[**SkipMessages**](PersistentTopicApi.md#SkipMessages) | **Post** /persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip/{numMessages} | Skipping messages on a topic subscription.
[**Terminate**](PersistentTopicApi.md#Terminate) | **Post** /persistent/{tenant}/{namespace}/{topic}/terminate | Terminate a topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog
[**TerminatePartitionedTopic**](PersistentTopicApi.md#TerminatePartitionedTopic) | **Post** /persistent/{tenant}/{namespace}/{topic}/terminate/partitions | Terminate all partitioned topic. A topic that is terminated will not accept any more messages to be published and will let consumer to drain existing messages in backlog
[**TriggerOffload**](PersistentTopicApi.md#TriggerOffload) | **Put** /persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**TruncateTopic**](PersistentTopicApi.md#TruncateTopic) | **Delete** /persistent/{tenant}/{namespace}/{topic}/truncate | Truncate a topic.
[**UnloadTopic**](PersistentTopicApi.md#UnloadTopic) | **Put** /persistent/{tenant}/{namespace}/{topic}/unload | Unload a topic
[**UpdatePartitionedTopic**](PersistentTopicApi.md#UpdatePartitionedTopic) | **Post** /persistent/{tenant}/{namespace}/{topic}/partitions | Increment partitions of an existing partitioned topic.



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
    resp, r, err := api_client.PersistentTopicApi.Compact(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.Compact``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.CompactionStatus(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.CompactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompactionStatus`: LongRunningProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.CompactionStatus`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.CreateMissedPartitions(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.CreateMissedPartitions``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.CreateNonPartitionedTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.CreateNonPartitionedTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.CreatePartitionedTopic(context.Background(), tenant, namespace, topic).Body(body).CreateLocalTopicOnly(createLocalTopicOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.CreatePartitionedTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.CreateSubscription(context.Background(), tenant, namespace, topic, subscriptionName).Authoritative(authoritative).Replicated(replicated).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.CreateSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteDeduplicationSnapshotInterval``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteDelayedDeliveryPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteInactiveTopicPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteMaxUnackedMessagesOnConsumer``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteMaxUnackedMessagesOnSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeletePartitionedTopic(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeletePartitionedTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteSubscription(context.Background(), tenant, namespace, topic, subName).Force(force).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.DeleteTopic(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.DeleteTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ExamineMessage(context.Background(), tenant, namespace, topic).InitialPosition(initialPosition).MessagePosition(messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ExamineMessage``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ExpireMessagesForAllSubscriptions(context.Background(), tenant, namespace, topic, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ExpireMessagesForAllSubscriptions``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ExpireTopicMessages(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ExpireTopicMessages``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ExpireTopicMessages_0(context.Background(), tenant, namespace, topic, subName, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ExpireTopicMessages_0``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetBacklog(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetBacklog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBacklog`: PersistentOfflineTopicStats
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetBacklog`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetBacklogQuotaMap(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetBacklogQuotaMap``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetCompactionThreshold(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetCompactionThreshold``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetDeduplication(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetDeduplication``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetDeduplicationSnapshotInterval``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetDelayedDeliveryPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetInactiveTopicPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetInternalStats(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetInternalStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInternalStats`: PersistentTopicInternalStats
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetInternalStats`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetLastMessageId(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetLastMessageId``: %v\n", err)
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

Get the list of topics under a namespace.

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
    resp, r, err := api_client.PersistentTopicApi.GetList(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetList`: []string
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetList`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetManagedLedgerInfo(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetManagedLedgerInfo``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxConsumers(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxConsumers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxConsumersPerSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxMessageSize``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxProducers(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxProducers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxSubscriptionsPerTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxUnackedMessagesOnConsumer``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMaxUnackedMessagesOnSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMessageById(context.Background(), tenant, namespace, topic, ledgerId, entryId).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMessageById``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetMessageTTL(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetMessageTTL``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetOffloadPolicies(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetOffloadPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetPartitionedMetadata(context.Background(), tenant, namespace, topic).Authoritative(authoritative).CheckAllowAutoCreation(checkAllowAutoCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPartitionedMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedMetadata`: PartitionedTopicMetadata
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetPartitionedMetadata`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetPartitionedStats(context.Background(), tenant, namespace, topic).PerPartition(perPartition).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPartitionedStats``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetPartitionedTopicList(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPartitionedTopicList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedTopicList`: []string
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetPartitionedTopicList`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetPermissionsOnTopic(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPermissionsOnTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionsOnTopic`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetPermissionsOnTopic`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.GetPersistence(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPersistence``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetPublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetPublishRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetReplicatorDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetRetention(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetRetention``: %v\n", err)
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

> TopicStats GetStats(ctx, tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()

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
    resp, r, err := api_client.PersistentTopicApi.GetStats(context.Background(), tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).SubscriptionBacklogSize(subscriptionBacklogSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStats`: TopicStats
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.GetStats`: %v\n", resp)
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

[**TopicStats**](TopicStats.md)

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
    resp, r, err := api_client.PersistentTopicApi.GetSubscribeRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetSubscribeRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Applied(applied).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetSubscriptionDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetSubscriptionTypesEnabled(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetSubscriptionTypesEnabled``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GetSubscriptions(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GetSubscriptions``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.GrantPermissionsOnTopic(context.Background(), tenant, namespace, topic, role).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.GrantPermissionsOnTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.OffloadStatus(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.OffloadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OffloadStatus`: OffloadProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.OffloadStatus`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.PeekNthMessage(context.Background(), tenant, namespace, topic, subName, messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.PeekNthMessage``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveBacklogQuota(context.Background(), tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveBacklogQuota``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveCompactionThreshold(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveCompactionThreshold``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveDeduplication(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveDeduplication``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMaxConsumers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMaxConsumers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMaxConsumersPerSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMaxMessageSize``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMaxProducers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMaxProducers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMaxSubscriptionsPerTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveMessageTTL(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveMessageTTL``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveOffloadPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveOffloadPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemovePersistence(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemovePersistence``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemovePublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemovePublishRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveReplicatorDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveRetention(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveRetention``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveSubscribeRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveSubscribeRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RemoveSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RemoveSubscriptionDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ResetCursor(context.Background(), tenant, namespace, topic, subName, timestamp).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ResetCursor``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.ResetCursorOnPosition(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.ResetCursorOnPosition``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.RevokePermissionsOnTopic(context.Background(), tenant, namespace, topic, role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.RevokePermissionsOnTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetBacklogQuota(context.Background(), tenant, namespace, topic).Authoritative(authoritative).BacklogQuotaType(backlogQuotaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetBacklogQuota``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetCompactionThreshold(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetCompactionThreshold``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetDeduplication(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetDeduplication``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetDeduplicationSnapshotInterval(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetDeduplicationSnapshotInterval``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetDelayedDeliveryPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetDelayedDeliveryPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetInactiveTopicPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetInactiveTopicPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxConsumers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxConsumers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxConsumersPerSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxConsumersPerSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxMessageSize(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxMessageSize``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxProducers(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxProducers``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxSubscriptionsPerTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxUnackedMessagesOnConsumer(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxUnackedMessagesOnConsumer``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMaxUnackedMessagesOnSubscription(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMaxUnackedMessagesOnSubscription``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetMessageTTL(context.Background(), tenant, namespace, topic).MessageTTL(messageTTL).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetMessageTTL``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetOffloadPolicies(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetOffloadPolicies``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetPersistence(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetPersistence``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetPublishRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetPublishRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetReplicatedSubscriptionStatus(context.Background(), tenant, namespace, topic, subName).Body(body).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetReplicatedSubscriptionStatus``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetReplicatorDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetReplicatorDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetRetention(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetRetention``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetSubscribeRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetSubscribeRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetSubscriptionDispatchRate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetSubscriptionDispatchRate``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SetSubscriptionTypesEnabled(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SetSubscriptionTypesEnabled``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SkipAllMessages(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SkipAllMessages``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.SkipMessages(context.Background(), tenant, namespace, topic, subName, numMessages).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.SkipMessages``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.Terminate(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.Terminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Terminate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PersistentTopicApi.Terminate`: %v\n", resp)
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
    resp, r, err := api_client.PersistentTopicApi.TerminatePartitionedTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.TerminatePartitionedTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.TriggerOffload(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.TriggerOffload``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.TruncateTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.TruncateTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.UnloadTopic(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.UnloadTopic``: %v\n", err)
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
    resp, r, err := api_client.PersistentTopicApi.UpdatePartitionedTopic(context.Background(), tenant, namespace, topic).Body(body).UpdateLocalTopicOnly(updateLocalTopicOnly).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersistentTopicApi.UpdatePartitionedTopic``: %v\n", err)
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

