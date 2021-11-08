# LedgerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedgerId** | Pointer to **int64** |  | [optional] 
**Entries** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Offloaded** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**UnderReplicated** | Pointer to **bool** |  | [optional] 

## Methods

### NewLedgerInfo

`func NewLedgerInfo() *LedgerInfo`

NewLedgerInfo instantiates a new LedgerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerInfoWithDefaults

`func NewLedgerInfoWithDefaults() *LedgerInfo`

NewLedgerInfoWithDefaults instantiates a new LedgerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgerId

`func (o *LedgerInfo) GetLedgerId() int64`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *LedgerInfo) GetLedgerIdOk() (*int64, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *LedgerInfo) SetLedgerId(v int64)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *LedgerInfo) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### GetEntries

`func (o *LedgerInfo) GetEntries() int64`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *LedgerInfo) GetEntriesOk() (*int64, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *LedgerInfo) SetEntries(v int64)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *LedgerInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetSize

`func (o *LedgerInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LedgerInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LedgerInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *LedgerInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetOffloaded

`func (o *LedgerInfo) GetOffloaded() bool`

GetOffloaded returns the Offloaded field if non-nil, zero value otherwise.

### GetOffloadedOk

`func (o *LedgerInfo) GetOffloadedOk() (*bool, bool)`

GetOffloadedOk returns a tuple with the Offloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloaded

`func (o *LedgerInfo) SetOffloaded(v bool)`

SetOffloaded sets Offloaded field to given value.

### HasOffloaded

`func (o *LedgerInfo) HasOffloaded() bool`

HasOffloaded returns a boolean if a field has been set.

### GetMetadata

`func (o *LedgerInfo) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LedgerInfo) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LedgerInfo) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LedgerInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUnderReplicated

`func (o *LedgerInfo) GetUnderReplicated() bool`

GetUnderReplicated returns the UnderReplicated field if non-nil, zero value otherwise.

### GetUnderReplicatedOk

`func (o *LedgerInfo) GetUnderReplicatedOk() (*bool, bool)`

GetUnderReplicatedOk returns a tuple with the UnderReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReplicated

`func (o *LedgerInfo) SetUnderReplicated(v bool)`

SetUnderReplicated sets UnderReplicated field to given value.

### HasUnderReplicated

`func (o *LedgerInfo) HasUnderReplicated() bool`

HasUnderReplicated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


