# PersistentTopicInternalStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntriesAddedCounter** | Pointer to **int64** |  | [optional] 
**NumberOfEntries** | Pointer to **int64** |  | [optional] 
**TotalSize** | Pointer to **int64** |  | [optional] 
**CurrentLedgerEntries** | Pointer to **int64** |  | [optional] 
**CurrentLedgerSize** | Pointer to **int64** |  | [optional] 
**LastLedgerCreatedTimestamp** | Pointer to **string** |  | [optional] 
**LastLedgerCreationFailureTimestamp** | Pointer to **string** |  | [optional] 
**WaitingCursorsCount** | Pointer to **int32** |  | [optional] 
**PendingAddEntriesCount** | Pointer to **int32** |  | [optional] 
**LastConfirmedEntry** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Ledgers** | Pointer to [**[]LedgerInfo**](LedgerInfo.md) |  | [optional] 
**Cursors** | Pointer to [**map[string]CursorStats**](CursorStats.md) |  | [optional] 
**SchemaLedgers** | Pointer to [**[]LedgerInfo**](LedgerInfo.md) |  | [optional] 
**CompactedLedger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | [optional] 

## Methods

### NewPersistentTopicInternalStats

`func NewPersistentTopicInternalStats() *PersistentTopicInternalStats`

NewPersistentTopicInternalStats instantiates a new PersistentTopicInternalStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentTopicInternalStatsWithDefaults

`func NewPersistentTopicInternalStatsWithDefaults() *PersistentTopicInternalStats`

NewPersistentTopicInternalStatsWithDefaults instantiates a new PersistentTopicInternalStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntriesAddedCounter

`func (o *PersistentTopicInternalStats) GetEntriesAddedCounter() int64`

GetEntriesAddedCounter returns the EntriesAddedCounter field if non-nil, zero value otherwise.

### GetEntriesAddedCounterOk

`func (o *PersistentTopicInternalStats) GetEntriesAddedCounterOk() (*int64, bool)`

GetEntriesAddedCounterOk returns a tuple with the EntriesAddedCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesAddedCounter

`func (o *PersistentTopicInternalStats) SetEntriesAddedCounter(v int64)`

SetEntriesAddedCounter sets EntriesAddedCounter field to given value.

### HasEntriesAddedCounter

`func (o *PersistentTopicInternalStats) HasEntriesAddedCounter() bool`

HasEntriesAddedCounter returns a boolean if a field has been set.

### GetNumberOfEntries

`func (o *PersistentTopicInternalStats) GetNumberOfEntries() int64`

GetNumberOfEntries returns the NumberOfEntries field if non-nil, zero value otherwise.

### GetNumberOfEntriesOk

`func (o *PersistentTopicInternalStats) GetNumberOfEntriesOk() (*int64, bool)`

GetNumberOfEntriesOk returns a tuple with the NumberOfEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEntries

`func (o *PersistentTopicInternalStats) SetNumberOfEntries(v int64)`

SetNumberOfEntries sets NumberOfEntries field to given value.

### HasNumberOfEntries

`func (o *PersistentTopicInternalStats) HasNumberOfEntries() bool`

HasNumberOfEntries returns a boolean if a field has been set.

### GetTotalSize

`func (o *PersistentTopicInternalStats) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *PersistentTopicInternalStats) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *PersistentTopicInternalStats) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *PersistentTopicInternalStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetCurrentLedgerEntries

`func (o *PersistentTopicInternalStats) GetCurrentLedgerEntries() int64`

GetCurrentLedgerEntries returns the CurrentLedgerEntries field if non-nil, zero value otherwise.

### GetCurrentLedgerEntriesOk

`func (o *PersistentTopicInternalStats) GetCurrentLedgerEntriesOk() (*int64, bool)`

GetCurrentLedgerEntriesOk returns a tuple with the CurrentLedgerEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLedgerEntries

`func (o *PersistentTopicInternalStats) SetCurrentLedgerEntries(v int64)`

SetCurrentLedgerEntries sets CurrentLedgerEntries field to given value.

### HasCurrentLedgerEntries

`func (o *PersistentTopicInternalStats) HasCurrentLedgerEntries() bool`

HasCurrentLedgerEntries returns a boolean if a field has been set.

### GetCurrentLedgerSize

`func (o *PersistentTopicInternalStats) GetCurrentLedgerSize() int64`

GetCurrentLedgerSize returns the CurrentLedgerSize field if non-nil, zero value otherwise.

### GetCurrentLedgerSizeOk

`func (o *PersistentTopicInternalStats) GetCurrentLedgerSizeOk() (*int64, bool)`

GetCurrentLedgerSizeOk returns a tuple with the CurrentLedgerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLedgerSize

`func (o *PersistentTopicInternalStats) SetCurrentLedgerSize(v int64)`

SetCurrentLedgerSize sets CurrentLedgerSize field to given value.

### HasCurrentLedgerSize

`func (o *PersistentTopicInternalStats) HasCurrentLedgerSize() bool`

HasCurrentLedgerSize returns a boolean if a field has been set.

### GetLastLedgerCreatedTimestamp

`func (o *PersistentTopicInternalStats) GetLastLedgerCreatedTimestamp() string`

GetLastLedgerCreatedTimestamp returns the LastLedgerCreatedTimestamp field if non-nil, zero value otherwise.

### GetLastLedgerCreatedTimestampOk

`func (o *PersistentTopicInternalStats) GetLastLedgerCreatedTimestampOk() (*string, bool)`

GetLastLedgerCreatedTimestampOk returns a tuple with the LastLedgerCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLedgerCreatedTimestamp

`func (o *PersistentTopicInternalStats) SetLastLedgerCreatedTimestamp(v string)`

SetLastLedgerCreatedTimestamp sets LastLedgerCreatedTimestamp field to given value.

### HasLastLedgerCreatedTimestamp

`func (o *PersistentTopicInternalStats) HasLastLedgerCreatedTimestamp() bool`

HasLastLedgerCreatedTimestamp returns a boolean if a field has been set.

### GetLastLedgerCreationFailureTimestamp

`func (o *PersistentTopicInternalStats) GetLastLedgerCreationFailureTimestamp() string`

GetLastLedgerCreationFailureTimestamp returns the LastLedgerCreationFailureTimestamp field if non-nil, zero value otherwise.

### GetLastLedgerCreationFailureTimestampOk

`func (o *PersistentTopicInternalStats) GetLastLedgerCreationFailureTimestampOk() (*string, bool)`

GetLastLedgerCreationFailureTimestampOk returns a tuple with the LastLedgerCreationFailureTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLedgerCreationFailureTimestamp

`func (o *PersistentTopicInternalStats) SetLastLedgerCreationFailureTimestamp(v string)`

SetLastLedgerCreationFailureTimestamp sets LastLedgerCreationFailureTimestamp field to given value.

### HasLastLedgerCreationFailureTimestamp

`func (o *PersistentTopicInternalStats) HasLastLedgerCreationFailureTimestamp() bool`

HasLastLedgerCreationFailureTimestamp returns a boolean if a field has been set.

### GetWaitingCursorsCount

`func (o *PersistentTopicInternalStats) GetWaitingCursorsCount() int32`

GetWaitingCursorsCount returns the WaitingCursorsCount field if non-nil, zero value otherwise.

### GetWaitingCursorsCountOk

`func (o *PersistentTopicInternalStats) GetWaitingCursorsCountOk() (*int32, bool)`

GetWaitingCursorsCountOk returns a tuple with the WaitingCursorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingCursorsCount

`func (o *PersistentTopicInternalStats) SetWaitingCursorsCount(v int32)`

SetWaitingCursorsCount sets WaitingCursorsCount field to given value.

### HasWaitingCursorsCount

`func (o *PersistentTopicInternalStats) HasWaitingCursorsCount() bool`

HasWaitingCursorsCount returns a boolean if a field has been set.

### GetPendingAddEntriesCount

`func (o *PersistentTopicInternalStats) GetPendingAddEntriesCount() int32`

GetPendingAddEntriesCount returns the PendingAddEntriesCount field if non-nil, zero value otherwise.

### GetPendingAddEntriesCountOk

`func (o *PersistentTopicInternalStats) GetPendingAddEntriesCountOk() (*int32, bool)`

GetPendingAddEntriesCountOk returns a tuple with the PendingAddEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAddEntriesCount

`func (o *PersistentTopicInternalStats) SetPendingAddEntriesCount(v int32)`

SetPendingAddEntriesCount sets PendingAddEntriesCount field to given value.

### HasPendingAddEntriesCount

`func (o *PersistentTopicInternalStats) HasPendingAddEntriesCount() bool`

HasPendingAddEntriesCount returns a boolean if a field has been set.

### GetLastConfirmedEntry

`func (o *PersistentTopicInternalStats) GetLastConfirmedEntry() string`

GetLastConfirmedEntry returns the LastConfirmedEntry field if non-nil, zero value otherwise.

### GetLastConfirmedEntryOk

`func (o *PersistentTopicInternalStats) GetLastConfirmedEntryOk() (*string, bool)`

GetLastConfirmedEntryOk returns a tuple with the LastConfirmedEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfirmedEntry

`func (o *PersistentTopicInternalStats) SetLastConfirmedEntry(v string)`

SetLastConfirmedEntry sets LastConfirmedEntry field to given value.

### HasLastConfirmedEntry

`func (o *PersistentTopicInternalStats) HasLastConfirmedEntry() bool`

HasLastConfirmedEntry returns a boolean if a field has been set.

### GetState

`func (o *PersistentTopicInternalStats) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PersistentTopicInternalStats) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PersistentTopicInternalStats) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PersistentTopicInternalStats) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLedgers

`func (o *PersistentTopicInternalStats) GetLedgers() []LedgerInfo`

GetLedgers returns the Ledgers field if non-nil, zero value otherwise.

### GetLedgersOk

`func (o *PersistentTopicInternalStats) GetLedgersOk() (*[]LedgerInfo, bool)`

GetLedgersOk returns a tuple with the Ledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgers

`func (o *PersistentTopicInternalStats) SetLedgers(v []LedgerInfo)`

SetLedgers sets Ledgers field to given value.

### HasLedgers

`func (o *PersistentTopicInternalStats) HasLedgers() bool`

HasLedgers returns a boolean if a field has been set.

### GetCursors

`func (o *PersistentTopicInternalStats) GetCursors() map[string]CursorStats`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *PersistentTopicInternalStats) GetCursorsOk() (*map[string]CursorStats, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *PersistentTopicInternalStats) SetCursors(v map[string]CursorStats)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *PersistentTopicInternalStats) HasCursors() bool`

HasCursors returns a boolean if a field has been set.

### GetSchemaLedgers

`func (o *PersistentTopicInternalStats) GetSchemaLedgers() []LedgerInfo`

GetSchemaLedgers returns the SchemaLedgers field if non-nil, zero value otherwise.

### GetSchemaLedgersOk

`func (o *PersistentTopicInternalStats) GetSchemaLedgersOk() (*[]LedgerInfo, bool)`

GetSchemaLedgersOk returns a tuple with the SchemaLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLedgers

`func (o *PersistentTopicInternalStats) SetSchemaLedgers(v []LedgerInfo)`

SetSchemaLedgers sets SchemaLedgers field to given value.

### HasSchemaLedgers

`func (o *PersistentTopicInternalStats) HasSchemaLedgers() bool`

HasSchemaLedgers returns a boolean if a field has been set.

### GetCompactedLedger

`func (o *PersistentTopicInternalStats) GetCompactedLedger() LedgerInfo`

GetCompactedLedger returns the CompactedLedger field if non-nil, zero value otherwise.

### GetCompactedLedgerOk

`func (o *PersistentTopicInternalStats) GetCompactedLedgerOk() (*LedgerInfo, bool)`

GetCompactedLedgerOk returns a tuple with the CompactedLedger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactedLedger

`func (o *PersistentTopicInternalStats) SetCompactedLedger(v LedgerInfo)`

SetCompactedLedger sets CompactedLedger field to given value.

### HasCompactedLedger

`func (o *PersistentTopicInternalStats) HasCompactedLedger() bool`

HasCompactedLedger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


