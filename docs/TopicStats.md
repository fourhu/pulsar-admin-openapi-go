# TopicStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**BacklogSize** | Pointer to **int64** |  | [optional] 
**Subscriptions** | Pointer to [**map[string]SubscriptionStats**](SubscriptionStats.md) |  | [optional] 
**BytesInCounter** | Pointer to **int64** |  | [optional] 
**MsgInCounter** | Pointer to **int64** |  | [optional] 
**BytesOutCounter** | Pointer to **int64** |  | [optional] 
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**OffloadedStorageSize** | Pointer to **int64** |  | [optional] 
**Publishers** | Pointer to [**[]PublisherStats**](PublisherStats.md) |  | [optional] 
**WaitingPublishers** | Pointer to **int32** |  | [optional] 
**Replication** | Pointer to [**map[string]ReplicatorStats**](ReplicatorStats.md) |  | [optional] 
**DeduplicationStatus** | Pointer to **string** |  | [optional] 
**TopicEpoch** | Pointer to **int64** |  | [optional] 
**NonContiguousDeletedMessagesRanges** | Pointer to **int32** |  | [optional] 
**NonContiguousDeletedMessagesRangesSerializedSize** | Pointer to **int32** |  | [optional] 
**Compaction** | Pointer to [**CompactionStats**](CompactionStats.md) |  | [optional] 
**AverageMsgSize** | Pointer to **float64** |  | [optional] 
**MsgChunkPublished** | Pointer to **bool** |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 

## Methods

### NewTopicStats

`func NewTopicStats() *TopicStats`

NewTopicStats instantiates a new TopicStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicStatsWithDefaults

`func NewTopicStatsWithDefaults() *TopicStats`

NewTopicStatsWithDefaults instantiates a new TopicStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgThroughputIn

`func (o *TopicStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *TopicStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *TopicStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *TopicStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *TopicStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *TopicStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *TopicStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *TopicStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetBacklogSize

`func (o *TopicStats) GetBacklogSize() int64`

GetBacklogSize returns the BacklogSize field if non-nil, zero value otherwise.

### GetBacklogSizeOk

`func (o *TopicStats) GetBacklogSizeOk() (*int64, bool)`

GetBacklogSizeOk returns a tuple with the BacklogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogSize

`func (o *TopicStats) SetBacklogSize(v int64)`

SetBacklogSize sets BacklogSize field to given value.

### HasBacklogSize

`func (o *TopicStats) HasBacklogSize() bool`

HasBacklogSize returns a boolean if a field has been set.

### GetSubscriptions

`func (o *TopicStats) GetSubscriptions() map[string]SubscriptionStats`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *TopicStats) GetSubscriptionsOk() (*map[string]SubscriptionStats, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *TopicStats) SetSubscriptions(v map[string]SubscriptionStats)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *TopicStats) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetBytesInCounter

`func (o *TopicStats) GetBytesInCounter() int64`

GetBytesInCounter returns the BytesInCounter field if non-nil, zero value otherwise.

### GetBytesInCounterOk

`func (o *TopicStats) GetBytesInCounterOk() (*int64, bool)`

GetBytesInCounterOk returns a tuple with the BytesInCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesInCounter

`func (o *TopicStats) SetBytesInCounter(v int64)`

SetBytesInCounter sets BytesInCounter field to given value.

### HasBytesInCounter

`func (o *TopicStats) HasBytesInCounter() bool`

HasBytesInCounter returns a boolean if a field has been set.

### GetMsgInCounter

`func (o *TopicStats) GetMsgInCounter() int64`

GetMsgInCounter returns the MsgInCounter field if non-nil, zero value otherwise.

### GetMsgInCounterOk

`func (o *TopicStats) GetMsgInCounterOk() (*int64, bool)`

GetMsgInCounterOk returns a tuple with the MsgInCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgInCounter

`func (o *TopicStats) SetMsgInCounter(v int64)`

SetMsgInCounter sets MsgInCounter field to given value.

### HasMsgInCounter

`func (o *TopicStats) HasMsgInCounter() bool`

HasMsgInCounter returns a boolean if a field has been set.

### GetBytesOutCounter

`func (o *TopicStats) GetBytesOutCounter() int64`

GetBytesOutCounter returns the BytesOutCounter field if non-nil, zero value otherwise.

### GetBytesOutCounterOk

`func (o *TopicStats) GetBytesOutCounterOk() (*int64, bool)`

GetBytesOutCounterOk returns a tuple with the BytesOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutCounter

`func (o *TopicStats) SetBytesOutCounter(v int64)`

SetBytesOutCounter sets BytesOutCounter field to given value.

### HasBytesOutCounter

`func (o *TopicStats) HasBytesOutCounter() bool`

HasBytesOutCounter returns a boolean if a field has been set.

### GetMsgOutCounter

`func (o *TopicStats) GetMsgOutCounter() int64`

GetMsgOutCounter returns the MsgOutCounter field if non-nil, zero value otherwise.

### GetMsgOutCounterOk

`func (o *TopicStats) GetMsgOutCounterOk() (*int64, bool)`

GetMsgOutCounterOk returns a tuple with the MsgOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgOutCounter

`func (o *TopicStats) SetMsgOutCounter(v int64)`

SetMsgOutCounter sets MsgOutCounter field to given value.

### HasMsgOutCounter

`func (o *TopicStats) HasMsgOutCounter() bool`

HasMsgOutCounter returns a boolean if a field has been set.

### GetOffloadedStorageSize

`func (o *TopicStats) GetOffloadedStorageSize() int64`

GetOffloadedStorageSize returns the OffloadedStorageSize field if non-nil, zero value otherwise.

### GetOffloadedStorageSizeOk

`func (o *TopicStats) GetOffloadedStorageSizeOk() (*int64, bool)`

GetOffloadedStorageSizeOk returns a tuple with the OffloadedStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadedStorageSize

`func (o *TopicStats) SetOffloadedStorageSize(v int64)`

SetOffloadedStorageSize sets OffloadedStorageSize field to given value.

### HasOffloadedStorageSize

`func (o *TopicStats) HasOffloadedStorageSize() bool`

HasOffloadedStorageSize returns a boolean if a field has been set.

### GetPublishers

`func (o *TopicStats) GetPublishers() []PublisherStats`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *TopicStats) GetPublishersOk() (*[]PublisherStats, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *TopicStats) SetPublishers(v []PublisherStats)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *TopicStats) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### GetWaitingPublishers

`func (o *TopicStats) GetWaitingPublishers() int32`

GetWaitingPublishers returns the WaitingPublishers field if non-nil, zero value otherwise.

### GetWaitingPublishersOk

`func (o *TopicStats) GetWaitingPublishersOk() (*int32, bool)`

GetWaitingPublishersOk returns a tuple with the WaitingPublishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingPublishers

`func (o *TopicStats) SetWaitingPublishers(v int32)`

SetWaitingPublishers sets WaitingPublishers field to given value.

### HasWaitingPublishers

`func (o *TopicStats) HasWaitingPublishers() bool`

HasWaitingPublishers returns a boolean if a field has been set.

### GetReplication

`func (o *TopicStats) GetReplication() map[string]ReplicatorStats`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *TopicStats) GetReplicationOk() (*map[string]ReplicatorStats, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *TopicStats) SetReplication(v map[string]ReplicatorStats)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *TopicStats) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetDeduplicationStatus

`func (o *TopicStats) GetDeduplicationStatus() string`

GetDeduplicationStatus returns the DeduplicationStatus field if non-nil, zero value otherwise.

### GetDeduplicationStatusOk

`func (o *TopicStats) GetDeduplicationStatusOk() (*string, bool)`

GetDeduplicationStatusOk returns a tuple with the DeduplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationStatus

`func (o *TopicStats) SetDeduplicationStatus(v string)`

SetDeduplicationStatus sets DeduplicationStatus field to given value.

### HasDeduplicationStatus

`func (o *TopicStats) HasDeduplicationStatus() bool`

HasDeduplicationStatus returns a boolean if a field has been set.

### GetTopicEpoch

`func (o *TopicStats) GetTopicEpoch() int64`

GetTopicEpoch returns the TopicEpoch field if non-nil, zero value otherwise.

### GetTopicEpochOk

`func (o *TopicStats) GetTopicEpochOk() (*int64, bool)`

GetTopicEpochOk returns a tuple with the TopicEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEpoch

`func (o *TopicStats) SetTopicEpoch(v int64)`

SetTopicEpoch sets TopicEpoch field to given value.

### HasTopicEpoch

`func (o *TopicStats) HasTopicEpoch() bool`

HasTopicEpoch returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRanges

`func (o *TopicStats) GetNonContiguousDeletedMessagesRanges() int32`

GetNonContiguousDeletedMessagesRanges returns the NonContiguousDeletedMessagesRanges field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesOk

`func (o *TopicStats) GetNonContiguousDeletedMessagesRangesOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesOk returns a tuple with the NonContiguousDeletedMessagesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRanges

`func (o *TopicStats) SetNonContiguousDeletedMessagesRanges(v int32)`

SetNonContiguousDeletedMessagesRanges sets NonContiguousDeletedMessagesRanges field to given value.

### HasNonContiguousDeletedMessagesRanges

`func (o *TopicStats) HasNonContiguousDeletedMessagesRanges() bool`

HasNonContiguousDeletedMessagesRanges returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *TopicStats) GetNonContiguousDeletedMessagesRangesSerializedSize() int32`

GetNonContiguousDeletedMessagesRangesSerializedSize returns the NonContiguousDeletedMessagesRangesSerializedSize field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesSerializedSizeOk

`func (o *TopicStats) GetNonContiguousDeletedMessagesRangesSerializedSizeOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesSerializedSizeOk returns a tuple with the NonContiguousDeletedMessagesRangesSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *TopicStats) SetNonContiguousDeletedMessagesRangesSerializedSize(v int32)`

SetNonContiguousDeletedMessagesRangesSerializedSize sets NonContiguousDeletedMessagesRangesSerializedSize field to given value.

### HasNonContiguousDeletedMessagesRangesSerializedSize

`func (o *TopicStats) HasNonContiguousDeletedMessagesRangesSerializedSize() bool`

HasNonContiguousDeletedMessagesRangesSerializedSize returns a boolean if a field has been set.

### GetCompaction

`func (o *TopicStats) GetCompaction() CompactionStats`

GetCompaction returns the Compaction field if non-nil, zero value otherwise.

### GetCompactionOk

`func (o *TopicStats) GetCompactionOk() (*CompactionStats, bool)`

GetCompactionOk returns a tuple with the Compaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompaction

`func (o *TopicStats) SetCompaction(v CompactionStats)`

SetCompaction sets Compaction field to given value.

### HasCompaction

`func (o *TopicStats) HasCompaction() bool`

HasCompaction returns a boolean if a field has been set.

### GetAverageMsgSize

`func (o *TopicStats) GetAverageMsgSize() float64`

GetAverageMsgSize returns the AverageMsgSize field if non-nil, zero value otherwise.

### GetAverageMsgSizeOk

`func (o *TopicStats) GetAverageMsgSizeOk() (*float64, bool)`

GetAverageMsgSizeOk returns a tuple with the AverageMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMsgSize

`func (o *TopicStats) SetAverageMsgSize(v float64)`

SetAverageMsgSize sets AverageMsgSize field to given value.

### HasAverageMsgSize

`func (o *TopicStats) HasAverageMsgSize() bool`

HasAverageMsgSize returns a boolean if a field has been set.

### GetMsgChunkPublished

`func (o *TopicStats) GetMsgChunkPublished() bool`

GetMsgChunkPublished returns the MsgChunkPublished field if non-nil, zero value otherwise.

### GetMsgChunkPublishedOk

`func (o *TopicStats) GetMsgChunkPublishedOk() (*bool, bool)`

GetMsgChunkPublishedOk returns a tuple with the MsgChunkPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgChunkPublished

`func (o *TopicStats) SetMsgChunkPublished(v bool)`

SetMsgChunkPublished sets MsgChunkPublished field to given value.

### HasMsgChunkPublished

`func (o *TopicStats) HasMsgChunkPublished() bool`

HasMsgChunkPublished returns a boolean if a field has been set.

### GetStorageSize

`func (o *TopicStats) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *TopicStats) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *TopicStats) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *TopicStats) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *TopicStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *TopicStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *TopicStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *TopicStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *TopicStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *TopicStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *TopicStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *TopicStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


