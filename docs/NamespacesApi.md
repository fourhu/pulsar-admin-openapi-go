# \NamespacesApi

All URIs are relative to *http://localhost/admin/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearNamespaceBacklog**](NamespacesApi.md#ClearNamespaceBacklog) | **Post** /namespaces/{tenant}/{namespace}/clearBacklog | Clear backlog for all topics on a namespace.
[**ClearNamespaceBacklogForSubscription**](NamespacesApi.md#ClearNamespaceBacklogForSubscription) | **Post** /namespaces/{tenant}/{namespace}/clearBacklog/{subscription} | Clear backlog for a given subscription on all topics on a namespace.
[**ClearNamespaceBundleBacklog**](NamespacesApi.md#ClearNamespaceBundleBacklog) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/clearBacklog | Clear backlog for all topics on a namespace bundle.
[**ClearNamespaceBundleBacklogForSubscription**](NamespacesApi.md#ClearNamespaceBundleBacklogForSubscription) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/clearBacklog/{subscription} | Clear backlog for a given subscription on all topics on a namespace bundle.
[**ClearOffloadDeletionLag**](NamespacesApi.md#ClearOffloadDeletionLag) | **Delete** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Clear the namespace configured offload deletion lag. The topics in the namespace will fallback to using the default configured deletion lag for the broker
[**ClearProperties**](NamespacesApi.md#ClearProperties) | **Delete** /namespaces/{tenant}/{namespace}/properties | Get property value for a given key on a namespace.
[**CreateNamespace**](NamespacesApi.md#CreateNamespace) | **Put** /namespaces/{tenant}/{namespace} | Creates a new namespace with the specified policies
[**DeleteBookieAffinityGroup**](NamespacesApi.md#DeleteBookieAffinityGroup) | **Delete** /namespaces/{property}/{namespace}/persistence/bookieAffinity | Delete the bookie-affinity-group from namespace-local policy.
[**DeleteCompactionThreshold**](NamespacesApi.md#DeleteCompactionThreshold) | **Delete** /namespaces/{tenant}/{namespace}/compactionThreshold | Delete maximum number of uncompacted bytes in a topic before compaction is triggered.
[**DeleteDispatchRate**](NamespacesApi.md#DeleteDispatchRate) | **Delete** /namespaces/{tenant}/{namespace}/dispatchRate | Delete dispatch-rate throttling for all topics of the namespace
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /namespaces/{tenant}/{namespace} | Delete a namespace and all the topics under it.
[**DeleteNamespaceBundle**](NamespacesApi.md#DeleteNamespaceBundle) | **Delete** /namespaces/{tenant}/{namespace}/{bundle} | Delete a namespace bundle and all the topics under it.
[**DeletePersistence**](NamespacesApi.md#DeletePersistence) | **Delete** /namespaces/{tenant}/{namespace}/persistence | Delete the persistence configuration for all topics on a namespace
[**DeleteSubscribeRate**](NamespacesApi.md#DeleteSubscribeRate) | **Delete** /namespaces/{tenant}/{namespace}/subscribeRate | Delete subscribe-rate throttling for all topics of the namespace
[**DeleteSubscriptionDispatchRate**](NamespacesApi.md#DeleteSubscriptionDispatchRate) | **Delete** /namespaces/{tenant}/{namespace}/subscriptionDispatchRate | Delete Subscription dispatch-rate throttling for all topics of the namespace
[**GetAntiAffinityNamespaces**](NamespacesApi.md#GetAntiAffinityNamespaces) | **Get** /namespaces/{cluster}/antiAffinity/{group} | Get all namespaces that are grouped by given anti-affinity group in a given cluster. api can be only accessed by admin of any of the existing tenant
[**GetBacklogQuotaMap**](NamespacesApi.md#GetBacklogQuotaMap) | **Get** /namespaces/{tenant}/{namespace}/backlogQuotaMap | Get backlog quota map on a namespace.
[**GetBookieAffinityGroup**](NamespacesApi.md#GetBookieAffinityGroup) | **Get** /namespaces/{property}/{namespace}/persistence/bookieAffinity | Get the bookie-affinity-group from namespace-local policy.
[**GetBundlesData**](NamespacesApi.md#GetBundlesData) | **Get** /namespaces/{tenant}/{namespace}/bundles | Get the bundles split data.
[**GetCompactionThreshold**](NamespacesApi.md#GetCompactionThreshold) | **Get** /namespaces/{tenant}/{namespace}/compactionThreshold | Maximum number of uncompacted bytes in topics before compaction is triggered.
[**GetDeduplication**](NamespacesApi.md#GetDeduplication) | **Get** /namespaces/{tenant}/{namespace}/deduplication | Get broker side deduplication for all topics in a namespace
[**GetDeduplicationSnapshotInterval**](NamespacesApi.md#GetDeduplicationSnapshotInterval) | **Get** /namespaces/{tenant}/{namespace}/deduplicationSnapshotInterval | Get deduplicationSnapshotInterval config on a namespace.
[**GetDelayedDeliveryPolicies**](NamespacesApi.md#GetDelayedDeliveryPolicies) | **Get** /namespaces/{tenant}/{namespace}/delayedDelivery | Get delayed delivery messages config on a namespace.
[**GetDispatchRate**](NamespacesApi.md#GetDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/dispatchRate | Get dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetInactiveTopicPolicies**](NamespacesApi.md#GetInactiveTopicPolicies) | **Get** /namespaces/{tenant}/{namespace}/inactiveTopicPolicies | Get inactive topic policies config on a namespace.
[**GetIsAllowAutoUpdateSchema**](NamespacesApi.md#GetIsAllowAutoUpdateSchema) | **Get** /namespaces/{tenant}/{namespace}/isAllowAutoUpdateSchema | The flag of whether allow auto update schema
[**GetMaxConsumersPerSubscription**](NamespacesApi.md#GetMaxConsumersPerSubscription) | **Get** /namespaces/{tenant}/{namespace}/maxConsumersPerSubscription | Get maxConsumersPerSubscription config on a namespace.
[**GetMaxConsumersPerTopic**](NamespacesApi.md#GetMaxConsumersPerTopic) | **Get** /namespaces/{tenant}/{namespace}/maxConsumersPerTopic | Get maxConsumersPerTopic config on a namespace.
[**GetMaxProducersPerTopic**](NamespacesApi.md#GetMaxProducersPerTopic) | **Get** /namespaces/{tenant}/{namespace}/maxProducersPerTopic | Get maxProducersPerTopic config on a namespace.
[**GetMaxSubscriptionsPerTopic**](NamespacesApi.md#GetMaxSubscriptionsPerTopic) | **Get** /namespaces/{tenant}/{namespace}/maxSubscriptionsPerTopic | Get maxSubscriptionsPerTopic config on a namespace.
[**GetMaxTopicsPerNamespace**](NamespacesApi.md#GetMaxTopicsPerNamespace) | **Get** /namespaces/{tenant}/{namespace}/maxTopicsPerNamespace | Get maxTopicsPerNamespace config on a namespace.
[**GetMaxUnackedMessagesPerConsumer**](NamespacesApi.md#GetMaxUnackedMessagesPerConsumer) | **Get** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerConsumer | Get maxUnackedMessagesPerConsumer config on a namespace.
[**GetMaxUnackedmessagesPerSubscription**](NamespacesApi.md#GetMaxUnackedmessagesPerSubscription) | **Get** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerSubscription | Get maxUnackedMessagesPerSubscription config on a namespace.
[**GetNamespaceAntiAffinityGroup**](NamespacesApi.md#GetNamespaceAntiAffinityGroup) | **Get** /namespaces/{tenant}/{namespace}/antiAffinity | Get anti-affinity group of a namespace.
[**GetNamespaceMessageTTL**](NamespacesApi.md#GetNamespaceMessageTTL) | **Get** /namespaces/{tenant}/{namespace}/messageTTL | Get the message TTL for the namespace
[**GetNamespaceReplicationClusters**](NamespacesApi.md#GetNamespaceReplicationClusters) | **Get** /namespaces/{tenant}/{namespace}/replication | Get the replication clusters for a namespace.
[**GetNamespaceResourceGroup**](NamespacesApi.md#GetNamespaceResourceGroup) | **Get** /namespaces/{tenant}/{namespace}/resourcegroup | Get the resourcegroup attached to the namespace
[**GetOffloadDeletionLag**](NamespacesApi.md#GetOffloadDeletionLag) | **Get** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster&#39;s local storage (i.e. BookKeeper)
[**GetOffloadPolicies**](NamespacesApi.md#GetOffloadPolicies) | **Get** /namespaces/{tenant}/{namespace}/offloadPolicies | Get offload configuration on a namespace.
[**GetOffloadThreshold**](NamespacesApi.md#GetOffloadThreshold) | **Get** /namespaces/{tenant}/{namespace}/offloadThreshold | Maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage
[**GetPermissions**](NamespacesApi.md#GetPermissions) | **Get** /namespaces/{tenant}/{namespace}/permissions | Retrieve the permissions for a namespace.
[**GetPersistence**](NamespacesApi.md#GetPersistence) | **Get** /namespaces/{tenant}/{namespace}/persistence | Get the persistence configuration for a namespace.
[**GetPolicies**](NamespacesApi.md#GetPolicies) | **Get** /namespaces/{tenant}/{namespace} | Get the dump all the policies specified for a namespace.
[**GetProperties**](NamespacesApi.md#GetProperties) | **Get** /namespaces/{tenant}/{namespace}/properties | Get key value pair properties for a given namespace.
[**GetProperty**](NamespacesApi.md#GetProperty) | **Get** /namespaces/{tenant}/{namespace}/property/{key} | Get property value for a given key on a namespace.
[**GetReplicatorDispatchRate**](NamespacesApi.md#GetReplicatorDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/replicatorDispatchRate | Get replicator dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetRetention**](NamespacesApi.md#GetRetention) | **Get** /namespaces/{tenant}/{namespace}/retention | Get retention config on a namespace.
[**GetSchemaAutoUpdateCompatibilityStrategy**](NamespacesApi.md#GetSchemaAutoUpdateCompatibilityStrategy) | **Get** /namespaces/{tenant}/{namespace}/schemaAutoUpdateCompatibilityStrategy | The strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema
[**GetSchemaCompatibilityStrategy**](NamespacesApi.md#GetSchemaCompatibilityStrategy) | **Get** /namespaces/{tenant}/{namespace}/schemaCompatibilityStrategy | The strategy of the namespace schema compatibility 
[**GetSchemaValidtionEnforced**](NamespacesApi.md#GetSchemaValidtionEnforced) | **Get** /namespaces/{tenant}/{namespace}/schemaValidationEnforced | Get schema validation enforced flag for namespace.
[**GetSubscribeRate**](NamespacesApi.md#GetSubscribeRate) | **Get** /namespaces/{tenant}/{namespace}/subscribeRate | Get subscribe-rate configured for the namespace
[**GetSubscriptionDispatchRate**](NamespacesApi.md#GetSubscriptionDispatchRate) | **Get** /namespaces/{tenant}/{namespace}/subscriptionDispatchRate | Get Subscription dispatch-rate configured for the namespace, -1 represents not configured yet
[**GetSubscriptionExpirationTime**](NamespacesApi.md#GetSubscriptionExpirationTime) | **Get** /namespaces/{tenant}/{namespace}/subscriptionExpirationTime | Get the subscription expiration time for the namespace
[**GetSubscriptionTypesEnabled**](NamespacesApi.md#GetSubscriptionTypesEnabled) | **Get** /namespaces/{tenant}/{namespace}/subscriptionTypesEnabled | The set of whether allow subscription types
[**GetTenantNamespaces**](NamespacesApi.md#GetTenantNamespaces) | **Get** /namespaces/{tenant} | Get the list of all the namespaces for a certain tenant.
[**GetTopics**](NamespacesApi.md#GetTopics) | **Get** /namespaces/{tenant}/{namespace}/topics | Get the list of all the topics under a certain namespace.
[**GrantPermissionOnNamespace**](NamespacesApi.md#GrantPermissionOnNamespace) | **Post** /namespaces/{tenant}/{namespace}/permissions/{role} | Grant a new permission to a role on a namespace.
[**ModifyDeduplication**](NamespacesApi.md#ModifyDeduplication) | **Post** /namespaces/{tenant}/{namespace}/deduplication | Enable or disable broker side deduplication for all topics in a namespace
[**ModifyEncryptionRequired**](NamespacesApi.md#ModifyEncryptionRequired) | **Post** /namespaces/{tenant}/{namespace}/encryptionRequired | Message encryption is required or not for all topics in a namespace
[**RemoveAutoSubscriptionCreation**](NamespacesApi.md#RemoveAutoSubscriptionCreation) | **Delete** /namespaces/{tenant}/{namespace}/autoSubscriptionCreation | Remove override of broker&#39;s allowAutoSubscriptionCreation in a namespace
[**RemoveAutoTopicCreation**](NamespacesApi.md#RemoveAutoTopicCreation) | **Delete** /namespaces/{tenant}/{namespace}/autoTopicCreation | Remove override of broker&#39;s allowAutoTopicCreation in a namespace
[**RemoveBacklogQuota**](NamespacesApi.md#RemoveBacklogQuota) | **Delete** /namespaces/{tenant}/{namespace}/backlogQuota | Remove a backlog quota policy from a namespace.
[**RemoveDeduplication**](NamespacesApi.md#RemoveDeduplication) | **Delete** /namespaces/{tenant}/{namespace}/deduplication | Remove broker side deduplication for all topics in a namespace
[**RemoveDelayedDeliveryPolicies**](NamespacesApi.md#RemoveDelayedDeliveryPolicies) | **Delete** /namespaces/{tenant}/{namespace}/delayedDelivery | Delete delayed delivery messages config on a namespace.
[**RemoveInactiveTopicPolicies**](NamespacesApi.md#RemoveInactiveTopicPolicies) | **Delete** /namespaces/{tenant}/{namespace}/inactiveTopicPolicies | Remove inactive topic policies from a namespace.
[**RemoveMaxConsumersPerSubscription**](NamespacesApi.md#RemoveMaxConsumersPerSubscription) | **Delete** /namespaces/{tenant}/{namespace}/maxConsumersPerSubscription |  Set maxConsumersPerSubscription configuration on a namespace.
[**RemoveMaxConsumersPerTopic**](NamespacesApi.md#RemoveMaxConsumersPerTopic) | **Delete** /namespaces/{tenant}/{namespace}/maxConsumersPerTopic | Remove maxConsumersPerTopic configuration on a namespace.
[**RemoveMaxProducersPerTopic**](NamespacesApi.md#RemoveMaxProducersPerTopic) | **Delete** /namespaces/{tenant}/{namespace}/maxProducersPerTopic | Remove maxProducersPerTopic configuration on a namespace.
[**RemoveMaxSubscriptionsPerTopic**](NamespacesApi.md#RemoveMaxSubscriptionsPerTopic) | **Delete** /namespaces/{tenant}/{namespace}/maxSubscriptionsPerTopic | Remove maxSubscriptionsPerTopic configuration on a namespace.
[**RemoveMaxUnackedmessagesPerConsumer**](NamespacesApi.md#RemoveMaxUnackedmessagesPerConsumer) | **Delete** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerConsumer | Remove maxUnackedMessagesPerConsumer config on a namespace.
[**RemoveMaxUnackedmessagesPerSubscription**](NamespacesApi.md#RemoveMaxUnackedmessagesPerSubscription) | **Delete** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerSubscription | Remove maxUnackedMessagesPerSubscription config on a namespace.
[**RemoveNamespaceAntiAffinityGroup**](NamespacesApi.md#RemoveNamespaceAntiAffinityGroup) | **Delete** /namespaces/{tenant}/{namespace}/antiAffinity | Remove anti-affinity group of a namespace.
[**RemoveNamespaceMessageTTL**](NamespacesApi.md#RemoveNamespaceMessageTTL) | **Delete** /namespaces/{tenant}/{namespace}/messageTTL | Set message TTL in seconds for namespace
[**RemoveNamespaceResourceGroup**](NamespacesApi.md#RemoveNamespaceResourceGroup) | **Delete** /namespaces/{tenant}/{namespace}/resourcegroup | Delete resourcegroup for a namespace
[**RemoveOffloadPolicies**](NamespacesApi.md#RemoveOffloadPolicies) | **Delete** /namespaces/{tenant}/{namespace}/removeOffloadPolicies |  Set offload configuration on a namespace.
[**RemoveProperty**](NamespacesApi.md#RemoveProperty) | **Delete** /namespaces/{tenant}/{namespace}/property/{key} | Get property value for a given key on a namespace.
[**RemoveReplicatorDispatchRate**](NamespacesApi.md#RemoveReplicatorDispatchRate) | **Delete** /namespaces/{tenant}/{namespace}/replicatorDispatchRate | Remove replicator dispatch-rate throttling for all topics of the namespace
[**RemoveRetention**](NamespacesApi.md#RemoveRetention) | **Delete** /namespaces/{tenant}/{namespace}/retention |  Remove retention configuration on a namespace.
[**RemoveSubscriptionExpirationTime**](NamespacesApi.md#RemoveSubscriptionExpirationTime) | **Delete** /namespaces/{tenant}/{namespace}/subscriptionExpirationTime | Remove subscription expiration time for namespace
[**RevokePermissionsOnNamespace**](NamespacesApi.md#RevokePermissionsOnNamespace) | **Delete** /namespaces/{tenant}/{namespace}/permissions/{role} | Revoke all permissions to a role on a namespace.
[**SetAutoSubscriptionCreation**](NamespacesApi.md#SetAutoSubscriptionCreation) | **Post** /namespaces/{tenant}/{namespace}/autoSubscriptionCreation | Override broker&#39;s allowAutoSubscriptionCreation setting for a namespace
[**SetAutoTopicCreation**](NamespacesApi.md#SetAutoTopicCreation) | **Post** /namespaces/{tenant}/{namespace}/autoTopicCreation | Override broker&#39;s allowAutoTopicCreation setting for a namespace
[**SetBacklogQuota**](NamespacesApi.md#SetBacklogQuota) | **Post** /namespaces/{tenant}/{namespace}/backlogQuota |  Set a backlog quota for all the topics on a namespace.
[**SetBookieAffinityGroup**](NamespacesApi.md#SetBookieAffinityGroup) | **Post** /namespaces/{tenant}/{namespace}/persistence/bookieAffinity | Set the bookie-affinity-group to namespace-persistent policy.
[**SetCompactionThreshold**](NamespacesApi.md#SetCompactionThreshold) | **Put** /namespaces/{tenant}/{namespace}/compactionThreshold | Set maximum number of uncompacted bytes in a topic before compaction is triggered.
[**SetDeduplicationSnapshotInterval**](NamespacesApi.md#SetDeduplicationSnapshotInterval) | **Post** /namespaces/{tenant}/{namespace}/deduplicationSnapshotInterval | Set deduplicationSnapshotInterval config on a namespace.
[**SetDelayedDeliveryPolicies**](NamespacesApi.md#SetDelayedDeliveryPolicies) | **Post** /namespaces/{tenant}/{namespace}/delayedDelivery | Set delayed delivery messages config on a namespace.
[**SetDispatchRate**](NamespacesApi.md#SetDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/dispatchRate | Set dispatch-rate throttling for all topics of the namespace
[**SetInactiveTopicPolicies**](NamespacesApi.md#SetInactiveTopicPolicies) | **Post** /namespaces/{tenant}/{namespace}/inactiveTopicPolicies | Set inactive topic policies config on a namespace.
[**SetInactiveTopicPolicies_0**](NamespacesApi.md#SetInactiveTopicPolicies_0) | **Post** /namespaces/{tenant}/{namespace}/maxTopicsPerNamespace | Set maxTopicsPerNamespace config on a namespace.
[**SetInactiveTopicPolicies_1**](NamespacesApi.md#SetInactiveTopicPolicies_1) | **Delete** /namespaces/{tenant}/{namespace}/maxTopicsPerNamespace | Set maxTopicsPerNamespace config on a namespace.
[**SetIsAllowAutoUpdateSchema**](NamespacesApi.md#SetIsAllowAutoUpdateSchema) | **Post** /namespaces/{tenant}/{namespace}/isAllowAutoUpdateSchema | Update flag of whether allow auto update schema
[**SetMaxConsumersPerSubscription**](NamespacesApi.md#SetMaxConsumersPerSubscription) | **Post** /namespaces/{tenant}/{namespace}/maxConsumersPerSubscription |  Set maxConsumersPerSubscription configuration on a namespace.
[**SetMaxConsumersPerTopic**](NamespacesApi.md#SetMaxConsumersPerTopic) | **Post** /namespaces/{tenant}/{namespace}/maxConsumersPerTopic |  Set maxConsumersPerTopic configuration on a namespace.
[**SetMaxProducersPerTopic**](NamespacesApi.md#SetMaxProducersPerTopic) | **Post** /namespaces/{tenant}/{namespace}/maxProducersPerTopic |  Set maxProducersPerTopic configuration on a namespace.
[**SetMaxSubscriptionsPerTopic**](NamespacesApi.md#SetMaxSubscriptionsPerTopic) | **Post** /namespaces/{tenant}/{namespace}/maxSubscriptionsPerTopic |  Set maxSubscriptionsPerTopic configuration on a namespace.
[**SetMaxUnackedMessagesPerConsumer**](NamespacesApi.md#SetMaxUnackedMessagesPerConsumer) | **Post** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerConsumer |  Set maxConsumersPerTopic configuration on a namespace.
[**SetMaxUnackedMessagesPerSubscription**](NamespacesApi.md#SetMaxUnackedMessagesPerSubscription) | **Post** /namespaces/{tenant}/{namespace}/maxUnackedMessagesPerSubscription |  Set maxUnackedMessagesPerSubscription configuration on a namespace.
[**SetNamespaceAntiAffinityGroup**](NamespacesApi.md#SetNamespaceAntiAffinityGroup) | **Post** /namespaces/{tenant}/{namespace}/antiAffinity | Set anti-affinity group for a namespace
[**SetNamespaceMessageTTL**](NamespacesApi.md#SetNamespaceMessageTTL) | **Post** /namespaces/{tenant}/{namespace}/messageTTL | Set message TTL in seconds for namespace
[**SetNamespaceReplicationClusters**](NamespacesApi.md#SetNamespaceReplicationClusters) | **Post** /namespaces/{tenant}/{namespace}/replication | Set the replication clusters for a namespace.
[**SetNamespaceResourceGroup**](NamespacesApi.md#SetNamespaceResourceGroup) | **Post** /namespaces/{tenant}/{namespace}/resourcegroup | Set resourcegroup for a namespace
[**SetOffloadDeletionLag**](NamespacesApi.md#SetOffloadDeletionLag) | **Put** /namespaces/{tenant}/{namespace}/offloadDeletionLagMs | Set number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster&#39;s local storage (i.e. BookKeeper)
[**SetOffloadPolicies**](NamespacesApi.md#SetOffloadPolicies) | **Post** /namespaces/{tenant}/{namespace}/offloadPolicies |  Set offload configuration on a namespace.
[**SetOffloadThreshold**](NamespacesApi.md#SetOffloadThreshold) | **Put** /namespaces/{tenant}/{namespace}/offloadThreshold | Set maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage
[**SetPersistence**](NamespacesApi.md#SetPersistence) | **Post** /namespaces/{tenant}/{namespace}/persistence | Set the persistence configuration for all the topics on a namespace.
[**SetProperties**](NamespacesApi.md#SetProperties) | **Put** /namespaces/{tenant}/{namespace}/properties | Put key value pairs property on a namespace.
[**SetProperty**](NamespacesApi.md#SetProperty) | **Put** /namespaces/{tenant}/{namespace}/property/{key}/{value} | Put a key value pair property on a namespace.
[**SetReplicatorDispatchRate**](NamespacesApi.md#SetReplicatorDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/replicatorDispatchRate | Set replicator dispatch-rate throttling for all topics of the namespace
[**SetRetention**](NamespacesApi.md#SetRetention) | **Post** /namespaces/{tenant}/{namespace}/retention |  Set retention configuration on a namespace.
[**SetSchemaAutoUpdateCompatibilityStrategy**](NamespacesApi.md#SetSchemaAutoUpdateCompatibilityStrategy) | **Put** /namespaces/{tenant}/{namespace}/schemaAutoUpdateCompatibilityStrategy | Update the strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema
[**SetSchemaCompatibilityStrategy**](NamespacesApi.md#SetSchemaCompatibilityStrategy) | **Put** /namespaces/{tenant}/{namespace}/schemaCompatibilityStrategy | Update the strategy used to check the compatibility of new schema
[**SetSchemaValidtionEnforced**](NamespacesApi.md#SetSchemaValidtionEnforced) | **Post** /namespaces/{tenant}/{namespace}/schemaValidationEnforced | Set schema validation enforced flag on namespace.
[**SetSubscribeRate**](NamespacesApi.md#SetSubscribeRate) | **Post** /namespaces/{tenant}/{namespace}/subscribeRate | Set subscribe-rate throttling for all topics of the namespace
[**SetSubscriptionAuthMode**](NamespacesApi.md#SetSubscriptionAuthMode) | **Post** /namespaces/{tenant}/{namespace}/subscriptionAuthMode |  Set a subscription auth mode for all the topics on a namespace.
[**SetSubscriptionDispatchRate**](NamespacesApi.md#SetSubscriptionDispatchRate) | **Post** /namespaces/{tenant}/{namespace}/subscriptionDispatchRate | Set Subscription dispatch-rate throttling for all topics of the namespace
[**SetSubscriptionExpirationTime**](NamespacesApi.md#SetSubscriptionExpirationTime) | **Post** /namespaces/{tenant}/{namespace}/subscriptionExpirationTime | Set subscription expiration time in minutes for namespace
[**SetSubscriptionTypesEnabled**](NamespacesApi.md#SetSubscriptionTypesEnabled) | **Post** /namespaces/{tenant}/{namespace}/subscriptionTypesEnabled | Update set of whether allow share sub type
[**SplitNamespaceBundle**](NamespacesApi.md#SplitNamespaceBundle) | **Put** /namespaces/{tenant}/{namespace}/{bundle}/split | Split a namespace bundle
[**UnloadNamespace**](NamespacesApi.md#UnloadNamespace) | **Put** /namespaces/{tenant}/{namespace}/unload | Unload namespace
[**UnloadNamespaceBundle**](NamespacesApi.md#UnloadNamespaceBundle) | **Put** /namespaces/{tenant}/{namespace}/{bundle}/unload | Unload a namespace bundle
[**UnsubscribeNamespace**](NamespacesApi.md#UnsubscribeNamespace) | **Post** /namespaces/{tenant}/{namespace}/unsubscribe/{subscription} | Unsubscribes the given subscription on all topics on a namespace.
[**UnsubscribeNamespaceBundle**](NamespacesApi.md#UnsubscribeNamespaceBundle) | **Post** /namespaces/{tenant}/{namespace}/{bundle}/unsubscribe/{subscription} | Unsubscribes the given subscription on all topics on a namespace bundle.



## ClearNamespaceBacklog

> ClearNamespaceBacklog(ctx, tenant, namespace).Authoritative(authoritative).Execute()

Clear backlog for all topics on a namespace.

### Example

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
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearNamespaceBacklog(context.Background(), tenant, namespace).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearNamespaceBacklog``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiClearNamespaceBacklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNamespaceBacklogForSubscription

> ClearNamespaceBacklogForSubscription(ctx, tenant, namespace, subscription).Authoritative(authoritative).Execute()

Clear backlog for a given subscription on all topics on a namespace.

### Example

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
    subscription := "subscription_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearNamespaceBacklogForSubscription(context.Background(), tenant, namespace, subscription).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearNamespaceBacklogForSubscription``: %v\n", err)
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
**subscription** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearNamespaceBacklogForSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNamespaceBundleBacklog

> ClearNamespaceBundleBacklog(ctx, tenant, namespace, bundle).Authoritative(authoritative).Execute()

Clear backlog for all topics on a namespace bundle.

### Example

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
    bundle := "bundle_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearNamespaceBundleBacklog(context.Background(), tenant, namespace, bundle).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearNamespaceBundleBacklog``: %v\n", err)
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
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearNamespaceBundleBacklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNamespaceBundleBacklogForSubscription

> ClearNamespaceBundleBacklogForSubscription(ctx, tenant, namespace, subscription, bundle).Authoritative(authoritative).Execute()

Clear backlog for a given subscription on all topics on a namespace bundle.

### Example

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
    subscription := "subscription_example" // string | 
    bundle := "bundle_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearNamespaceBundleBacklogForSubscription(context.Background(), tenant, namespace, subscription, bundle).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearNamespaceBundleBacklogForSubscription``: %v\n", err)
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
**subscription** | **string** |  | 
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearNamespaceBundleBacklogForSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearOffloadDeletionLag

> ClearOffloadDeletionLag(ctx, tenant, namespace).Execute()

Clear the namespace configured offload deletion lag. The topics in the namespace will fallback to using the default configured deletion lag for the broker

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearOffloadDeletionLag(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearOffloadDeletionLag``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiClearOffloadDeletionLagRequest struct via the builder pattern


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


## ClearProperties

> ClearProperties(ctx, tenant, namespace).Execute()

Get property value for a given key on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ClearProperties(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ClearProperties``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiClearPropertiesRequest struct via the builder pattern


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


## CreateNamespace

> CreateNamespace(ctx, tenant, namespace).Body(body).Execute()

Creates a new namespace with the specified policies

### Example

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
    body := *openapiclient.NewPolicies() // Policies | Policies for the namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.CreateNamespace(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.CreateNamespace``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Policies**](Policies.md) | Policies for the namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBookieAffinityGroup

> DeleteBookieAffinityGroup(ctx, property, namespace).Execute()

Delete the bookie-affinity-group from namespace-local policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    property := "property_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteBookieAffinityGroup(context.Background(), property, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteBookieAffinityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookieAffinityGroupRequest struct via the builder pattern


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


## DeleteCompactionThreshold

> DeleteCompactionThreshold(ctx, tenant, namespace).Execute()

Delete maximum number of uncompacted bytes in a topic before compaction is triggered.



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteCompactionThreshold(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteCompactionThreshold``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompactionThresholdRequest struct via the builder pattern


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


## DeleteDispatchRate

> DeleteDispatchRate(ctx, tenant, namespace).Execute()

Delete dispatch-rate throttling for all topics of the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDispatchRateRequest struct via the builder pattern


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


## DeleteNamespace

> DeleteNamespace(ctx, tenant, namespace).Force(force).Authoritative(authoritative).Execute()

Delete a namespace and all the topics under it.

### Example

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
    force := true // bool |  (optional) (default to false)
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteNamespace(context.Background(), tenant, namespace).Force(force).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteNamespace``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** |  | [default to false]
 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespaceBundle

> DeleteNamespaceBundle(ctx, tenant, namespace, bundle).Force(force).Authoritative(authoritative).Execute()

Delete a namespace bundle and all the topics under it.

### Example

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
    bundle := "bundle_example" // string | 
    force := true // bool |  (optional) (default to false)
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteNamespaceBundle(context.Background(), tenant, namespace, bundle).Force(force).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteNamespaceBundle``: %v\n", err)
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
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** |  | [default to false]
 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersistence

> DeletePersistence(ctx, tenant, namespace).Execute()

Delete the persistence configuration for all topics on a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeletePersistence(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeletePersistence``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersistenceRequest struct via the builder pattern


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


## DeleteSubscribeRate

> DeleteSubscribeRate(ctx, tenant, namespace).Execute()

Delete subscribe-rate throttling for all topics of the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteSubscribeRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteSubscribeRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscribeRateRequest struct via the builder pattern


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


## DeleteSubscriptionDispatchRate

> DeleteSubscriptionDispatchRate(ctx, tenant, namespace).Execute()

Delete Subscription dispatch-rate throttling for all topics of the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.DeleteSubscriptionDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteSubscriptionDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionDispatchRateRequest struct via the builder pattern


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


## GetAntiAffinityNamespaces

> []string GetAntiAffinityNamespaces(ctx, cluster, group).Tenant(tenant).Execute()

Get all namespaces that are grouped by given anti-affinity group in a given cluster. api can be only accessed by admin of any of the existing tenant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cluster := "cluster_example" // string | 
    group := "group_example" // string | 
    tenant := "tenant_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetAntiAffinityNamespaces(context.Background(), cluster, group).Tenant(tenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetAntiAffinityNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAntiAffinityNamespaces`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetAntiAffinityNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cluster** | **string** |  | 
**group** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntiAffinityNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenant** | **string** |  | 

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


## GetBacklogQuotaMap

> map[string]BacklogQuota GetBacklogQuotaMap(ctx, tenant, namespace).Execute()

Get backlog quota map on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetBacklogQuotaMap(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetBacklogQuotaMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBacklogQuotaMap`: map[string]BacklogQuota
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetBacklogQuotaMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBacklogQuotaMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]BacklogQuota**](BacklogQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookieAffinityGroup

> BookieAffinityGroupData GetBookieAffinityGroup(ctx, property, namespace).Execute()

Get the bookie-affinity-group from namespace-local policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    property := "property_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetBookieAffinityGroup(context.Background(), property, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetBookieAffinityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookieAffinityGroup`: BookieAffinityGroupData
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetBookieAffinityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookieAffinityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BookieAffinityGroupData**](BookieAffinityGroupData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundlesData

> BundlesData GetBundlesData(ctx, tenant, namespace).Execute()

Get the bundles split data.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetBundlesData(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetBundlesData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundlesData`: BundlesData
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetBundlesData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundlesDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BundlesData**](BundlesData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompactionThreshold

> int64 GetCompactionThreshold(ctx, tenant, namespace).Execute()

Maximum number of uncompacted bytes in topics before compaction is triggered.



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetCompactionThreshold(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetCompactionThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompactionThreshold`: int64
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetCompactionThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompactionThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeduplication

> bool GetDeduplication(ctx, tenant, namespace).Execute()

Get broker side deduplication for all topics in a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetDeduplication(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetDeduplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeduplication`: bool
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetDeduplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeduplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeduplicationSnapshotInterval

> int32 GetDeduplicationSnapshotInterval(ctx, tenant, namespace).Execute()

Get deduplicationSnapshotInterval config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetDeduplicationSnapshotInterval(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetDeduplicationSnapshotInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeduplicationSnapshotInterval`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetDeduplicationSnapshotInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeduplicationSnapshotIntervalRequest struct via the builder pattern


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


## GetDelayedDeliveryPolicies

> DelayedDeliveryPolicies GetDelayedDeliveryPolicies(ctx, tenant, namespace).Execute()

Get delayed delivery messages config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetDelayedDeliveryPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetDelayedDeliveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelayedDeliveryPolicies`: DelayedDeliveryPolicies
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetDelayedDeliveryPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelayedDeliveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DelayedDeliveryPolicies**](DelayedDeliveryPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispatchRate

> DispatchRate GetDispatchRate(ctx, tenant, namespace).Execute()

Get dispatch-rate configured for the namespace, -1 represents not configured yet

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDispatchRate`: DispatchRate
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetDispatchRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInactiveTopicPolicies

> InactiveTopicPolicies GetInactiveTopicPolicies(ctx, tenant, namespace).Execute()

Get inactive topic policies config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetInactiveTopicPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetInactiveTopicPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInactiveTopicPolicies`: InactiveTopicPolicies
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetInactiveTopicPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInactiveTopicPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InactiveTopicPolicies**](InactiveTopicPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsAllowAutoUpdateSchema

> bool GetIsAllowAutoUpdateSchema(ctx, tenant, namespace).Execute()

The flag of whether allow auto update schema

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetIsAllowAutoUpdateSchema(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetIsAllowAutoUpdateSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIsAllowAutoUpdateSchema`: bool
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetIsAllowAutoUpdateSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsAllowAutoUpdateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxConsumersPerSubscription

> int32 GetMaxConsumersPerSubscription(ctx, tenant, namespace).Execute()

Get maxConsumersPerSubscription config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxConsumersPerSubscription(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxConsumersPerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxConsumersPerSubscription`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxConsumersPerSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxConsumersPerSubscriptionRequest struct via the builder pattern


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


## GetMaxConsumersPerTopic

> int32 GetMaxConsumersPerTopic(ctx, tenant, namespace).Execute()

Get maxConsumersPerTopic config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxConsumersPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxConsumersPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxConsumersPerTopic`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxConsumersPerTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxConsumersPerTopicRequest struct via the builder pattern


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


## GetMaxProducersPerTopic

> int32 GetMaxProducersPerTopic(ctx, tenant, namespace).Execute()

Get maxProducersPerTopic config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxProducersPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxProducersPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxProducersPerTopic`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxProducersPerTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxProducersPerTopicRequest struct via the builder pattern


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


## GetMaxSubscriptionsPerTopic

> int32 GetMaxSubscriptionsPerTopic(ctx, tenant, namespace).Execute()

Get maxSubscriptionsPerTopic config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxSubscriptionsPerTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxSubscriptionsPerTopic`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxSubscriptionsPerTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxSubscriptionsPerTopicRequest struct via the builder pattern


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


## GetMaxTopicsPerNamespace

> int32 GetMaxTopicsPerNamespace(ctx, tenant, namespace).Execute()

Get maxTopicsPerNamespace config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxTopicsPerNamespace(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxTopicsPerNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxTopicsPerNamespace`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxTopicsPerNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxTopicsPerNamespaceRequest struct via the builder pattern


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


## GetMaxUnackedMessagesPerConsumer

> int32 GetMaxUnackedMessagesPerConsumer(ctx, tenant, namespace).Execute()

Get maxUnackedMessagesPerConsumer config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxUnackedMessagesPerConsumer(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxUnackedMessagesPerConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxUnackedMessagesPerConsumer`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxUnackedMessagesPerConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxUnackedMessagesPerConsumerRequest struct via the builder pattern


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


## GetMaxUnackedmessagesPerSubscription

> int32 GetMaxUnackedmessagesPerSubscription(ctx, tenant, namespace).Execute()

Get maxUnackedMessagesPerSubscription config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetMaxUnackedmessagesPerSubscription(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetMaxUnackedmessagesPerSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxUnackedmessagesPerSubscription`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetMaxUnackedmessagesPerSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxUnackedmessagesPerSubscriptionRequest struct via the builder pattern


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


## GetNamespaceAntiAffinityGroup

> string GetNamespaceAntiAffinityGroup(ctx, tenant, namespace).Execute()

Get anti-affinity group of a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetNamespaceAntiAffinityGroup(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespaceAntiAffinityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceAntiAffinityGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespaceAntiAffinityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceAntiAffinityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaceMessageTTL

> int32 GetNamespaceMessageTTL(ctx, tenant, namespace).Execute()

Get the message TTL for the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetNamespaceMessageTTL(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespaceMessageTTL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceMessageTTL`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespaceMessageTTL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceMessageTTLRequest struct via the builder pattern


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


## GetNamespaceReplicationClusters

> []string GetNamespaceReplicationClusters(ctx, tenant, namespace).Execute()

Get the replication clusters for a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetNamespaceReplicationClusters(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespaceReplicationClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceReplicationClusters`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespaceReplicationClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceReplicationClustersRequest struct via the builder pattern


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


## GetNamespaceResourceGroup

> string GetNamespaceResourceGroup(ctx, tenant, namespace).Execute()

Get the resourcegroup attached to the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetNamespaceResourceGroup(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespaceResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaceResourceGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespaceResourceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadDeletionLag

> int64 GetOffloadDeletionLag(ctx, tenant, namespace).Execute()

Number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster's local storage (i.e. BookKeeper)



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetOffloadDeletionLag(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetOffloadDeletionLag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffloadDeletionLag`: int64
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetOffloadDeletionLag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffloadDeletionLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadPolicies

> OffloadPoliciesImpl GetOffloadPolicies(ctx, tenant, namespace).Execute()

Get offload configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetOffloadPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetOffloadPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffloadPolicies`: OffloadPoliciesImpl
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetOffloadPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffloadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OffloadPoliciesImpl**](OffloadPoliciesImpl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffloadThreshold

> int64 GetOffloadThreshold(ctx, tenant, namespace).Execute()

Maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetOffloadThreshold(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetOffloadThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffloadThreshold`: int64
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetOffloadThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffloadThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> map[string][]string GetPermissions(ctx, tenant, namespace).Execute()

Retrieve the permissions for a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetPermissions(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


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

> PersistencePolicies GetPersistence(ctx, tenant, namespace).Execute()

Get the persistence configuration for a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetPersistence(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetPersistence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersistence`: PersistencePolicies
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetPersistence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersistencePolicies**](PersistencePolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> Policies GetPolicies(ctx, tenant, namespace).Execute()

Get the dump all the policies specified for a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicies`: Policies
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Policies**](Policies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProperties

> GetProperties(ctx, tenant, namespace).Execute()

Get key value pair properties for a given namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetProperties(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetProperties``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesRequest struct via the builder pattern


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


## GetProperty

> GetProperty(ctx, tenant, namespace, key).Execute()

Get property value for a given key on a namespace.

### Example

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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetProperty(context.Background(), tenant, namespace, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetProperty``: %v\n", err)
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
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyRequest struct via the builder pattern


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


## GetReplicatorDispatchRate

> DispatchRate GetReplicatorDispatchRate(ctx, tenant, namespace).Execute()

Get replicator dispatch-rate configured for the namespace, -1 represents not configured yet

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetReplicatorDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetReplicatorDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicatorDispatchRate`: DispatchRate
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetReplicatorDispatchRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicatorDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetention

> RetentionPolicies GetRetention(ctx, tenant, namespace).Execute()

Get retention config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetRetention(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRetention`: RetentionPolicies
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RetentionPolicies**](RetentionPolicies.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaAutoUpdateCompatibilityStrategy

> string GetSchemaAutoUpdateCompatibilityStrategy(ctx, tenant, namespace).Execute()

The strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSchemaAutoUpdateCompatibilityStrategy(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSchemaAutoUpdateCompatibilityStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaAutoUpdateCompatibilityStrategy`: string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSchemaAutoUpdateCompatibilityStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaAutoUpdateCompatibilityStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaCompatibilityStrategy

> string GetSchemaCompatibilityStrategy(ctx, tenant, namespace).Execute()

The strategy of the namespace schema compatibility 

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSchemaCompatibilityStrategy(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSchemaCompatibilityStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaCompatibilityStrategy`: string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSchemaCompatibilityStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaCompatibilityStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaValidtionEnforced

> bool GetSchemaValidtionEnforced(ctx, tenant, namespace).Execute()

Get schema validation enforced flag for namespace.



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSchemaValidtionEnforced(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSchemaValidtionEnforced``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaValidtionEnforced`: bool
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSchemaValidtionEnforced`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaValidtionEnforcedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribeRate

> SubscribeRate GetSubscribeRate(ctx, tenant, namespace).Execute()

Get subscribe-rate configured for the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSubscribeRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSubscribeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscribeRate`: SubscribeRate
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSubscribeRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscribeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscribeRate**](SubscribeRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionDispatchRate

> DispatchRate GetSubscriptionDispatchRate(ctx, tenant, namespace).Execute()

Get Subscription dispatch-rate configured for the namespace, -1 represents not configured yet

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSubscriptionDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSubscriptionDispatchRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionDispatchRate`: DispatchRate
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSubscriptionDispatchRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DispatchRate**](DispatchRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionExpirationTime

> int32 GetSubscriptionExpirationTime(ctx, tenant, namespace).Execute()

Get the subscription expiration time for the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSubscriptionExpirationTime(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSubscriptionExpirationTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionExpirationTime`: int32
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSubscriptionExpirationTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionExpirationTimeRequest struct via the builder pattern


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


## GetSubscriptionTypesEnabled

> []string GetSubscriptionTypesEnabled(ctx, tenant, namespace).Execute()

The set of whether allow subscription types

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetSubscriptionTypesEnabled(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetSubscriptionTypesEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionTypesEnabled`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetSubscriptionTypesEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionTypesEnabledRequest struct via the builder pattern


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


## GetTenantNamespaces

> []string GetTenantNamespaces(ctx, tenant).Execute()

Get the list of all the namespaces for a certain tenant.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetTenantNamespaces(context.Background(), tenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetTenantNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantNamespaces`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetTenantNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantNamespacesRequest struct via the builder pattern


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


## GetTopics

> []string GetTopics(ctx, tenant, namespace).Mode(mode).Execute()

Get the list of all the topics under a certain namespace.

### Example

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
    mode := "mode_example" // string |  (optional) (default to "PERSISTENT")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GetTopics(context.Background(), tenant, namespace).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopics`: []string
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **string** |  | [default to &quot;PERSISTENT&quot;]

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


## GrantPermissionOnNamespace

> GrantPermissionOnNamespace(ctx, tenant, namespace, role).Body(body).Execute()

Grant a new permission to a role on a namespace.

### Example

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
    role := "role_example" // string | 
    body := []string{"Property_example"} // []string | List of permissions for the specified role (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.GrantPermissionOnNamespace(context.Background(), tenant, namespace, role).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GrantPermissionOnNamespace``: %v\n", err)
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
**role** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPermissionOnNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **[]string** | List of permissions for the specified role | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDeduplication

> ModifyDeduplication(ctx, tenant, namespace).Body(body).Execute()

Enable or disable broker side deduplication for all topics in a namespace

### Example

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
    body := true // bool | Flag for disabling or enabling broker side deduplication for all topics in the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ModifyDeduplication(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ModifyDeduplication``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDeduplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **bool** | Flag for disabling or enabling broker side deduplication for all topics in the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyEncryptionRequired

> ModifyEncryptionRequired(ctx, tenant, namespace).Body(body).Execute()

Message encryption is required or not for all topics in a namespace

### Example

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
    body := true // bool | Flag defining if message encryption is required

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.ModifyEncryptionRequired(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.ModifyEncryptionRequired``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiModifyEncryptionRequiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **bool** | Flag defining if message encryption is required | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAutoSubscriptionCreation

> RemoveAutoSubscriptionCreation(ctx, tenant, namespace).Execute()

Remove override of broker's allowAutoSubscriptionCreation in a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveAutoSubscriptionCreation(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveAutoSubscriptionCreation``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAutoSubscriptionCreationRequest struct via the builder pattern


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


## RemoveAutoTopicCreation

> RemoveAutoTopicCreation(ctx, tenant, namespace).Execute()

Remove override of broker's allowAutoTopicCreation in a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveAutoTopicCreation(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveAutoTopicCreation``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAutoTopicCreationRequest struct via the builder pattern


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


## RemoveBacklogQuota

> RemoveBacklogQuota(ctx, tenant, namespace).BacklogQuotaType(backlogQuotaType).Execute()

Remove a backlog quota policy from a namespace.

### Example

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
    backlogQuotaType := "backlogQuotaType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveBacklogQuota(context.Background(), tenant, namespace).BacklogQuotaType(backlogQuotaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveBacklogQuota``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBacklogQuotaRequest struct via the builder pattern


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


## RemoveDeduplication

> RemoveDeduplication(ctx, tenant, namespace).Execute()

Remove broker side deduplication for all topics in a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveDeduplication(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveDeduplication``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeduplicationRequest struct via the builder pattern


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


## RemoveDelayedDeliveryPolicies

> RemoveDelayedDeliveryPolicies(ctx, tenant, namespace).Execute()

Delete delayed delivery messages config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveDelayedDeliveryPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveDelayedDeliveryPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDelayedDeliveryPoliciesRequest struct via the builder pattern


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


## RemoveInactiveTopicPolicies

> RemoveInactiveTopicPolicies(ctx, tenant, namespace).Execute()

Remove inactive topic policies from a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveInactiveTopicPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveInactiveTopicPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInactiveTopicPoliciesRequest struct via the builder pattern


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


## RemoveMaxConsumersPerSubscription

> RemoveMaxConsumersPerSubscription(ctx, tenant, namespace).Execute()

 Set maxConsumersPerSubscription configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxConsumersPerSubscription(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxConsumersPerSubscription``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxConsumersPerSubscriptionRequest struct via the builder pattern


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


## RemoveMaxConsumersPerTopic

> RemoveMaxConsumersPerTopic(ctx, tenant, namespace).Execute()

Remove maxConsumersPerTopic configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxConsumersPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxConsumersPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxConsumersPerTopicRequest struct via the builder pattern


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


## RemoveMaxProducersPerTopic

> RemoveMaxProducersPerTopic(ctx, tenant, namespace).Execute()

Remove maxProducersPerTopic configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxProducersPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxProducersPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxProducersPerTopicRequest struct via the builder pattern


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


## RemoveMaxSubscriptionsPerTopic

> RemoveMaxSubscriptionsPerTopic(ctx, tenant, namespace).Execute()

Remove maxSubscriptionsPerTopic configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxSubscriptionsPerTopic(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxSubscriptionsPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxSubscriptionsPerTopicRequest struct via the builder pattern


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


## RemoveMaxUnackedmessagesPerConsumer

> RemoveMaxUnackedmessagesPerConsumer(ctx, tenant, namespace).Execute()

Remove maxUnackedMessagesPerConsumer config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxUnackedmessagesPerConsumer(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxUnackedmessagesPerConsumer``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxUnackedmessagesPerConsumerRequest struct via the builder pattern


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


## RemoveMaxUnackedmessagesPerSubscription

> RemoveMaxUnackedmessagesPerSubscription(ctx, tenant, namespace).Execute()

Remove maxUnackedMessagesPerSubscription config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveMaxUnackedmessagesPerSubscription(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveMaxUnackedmessagesPerSubscription``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaxUnackedmessagesPerSubscriptionRequest struct via the builder pattern


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


## RemoveNamespaceAntiAffinityGroup

> RemoveNamespaceAntiAffinityGroup(ctx, tenant, namespace).Execute()

Remove anti-affinity group of a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveNamespaceAntiAffinityGroup(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveNamespaceAntiAffinityGroup``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceAntiAffinityGroupRequest struct via the builder pattern


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


## RemoveNamespaceMessageTTL

> RemoveNamespaceMessageTTL(ctx, tenant, namespace).Execute()

Set message TTL in seconds for namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveNamespaceMessageTTL(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveNamespaceMessageTTL``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceMessageTTLRequest struct via the builder pattern


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


## RemoveNamespaceResourceGroup

> RemoveNamespaceResourceGroup(ctx, tenant, namespace).Execute()

Delete resourcegroup for a namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveNamespaceResourceGroup(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveNamespaceResourceGroup``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNamespaceResourceGroupRequest struct via the builder pattern


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


## RemoveOffloadPolicies

> RemoveOffloadPolicies(ctx, tenant, namespace).Execute()

 Set offload configuration on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveOffloadPolicies(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveOffloadPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOffloadPoliciesRequest struct via the builder pattern


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


## RemoveProperty

> RemoveProperty(ctx, tenant, namespace, key).Execute()

Get property value for a given key on a namespace.

### Example

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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveProperty(context.Background(), tenant, namespace, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveProperty``: %v\n", err)
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
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePropertyRequest struct via the builder pattern


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


## RemoveReplicatorDispatchRate

> RemoveReplicatorDispatchRate(ctx, tenant, namespace).Execute()

Remove replicator dispatch-rate throttling for all topics of the namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveReplicatorDispatchRate(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveReplicatorDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReplicatorDispatchRateRequest struct via the builder pattern


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


## RemoveRetention

> RemoveRetention(ctx, tenant, namespace).Body(body).Execute()

 Remove retention configuration on a namespace.

### Example

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
    body := *openapiclient.NewRetentionPolicies() // RetentionPolicies | Retention policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveRetention(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveRetention``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RetentionPolicies**](RetentionPolicies.md) | Retention policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSubscriptionExpirationTime

> RemoveSubscriptionExpirationTime(ctx, tenant, namespace).Execute()

Remove subscription expiration time for namespace

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RemoveSubscriptionExpirationTime(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RemoveSubscriptionExpirationTime``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscriptionExpirationTimeRequest struct via the builder pattern


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


## RevokePermissionsOnNamespace

> RevokePermissionsOnNamespace(ctx, tenant, namespace, role).Execute()

Revoke all permissions to a role on a namespace.

### Example

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
    role := "role_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.RevokePermissionsOnNamespace(context.Background(), tenant, namespace, role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.RevokePermissionsOnNamespace``: %v\n", err)
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
**role** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsOnNamespaceRequest struct via the builder pattern


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


## SetAutoSubscriptionCreation

> SetAutoSubscriptionCreation(ctx, tenant, namespace).Body(body).Execute()

Override broker's allowAutoSubscriptionCreation setting for a namespace

### Example

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
    body := *openapiclient.NewAutoSubscriptionCreationOverride() // AutoSubscriptionCreationOverride | Settings for automatic subscription creation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetAutoSubscriptionCreation(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetAutoSubscriptionCreation``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoSubscriptionCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AutoSubscriptionCreationOverride**](AutoSubscriptionCreationOverride.md) | Settings for automatic subscription creation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAutoTopicCreation

> SetAutoTopicCreation(ctx, tenant, namespace).Body(body).Execute()

Override broker's allowAutoTopicCreation setting for a namespace

### Example

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
    body := *openapiclient.NewAutoTopicCreationOverride() // AutoTopicCreationOverride | Settings for automatic topic creation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetAutoTopicCreation(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetAutoTopicCreation``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoTopicCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AutoTopicCreationOverride**](AutoTopicCreationOverride.md) | Settings for automatic topic creation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBacklogQuota

> SetBacklogQuota(ctx, tenant, namespace).BacklogQuotaType(backlogQuotaType).Body(body).Execute()

 Set a backlog quota for all the topics on a namespace.

### Example

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
    backlogQuotaType := "backlogQuotaType_example" // string |  (optional)
    body := *openapiclient.NewBacklogQuota() // BacklogQuota | Backlog quota for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetBacklogQuota(context.Background(), tenant, namespace).BacklogQuotaType(backlogQuotaType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetBacklogQuota``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetBacklogQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backlogQuotaType** | **string** |  | 
 **body** | [**BacklogQuota**](BacklogQuota.md) | Backlog quota for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBookieAffinityGroup

> SetBookieAffinityGroup(ctx, tenant, namespace).Body(body).Execute()

Set the bookie-affinity-group to namespace-persistent policy.

### Example

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
    body := *openapiclient.NewBookieAffinityGroupData() // BookieAffinityGroupData | Bookie affinity group for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetBookieAffinityGroup(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetBookieAffinityGroup``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetBookieAffinityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**BookieAffinityGroupData**](BookieAffinityGroupData.md) | Bookie affinity group for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCompactionThreshold

> SetCompactionThreshold(ctx, tenant, namespace).Body(body).Execute()

Set maximum number of uncompacted bytes in a topic before compaction is triggered.



### Example

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
    body := int64(789) // int64 | Maximum number of uncompacted bytes in a topic of the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetCompactionThreshold(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetCompactionThreshold``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetCompactionThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int64** | Maximum number of uncompacted bytes in a topic of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeduplicationSnapshotInterval

> SetDeduplicationSnapshotInterval(ctx, tenant, namespace).Body(body).Execute()

Set deduplicationSnapshotInterval config on a namespace.

### Example

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
    body := int32(56) // int32 | Interval to take deduplication snapshot per topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetDeduplicationSnapshotInterval(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetDeduplicationSnapshotInterval``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeduplicationSnapshotIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Interval to take deduplication snapshot per topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDelayedDeliveryPolicies

> SetDelayedDeliveryPolicies(ctx, tenant, namespace).Body(body).Execute()

Set delayed delivery messages config on a namespace.

### Example

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
    body := *openapiclient.NewDelayedDeliveryPolicies() // DelayedDeliveryPolicies | Delayed delivery policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetDelayedDeliveryPolicies(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetDelayedDeliveryPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetDelayedDeliveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DelayedDeliveryPolicies**](DelayedDeliveryPolicies.md) | Delayed delivery policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDispatchRate

> SetDispatchRate(ctx, tenant, namespace).Body(body).Execute()

Set dispatch-rate throttling for all topics of the namespace

### Example

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
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Dispatch rate for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetDispatchRate(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Dispatch rate for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInactiveTopicPolicies

> SetInactiveTopicPolicies(ctx, tenant, namespace).Body(body).Execute()

Set inactive topic policies config on a namespace.

### Example

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
    body := *openapiclient.NewInactiveTopicPolicies() // InactiveTopicPolicies | Inactive topic policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetInactiveTopicPolicies(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetInactiveTopicPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetInactiveTopicPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**InactiveTopicPolicies**](InactiveTopicPolicies.md) | Inactive topic policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInactiveTopicPolicies_0

> SetInactiveTopicPolicies_0(ctx, tenant, namespace).Body(body).Execute()

Set maxTopicsPerNamespace config on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum topics for specific namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetInactiveTopicPolicies_0(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetInactiveTopicPolicies_0``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetInactiveTopicPolicies_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum topics for specific namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInactiveTopicPolicies_1

> SetInactiveTopicPolicies_1(ctx, tenant, namespace).Execute()

Set maxTopicsPerNamespace config on a namespace.

### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetInactiveTopicPolicies_1(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetInactiveTopicPolicies_1``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetInactiveTopicPolicies_2Request struct via the builder pattern


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


## SetIsAllowAutoUpdateSchema

> SetIsAllowAutoUpdateSchema(ctx, tenant, namespace).Body(body).Execute()

Update flag of whether allow auto update schema

### Example

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
    body := true // bool | Flag of whether to allow auto update schema

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetIsAllowAutoUpdateSchema(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetIsAllowAutoUpdateSchema``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetIsAllowAutoUpdateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **bool** | Flag of whether to allow auto update schema | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxConsumersPerSubscription

> SetMaxConsumersPerSubscription(ctx, tenant, namespace).Body(body).Execute()

 Set maxConsumersPerSubscription configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum consumers per subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxConsumersPerSubscription(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxConsumersPerSubscription``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxConsumersPerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum consumers per subscription | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxConsumersPerTopic

> SetMaxConsumersPerTopic(ctx, tenant, namespace).Body(body).Execute()

 Set maxConsumersPerTopic configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum consumers per topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxConsumersPerTopic(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxConsumersPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxConsumersPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum consumers per topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxProducersPerTopic

> SetMaxProducersPerTopic(ctx, tenant, namespace).Body(body).Execute()

 Set maxProducersPerTopic configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum producers per topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxProducersPerTopic(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxProducersPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxProducersPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum producers per topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxSubscriptionsPerTopic

> SetMaxSubscriptionsPerTopic(ctx, tenant, namespace).Body(body).Execute()

 Set maxSubscriptionsPerTopic configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum subscriptions per topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxSubscriptionsPerTopic(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxSubscriptionsPerTopic``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxSubscriptionsPerTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum subscriptions per topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxUnackedMessagesPerConsumer

> SetMaxUnackedMessagesPerConsumer(ctx, tenant, namespace).Body(body).Execute()

 Set maxConsumersPerTopic configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum unacked messages per consumer

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxUnackedMessagesPerConsumer(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxUnackedMessagesPerConsumer``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesPerConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum unacked messages per consumer | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxUnackedMessagesPerSubscription

> SetMaxUnackedMessagesPerSubscription(ctx, tenant, namespace).Body(body).Execute()

 Set maxUnackedMessagesPerSubscription configuration on a namespace.

### Example

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
    body := int32(56) // int32 | Number of maximum unacked messages per subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetMaxUnackedMessagesPerSubscription(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetMaxUnackedMessagesPerSubscription``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxUnackedMessagesPerSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Number of maximum unacked messages per subscription | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNamespaceAntiAffinityGroup

> SetNamespaceAntiAffinityGroup(ctx, tenant, namespace).Body(body).Execute()

Set anti-affinity group for a namespace

### Example

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
    body := "body_example" // string | Anti-affinity group for the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetNamespaceAntiAffinityGroup(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetNamespaceAntiAffinityGroup``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceAntiAffinityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Anti-affinity group for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNamespaceMessageTTL

> SetNamespaceMessageTTL(ctx, tenant, namespace).Body(body).Execute()

Set message TTL in seconds for namespace

### Example

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
    body := int32(56) // int32 | TTL in seconds for the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetNamespaceMessageTTL(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetNamespaceMessageTTL``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceMessageTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | TTL in seconds for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNamespaceReplicationClusters

> SetNamespaceReplicationClusters(ctx, tenant, namespace).Body(body).Execute()

Set the replication clusters for a namespace.

### Example

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
    body := []string{"Property_example"} // []string | List of replication clusters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetNamespaceReplicationClusters(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetNamespaceReplicationClusters``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceReplicationClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **[]string** | List of replication clusters | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNamespaceResourceGroup

> SetNamespaceResourceGroup(ctx, tenant, namespace).Body(body).Execute()

Set resourcegroup for a namespace

### Example

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
    body := "body_example" // string | Name of resourcegroup

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetNamespaceResourceGroup(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetNamespaceResourceGroup``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetNamespaceResourceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Name of resourcegroup | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOffloadDeletionLag

> SetOffloadDeletionLag(ctx, tenant, namespace).Body(body).Execute()

Set number of milliseconds to wait before deleting a ledger segment which has been offloaded from the Pulsar cluster's local storage (i.e. BookKeeper)



### Example

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
    body := int64(789) // int64 | New number of milliseconds to wait before deleting a ledger segment which has been offloaded

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetOffloadDeletionLag(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetOffloadDeletionLag``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetOffloadDeletionLagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int64** | New number of milliseconds to wait before deleting a ledger segment which has been offloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOffloadPolicies

> SetOffloadPolicies(ctx, tenant, namespace).Body(body).Execute()

 Set offload configuration on a namespace.

### Example

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
    body := *openapiclient.NewOffloadPoliciesImpl() // OffloadPoliciesImpl | Offload policies for the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetOffloadPolicies(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetOffloadPolicies``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetOffloadPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OffloadPoliciesImpl**](OffloadPoliciesImpl.md) | Offload policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOffloadThreshold

> SetOffloadThreshold(ctx, tenant, namespace).Body(body).Execute()

Set maximum number of bytes stored on the pulsar cluster for a topic, before the broker will start offloading to longterm storage



### Example

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
    body := int64(789) // int64 | Maximum number of bytes stored on the pulsar cluster for a topic of the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetOffloadThreshold(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetOffloadThreshold``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetOffloadThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int64** | Maximum number of bytes stored on the pulsar cluster for a topic of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPersistence

> SetPersistence(ctx, tenant, namespace).Body(body).Execute()

Set the persistence configuration for all the topics on a namespace.

### Example

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
    body := *openapiclient.NewPersistencePolicies() // PersistencePolicies | Persistence policies for the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetPersistence(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetPersistence``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetPersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**PersistencePolicies**](PersistencePolicies.md) | Persistence policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProperties

> SetProperties(ctx, tenant, namespace).Body(body).Execute()

Put key value pairs property on a namespace.

### Example

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
    body := map[string]string{"key": "Inner_example"} // map[string]string | Key value pair properties for the namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetProperties(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetProperties``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]string** | Key value pair properties for the namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProperty

> SetProperty(ctx, tenant, namespace, key, value).Execute()

Put a key value pair property on a namespace.

### Example

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
    key := "key_example" // string | 
    value := "value_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetProperty(context.Background(), tenant, namespace, key, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetProperty``: %v\n", err)
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
**key** | **string** |  | 
**value** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPropertyRequest struct via the builder pattern


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


## SetReplicatorDispatchRate

> SetReplicatorDispatchRate(ctx, tenant, namespace).Body(body).Execute()

Set replicator dispatch-rate throttling for all topics of the namespace

### Example

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
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Replicator dispatch rate for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetReplicatorDispatchRate(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetReplicatorDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetReplicatorDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Replicator dispatch rate for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetention

> SetRetention(ctx, tenant, namespace).Body(body).Execute()

 Set retention configuration on a namespace.

### Example

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
    body := *openapiclient.NewRetentionPolicies() // RetentionPolicies | Retention policies for the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetRetention(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetRetention``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RetentionPolicies**](RetentionPolicies.md) | Retention policies for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSchemaAutoUpdateCompatibilityStrategy

> SetSchemaAutoUpdateCompatibilityStrategy(ctx, tenant, namespace).Body(body).Execute()

Update the strategy used to check the compatibility of new schemas, provided by producers, before automatically updating the schema



### Example

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
    body := "body_example" // string | Strategy used to check the compatibility of new schemas (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSchemaAutoUpdateCompatibilityStrategy(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSchemaAutoUpdateCompatibilityStrategy``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSchemaAutoUpdateCompatibilityStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Strategy used to check the compatibility of new schemas | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSchemaCompatibilityStrategy

> SetSchemaCompatibilityStrategy(ctx, tenant, namespace).Body(body).Execute()

Update the strategy used to check the compatibility of new schema

### Example

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
    body := "body_example" // string | Strategy used to check the compatibility of new schema (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSchemaCompatibilityStrategy(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSchemaCompatibilityStrategy``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSchemaCompatibilityStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Strategy used to check the compatibility of new schema | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSchemaValidtionEnforced

> SetSchemaValidtionEnforced(ctx, tenant, namespace).Body(body).Execute()

Set schema validation enforced flag on namespace.



### Example

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
    body := true // bool | Flag of whether validation is enforced on the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSchemaValidtionEnforced(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSchemaValidtionEnforced``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSchemaValidtionEnforcedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **bool** | Flag of whether validation is enforced on the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscribeRate

> SetSubscribeRate(ctx, tenant, namespace).Body(body).Execute()

Set subscribe-rate throttling for all topics of the namespace

### Example

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
    body := *openapiclient.NewSubscribeRate() // SubscribeRate | Subscribe rate for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSubscribeRate(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSubscribeRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscribeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SubscribeRate**](SubscribeRate.md) | Subscribe rate for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionAuthMode

> SetSubscriptionAuthMode(ctx, tenant, namespace).Body(body).Execute()

 Set a subscription auth mode for all the topics on a namespace.

### Example

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
    body := "body_example" // string | Subscription auth mode for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSubscriptionAuthMode(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSubscriptionAuthMode``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionAuthModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Subscription auth mode for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionDispatchRate

> SetSubscriptionDispatchRate(ctx, tenant, namespace).Body(body).Execute()

Set Subscription dispatch-rate throttling for all topics of the namespace

### Example

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
    body := *openapiclient.NewDispatchRateImpl() // DispatchRateImpl | Subscription dispatch rate for all topics of the specified namespace (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSubscriptionDispatchRate(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSubscriptionDispatchRate``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionDispatchRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DispatchRateImpl**](DispatchRateImpl.md) | Subscription dispatch rate for all topics of the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionExpirationTime

> SetSubscriptionExpirationTime(ctx, tenant, namespace).Body(body).Execute()

Set subscription expiration time in minutes for namespace

### Example

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
    body := int32(56) // int32 | Expiration time in minutes for the specified namespace

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSubscriptionExpirationTime(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSubscriptionExpirationTime``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionExpirationTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **int32** | Expiration time in minutes for the specified namespace | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriptionTypesEnabled

> SetSubscriptionTypesEnabled(ctx, tenant, namespace).Body(body).Execute()

Update set of whether allow share sub type

### Example

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
    body := []string{"Property_example"} // []string | Set of whether allow subscription types

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SetSubscriptionTypesEnabled(context.Background(), tenant, namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SetSubscriptionTypesEnabled``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriptionTypesEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **[]string** | Set of whether allow subscription types | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SplitNamespaceBundle

> SplitNamespaceBundle(ctx, tenant, namespace, bundle).Authoritative(authoritative).Unload(unload).SplitAlgorithmName(splitAlgorithmName).Execute()

Split a namespace bundle

### Example

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
    bundle := "bundle_example" // string | 
    authoritative := true // bool |  (optional) (default to false)
    unload := true // bool |  (optional) (default to false)
    splitAlgorithmName := "splitAlgorithmName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.SplitNamespaceBundle(context.Background(), tenant, namespace, bundle).Authoritative(authoritative).Unload(unload).SplitAlgorithmName(splitAlgorithmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.SplitNamespaceBundle``: %v\n", err)
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
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitNamespaceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]
 **unload** | **bool** |  | [default to false]
 **splitAlgorithmName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnloadNamespace

> UnloadNamespace(ctx, tenant, namespace).Execute()

Unload namespace



### Example

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.UnloadNamespace(context.Background(), tenant, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.UnloadNamespace``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUnloadNamespaceRequest struct via the builder pattern


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


## UnloadNamespaceBundle

> UnloadNamespaceBundle(ctx, tenant, namespace, bundle).Authoritative(authoritative).Execute()

Unload a namespace bundle

### Example

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
    bundle := "bundle_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.UnloadNamespaceBundle(context.Background(), tenant, namespace, bundle).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.UnloadNamespaceBundle``: %v\n", err)
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
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnloadNamespaceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeNamespace

> UnsubscribeNamespace(ctx, tenant, namespace, subscription).Authoritative(authoritative).Execute()

Unsubscribes the given subscription on all topics on a namespace.

### Example

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
    subscription := "subscription_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.UnsubscribeNamespace(context.Background(), tenant, namespace, subscription).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.UnsubscribeNamespace``: %v\n", err)
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
**subscription** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeNamespaceBundle

> UnsubscribeNamespaceBundle(ctx, tenant, namespace, subscription, bundle).Authoritative(authoritative).Execute()

Unsubscribes the given subscription on all topics on a namespace bundle.

### Example

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
    subscription := "subscription_example" // string | 
    bundle := "bundle_example" // string | 
    authoritative := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NamespacesApi.UnsubscribeNamespaceBundle(context.Background(), tenant, namespace, subscription, bundle).Authoritative(authoritative).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.UnsubscribeNamespaceBundle``: %v\n", err)
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
**subscription** | **string** |  | 
**bundle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeNamespaceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authoritative** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

