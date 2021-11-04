# \NonPersistentTopicApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompactNP**](NonPersistentTopicApi.md#CompactNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Trigger a compaction operation on a topic.
[**CompactionStatusNP**](NonPersistentTopicApi.md#CompactionStatusNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/compaction | Get the status of a compaction operation for a topic.
[**CreateMissedPartitionsNP**](NonPersistentTopicApi.md#CreateMissedPartitionsNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/createMissedPartitions | Create missed partitions of an existing partitioned topic.
[**CreateNonPartitionedTopicNP**](NonPersistentTopicApi.md#CreateNonPartitionedTopicNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic} | Create a non-partitioned topic.
[**CreatePartitionedTopicNP**](NonPersistentTopicApi.md#CreatePartitionedTopicNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Create a partitioned topic.
[**CreateSubscriptionNP**](NonPersistentTopicApi.md#CreateSubscriptionNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subscriptionName} | Create a subscription on the topic.
[**DeleteDeduplicationSnapshotIntervalNP**](NonPersistentTopicApi.md#DeleteDeduplicationSnapshotIntervalNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Delete deduplicationSnapshotInterval config on a topic.
[**DeleteDelayedDeliveryPoliciesNP**](NonPersistentTopicApi.md#DeleteDelayedDeliveryPoliciesNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**DeleteInactiveTopicPoliciesNP**](NonPersistentTopicApi.md#DeleteInactiveTopicPoliciesNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Delete inactive topic policies on a topic.
[**DeleteMaxUnackedMessagesOnConsumerNP**](NonPersistentTopicApi.md#DeleteMaxUnackedMessagesOnConsumerNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Delete max unacked messages per consumer config on a topic.
[**DeleteMaxUnackedMessagesOnSubscriptionNP**](NonPersistentTopicApi.md#DeleteMaxUnackedMessagesOnSubscriptionNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Delete max unacked messages per subscription config on a topic.
[**DeletePartitionedTopicNP**](NonPersistentTopicApi.md#DeletePartitionedTopicNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Delete a partitioned topic.
[**DeleteSubscriptionNP**](NonPersistentTopicApi.md#DeleteSubscriptionNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName} | Delete a subscription.
[**DeleteTopicNP**](NonPersistentTopicApi.md#DeleteTopicNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic} | Delete a topic.
[**ExamineMessageNP**](NonPersistentTopicApi.md#ExamineMessageNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/examinemessage | Examine a specific message on a topic by position relative to the earliest or the latest message.
[**ExpireMessagesForAllSubscriptionsNP**](NonPersistentTopicApi.md#ExpireMessagesForAllSubscriptionsNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/all_subscription/expireMessages/{expireTimeInSeconds} | Expiry messages on all subscriptions of topic.
[**ExpireTopicMessagesNP**](NonPersistentTopicApi.md#ExpireTopicMessagesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/expireMessages/{expireTimeInSeconds} | Expiry messages on a topic subscription.
[**GetBacklogNP**](NonPersistentTopicApi.md#GetBacklogNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/backlog | Get estimated backlog for offline topic.
[**GetBacklogQuotaMapNP**](NonPersistentTopicApi.md#GetBacklogQuotaMapNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuotaMap | Get backlog quota map on a topic.
[**GetCompactionThresholdNP**](NonPersistentTopicApi.md#GetCompactionThresholdNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Get compaction threshold configuration for specified topic.
[**GetDeduplicationEnabledNP**](NonPersistentTopicApi.md#GetDeduplicationEnabledNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Get deduplication configuration of a topic.
[**GetDeduplicationSnapshotIntervalNP**](NonPersistentTopicApi.md#GetDeduplicationSnapshotIntervalNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Get deduplicationSnapshotInterval config on a topic.
[**GetDelayedDeliveryPoliciesNP**](NonPersistentTopicApi.md#GetDelayedDeliveryPoliciesNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Get delayed delivery messages config on a topic.
[**GetDispatchRateNP**](NonPersistentTopicApi.md#GetDispatchRateNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Get dispatch rate configuration for specified topic.
[**GetInactiveTopicPoliciesNP**](NonPersistentTopicApi.md#GetInactiveTopicPoliciesNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Get inactive topic policies on a topic.
[**GetInternalStatsNP**](NonPersistentTopicApi.md#GetInternalStatsNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internalStats | Get the internal stats for the topic.
[**GetLastMessageIdNP**](NonPersistentTopicApi.md#GetLastMessageIdNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/lastMessageId | Return the last commit message id of topic
[**GetListFromBundle**](NonPersistentTopicApi.md#GetListFromBundle) | **Get** /non-persistent/{tenant}/{namespace}/{bundle} | Get the list of non-persistent topics under a namespace bundle.
[**GetListNP**](NonPersistentTopicApi.md#GetListNP) | **Get** /non-persistent/{tenant}/{namespace} | Get the list of non-persistent topics under a namespace.
[**GetManagedLedgerInfoNP**](NonPersistentTopicApi.md#GetManagedLedgerInfoNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/internal-info | Get the stored topic metadata.
[**GetMaxConsumersNP**](NonPersistentTopicApi.md#GetMaxConsumersNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Get maxConsumers config for specified topic.
[**GetMaxConsumersPerSubscriptionNP**](NonPersistentTopicApi.md#GetMaxConsumersPerSubscriptionNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Get max consumers per subscription configuration for specified topic.
[**GetMaxProducersNP**](NonPersistentTopicApi.md#GetMaxProducersNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Get maxProducers config for specified topic.
[**GetMaxUnackedMessagesOnConsumerNP**](NonPersistentTopicApi.md#GetMaxUnackedMessagesOnConsumerNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Get max unacked messages per consumer config on a topic.
[**GetMaxUnackedMessagesOnSubscriptionNP**](NonPersistentTopicApi.md#GetMaxUnackedMessagesOnSubscriptionNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Get max unacked messages per subscription config on a topic.
[**GetMessageByIdNP**](NonPersistentTopicApi.md#GetMessageByIdNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/ledger/{ledgerId}/entry/{entryId} | Get message by its messageId.
[**GetMessageTTLNP**](NonPersistentTopicApi.md#GetMessageTTLNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Get message TTL in seconds for a topic
[**GetOffloadPoliciesNP**](NonPersistentTopicApi.md#GetOffloadPoliciesNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Get offload policies on a topic.
[**GetPartitionedMetadataNP**](NonPersistentTopicApi.md#GetPartitionedMetadataNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Get partitioned topic metadata.
[**GetPartitionedStatsNP**](NonPersistentTopicApi.md#GetPartitionedStatsNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/partitioned-stats | Get the stats for the partitioned topic.
[**GetPartitionedTopicListNP**](NonPersistentTopicApi.md#GetPartitionedTopicListNP) | **Get** /non-persistent/{tenant}/{namespace}/partitioned | Get the list of partitioned topics under a namespace.
[**GetPermissionsOnTopicNP**](NonPersistentTopicApi.md#GetPermissionsOnTopicNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/permissions | Get permissions on a topic.
[**GetPersistenceNP**](NonPersistentTopicApi.md#GetPersistenceNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Get configuration of persistence policies for specified topic.
[**GetPublishRateNP**](NonPersistentTopicApi.md#GetPublishRateNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Get publish rate configuration for specified topic.
[**GetRetentionNP**](NonPersistentTopicApi.md#GetRetentionNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/retention | Get retention configuration for specified topic.
[**GetStatsNP**](NonPersistentTopicApi.md#GetStatsNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/stats | Get the stats for the topic.
[**GetSubscribeRateNP**](NonPersistentTopicApi.md#GetSubscribeRateNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Get subscribe rate configuration for specified topic.
[**GetSubscriptionDispatchRateNP**](NonPersistentTopicApi.md#GetSubscriptionDispatchRateNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Get subscription message dispatch rate configuration for specified topic.
[**GetSubscriptionsNP**](NonPersistentTopicApi.md#GetSubscriptionsNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscriptions | Get the list of persistent subscriptions for a given topic.
[**GrantPermissionsOnTopicNP**](NonPersistentTopicApi.md#GrantPermissionsOnTopicNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Grant a new permission to a role on a single topic.
[**OffloadStatusNP**](NonPersistentTopicApi.md#OffloadStatusNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**PeekNthMessageNP**](NonPersistentTopicApi.md#PeekNthMessageNP) | **Get** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/position/{messagePosition} | Peek nth message on a topic subscription.
[**RemoveBacklogQuotaNP**](NonPersistentTopicApi.md#RemoveBacklogQuotaNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuota | Remove a backlog quota policy from a topic.
[**RemoveCompactionThresholdNP**](NonPersistentTopicApi.md#RemoveCompactionThresholdNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Remove compaction threshold configuration for specified topic.
[**RemoveDeduplicationEnabledNP**](NonPersistentTopicApi.md#RemoveDeduplicationEnabledNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Remove deduplication configuration for specified topic.
[**RemoveDispatchRateNP**](NonPersistentTopicApi.md#RemoveDispatchRateNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Remove message dispatch rate configuration for specified topic.
[**RemoveMaxConsumersNP**](NonPersistentTopicApi.md#RemoveMaxConsumersNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Remove maxConsumers config for specified topic.
[**RemoveMaxConsumersPerSubscriptionNP**](NonPersistentTopicApi.md#RemoveMaxConsumersPerSubscriptionNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Remove max consumers per subscription configuration for specified topic.
[**RemoveMaxProducersNP**](NonPersistentTopicApi.md#RemoveMaxProducersNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Remove maxProducers config for specified topic.
[**RemoveMessageTTLNP**](NonPersistentTopicApi.md#RemoveMessageTTLNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Remove message TTL in seconds for a topic
[**RemoveOffloadPoliciesNP**](NonPersistentTopicApi.md#RemoveOffloadPoliciesNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Delete offload policies on a topic.
[**RemovePersistenceNP**](NonPersistentTopicApi.md#RemovePersistenceNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Remove configuration of persistence policies for specified topic.
[**RemovePublishRateNP**](NonPersistentTopicApi.md#RemovePublishRateNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Remove message publish rate configuration for specified topic.
[**RemoveRetentionNP**](NonPersistentTopicApi.md#RemoveRetentionNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/retention | Remove retention configuration for specified topic.
[**RemoveSubscribeRateNP**](NonPersistentTopicApi.md#RemoveSubscribeRateNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Remove subscribe rate configuration for specified topic.
[**RemoveSubscriptionDispatchRateNP**](NonPersistentTopicApi.md#RemoveSubscriptionDispatchRateNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Remove subscription message dispatch rate configuration for specified topic.
[**ResetCursorNP**](NonPersistentTopicApi.md#ResetCursorNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor/{timestamp} | Reset subscription to message position closest to absolute timestamp (in ms).
[**ResetCursorOnPositionNP**](NonPersistentTopicApi.md#ResetCursorOnPositionNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/resetcursor | Reset subscription to message position closest to given position.
[**RevokePermissionsOnTopicNP**](NonPersistentTopicApi.md#RevokePermissionsOnTopicNP) | **Delete** /non-persistent/{tenant}/{namespace}/{topic}/permissions/{role} | Revoke permissions on a topic.
[**SetBacklogQuotaNP**](NonPersistentTopicApi.md#SetBacklogQuotaNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/backlogQuota | Set a backlog quota for a topic.
[**SetCompactionThresholdNP**](NonPersistentTopicApi.md#SetCompactionThresholdNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/compactionThreshold | Set compaction threshold configuration for specified topic.
[**SetDeduplicationEnabledNP**](NonPersistentTopicApi.md#SetDeduplicationEnabledNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationEnabled | Set deduplication enabled on a topic.
[**SetDeduplicationSnapshotIntervalNP**](NonPersistentTopicApi.md#SetDeduplicationSnapshotIntervalNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/deduplicationSnapshotInterval | Set deduplicationSnapshotInterval config on a topic.
[**SetDelayedDeliveryPoliciesNP**](NonPersistentTopicApi.md#SetDelayedDeliveryPoliciesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/delayedDelivery | Set delayed delivery messages config on a topic.
[**SetDispatchRateNP**](NonPersistentTopicApi.md#SetDispatchRateNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/dispatchRate | Set message dispatch rate configuration for specified topic.
[**SetInactiveTopicPoliciesNP**](NonPersistentTopicApi.md#SetInactiveTopicPoliciesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/inactiveTopicPolicies | Set inactive topic policies on a topic.
[**SetMaxConsumersNP**](NonPersistentTopicApi.md#SetMaxConsumersNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumers | Set maxConsumers config for specified topic.
[**SetMaxConsumersPerSubscriptionNP**](NonPersistentTopicApi.md#SetMaxConsumersPerSubscriptionNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxConsumersPerSubscription | Set max consumers per subscription configuration for specified topic.
[**SetMaxProducersNP**](NonPersistentTopicApi.md#SetMaxProducersNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxProducers | Set maxProducers config for specified topic.
[**SetMaxUnackedMessagesOnConsumerNP**](NonPersistentTopicApi.md#SetMaxUnackedMessagesOnConsumerNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnConsumer | Set max unacked messages per consumer config on a topic.
[**SetMaxUnackedMessagesOnSubscriptionNP**](NonPersistentTopicApi.md#SetMaxUnackedMessagesOnSubscriptionNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/maxUnackedMessagesOnSubscription | Set max unacked messages per subscription config on a topic.
[**SetMessageTTLNP**](NonPersistentTopicApi.md#SetMessageTTLNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/messageTTL | Set message TTL in seconds for a topic
[**SetOffloadPoliciesNP**](NonPersistentTopicApi.md#SetOffloadPoliciesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/offloadPolicies | Set offload policies on a topic.
[**SetPersistenceNP**](NonPersistentTopicApi.md#SetPersistenceNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/persistence | Set configuration of persistence policies for specified topic.
[**SetPublishRateNP**](NonPersistentTopicApi.md#SetPublishRateNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/publishRate | Set message publish rate configuration for specified topic.
[**SetRetentionNP**](NonPersistentTopicApi.md#SetRetentionNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/retention | Set retention configuration for specified topic.
[**SetSubscribeRateNP**](NonPersistentTopicApi.md#SetSubscribeRateNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscribeRate | Set subscribe rate configuration for specified topic.
[**SetSubscriptionDispatchRateNP**](NonPersistentTopicApi.md#SetSubscriptionDispatchRateNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscriptionDispatchRate | Set subscription message dispatch rate configuration for specified topic.
[**SkipAllMessagesNP**](NonPersistentTopicApi.md#SkipAllMessagesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip_all | Skip all messages on a topic subscription.
[**SkipMessagesNP**](NonPersistentTopicApi.md#SkipMessagesNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/subscription/{subName}/skip/{numMessages} | Skipping messages on a topic subscription.
[**TriggerOffloadNP**](NonPersistentTopicApi.md#TriggerOffloadNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/offload | Offload a prefix of a topic to long term storage
[**UnloadTopicNP**](NonPersistentTopicApi.md#UnloadTopicNP) | **Put** /non-persistent/{tenant}/{namespace}/{topic}/unload | Unload a topic
[**UpdatePartitionedTopicNP**](NonPersistentTopicApi.md#UpdatePartitionedTopicNP) | **Post** /non-persistent/{tenant}/{namespace}/{topic}/partitions | Increment partitions of an existing partitioned topic.



## CompactNP

> CompactNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.CompactNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CompactNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCompactNPRequest struct via the builder pattern


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


## CompactionStatusNP

> LongRunningProcessStatus CompactionStatusNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.CompactionStatusNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CompactionStatusNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompactionStatusNP`: LongRunningProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.CompactionStatusNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCompactionStatusNPRequest struct via the builder pattern


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


## CreateMissedPartitionsNP

> CreateMissedPartitionsNP(ctx, tenant, namespace, topic).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.CreateMissedPartitionsNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateMissedPartitionsNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateMissedPartitionsNPRequest struct via the builder pattern


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


## CreateNonPartitionedTopicNP

> CreateNonPartitionedTopicNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.CreateNonPartitionedTopicNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateNonPartitionedTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateNonPartitionedTopicNPRequest struct via the builder pattern


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


## CreatePartitionedTopicNP

> CreatePartitionedTopicNP(ctx, tenant, namespace, topic).Body(body).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.CreatePartitionedTopicNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreatePartitionedTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreatePartitionedTopicNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **int32** | The number of partitions for the topic | 

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


## CreateSubscriptionNP

> CreateSubscriptionNP(ctx, tenant, namespace, topic, subscriptionName).Authoritative(authoritative).Replicated(replicated).MessageId(messageId).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.CreateSubscriptionNP(context.Background(), tenant, namespace, topic, subscriptionName).Authoritative(authoritative).Replicated(replicated).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.CreateSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateSubscriptionNPRequest struct via the builder pattern


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


## DeleteDeduplicationSnapshotIntervalNP

> DeleteDeduplicationSnapshotIntervalNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteDeduplicationSnapshotIntervalNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteDeduplicationSnapshotIntervalNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeduplicationSnapshotIntervalNPRequest struct via the builder pattern


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


## DeleteDelayedDeliveryPoliciesNP

> DeleteDelayedDeliveryPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteDelayedDeliveryPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteDelayedDeliveryPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDelayedDeliveryPoliciesNPRequest struct via the builder pattern


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


## DeleteInactiveTopicPoliciesNP

> DeleteInactiveTopicPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteInactiveTopicPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteInactiveTopicPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteInactiveTopicPoliciesNPRequest struct via the builder pattern


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


## DeleteMaxUnackedMessagesOnConsumerNP

> DeleteMaxUnackedMessagesOnConsumerNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteMaxUnackedMessagesOnConsumerNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteMaxUnackedMessagesOnConsumerNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMaxUnackedMessagesOnConsumerNPRequest struct via the builder pattern


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


## DeleteMaxUnackedMessagesOnSubscriptionNP

> DeleteMaxUnackedMessagesOnSubscriptionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.DeleteMaxUnackedMessagesOnSubscriptionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteMaxUnackedMessagesOnSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMaxUnackedMessagesOnSubscriptionNPRequest struct via the builder pattern


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


## DeletePartitionedTopicNP

> DeletePartitionedTopicNP(ctx, tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.DeletePartitionedTopicNP(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeletePartitionedTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePartitionedTopicNPRequest struct via the builder pattern


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


## DeleteSubscriptionNP

> DeleteSubscriptionNP(ctx, tenant, namespace, topic, subName).Force(force).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.DeleteSubscriptionNP(context.Background(), tenant, namespace, topic, subName).Force(force).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSubscriptionNPRequest struct via the builder pattern


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


## DeleteTopicNP

> DeleteTopicNP(ctx, tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.DeleteTopicNP(context.Background(), tenant, namespace, topic).Force(force).Authoritative(authoritative).DeleteSchema(deleteSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.DeleteTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTopicNPRequest struct via the builder pattern


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


## ExamineMessageNP

> ExamineMessageNP(ctx, tenant, namespace, topic).InitialPosition(initialPosition).MessagePosition(messagePosition).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.ExamineMessageNP(context.Background(), tenant, namespace, topic).InitialPosition(initialPosition).MessagePosition(messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExamineMessageNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExamineMessageNPRequest struct via the builder pattern


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


## ExpireMessagesForAllSubscriptionsNP

> ExpireMessagesForAllSubscriptionsNP(ctx, tenant, namespace, topic, expireTimeInSeconds).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.ExpireMessagesForAllSubscriptionsNP(context.Background(), tenant, namespace, topic, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExpireMessagesForAllSubscriptionsNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExpireMessagesForAllSubscriptionsNPRequest struct via the builder pattern


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


## ExpireTopicMessagesNP

> ExpireTopicMessagesNP(ctx, tenant, namespace, topic, subName, expireTimeInSeconds).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.ExpireTopicMessagesNP(context.Background(), tenant, namespace, topic, subName, expireTimeInSeconds).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ExpireTopicMessagesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExpireTopicMessagesNPRequest struct via the builder pattern


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


## GetBacklogNP

> PersistentOfflineTopicStats GetBacklogNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetBacklogNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetBacklogNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBacklogNP`: PersistentOfflineTopicStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetBacklogNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBacklogNPRequest struct via the builder pattern


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


## GetBacklogQuotaMapNP

> map[string]map[string]interface{} GetBacklogQuotaMapNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetBacklogQuotaMapNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetBacklogQuotaMapNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBacklogQuotaMapNP`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetBacklogQuotaMapNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBacklogQuotaMapNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompactionThresholdNP

> GetCompactionThresholdNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetCompactionThresholdNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetCompactionThresholdNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetCompactionThresholdNPRequest struct via the builder pattern


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


## GetDeduplicationEnabledNP

> GetDeduplicationEnabledNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDeduplicationEnabledNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDeduplicationEnabledNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDeduplicationEnabledNPRequest struct via the builder pattern


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


## GetDeduplicationSnapshotIntervalNP

> GetDeduplicationSnapshotIntervalNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDeduplicationSnapshotIntervalNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDeduplicationSnapshotIntervalNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDeduplicationSnapshotIntervalNPRequest struct via the builder pattern


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


## GetDelayedDeliveryPoliciesNP

> GetDelayedDeliveryPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDelayedDeliveryPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDelayedDeliveryPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDelayedDeliveryPoliciesNPRequest struct via the builder pattern


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


## GetDispatchRateNP

> GetDispatchRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetDispatchRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDispatchRateNPRequest struct via the builder pattern


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


## GetInactiveTopicPoliciesNP

> GetInactiveTopicPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetInactiveTopicPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetInactiveTopicPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetInactiveTopicPoliciesNPRequest struct via the builder pattern


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


## GetInternalStatsNP

> PersistentTopicInternalStats GetInternalStatsNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Metadata(metadata).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetInternalStatsNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetInternalStatsNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInternalStatsNP`: PersistentTopicInternalStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetInternalStatsNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetInternalStatsNPRequest struct via the builder pattern


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


## GetLastMessageIdNP

> GetLastMessageIdNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetLastMessageIdNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetLastMessageIdNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLastMessageIdNPRequest struct via the builder pattern


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


## GetListNP

> []string GetListNP(ctx, tenant, namespace).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetListNP(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetListNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListNP`: []string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetListNP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListNPRequest struct via the builder pattern


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


## GetManagedLedgerInfoNP

> GetManagedLedgerInfoNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetManagedLedgerInfoNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetManagedLedgerInfoNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetManagedLedgerInfoNPRequest struct via the builder pattern


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


## GetMaxConsumersNP

> GetMaxConsumersNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxConsumersNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxConsumersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMaxConsumersNPRequest struct via the builder pattern


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


## GetMaxConsumersPerSubscriptionNP

> GetMaxConsumersPerSubscriptionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxConsumersPerSubscriptionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxConsumersPerSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMaxConsumersPerSubscriptionNPRequest struct via the builder pattern


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


## GetMaxProducersNP

> GetMaxProducersNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxProducersNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxProducersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMaxProducersNPRequest struct via the builder pattern


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


## GetMaxUnackedMessagesOnConsumerNP

> GetMaxUnackedMessagesOnConsumerNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxUnackedMessagesOnConsumerNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxUnackedMessagesOnConsumerNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMaxUnackedMessagesOnConsumerNPRequest struct via the builder pattern


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


## GetMaxUnackedMessagesOnSubscriptionNP

> GetMaxUnackedMessagesOnSubscriptionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMaxUnackedMessagesOnSubscriptionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMaxUnackedMessagesOnSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMaxUnackedMessagesOnSubscriptionNPRequest struct via the builder pattern


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


## GetMessageByIdNP

> GetMessageByIdNP(ctx, tenant, namespace, topic, ledgerId, entryId).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetMessageByIdNP(context.Background(), tenant, namespace, topic, ledgerId, entryId).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMessageByIdNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMessageByIdNPRequest struct via the builder pattern


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


## GetMessageTTLNP

> int32 GetMessageTTLNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetMessageTTLNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetMessageTTLNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageTTLNP`: int32
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetMessageTTLNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMessageTTLNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadPoliciesNP

> GetOffloadPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetOffloadPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetOffloadPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOffloadPoliciesNPRequest struct via the builder pattern


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


## GetPartitionedMetadataNP

> PartitionedTopicMetadata GetPartitionedMetadataNP(ctx, tenant, namespace, topic).Authoritative(authoritative).CheckAllowAutoCreation(checkAllowAutoCreation).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedMetadataNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).CheckAllowAutoCreation(checkAllowAutoCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedMetadataNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedMetadataNP`: PartitionedTopicMetadata
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPartitionedMetadataNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartitionedMetadataNPRequest struct via the builder pattern


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


## GetPartitionedStatsNP

> GetPartitionedStatsNP(ctx, tenant, namespace, topic).PerPartition(perPartition).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).Execute()

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
    getPreciseBacklog := true // bool | Is return precise backlog or imprecise backlog (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedStatsNP(context.Background(), tenant, namespace, topic).PerPartition(perPartition).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedStatsNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetPartitionedStatsNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPartition** | **bool** | Get per partition stats | [default to true]
 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **getPreciseBacklog** | **bool** | Is return precise backlog or imprecise backlog | [default to false]

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


## GetPartitionedTopicListNP

> []string GetPartitionedTopicListNP(ctx, tenant, namespace).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetPartitionedTopicListNP(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPartitionedTopicListNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartitionedTopicListNP`: []string
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPartitionedTopicListNP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Specify the tenant | 
**namespace** | **string** | Specify the namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionedTopicListNPRequest struct via the builder pattern


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


## GetPermissionsOnTopicNP

> map[string]map[string]interface{} GetPermissionsOnTopicNP(ctx, tenant, namespace, topic).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetPermissionsOnTopicNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPermissionsOnTopicNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionsOnTopicNP`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetPermissionsOnTopicNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPermissionsOnTopicNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersistenceNP

> GetPersistenceNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPersistenceNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPersistenceNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetPersistenceNPRequest struct via the builder pattern


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


## GetPublishRateNP

> GetPublishRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetPublishRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetPublishRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetPublishRateNPRequest struct via the builder pattern


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


## GetRetentionNP

> GetRetentionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetRetentionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetRetentionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetRetentionNPRequest struct via the builder pattern


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


## GetStatsNP

> TopicStats GetStatsNP(ctx, tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).Execute()

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
    getPreciseBacklog := true // bool | Is return precise backlog or imprecise backlog (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetStatsNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).GetPreciseBacklog(getPreciseBacklog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetStatsNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsNP`: TopicStats
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.GetStatsNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetStatsNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** | Is authentication required to perform this operation | [default to false]
 **getPreciseBacklog** | **bool** | Is return precise backlog or imprecise backlog | [default to false]

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


## GetSubscribeRateNP

> GetSubscribeRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscribeRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscribeRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSubscribeRateNPRequest struct via the builder pattern


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


## GetSubscriptionDispatchRateNP

> GetSubscriptionDispatchRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscriptionDispatchRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscriptionDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSubscriptionDispatchRateNPRequest struct via the builder pattern


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


## GetSubscriptionsNP

> GetSubscriptionsNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GetSubscriptionsNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GetSubscriptionsNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSubscriptionsNPRequest struct via the builder pattern


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


## GrantPermissionsOnTopicNP

> GrantPermissionsOnTopicNP(ctx, tenant, namespace, topic, role).Body(body).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.GrantPermissionsOnTopicNP(context.Background(), tenant, namespace, topic, role).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.GrantPermissionsOnTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGrantPermissionsOnTopicNPRequest struct via the builder pattern


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


## OffloadStatusNP

> OffloadProcessStatus OffloadStatusNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.OffloadStatusNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.OffloadStatusNP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OffloadStatusNP`: OffloadProcessStatus
    fmt.Fprintf(os.Stdout, "Response from `NonPersistentTopicApi.OffloadStatusNP`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOffloadStatusNPRequest struct via the builder pattern


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


## PeekNthMessageNP

> PeekNthMessageNP(ctx, tenant, namespace, topic, subName, messagePosition).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.PeekNthMessageNP(context.Background(), tenant, namespace, topic, subName, messagePosition).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.PeekNthMessageNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPeekNthMessageNPRequest struct via the builder pattern


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


## RemoveBacklogQuotaNP

> RemoveBacklogQuotaNP(ctx, tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveBacklogQuotaNP(context.Background(), tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveBacklogQuotaNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveBacklogQuotaNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## RemoveCompactionThresholdNP

> RemoveCompactionThresholdNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveCompactionThresholdNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveCompactionThresholdNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCompactionThresholdNPRequest struct via the builder pattern


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


## RemoveDeduplicationEnabledNP

> RemoveDeduplicationEnabledNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveDeduplicationEnabledNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveDeduplicationEnabledNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveDeduplicationEnabledNPRequest struct via the builder pattern


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


## RemoveDispatchRateNP

> RemoveDispatchRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveDispatchRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveDispatchRateNPRequest struct via the builder pattern


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


## RemoveMaxConsumersNP

> RemoveMaxConsumersNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxConsumersNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxConsumersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveMaxConsumersNPRequest struct via the builder pattern


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


## RemoveMaxConsumersPerSubscriptionNP

> RemoveMaxConsumersPerSubscriptionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxConsumersPerSubscriptionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxConsumersPerSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveMaxConsumersPerSubscriptionNPRequest struct via the builder pattern


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


## RemoveMaxProducersNP

> RemoveMaxProducersNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMaxProducersNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMaxProducersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveMaxProducersNPRequest struct via the builder pattern


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


## RemoveMessageTTLNP

> RemoveMessageTTLNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveMessageTTLNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveMessageTTLNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveMessageTTLNPRequest struct via the builder pattern


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


## RemoveOffloadPoliciesNP

> RemoveOffloadPoliciesNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveOffloadPoliciesNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveOffloadPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveOffloadPoliciesNPRequest struct via the builder pattern


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


## RemovePersistenceNP

> RemovePersistenceNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemovePersistenceNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemovePersistenceNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemovePersistenceNPRequest struct via the builder pattern


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


## RemovePublishRateNP

> RemovePublishRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemovePublishRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemovePublishRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemovePublishRateNPRequest struct via the builder pattern


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


## RemoveRetentionNP

> RemoveRetentionNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveRetentionNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveRetentionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveRetentionNPRequest struct via the builder pattern


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


## RemoveSubscribeRateNP

> RemoveSubscribeRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveSubscribeRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveSubscribeRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveSubscribeRateNPRequest struct via the builder pattern


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


## RemoveSubscriptionDispatchRateNP

> RemoveSubscriptionDispatchRateNP(ctx, tenant, namespace, topic).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.RemoveSubscriptionDispatchRateNP(context.Background(), tenant, namespace, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RemoveSubscriptionDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveSubscriptionDispatchRateNPRequest struct via the builder pattern


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


## ResetCursorNP

> ResetCursorNP(ctx, tenant, namespace, topic, subName, timestamp).Authoritative(authoritative).Execute()

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
    timestamp := int64(789) // int64 | time in minutes to reset back to (or minutes, hours,days,weeks eg:100m, 3h, 2d, 5w)
    authoritative := true // bool | Is authentication required to perform this operation (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.ResetCursorNP(context.Background(), tenant, namespace, topic, subName, timestamp).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ResetCursorNP``: %v\n", err)
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
**timestamp** | **int64** | time in minutes to reset back to (or minutes, hours,days,weeks eg:100m, 3h, 2d, 5w) | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCursorNPRequest struct via the builder pattern


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


## ResetCursorOnPositionNP

> ResetCursorOnPositionNP(ctx, tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.ResetCursorOnPositionNP(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).MessageId(messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.ResetCursorOnPositionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetCursorOnPositionNPRequest struct via the builder pattern


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


## RevokePermissionsOnTopicNP

> RevokePermissionsOnTopicNP(ctx, tenant, namespace, topic, role).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.RevokePermissionsOnTopicNP(context.Background(), tenant, namespace, topic, role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.RevokePermissionsOnTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokePermissionsOnTopicNPRequest struct via the builder pattern


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


## SetBacklogQuotaNP

> SetBacklogQuotaNP(ctx, tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Execute()

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
    backlogQuotaType := "backlogQuotaType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetBacklogQuotaNP(context.Background(), tenant, namespace, topic).BacklogQuotaType(backlogQuotaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetBacklogQuotaNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetBacklogQuotaNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetCompactionThresholdNP

> SetCompactionThresholdNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int64(789) // int64 | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetCompactionThresholdNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetCompactionThresholdNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetCompactionThresholdNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetDeduplicationEnabledNP

> SetDeduplicationEnabledNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := true // bool | DeduplicationEnabled policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDeduplicationEnabledNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDeduplicationEnabledNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDeduplicationEnabledNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetDeduplicationSnapshotIntervalNP

> SetDeduplicationSnapshotIntervalNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | Interval to take deduplication snapshot for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDeduplicationSnapshotIntervalNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDeduplicationSnapshotIntervalNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDeduplicationSnapshotIntervalNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetDelayedDeliveryPoliciesNP

> SetDelayedDeliveryPoliciesNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewDelayedDeliveryPolicies() // DelayedDeliveryPolicies | Delayed delivery policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDelayedDeliveryPoliciesNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDelayedDeliveryPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDelayedDeliveryPoliciesNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetDispatchRateNP

> SetDispatchRateNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewDispatchRate() // DispatchRate | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetDispatchRateNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDispatchRateNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**DispatchRate**](DispatchRate.md) | Dispatch rate for the specified topic | 

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


## SetInactiveTopicPoliciesNP

> SetInactiveTopicPoliciesNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewInactiveTopicPolicies() // InactiveTopicPolicies | inactive topic policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetInactiveTopicPoliciesNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetInactiveTopicPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetInactiveTopicPoliciesNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMaxConsumersNP

> SetMaxConsumersNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | The max consumers of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxConsumersNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxConsumersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMaxConsumersNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMaxConsumersPerSubscriptionNP

> SetMaxConsumersPerSubscriptionNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxConsumersPerSubscriptionNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxConsumersPerSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMaxConsumersPerSubscriptionNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMaxProducersNP

> SetMaxProducersNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | The max producers of the topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxProducersNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxProducersNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMaxProducersNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMaxUnackedMessagesOnConsumerNP

> SetMaxUnackedMessagesOnConsumerNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | Max unacked messages on consumer policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxUnackedMessagesOnConsumerNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxUnackedMessagesOnConsumerNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesOnConsumerNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMaxUnackedMessagesOnSubscriptionNP

> SetMaxUnackedMessagesOnSubscriptionNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := int32(56) // int32 | Max unacked messages on subscription policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMaxUnackedMessagesOnSubscriptionNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMaxUnackedMessagesOnSubscriptionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesOnSubscriptionNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetMessageTTLNP

> SetMessageTTLNP(ctx, tenant, namespace, topic).MessageTTL(messageTTL).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetMessageTTLNP(context.Background(), tenant, namespace, topic).MessageTTL(messageTTL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetMessageTTLNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMessageTTLNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **messageTTL** | **int32** | TTL in seconds for the specified namespace | 

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


## SetOffloadPoliciesNP

> SetOffloadPoliciesNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewOffloadPolicies() // OffloadPolicies | Offload policies for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetOffloadPoliciesNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetOffloadPoliciesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetOffloadPoliciesNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**OffloadPolicies**](OffloadPolicies.md) | Offload policies for the specified topic | 

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


## SetPersistenceNP

> SetPersistenceNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewPersistencePolicies() // PersistencePolicies | Bookkeeper persistence policies for specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetPersistenceNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetPersistenceNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetPersistenceNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetPublishRateNP

> SetPublishRateNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewPublishRate() // PublishRate | Dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetPublishRateNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetPublishRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetPublishRateNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetRetentionNP

> SetRetentionNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewRetentionPolicies() // RetentionPolicies | Retention policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetRetentionNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetRetentionNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetRetentionNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetSubscribeRateNP

> SetSubscribeRateNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewSubscribeRate() // SubscribeRate | Subscribe rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetSubscribeRateNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetSubscribeRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetSubscribeRateNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SetSubscriptionDispatchRateNP

> SetSubscriptionDispatchRateNP(ctx, tenant, namespace, topic).Body(body).Execute()

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
    body := *openapiclient.NewDispatchRate() // DispatchRate | Subscription message dispatch rate for the specified topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NonPersistentTopicApi.SetSubscriptionDispatchRateNP(context.Background(), tenant, namespace, topic).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SetSubscriptionDispatchRateNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetSubscriptionDispatchRateNPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**DispatchRate**](DispatchRate.md) | Subscription message dispatch rate for the specified topic | 

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


## SkipAllMessagesNP

> SkipAllMessagesNP(ctx, tenant, namespace, topic, subName).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.SkipAllMessagesNP(context.Background(), tenant, namespace, topic, subName).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SkipAllMessagesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSkipAllMessagesNPRequest struct via the builder pattern


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


## SkipMessagesNP

> SkipMessagesNP(ctx, tenant, namespace, topic, subName, numMessages).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.SkipMessagesNP(context.Background(), tenant, namespace, topic, subName, numMessages).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.SkipMessagesNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSkipMessagesNPRequest struct via the builder pattern


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


## TriggerOffloadNP

> TriggerOffloadNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.TriggerOffloadNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.TriggerOffloadNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTriggerOffloadNPRequest struct via the builder pattern


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


## UnloadTopicNP

> UnloadTopicNP(ctx, tenant, namespace, topic).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.UnloadTopicNP(context.Background(), tenant, namespace, topic).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.UnloadTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnloadTopicNPRequest struct via the builder pattern


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


## UpdatePartitionedTopicNP

> UpdatePartitionedTopicNP(ctx, tenant, namespace, topic).Body(body).UpdateLocalTopicOnly(updateLocalTopicOnly).Authoritative(authoritative).Execute()

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
    resp, r, err := api_client.NonPersistentTopicApi.UpdatePartitionedTopicNP(context.Background(), tenant, namespace, topic).Body(body).UpdateLocalTopicOnly(updateLocalTopicOnly).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonPersistentTopicApi.UpdatePartitionedTopicNP``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePartitionedTopicNPRequest struct via the builder pattern


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

