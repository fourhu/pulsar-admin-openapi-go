# ReplicatorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** |  | [optional] 
**InboundConnectedSince** | Pointer to **string** |  | [optional] 
**InboundConnection** | Pointer to **string** |  | [optional] 
**MsgRateExpired** | Pointer to **float64** |  | [optional] 
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**OutboundConnectedSince** | Pointer to **string** |  | [optional] 
**OutboundConnection** | Pointer to **string** |  | [optional] 
**ReplicationBacklog** | Pointer to **int64** |  | [optional] 
**ReplicationDelayInSeconds** | Pointer to **int64** |  | [optional] 

## Methods

### NewReplicatorStats

`func NewReplicatorStats() *ReplicatorStats`

NewReplicatorStats instantiates a new ReplicatorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicatorStatsWithDefaults

`func NewReplicatorStatsWithDefaults() *ReplicatorStats`

NewReplicatorStatsWithDefaults instantiates a new ReplicatorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *ReplicatorStats) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ReplicatorStats) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ReplicatorStats) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ReplicatorStats) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetInboundConnectedSince

`func (o *ReplicatorStats) GetInboundConnectedSince() string`

GetInboundConnectedSince returns the InboundConnectedSince field if non-nil, zero value otherwise.

### GetInboundConnectedSinceOk

`func (o *ReplicatorStats) GetInboundConnectedSinceOk() (*string, bool)`

GetInboundConnectedSinceOk returns a tuple with the InboundConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnectedSince

`func (o *ReplicatorStats) SetInboundConnectedSince(v string)`

SetInboundConnectedSince sets InboundConnectedSince field to given value.

### HasInboundConnectedSince

`func (o *ReplicatorStats) HasInboundConnectedSince() bool`

HasInboundConnectedSince returns a boolean if a field has been set.

### GetInboundConnection

`func (o *ReplicatorStats) GetInboundConnection() string`

GetInboundConnection returns the InboundConnection field if non-nil, zero value otherwise.

### GetInboundConnectionOk

`func (o *ReplicatorStats) GetInboundConnectionOk() (*string, bool)`

GetInboundConnectionOk returns a tuple with the InboundConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnection

`func (o *ReplicatorStats) SetInboundConnection(v string)`

SetInboundConnection sets InboundConnection field to given value.

### HasInboundConnection

`func (o *ReplicatorStats) HasInboundConnection() bool`

HasInboundConnection returns a boolean if a field has been set.

### GetMsgRateExpired

`func (o *ReplicatorStats) GetMsgRateExpired() float64`

GetMsgRateExpired returns the MsgRateExpired field if non-nil, zero value otherwise.

### GetMsgRateExpiredOk

`func (o *ReplicatorStats) GetMsgRateExpiredOk() (*float64, bool)`

GetMsgRateExpiredOk returns a tuple with the MsgRateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateExpired

`func (o *ReplicatorStats) SetMsgRateExpired(v float64)`

SetMsgRateExpired sets MsgRateExpired field to given value.

### HasMsgRateExpired

`func (o *ReplicatorStats) HasMsgRateExpired() bool`

HasMsgRateExpired returns a boolean if a field has been set.

### GetMsgRateIn

`func (o *ReplicatorStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *ReplicatorStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *ReplicatorStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *ReplicatorStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *ReplicatorStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *ReplicatorStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *ReplicatorStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *ReplicatorStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *ReplicatorStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *ReplicatorStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *ReplicatorStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *ReplicatorStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *ReplicatorStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *ReplicatorStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *ReplicatorStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *ReplicatorStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetOutboundConnectedSince

`func (o *ReplicatorStats) GetOutboundConnectedSince() string`

GetOutboundConnectedSince returns the OutboundConnectedSince field if non-nil, zero value otherwise.

### GetOutboundConnectedSinceOk

`func (o *ReplicatorStats) GetOutboundConnectedSinceOk() (*string, bool)`

GetOutboundConnectedSinceOk returns a tuple with the OutboundConnectedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnectedSince

`func (o *ReplicatorStats) SetOutboundConnectedSince(v string)`

SetOutboundConnectedSince sets OutboundConnectedSince field to given value.

### HasOutboundConnectedSince

`func (o *ReplicatorStats) HasOutboundConnectedSince() bool`

HasOutboundConnectedSince returns a boolean if a field has been set.

### GetOutboundConnection

`func (o *ReplicatorStats) GetOutboundConnection() string`

GetOutboundConnection returns the OutboundConnection field if non-nil, zero value otherwise.

### GetOutboundConnectionOk

`func (o *ReplicatorStats) GetOutboundConnectionOk() (*string, bool)`

GetOutboundConnectionOk returns a tuple with the OutboundConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnection

`func (o *ReplicatorStats) SetOutboundConnection(v string)`

SetOutboundConnection sets OutboundConnection field to given value.

### HasOutboundConnection

`func (o *ReplicatorStats) HasOutboundConnection() bool`

HasOutboundConnection returns a boolean if a field has been set.

### GetReplicationBacklog

`func (o *ReplicatorStats) GetReplicationBacklog() int64`

GetReplicationBacklog returns the ReplicationBacklog field if non-nil, zero value otherwise.

### GetReplicationBacklogOk

`func (o *ReplicatorStats) GetReplicationBacklogOk() (*int64, bool)`

GetReplicationBacklogOk returns a tuple with the ReplicationBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBacklog

`func (o *ReplicatorStats) SetReplicationBacklog(v int64)`

SetReplicationBacklog sets ReplicationBacklog field to given value.

### HasReplicationBacklog

`func (o *ReplicatorStats) HasReplicationBacklog() bool`

HasReplicationBacklog returns a boolean if a field has been set.

### GetReplicationDelayInSeconds

`func (o *ReplicatorStats) GetReplicationDelayInSeconds() int64`

GetReplicationDelayInSeconds returns the ReplicationDelayInSeconds field if non-nil, zero value otherwise.

### GetReplicationDelayInSecondsOk

`func (o *ReplicatorStats) GetReplicationDelayInSecondsOk() (*int64, bool)`

GetReplicationDelayInSecondsOk returns a tuple with the ReplicationDelayInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDelayInSeconds

`func (o *ReplicatorStats) SetReplicationDelayInSeconds(v int64)`

SetReplicationDelayInSeconds sets ReplicationDelayInSeconds field to given value.

### HasReplicationDelayInSeconds

`func (o *ReplicatorStats) HasReplicationDelayInSeconds() bool`

HasReplicationDelayInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


