# WorkerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | Pointer to **string** |  | [optional] 
**WorkerHostname** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewWorkerInfo

`func NewWorkerInfo() *WorkerInfo`

NewWorkerInfo instantiates a new WorkerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerInfoWithDefaults

`func NewWorkerInfoWithDefaults() *WorkerInfo`

NewWorkerInfoWithDefaults instantiates a new WorkerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *WorkerInfo) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *WorkerInfo) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *WorkerInfo) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *WorkerInfo) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkerHostname

`func (o *WorkerInfo) GetWorkerHostname() string`

GetWorkerHostname returns the WorkerHostname field if non-nil, zero value otherwise.

### GetWorkerHostnameOk

`func (o *WorkerInfo) GetWorkerHostnameOk() (*string, bool)`

GetWorkerHostnameOk returns a tuple with the WorkerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerHostname

`func (o *WorkerInfo) SetWorkerHostname(v string)`

SetWorkerHostname sets WorkerHostname field to given value.

### HasWorkerHostname

`func (o *WorkerInfo) HasWorkerHostname() bool`

HasWorkerHostname returns a boolean if a field has been set.

### GetPort

`func (o *WorkerInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WorkerInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WorkerInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WorkerInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


