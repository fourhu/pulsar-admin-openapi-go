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

// WorkerFunctionInstanceStats struct for WorkerFunctionInstanceStats
type WorkerFunctionInstanceStats struct {
	Name *string `json:"name,omitempty"`
	Metrics *FunctionInstanceStatsData `json:"metrics,omitempty"`
}

// NewWorkerFunctionInstanceStats instantiates a new WorkerFunctionInstanceStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerFunctionInstanceStats() *WorkerFunctionInstanceStats {
	this := WorkerFunctionInstanceStats{}
	return &this
}

// NewWorkerFunctionInstanceStatsWithDefaults instantiates a new WorkerFunctionInstanceStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerFunctionInstanceStatsWithDefaults() *WorkerFunctionInstanceStats {
	this := WorkerFunctionInstanceStats{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkerFunctionInstanceStats) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerFunctionInstanceStats) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkerFunctionInstanceStats) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkerFunctionInstanceStats) SetName(v string) {
	o.Name = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *WorkerFunctionInstanceStats) GetMetrics() FunctionInstanceStatsData {
	if o == nil || o.Metrics == nil {
		var ret FunctionInstanceStatsData
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerFunctionInstanceStats) GetMetricsOk() (*FunctionInstanceStatsData, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *WorkerFunctionInstanceStats) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given FunctionInstanceStatsData and assigns it to the Metrics field.
func (o *WorkerFunctionInstanceStats) SetMetrics(v FunctionInstanceStatsData) {
	o.Metrics = &v
}

func (o WorkerFunctionInstanceStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableWorkerFunctionInstanceStats struct {
	value *WorkerFunctionInstanceStats
	isSet bool
}

func (v NullableWorkerFunctionInstanceStats) Get() *WorkerFunctionInstanceStats {
	return v.value
}

func (v *NullableWorkerFunctionInstanceStats) Set(val *WorkerFunctionInstanceStats) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerFunctionInstanceStats) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerFunctionInstanceStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerFunctionInstanceStats(val *WorkerFunctionInstanceStats) *NullableWorkerFunctionInstanceStats {
	return &NullableWorkerFunctionInstanceStats{value: val, isSet: true}
}

func (v NullableWorkerFunctionInstanceStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerFunctionInstanceStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


