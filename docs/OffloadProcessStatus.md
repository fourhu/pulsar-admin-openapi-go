# OffloadProcessStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstUnoffloadedMessage** | Pointer to [**MessageIdImpl**](MessageIdImpl.md) |  | [optional] 
**LastError** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOffloadProcessStatus

`func NewOffloadProcessStatus() *OffloadProcessStatus`

NewOffloadProcessStatus instantiates a new OffloadProcessStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffloadProcessStatusWithDefaults

`func NewOffloadProcessStatusWithDefaults() *OffloadProcessStatus`

NewOffloadProcessStatusWithDefaults instantiates a new OffloadProcessStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstUnoffloadedMessage

`func (o *OffloadProcessStatus) GetFirstUnoffloadedMessage() MessageIdImpl`

GetFirstUnoffloadedMessage returns the FirstUnoffloadedMessage field if non-nil, zero value otherwise.

### GetFirstUnoffloadedMessageOk

`func (o *OffloadProcessStatus) GetFirstUnoffloadedMessageOk() (*MessageIdImpl, bool)`

GetFirstUnoffloadedMessageOk returns a tuple with the FirstUnoffloadedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnoffloadedMessage

`func (o *OffloadProcessStatus) SetFirstUnoffloadedMessage(v MessageIdImpl)`

SetFirstUnoffloadedMessage sets FirstUnoffloadedMessage field to given value.

### HasFirstUnoffloadedMessage

`func (o *OffloadProcessStatus) HasFirstUnoffloadedMessage() bool`

HasFirstUnoffloadedMessage returns a boolean if a field has been set.

### GetLastError

`func (o *OffloadProcessStatus) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *OffloadProcessStatus) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *OffloadProcessStatus) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *OffloadProcessStatus) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetStatus

`func (o *OffloadProcessStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OffloadProcessStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OffloadProcessStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OffloadProcessStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


