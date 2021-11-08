/*
Pulsar Admin REST API

This provides the REST API for admin operations

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FunctionInstanceStatsData struct for FunctionInstanceStatsData
type FunctionInstanceStatsData struct {
	OneMin *FunctionInstanceStatsDataBase `json:"oneMin,omitempty"`
	UserMetrics *map[string]float64 `json:"userMetrics,omitempty"`
	LastInvocation *int64 `json:"lastInvocation,omitempty"`
	UserExceptionsTotal *int64 `json:"userExceptionsTotal,omitempty"`
	AvgProcessLatency *float64 `json:"avgProcessLatency,omitempty"`
	ProcessedSuccessfullyTotal *int64 `json:"processedSuccessfullyTotal,omitempty"`
	ReceivedTotal *int64 `json:"receivedTotal,omitempty"`
	SystemExceptionsTotal *int64 `json:"systemExceptionsTotal,omitempty"`
}

// NewFunctionInstanceStatsData instantiates a new FunctionInstanceStatsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionInstanceStatsData() *FunctionInstanceStatsData {
	this := FunctionInstanceStatsData{}
	return &this
}

// NewFunctionInstanceStatsDataWithDefaults instantiates a new FunctionInstanceStatsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionInstanceStatsDataWithDefaults() *FunctionInstanceStatsData {
	this := FunctionInstanceStatsData{}
	return &this
}

// GetOneMin returns the OneMin field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetOneMin() FunctionInstanceStatsDataBase {
	if o == nil || o.OneMin == nil {
		var ret FunctionInstanceStatsDataBase
		return ret
	}
	return *o.OneMin
}

// GetOneMinOk returns a tuple with the OneMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetOneMinOk() (*FunctionInstanceStatsDataBase, bool) {
	if o == nil || o.OneMin == nil {
		return nil, false
	}
	return o.OneMin, true
}

// HasOneMin returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasOneMin() bool {
	if o != nil && o.OneMin != nil {
		return true
	}

	return false
}

// SetOneMin gets a reference to the given FunctionInstanceStatsDataBase and assigns it to the OneMin field.
func (o *FunctionInstanceStatsData) SetOneMin(v FunctionInstanceStatsDataBase) {
	o.OneMin = &v
}

// GetUserMetrics returns the UserMetrics field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetUserMetrics() map[string]float64 {
	if o == nil || o.UserMetrics == nil {
		var ret map[string]float64
		return ret
	}
	return *o.UserMetrics
}

// GetUserMetricsOk returns a tuple with the UserMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetUserMetricsOk() (*map[string]float64, bool) {
	if o == nil || o.UserMetrics == nil {
		return nil, false
	}
	return o.UserMetrics, true
}

// HasUserMetrics returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasUserMetrics() bool {
	if o != nil && o.UserMetrics != nil {
		return true
	}

	return false
}

// SetUserMetrics gets a reference to the given map[string]float64 and assigns it to the UserMetrics field.
func (o *FunctionInstanceStatsData) SetUserMetrics(v map[string]float64) {
	o.UserMetrics = &v
}

// GetLastInvocation returns the LastInvocation field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetLastInvocation() int64 {
	if o == nil || o.LastInvocation == nil {
		var ret int64
		return ret
	}
	return *o.LastInvocation
}

// GetLastInvocationOk returns a tuple with the LastInvocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetLastInvocationOk() (*int64, bool) {
	if o == nil || o.LastInvocation == nil {
		return nil, false
	}
	return o.LastInvocation, true
}

// HasLastInvocation returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasLastInvocation() bool {
	if o != nil && o.LastInvocation != nil {
		return true
	}

	return false
}

// SetLastInvocation gets a reference to the given int64 and assigns it to the LastInvocation field.
func (o *FunctionInstanceStatsData) SetLastInvocation(v int64) {
	o.LastInvocation = &v
}

// GetUserExceptionsTotal returns the UserExceptionsTotal field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetUserExceptionsTotal() int64 {
	if o == nil || o.UserExceptionsTotal == nil {
		var ret int64
		return ret
	}
	return *o.UserExceptionsTotal
}

// GetUserExceptionsTotalOk returns a tuple with the UserExceptionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetUserExceptionsTotalOk() (*int64, bool) {
	if o == nil || o.UserExceptionsTotal == nil {
		return nil, false
	}
	return o.UserExceptionsTotal, true
}

// HasUserExceptionsTotal returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasUserExceptionsTotal() bool {
	if o != nil && o.UserExceptionsTotal != nil {
		return true
	}

	return false
}

// SetUserExceptionsTotal gets a reference to the given int64 and assigns it to the UserExceptionsTotal field.
func (o *FunctionInstanceStatsData) SetUserExceptionsTotal(v int64) {
	o.UserExceptionsTotal = &v
}

// GetAvgProcessLatency returns the AvgProcessLatency field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetAvgProcessLatency() float64 {
	if o == nil || o.AvgProcessLatency == nil {
		var ret float64
		return ret
	}
	return *o.AvgProcessLatency
}

// GetAvgProcessLatencyOk returns a tuple with the AvgProcessLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetAvgProcessLatencyOk() (*float64, bool) {
	if o == nil || o.AvgProcessLatency == nil {
		return nil, false
	}
	return o.AvgProcessLatency, true
}

// HasAvgProcessLatency returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasAvgProcessLatency() bool {
	if o != nil && o.AvgProcessLatency != nil {
		return true
	}

	return false
}

// SetAvgProcessLatency gets a reference to the given float64 and assigns it to the AvgProcessLatency field.
func (o *FunctionInstanceStatsData) SetAvgProcessLatency(v float64) {
	o.AvgProcessLatency = &v
}

// GetProcessedSuccessfullyTotal returns the ProcessedSuccessfullyTotal field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetProcessedSuccessfullyTotal() int64 {
	if o == nil || o.ProcessedSuccessfullyTotal == nil {
		var ret int64
		return ret
	}
	return *o.ProcessedSuccessfullyTotal
}

// GetProcessedSuccessfullyTotalOk returns a tuple with the ProcessedSuccessfullyTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetProcessedSuccessfullyTotalOk() (*int64, bool) {
	if o == nil || o.ProcessedSuccessfullyTotal == nil {
		return nil, false
	}
	return o.ProcessedSuccessfullyTotal, true
}

// HasProcessedSuccessfullyTotal returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasProcessedSuccessfullyTotal() bool {
	if o != nil && o.ProcessedSuccessfullyTotal != nil {
		return true
	}

	return false
}

// SetProcessedSuccessfullyTotal gets a reference to the given int64 and assigns it to the ProcessedSuccessfullyTotal field.
func (o *FunctionInstanceStatsData) SetProcessedSuccessfullyTotal(v int64) {
	o.ProcessedSuccessfullyTotal = &v
}

// GetReceivedTotal returns the ReceivedTotal field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetReceivedTotal() int64 {
	if o == nil || o.ReceivedTotal == nil {
		var ret int64
		return ret
	}
	return *o.ReceivedTotal
}

// GetReceivedTotalOk returns a tuple with the ReceivedTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetReceivedTotalOk() (*int64, bool) {
	if o == nil || o.ReceivedTotal == nil {
		return nil, false
	}
	return o.ReceivedTotal, true
}

// HasReceivedTotal returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasReceivedTotal() bool {
	if o != nil && o.ReceivedTotal != nil {
		return true
	}

	return false
}

// SetReceivedTotal gets a reference to the given int64 and assigns it to the ReceivedTotal field.
func (o *FunctionInstanceStatsData) SetReceivedTotal(v int64) {
	o.ReceivedTotal = &v
}

// GetSystemExceptionsTotal returns the SystemExceptionsTotal field value if set, zero value otherwise.
func (o *FunctionInstanceStatsData) GetSystemExceptionsTotal() int64 {
	if o == nil || o.SystemExceptionsTotal == nil {
		var ret int64
		return ret
	}
	return *o.SystemExceptionsTotal
}

// GetSystemExceptionsTotalOk returns a tuple with the SystemExceptionsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInstanceStatsData) GetSystemExceptionsTotalOk() (*int64, bool) {
	if o == nil || o.SystemExceptionsTotal == nil {
		return nil, false
	}
	return o.SystemExceptionsTotal, true
}

// HasSystemExceptionsTotal returns a boolean if a field has been set.
func (o *FunctionInstanceStatsData) HasSystemExceptionsTotal() bool {
	if o != nil && o.SystemExceptionsTotal != nil {
		return true
	}

	return false
}

// SetSystemExceptionsTotal gets a reference to the given int64 and assigns it to the SystemExceptionsTotal field.
func (o *FunctionInstanceStatsData) SetSystemExceptionsTotal(v int64) {
	o.SystemExceptionsTotal = &v
}

func (o FunctionInstanceStatsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OneMin != nil {
		toSerialize["oneMin"] = o.OneMin
	}
	if o.UserMetrics != nil {
		toSerialize["userMetrics"] = o.UserMetrics
	}
	if o.LastInvocation != nil {
		toSerialize["lastInvocation"] = o.LastInvocation
	}
	if o.UserExceptionsTotal != nil {
		toSerialize["userExceptionsTotal"] = o.UserExceptionsTotal
	}
	if o.AvgProcessLatency != nil {
		toSerialize["avgProcessLatency"] = o.AvgProcessLatency
	}
	if o.ProcessedSuccessfullyTotal != nil {
		toSerialize["processedSuccessfullyTotal"] = o.ProcessedSuccessfullyTotal
	}
	if o.ReceivedTotal != nil {
		toSerialize["receivedTotal"] = o.ReceivedTotal
	}
	if o.SystemExceptionsTotal != nil {
		toSerialize["systemExceptionsTotal"] = o.SystemExceptionsTotal
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionInstanceStatsData struct {
	value *FunctionInstanceStatsData
	isSet bool
}

func (v NullableFunctionInstanceStatsData) Get() *FunctionInstanceStatsData {
	return v.value
}

func (v *NullableFunctionInstanceStatsData) Set(val *FunctionInstanceStatsData) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionInstanceStatsData) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionInstanceStatsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionInstanceStatsData(val *FunctionInstanceStatsData) *NullableFunctionInstanceStatsData {
	return &NullableFunctionInstanceStatsData{value: val, isSet: true}
}

func (v NullableFunctionInstanceStatsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionInstanceStatsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


