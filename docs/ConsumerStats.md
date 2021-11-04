# ConsumerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AvailablePermits** | Pointer to **int32** |  | [optional] 
**AvgMessagesPerEntry** | Pointer to **int32** |  | [optional] 
**BlockedConsumerOnUnackedMsgs** | Pointer to **bool** |  | [optional] 
**BytesOutCounter** | Pointer to **int64** |  | [optional] 
**ChuckedMessageRate** | Pointer to **float64** |  | [optional] 
**ClientVersion** | Pointer to **string** |  | [optional] 
**ConnectedSince** | Pointer to **string** |  | [optional] 
**ConsumerName** | Pointer to **string** |  | [optional] 
**KeyHashRanges** | Pointer to **[]string** |  | [optional] 
**LastAckedTimestamp** | Pointer to **int64** |  | [optional] 
**LastConsumedTimestamp** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MsgOutCounter** | Pointer to **int64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgRateRedeliver** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**ReadPositionWhenJoining** | Pointer to **string** |  | [optional] 
**UnackedMessages** | Pointer to **int32** |  | [optional] 

## Methods

### NewConsumerStats

`func NewConsumerStats() *ConsumerStats`

NewConsumerStats instantiates a new ConsumerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerStatsWithDefaults

`func NewConsumerStatsWithDefaults() *ConsumerStats`

NewConsumerStatsWithDefaults instantiates a new ConsumerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConsumerStats) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsumerStats) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsumerStats) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsumerStats) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvailablePermits

`func (o *ConsumerStats) GetAvailablePermits() int32`

GetAvailablePermits returns the AvailablePermits field if non-nil, zero value otherwise.

### GetAvailablePermitsOk

`func (o *ConsumerStats) GetAvailablePermitsOk() (*int32, bool)`

GetAvailablePermitsOk returns a tuple with the AvailablePermits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePermits

`func (o *ConsumerStats) SetAvailablePermits(v int32)`

SetAvailablePermits sets AvailablePermits field to given value.

### HasAvailablePermits

`func (o *ConsumerStats) HasAvailablePermits() bool`

HasAvailablePermits returns a boolean if a field has been set.

### GetAvgMessagesPerEntry

`func (o *ConsumerStats) GetAvgMessagesPerEntry() int32`

GetAvgMessagesPerEntry returns the AvgMessagesPerEntry field if non-nil, zero value otherwise.

### GetAvgMessagesPerEntryOk

`func (o *ConsumerStats) GetAvgMessagesPerEntryOk() (*int32, bool)`

GetAvgMessagesPerEntryOk returns a tuple with the AvgMessagesPerEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMessagesPerEntry

`func (o *ConsumerStats) SetAvgMessagesPerEntry(v int32)`

SetAvgMessagesPerEntry sets AvgMessagesPerEntry field to given value.

### HasAvgMessagesPerEntry

`func (o *ConsumerStats) HasAvgMessagesPerEntry() bool`

HasAvgMessagesPerEntry returns a boolean if a field has been set.

### GetBlockedConsumerOnUnackedMsgs

`func (o *ConsumerStats) GetBlockedConsumerOnUnackedMsgs() bool`

GetBlockedConsumerOnUnackedMsgs returns the BlockedConsumerOnUnackedMsgs field if non-nil, zero value otherwise.

### GetBlockedConsumerOnUnackedMsgsOk

`func (o *ConsumerStats) GetBlockedConsumerOnUnackedMsgsOk() (*bool, bool)`

GetBlockedConsumerOnUnackedMsgsOk returns a tuple with the BlockedConsumerOnUnackedMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedConsumerOnUnackedMsgs

`func (o *ConsumerStats) SetBlockedConsumerOnUnackedMsgs(v bool)`

SetBlockedConsumerOnUnackedMsgs sets BlockedConsumerOnUnackedMsgs field to given value.

### HasBlockedConsumerOnUnackedMsgs

`func (o *ConsumerStats) HasBlockedConsumerOnUnackedMsgs() bool`

HasBlockedConsumerOnUnackedMsgs returns a boolean if a field has been set.

### GetBytesOutCounter

`func (o *ConsumerStats) GetBytesOutCounter() int64`

GetBytesOutCounter returns the BytesOutCounter field if non-nil, zero value otherwise.

### GetBytesOutCounterOk

`func (o *ConsumerStats) GetBytesOutCounterOk() (*int64, bool)`

GetBytesOutCounterOk returns a tuple with the BytesOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOutCounter

`func (o *ConsumerStats) SetBytesOutCounter(v int64)`

SetBytesOutCounter sets BytesOutCounter field to given value.

### HasBytesOutCounter

`func (o *ConsumerStats) HasBytesOutCounter() bool`

HasBytesOutCounter returns a boolean if a field has been set.

### GetChuckedMessageRate

`func (o *ConsumerStats) GetChuckedMessageRate() float64`

GetChuckedMessageRate returns the ChuckedMessageRate field if non-nil, zero value otherwise.

### GetChuckedMessageRateOk

`func (o *ConsumerStats) GetChuckedMessageRateOk() (*float64, bool)`

GetChuckedMessageRateOk returns a tuple with the ChuckedMessageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChuckedMessageRate

`func (o *ConsumerStats) SetChuckedMessageRate(v float64)`

SetChuckedMessageRate sets ChuckedMessageRate field to given value.

### HasChuckedMessageRate

`func (o *ConsumerStats) HasChuckedMessageRate() bool`

HasChuckedMessageRate returns a boolean if a field has been set.

### GetClientVersion

`func (o *ConsumerStats) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *ConsumerStats) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *ConsumerStats) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *ConsumerStats) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetConnectedSince

`func (o *ConsumerStats) GetConnectedSince() string`

GetConnectedSince returns the ConnectedSince field if non-nil, zero value otherwise.

### GetConnectedSinceOk

`func (o *ConsumerStats) GetConnectedSinceOk() (*string, bool)`

GetConnectedSinceOk returns a tuple with the ConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSince

`func (o *ConsumerStats) SetConnectedSince(v string)`

SetConnectedSince sets ConnectedSince field to given value.

### HasConnectedSince

`func (o *ConsumerStats) HasConnectedSince() bool`

HasConnectedSince returns a boolean if a field has been set.

### GetConsumerName

`func (o *ConsumerStats) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *ConsumerStats) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *ConsumerStats) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *ConsumerStats) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetKeyHashRanges

`func (o *ConsumerStats) GetKeyHashRanges() []string`

GetKeyHashRanges returns the KeyHashRanges field if non-nil, zero value otherwise.

### GetKeyHashRangesOk

`func (o *ConsumerStats) GetKeyHashRangesOk() (*[]string, bool)`

GetKeyHashRangesOk returns a tuple with the KeyHashRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyHashRanges

`func (o *ConsumerStats) SetKeyHashRanges(v []string)`

SetKeyHashRanges sets KeyHashRanges field to given value.

### HasKeyHashRanges

`func (o *ConsumerStats) HasKeyHashRanges() bool`

HasKeyHashRanges returns a boolean if a field has been set.

### GetLastAckedTimestamp

`func (o *ConsumerStats) GetLastAckedTimestamp() int64`

GetLastAckedTimestamp returns the LastAckedTimestamp field if non-nil, zero value otherwise.

### GetLastAckedTimestampOk

`func (o *ConsumerStats) GetLastAckedTimestampOk() (*int64, bool)`

GetLastAckedTimestampOk returns a tuple with the LastAckedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAckedTimestamp

`func (o *ConsumerStats) SetLastAckedTimestamp(v int64)`

SetLastAckedTimestamp sets LastAckedTimestamp field to given value.

### HasLastAckedTimestamp

`func (o *ConsumerStats) HasLastAckedTimestamp() bool`

HasLastAckedTimestamp returns a boolean if a field has been set.

### GetLastConsumedTimestamp

`func (o *ConsumerStats) GetLastConsumedTimestamp() int64`

GetLastConsumedTimestamp returns the LastConsumedTimestamp field if non-nil, zero value otherwise.

### GetLastConsumedTimestampOk

`func (o *ConsumerStats) GetLastConsumedTimestampOk() (*int64, bool)`

GetLastConsumedTimestampOk returns a tuple with the LastConsumedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTimestamp

`func (o *ConsumerStats) SetLastConsumedTimestamp(v int64)`

SetLastConsumedTimestamp sets LastConsumedTimestamp field to given value.

### HasLastConsumedTimestamp

`func (o *ConsumerStats) HasLastConsumedTimestamp() bool`

HasLastConsumedTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsumerStats) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsumerStats) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsumerStats) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsumerStats) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMsgOutCounter

`func (o *ConsumerStats) GetMsgOutCounter() int64`

GetMsgOutCounter returns the MsgOutCounter field if non-nil, zero value otherwise.

### GetMsgOutCounterOk

`func (o *ConsumerStats) GetMsgOutCounterOk() (*int64, bool)`

GetMsgOutCounterOk returns a tuple with the MsgOutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgOutCounter

`func (o *ConsumerStats) SetMsgOutCounter(v int64)`

SetMsgOutCounter sets MsgOutCounter field to given value.

### HasMsgOutCounter

`func (o *ConsumerStats) HasMsgOutCounter() bool`

HasMsgOutCounter returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *ConsumerStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *ConsumerStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *ConsumerStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *ConsumerStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgRateRedeliver

`func (o *ConsumerStats) GetMsgRateRedeliver() float64`

GetMsgRateRedeliver returns the MsgRateRedeliver field if non-nil, zero value otherwise.

### GetMsgRateRedeliverOk

`func (o *ConsumerStats) GetMsgRateRedeliverOk() (*float64, bool)`

GetMsgRateRedeliverOk returns a tuple with the MsgRateRedeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateRedeliver

`func (o *ConsumerStats) SetMsgRateRedeliver(v float64)`

SetMsgRateRedeliver sets MsgRateRedeliver field to given value.

### HasMsgRateRedeliver

`func (o *ConsumerStats) HasMsgRateRedeliver() bool`

HasMsgRateRedeliver returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *ConsumerStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *ConsumerStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *ConsumerStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *ConsumerStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetReadPositionWhenJoining

`func (o *ConsumerStats) GetReadPositionWhenJoining() string`

GetReadPositionWhenJoining returns the ReadPositionWhenJoining field if non-nil, zero value otherwise.

### GetReadPositionWhenJoiningOk

`func (o *ConsumerStats) GetReadPositionWhenJoiningOk() (*string, bool)`

GetReadPositionWhenJoiningOk returns a tuple with the ReadPositionWhenJoining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPositionWhenJoining

`func (o *ConsumerStats) SetReadPositionWhenJoining(v string)`

SetReadPositionWhenJoining sets ReadPositionWhenJoining field to given value.

### HasReadPositionWhenJoining

`func (o *ConsumerStats) HasReadPositionWhenJoining() bool`

HasReadPositionWhenJoining returns a boolean if a field has been set.

### GetUnackedMessages

`func (o *ConsumerStats) GetUnackedMessages() int32`

GetUnackedMessages returns the UnackedMessages field if non-nil, zero value otherwise.

### GetUnackedMessagesOk

`func (o *ConsumerStats) GetUnackedMessagesOk() (*int32, bool)`

GetUnackedMessagesOk returns a tuple with the UnackedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnackedMessages

`func (o *ConsumerStats) SetUnackedMessages(v int32)`

SetUnackedMessages sets UnackedMessages field to given value.

### HasUnackedMessages

`func (o *ConsumerStats) HasUnackedMessages() bool`

HasUnackedMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


