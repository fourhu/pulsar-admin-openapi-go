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

// PoolChunkListStats struct for PoolChunkListStats
type PoolChunkListStats struct {
	MinUsage *int32 `json:"minUsage,omitempty"`
	MaxUsage *int32 `json:"maxUsage,omitempty"`
	Chunks *[]PoolChunkStats `json:"chunks,omitempty"`
}

// NewPoolChunkListStats instantiates a new PoolChunkListStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolChunkListStats() *PoolChunkListStats {
	this := PoolChunkListStats{}
	return &this
}

// NewPoolChunkListStatsWithDefaults instantiates a new PoolChunkListStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolChunkListStatsWithDefaults() *PoolChunkListStats {
	this := PoolChunkListStats{}
	return &this
}

// GetMinUsage returns the MinUsage field value if set, zero value otherwise.
func (o *PoolChunkListStats) GetMinUsage() int32 {
	if o == nil || o.MinUsage == nil {
		var ret int32
		return ret
	}
	return *o.MinUsage
}

// GetMinUsageOk returns a tuple with the MinUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolChunkListStats) GetMinUsageOk() (*int32, bool) {
	if o == nil || o.MinUsage == nil {
		return nil, false
	}
	return o.MinUsage, true
}

// HasMinUsage returns a boolean if a field has been set.
func (o *PoolChunkListStats) HasMinUsage() bool {
	if o != nil && o.MinUsage != nil {
		return true
	}

	return false
}

// SetMinUsage gets a reference to the given int32 and assigns it to the MinUsage field.
func (o *PoolChunkListStats) SetMinUsage(v int32) {
	o.MinUsage = &v
}

// GetMaxUsage returns the MaxUsage field value if set, zero value otherwise.
func (o *PoolChunkListStats) GetMaxUsage() int32 {
	if o == nil || o.MaxUsage == nil {
		var ret int32
		return ret
	}
	return *o.MaxUsage
}

// GetMaxUsageOk returns a tuple with the MaxUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolChunkListStats) GetMaxUsageOk() (*int32, bool) {
	if o == nil || o.MaxUsage == nil {
		return nil, false
	}
	return o.MaxUsage, true
}

// HasMaxUsage returns a boolean if a field has been set.
func (o *PoolChunkListStats) HasMaxUsage() bool {
	if o != nil && o.MaxUsage != nil {
		return true
	}

	return false
}

// SetMaxUsage gets a reference to the given int32 and assigns it to the MaxUsage field.
func (o *PoolChunkListStats) SetMaxUsage(v int32) {
	o.MaxUsage = &v
}

// GetChunks returns the Chunks field value if set, zero value otherwise.
func (o *PoolChunkListStats) GetChunks() []PoolChunkStats {
	if o == nil || o.Chunks == nil {
		var ret []PoolChunkStats
		return ret
	}
	return *o.Chunks
}

// GetChunksOk returns a tuple with the Chunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolChunkListStats) GetChunksOk() (*[]PoolChunkStats, bool) {
	if o == nil || o.Chunks == nil {
		return nil, false
	}
	return o.Chunks, true
}

// HasChunks returns a boolean if a field has been set.
func (o *PoolChunkListStats) HasChunks() bool {
	if o != nil && o.Chunks != nil {
		return true
	}

	return false
}

// SetChunks gets a reference to the given []PoolChunkStats and assigns it to the Chunks field.
func (o *PoolChunkListStats) SetChunks(v []PoolChunkStats) {
	o.Chunks = &v
}

func (o PoolChunkListStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinUsage != nil {
		toSerialize["minUsage"] = o.MinUsage
	}
	if o.MaxUsage != nil {
		toSerialize["maxUsage"] = o.MaxUsage
	}
	if o.Chunks != nil {
		toSerialize["chunks"] = o.Chunks
	}
	return json.Marshal(toSerialize)
}

type NullablePoolChunkListStats struct {
	value *PoolChunkListStats
	isSet bool
}

func (v NullablePoolChunkListStats) Get() *PoolChunkListStats {
	return v.value
}

func (v *NullablePoolChunkListStats) Set(val *PoolChunkListStats) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolChunkListStats) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolChunkListStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolChunkListStats(val *PoolChunkListStats) *NullablePoolChunkListStats {
	return &NullablePoolChunkListStats{value: val, isSet: true}
}

func (v NullablePoolChunkListStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolChunkListStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


