# PoolChunkListStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinUsage** | Pointer to **int32** |  | [optional] 
**MaxUsage** | Pointer to **int32** |  | [optional] 
**Chunks** | Pointer to [**[]PoolChunkStats**](PoolChunkStats.md) |  | [optional] 

## Methods

### NewPoolChunkListStats

`func NewPoolChunkListStats() *PoolChunkListStats`

NewPoolChunkListStats instantiates a new PoolChunkListStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolChunkListStatsWithDefaults

`func NewPoolChunkListStatsWithDefaults() *PoolChunkListStats`

NewPoolChunkListStatsWithDefaults instantiates a new PoolChunkListStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinUsage

`func (o *PoolChunkListStats) GetMinUsage() int32`

GetMinUsage returns the MinUsage field if non-nil, zero value otherwise.

### GetMinUsageOk

`func (o *PoolChunkListStats) GetMinUsageOk() (*int32, bool)`

GetMinUsageOk returns a tuple with the MinUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUsage

`func (o *PoolChunkListStats) SetMinUsage(v int32)`

SetMinUsage sets MinUsage field to given value.

### HasMinUsage

`func (o *PoolChunkListStats) HasMinUsage() bool`

HasMinUsage returns a boolean if a field has been set.

### GetMaxUsage

`func (o *PoolChunkListStats) GetMaxUsage() int32`

GetMaxUsage returns the MaxUsage field if non-nil, zero value otherwise.

### GetMaxUsageOk

`func (o *PoolChunkListStats) GetMaxUsageOk() (*int32, bool)`

GetMaxUsageOk returns a tuple with the MaxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsage

`func (o *PoolChunkListStats) SetMaxUsage(v int32)`

SetMaxUsage sets MaxUsage field to given value.

### HasMaxUsage

`func (o *PoolChunkListStats) HasMaxUsage() bool`

HasMaxUsage returns a boolean if a field has been set.

### GetChunks

`func (o *PoolChunkListStats) GetChunks() []PoolChunkStats`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *PoolChunkListStats) GetChunksOk() (*[]PoolChunkStats, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *PoolChunkListStats) SetChunks(v []PoolChunkStats)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *PoolChunkListStats) HasChunks() bool`

HasChunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


