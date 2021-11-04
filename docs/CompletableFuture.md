# CompletableFuture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **bool** |  | [optional] 
**CompletedExceptionally** | Pointer to **bool** |  | [optional] 
**NumberOfDependents** | Pointer to **int32** |  | [optional] 
**Cancelled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCompletableFuture

`func NewCompletableFuture() *CompletableFuture`

NewCompletableFuture instantiates a new CompletableFuture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletableFutureWithDefaults

`func NewCompletableFutureWithDefaults() *CompletableFuture`

NewCompletableFutureWithDefaults instantiates a new CompletableFuture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *CompletableFuture) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *CompletableFuture) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *CompletableFuture) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *CompletableFuture) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetCompletedExceptionally

`func (o *CompletableFuture) GetCompletedExceptionally() bool`

GetCompletedExceptionally returns the CompletedExceptionally field if non-nil, zero value otherwise.

### GetCompletedExceptionallyOk

`func (o *CompletableFuture) GetCompletedExceptionallyOk() (*bool, bool)`

GetCompletedExceptionallyOk returns a tuple with the CompletedExceptionally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedExceptionally

`func (o *CompletableFuture) SetCompletedExceptionally(v bool)`

SetCompletedExceptionally sets CompletedExceptionally field to given value.

### HasCompletedExceptionally

`func (o *CompletableFuture) HasCompletedExceptionally() bool`

HasCompletedExceptionally returns a boolean if a field has been set.

### GetNumberOfDependents

`func (o *CompletableFuture) GetNumberOfDependents() int32`

GetNumberOfDependents returns the NumberOfDependents field if non-nil, zero value otherwise.

### GetNumberOfDependentsOk

`func (o *CompletableFuture) GetNumberOfDependentsOk() (*int32, bool)`

GetNumberOfDependentsOk returns a tuple with the NumberOfDependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDependents

`func (o *CompletableFuture) SetNumberOfDependents(v int32)`

SetNumberOfDependents sets NumberOfDependents field to given value.

### HasNumberOfDependents

`func (o *CompletableFuture) HasNumberOfDependents() bool`

HasNumberOfDependents returns a boolean if a field has been set.

### GetCancelled

`func (o *CompletableFuture) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *CompletableFuture) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *CompletableFuture) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *CompletableFuture) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


