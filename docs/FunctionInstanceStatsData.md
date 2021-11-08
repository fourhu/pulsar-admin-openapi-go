# FunctionInstanceStatsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneMin** | Pointer to [**FunctionInstanceStatsDataBase**](FunctionInstanceStatsDataBase.md) |  | [optional] 
**UserMetrics** | Pointer to **map[string]float64** |  | [optional] 
**LastInvocation** | Pointer to **int64** |  | [optional] 
**UserExceptionsTotal** | Pointer to **int64** |  | [optional] 
**AvgProcessLatency** | Pointer to **float64** |  | [optional] 
**ProcessedSuccessfullyTotal** | Pointer to **int64** |  | [optional] 
**ReceivedTotal** | Pointer to **int64** |  | [optional] 
**SystemExceptionsTotal** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunctionInstanceStatsData

`func NewFunctionInstanceStatsData() *FunctionInstanceStatsData`

NewFunctionInstanceStatsData instantiates a new FunctionInstanceStatsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInstanceStatsDataWithDefaults

`func NewFunctionInstanceStatsDataWithDefaults() *FunctionInstanceStatsData`

NewFunctionInstanceStatsDataWithDefaults instantiates a new FunctionInstanceStatsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneMin

`func (o *FunctionInstanceStatsData) GetOneMin() FunctionInstanceStatsDataBase`

GetOneMin returns the OneMin field if non-nil, zero value otherwise.

### GetOneMinOk

`func (o *FunctionInstanceStatsData) GetOneMinOk() (*FunctionInstanceStatsDataBase, bool)`

GetOneMinOk returns a tuple with the OneMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneMin

`func (o *FunctionInstanceStatsData) SetOneMin(v FunctionInstanceStatsDataBase)`

SetOneMin sets OneMin field to given value.

### HasOneMin

`func (o *FunctionInstanceStatsData) HasOneMin() bool`

HasOneMin returns a boolean if a field has been set.

### GetUserMetrics

`func (o *FunctionInstanceStatsData) GetUserMetrics() map[string]float64`

GetUserMetrics returns the UserMetrics field if non-nil, zero value otherwise.

### GetUserMetricsOk

`func (o *FunctionInstanceStatsData) GetUserMetricsOk() (*map[string]float64, bool)`

GetUserMetricsOk returns a tuple with the UserMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetrics

`func (o *FunctionInstanceStatsData) SetUserMetrics(v map[string]float64)`

SetUserMetrics sets UserMetrics field to given value.

### HasUserMetrics

`func (o *FunctionInstanceStatsData) HasUserMetrics() bool`

HasUserMetrics returns a boolean if a field has been set.

### GetLastInvocation

`func (o *FunctionInstanceStatsData) GetLastInvocation() int64`

GetLastInvocation returns the LastInvocation field if non-nil, zero value otherwise.

### GetLastInvocationOk

`func (o *FunctionInstanceStatsData) GetLastInvocationOk() (*int64, bool)`

GetLastInvocationOk returns a tuple with the LastInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInvocation

`func (o *FunctionInstanceStatsData) SetLastInvocation(v int64)`

SetLastInvocation sets LastInvocation field to given value.

### HasLastInvocation

`func (o *FunctionInstanceStatsData) HasLastInvocation() bool`

HasLastInvocation returns a boolean if a field has been set.

### GetUserExceptionsTotal

`func (o *FunctionInstanceStatsData) GetUserExceptionsTotal() int64`

GetUserExceptionsTotal returns the UserExceptionsTotal field if non-nil, zero value otherwise.

### GetUserExceptionsTotalOk

`func (o *FunctionInstanceStatsData) GetUserExceptionsTotalOk() (*int64, bool)`

GetUserExceptionsTotalOk returns a tuple with the UserExceptionsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExceptionsTotal

`func (o *FunctionInstanceStatsData) SetUserExceptionsTotal(v int64)`

SetUserExceptionsTotal sets UserExceptionsTotal field to given value.

### HasUserExceptionsTotal

`func (o *FunctionInstanceStatsData) HasUserExceptionsTotal() bool`

HasUserExceptionsTotal returns a boolean if a field has been set.

### GetAvgProcessLatency

`func (o *FunctionInstanceStatsData) GetAvgProcessLatency() float64`

GetAvgProcessLatency returns the AvgProcessLatency field if non-nil, zero value otherwise.

### GetAvgProcessLatencyOk

`func (o *FunctionInstanceStatsData) GetAvgProcessLatencyOk() (*float64, bool)`

GetAvgProcessLatencyOk returns a tuple with the AvgProcessLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgProcessLatency

`func (o *FunctionInstanceStatsData) SetAvgProcessLatency(v float64)`

SetAvgProcessLatency sets AvgProcessLatency field to given value.

### HasAvgProcessLatency

`func (o *FunctionInstanceStatsData) HasAvgProcessLatency() bool`

HasAvgProcessLatency returns a boolean if a field has been set.

### GetProcessedSuccessfullyTotal

`func (o *FunctionInstanceStatsData) GetProcessedSuccessfullyTotal() int64`

GetProcessedSuccessfullyTotal returns the ProcessedSuccessfullyTotal field if non-nil, zero value otherwise.

### GetProcessedSuccessfullyTotalOk

`func (o *FunctionInstanceStatsData) GetProcessedSuccessfullyTotalOk() (*int64, bool)`

GetProcessedSuccessfullyTotalOk returns a tuple with the ProcessedSuccessfullyTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedSuccessfullyTotal

`func (o *FunctionInstanceStatsData) SetProcessedSuccessfullyTotal(v int64)`

SetProcessedSuccessfullyTotal sets ProcessedSuccessfullyTotal field to given value.

### HasProcessedSuccessfullyTotal

`func (o *FunctionInstanceStatsData) HasProcessedSuccessfullyTotal() bool`

HasProcessedSuccessfullyTotal returns a boolean if a field has been set.

### GetReceivedTotal

`func (o *FunctionInstanceStatsData) GetReceivedTotal() int64`

GetReceivedTotal returns the ReceivedTotal field if non-nil, zero value otherwise.

### GetReceivedTotalOk

`func (o *FunctionInstanceStatsData) GetReceivedTotalOk() (*int64, bool)`

GetReceivedTotalOk returns a tuple with the ReceivedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedTotal

`func (o *FunctionInstanceStatsData) SetReceivedTotal(v int64)`

SetReceivedTotal sets ReceivedTotal field to given value.

### HasReceivedTotal

`func (o *FunctionInstanceStatsData) HasReceivedTotal() bool`

HasReceivedTotal returns a boolean if a field has been set.

### GetSystemExceptionsTotal

`func (o *FunctionInstanceStatsData) GetSystemExceptionsTotal() int64`

GetSystemExceptionsTotal returns the SystemExceptionsTotal field if non-nil, zero value otherwise.

### GetSystemExceptionsTotalOk

`func (o *FunctionInstanceStatsData) GetSystemExceptionsTotalOk() (*int64, bool)`

GetSystemExceptionsTotalOk returns a tuple with the SystemExceptionsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemExceptionsTotal

`func (o *FunctionInstanceStatsData) SetSystemExceptionsTotal(v int64)`

SetSystemExceptionsTotal sets SystemExceptionsTotal field to given value.

### HasSystemExceptionsTotal

`func (o *FunctionInstanceStatsData) HasSystemExceptionsTotal() bool`

HasSystemExceptionsTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


