# WorkerFunctionInstanceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**FunctionInstanceStatsData**](FunctionInstanceStatsData.md) |  | [optional] 

## Methods

### NewWorkerFunctionInstanceStats

`func NewWorkerFunctionInstanceStats() *WorkerFunctionInstanceStats`

NewWorkerFunctionInstanceStats instantiates a new WorkerFunctionInstanceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerFunctionInstanceStatsWithDefaults

`func NewWorkerFunctionInstanceStatsWithDefaults() *WorkerFunctionInstanceStats`

NewWorkerFunctionInstanceStatsWithDefaults instantiates a new WorkerFunctionInstanceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkerFunctionInstanceStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkerFunctionInstanceStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkerFunctionInstanceStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkerFunctionInstanceStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetrics

`func (o *WorkerFunctionInstanceStats) GetMetrics() FunctionInstanceStatsData`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WorkerFunctionInstanceStats) GetMetricsOk() (*FunctionInstanceStatsData, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WorkerFunctionInstanceStats) SetMetrics(v FunctionInstanceStatsData)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *WorkerFunctionInstanceStats) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


