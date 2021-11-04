# PoolChunkStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChunkSize** | Pointer to **int32** |  | [optional] 
**FreeBytes** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to **int32** |  | [optional] 

## Methods

### NewPoolChunkStats

`func NewPoolChunkStats() *PoolChunkStats`

NewPoolChunkStats instantiates a new PoolChunkStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolChunkStatsWithDefaults

`func NewPoolChunkStatsWithDefaults() *PoolChunkStats`

NewPoolChunkStatsWithDefaults instantiates a new PoolChunkStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunkSize

`func (o *PoolChunkStats) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *PoolChunkStats) GetChunkSizeOk() (*int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *PoolChunkStats) SetChunkSize(v int32)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *PoolChunkStats) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### GetFreeBytes

`func (o *PoolChunkStats) GetFreeBytes() int32`

GetFreeBytes returns the FreeBytes field if non-nil, zero value otherwise.

### GetFreeBytesOk

`func (o *PoolChunkStats) GetFreeBytesOk() (*int32, bool)`

GetFreeBytesOk returns a tuple with the FreeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeBytes

`func (o *PoolChunkStats) SetFreeBytes(v int32)`

SetFreeBytes sets FreeBytes field to given value.

### HasFreeBytes

`func (o *PoolChunkStats) HasFreeBytes() bool`

HasFreeBytes returns a boolean if a field has been set.

### GetUsage

`func (o *PoolChunkStats) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PoolChunkStats) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PoolChunkStats) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PoolChunkStats) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


