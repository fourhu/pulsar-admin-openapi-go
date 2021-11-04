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

// AllocatorStats struct for AllocatorStats
type AllocatorStats struct {
	DirectArenas *[]PoolArenaStats `json:"directArenas,omitempty"`
	HeapArenas *[]PoolArenaStats `json:"heapArenas,omitempty"`
	NormalCacheSize *int32 `json:"normalCacheSize,omitempty"`
	NumDirectArenas *int32 `json:"numDirectArenas,omitempty"`
	NumHeapArenas *int32 `json:"numHeapArenas,omitempty"`
	NumThreadLocalCaches *int32 `json:"numThreadLocalCaches,omitempty"`
	SmallCacheSize *int32 `json:"smallCacheSize,omitempty"`
	TinyCacheSize *int32 `json:"tinyCacheSize,omitempty"`
}

// NewAllocatorStats instantiates a new AllocatorStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatorStats() *AllocatorStats {
	this := AllocatorStats{}
	return &this
}

// NewAllocatorStatsWithDefaults instantiates a new AllocatorStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatorStatsWithDefaults() *AllocatorStats {
	this := AllocatorStats{}
	return &this
}

// GetDirectArenas returns the DirectArenas field value if set, zero value otherwise.
func (o *AllocatorStats) GetDirectArenas() []PoolArenaStats {
	if o == nil || o.DirectArenas == nil {
		var ret []PoolArenaStats
		return ret
	}
	return *o.DirectArenas
}

// GetDirectArenasOk returns a tuple with the DirectArenas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetDirectArenasOk() (*[]PoolArenaStats, bool) {
	if o == nil || o.DirectArenas == nil {
		return nil, false
	}
	return o.DirectArenas, true
}

// HasDirectArenas returns a boolean if a field has been set.
func (o *AllocatorStats) HasDirectArenas() bool {
	if o != nil && o.DirectArenas != nil {
		return true
	}

	return false
}

// SetDirectArenas gets a reference to the given []PoolArenaStats and assigns it to the DirectArenas field.
func (o *AllocatorStats) SetDirectArenas(v []PoolArenaStats) {
	o.DirectArenas = &v
}

// GetHeapArenas returns the HeapArenas field value if set, zero value otherwise.
func (o *AllocatorStats) GetHeapArenas() []PoolArenaStats {
	if o == nil || o.HeapArenas == nil {
		var ret []PoolArenaStats
		return ret
	}
	return *o.HeapArenas
}

// GetHeapArenasOk returns a tuple with the HeapArenas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetHeapArenasOk() (*[]PoolArenaStats, bool) {
	if o == nil || o.HeapArenas == nil {
		return nil, false
	}
	return o.HeapArenas, true
}

// HasHeapArenas returns a boolean if a field has been set.
func (o *AllocatorStats) HasHeapArenas() bool {
	if o != nil && o.HeapArenas != nil {
		return true
	}

	return false
}

// SetHeapArenas gets a reference to the given []PoolArenaStats and assigns it to the HeapArenas field.
func (o *AllocatorStats) SetHeapArenas(v []PoolArenaStats) {
	o.HeapArenas = &v
}

// GetNormalCacheSize returns the NormalCacheSize field value if set, zero value otherwise.
func (o *AllocatorStats) GetNormalCacheSize() int32 {
	if o == nil || o.NormalCacheSize == nil {
		var ret int32
		return ret
	}
	return *o.NormalCacheSize
}

// GetNormalCacheSizeOk returns a tuple with the NormalCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetNormalCacheSizeOk() (*int32, bool) {
	if o == nil || o.NormalCacheSize == nil {
		return nil, false
	}
	return o.NormalCacheSize, true
}

// HasNormalCacheSize returns a boolean if a field has been set.
func (o *AllocatorStats) HasNormalCacheSize() bool {
	if o != nil && o.NormalCacheSize != nil {
		return true
	}

	return false
}

// SetNormalCacheSize gets a reference to the given int32 and assigns it to the NormalCacheSize field.
func (o *AllocatorStats) SetNormalCacheSize(v int32) {
	o.NormalCacheSize = &v
}

// GetNumDirectArenas returns the NumDirectArenas field value if set, zero value otherwise.
func (o *AllocatorStats) GetNumDirectArenas() int32 {
	if o == nil || o.NumDirectArenas == nil {
		var ret int32
		return ret
	}
	return *o.NumDirectArenas
}

// GetNumDirectArenasOk returns a tuple with the NumDirectArenas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetNumDirectArenasOk() (*int32, bool) {
	if o == nil || o.NumDirectArenas == nil {
		return nil, false
	}
	return o.NumDirectArenas, true
}

// HasNumDirectArenas returns a boolean if a field has been set.
func (o *AllocatorStats) HasNumDirectArenas() bool {
	if o != nil && o.NumDirectArenas != nil {
		return true
	}

	return false
}

// SetNumDirectArenas gets a reference to the given int32 and assigns it to the NumDirectArenas field.
func (o *AllocatorStats) SetNumDirectArenas(v int32) {
	o.NumDirectArenas = &v
}

// GetNumHeapArenas returns the NumHeapArenas field value if set, zero value otherwise.
func (o *AllocatorStats) GetNumHeapArenas() int32 {
	if o == nil || o.NumHeapArenas == nil {
		var ret int32
		return ret
	}
	return *o.NumHeapArenas
}

// GetNumHeapArenasOk returns a tuple with the NumHeapArenas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetNumHeapArenasOk() (*int32, bool) {
	if o == nil || o.NumHeapArenas == nil {
		return nil, false
	}
	return o.NumHeapArenas, true
}

// HasNumHeapArenas returns a boolean if a field has been set.
func (o *AllocatorStats) HasNumHeapArenas() bool {
	if o != nil && o.NumHeapArenas != nil {
		return true
	}

	return false
}

// SetNumHeapArenas gets a reference to the given int32 and assigns it to the NumHeapArenas field.
func (o *AllocatorStats) SetNumHeapArenas(v int32) {
	o.NumHeapArenas = &v
}

// GetNumThreadLocalCaches returns the NumThreadLocalCaches field value if set, zero value otherwise.
func (o *AllocatorStats) GetNumThreadLocalCaches() int32 {
	if o == nil || o.NumThreadLocalCaches == nil {
		var ret int32
		return ret
	}
	return *o.NumThreadLocalCaches
}

// GetNumThreadLocalCachesOk returns a tuple with the NumThreadLocalCaches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetNumThreadLocalCachesOk() (*int32, bool) {
	if o == nil || o.NumThreadLocalCaches == nil {
		return nil, false
	}
	return o.NumThreadLocalCaches, true
}

// HasNumThreadLocalCaches returns a boolean if a field has been set.
func (o *AllocatorStats) HasNumThreadLocalCaches() bool {
	if o != nil && o.NumThreadLocalCaches != nil {
		return true
	}

	return false
}

// SetNumThreadLocalCaches gets a reference to the given int32 and assigns it to the NumThreadLocalCaches field.
func (o *AllocatorStats) SetNumThreadLocalCaches(v int32) {
	o.NumThreadLocalCaches = &v
}

// GetSmallCacheSize returns the SmallCacheSize field value if set, zero value otherwise.
func (o *AllocatorStats) GetSmallCacheSize() int32 {
	if o == nil || o.SmallCacheSize == nil {
		var ret int32
		return ret
	}
	return *o.SmallCacheSize
}

// GetSmallCacheSizeOk returns a tuple with the SmallCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetSmallCacheSizeOk() (*int32, bool) {
	if o == nil || o.SmallCacheSize == nil {
		return nil, false
	}
	return o.SmallCacheSize, true
}

// HasSmallCacheSize returns a boolean if a field has been set.
func (o *AllocatorStats) HasSmallCacheSize() bool {
	if o != nil && o.SmallCacheSize != nil {
		return true
	}

	return false
}

// SetSmallCacheSize gets a reference to the given int32 and assigns it to the SmallCacheSize field.
func (o *AllocatorStats) SetSmallCacheSize(v int32) {
	o.SmallCacheSize = &v
}

// GetTinyCacheSize returns the TinyCacheSize field value if set, zero value otherwise.
func (o *AllocatorStats) GetTinyCacheSize() int32 {
	if o == nil || o.TinyCacheSize == nil {
		var ret int32
		return ret
	}
	return *o.TinyCacheSize
}

// GetTinyCacheSizeOk returns a tuple with the TinyCacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatorStats) GetTinyCacheSizeOk() (*int32, bool) {
	if o == nil || o.TinyCacheSize == nil {
		return nil, false
	}
	return o.TinyCacheSize, true
}

// HasTinyCacheSize returns a boolean if a field has been set.
func (o *AllocatorStats) HasTinyCacheSize() bool {
	if o != nil && o.TinyCacheSize != nil {
		return true
	}

	return false
}

// SetTinyCacheSize gets a reference to the given int32 and assigns it to the TinyCacheSize field.
func (o *AllocatorStats) SetTinyCacheSize(v int32) {
	o.TinyCacheSize = &v
}

func (o AllocatorStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectArenas != nil {
		toSerialize["directArenas"] = o.DirectArenas
	}
	if o.HeapArenas != nil {
		toSerialize["heapArenas"] = o.HeapArenas
	}
	if o.NormalCacheSize != nil {
		toSerialize["normalCacheSize"] = o.NormalCacheSize
	}
	if o.NumDirectArenas != nil {
		toSerialize["numDirectArenas"] = o.NumDirectArenas
	}
	if o.NumHeapArenas != nil {
		toSerialize["numHeapArenas"] = o.NumHeapArenas
	}
	if o.NumThreadLocalCaches != nil {
		toSerialize["numThreadLocalCaches"] = o.NumThreadLocalCaches
	}
	if o.SmallCacheSize != nil {
		toSerialize["smallCacheSize"] = o.SmallCacheSize
	}
	if o.TinyCacheSize != nil {
		toSerialize["tinyCacheSize"] = o.TinyCacheSize
	}
	return json.Marshal(toSerialize)
}

type NullableAllocatorStats struct {
	value *AllocatorStats
	isSet bool
}

func (v NullableAllocatorStats) Get() *AllocatorStats {
	return v.value
}

func (v *NullableAllocatorStats) Set(val *AllocatorStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatorStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatorStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatorStats(val *AllocatorStats) *NullableAllocatorStats {
	return &NullableAllocatorStats{value: val, isSet: true}
}

func (v NullableAllocatorStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatorStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

