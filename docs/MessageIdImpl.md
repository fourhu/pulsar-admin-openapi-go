# MessageIdImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedgerId** | Pointer to **int64** |  | [optional] 
**EntryId** | Pointer to **int64** |  | [optional] 
**PartitionIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewMessageIdImpl

`func NewMessageIdImpl() *MessageIdImpl`

NewMessageIdImpl instantiates a new MessageIdImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageIdImplWithDefaults

`func NewMessageIdImplWithDefaults() *MessageIdImpl`

NewMessageIdImplWithDefaults instantiates a new MessageIdImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgerId

`func (o *MessageIdImpl) GetLedgerId() int64`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *MessageIdImpl) GetLedgerIdOk() (*int64, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *MessageIdImpl) SetLedgerId(v int64)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *MessageIdImpl) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### GetEntryId

`func (o *MessageIdImpl) GetEntryId() int64`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MessageIdImpl) GetEntryIdOk() (*int64, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MessageIdImpl) SetEntryId(v int64)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *MessageIdImpl) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetPartitionIndex

`func (o *MessageIdImpl) GetPartitionIndex() int32`

GetPartitionIndex returns the PartitionIndex field if non-nil, zero value otherwise.

### GetPartitionIndexOk

`func (o *MessageIdImpl) GetPartitionIndexOk() (*int32, bool)`

GetPartitionIndexOk returns a tuple with the PartitionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionIndex

`func (o *MessageIdImpl) SetPartitionIndex(v int32)`

SetPartitionIndex sets PartitionIndex field to given value.

### HasPartitionIndex

`func (o *MessageIdImpl) HasPartitionIndex() bool`

HasPartitionIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


