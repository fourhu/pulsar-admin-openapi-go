# NonPersistentSubscriptionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgDropRate** | Pointer to **float64** |  | [optional] 
**Durable** | Pointer to **bool** |  | [optional] 
**Replicated** | Pointer to **bool** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**BacklogSize** | Pointer to **int64** |  | [optional] 
**BytesOutCounter** | Pointer to **int64** |  | [optional] 
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**NonContiguousDeletedMessagesRanges** | Pointer to **int32** |  | [optional] 
**NonContiguousDeletedMessagesRangesSerializedSize** | Pointer to **int32** |  | [optional] 
**Consumers** | Pointer to [**[]ConsumerStats**](ConsumerStats.md) |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgRateRedeliver** | Pointer to **float64** |  | [optional] 
**ChunkedMessageRate** | Pointer to **int32** |  | [optional] 
**MsgBacklog** | Pointer to **int64** |  | [optional] 
**MsgBacklogNoDelayed** | Pointer to **int64** |  | [optional] 
**BlockedSubscriptionOnUnackedMsgs** | Pointer to **bool** |  | [optional] 
**MsgDelayed** | Pointer to **int64** |  | [optional] 
**UnackedMessages** | Pointer to **int64** |  | [optional] 
**ActiveConsumerName** | Pointer to **string** |  | [optional] 
**MsgRateExpired** | Pointer to **float64** |  | [optional] 
**TotalMsgExpired** | Pointer to **int64** |  | [optional] 
**LastExpireTimestamp** | Pointer to **int64** |  | [optional] 
**LastConsumedFlowTimestamp** | Pointer to **int64** |  | [optional] 
**LastConsumedTimestamp** | Pointer to **int64** |  | [optional] 
**LastAckedTimestamp** | Pointer to **int64** |  | [optional] 
**LastMarkDeleteAdvancedTimestamp** | Pointer to **int64** |  | [optional] 
**ConsumersAfterMarkDeletePosition** | Pointer to **map[string]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNonPersistentSubscriptionStats

`func NewNonPersistentSubscriptionStats() *NonPersistentSubscriptionStats`

NewNonPersistentSubscriptionStats instantiates a new NonPersistentSubscriptionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPersistentSubscriptionStatsWithDefaults

`func NewNonPersistentSubscriptionStatsWithDefaults() *NonPersistentSubscriptionStats`

NewNonPersistentSubscriptionStatsWithDefaults instantiates a new NonPersistentSubscriptionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgDropRate

`func (o *NonPersistentSubscriptionStats) GetMsgDropRate() float64`

GetMsgDropRate returns the MsgDropRate field if non-nil, zero value otherwise.

### GetMsgDropRateOk

`func (o *NonPersistentSubscriptionStats) GetMsgDropRateOk() (*float64, bool)`

GetMsgDropRateOk returns a tuple with the MsgDropRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDropRate

`func (o *NonPersistentSubscriptionStats) SetMsgDropRate(v float64)`

SetMsgDropRate sets MsgDropRate field to given value.

### HasMsgDropRate

`func (o *NonPersistentSubscriptionStats) HasMsgDropRate() bool`

HasMsgDropRate returns a boolean if a field has been set.

### GetDurable

`func (o *NonPersistentSubscriptionStats) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *NonPersistentSubscriptionStats) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *NonPersistentSubscriptionStats) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *NonPersistentSubscriptionStats) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetReplicated

`func (o *NonPersistentSubscriptionStats) GetReplicated() bool`

GetReplicated returns the Replicated field if non-nil, zero value otherwise.

### GetReplicatedOk

`func (o *NonPersistentSubscriptionStats) GetReplicatedOk() (*bool, bool)`

GetReplicatedOk returns a tuple with the Replicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicated

`func (o *NonPersistentSubscriptionStats) SetReplicated(v bool)`

SetReplicated sets Replicated field to given value.

### HasReplicated

`func (o *NonPersistentSubscriptionStats) HasReplicated() bool`

HasReplicated returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *NonPersistentSubscriptionStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *NonPersistentSubscriptionStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *NonPersistentSubscriptionStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *NonPersistentSubscriptionStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetBacklogSize

`func (o *NonPersistentSubscriptionStats) GetBacklogSize() int64`

GetBacklogSize returns the BacklogSize field if non-nil, zero value otherwise.

### GetBacklogSizeOk

`func (o *NonPersistentSubscriptionStats) GetBacklogSizeOk() (*int64, bool)`

GetBacklogSizeOk returns a tuple with the BacklogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogSize

`func (o *NonPersistentSubscriptionStats) SetBacklogSize(v int64)`

SetBacklogSize sets BacklogSize field to given value.

### HasBacklogSize

`func (o *NonPersistentSubscriptionStats) HasBacklogSize() bool`

HasBacklogSize returns a boolean if a field has been set.

### GetBytesOutCounter

`func (o *NonPersistentSubscriptionStats) GetBytesOutCounter() int64`

GetBytesOutCounter returns the BytesOutCounter field if non-nil, zero value otherwise.

### GetBytesOutCounterOk

`func (o *NonPersistentSubscriptionStats) GetBytesOutCounterOk() (*int64, bool)`

GetBytesOutCounterOk returns a tuple with the BytesOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutCounter

`func (o *NonPersistentSubscriptionStats) SetBytesOutCounter(v int64)`

SetBytesOutCounter sets BytesOutCounter field to given value.

### HasBytesOutCounter

`func (o *NonPersistentSubscriptionStats) HasBytesOutCounter() bool`

HasBytesOutCounter returns a boolean if a field has been set.

### GetMsgOutCounter

`func (o *NonPersistentSubscriptionStats) GetMsgOutCounter() int64`

GetMsgOutCounter returns the MsgOutCounter field if non-nil, zero value otherwise.

### GetMsgOutCounterOk

`func (o *NonPersistentSubscriptionStats) GetMsgOutCounterOk() (*int64, bool)`

GetMsgOutCounterOk returns a tuple with the MsgOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgOutCounter

`func (o *NonPersistentSubscriptionStats) SetMsgOutCounter(v int64)`

SetMsgOutCounter sets MsgOutCounter field to given value.

### HasMsgOutCounter

`func (o *NonPersistentSubscriptionStats) HasMsgOutCounter() bool`

HasMsgOutCounter returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRanges

`func (o *NonPersistentSubscriptionStats) GetNonContiguousDeletedMessagesRanges() int32`

GetNonContiguousDeletedMessagesRanges returns the NonContiguousDeletedMessagesRanges field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesOk

`func (o *NonPersistentSubscriptionStats) GetNonContiguousDeletedMessagesRangesOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesOk returns a tuple with the NonContiguousDeletedMessagesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRanges

`func (o *NonPersistentSubscriptionStats) SetNonContiguousDeletedMessagesRanges(v int32)`

SetNonContiguousDeletedMessagesRanges sets NonContiguousDeletedMessagesRanges field to given value.

### HasNonContiguousDeletedMessagesRanges

`func (o *NonPersistentSubscriptionStats) HasNonContiguousDeletedMessagesRanges() bool`

HasNonContiguousDeletedMessagesRanges returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *NonPersistentSubscriptionStats) GetNonContiguousDeletedMessagesRangesSerializedSize() int32`

GetNonContiguousDeletedMessagesRangesSerializedSize returns the NonContiguousDeletedMessagesRangesSerializedSize field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesSerializedSizeOk

`func (o *NonPersistentSubscriptionStats) GetNonContiguousDeletedMessagesRangesSerializedSizeOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesSerializedSizeOk returns a tuple with the NonContiguousDeletedMessagesRangesSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *NonPersistentSubscriptionStats) SetNonContiguousDeletedMessagesRangesSerializedSize(v int32)`

SetNonContiguousDeletedMessagesRangesSerializedSize sets NonContiguousDeletedMessagesRangesSerializedSize field to given value.

### HasNonContiguousDeletedMessagesRangesSerializedSize

`func (o *NonPersistentSubscriptionStats) HasNonContiguousDeletedMessagesRangesSerializedSize() bool`

HasNonContiguousDeletedMessagesRangesSerializedSize returns a boolean if a field has been set.

### GetConsumers

`func (o *NonPersistentSubscriptionStats) GetConsumers() []ConsumerStats`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *NonPersistentSubscriptionStats) GetConsumersOk() (*[]ConsumerStats, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *NonPersistentSubscriptionStats) SetConsumers(v []ConsumerStats)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *NonPersistentSubscriptionStats) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *NonPersistentSubscriptionStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *NonPersistentSubscriptionStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *NonPersistentSubscriptionStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *NonPersistentSubscriptionStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgRateRedeliver

`func (o *NonPersistentSubscriptionStats) GetMsgRateRedeliver() float64`

GetMsgRateRedeliver returns the MsgRateRedeliver field if non-nil, zero value otherwise.

### GetMsgRateRedeliverOk

`func (o *NonPersistentSubscriptionStats) GetMsgRateRedeliverOk() (*float64, bool)`

GetMsgRateRedeliverOk returns a tuple with the MsgRateRedeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateRedeliver

`func (o *NonPersistentSubscriptionStats) SetMsgRateRedeliver(v float64)`

SetMsgRateRedeliver sets MsgRateRedeliver field to given value.

### HasMsgRateRedeliver

`func (o *NonPersistentSubscriptionStats) HasMsgRateRedeliver() bool`

HasMsgRateRedeliver returns a boolean if a field has been set.

### GetChunkedMessageRate

`func (o *NonPersistentSubscriptionStats) GetChunkedMessageRate() int32`

GetChunkedMessageRate returns the ChunkedMessageRate field if non-nil, zero value otherwise.

### GetChunkedMessageRateOk

`func (o *NonPersistentSubscriptionStats) GetChunkedMessageRateOk() (*int32, bool)`

GetChunkedMessageRateOk returns a tuple with the ChunkedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkedMessageRate

`func (o *NonPersistentSubscriptionStats) SetChunkedMessageRate(v int32)`

SetChunkedMessageRate sets ChunkedMessageRate field to given value.

### HasChunkedMessageRate

`func (o *NonPersistentSubscriptionStats) HasChunkedMessageRate() bool`

HasChunkedMessageRate returns a boolean if a field has been set.

### GetMsgBacklog

`func (o *NonPersistentSubscriptionStats) GetMsgBacklog() int64`

GetMsgBacklog returns the MsgBacklog field if non-nil, zero value otherwise.

### GetMsgBacklogOk

`func (o *NonPersistentSubscriptionStats) GetMsgBacklogOk() (*int64, bool)`

GetMsgBacklogOk returns a tuple with the MsgBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgBacklog

`func (o *NonPersistentSubscriptionStats) SetMsgBacklog(v int64)`

SetMsgBacklog sets MsgBacklog field to given value.

### HasMsgBacklog

`func (o *NonPersistentSubscriptionStats) HasMsgBacklog() bool`

HasMsgBacklog returns a boolean if a field has been set.

### GetMsgBacklogNoDelayed

`func (o *NonPersistentSubscriptionStats) GetMsgBacklogNoDelayed() int64`

GetMsgBacklogNoDelayed returns the MsgBacklogNoDelayed field if non-nil, zero value otherwise.

### GetMsgBacklogNoDelayedOk

`func (o *NonPersistentSubscriptionStats) GetMsgBacklogNoDelayedOk() (*int64, bool)`

GetMsgBacklogNoDelayedOk returns a tuple with the MsgBacklogNoDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgBacklogNoDelayed

`func (o *NonPersistentSubscriptionStats) SetMsgBacklogNoDelayed(v int64)`

SetMsgBacklogNoDelayed sets MsgBacklogNoDelayed field to given value.

### HasMsgBacklogNoDelayed

`func (o *NonPersistentSubscriptionStats) HasMsgBacklogNoDelayed() bool`

HasMsgBacklogNoDelayed returns a boolean if a field has been set.

### GetBlockedSubscriptionOnUnackedMsgs

`func (o *NonPersistentSubscriptionStats) GetBlockedSubscriptionOnUnackedMsgs() bool`

GetBlockedSubscriptionOnUnackedMsgs returns the BlockedSubscriptionOnUnackedMsgs field if non-nil, zero value otherwise.

### GetBlockedSubscriptionOnUnackedMsgsOk

`func (o *NonPersistentSubscriptionStats) GetBlockedSubscriptionOnUnackedMsgsOk() (*bool, bool)`

GetBlockedSubscriptionOnUnackedMsgsOk returns a tuple with the BlockedSubscriptionOnUnackedMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedSubscriptionOnUnackedMsgs

`func (o *NonPersistentSubscriptionStats) SetBlockedSubscriptionOnUnackedMsgs(v bool)`

SetBlockedSubscriptionOnUnackedMsgs sets BlockedSubscriptionOnUnackedMsgs field to given value.

### HasBlockedSubscriptionOnUnackedMsgs

`func (o *NonPersistentSubscriptionStats) HasBlockedSubscriptionOnUnackedMsgs() bool`

HasBlockedSubscriptionOnUnackedMsgs returns a boolean if a field has been set.

### GetMsgDelayed

`func (o *NonPersistentSubscriptionStats) GetMsgDelayed() int64`

GetMsgDelayed returns the MsgDelayed field if non-nil, zero value otherwise.

### GetMsgDelayedOk

`func (o *NonPersistentSubscriptionStats) GetMsgDelayedOk() (*int64, bool)`

GetMsgDelayedOk returns a tuple with the MsgDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDelayed

`func (o *NonPersistentSubscriptionStats) SetMsgDelayed(v int64)`

SetMsgDelayed sets MsgDelayed field to given value.

### HasMsgDelayed

`func (o *NonPersistentSubscriptionStats) HasMsgDelayed() bool`

HasMsgDelayed returns a boolean if a field has been set.

### GetUnackedMessages

`func (o *NonPersistentSubscriptionStats) GetUnackedMessages() int64`

GetUnackedMessages returns the UnackedMessages field if non-nil, zero value otherwise.

### GetUnackedMessagesOk

`func (o *NonPersistentSubscriptionStats) GetUnackedMessagesOk() (*int64, bool)`

GetUnackedMessagesOk returns a tuple with the UnackedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnackedMessages

`func (o *NonPersistentSubscriptionStats) SetUnackedMessages(v int64)`

SetUnackedMessages sets UnackedMessages field to given value.

### HasUnackedMessages

`func (o *NonPersistentSubscriptionStats) HasUnackedMessages() bool`

HasUnackedMessages returns a boolean if a field has been set.

### GetActiveConsumerName

`func (o *NonPersistentSubscriptionStats) GetActiveConsumerName() string`

GetActiveConsumerName returns the ActiveConsumerName field if non-nil, zero value otherwise.

### GetActiveConsumerNameOk

`func (o *NonPersistentSubscriptionStats) GetActiveConsumerNameOk() (*string, bool)`

GetActiveConsumerNameOk returns a tuple with the ActiveConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConsumerName

`func (o *NonPersistentSubscriptionStats) SetActiveConsumerName(v string)`

SetActiveConsumerName sets ActiveConsumerName field to given value.

### HasActiveConsumerName

`func (o *NonPersistentSubscriptionStats) HasActiveConsumerName() bool`

HasActiveConsumerName returns a boolean if a field has been set.

### GetMsgRateExpired

`func (o *NonPersistentSubscriptionStats) GetMsgRateExpired() float64`

GetMsgRateExpired returns the MsgRateExpired field if non-nil, zero value otherwise.

### GetMsgRateExpiredOk

`func (o *NonPersistentSubscriptionStats) GetMsgRateExpiredOk() (*float64, bool)`

GetMsgRateExpiredOk returns a tuple with the MsgRateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateExpired

`func (o *NonPersistentSubscriptionStats) SetMsgRateExpired(v float64)`

SetMsgRateExpired sets MsgRateExpired field to given value.

### HasMsgRateExpired

`func (o *NonPersistentSubscriptionStats) HasMsgRateExpired() bool`

HasMsgRateExpired returns a boolean if a field has been set.

### GetTotalMsgExpired

`func (o *NonPersistentSubscriptionStats) GetTotalMsgExpired() int64`

GetTotalMsgExpired returns the TotalMsgExpired field if non-nil, zero value otherwise.

### GetTotalMsgExpiredOk

`func (o *NonPersistentSubscriptionStats) GetTotalMsgExpiredOk() (*int64, bool)`

GetTotalMsgExpiredOk returns a tuple with the TotalMsgExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMsgExpired

`func (o *NonPersistentSubscriptionStats) SetTotalMsgExpired(v int64)`

SetTotalMsgExpired sets TotalMsgExpired field to given value.

### HasTotalMsgExpired

`func (o *NonPersistentSubscriptionStats) HasTotalMsgExpired() bool`

HasTotalMsgExpired returns a boolean if a field has been set.

### GetLastExpireTimestamp

`func (o *NonPersistentSubscriptionStats) GetLastExpireTimestamp() int64`

GetLastExpireTimestamp returns the LastExpireTimestamp field if non-nil, zero value otherwise.

### GetLastExpireTimestampOk

`func (o *NonPersistentSubscriptionStats) GetLastExpireTimestampOk() (*int64, bool)`

GetLastExpireTimestampOk returns a tuple with the LastExpireTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExpireTimestamp

`func (o *NonPersistentSubscriptionStats) SetLastExpireTimestamp(v int64)`

SetLastExpireTimestamp sets LastExpireTimestamp field to given value.

### HasLastExpireTimestamp

`func (o *NonPersistentSubscriptionStats) HasLastExpireTimestamp() bool`

HasLastExpireTimestamp returns a boolean if a field has been set.

### GetLastConsumedFlowTimestamp

`func (o *NonPersistentSubscriptionStats) GetLastConsumedFlowTimestamp() int64`

GetLastConsumedFlowTimestamp returns the LastConsumedFlowTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedFlowTimestampOk

`func (o *NonPersistentSubscriptionStats) GetLastConsumedFlowTimestampOk() (*int64, bool)`

GetLastConsumedFlowTimestampOk returns a tuple with the LastConsumedFlowTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedFlowTimestamp

`func (o *NonPersistentSubscriptionStats) SetLastConsumedFlowTimestamp(v int64)`

SetLastConsumedFlowTimestamp sets LastConsumedFlowTimestamp field to given value.

### HasLastConsumedFlowTimestamp

`func (o *NonPersistentSubscriptionStats) HasLastConsumedFlowTimestamp() bool`

HasLastConsumedFlowTimestamp returns a boolean if a field has been set.

### GetLastConsumedTimestamp

`func (o *NonPersistentSubscriptionStats) GetLastConsumedTimestamp() int64`

GetLastConsumedTimestamp returns the LastConsumedTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedTimestampOk

`func (o *NonPersistentSubscriptionStats) GetLastConsumedTimestampOk() (*int64, bool)`

GetLastConsumedTimestampOk returns a tuple with the LastConsumedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTimestamp

`func (o *NonPersistentSubscriptionStats) SetLastConsumedTimestamp(v int64)`

SetLastConsumedTimestamp sets LastConsumedTimestamp field to given value.

### HasLastConsumedTimestamp

`func (o *NonPersistentSubscriptionStats) HasLastConsumedTimestamp() bool`

HasLastConsumedTimestamp returns a boolean if a field has been set.

### GetLastAckedTimestamp

`func (o *NonPersistentSubscriptionStats) GetLastAckedTimestamp() int64`

GetLastAckedTimestamp returns the LastAckedTimestamp field if non-nil, zero value otherwise.

### GetLastAckedTimestampOk

`func (o *NonPersistentSubscriptionStats) GetLastAckedTimestampOk() (*int64, bool)`

GetLastAckedTimestampOk returns a tuple with the LastAckedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAckedTimestamp

`func (o *NonPersistentSubscriptionStats) SetLastAckedTimestamp(v int64)`

SetLastAckedTimestamp sets LastAckedTimestamp field to given value.

### HasLastAckedTimestamp

`func (o *NonPersistentSubscriptionStats) HasLastAckedTimestamp() bool`

HasLastAckedTimestamp returns a boolean if a field has been set.

### GetLastMarkDeleteAdvancedTimestamp

`func (o *NonPersistentSubscriptionStats) GetLastMarkDeleteAdvancedTimestamp() int64`

GetLastMarkDeleteAdvancedTimestamp returns the LastMarkDeleteAdvancedTimestamp field if non-nil, zero value otherwise.

### GetLastMarkDeleteAdvancedTimestampOk

`func (o *NonPersistentSubscriptionStats) GetLastMarkDeleteAdvancedTimestampOk() (*int64, bool)`

GetLastMarkDeleteAdvancedTimestampOk returns a tuple with the LastMarkDeleteAdvancedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMarkDeleteAdvancedTimestamp

`func (o *NonPersistentSubscriptionStats) SetLastMarkDeleteAdvancedTimestamp(v int64)`

SetLastMarkDeleteAdvancedTimestamp sets LastMarkDeleteAdvancedTimestamp field to given value.

### HasLastMarkDeleteAdvancedTimestamp

`func (o *NonPersistentSubscriptionStats) HasLastMarkDeleteAdvancedTimestamp() bool`

HasLastMarkDeleteAdvancedTimestamp returns a boolean if a field has been set.

### GetConsumersAfterMarkDeletePosition

`func (o *NonPersistentSubscriptionStats) GetConsumersAfterMarkDeletePosition() map[string]string`

GetConsumersAfterMarkDeletePosition returns the ConsumersAfterMarkDeletePosition field if non-nil, zero value otherwise.

### GetConsumersAfterMarkDeletePositionOk

`func (o *NonPersistentSubscriptionStats) GetConsumersAfterMarkDeletePositionOk() (*map[string]string, bool)`

GetConsumersAfterMarkDeletePositionOk returns a tuple with the ConsumersAfterMarkDeletePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumersAfterMarkDeletePosition

`func (o *NonPersistentSubscriptionStats) SetConsumersAfterMarkDeletePosition(v map[string]string)`

SetConsumersAfterMarkDeletePosition sets ConsumersAfterMarkDeletePosition field to given value.

### HasConsumersAfterMarkDeletePosition

`func (o *NonPersistentSubscriptionStats) HasConsumersAfterMarkDeletePosition() bool`

HasConsumersAfterMarkDeletePosition returns a boolean if a field has been set.

### GetType

`func (o *NonPersistentSubscriptionStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NonPersistentSubscriptionStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NonPersistentSubscriptionStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NonPersistentSubscriptionStats) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


