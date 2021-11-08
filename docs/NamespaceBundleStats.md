# NamespaceBundleStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgRateIn** | Pointer to **float64** |  | [optional] 
**MsgThroughputIn** | Pointer to **float64** |  | [optional] 
**MsgRateOut** | Pointer to **float64** |  | [optional] 
**MsgThroughputOut** | Pointer to **float64** |  | [optional] 
**ConsumerCount** | Pointer to **int32** |  | [optional] 
**ProducerCount** | Pointer to **int32** |  | [optional] 
**Topics** | Pointer to **int64** |  | [optional] 
**CacheSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewNamespaceBundleStats

`func NewNamespaceBundleStats() *NamespaceBundleStats`

NewNamespaceBundleStats instantiates a new NamespaceBundleStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceBundleStatsWithDefaults

`func NewNamespaceBundleStatsWithDefaults() *NamespaceBundleStats`

NewNamespaceBundleStatsWithDefaults instantiates a new NamespaceBundleStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgRateIn

`func (o *NamespaceBundleStats) GetMsgRateIn() float64`

GetMsgRateIn returns the MsgRateIn field if non-nil, zero value otherwise.

### GetMsgRateInOk

`func (o *NamespaceBundleStats) GetMsgRateInOk() (*float64, bool)`

GetMsgRateInOk returns a tuple with the MsgRateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateIn

`func (o *NamespaceBundleStats) SetMsgRateIn(v float64)`

SetMsgRateIn sets MsgRateIn field to given value.

### HasMsgRateIn

`func (o *NamespaceBundleStats) HasMsgRateIn() bool`

HasMsgRateIn returns a boolean if a field has been set.

### GetMsgThroughputIn

`func (o *NamespaceBundleStats) GetMsgThroughputIn() float64`

GetMsgThroughputIn returns the MsgThroughputIn field if non-nil, zero value otherwise.

### GetMsgThroughputInOk

`func (o *NamespaceBundleStats) GetMsgThroughputInOk() (*float64, bool)`

GetMsgThroughputInOk returns a tuple with the MsgThroughputIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputIn

`func (o *NamespaceBundleStats) SetMsgThroughputIn(v float64)`

SetMsgThroughputIn sets MsgThroughputIn field to given value.

### HasMsgThroughputIn

`func (o *NamespaceBundleStats) HasMsgThroughputIn() bool`

HasMsgThroughputIn returns a boolean if a field has been set.

### GetMsgRateOut

`func (o *NamespaceBundleStats) GetMsgRateOut() float64`

GetMsgRateOut returns the MsgRateOut field if non-nil, zero value otherwise.

### GetMsgRateOutOk

`func (o *NamespaceBundleStats) GetMsgRateOutOk() (*float64, bool)`

GetMsgRateOutOk returns a tuple with the MsgRateOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgRateOut

`func (o *NamespaceBundleStats) SetMsgRateOut(v float64)`

SetMsgRateOut sets MsgRateOut field to given value.

### HasMsgRateOut

`func (o *NamespaceBundleStats) HasMsgRateOut() bool`

HasMsgRateOut returns a boolean if a field has been set.

### GetMsgThroughputOut

`func (o *NamespaceBundleStats) GetMsgThroughputOut() float64`

GetMsgThroughputOut returns the MsgThroughputOut field if non-nil, zero value otherwise.

### GetMsgThroughputOutOk

`func (o *NamespaceBundleStats) GetMsgThroughputOutOk() (*float64, bool)`

GetMsgThroughputOutOk returns a tuple with the MsgThroughputOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgThroughputOut

`func (o *NamespaceBundleStats) SetMsgThroughputOut(v float64)`

SetMsgThroughputOut sets MsgThroughputOut field to given value.

### HasMsgThroughputOut

`func (o *NamespaceBundleStats) HasMsgThroughputOut() bool`

HasMsgThroughputOut returns a boolean if a field has been set.

### GetConsumerCount

`func (o *NamespaceBundleStats) GetConsumerCount() int32`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *NamespaceBundleStats) GetConsumerCountOk() (*int32, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *NamespaceBundleStats) SetConsumerCount(v int32)`

SetConsumerCount sets ConsumerCount field to given value.

### HasConsumerCount

`func (o *NamespaceBundleStats) HasConsumerCount() bool`

HasConsumerCount returns a boolean if a field has been set.

### GetProducerCount

`func (o *NamespaceBundleStats) GetProducerCount() int32`

GetProducerCount returns the ProducerCount field if non-nil, zero value otherwise.

### GetProducerCountOk

`func (o *NamespaceBundleStats) GetProducerCountOk() (*int32, bool)`

GetProducerCountOk returns a tuple with the ProducerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerCount

`func (o *NamespaceBundleStats) SetProducerCount(v int32)`

SetProducerCount sets ProducerCount field to given value.

### HasProducerCount

`func (o *NamespaceBundleStats) HasProducerCount() bool`

HasProducerCount returns a boolean if a field has been set.

### GetTopics

`func (o *NamespaceBundleStats) GetTopics() int64`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *NamespaceBundleStats) GetTopicsOk() (*int64, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *NamespaceBundleStats) SetTopics(v int64)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *NamespaceBundleStats) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetCacheSize

`func (o *NamespaceBundleStats) GetCacheSize() int64`

GetCacheSize returns the CacheSize field if non-nil, zero value otherwise.

### GetCacheSizeOk

`func (o *NamespaceBundleStats) GetCacheSizeOk() (*int64, bool)`

GetCacheSizeOk returns a tuple with the CacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSize

`func (o *NamespaceBundleStats) SetCacheSize(v int64)`

SetCacheSize sets CacheSize field to given value.

### HasCacheSize

`func (o *NamespaceBundleStats) HasCacheSize() bool`

HasCacheSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


