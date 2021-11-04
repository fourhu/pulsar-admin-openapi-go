# SubscriptionStats

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
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**MsgRateExpired** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgRateRedeliver** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnackedMessages** | Pointer to **int64** |  | [optional] 

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

### GetChuckedMessageRate

`func (o *SubscriptionStats) GetChuckedMessageRate() int32`

GetChuckedMessageRate returns the ChuckedMessageRate field if non-nil, zero value otherwise.

### GetChuckedMessageRateOk

`func (o *SubscriptionStats) GetChuckedMessageRateOk() (*int32, bool)`

GetChuckedMessageRateOk returns a tuple with the ChuckedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChuckedMessageRate

`func (o *SubscriptionStats) SetChuckedMessageRate(v int32)`

SetChuckedMessageRate sets ChuckedMessageRate field to given value.

### HasChuckedMessageRate

`func (o *SubscriptionStats) HasChuckedMessageRate() bool`

HasChuckedMessageRate returns a boolean if a field has been set.

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

### GetIsDurable

`func (o *SubscriptionStats) GetIsDurable() bool`

GetIsDurable returns the IsDurable field if non-nil, zero value otherwise.

### GetIsDurableOk

`func (o *SubscriptionStats) GetIsDurableOk() (*bool, bool)`

GetIsDurableOk returns a tuple with the IsDurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDurable

`func (o *SubscriptionStats) SetIsDurable(v bool)`

SetIsDurable sets IsDurable field to given value.

### HasIsDurable

`func (o *SubscriptionStats) HasIsDurable() bool`

HasIsDurable returns a boolean if a field has been set.

### GetIsReplicated

`func (o *SubscriptionStats) GetIsReplicated() bool`

GetIsReplicated returns the IsReplicated field if non-nil, zero value otherwise.

### GetIsReplicatedOk

`func (o *SubscriptionStats) GetIsReplicatedOk() (*bool, bool)`

GetIsReplicatedOk returns a tuple with the IsReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicated

`func (o *SubscriptionStats) SetIsReplicated(v bool)`

SetIsReplicated sets IsReplicated field to given value.

### HasIsReplicated

`func (o *SubscriptionStats) HasIsReplicated() bool`

HasIsReplicated returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


