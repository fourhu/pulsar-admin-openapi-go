# NonPersistentSubscriptionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConsumerName** | Pointer to **string** |  | [optional] 
**BlockedSubscriptionOnUnackedMsgs** | Pointer to **bool** |  | [optional] 
**BytesOutCounter** | Pointer to **int64** |  | [optional] 
**ChuckedMessageRate** | Pointer to **int32** |  | [optional] 
**Consumers** | Pointer to [**[]ConsumerStats**](ConsumerStats.md) |  | [optional] 
**ConsumersAfterMarkDeletePosition** | Pointer to **map[string]string** |  | [optional] 
**IsDurable** | Pointer to **bool** |  | [optional] 
**IsReplicated** | Pointer to **bool** |  | [optional] 
**LastAckedTimestamp** | Pointer to **int64** |  | [optional] 
**LastConsumedFlowTimestamp** | Pointer to **int64** |  | [optional] 
**LastConsumedTimestamp** | Pointer to **int64** |  | [optional] 
**LastExpireTimestamp** | Pointer to **int64** |  | [optional] 
**MsgBacklog** | Pointer to **int64** |  | [optional] 
**MsgBacklogNoDelayed** | Pointer to **int64** |  | [optional] 
**MsgDelayed** | Pointer to **int64** |  | [optional] 
**MsgDropRate** | Pointer to **float64** |  | [optional] 
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**MsgRateExpired** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgRateRedeliver** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnackedMessages** | Pointer to **int64** |  | [optional] 

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

### GetChuckedMessageRate

`func (o *NonPersistentSubscriptionStats) GetChuckedMessageRate() int32`

GetChuckedMessageRate returns the ChuckedMessageRate field if non-nil, zero value otherwise.

### GetChuckedMessageRateOk

`func (o *NonPersistentSubscriptionStats) GetChuckedMessageRateOk() (*int32, bool)`

GetChuckedMessageRateOk returns a tuple with the ChuckedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChuckedMessageRate

`func (o *NonPersistentSubscriptionStats) SetChuckedMessageRate(v int32)`

SetChuckedMessageRate sets ChuckedMessageRate field to given value.

### HasChuckedMessageRate

`func (o *NonPersistentSubscriptionStats) HasChuckedMessageRate() bool`

HasChuckedMessageRate returns a boolean if a field has been set.

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

### GetIsDurable

`func (o *NonPersistentSubscriptionStats) GetIsDurable() bool`

GetIsDurable returns the IsDurable field if non-nil, zero value otherwise.

### GetIsDurableOk

`func (o *NonPersistentSubscriptionStats) GetIsDurableOk() (*bool, bool)`

GetIsDurableOk returns a tuple with the IsDurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDurable

`func (o *NonPersistentSubscriptionStats) SetIsDurable(v bool)`

SetIsDurable sets IsDurable field to given value.

### HasIsDurable

`func (o *NonPersistentSubscriptionStats) HasIsDurable() bool`

HasIsDurable returns a boolean if a field has been set.

### GetIsReplicated

`func (o *NonPersistentSubscriptionStats) GetIsReplicated() bool`

GetIsReplicated returns the IsReplicated field if non-nil, zero value otherwise.

### GetIsReplicatedOk

`func (o *NonPersistentSubscriptionStats) GetIsReplicatedOk() (*bool, bool)`

GetIsReplicatedOk returns a tuple with the IsReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicated

`func (o *NonPersistentSubscriptionStats) SetIsReplicated(v bool)`

SetIsReplicated sets IsReplicated field to given value.

### HasIsReplicated

`func (o *NonPersistentSubscriptionStats) HasIsReplicated() bool`

HasIsReplicated returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


