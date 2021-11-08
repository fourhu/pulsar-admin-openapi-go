# PoolArenaStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumTinySubpages** | Pointer to **int32** |  | [optional] 
**NumSmallSubpages** | Pointer to **int32** |  | [optional] 
**NumChunkLists** | Pointer to **int32** |  | [optional] 
**TinySubpages** | Pointer to [**[]PoolSubpageStats**](PoolSubpageStats.md) |  | [optional] 
**SmallSubpages** | Pointer to [**[]PoolSubpageStats**](PoolSubpageStats.md) |  | [optional] 
**ChunkLists** | Pointer to [**[]PoolChunkListStats**](PoolChunkListStats.md) |  | [optional] 
**NumAllocations** | Pointer to **int64** |  | [optional] 
**NumTinyAllocations** | Pointer to **int64** |  | [optional] 
**NumSmallAllocations** | Pointer to **int64** |  | [optional] 
**NumNormalAllocations** | Pointer to **int64** |  | [optional] 
**NumHugeAllocations** | Pointer to **int64** |  | [optional] 
**NumDeallocations** | Pointer to **int64** |  | [optional] 
**NumTinyDeallocations** | Pointer to **int64** |  | [optional] 
**NumSmallDeallocations** | Pointer to **int64** |  | [optional] 
**NumNormalDeallocations** | Pointer to **int64** |  | [optional] 
**NumHugeDeallocations** | Pointer to **int64** |  | [optional] 
**NumActiveAllocations** | Pointer to **int64** |  | [optional] 
**NumActiveTinyAllocations** | Pointer to **int64** |  | [optional] 
**NumActiveSmallAllocations** | Pointer to **int64** |  | [optional] 
**NumActiveNormalAllocations** | Pointer to **int64** |  | [optional] 
**NumActiveHugeAllocations** | Pointer to **int64** |  | [optional] 

## Methods

### NewPoolArenaStats

`func NewPoolArenaStats() *PoolArenaStats`

NewPoolArenaStats instantiates a new PoolArenaStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolArenaStatsWithDefaults

`func NewPoolArenaStatsWithDefaults() *PoolArenaStats`

NewPoolArenaStatsWithDefaults instantiates a new PoolArenaStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumTinySubpages

`func (o *PoolArenaStats) GetNumTinySubpages() int32`

GetNumTinySubpages returns the NumTinySubpages field if non-nil, zero value otherwise.

### GetNumTinySubpagesOk

`func (o *PoolArenaStats) GetNumTinySubpagesOk() (*int32, bool)`

GetNumTinySubpagesOk returns a tuple with the NumTinySubpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTinySubpages

`func (o *PoolArenaStats) SetNumTinySubpages(v int32)`

SetNumTinySubpages sets NumTinySubpages field to given value.

### HasNumTinySubpages

`func (o *PoolArenaStats) HasNumTinySubpages() bool`

HasNumTinySubpages returns a boolean if a field has been set.

### GetNumSmallSubpages

`func (o *PoolArenaStats) GetNumSmallSubpages() int32`

GetNumSmallSubpages returns the NumSmallSubpages field if non-nil, zero value otherwise.

### GetNumSmallSubpagesOk

`func (o *PoolArenaStats) GetNumSmallSubpagesOk() (*int32, bool)`

GetNumSmallSubpagesOk returns a tuple with the NumSmallSubpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSmallSubpages

`func (o *PoolArenaStats) SetNumSmallSubpages(v int32)`

SetNumSmallSubpages sets NumSmallSubpages field to given value.

### HasNumSmallSubpages

`func (o *PoolArenaStats) HasNumSmallSubpages() bool`

HasNumSmallSubpages returns a boolean if a field has been set.

### GetNumChunkLists

`func (o *PoolArenaStats) GetNumChunkLists() int32`

GetNumChunkLists returns the NumChunkLists field if non-nil, zero value otherwise.

### GetNumChunkListsOk

`func (o *PoolArenaStats) GetNumChunkListsOk() (*int32, bool)`

GetNumChunkListsOk returns a tuple with the NumChunkLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumChunkLists

`func (o *PoolArenaStats) SetNumChunkLists(v int32)`

SetNumChunkLists sets NumChunkLists field to given value.

### HasNumChunkLists

`func (o *PoolArenaStats) HasNumChunkLists() bool`

HasNumChunkLists returns a boolean if a field has been set.

### GetTinySubpages

`func (o *PoolArenaStats) GetTinySubpages() []PoolSubpageStats`

GetTinySubpages returns the TinySubpages field if non-nil, zero value otherwise.

### GetTinySubpagesOk

`func (o *PoolArenaStats) GetTinySubpagesOk() (*[]PoolSubpageStats, bool)`

GetTinySubpagesOk returns a tuple with the TinySubpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinySubpages

`func (o *PoolArenaStats) SetTinySubpages(v []PoolSubpageStats)`

SetTinySubpages sets TinySubpages field to given value.

### HasTinySubpages

`func (o *PoolArenaStats) HasTinySubpages() bool`

HasTinySubpages returns a boolean if a field has been set.

### GetSmallSubpages

`func (o *PoolArenaStats) GetSmallSubpages() []PoolSubpageStats`

GetSmallSubpages returns the SmallSubpages field if non-nil, zero value otherwise.

### GetSmallSubpagesOk

`func (o *PoolArenaStats) GetSmallSubpagesOk() (*[]PoolSubpageStats, bool)`

GetSmallSubpagesOk returns a tuple with the SmallSubpages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallSubpages

`func (o *PoolArenaStats) SetSmallSubpages(v []PoolSubpageStats)`

SetSmallSubpages sets SmallSubpages field to given value.

### HasSmallSubpages

`func (o *PoolArenaStats) HasSmallSubpages() bool`

HasSmallSubpages returns a boolean if a field has been set.

### GetChunkLists

`func (o *PoolArenaStats) GetChunkLists() []PoolChunkListStats`

GetChunkLists returns the ChunkLists field if non-nil, zero value otherwise.

### GetChunkListsOk

`func (o *PoolArenaStats) GetChunkListsOk() (*[]PoolChunkListStats, bool)`

GetChunkListsOk returns a tuple with the ChunkLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkLists

`func (o *PoolArenaStats) SetChunkLists(v []PoolChunkListStats)`

SetChunkLists sets ChunkLists field to given value.

### HasChunkLists

`func (o *PoolArenaStats) HasChunkLists() bool`

HasChunkLists returns a boolean if a field has been set.

### GetNumAllocations

`func (o *PoolArenaStats) GetNumAllocations() int64`

GetNumAllocations returns the NumAllocations field if non-nil, zero value otherwise.

### GetNumAllocationsOk

`func (o *PoolArenaStats) GetNumAllocationsOk() (*int64, bool)`

GetNumAllocationsOk returns a tuple with the NumAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAllocations

`func (o *PoolArenaStats) SetNumAllocations(v int64)`

SetNumAllocations sets NumAllocations field to given value.

### HasNumAllocations

`func (o *PoolArenaStats) HasNumAllocations() bool`

HasNumAllocations returns a boolean if a field has been set.

### GetNumTinyAllocations

`func (o *PoolArenaStats) GetNumTinyAllocations() int64`

GetNumTinyAllocations returns the NumTinyAllocations field if non-nil, zero value otherwise.

### GetNumTinyAllocationsOk

`func (o *PoolArenaStats) GetNumTinyAllocationsOk() (*int64, bool)`

GetNumTinyAllocationsOk returns a tuple with the NumTinyAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTinyAllocations

`func (o *PoolArenaStats) SetNumTinyAllocations(v int64)`

SetNumTinyAllocations sets NumTinyAllocations field to given value.

### HasNumTinyAllocations

`func (o *PoolArenaStats) HasNumTinyAllocations() bool`

HasNumTinyAllocations returns a boolean if a field has been set.

### GetNumSmallAllocations

`func (o *PoolArenaStats) GetNumSmallAllocations() int64`

GetNumSmallAllocations returns the NumSmallAllocations field if non-nil, zero value otherwise.

### GetNumSmallAllocationsOk

`func (o *PoolArenaStats) GetNumSmallAllocationsOk() (*int64, bool)`

GetNumSmallAllocationsOk returns a tuple with the NumSmallAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSmallAllocations

`func (o *PoolArenaStats) SetNumSmallAllocations(v int64)`

SetNumSmallAllocations sets NumSmallAllocations field to given value.

### HasNumSmallAllocations

`func (o *PoolArenaStats) HasNumSmallAllocations() bool`

HasNumSmallAllocations returns a boolean if a field has been set.

### GetNumNormalAllocations

`func (o *PoolArenaStats) GetNumNormalAllocations() int64`

GetNumNormalAllocations returns the NumNormalAllocations field if non-nil, zero value otherwise.

### GetNumNormalAllocationsOk

`func (o *PoolArenaStats) GetNumNormalAllocationsOk() (*int64, bool)`

GetNumNormalAllocationsOk returns a tuple with the NumNormalAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNormalAllocations

`func (o *PoolArenaStats) SetNumNormalAllocations(v int64)`

SetNumNormalAllocations sets NumNormalAllocations field to given value.

### HasNumNormalAllocations

`func (o *PoolArenaStats) HasNumNormalAllocations() bool`

HasNumNormalAllocations returns a boolean if a field has been set.

### GetNumHugeAllocations

`func (o *PoolArenaStats) GetNumHugeAllocations() int64`

GetNumHugeAllocations returns the NumHugeAllocations field if non-nil, zero value otherwise.

### GetNumHugeAllocationsOk

`func (o *PoolArenaStats) GetNumHugeAllocationsOk() (*int64, bool)`

GetNumHugeAllocationsOk returns a tuple with the NumHugeAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHugeAllocations

`func (o *PoolArenaStats) SetNumHugeAllocations(v int64)`

SetNumHugeAllocations sets NumHugeAllocations field to given value.

### HasNumHugeAllocations

`func (o *PoolArenaStats) HasNumHugeAllocations() bool`

HasNumHugeAllocations returns a boolean if a field has been set.

### GetNumDeallocations

`func (o *PoolArenaStats) GetNumDeallocations() int64`

GetNumDeallocations returns the NumDeallocations field if non-nil, zero value otherwise.

### GetNumDeallocationsOk

`func (o *PoolArenaStats) GetNumDeallocationsOk() (*int64, bool)`

GetNumDeallocationsOk returns a tuple with the NumDeallocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeallocations

`func (o *PoolArenaStats) SetNumDeallocations(v int64)`

SetNumDeallocations sets NumDeallocations field to given value.

### HasNumDeallocations

`func (o *PoolArenaStats) HasNumDeallocations() bool`

HasNumDeallocations returns a boolean if a field has been set.

### GetNumTinyDeallocations

`func (o *PoolArenaStats) GetNumTinyDeallocations() int64`

GetNumTinyDeallocations returns the NumTinyDeallocations field if non-nil, zero value otherwise.

### GetNumTinyDeallocationsOk

`func (o *PoolArenaStats) GetNumTinyDeallocationsOk() (*int64, bool)`

GetNumTinyDeallocationsOk returns a tuple with the NumTinyDeallocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTinyDeallocations

`func (o *PoolArenaStats) SetNumTinyDeallocations(v int64)`

SetNumTinyDeallocations sets NumTinyDeallocations field to given value.

### HasNumTinyDeallocations

`func (o *PoolArenaStats) HasNumTinyDeallocations() bool`

HasNumTinyDeallocations returns a boolean if a field has been set.

### GetNumSmallDeallocations

`func (o *PoolArenaStats) GetNumSmallDeallocations() int64`

GetNumSmallDeallocations returns the NumSmallDeallocations field if non-nil, zero value otherwise.

### GetNumSmallDeallocationsOk

`func (o *PoolArenaStats) GetNumSmallDeallocationsOk() (*int64, bool)`

GetNumSmallDeallocationsOk returns a tuple with the NumSmallDeallocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSmallDeallocations

`func (o *PoolArenaStats) SetNumSmallDeallocations(v int64)`

SetNumSmallDeallocations sets NumSmallDeallocations field to given value.

### HasNumSmallDeallocations

`func (o *PoolArenaStats) HasNumSmallDeallocations() bool`

HasNumSmallDeallocations returns a boolean if a field has been set.

### GetNumNormalDeallocations

`func (o *PoolArenaStats) GetNumNormalDeallocations() int64`

GetNumNormalDeallocations returns the NumNormalDeallocations field if non-nil, zero value otherwise.

### GetNumNormalDeallocationsOk

`func (o *PoolArenaStats) GetNumNormalDeallocationsOk() (*int64, bool)`

GetNumNormalDeallocationsOk returns a tuple with the NumNormalDeallocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNormalDeallocations

`func (o *PoolArenaStats) SetNumNormalDeallocations(v int64)`

SetNumNormalDeallocations sets NumNormalDeallocations field to given value.

### HasNumNormalDeallocations

`func (o *PoolArenaStats) HasNumNormalDeallocations() bool`

HasNumNormalDeallocations returns a boolean if a field has been set.

### GetNumHugeDeallocations

`func (o *PoolArenaStats) GetNumHugeDeallocations() int64`

GetNumHugeDeallocations returns the NumHugeDeallocations field if non-nil, zero value otherwise.

### GetNumHugeDeallocationsOk

`func (o *PoolArenaStats) GetNumHugeDeallocationsOk() (*int64, bool)`

GetNumHugeDeallocationsOk returns a tuple with the NumHugeDeallocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHugeDeallocations

`func (o *PoolArenaStats) SetNumHugeDeallocations(v int64)`

SetNumHugeDeallocations sets NumHugeDeallocations field to given value.

### HasNumHugeDeallocations

`func (o *PoolArenaStats) HasNumHugeDeallocations() bool`

HasNumHugeDeallocations returns a boolean if a field has been set.

### GetNumActiveAllocations

`func (o *PoolArenaStats) GetNumActiveAllocations() int64`

GetNumActiveAllocations returns the NumActiveAllocations field if non-nil, zero value otherwise.

### GetNumActiveAllocationsOk

`func (o *PoolArenaStats) GetNumActiveAllocationsOk() (*int64, bool)`

GetNumActiveAllocationsOk returns a tuple with the NumActiveAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveAllocations

`func (o *PoolArenaStats) SetNumActiveAllocations(v int64)`

SetNumActiveAllocations sets NumActiveAllocations field to given value.

### HasNumActiveAllocations

`func (o *PoolArenaStats) HasNumActiveAllocations() bool`

HasNumActiveAllocations returns a boolean if a field has been set.

### GetNumActiveTinyAllocations

`func (o *PoolArenaStats) GetNumActiveTinyAllocations() int64`

GetNumActiveTinyAllocations returns the NumActiveTinyAllocations field if non-nil, zero value otherwise.

### GetNumActiveTinyAllocationsOk

`func (o *PoolArenaStats) GetNumActiveTinyAllocationsOk() (*int64, bool)`

GetNumActiveTinyAllocationsOk returns a tuple with the NumActiveTinyAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveTinyAllocations

`func (o *PoolArenaStats) SetNumActiveTinyAllocations(v int64)`

SetNumActiveTinyAllocations sets NumActiveTinyAllocations field to given value.

### HasNumActiveTinyAllocations

`func (o *PoolArenaStats) HasNumActiveTinyAllocations() bool`

HasNumActiveTinyAllocations returns a boolean if a field has been set.

### GetNumActiveSmallAllocations

`func (o *PoolArenaStats) GetNumActiveSmallAllocations() int64`

GetNumActiveSmallAllocations returns the NumActiveSmallAllocations field if non-nil, zero value otherwise.

### GetNumActiveSmallAllocationsOk

`func (o *PoolArenaStats) GetNumActiveSmallAllocationsOk() (*int64, bool)`

GetNumActiveSmallAllocationsOk returns a tuple with the NumActiveSmallAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveSmallAllocations

`func (o *PoolArenaStats) SetNumActiveSmallAllocations(v int64)`

SetNumActiveSmallAllocations sets NumActiveSmallAllocations field to given value.

### HasNumActiveSmallAllocations

`func (o *PoolArenaStats) HasNumActiveSmallAllocations() bool`

HasNumActiveSmallAllocations returns a boolean if a field has been set.

### GetNumActiveNormalAllocations

`func (o *PoolArenaStats) GetNumActiveNormalAllocations() int64`

GetNumActiveNormalAllocations returns the NumActiveNormalAllocations field if non-nil, zero value otherwise.

### GetNumActiveNormalAllocationsOk

`func (o *PoolArenaStats) GetNumActiveNormalAllocationsOk() (*int64, bool)`

GetNumActiveNormalAllocationsOk returns a tuple with the NumActiveNormalAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveNormalAllocations

`func (o *PoolArenaStats) SetNumActiveNormalAllocations(v int64)`

SetNumActiveNormalAllocations sets NumActiveNormalAllocations field to given value.

### HasNumActiveNormalAllocations

`func (o *PoolArenaStats) HasNumActiveNormalAllocations() bool`

HasNumActiveNormalAllocations returns a boolean if a field has been set.

### GetNumActiveHugeAllocations

`func (o *PoolArenaStats) GetNumActiveHugeAllocations() int64`

GetNumActiveHugeAllocations returns the NumActiveHugeAllocations field if non-nil, zero value otherwise.

### GetNumActiveHugeAllocationsOk

`func (o *PoolArenaStats) GetNumActiveHugeAllocationsOk() (*int64, bool)`

GetNumActiveHugeAllocationsOk returns a tuple with the NumActiveHugeAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumActiveHugeAllocations

`func (o *PoolArenaStats) SetNumActiveHugeAllocations(v int64)`

SetNumActiveHugeAllocations sets NumActiveHugeAllocations field to given value.

### HasNumActiveHugeAllocations

`func (o *PoolArenaStats) HasNumActiveHugeAllocations() bool`

HasNumActiveHugeAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


