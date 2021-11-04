# CursorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorBacklog** | Pointer to **int64** |  | [optional] 
**CursorLedgerId** | Pointer to **int64** |  | [optional] 

## Methods

### NewCursorDetails

`func NewCursorDetails() *CursorDetails`

NewCursorDetails instantiates a new CursorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorDetailsWithDefaults

`func NewCursorDetailsWithDefaults() *CursorDetails`

NewCursorDetailsWithDefaults instantiates a new CursorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorBacklog

`func (o *CursorDetails) GetCursorBacklog() int64`

GetCursorBacklog returns the CursorBacklog field if non-nil, zero value otherwise.

### GetCursorBacklogOk

`func (o *CursorDetails) GetCursorBacklogOk() (*int64, bool)`

GetCursorBacklogOk returns a tuple with the CursorBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorBacklog

`func (o *CursorDetails) SetCursorBacklog(v int64)`

SetCursorBacklog sets CursorBacklog field to given value.

### HasCursorBacklog

`func (o *CursorDetails) HasCursorBacklog() bool`

HasCursorBacklog returns a boolean if a field has been set.

### GetCursorLedgerId

`func (o *CursorDetails) GetCursorLedgerId() int64`

GetCursorLedgerId returns the CursorLedgerId field if non-nil, zero value otherwise.

### GetCursorLedgerIdOk

`func (o *CursorDetails) GetCursorLedgerIdOk() (*int64, bool)`

GetCursorLedgerIdOk returns a tuple with the CursorLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorLedgerId

`func (o *CursorDetails) SetCursorLedgerId(v int64)`

SetCursorLedgerId sets CursorLedgerId field to given value.

### HasCursorLedgerId

`func (o *CursorDetails) HasCursorLedgerId() bool`

HasCursorLedgerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


