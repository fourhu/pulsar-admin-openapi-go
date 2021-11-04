# NonPersistentTopicStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageMsgSize** | Pointer to **float64** |  | [optional] 
**BacklogSize** | Pointer to **int64** |  | [optional] 
**BytesInCounter** | Pointer to **int64** |  | [optional] 
**BytesOutCounter** | Pointer to **int64** |  | [optional] 
**DeduplicationStatus** | Pointer to **string** |  | [optional] 
**MsgChunkPublished** | Pointer to **bool** |  | [optional] 
**MsgDropRate** | Pointer to **float64** |  | [optional] 
**MsgInCounter** | Pointer to **int64** |  | [optional] 
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**Publishers** | Pointer to [**[]NonPersistentPublisherStats**](NonPersistentPublisherStats.md) |  | [optional] 
**Replication** | Pointer to [**map[string]NonPersistentReplicatorStats**](NonPersistentReplicatorStats.md) |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**Subscriptions** | Pointer to [**map[string]NonPersistentSubscriptionStats**](NonPersistentSubscriptionStats.md) |  | [optional] 

## Methods

### NewNonPersistentTopicStats

`func NewNonPersistentTopicStats() *NonPersistentTopicStats`

NewNonPersistentTopicStats instantiates a new NonPersistentTopicStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPersistentTopicStatsWithDefaults

`func NewNonPersistentTopicStatsWithDefaults() *NonPersistentTopicStats`

NewNonPersistentTopicStatsWithDefaults instantiates a new NonPersistentTopicStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageMsgSize

`func (o *NonPersistentTopicStats) GetAverageMsgSize() float64`

GetAverageMsgSize returns the AverageMsgSize field if non-nil, zero value otherwise.

### GetAverageMsgSizeOk

`func (o *NonPersistentTopicStats) GetAverageMsgSizeOk() (*float64, bool)`

GetAverageMsgSizeOk returns a tuple with the AverageMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageMsgSize

`func (o *NonPersistentTopicStats) SetAverageMsgSize(v float64)`

SetAverageMsgSize sets AverageMsgSize field to given value.

### HasAverageMsgSize

`func (o *NonPersistentTopicStats) HasAverageMsgSize() bool`

HasAverageMsgSize returns a boolean if a field has been set.

### GetBacklogSize

`func (o *NonPersistentTopicStats) GetBacklogSize() int64`

GetBacklogSize returns the BacklogSize field if non-nil, zero value otherwise.

### GetBacklogSizeOk

`func (o *NonPersistentTopicStats) GetBacklogSizeOk() (*int64, bool)`

GetBacklogSizeOk returns a tuple with the BacklogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogSize

`func (o *NonPersistentTopicStats) SetBacklogSize(v int64)`

SetBacklogSize sets BacklogSize field to given value.

### HasBacklogSize

`func (o *NonPersistentTopicStats) HasBacklogSize() bool`

HasBacklogSize returns a boolean if a field has been set.

### GetBytesInCounter

`func (o *NonPersistentTopicStats) GetBytesInCounter() int64`

GetBytesInCounter returns the BytesInCounter field if non-nil, zero value otherwise.

### GetBytesInCounterOk

`func (o *NonPersistentTopicStats) GetBytesInCounterOk() (*int64, bool)`

GetBytesInCounterOk returns a tuple with the BytesInCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesInCounter

`func (o *NonPersistentTopicStats) SetBytesInCounter(v int64)`

SetBytesInCounter sets BytesInCounter field to given value.

### HasBytesInCounter

`func (o *NonPersistentTopicStats) HasBytesInCounter() bool`

HasBytesInCounter returns a boolean if a field has been set.

### GetBytesOutCounter

`func (o *NonPersistentTopicStats) GetBytesOutCounter() int64`

GetBytesOutCounter returns the BytesOutCounter field if non-nil, zero value otherwise.

### GetBytesOutCounterOk

`func (o *NonPersistentTopicStats) GetBytesOutCounterOk() (*int64, bool)`

GetBytesOutCounterOk returns a tuple with the BytesOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutCounter

`func (o *NonPersistentTopicStats) SetBytesOutCounter(v int64)`

SetBytesOutCounter sets BytesOutCounter field to given value.

### HasBytesOutCounter

`func (o *NonPersistentTopicStats) HasBytesOutCounter() bool`

HasBytesOutCounter returns a boolean if a field has been set.

### GetDeduplicationStatus

`func (o *NonPersistentTopicStats) GetDeduplicationStatus() string`

GetDeduplicationStatus returns the DeduplicationStatus field if non-nil, zero value otherwise.

### GetDeduplicationStatusOk

`func (o *NonPersistentTopicStats) GetDeduplicationStatusOk() (*string, bool)`

GetDeduplicationStatusOk returns a tuple with the DeduplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationStatus

`func (o *NonPersistentTopicStats) SetDeduplicationStatus(v string)`

SetDeduplicationStatus sets DeduplicationStatus field to given value.

### HasDeduplicationStatus

`func (o *NonPersistentTopicStats) HasDeduplicationStatus() bool`

HasDeduplicationStatus returns a boolean if a field has been set.

### GetMsgChunkPublished

`func (o *NonPersistentTopicStats) GetMsgChunkPublished() bool`

GetMsgChunkPublished returns the MsgChunkPublished field if non-nil, zero value otherwise.

### GetMsgChunkPublishedOk

`func (o *NonPersistentTopicStats) GetMsgChunkPublishedOk() (*bool, bool)`

GetMsgChunkPublishedOk returns a tuple with the MsgChunkPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgChunkPublished

`func (o *NonPersistentTopicStats) SetMsgChunkPublished(v bool)`

SetMsgChunkPublished sets MsgChunkPublished field to given value.

### HasMsgChunkPublished

`func (o *NonPersistentTopicStats) HasMsgChunkPublished() bool`

HasMsgChunkPublished returns a boolean if a field has been set.

### GetMsgDropRate

`func (o *NonPersistentTopicStats) GetMsgDropRate() float64`

GetMsgDropRate returns the MsgDropRate field if non-nil, zero value otherwise.

### GetMsgDropRateOk

`func (o *NonPersistentTopicStats) GetMsgDropRateOk() (*float64, bool)`

GetMsgDropRateOk returns a tuple with the MsgDropRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDropRate

`func (o *NonPersistentTopicStats) SetMsgDropRate(v float64)`

SetMsgDropRate sets MsgDropRate field to given value.

### HasMsgDropRate

`func (o *NonPersistentTopicStats) HasMsgDropRate() bool`

HasMsgDropRate returns a boolean if a field has been set.

### GetMsgInCounter

`func (o *NonPersistentTopicStats) GetMsgInCounter() int64`

GetMsgInCounter returns the MsgInCounter field if non-nil, zero value otherwise.

### GetMsgInCounterOk

`func (o *NonPersistentTopicStats) GetMsgInCounterOk() (*int64, bool)`

GetMsgInCounterOk returns a tuple with the MsgInCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgInCounter

`func (o *NonPersistentTopicStats) SetMsgInCounter(v int64)`

SetMsgInCounter sets MsgInCounter field to given value.

### HasMsgInCounter

`func (o *NonPersistentTopicStats) HasMsgInCounter() bool`

HasMsgInCounter returns a boolean if a field has been set.

### GetMsgOutCounter

`func (o *NonPersistentTopicStats) GetMsgOutCounter() int64`

GetMsgOutCounter returns the MsgOutCounter field if non-nil, zero value otherwise.

### GetMsgOutCounterOk

`func (o *NonPersistentTopicStats) GetMsgOutCounterOk() (*int64, bool)`

GetMsgOutCounterOk returns a tuple with the MsgOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgOutCounter

`func (o *NonPersistentTopicStats) SetMsgOutCounter(v int64)`

SetMsgOutCounter sets MsgOutCounter field to given value.

### HasMsgOutCounter

`func (o *NonPersistentTopicStats) HasMsgOutCounter() bool`

HasMsgOutCounter returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *NonPersistentTopicStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *NonPersistentTopicStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *NonPersistentTopicStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *NonPersistentTopicStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *NonPersistentTopicStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *NonPersistentTopicStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *NonPersistentTopicStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *NonPersistentTopicStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *NonPersistentTopicStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *NonPersistentTopicStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *NonPersistentTopicStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *NonPersistentTopicStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *NonPersistentTopicStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *NonPersistentTopicStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *NonPersistentTopicStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *NonPersistentTopicStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetPublishers

`func (o *NonPersistentTopicStats) GetPublishers() []NonPersistentPublisherStats`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *NonPersistentTopicStats) GetPublishersOk() (*[]NonPersistentPublisherStats, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *NonPersistentTopicStats) SetPublishers(v []NonPersistentPublisherStats)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *NonPersistentTopicStats) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### GetReplication

`func (o *NonPersistentTopicStats) GetReplication() map[string]NonPersistentReplicatorStats`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *NonPersistentTopicStats) GetReplicationOk() (*map[string]NonPersistentReplicatorStats, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *NonPersistentTopicStats) SetReplication(v map[string]NonPersistentReplicatorStats)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *NonPersistentTopicStats) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetStorageSize

`func (o *NonPersistentTopicStats) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *NonPersistentTopicStats) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *NonPersistentTopicStats) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *NonPersistentTopicStats) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetSubscriptions

`func (o *NonPersistentTopicStats) GetSubscriptions() map[string]NonPersistentSubscriptionStats`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *NonPersistentTopicStats) GetSubscriptionsOk() (*map[string]NonPersistentSubscriptionStats, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *NonPersistentTopicStats) SetSubscriptions(v map[string]NonPersistentSubscriptionStats)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *NonPersistentTopicStats) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


