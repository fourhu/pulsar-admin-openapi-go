# CompactionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCompactionRemovedEventCount** | Pointer to **int64** |  | [optional] 
**LastCompactionSucceedTimestamp** | Pointer to **int64** |  | [optional] 
**LastCompactionFailedTimestamp** | Pointer to **int64** |  | [optional] 
**LastCompactionDurationTimeInMills** | Pointer to **int64** |  | [optional] 

## Methods

### NewCompactionStats

`func NewCompactionStats() *CompactionStats`

NewCompactionStats instantiates a new CompactionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompactionStatsWithDefaults

`func NewCompactionStatsWithDefaults() *CompactionStats`

NewCompactionStatsWithDefaults instantiates a new CompactionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCompactionRemovedEventCount

`func (o *CompactionStats) GetLastCompactionRemovedEventCount() int64`

GetLastCompactionRemovedEventCount returns the LastCompactionRemovedEventCount field if non-nil, zero value otherwise.

### GetLastCompactionRemovedEventCountOk

`func (o *CompactionStats) GetLastCompactionRemovedEventCountOk() (*int64, bool)`

GetLastCompactionRemovedEventCountOk returns a tuple with the LastCompactionRemovedEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompactionRemovedEventCount

`func (o *CompactionStats) SetLastCompactionRemovedEventCount(v int64)`

SetLastCompactionRemovedEventCount sets LastCompactionRemovedEventCount field to given value.

### HasLastCompactionRemovedEventCount

`func (o *CompactionStats) HasLastCompactionRemovedEventCount() bool`

HasLastCompactionRemovedEventCount returns a boolean if a field has been set.

### GetLastCompactionSucceedTimestamp

`func (o *CompactionStats) GetLastCompactionSucceedTimestamp() int64`

GetLastCompactionSucceedTimestamp returns the LastCompactionSucceedTimestamp field if non-nil, zero value otherwise.

### GetLastCompactionSucceedTimestampOk

`func (o *CompactionStats) GetLastCompactionSucceedTimestampOk() (*int64, bool)`

GetLastCompactionSucceedTimestampOk returns a tuple with the LastCompactionSucceedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompactionSucceedTimestamp

`func (o *CompactionStats) SetLastCompactionSucceedTimestamp(v int64)`

SetLastCompactionSucceedTimestamp sets LastCompactionSucceedTimestamp field to given value.

### HasLastCompactionSucceedTimestamp

`func (o *CompactionStats) HasLastCompactionSucceedTimestamp() bool`

HasLastCompactionSucceedTimestamp returns a boolean if a field has been set.

### GetLastCompactionFailedTimestamp

`func (o *CompactionStats) GetLastCompactionFailedTimestamp() int64`

GetLastCompactionFailedTimestamp returns the LastCompactionFailedTimestamp field if non-nil, zero value otherwise.

### GetLastCompactionFailedTimestampOk

`func (o *CompactionStats) GetLastCompactionFailedTimestampOk() (*int64, bool)`

GetLastCompactionFailedTimestampOk returns a tuple with the LastCompactionFailedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompactionFailedTimestamp

`func (o *CompactionStats) SetLastCompactionFailedTimestamp(v int64)`

SetLastCompactionFailedTimestamp sets LastCompactionFailedTimestamp field to given value.

### HasLastCompactionFailedTimestamp

`func (o *CompactionStats) HasLastCompactionFailedTimestamp() bool`

HasLastCompactionFailedTimestamp returns a boolean if a field has been set.

### GetLastCompactionDurationTimeInMills

`func (o *CompactionStats) GetLastCompactionDurationTimeInMills() int64`

GetLastCompactionDurationTimeInMills returns the LastCompactionDurationTimeInMills field if non-nil, zero value otherwise.

### GetLastCompactionDurationTimeInMillsOk

`func (o *CompactionStats) GetLastCompactionDurationTimeInMillsOk() (*int64, bool)`

GetLastCompactionDurationTimeInMillsOk returns a tuple with the LastCompactionDurationTimeInMills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompactionDurationTimeInMills

`func (o *CompactionStats) SetLastCompactionDurationTimeInMills(v int64)`

SetLastCompactionDurationTimeInMills sets LastCompactionDurationTimeInMills field to given value.

### HasLastCompactionDurationTimeInMills

`func (o *CompactionStats) HasLastCompactionDurationTimeInMills() bool`

HasLastCompactionDurationTimeInMills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


