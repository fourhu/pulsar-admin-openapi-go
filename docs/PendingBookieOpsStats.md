# PendingBookieOpsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorLedgerCloseOp** | Pointer to **int64** |  | [optional] 
**CursorLedgerCreateOp** | Pointer to **int64** |  | [optional] 
**CursorLedgerDeleteOp** | Pointer to **int64** |  | [optional] 
**CursorLedgerOpenOp** | Pointer to **int64** |  | [optional] 
**DataLedgerCloseOp** | Pointer to **int64** |  | [optional] 
**DataLedgerCreateOp** | Pointer to **int64** |  | [optional] 
**DataLedgerDeleteOp** | Pointer to **int64** |  | [optional] 
**DataLedgerOpenOp** | Pointer to **int64** |  | [optional] 

## Methods

### NewPendingBookieOpsStats

`func NewPendingBookieOpsStats() *PendingBookieOpsStats`

NewPendingBookieOpsStats instantiates a new PendingBookieOpsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingBookieOpsStatsWithDefaults

`func NewPendingBookieOpsStatsWithDefaults() *PendingBookieOpsStats`

NewPendingBookieOpsStatsWithDefaults instantiates a new PendingBookieOpsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorLedgerCloseOp

`func (o *PendingBookieOpsStats) GetCursorLedgerCloseOp() int64`

GetCursorLedgerCloseOp returns the CursorLedgerCloseOp field if non-nil, zero value otherwise.

### GetCursorLedgerCloseOpOk

`func (o *PendingBookieOpsStats) GetCursorLedgerCloseOpOk() (*int64, bool)`

GetCursorLedgerCloseOpOk returns a tuple with the CursorLedgerCloseOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerCloseOp

`func (o *PendingBookieOpsStats) SetCursorLedgerCloseOp(v int64)`

SetCursorLedgerCloseOp sets CursorLedgerCloseOp field to given value.

### HasCursorLedgerCloseOp

`func (o *PendingBookieOpsStats) HasCursorLedgerCloseOp() bool`

HasCursorLedgerCloseOp returns a boolean if a field has been set.

### GetCursorLedgerCreateOp

`func (o *PendingBookieOpsStats) GetCursorLedgerCreateOp() int64`

GetCursorLedgerCreateOp returns the CursorLedgerCreateOp field if non-nil, zero value otherwise.

### GetCursorLedgerCreateOpOk

`func (o *PendingBookieOpsStats) GetCursorLedgerCreateOpOk() (*int64, bool)`

GetCursorLedgerCreateOpOk returns a tuple with the CursorLedgerCreateOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerCreateOp

`func (o *PendingBookieOpsStats) SetCursorLedgerCreateOp(v int64)`

SetCursorLedgerCreateOp sets CursorLedgerCreateOp field to given value.

### HasCursorLedgerCreateOp

`func (o *PendingBookieOpsStats) HasCursorLedgerCreateOp() bool`

HasCursorLedgerCreateOp returns a boolean if a field has been set.

### GetCursorLedgerDeleteOp

`func (o *PendingBookieOpsStats) GetCursorLedgerDeleteOp() int64`

GetCursorLedgerDeleteOp returns the CursorLedgerDeleteOp field if non-nil, zero value otherwise.

### GetCursorLedgerDeleteOpOk

`func (o *PendingBookieOpsStats) GetCursorLedgerDeleteOpOk() (*int64, bool)`

GetCursorLedgerDeleteOpOk returns a tuple with the CursorLedgerDeleteOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerDeleteOp

`func (o *PendingBookieOpsStats) SetCursorLedgerDeleteOp(v int64)`

SetCursorLedgerDeleteOp sets CursorLedgerDeleteOp field to given value.

### HasCursorLedgerDeleteOp

`func (o *PendingBookieOpsStats) HasCursorLedgerDeleteOp() bool`

HasCursorLedgerDeleteOp returns a boolean if a field has been set.

### GetCursorLedgerOpenOp

`func (o *PendingBookieOpsStats) GetCursorLedgerOpenOp() int64`

GetCursorLedgerOpenOp returns the CursorLedgerOpenOp field if non-nil, zero value otherwise.

### GetCursorLedgerOpenOpOk

`func (o *PendingBookieOpsStats) GetCursorLedgerOpenOpOk() (*int64, bool)`

GetCursorLedgerOpenOpOk returns a tuple with the CursorLedgerOpenOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerOpenOp

`func (o *PendingBookieOpsStats) SetCursorLedgerOpenOp(v int64)`

SetCursorLedgerOpenOp sets CursorLedgerOpenOp field to given value.

### HasCursorLedgerOpenOp

`func (o *PendingBookieOpsStats) HasCursorLedgerOpenOp() bool`

HasCursorLedgerOpenOp returns a boolean if a field has been set.

### GetDataLedgerCloseOp

`func (o *PendingBookieOpsStats) GetDataLedgerCloseOp() int64`

GetDataLedgerCloseOp returns the DataLedgerCloseOp field if non-nil, zero value otherwise.

### GetDataLedgerCloseOpOk

`func (o *PendingBookieOpsStats) GetDataLedgerCloseOpOk() (*int64, bool)`

GetDataLedgerCloseOpOk returns a tuple with the DataLedgerCloseOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLedgerCloseOp

`func (o *PendingBookieOpsStats) SetDataLedgerCloseOp(v int64)`

SetDataLedgerCloseOp sets DataLedgerCloseOp field to given value.

### HasDataLedgerCloseOp

`func (o *PendingBookieOpsStats) HasDataLedgerCloseOp() bool`

HasDataLedgerCloseOp returns a boolean if a field has been set.

### GetDataLedgerCreateOp

`func (o *PendingBookieOpsStats) GetDataLedgerCreateOp() int64`

GetDataLedgerCreateOp returns the DataLedgerCreateOp field if non-nil, zero value otherwise.

### GetDataLedgerCreateOpOk

`func (o *PendingBookieOpsStats) GetDataLedgerCreateOpOk() (*int64, bool)`

GetDataLedgerCreateOpOk returns a tuple with the DataLedgerCreateOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLedgerCreateOp

`func (o *PendingBookieOpsStats) SetDataLedgerCreateOp(v int64)`

SetDataLedgerCreateOp sets DataLedgerCreateOp field to given value.

### HasDataLedgerCreateOp

`func (o *PendingBookieOpsStats) HasDataLedgerCreateOp() bool`

HasDataLedgerCreateOp returns a boolean if a field has been set.

### GetDataLedgerDeleteOp

`func (o *PendingBookieOpsStats) GetDataLedgerDeleteOp() int64`

GetDataLedgerDeleteOp returns the DataLedgerDeleteOp field if non-nil, zero value otherwise.

### GetDataLedgerDeleteOpOk

`func (o *PendingBookieOpsStats) GetDataLedgerDeleteOpOk() (*int64, bool)`

GetDataLedgerDeleteOpOk returns a tuple with the DataLedgerDeleteOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLedgerDeleteOp

`func (o *PendingBookieOpsStats) SetDataLedgerDeleteOp(v int64)`

SetDataLedgerDeleteOp sets DataLedgerDeleteOp field to given value.

### HasDataLedgerDeleteOp

`func (o *PendingBookieOpsStats) HasDataLedgerDeleteOp() bool`

HasDataLedgerDeleteOp returns a boolean if a field has been set.

### GetDataLedgerOpenOp

`func (o *PendingBookieOpsStats) GetDataLedgerOpenOp() int64`

GetDataLedgerOpenOp returns the DataLedgerOpenOp field if non-nil, zero value otherwise.

### GetDataLedgerOpenOpOk

`func (o *PendingBookieOpsStats) GetDataLedgerOpenOpOk() (*int64, bool)`

GetDataLedgerOpenOpOk returns a tuple with the DataLedgerOpenOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLedgerOpenOp

`func (o *PendingBookieOpsStats) SetDataLedgerOpenOp(v int64)`

SetDataLedgerOpenOp sets DataLedgerOpenOp field to given value.

### HasDataLedgerOpenOp

`func (o *PendingBookieOpsStats) HasDataLedgerOpenOp() bool`

HasDataLedgerOpenOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


