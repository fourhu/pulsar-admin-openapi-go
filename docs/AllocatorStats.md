# AllocatorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDirectArenas** | Pointer to **int32** |  | [optional] 
**NumHeapArenas** | Pointer to **int32** |  | [optional] 
**NumThreadLocalCaches** | Pointer to **int32** |  | [optional] 
**NormalCacheSize** | Pointer to **int32** |  | [optional] 
**SmallCacheSize** | Pointer to **int32** |  | [optional] 
**TinyCacheSize** | Pointer to **int32** |  | [optional] 
**DirectArenas** | Pointer to [**[]PoolArenaStats**](PoolArenaStats.md) |  | [optional] 
**HeapArenas** | Pointer to [**[]PoolArenaStats**](PoolArenaStats.md) |  | [optional] 

## Methods

### NewAllocatorStats

`func NewAllocatorStats() *AllocatorStats`

NewAllocatorStats instantiates a new AllocatorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocatorStatsWithDefaults

`func NewAllocatorStatsWithDefaults() *AllocatorStats`

NewAllocatorStatsWithDefaults instantiates a new AllocatorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDirectArenas

`func (o *AllocatorStats) GetNumDirectArenas() int32`

GetNumDirectArenas returns the NumDirectArenas field if non-nil, zero value otherwise.

### GetNumDirectArenasOk

`func (o *AllocatorStats) GetNumDirectArenasOk() (*int32, bool)`

GetNumDirectArenasOk returns a tuple with the NumDirectArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectArenas

`func (o *AllocatorStats) SetNumDirectArenas(v int32)`

SetNumDirectArenas sets NumDirectArenas field to given value.

### HasNumDirectArenas

`func (o *AllocatorStats) HasNumDirectArenas() bool`

HasNumDirectArenas returns a boolean if a field has been set.

### GetNumHeapArenas

`func (o *AllocatorStats) GetNumHeapArenas() int32`

GetNumHeapArenas returns the NumHeapArenas field if non-nil, zero value otherwise.

### GetNumHeapArenasOk

`func (o *AllocatorStats) GetNumHeapArenasOk() (*int32, bool)`

GetNumHeapArenasOk returns a tuple with the NumHeapArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHeapArenas

`func (o *AllocatorStats) SetNumHeapArenas(v int32)`

SetNumHeapArenas sets NumHeapArenas field to given value.

### HasNumHeapArenas

`func (o *AllocatorStats) HasNumHeapArenas() bool`

HasNumHeapArenas returns a boolean if a field has been set.

### GetNumThreadLocalCaches

`func (o *AllocatorStats) GetNumThreadLocalCaches() int32`

GetNumThreadLocalCaches returns the NumThreadLocalCaches field if non-nil, zero value otherwise.

### GetNumThreadLocalCachesOk

`func (o *AllocatorStats) GetNumThreadLocalCachesOk() (*int32, bool)`

GetNumThreadLocalCachesOk returns a tuple with the NumThreadLocalCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreadLocalCaches

`func (o *AllocatorStats) SetNumThreadLocalCaches(v int32)`

SetNumThreadLocalCaches sets NumThreadLocalCaches field to given value.

### HasNumThreadLocalCaches

`func (o *AllocatorStats) HasNumThreadLocalCaches() bool`

HasNumThreadLocalCaches returns a boolean if a field has been set.

### GetNormalCacheSize

`func (o *AllocatorStats) GetNormalCacheSize() int32`

GetNormalCacheSize returns the NormalCacheSize field if non-nil, zero value otherwise.

### GetNormalCacheSizeOk

`func (o *AllocatorStats) GetNormalCacheSizeOk() (*int32, bool)`

GetNormalCacheSizeOk returns a tuple with the NormalCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalCacheSize

`func (o *AllocatorStats) SetNormalCacheSize(v int32)`

SetNormalCacheSize sets NormalCacheSize field to given value.

### HasNormalCacheSize

`func (o *AllocatorStats) HasNormalCacheSize() bool`

HasNormalCacheSize returns a boolean if a field has been set.

### GetSmallCacheSize

`func (o *AllocatorStats) GetSmallCacheSize() int32`

GetSmallCacheSize returns the SmallCacheSize field if non-nil, zero value otherwise.

### GetSmallCacheSizeOk

`func (o *AllocatorStats) GetSmallCacheSizeOk() (*int32, bool)`

GetSmallCacheSizeOk returns a tuple with the SmallCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallCacheSize

`func (o *AllocatorStats) SetSmallCacheSize(v int32)`

SetSmallCacheSize sets SmallCacheSize field to given value.

### HasSmallCacheSize

`func (o *AllocatorStats) HasSmallCacheSize() bool`

HasSmallCacheSize returns a boolean if a field has been set.

### GetTinyCacheSize

`func (o *AllocatorStats) GetTinyCacheSize() int32`

GetTinyCacheSize returns the TinyCacheSize field if non-nil, zero value otherwise.

### GetTinyCacheSizeOk

`func (o *AllocatorStats) GetTinyCacheSizeOk() (*int32, bool)`

GetTinyCacheSizeOk returns a tuple with the TinyCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinyCacheSize

`func (o *AllocatorStats) SetTinyCacheSize(v int32)`

SetTinyCacheSize sets TinyCacheSize field to given value.

### HasTinyCacheSize

`func (o *AllocatorStats) HasTinyCacheSize() bool`

HasTinyCacheSize returns a boolean if a field has been set.

### GetDirectArenas

`func (o *AllocatorStats) GetDirectArenas() []PoolArenaStats`

GetDirectArenas returns the DirectArenas field if non-nil, zero value otherwise.

### GetDirectArenasOk

`func (o *AllocatorStats) GetDirectArenasOk() (*[]PoolArenaStats, bool)`

GetDirectArenasOk returns a tuple with the DirectArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectArenas

`func (o *AllocatorStats) SetDirectArenas(v []PoolArenaStats)`

SetDirectArenas sets DirectArenas field to given value.

### HasDirectArenas

`func (o *AllocatorStats) HasDirectArenas() bool`

HasDirectArenas returns a boolean if a field has been set.

### GetHeapArenas

`func (o *AllocatorStats) GetHeapArenas() []PoolArenaStats`

GetHeapArenas returns the HeapArenas field if non-nil, zero value otherwise.

### GetHeapArenasOk

`func (o *AllocatorStats) GetHeapArenasOk() (*[]PoolArenaStats, bool)`

GetHeapArenasOk returns a tuple with the HeapArenas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeapArenas

`func (o *AllocatorStats) SetHeapArenas(v []PoolArenaStats)`

SetHeapArenas sets HeapArenas field to given value.

### HasHeapArenas

`func (o *AllocatorStats) HasHeapArenas() bool`

HasHeapArenas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


