# NonPersistentReplicatorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgDropRate** | Pointer to **float64** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**ReplicationDelayInSeconds** | Pointer to **int64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgRateExpired** | Pointer to **float64** |  | [optional] 
**OutboundConnection** | Pointer to **string** |  | [optional] 
**InboundConnectedSince** | Pointer to **string** |  | [optional] 
**InboundConnection** | Pointer to **string** |  | [optional] 
**OutboundConnectedSince** | Pointer to **string** |  | [optional] 
**ReplicationBacklog** | Pointer to **int64** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 

## Methods

### NewNonPersistentReplicatorStats

`func NewNonPersistentReplicatorStats() *NonPersistentReplicatorStats`

NewNonPersistentReplicatorStats instantiates a new NonPersistentReplicatorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPersistentReplicatorStatsWithDefaults

`func NewNonPersistentReplicatorStatsWithDefaults() *NonPersistentReplicatorStats`

NewNonPersistentReplicatorStatsWithDefaults instantiates a new NonPersistentReplicatorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgDropRate

`func (o *NonPersistentReplicatorStats) GetMsgDropRate() float64`

GetMsgDropRate returns the MsgDropRate field if non-nil, zero value otherwise.

### GetMsgDropRateOk

`func (o *NonPersistentReplicatorStats) GetMsgDropRateOk() (*float64, bool)`

GetMsgDropRateOk returns a tuple with the MsgDropRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDropRate

`func (o *NonPersistentReplicatorStats) SetMsgDropRate(v float64)`

SetMsgDropRate sets MsgDropRate field to given value.

### HasMsgDropRate

`func (o *NonPersistentReplicatorStats) HasMsgDropRate() bool`

HasMsgDropRate returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *NonPersistentReplicatorStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *NonPersistentReplicatorStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *NonPersistentReplicatorStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *NonPersistentReplicatorStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *NonPersistentReplicatorStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *NonPersistentReplicatorStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *NonPersistentReplicatorStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *NonPersistentReplicatorStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetReplicationDelayInSeconds

`func (o *NonPersistentReplicatorStats) GetReplicationDelayInSeconds() int64`

GetReplicationDelayInSeconds returns the ReplicationDelayInSeconds field if non-nil, zero value otherwise.

### GetReplicationDelayInSecondsOk

`func (o *NonPersistentReplicatorStats) GetReplicationDelayInSecondsOk() (*int64, bool)`

GetReplicationDelayInSecondsOk returns a tuple with the ReplicationDelayInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDelayInSeconds

`func (o *NonPersistentReplicatorStats) SetReplicationDelayInSeconds(v int64)`

SetReplicationDelayInSeconds sets ReplicationDelayInSeconds field to given value.

### HasReplicationDelayInSeconds

`func (o *NonPersistentReplicatorStats) HasReplicationDelayInSeconds() bool`

HasReplicationDelayInSeconds returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *NonPersistentReplicatorStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *NonPersistentReplicatorStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *NonPersistentReplicatorStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *NonPersistentReplicatorStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *NonPersistentReplicatorStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *NonPersistentReplicatorStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *NonPersistentReplicatorStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *NonPersistentReplicatorStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgRateExpired

`func (o *NonPersistentReplicatorStats) GetMsgRateExpired() float64`

GetMsgRateExpired returns the MsgRateExpired field if non-nil, zero value otherwise.

### GetMsgRateExpiredOk

`func (o *NonPersistentReplicatorStats) GetMsgRateExpiredOk() (*float64, bool)`

GetMsgRateExpiredOk returns a tuple with the MsgRateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateExpired

`func (o *NonPersistentReplicatorStats) SetMsgRateExpired(v float64)`

SetMsgRateExpired sets MsgRateExpired field to given value.

### HasMsgRateExpired

`func (o *NonPersistentReplicatorStats) HasMsgRateExpired() bool`

HasMsgRateExpired returns a boolean if a field has been set.

### GetOutboundConnection

`func (o *NonPersistentReplicatorStats) GetOutboundConnection() string`

GetOutboundConnection returns the OutboundConnection field if non-nil, zero value otherwise.

### GetOutboundConnectionOk

`func (o *NonPersistentReplicatorStats) GetOutboundConnectionOk() (*string, bool)`

GetOutboundConnectionOk returns a tuple with the OutboundConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnection

`func (o *NonPersistentReplicatorStats) SetOutboundConnection(v string)`

SetOutboundConnection sets OutboundConnection field to given value.

### HasOutboundConnection

`func (o *NonPersistentReplicatorStats) HasOutboundConnection() bool`

HasOutboundConnection returns a boolean if a field has been set.

### GetInboundConnectedSince

`func (o *NonPersistentReplicatorStats) GetInboundConnectedSince() string`

GetInboundConnectedSince returns the InboundConnectedSince field if non-nil, zero value otherwise.

### GetInboundConnectedSinceOk

`func (o *NonPersistentReplicatorStats) GetInboundConnectedSinceOk() (*string, bool)`

GetInboundConnectedSinceOk returns a tuple with the InboundConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnectedSince

`func (o *NonPersistentReplicatorStats) SetInboundConnectedSince(v string)`

SetInboundConnectedSince sets InboundConnectedSince field to given value.

### HasInboundConnectedSince

`func (o *NonPersistentReplicatorStats) HasInboundConnectedSince() bool`

HasInboundConnectedSince returns a boolean if a field has been set.

### GetInboundConnection

`func (o *NonPersistentReplicatorStats) GetInboundConnection() string`

GetInboundConnection returns the InboundConnection field if non-nil, zero value otherwise.

### GetInboundConnectionOk

`func (o *NonPersistentReplicatorStats) GetInboundConnectionOk() (*string, bool)`

GetInboundConnectionOk returns a tuple with the InboundConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnection

`func (o *NonPersistentReplicatorStats) SetInboundConnection(v string)`

SetInboundConnection sets InboundConnection field to given value.

### HasInboundConnection

`func (o *NonPersistentReplicatorStats) HasInboundConnection() bool`

HasInboundConnection returns a boolean if a field has been set.

### GetOutboundConnectedSince

`func (o *NonPersistentReplicatorStats) GetOutboundConnectedSince() string`

GetOutboundConnectedSince returns the OutboundConnectedSince field if non-nil, zero value otherwise.

### GetOutboundConnectedSinceOk

`func (o *NonPersistentReplicatorStats) GetOutboundConnectedSinceOk() (*string, bool)`

GetOutboundConnectedSinceOk returns a tuple with the OutboundConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnectedSince

`func (o *NonPersistentReplicatorStats) SetOutboundConnectedSince(v string)`

SetOutboundConnectedSince sets OutboundConnectedSince field to given value.

### HasOutboundConnectedSince

`func (o *NonPersistentReplicatorStats) HasOutboundConnectedSince() bool`

HasOutboundConnectedSince returns a boolean if a field has been set.

### GetReplicationBacklog

`func (o *NonPersistentReplicatorStats) GetReplicationBacklog() int64`

GetReplicationBacklog returns the ReplicationBacklog field if non-nil, zero value otherwise.

### GetReplicationBacklogOk

`func (o *NonPersistentReplicatorStats) GetReplicationBacklogOk() (*int64, bool)`

GetReplicationBacklogOk returns a tuple with the ReplicationBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBacklog

`func (o *NonPersistentReplicatorStats) SetReplicationBacklog(v int64)`

SetReplicationBacklog sets ReplicationBacklog field to given value.

### HasReplicationBacklog

`func (o *NonPersistentReplicatorStats) HasReplicationBacklog() bool`

HasReplicationBacklog returns a boolean if a field has been set.

### GetConnected

`func (o *NonPersistentReplicatorStats) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *NonPersistentReplicatorStats) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *NonPersistentReplicatorStats) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *NonPersistentReplicatorStats) HasConnected() bool`

HasConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


