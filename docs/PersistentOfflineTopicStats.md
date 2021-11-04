# PersistentOfflineTopicStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokerName** | Pointer to **string** |  | [optional] 
**CursorDetails** | Pointer to [**map[string]CursorDetails**](CursorDetails.md) |  | [optional] 
**DataLedgerDetails** | Pointer to [**[]LedgerDetails**](LedgerDetails.md) |  | [optional] 
**MessageBacklog** | Pointer to **int64** |  | [optional] 
**StatGeneratedAt** | Pointer to **time.Time** |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**TopicName** | Pointer to **string** |  | [optional] 
**TotalMessages** | Pointer to **int64** |  | [optional] 

## Methods

### NewPersistentOfflineTopicStats

`func NewPersistentOfflineTopicStats() *PersistentOfflineTopicStats`

NewPersistentOfflineTopicStats instantiates a new PersistentOfflineTopicStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentOfflineTopicStatsWithDefaults

`func NewPersistentOfflineTopicStatsWithDefaults() *PersistentOfflineTopicStats`

NewPersistentOfflineTopicStatsWithDefaults instantiates a new PersistentOfflineTopicStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokerName

`func (o *PersistentOfflineTopicStats) GetBrokerName() string`

GetBrokerName returns the BrokerName field if non-nil, zero value otherwise.

### GetBrokerNameOk

`func (o *PersistentOfflineTopicStats) GetBrokerNameOk() (*string, bool)`

GetBrokerNameOk returns a tuple with the BrokerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerName

`func (o *PersistentOfflineTopicStats) SetBrokerName(v string)`

SetBrokerName sets BrokerName field to given value.

### HasBrokerName

`func (o *PersistentOfflineTopicStats) HasBrokerName() bool`

HasBrokerName returns a boolean if a field has been set.

### GetCursorDetails

`func (o *PersistentOfflineTopicStats) GetCursorDetails() map[string]CursorDetails`

GetCursorDetails returns the CursorDetails field if non-nil, zero value otherwise.

### GetCursorDetailsOk

`func (o *PersistentOfflineTopicStats) GetCursorDetailsOk() (*map[string]CursorDetails, bool)`

GetCursorDetailsOk returns a tuple with the CursorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorDetails

`func (o *PersistentOfflineTopicStats) SetCursorDetails(v map[string]CursorDetails)`

SetCursorDetails sets CursorDetails field to given value.

### HasCursorDetails

`func (o *PersistentOfflineTopicStats) HasCursorDetails() bool`

HasCursorDetails returns a boolean if a field has been set.

### GetDataLedgerDetails

`func (o *PersistentOfflineTopicStats) GetDataLedgerDetails() []LedgerDetails`

GetDataLedgerDetails returns the DataLedgerDetails field if non-nil, zero value otherwise.

### GetDataLedgerDetailsOk

`func (o *PersistentOfflineTopicStats) GetDataLedgerDetailsOk() (*[]LedgerDetails, bool)`

GetDataLedgerDetailsOk returns a tuple with the DataLedgerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLedgerDetails

`func (o *PersistentOfflineTopicStats) SetDataLedgerDetails(v []LedgerDetails)`

SetDataLedgerDetails sets DataLedgerDetails field to given value.

### HasDataLedgerDetails

`func (o *PersistentOfflineTopicStats) HasDataLedgerDetails() bool`

HasDataLedgerDetails returns a boolean if a field has been set.

### GetMessageBacklog

`func (o *PersistentOfflineTopicStats) GetMessageBacklog() int64`

GetMessageBacklog returns the MessageBacklog field if non-nil, zero value otherwise.

### GetMessageBacklogOk

`func (o *PersistentOfflineTopicStats) GetMessageBacklogOk() (*int64, bool)`

GetMessageBacklogOk returns a tuple with the MessageBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBacklog

`func (o *PersistentOfflineTopicStats) SetMessageBacklog(v int64)`

SetMessageBacklog sets MessageBacklog field to given value.

### HasMessageBacklog

`func (o *PersistentOfflineTopicStats) HasMessageBacklog() bool`

HasMessageBacklog returns a boolean if a field has been set.

### GetStatGeneratedAt

`func (o *PersistentOfflineTopicStats) GetStatGeneratedAt() time.Time`

GetStatGeneratedAt returns the StatGeneratedAt field if non-nil, zero value otherwise.

### GetStatGeneratedAtOk

`func (o *PersistentOfflineTopicStats) GetStatGeneratedAtOk() (*time.Time, bool)`

GetStatGeneratedAtOk returns a tuple with the StatGeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatGeneratedAt

`func (o *PersistentOfflineTopicStats) SetStatGeneratedAt(v time.Time)`

SetStatGeneratedAt sets StatGeneratedAt field to given value.

### HasStatGeneratedAt

`func (o *PersistentOfflineTopicStats) HasStatGeneratedAt() bool`

HasStatGeneratedAt returns a boolean if a field has been set.

### GetStorageSize

`func (o *PersistentOfflineTopicStats) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *PersistentOfflineTopicStats) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *PersistentOfflineTopicStats) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *PersistentOfflineTopicStats) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetTopicName

`func (o *PersistentOfflineTopicStats) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *PersistentOfflineTopicStats) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *PersistentOfflineTopicStats) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.

### HasTopicName

`func (o *PersistentOfflineTopicStats) HasTopicName() bool`

HasTopicName returns a boolean if a field has been set.

### GetTotalMessages

`func (o *PersistentOfflineTopicStats) GetTotalMessages() int64`

GetTotalMessages returns the TotalMessages field if non-nil, zero value otherwise.

### GetTotalMessagesOk

`func (o *PersistentOfflineTopicStats) GetTotalMessagesOk() (*int64, bool)`

GetTotalMessagesOk returns a tuple with the TotalMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessages

`func (o *PersistentOfflineTopicStats) SetTotalMessages(v int64)`

SetTotalMessages sets TotalMessages field to given value.

### HasTotalMessages

`func (o *PersistentOfflineTopicStats) HasTotalMessages() bool`

HasTotalMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


