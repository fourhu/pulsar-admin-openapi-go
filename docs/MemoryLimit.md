# MemoryLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteValue** | Pointer to **int64** |  | [optional] 
**PercentOfMaxDirectMemory** | Pointer to **float64** |  | [optional] 

## Methods

### NewMemoryLimit

`func NewMemoryLimit() *MemoryLimit`

NewMemoryLimit instantiates a new MemoryLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryLimitWithDefaults

`func NewMemoryLimitWithDefaults() *MemoryLimit`

NewMemoryLimitWithDefaults instantiates a new MemoryLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteValue

`func (o *MemoryLimit) GetAbsoluteValue() int64`

GetAbsoluteValue returns the AbsoluteValue field if non-nil, zero value otherwise.

### GetAbsoluteValueOk

`func (o *MemoryLimit) GetAbsoluteValueOk() (*int64, bool)`

GetAbsoluteValueOk returns a tuple with the AbsoluteValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteValue

`func (o *MemoryLimit) SetAbsoluteValue(v int64)`

SetAbsoluteValue sets AbsoluteValue field to given value.

### HasAbsoluteValue

`func (o *MemoryLimit) HasAbsoluteValue() bool`

HasAbsoluteValue returns a boolean if a field has been set.

### GetPercentOfMaxDirectMemory

`func (o *MemoryLimit) GetPercentOfMaxDirectMemory() float64`

GetPercentOfMaxDirectMemory returns the PercentOfMaxDirectMemory field if non-nil, zero value otherwise.

### GetPercentOfMaxDirectMemoryOk

`func (o *MemoryLimit) GetPercentOfMaxDirectMemoryOk() (*float64, bool)`

GetPercentOfMaxDirectMemoryOk returns a tuple with the PercentOfMaxDirectMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfMaxDirectMemory

`func (o *MemoryLimit) SetPercentOfMaxDirectMemory(v float64)`

SetPercentOfMaxDirectMemory sets PercentOfMaxDirectMemory field to given value.

### HasPercentOfMaxDirectMemory

`func (o *MemoryLimit) HasPercentOfMaxDirectMemory() bool`

HasPercentOfMaxDirectMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


