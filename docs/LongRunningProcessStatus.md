# LongRunningProcessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastError** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewLongRunningProcessStatus

`func NewLongRunningProcessStatus() *LongRunningProcessStatus`

NewLongRunningProcessStatus instantiates a new LongRunningProcessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongRunningProcessStatusWithDefaults

`func NewLongRunningProcessStatusWithDefaults() *LongRunningProcessStatus`

NewLongRunningProcessStatusWithDefaults instantiates a new LongRunningProcessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastError

`func (o *LongRunningProcessStatus) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *LongRunningProcessStatus) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *LongRunningProcessStatus) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *LongRunningProcessStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetStatus

`func (o *LongRunningProcessStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LongRunningProcessStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LongRunningProcessStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LongRunningProcessStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


