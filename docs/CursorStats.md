# CursorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarkDeletePosition** | Pointer to **string** |  | [optional] 
**ReadPosition** | Pointer to **string** |  | [optional] 
**WaitingReadOp** | Pointer to **bool** |  | [optional] 
**PendingReadOps** | Pointer to **int32** |  | [optional] 
**MessagesConsumedCounter** | Pointer to **int64** |  | [optional] 
**CursorLedger** | Pointer to **int64** |  | [optional] 
**CursorLedgerLastEntry** | Pointer to **int64** |  | [optional] 
**IndividuallyDeletedMessages** | Pointer to **string** |  | [optional] 
**LastLedgerSwitchTimestamp** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**NumberOfEntriesSinceFirstNotAckedMessage** | Pointer to **int64** |  | [optional] 
**TotalNonContiguousDeletedMessagesRange** | Pointer to **int32** |  | [optional] 
**SubscriptionHavePendingRead** | Pointer to **bool** |  | [optional] 
**SubscriptionHavePendingReplayRead** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **map[string]int64** |  | [optional] 

## Methods

### NewCursorStats

`func NewCursorStats() *CursorStats`

NewCursorStats instantiates a new CursorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorStatsWithDefaults

`func NewCursorStatsWithDefaults() *CursorStats`

NewCursorStatsWithDefaults instantiates a new CursorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkDeletePosition

`func (o *CursorStats) GetMarkDeletePosition() string`

GetMarkDeletePosition returns the MarkDeletePosition field if non-nil, zero value otherwise.

### GetMarkDeletePositionOk

`func (o *CursorStats) GetMarkDeletePositionOk() (*string, bool)`

GetMarkDeletePositionOk returns a tuple with the MarkDeletePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkDeletePosition

`func (o *CursorStats) SetMarkDeletePosition(v string)`

SetMarkDeletePosition sets MarkDeletePosition field to given value.

### HasMarkDeletePosition

`func (o *CursorStats) HasMarkDeletePosition() bool`

HasMarkDeletePosition returns a boolean if a field has been set.

### GetReadPosition

`func (o *CursorStats) GetReadPosition() string`

GetReadPosition returns the ReadPosition field if non-nil, zero value otherwise.

### GetReadPositionOk

`func (o *CursorStats) GetReadPositionOk() (*string, bool)`

GetReadPositionOk returns a tuple with the ReadPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPosition

`func (o *CursorStats) SetReadPosition(v string)`

SetReadPosition sets ReadPosition field to given value.

### HasReadPosition

`func (o *CursorStats) HasReadPosition() bool`

HasReadPosition returns a boolean if a field has been set.

### GetWaitingReadOp

`func (o *CursorStats) GetWaitingReadOp() bool`

GetWaitingReadOp returns the WaitingReadOp field if non-nil, zero value otherwise.

### GetWaitingReadOpOk

`func (o *CursorStats) GetWaitingReadOpOk() (*bool, bool)`

GetWaitingReadOpOk returns a tuple with the WaitingReadOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingReadOp

`func (o *CursorStats) SetWaitingReadOp(v bool)`

SetWaitingReadOp sets WaitingReadOp field to given value.

### HasWaitingReadOp

`func (o *CursorStats) HasWaitingReadOp() bool`

HasWaitingReadOp returns a boolean if a field has been set.

### GetPendingReadOps

`func (o *CursorStats) GetPendingReadOps() int32`

GetPendingReadOps returns the PendingReadOps field if non-nil, zero value otherwise.

### GetPendingReadOpsOk

`func (o *CursorStats) GetPendingReadOpsOk() (*int32, bool)`

GetPendingReadOpsOk returns a tuple with the PendingReadOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReadOps

`func (o *CursorStats) SetPendingReadOps(v int32)`

SetPendingReadOps sets PendingReadOps field to given value.

### HasPendingReadOps

`func (o *CursorStats) HasPendingReadOps() bool`

HasPendingReadOps returns a boolean if a field has been set.

### GetMessagesConsumedCounter

`func (o *CursorStats) GetMessagesConsumedCounter() int64`

GetMessagesConsumedCounter returns the MessagesConsumedCounter field if non-nil, zero value otherwise.

### GetMessagesConsumedCounterOk

`func (o *CursorStats) GetMessagesConsumedCounterOk() (*int64, bool)`

GetMessagesConsumedCounterOk returns a tuple with the MessagesConsumedCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesConsumedCounter

`func (o *CursorStats) SetMessagesConsumedCounter(v int64)`

SetMessagesConsumedCounter sets MessagesConsumedCounter field to given value.

### HasMessagesConsumedCounter

`func (o *CursorStats) HasMessagesConsumedCounter() bool`

HasMessagesConsumedCounter returns a boolean if a field has been set.

### GetCursorLedger

`func (o *CursorStats) GetCursorLedger() int64`

GetCursorLedger returns the CursorLedger field if non-nil, zero value otherwise.

### GetCursorLedgerOk

`func (o *CursorStats) GetCursorLedgerOk() (*int64, bool)`

GetCursorLedgerOk returns a tuple with the CursorLedger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedger

`func (o *CursorStats) SetCursorLedger(v int64)`

SetCursorLedger sets CursorLedger field to given value.

### HasCursorLedger

`func (o *CursorStats) HasCursorLedger() bool`

HasCursorLedger returns a boolean if a field has been set.

### GetCursorLedgerLastEntry

`func (o *CursorStats) GetCursorLedgerLastEntry() int64`

GetCursorLedgerLastEntry returns the CursorLedgerLastEntry field if non-nil, zero value otherwise.

### GetCursorLedgerLastEntryOk

`func (o *CursorStats) GetCursorLedgerLastEntryOk() (*int64, bool)`

GetCursorLedgerLastEntryOk returns a tuple with the CursorLedgerLastEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerLastEntry

`func (o *CursorStats) SetCursorLedgerLastEntry(v int64)`

SetCursorLedgerLastEntry sets CursorLedgerLastEntry field to given value.

### HasCursorLedgerLastEntry

`func (o *CursorStats) HasCursorLedgerLastEntry() bool`

HasCursorLedgerLastEntry returns a boolean if a field has been set.

### GetIndividuallyDeletedMessages

`func (o *CursorStats) GetIndividuallyDeletedMessages() string`

GetIndividuallyDeletedMessages returns the IndividuallyDeletedMessages field if non-nil, zero value otherwise.

### GetIndividuallyDeletedMessagesOk

`func (o *CursorStats) GetIndividuallyDeletedMessagesOk() (*string, bool)`

GetIndividuallyDeletedMessagesOk returns a tuple with the IndividuallyDeletedMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividuallyDeletedMessages

`func (o *CursorStats) SetIndividuallyDeletedMessages(v string)`

SetIndividuallyDeletedMessages sets IndividuallyDeletedMessages field to given value.

### HasIndividuallyDeletedMessages

`func (o *CursorStats) HasIndividuallyDeletedMessages() bool`

HasIndividuallyDeletedMessages returns a boolean if a field has been set.

### GetLastLedgerSwitchTimestamp

`func (o *CursorStats) GetLastLedgerSwitchTimestamp() string`

GetLastLedgerSwitchTimestamp returns the LastLedgerSwitchTimestamp field if non-nil, zero value otherwise.

### GetLastLedgerSwitchTimestampOk

`func (o *CursorStats) GetLastLedgerSwitchTimestampOk() (*string, bool)`

GetLastLedgerSwitchTimestampOk returns a tuple with the LastLedgerSwitchTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLedgerSwitchTimestamp

`func (o *CursorStats) SetLastLedgerSwitchTimestamp(v string)`

SetLastLedgerSwitchTimestamp sets LastLedgerSwitchTimestamp field to given value.

### HasLastLedgerSwitchTimestamp

`func (o *CursorStats) HasLastLedgerSwitchTimestamp() bool`

HasLastLedgerSwitchTimestamp returns a boolean if a field has been set.

### GetState

`func (o *CursorStats) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CursorStats) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CursorStats) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CursorStats) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNumberOfEntriesSinceFirstNotAckedMessage

`func (o *CursorStats) GetNumberOfEntriesSinceFirstNotAckedMessage() int64`

GetNumberOfEntriesSinceFirstNotAckedMessage returns the NumberOfEntriesSinceFirstNotAckedMessage field if non-nil, zero value otherwise.

### GetNumberOfEntriesSinceFirstNotAckedMessageOk

`func (o *CursorStats) GetNumberOfEntriesSinceFirstNotAckedMessageOk() (*int64, bool)`

GetNumberOfEntriesSinceFirstNotAckedMessageOk returns a tuple with the NumberOfEntriesSinceFirstNotAckedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEntriesSinceFirstNotAckedMessage

`func (o *CursorStats) SetNumberOfEntriesSinceFirstNotAckedMessage(v int64)`

SetNumberOfEntriesSinceFirstNotAckedMessage sets NumberOfEntriesSinceFirstNotAckedMessage field to given value.

### HasNumberOfEntriesSinceFirstNotAckedMessage

`func (o *CursorStats) HasNumberOfEntriesSinceFirstNotAckedMessage() bool`

HasNumberOfEntriesSinceFirstNotAckedMessage returns a boolean if a field has been set.

### GetTotalNonContiguousDeletedMessagesRange

`func (o *CursorStats) GetTotalNonContiguousDeletedMessagesRange() int32`

GetTotalNonContiguousDeletedMessagesRange returns the TotalNonContiguousDeletedMessagesRange field if non-nil, zero value otherwise.

### GetTotalNonContiguousDeletedMessagesRangeOk

`func (o *CursorStats) GetTotalNonContiguousDeletedMessagesRangeOk() (*int32, bool)`

GetTotalNonContiguousDeletedMessagesRangeOk returns a tuple with the TotalNonContiguousDeletedMessagesRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNonContiguousDeletedMessagesRange

`func (o *CursorStats) SetTotalNonContiguousDeletedMessagesRange(v int32)`

SetTotalNonContiguousDeletedMessagesRange sets TotalNonContiguousDeletedMessagesRange field to given value.

### HasTotalNonContiguousDeletedMessagesRange

`func (o *CursorStats) HasTotalNonContiguousDeletedMessagesRange() bool`

HasTotalNonContiguousDeletedMessagesRange returns a boolean if a field has been set.

### GetSubscriptionHavePendingRead

`func (o *CursorStats) GetSubscriptionHavePendingRead() bool`

GetSubscriptionHavePendingRead returns the SubscriptionHavePendingRead field if non-nil, zero value otherwise.

### GetSubscriptionHavePendingReadOk

`func (o *CursorStats) GetSubscriptionHavePendingReadOk() (*bool, bool)`

GetSubscriptionHavePendingReadOk returns a tuple with the SubscriptionHavePendingRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionHavePendingRead

`func (o *CursorStats) SetSubscriptionHavePendingRead(v bool)`

SetSubscriptionHavePendingRead sets SubscriptionHavePendingRead field to given value.

### HasSubscriptionHavePendingRead

`func (o *CursorStats) HasSubscriptionHavePendingRead() bool`

HasSubscriptionHavePendingRead returns a boolean if a field has been set.

### GetSubscriptionHavePendingReplayRead

`func (o *CursorStats) GetSubscriptionHavePendingReplayRead() bool`

GetSubscriptionHavePendingReplayRead returns the SubscriptionHavePendingReplayRead field if non-nil, zero value otherwise.

### GetSubscriptionHavePendingReplayReadOk

`func (o *CursorStats) GetSubscriptionHavePendingReplayReadOk() (*bool, bool)`

GetSubscriptionHavePendingReplayReadOk returns a tuple with the SubscriptionHavePendingReplayRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionHavePendingReplayRead

`func (o *CursorStats) SetSubscriptionHavePendingReplayRead(v bool)`

SetSubscriptionHavePendingReplayRead sets SubscriptionHavePendingReplayRead field to given value.

### HasSubscriptionHavePendingReplayRead

`func (o *CursorStats) HasSubscriptionHavePendingReplayRead() bool`

HasSubscriptionHavePendingReplayRead returns a boolean if a field has been set.

### GetProperties

`func (o *CursorStats) GetProperties() map[string]int64`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CursorStats) GetPropertiesOk() (*map[string]int64, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CursorStats) SetProperties(v map[string]int64)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CursorStats) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


