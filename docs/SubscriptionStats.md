# SubscriptionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSubscriptionStats

`func NewSubscriptionStats() *SubscriptionStats`

NewSubscriptionStats instantiates a new SubscriptionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionStatsWithDefaults

`func NewSubscriptionStatsWithDefaults() *SubscriptionStats`

NewSubscriptionStatsWithDefaults instantiates a new SubscriptionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurable

`func (o *SubscriptionStats) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *SubscriptionStats) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *SubscriptionStats) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *SubscriptionStats) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetReplicated

`func (o *SubscriptionStats) GetReplicated() bool`

GetReplicated returns the Replicated field if non-nil, zero value otherwise.

### GetReplicatedOk

`func (o *SubscriptionStats) GetReplicatedOk() (*bool, bool)`

GetReplicatedOk returns a tuple with the Replicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicated

`func (o *SubscriptionStats) SetReplicated(v bool)`

SetReplicated sets Replicated field to given value.

### HasReplicated

`func (o *SubscriptionStats) HasReplicated() bool`

HasReplicated returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *SubscriptionStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *SubscriptionStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *SubscriptionStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *SubscriptionStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetBacklogSize

`func (o *SubscriptionStats) GetBacklogSize() int64`

GetBacklogSize returns the BacklogSize field if non-nil, zero value otherwise.

### GetBacklogSizeOk

`func (o *SubscriptionStats) GetBacklogSizeOk() (*int64, bool)`

GetBacklogSizeOk returns a tuple with the BacklogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklogSize

`func (o *SubscriptionStats) SetBacklogSize(v int64)`

SetBacklogSize sets BacklogSize field to given value.

### HasBacklogSize

`func (o *SubscriptionStats) HasBacklogSize() bool`

HasBacklogSize returns a boolean if a field has been set.

### GetBytesOutCounter

`func (o *SubscriptionStats) GetBytesOutCounter() int64`

GetBytesOutCounter returns the BytesOutCounter field if non-nil, zero value otherwise.

### GetBytesOutCounterOk

`func (o *SubscriptionStats) GetBytesOutCounterOk() (*int64, bool)`

GetBytesOutCounterOk returns a tuple with the BytesOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutCounter

`func (o *SubscriptionStats) SetBytesOutCounter(v int64)`

SetBytesOutCounter sets BytesOutCounter field to given value.

### HasBytesOutCounter

`func (o *SubscriptionStats) HasBytesOutCounter() bool`

HasBytesOutCounter returns a boolean if a field has been set.

### GetMsgOutCounter

`func (o *SubscriptionStats) GetMsgOutCounter() int64`

GetMsgOutCounter returns the MsgOutCounter field if non-nil, zero value otherwise.

### GetMsgOutCounterOk

`func (o *SubscriptionStats) GetMsgOutCounterOk() (*int64, bool)`

GetMsgOutCounterOk returns a tuple with the MsgOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgOutCounter

`func (o *SubscriptionStats) SetMsgOutCounter(v int64)`

SetMsgOutCounter sets MsgOutCounter field to given value.

### HasMsgOutCounter

`func (o *SubscriptionStats) HasMsgOutCounter() bool`

HasMsgOutCounter returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRanges

`func (o *SubscriptionStats) GetNonContiguousDeletedMessagesRanges() int32`

GetNonContiguousDeletedMessagesRanges returns the NonContiguousDeletedMessagesRanges field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesOk

`func (o *SubscriptionStats) GetNonContiguousDeletedMessagesRangesOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesOk returns a tuple with the NonContiguousDeletedMessagesRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRanges

`func (o *SubscriptionStats) SetNonContiguousDeletedMessagesRanges(v int32)`

SetNonContiguousDeletedMessagesRanges sets NonContiguousDeletedMessagesRanges field to given value.

### HasNonContiguousDeletedMessagesRanges

`func (o *SubscriptionStats) HasNonContiguousDeletedMessagesRanges() bool`

HasNonContiguousDeletedMessagesRanges returns a boolean if a field has been set.

### GetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *SubscriptionStats) GetNonContiguousDeletedMessagesRangesSerializedSize() int32`

GetNonContiguousDeletedMessagesRangesSerializedSize returns the NonContiguousDeletedMessagesRangesSerializedSize field if non-nil, zero value otherwise.

### GetNonContiguousDeletedMessagesRangesSerializedSizeOk

`func (o *SubscriptionStats) GetNonContiguousDeletedMessagesRangesSerializedSizeOk() (*int32, bool)`

GetNonContiguousDeletedMessagesRangesSerializedSizeOk returns a tuple with the NonContiguousDeletedMessagesRangesSerializedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonContiguousDeletedMessagesRangesSerializedSize

`func (o *SubscriptionStats) SetNonContiguousDeletedMessagesRangesSerializedSize(v int32)`

SetNonContiguousDeletedMessagesRangesSerializedSize sets NonContiguousDeletedMessagesRangesSerializedSize field to given value.

### HasNonContiguousDeletedMessagesRangesSerializedSize

`func (o *SubscriptionStats) HasNonContiguousDeletedMessagesRangesSerializedSize() bool`

HasNonContiguousDeletedMessagesRangesSerializedSize returns a boolean if a field has been set.

### GetConsumers

`func (o *SubscriptionStats) GetConsumers() []ConsumerStats`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *SubscriptionStats) GetConsumersOk() (*[]ConsumerStats, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *SubscriptionStats) SetConsumers(v []ConsumerStats)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *SubscriptionStats) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *SubscriptionStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *SubscriptionStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *SubscriptionStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *SubscriptionStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgRateRedeliver

`func (o *SubscriptionStats) GetMsgRateRedeliver() float64`

GetMsgRateRedeliver returns the MsgRateRedeliver field if non-nil, zero value otherwise.

### GetMsgRateRedeliverOk

`func (o *SubscriptionStats) GetMsgRateRedeliverOk() (*float64, bool)`

GetMsgRateRedeliverOk returns a tuple with the MsgRateRedeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateRedeliver

`func (o *SubscriptionStats) SetMsgRateRedeliver(v float64)`

SetMsgRateRedeliver sets MsgRateRedeliver field to given value.

### HasMsgRateRedeliver

`func (o *SubscriptionStats) HasMsgRateRedeliver() bool`

HasMsgRateRedeliver returns a boolean if a field has been set.

### GetChunkedMessageRate

`func (o *SubscriptionStats) GetChunkedMessageRate() int32`

GetChunkedMessageRate returns the ChunkedMessageRate field if non-nil, zero value otherwise.

### GetChunkedMessageRateOk

`func (o *SubscriptionStats) GetChunkedMessageRateOk() (*int32, bool)`

GetChunkedMessageRateOk returns a tuple with the ChunkedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkedMessageRate

`func (o *SubscriptionStats) SetChunkedMessageRate(v int32)`

SetChunkedMessageRate sets ChunkedMessageRate field to given value.

### HasChunkedMessageRate

`func (o *SubscriptionStats) HasChunkedMessageRate() bool`

HasChunkedMessageRate returns a boolean if a field has been set.

### GetMsgBacklog

`func (o *SubscriptionStats) GetMsgBacklog() int64`

GetMsgBacklog returns the MsgBacklog field if non-nil, zero value otherwise.

### GetMsgBacklogOk

`func (o *SubscriptionStats) GetMsgBacklogOk() (*int64, bool)`

GetMsgBacklogOk returns a tuple with the MsgBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgBacklog

`func (o *SubscriptionStats) SetMsgBacklog(v int64)`

SetMsgBacklog sets MsgBacklog field to given value.

### HasMsgBacklog

`func (o *SubscriptionStats) HasMsgBacklog() bool`

HasMsgBacklog returns a boolean if a field has been set.

### GetMsgBacklogNoDelayed

`func (o *SubscriptionStats) GetMsgBacklogNoDelayed() int64`

GetMsgBacklogNoDelayed returns the MsgBacklogNoDelayed field if non-nil, zero value otherwise.

### GetMsgBacklogNoDelayedOk

`func (o *SubscriptionStats) GetMsgBacklogNoDelayedOk() (*int64, bool)`

GetMsgBacklogNoDelayedOk returns a tuple with the MsgBacklogNoDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgBacklogNoDelayed

`func (o *SubscriptionStats) SetMsgBacklogNoDelayed(v int64)`

SetMsgBacklogNoDelayed sets MsgBacklogNoDelayed field to given value.

### HasMsgBacklogNoDelayed

`func (o *SubscriptionStats) HasMsgBacklogNoDelayed() bool`

HasMsgBacklogNoDelayed returns a boolean if a field has been set.

### GetBlockedSubscriptionOnUnackedMsgs

`func (o *SubscriptionStats) GetBlockedSubscriptionOnUnackedMsgs() bool`

GetBlockedSubscriptionOnUnackedMsgs returns the BlockedSubscriptionOnUnackedMsgs field if non-nil, zero value otherwise.

### GetBlockedSubscriptionOnUnackedMsgsOk

`func (o *SubscriptionStats) GetBlockedSubscriptionOnUnackedMsgsOk() (*bool, bool)`

GetBlockedSubscriptionOnUnackedMsgsOk returns a tuple with the BlockedSubscriptionOnUnackedMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedSubscriptionOnUnackedMsgs

`func (o *SubscriptionStats) SetBlockedSubscriptionOnUnackedMsgs(v bool)`

SetBlockedSubscriptionOnUnackedMsgs sets BlockedSubscriptionOnUnackedMsgs field to given value.

### HasBlockedSubscriptionOnUnackedMsgs

`func (o *SubscriptionStats) HasBlockedSubscriptionOnUnackedMsgs() bool`

HasBlockedSubscriptionOnUnackedMsgs returns a boolean if a field has been set.

### GetMsgDelayed

`func (o *SubscriptionStats) GetMsgDelayed() int64`

GetMsgDelayed returns the MsgDelayed field if non-nil, zero value otherwise.

### GetMsgDelayedOk

`func (o *SubscriptionStats) GetMsgDelayedOk() (*int64, bool)`

GetMsgDelayedOk returns a tuple with the MsgDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDelayed

`func (o *SubscriptionStats) SetMsgDelayed(v int64)`

SetMsgDelayed sets MsgDelayed field to given value.

### HasMsgDelayed

`func (o *SubscriptionStats) HasMsgDelayed() bool`

HasMsgDelayed returns a boolean if a field has been set.

### GetUnackedMessages

`func (o *SubscriptionStats) GetUnackedMessages() int64`

GetUnackedMessages returns the UnackedMessages field if non-nil, zero value otherwise.

### GetUnackedMessagesOk

`func (o *SubscriptionStats) GetUnackedMessagesOk() (*int64, bool)`

GetUnackedMessagesOk returns a tuple with the UnackedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnackedMessages

`func (o *SubscriptionStats) SetUnackedMessages(v int64)`

SetUnackedMessages sets UnackedMessages field to given value.

### HasUnackedMessages

`func (o *SubscriptionStats) HasUnackedMessages() bool`

HasUnackedMessages returns a boolean if a field has been set.

### GetActiveConsumerName

`func (o *SubscriptionStats) GetActiveConsumerName() string`

GetActiveConsumerName returns the ActiveConsumerName field if non-nil, zero value otherwise.

### GetActiveConsumerNameOk

`func (o *SubscriptionStats) GetActiveConsumerNameOk() (*string, bool)`

GetActiveConsumerNameOk returns a tuple with the ActiveConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConsumerName

`func (o *SubscriptionStats) SetActiveConsumerName(v string)`

SetActiveConsumerName sets ActiveConsumerName field to given value.

### HasActiveConsumerName

`func (o *SubscriptionStats) HasActiveConsumerName() bool`

HasActiveConsumerName returns a boolean if a field has been set.

### GetMsgRateExpired

`func (o *SubscriptionStats) GetMsgRateExpired() float64`

GetMsgRateExpired returns the MsgRateExpired field if non-nil, zero value otherwise.

### GetMsgRateExpiredOk

`func (o *SubscriptionStats) GetMsgRateExpiredOk() (*float64, bool)`

GetMsgRateExpiredOk returns a tuple with the MsgRateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateExpired

`func (o *SubscriptionStats) SetMsgRateExpired(v float64)`

SetMsgRateExpired sets MsgRateExpired field to given value.

### HasMsgRateExpired

`func (o *SubscriptionStats) HasMsgRateExpired() bool`

HasMsgRateExpired returns a boolean if a field has been set.

### GetTotalMsgExpired

`func (o *SubscriptionStats) GetTotalMsgExpired() int64`

GetTotalMsgExpired returns the TotalMsgExpired field if non-nil, zero value otherwise.

### GetTotalMsgExpiredOk

`func (o *SubscriptionStats) GetTotalMsgExpiredOk() (*int64, bool)`

GetTotalMsgExpiredOk returns a tuple with the TotalMsgExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMsgExpired

`func (o *SubscriptionStats) SetTotalMsgExpired(v int64)`

SetTotalMsgExpired sets TotalMsgExpired field to given value.

### HasTotalMsgExpired

`func (o *SubscriptionStats) HasTotalMsgExpired() bool`

HasTotalMsgExpired returns a boolean if a field has been set.

### GetLastExpireTimestamp

`func (o *SubscriptionStats) GetLastExpireTimestamp() int64`

GetLastExpireTimestamp returns the LastExpireTimestamp field if non-nil, zero value otherwise.

### GetLastExpireTimestampOk

`func (o *SubscriptionStats) GetLastExpireTimestampOk() (*int64, bool)`

GetLastExpireTimestampOk returns a tuple with the LastExpireTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExpireTimestamp

`func (o *SubscriptionStats) SetLastExpireTimestamp(v int64)`

SetLastExpireTimestamp sets LastExpireTimestamp field to given value.

### HasLastExpireTimestamp

`func (o *SubscriptionStats) HasLastExpireTimestamp() bool`

HasLastExpireTimestamp returns a boolean if a field has been set.

### GetLastConsumedFlowTimestamp

`func (o *SubscriptionStats) GetLastConsumedFlowTimestamp() int64`

GetLastConsumedFlowTimestamp returns the LastConsumedFlowTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedFlowTimestampOk

`func (o *SubscriptionStats) GetLastConsumedFlowTimestampOk() (*int64, bool)`

GetLastConsumedFlowTimestampOk returns a tuple with the LastConsumedFlowTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedFlowTimestamp

`func (o *SubscriptionStats) SetLastConsumedFlowTimestamp(v int64)`

SetLastConsumedFlowTimestamp sets LastConsumedFlowTimestamp field to given value.

### HasLastConsumedFlowTimestamp

`func (o *SubscriptionStats) HasLastConsumedFlowTimestamp() bool`

HasLastConsumedFlowTimestamp returns a boolean if a field has been set.

### GetLastConsumedTimestamp

`func (o *SubscriptionStats) GetLastConsumedTimestamp() int64`

GetLastConsumedTimestamp returns the LastConsumedTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedTimestampOk

`func (o *SubscriptionStats) GetLastConsumedTimestampOk() (*int64, bool)`

GetLastConsumedTimestampOk returns a tuple with the LastConsumedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTimestamp

`func (o *SubscriptionStats) SetLastConsumedTimestamp(v int64)`

SetLastConsumedTimestamp sets LastConsumedTimestamp field to given value.

### HasLastConsumedTimestamp

`func (o *SubscriptionStats) HasLastConsumedTimestamp() bool`

HasLastConsumedTimestamp returns a boolean if a field has been set.

### GetLastAckedTimestamp

`func (o *SubscriptionStats) GetLastAckedTimestamp() int64`

GetLastAckedTimestamp returns the LastAckedTimestamp field if non-nil, zero value otherwise.

### GetLastAckedTimestampOk

`func (o *SubscriptionStats) GetLastAckedTimestampOk() (*int64, bool)`

GetLastAckedTimestampOk returns a tuple with the LastAckedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAckedTimestamp

`func (o *SubscriptionStats) SetLastAckedTimestamp(v int64)`

SetLastAckedTimestamp sets LastAckedTimestamp field to given value.

### HasLastAckedTimestamp

`func (o *SubscriptionStats) HasLastAckedTimestamp() bool`

HasLastAckedTimestamp returns a boolean if a field has been set.

### GetLastMarkDeleteAdvancedTimestamp

`func (o *SubscriptionStats) GetLastMarkDeleteAdvancedTimestamp() int64`

GetLastMarkDeleteAdvancedTimestamp returns the LastMarkDeleteAdvancedTimestamp field if non-nil, zero value otherwise.

### GetLastMarkDeleteAdvancedTimestampOk

`func (o *SubscriptionStats) GetLastMarkDeleteAdvancedTimestampOk() (*int64, bool)`

GetLastMarkDeleteAdvancedTimestampOk returns a tuple with the LastMarkDeleteAdvancedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMarkDeleteAdvancedTimestamp

`func (o *SubscriptionStats) SetLastMarkDeleteAdvancedTimestamp(v int64)`

SetLastMarkDeleteAdvancedTimestamp sets LastMarkDeleteAdvancedTimestamp field to given value.

### HasLastMarkDeleteAdvancedTimestamp

`func (o *SubscriptionStats) HasLastMarkDeleteAdvancedTimestamp() bool`

HasLastMarkDeleteAdvancedTimestamp returns a boolean if a field has been set.

### GetConsumersAfterMarkDeletePosition

`func (o *SubscriptionStats) GetConsumersAfterMarkDeletePosition() map[string]string`

GetConsumersAfterMarkDeletePosition returns the ConsumersAfterMarkDeletePosition field if non-nil, zero value otherwise.

### GetConsumersAfterMarkDeletePositionOk

`func (o *SubscriptionStats) GetConsumersAfterMarkDeletePositionOk() (*map[string]string, bool)`

GetConsumersAfterMarkDeletePositionOk returns a tuple with the ConsumersAfterMarkDeletePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumersAfterMarkDeletePosition

`func (o *SubscriptionStats) SetConsumersAfterMarkDeletePosition(v map[string]string)`

SetConsumersAfterMarkDeletePosition sets ConsumersAfterMarkDeletePosition field to given value.

### HasConsumersAfterMarkDeletePosition

`func (o *SubscriptionStats) HasConsumersAfterMarkDeletePosition() bool`

HasConsumersAfterMarkDeletePosition returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionStats) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


